package scripts

var GetAllScript = `import groovy.json.JsonBuilder
import groovy.json.JsonSlurper
import org.slf4j.Logger
import org.sonatype.nexus.repository.Repository
import org.sonatype.nexus.repository.maven.internal.Maven2Format
import org.sonatype.nexus.repository.types.GroupType
import org.sonatype.nexus.repository.types.HostedType
import org.sonatype.nexus.repository.types.ProxyType
import org.sonatype.nexus.script.plugin.internal.provisioning.RepositoryApiImpl

public class GetAll {
    private final Logger log
    private final RepositoryApiImpl repo

    GetAll(log, repo) {
        this.log = log
        this.repo = repo
    }

    String doStuff(String data) {

        def request = new JsonSlurper().parseText(data) as Request

        if (request.id != null) {
            if(repo.repositoryManager.exists(request.id)) {
                def repositoryMap = convert(repo.repositoryManager.get(request.id))
                return new JsonBuilder(repositoryMap)
            } else {
                return null
            }

        } else {

            def all = repo.repositoryManager.browse()
                    .toSorted { it.name }
                    .collect { Repository repository -> convert(repository) }
                    .findAll { it != null }

            new JsonBuilder(all)

        }
    }

    private Map convert(Repository repository) {

        if (repository.format == new Maven2Format()) {
            if (repository.type == new HostedType()) {
                [
                        name: repository.name,
                        type: "maven_hosted",
                        data: [] as Map
                ]
            } else if (repository.type == new ProxyType()) {
                def remoteUrl = repository.configuration.attributes("proxy").get("remoteUrl") as String
                [
                        name: repository.name,
                        type: "maven_proxy",
                        data: [
                                remoteUrl: remoteUrl
                        ] as Map
                ]
            } else if (repository.type == new GroupType()) {
                def members = repository.configuration.attributes("group").get("memberNames") as List<String>
                [
                        name: repository.name,
                        type: "maven_group",
                        data: [
                                members: new JsonBuilder(members).toString()
                        ] as Map
                ]
            } else {
                null
            }
        } else {
            null
        }

    }

    class Request {
        public String id;
    }


}

return new GetAll(log, repository).doStuff(args)

`


