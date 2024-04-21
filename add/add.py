
def main():
    print("add of X AND Y")
    try:
        X = int(input("Enter X: "))
        Y = int(input("Enter Y: "))
    except ValueError as e:
        print(f"Error reading input: {e}")
        return

    # Calculating and displaying the result
    print(f"X + Y = {X} + {Y} = {X + Y}")

if __name__ == "__main__":
    main()