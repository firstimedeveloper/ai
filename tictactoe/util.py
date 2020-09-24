class Node():
    def __init__(self, state, parent, action):
        self.state = state
        self.parent = parent
        self.action = action

class StackFrontier():
    def __init__(self):
        self.frontier = []
    
    def add(self, node):
        self.frontier.append(node)

    def contains_state(self, node):
        if node in self.frontier:
            return True
        else:
            return False

    def empty(self):
        if len(self.frontier) == 0:
            return True
    
    def remove(self):
        if self.empty():
            return False
        else: 
            node = self.frontier[-1]
            self.frontier = self.frontier[:-2]
        return node

class QueueFrontier(StackFrontier):
    def remove(self):
        if self.empty():
            return False
        else:
            node = self.frontier[0]
            self.frontier = self.frontier[1:]
        return node

