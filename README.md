# easynix
An experimental Lisp to Nix transpiler

The [Nix](https://nixos.org) package manager is an absolutely amazing package manager
that brings a significant shift to how you manage software on your computer.
It makes you write all the software you want in a `*.nix` file in the Nix language.

The Nix language is purely functional language which is completely different from your
traditional procedural or object-oriented languages. Hence it can feel weird for newcomers
starting new on Nix or NixOS.

easynix is an attempt to write a transpiler that takes a more friendly language
like Lisp and convert it to a Nix file. The exact dialect of Lisp that easynix will
work with is still undecided, although Scheme is currently the top-pick.

easynix is highly experimental and practically of no use write now.

# Roadmap
- [ ] Decide the dialect of Lisp to be used
- [ ] Finish the lexer
- [ ] Work on the parser
- [ ] Code generation backend
