class Pet:
    def speak(self):
        print("...", end='')

    def speak_to(self, name):
        self.speak()
        print(name)


class Dog(Pet):
    def speak(self):
        print("Wang!", end='')


aDog = Dog()
aDog.speak()
aDog.speak_to("Chao")