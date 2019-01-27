package scripts

var CreateProxyMavenScript = `import groovy.json.JsonSlurper
import org.slf4j.Logger
import org.sonatype.nexus.script.plugin.internal.provisioning.RepositoryApiImpl

public class CreateProxyMaven {

    private final Logger log;
    private final RepositoryApiImpl repo

    CreateProxyMaven(log, repo) {
        this.log = log
        this.repo = repo
    }

    String doStuff(String data) {

        def req = new JsonSlurper().parseText(data) as Request

        repo.createMavenProxy(req.name, req.remote)

    }

    private class Request {
        public String name;
        public String remote;
    }

}

return new CreateProxyMaven(log, repository).doStuff(args)

`


