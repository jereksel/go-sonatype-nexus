package scripts

var CreateGroupMavenScript = `import groovy.json.JsonSlurper
import org.slf4j.Logger
import org.sonatype.nexus.script.plugin.internal.provisioning.RepositoryApiImpl

public class CreateGroupMaven {

    private final Logger log;
    private final RepositoryApiImpl repo

    CreateGroupMaven(log, repo) {
        this.log = log
        this.repo = repo
    }

    String doStuff(String data) {

        def req = new JsonSlurper().parseText(data) as Request

        repo.createMavenGroup(req.name, req.members)

    }

    private class Request {
        public String name
        public List<String> members
    }

}

return new CreateGroupMaven(log, repository).doStuff(args)

`


