from tictactoe import initial_state
import tictactoe as ttt

EMPTY = None

board = initial_state()
board = [['X', EMPTY, EMPTY],
            [EMPTY, 'O', EMPTY],
            [EMPTY, EMPTY, EMPTY]]

player = ttt.player(board)

moves = ttt.actions(board)

newBoard = ttt.result(board, (0,1))

print(f'nextplayer: {player}\n')
print(f'possible moves: {moves}\n')
print(f'board: {board} newBoard: {newBoard}\n')

board2 = [['X', 'X', 'O'],
            ['O', EMPTY, 'X'],
            ["X", 'O', 'X']]

winner = ttt.winner(board2)

print(f'winner: {winner}\n')

gameOver = ttt.terminal(board2)

print(f'gameOver: {gameOver}\n')