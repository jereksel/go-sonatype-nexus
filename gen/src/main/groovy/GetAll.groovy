import groovy.json.JsonBuilder
import org.slf4j.Logger
import org.sonatype.nexus.repository.Repository
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
            [
                    name  : repository.name,
                    format: repository.format.value,
                    type  : repository.type.value
            ]
        }

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

