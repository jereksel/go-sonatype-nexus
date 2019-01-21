import groovy.json.JsonBuilder
import groovy.json.JsonSlurper
import org.slf4j.Logger
import org.sonatype.nexus.script.plugin.internal.provisioning.RepositoryApiImpl

public class Remove {

    private final Logger log
    private final RepositoryApiImpl repo

    Remove(log, repo) {
        this.log = log
        this.repo = repo
    }

    String doStuff(String request) {

        def repositoryName = (new JsonSlurper().parseText(request) as Request).name

        def result

        if (!repo.repositoryManager.exists(repositoryName)) {
            result = [status: false]
        } else {
            repo.repositoryManager.delete(repositoryName)
            result = [status: true]
        }

        return new JsonBuilder(result)

    }

    private class Request {
        public String name;
    }

}


// #return new Remove(log, repository).doStuff(args)

// ###

class RemoveMain {

    static main(args) {
        Deployer.deploy("Remove")
    }

}

