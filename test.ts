// Generic "ok" and "err" shapes
type Ok<T> = {
    type: 'ok'
    value: T
}

type Err<E> = {
    type: 'err'
    error: E
}

// Generic Result type, similar to Rust's Result<T, E>
type Result<T, E> = Ok<T> | Err<E>

// Constructors
function ok<T>(value: T): Ok<T> {
    return { type: 'ok', value }
}

function err<E>(error: E): Err<E> {
    return { type: 'err', error }
}

// Example usage
function parseIntResult(input: string | number): Result<number, string> {
    const n = Number(input)

    if (Number.isNaN(n)) {
        return err(`Invalid number: ${input}`)
    }
    return ok(n)
}

// Pattern matching style   
const r = parseIntResult(3.2)

if (r.type === 'ok') {
    console.log('value =', r.value)
} else {
    console.error('error =', r.error)
}
