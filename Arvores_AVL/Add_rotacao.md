# Revisar codigo de BST e seguir os seguintes passos
## *(0) Adicionar adicionar altura e fator de balanco no struct* [x]

(1) implementar funcao func (node *BSTNode) UpdatePropieties() [ ]
    ->atualizar altura de node
    ->atualizar fator de balanco de node

(2) Implementar funcoes de rotacao:
    func (node *BSTNode) RotRight() *BSTNode
    func (node *BSTNode) RotLeft() *BSTNode
    -> Ambas usam UpdateProprieties
        -> (1) Raiz antiga
        -> (2) Nova Raiz

(3) Implementar acoes de Rebalanceamento para cada caso: [ ]
    func (node *BSTNode) RebLeftLeft() *BSTNode // caso 1
    func (node *BSTNode) RebLeftRight() *BSTNode // caso 2
    func (node *BSTNode) RebRightRight() *BSTNode // caso 3
    func (node *BSTNode) RebRightLeft() *BSTNode // caso 4

(4) Implementar funcao que identifica os casos e chama a funcao de Rebalanceamento correta [ ]
    func (node *BSTNode) Rebalance() *BSTNode

(5) Adicionar chamada para funcoes Rebalance() e UpdateProprieties() nas funcoes Add() e Remove() [ ]
