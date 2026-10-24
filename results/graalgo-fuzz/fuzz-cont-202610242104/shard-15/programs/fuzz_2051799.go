package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 uint16) int {
    var v4 uint16 = (^(p0))
    v4 += 32767
    var v5 uint8 = 193
    _ = p0
    _ = v4
    _ = v5
    return 451
}

func main() {
    var v1 int = 42
    var v2 uint64 = 432
    v3 := make(map[int]bool)
    v2 -= (v2 / ((0) | 1))
    if true {
        fmt.Println(v1)
    }
    var v6 uint32 = 2147483648
    v1 = 32768
    v1 += 4294967295
    var v7 int = (v1 | 4294967295)
    v2 = (v2 + (^(v2)))
    var v8 bool = (127 != 127)
    fmt.Println(v1, v6)
    if v8 {
        switch -2147483648 {
        case 140:
            v1 = -559
        case 42:
            fmt.Println(v6)
        case 128:
            if (628 != v6) {
                var v9 uint8 = 42
                delete(v3, 32767)
                _ = v9
            }
        default:
            for fi10 := 0; fi10 < 5; fi10++ {
                var v11 uint32 = v6
                v12 := GenS1{F0: ("\n\t" + ((("foo bar" + (("\n\t" + "foo bar") + (("a" + (("foo bar" + "\n\t") + ("foo bar" + "a"))) + ("" + "hello")))) + "") + ("\n\t" + (("" + "\n\t") + "a"))))}
                _ = fi10
                _ = v11
                _ = v12
            }
        }
        var v13 int8 = 100
        var v14 bool = false
        _ = v13
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
