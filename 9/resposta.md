 ## Resposta
 Levando em consideração o estudo de caso apresentado, podemos identificar algumas vulnerabilidades em potêncial, onde a maioria delas depende de detalhes de implementação da API.

 ### Contribuintes não precisam se autenticar
A maior falha do endpoint ao meu ver está na falta de obrigatoriedade de autenticação para os contribuintes. Ao obrigar que o contribuinte valide sua identidade, por exemplo através de um token, temos muito mais segurança na intenção e validade das requisições que estamos recebendo. Tendo em vista que um BOT não conseguiria reproduzir as mesmas requisições com sucesso.

 ### Riscos de SQL injection
 Outra vulnerabilidade encontrada é a possibilidade de o sistema sofrer SQL injections se o mesmo não tiver um tratamento adequado no backend, visto que os campos permitem que seja enviado um SQL razoavelmente grande e malicioso.

 ### Rate limit
 Não é citado em momento algum se o endpoint possui algum rate limit, o que é extremamente necessário levando em consideração que sem um rate limit, um usuário mal intencionado pode causar um ataque de Deny of service (DOS), sobrecarregando o servidor com requisições infinitas.