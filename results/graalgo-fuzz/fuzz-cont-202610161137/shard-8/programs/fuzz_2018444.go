package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64 }

func main() {
    v1 := []string{"", "", ("\n\t" + (("hello" + "a") + ("\n\t" + "\n\t"))), (("foo bar" + ("\n\t" + "a")) + ("a" + (("\n\t" + "hello") + "")))}
    var v2 bool = false
    var v3 uint = 4294967295
    var v4 uint8 = 246
    var v5 float32 = 0.1
    v2 = v2
    v2 = (775 == 871)
    var v6 float32 = v5
    v6 = (v5 + (((v6 * (3.14159 / 1e10)) / v5) / -2.25))
    fmt.Println(v2)
    switch 32767 {
    case -469:
        v3 -= v3
    case 126:
        switch -1 {
        case 0:
            _ = len(v1)
        case -9223372036854775808:
            if (42 >= 107) {
                v7 := GenS1{F0: false}
                _ = v7
            }
        case 42:
            var v8 int64 = -9223372036854775808
            _ = v8
        }
    default:
        var v9 int64 = -9223372036854775808
        _ = v9
    }
    var v10 uint32 = 4294967295
    v4 = v4
    var v11 uint16 = 65535
    fmt.Println(v6)
    var v12 float64 = 1e10
    v13 := GenS2{F0: 42}
    for fi14 := 0; fi14 < 19; fi14++ {
        var v15 bool = (5 > int8(fi14))
        var v16 float32 = 1e10
        _ = fi14
        _ = v15
        _ = v16
    }
    v6 = v5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
