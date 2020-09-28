from tictactoe import initial_state
import tictactoe as ttt

EMPTY = None

board = initial_state()
board = [['X', 'X', EMPTY],
            [EMPTY, 'O', EMPTY],
            [EMPTY, EMPTY, EMPTY]]

player = ttt.player(board)

moves = ttt.actions(board)


print(f'nextplayer: {player}\n')
print(f'possible moves: {moves}\n')

board2 = [['O', 'X', 'O'],
            [EMPTY, EMPTY, 'X'],
            ["X", EMPTY, 'X']]

winner = ttt.utility(board2)

print(f'winner: {winner}\n')
