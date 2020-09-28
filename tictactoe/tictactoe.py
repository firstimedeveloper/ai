"""
Tic Tac Toe Player
"""

import math
import copy
import util

X = "X"
O = "O"
EMPTY = None


def initial_state():
    """
    Returns starting state of the board.
    """
    return [[EMPTY, EMPTY, EMPTY],
            [EMPTY, EMPTY, EMPTY],
            [EMPTY, EMPTY, EMPTY]]


def player(board):
    """
    Returns player who has the next turn on a board.
    """
    x = 0
    o = 0
    for row in board:
        for box in row:
            if box == X:
                x += 1
            if box == O:
                o += 1
    
    if (x+o) % 2 == 0:
        return X
    else:
        return O


def actions(board):
    """
    Returns set of all possible actions (i, j) available on the board.
    """
    options = set()
    for i, row in enumerate(board):
        for j, box in enumerate(row):
            if box == EMPTY:
                options.add((i,j))
    return options


def result(board, action):
    """
    Returns the board that results from making move (i, j) on the board.
    """
    newBoard = copy.deepcopy(board)

    nextPlayer = player(newBoard)
    moves = actions(newBoard)
    
    if action in moves:
        newBoard[action[0]][action[1]] = nextPlayer
    else:
        raise ValueError("Invalid action.")

    return newBoard


def winner(board):
    """
    Returns the winner of the game, if there is one.
    """
    for i, row in enumerate(board):
        if row[0] == row[1] == row[2]:
            return row[0]
        if board[0][i] == board[1][i] == board[2][i]:
            return board[0][i]

    if board[0][0] == board[1][1] == board[2][2]:
        return board[0][0]
    if board[2][0] == board[1][1] == board [0][2]:
        return board[1][1]



    return None

def terminal(board):
    """
    Returns True if game is over, False otherwise.
    """
    if winner(board) != None:
        return True
    if len(actions(board)) == 0:
        return True
    return False

def utility(board):
    """
    Returns 1 if X has won the game, -1 if O has won, 0 otherwise.
    """
    if winner(board) == X:
        return 1
    elif winner(board) == O:
        return -1
    else:
        return 0


def minimax(board):
    """
    Returns the optimal action for the current player on the board.
    """

    def max(board):
        if terminal(board):
            return utility(board)
        v = -2
        for action in actions(board):
            a = min(result(board, action))
            # if a == 1:
            #     return a
            if v < a:
                v = a
        return v

    def min(board): 
        if terminal(board):
            return utility(board)
        v = 2 
        for action in actions(board):
            a = max(result(board, action))
            # if a == -1:
            #     return a
            if v > a:
                v = a
        return v

    coord = tuple
    p = player(board)
    if p == X: 
        v = -2
        for action in actions(board):
            a = max(result(board, action))
            if v < a:
                coord = action
    elif p == O:
        v = 2
        for action in actions(board):
            a = min(result(board, action))
            if v > a:
                coord = action
    return coord




        
        
