
import flask

# ======================================================================

def factors(n):
    """Yield the factors of a given number, including duplicates."""
    p = 2
    while n > 1:
        # We're working up, so n % p == 0 implies p is prime.
        while n % p == 0:
            n /= p
            yield p
        p += 1

# ----------------------------------------------------------------------

def get_factors(n):
    if not isinstance(n, int) or n < 1:
        flask.abort(404, "Give a positive integer, not " + repr(n))
    return list( factors(n) )
