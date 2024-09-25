import json
import random
import argparse

def generate_config(mode, lower, upper, size):
    config_data = []

    if mode == "client_streaming":
        for _ in range(size):
            number = "{{randomInt " + str(lower) + " " + str(upper) + "}}"
            config_data.append({"number": number})
    
    elif mode == "bi_streaming":
        for _ in range(size):
            real_rows = [{"val": "{{randomInt " + str(lower) + " " + str(upper) +"}}"} for _ in range(size)]
            config_data.append({"rows": real_rows})

    return config_data

def write_config_to_file(config_data, filename="data.json"):
    with open(filename, 'w') as json_file:
        json.dump(config_data, json_file, indent=4)

def main():
    parser = argparse.ArgumentParser(description="Generate config.json based on mode.")
    
    parser.add_argument("mode", type=str, choices=["client_streaming", "bi_streaming"], 
                        help="The mode of configuration, either 'client_streaming' or 'bi_streaming'.")
    parser.add_argument("lower", type=int, help="The lower limit for random number generation.")
    parser.add_argument("upper", type=int, help="The upper limit for random number generation.")
    parser.add_argument("size", type=int, help="The number of entries to generate.")
    parser.add_argument("path", type=str, help="File path.", default="data.json")
    args = parser.parse_args()

    config_data = generate_config(args.mode, args.lower, args.upper, args.size)
    write_config_to_file(config_data,args.path)

    print(f"Config file generated successfully with mode {args.mode}.")

if __name__ == "__main__":
    main()

