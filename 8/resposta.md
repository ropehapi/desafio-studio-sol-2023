## Resposta
Levando as premissas da questão em consideração, elaborei um DER bem simples para ser aplicado no domínio de uma biblioteca.

![DER da biblioteca](DER%20biblioteca.png)

O diagrama entidade-relacional (DER) acima tem como domínio um sistema de empréstimo de livros em uma biblioteca. Claro que um sistema completo não teria apenas essas entidades, portanto, vamos tratar o modelo acima como um serviço responsável apenas pelo empréstimo de livros.

Para esse domínio precisamos de 7 entidades, o próprio empréstimo, o leitor que emprestou o livro, o funcionário que registrou o empréstimo, e o livro, que deve obrigatoriamente conter uma categoria, um autor e uma editora.