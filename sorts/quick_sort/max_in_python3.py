def max(*args, **kwargs):
    """
    :param args: If one positional argument is provided, it should be an iterable.
                 If two or more positional arguments are provided, the largest of
                 the positional arguments is returned.
    :param kwargs: There are two optional keyword-only arguments.
                   The `key` argument specifies a one-argument ordering function like
                   that used for list.sort().
                   The `default` argument specifies an object to return if the provided
                   iterable is empty.
    :return: return The largest item in the iterable
    """
    if not args:
        raise TypeError('max expected 1 arguments, got 0')
    elif len(args) == 1:
        iterable = iter(args[0])
    else:
        iterable = args

    default = False
    key = lambda x: x
    for k, v in kwargs.items():
        if k == "key":
            key = v
        elif k == "default":
            default = True
            default_value = v
        else:
            raise TypeError(f"'{k}' is an invalid keyword argument for this function")
    if not callable(key):
        raise TypeError(f"'{type(v).__name__}' object is not callable")

    first = True
    for item in iterable:
        value = key(item)
        if first or max_value < value:
            max_value = value
            max_item = item
            first = False

    if first:
        if not default:
            raise ValueError("max() arg is an empty sequence")
        max_item = default_value

    return max_item
