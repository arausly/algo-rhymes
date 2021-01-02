def counting_valleys(n, path):
    valley_count, depth_tracker, sea_level = 0, 0, 0
    paths_arr = list(path)
    for path in paths_arr:
        if path == "U":
            sea_level += 1
        else:
            sea_level -= 1

        if sea_level == -1:
            depth_tracker = sea_level

        if depth_tracker == -1 and sea_level == 0:
            valley_count += 1
            depth_tracker = 0

    return valley_count


if __name__ == "__main__":
    pass
