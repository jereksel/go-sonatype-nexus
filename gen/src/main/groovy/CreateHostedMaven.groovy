import groovy.json.JsonSlurper
import org.slf4j.Logger
import org.sonatype.nexus.script.plugin.internal.provisioning.RepositoryApiImpl

public class CreateHostedMaven {
    private final Logger log;
    private final RepositoryApiImpl repo

    CreateHostedMaven(log, repo) {
        this.log = log
        this.repo = repo
    }

    String doStuff(String data) {

        def req = new JsonSlurper().parseText(data) as Request

        repo.createMavenHosted(req.name)

    }

    private class Request {
        public String name;
    }

}

// #return new CreateHostedMaven(log, repository).doStuff(args)

// ###

class CreateHostedMavenMain {

    static main(args) {
        Deployer.deploy("Create")
    }

}
