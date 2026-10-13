package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 int16 = 866
    v2 := make(map[uint32]string)
    var v3 uint = 0
    var v4 uint8 = 128
    var v5 string = (("a" + ((("foo bar" + (("a" + "") + ("" + ("foo bar" + "foo bar")))) + (("hello" + ("a" + ("hello" + ("a" + "foo bar")))) + "hello")) + "a")) + "")
    var v6 int = 0
    v6 = (v6 * (v6 ^ v6))
    for fi7 := 0; fi7 < 10; fi7++ {
        fmt.Println(v6, v5, v1)
        var v8 int8 = 124
        _ = fi7
        _ = v8
    }
    v2[363] = (v5 + v5)
    var v9 uint8 = 2
    v5 = ("\n\t" + "hello")
    if (v5 < "foo bar") {
        var v10 uint32 = 931
        _ = v10
    }
    switch 224 {
    case 128:
        delete(v2, 1)
    case -2147483648:
        v4 -= v4
    default:
        for fi11 := 0; fi11 < 2; fi11++ {
            v3 -= 65535
            var v12 uint64 = 2
            _ = fi11
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        fmt.Println(v6, v4)
        v14 := GenS1{F0: 1}
        v5 = "\n\t"
        _ = fi13
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
