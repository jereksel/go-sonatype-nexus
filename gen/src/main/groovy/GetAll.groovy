import groovy.json.JsonBuilder
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

    String doStuff() {

        def all = repo.repositoryManager.browse()
                .toSorted { it.name }
                .collect { Repository repository ->

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
        .findAll { it != null }

        return new JsonBuilder(all)
    }
}

// #return new GetAll(log, repository).doStuff()

// ###

class GetAllMain {

    static main(args) {
        Deployer.deploy("GetAll")
    }

}

