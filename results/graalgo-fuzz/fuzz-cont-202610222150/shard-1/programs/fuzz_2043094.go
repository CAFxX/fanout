package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint }
type GenS2 struct { F0 int16 }
func genFn0() int16 {
    var v7 uint8 = 127
    v7 += (v7 * (v7 / ((v7) | 1)))
    if (!false) {
        v8 := GenS1{F0: 564, F1: ("" + "\n\t"), F2: 1}
        _ = v8
    }
    _ = v7
    return 42
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]int)
    var v3 int32 = -949
    var v4 uint64 = 319
    var v5 int32 = v3
    v6 := []uint32{4294967295, 937}
    var v9 bool = false
    if false {
        fmt.Println(v9)
    } else {
        var v10 uint32 = 2147483647
        if v9 {
            v5 -= -2147483648
            v4 = (^(v4))
            v10 = 2147483648
        }
        _ = v10
    }
    fmt.Println(v4, v4)
    delete(v2, (("" + ("foo bar" + ("hello" + "hello"))) + (("a" + ("" + (("" + "foo bar") + ("\n\t" + ("hello" + ("hello" + ("" + ""))))))) + "")))
    v2["a"] = -451
    v4 = v4
    var v11 uint = 743
    v3 += (v3 << ((892) & 63))
    v4 -= 2
    for fi12 := 0; fi12 < 3; fi12++ {
        v13 := GenS1{F0: 1, F1: ("a" + "foo bar"), F2: uint(v4)}
        fmt.Println(v11, v9)
        _ = fi12
        _ = v13
    }
    var v14 int = 855
    v4 += (v4 / (((v4 >> ((4294967295) & 63))) | 1))
    var v15 uint8 = 127
    var v16 int32 = (-(v3))
    v17 := GenS2{F0: -269}
    v18 := GenS2{F0: 32767}
    switch "foo bar" {
    case "a":
        v19 := GenS1{F0: 201, F1: ("hello" + ("hello" + "\n\t")), F2: (v11 + 0)}
        _ = v19
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
