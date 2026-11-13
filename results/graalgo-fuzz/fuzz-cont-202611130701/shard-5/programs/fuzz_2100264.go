package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64 }

func main() {
    var v1 bool = (!true)
    var v2 float64 = 1e10
    v3 := make(map[int]bool)
    v4 := make(map[uint32]uint32)
    var v5 int16 = 1
    v6 := []string{"a"}
    switch ("hello" + ("foo bar" + ("a" + ("" + "\n\t")))) {
    case "\n\t":
        v2 = (1.5 + ((1.5 + float64(0.1)) / (float64((3.14159 - 1.5)) / 1e10)))
    default:
        v2 += v2
    }
    fmt.Println(v2, v2)
    if (96 < int8(v5)) {
        for fi7 := 0; fi7 < 1; fi7++ {
            if len(v6) > 0 {
                v6[0] = "hello"
            }
            _ = fi7
        }
        var v8 uint = 810
        var v9 int = -2147483648
        _ = v8
        _ = v9
    } else {
        v2 += v2
        v10 := GenS1{F0: 268}
        _ = v10
    }
    fmt.Println(v1)
    v1 = true
    var v11 float32 = -2.25
    if true {
        var v12 string = "\n\t"
        v4[2] = 0
        v12 = v12
        _ = v12
    } else {
        v5 -= v5
        var v13 int16 = v5
        v14 := GenS2{F0: -139}
        _ = v13
        _ = v14
    }
    var v15 uint64 = 602
    v15 -= 1
    if len(v6) > 0 {
        v6[0] = (("hello" + ("" + "a")) + "\n\t")
    }
    var v16 uint64 = v15
    var v17 int64 = 864
    fmt.Println(v11, v2)
    var v18 uint32 = 42
    switch "" {
    case "a":
        for fi19 := 0; fi19 < 20; fi19++ {
            if len(v6) > 0 {
                v6[0] = ("hello" + ("foo bar" + "a"))
            }
            v20 := GenS1{F0: (^(v18))}
            _ = fi19
            _ = v20
        }
    case "foo bar":
        var v21 int = 2147483647
        _ = v21
    default:
        var v22 int16 = v5
        _ = v22
    }
    v3[4294967295] = v1
    v23 := GenS1{F0: v18}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
