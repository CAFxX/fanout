package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 uint64 }
type GenS2 struct { F0 uint32; F1 int }
func genFn0(p0 int64, p1 int8) uint64 {
    for fi7 := 0; fi7 < 20; fi7++ {
        p1 = (p1 ^ (p1 | p1))
        switch "a" {
        case "foo bar":
            v8 := GenS1{F0: 42, F1: 2, F2: 217}
            _ = v8
        case "a":
            v9 := []int{}
            _ = v9
        case "\n\t":
            fmt.Println(fi7, p1)
        default:
            v10 := []int{fi7, -987}
            _ = v10
        }
        switch (^(fi7)) {
        case -206:
            v11 := GenS2{F0: 236, F1: 42}
            _ = v11
        default:
            if true {
                v12 := GenS1{F0: 198, F1: 0, F2: 18446744073709551615}
                _ = v12
            } else {
                var v13 int = (^(fi7))
                _ = v13
            }
        }
        _ = fi7
    }
    p0 += 736
    v14 := GenS1{F0: 28, F1: 1, F2: 214}
    _ = p0
    _ = p1
    _ = v14
    return 4294967295
}

func main() {
    v1 := []uint32{65535, 377, 257, 41}
    var v2 float32 = 0.0
    var v3 int64 = -180
    var v4 float64 = -2.25
    v5 := make(map[uint32]string)
    var v6 int = -2147483648
    switch ("hello" + "hello") {
    case "foo bar":
        v4 = (float64(0.0) + -2.25)
    default:
        var v15 string = (("a" + ("a" + "a")) + ("foo bar" + "a"))
        _ = v15
    }
    var v16 uint16 = 590
    delete(v5, 2147483647)
    var v17 uint16 = (v16 - v16)
    var v18 int = v6
    for fi19 := 0; fi19 < 17; fi19++ {
        v20 := v5[2]
        var v21 uint32 = 2147483647
        var v22 float64 = v4
        _ = fi19
        _ = v20
        _ = v21
        _ = v22
    }
    var v23 uint64 = 42
    v24 := GenS1{F0: 255, F1: 4294967295, F2: 2}
    if (!(("a" + "foo bar") >= "\n\t")) {
        fmt.Println(v23, v4, v4)
        if len(v1) > 0 {
            v1[0] = 2147483648
        }
    } else {
        if len(v1) > 0 {
            v1[0] = 537
        }
        if (2 == 500) {
            for fi25 := 0; fi25 < 13; fi25++ {
                var v26 int8 = 127
                var v27 int32 = 1
                _ = fi25
                _ = v26
                _ = v27
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
