from tictactoe import initial_state
import tictactoe as ttt

EMPTY = None

board = initial_state()
board = [['X', 'X', 'O'],
        [EMPTY, 'O', EMPTY],
        [EMPTY, EMPTY, EMPTY]]

player = ttt.player(board)

moves = ttt.actions(board)

print(f'board: {board}')
while True:
    if ttt.terminal(board):
        break
    optimal = ttt.minimax(board)
    print(optimal)
    board = ttt.result(board, optimal)
    print(f'board: {board}')

# print(ttt.utility(board))


