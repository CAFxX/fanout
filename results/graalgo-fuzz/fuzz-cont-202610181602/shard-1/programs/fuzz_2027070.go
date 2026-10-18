package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16; F1 int16; F2 uint16 }
type GenS4 struct { F0 int32; F1 uint32 }
func genFn0() int8 {
    var v7 int16 = -343
    v8 := GenS1{F0: "foo bar", F1: 733}
    if (0 < 113) {
        v7 += (v7 + (v7 / ((42) | 1)))
        switch 501 {
        case -492:
            var v9 uint = 128
            _ = v9
        case -762:
            v7 = (v7 * v7)
        default:
            v7 = (v7 * (-(v7)))
        }
        v10 := GenS1{F0: "a", F1: -719}
        _ = v10
    } else {
        var v11 uint16 = 0
        v11 -= v11
        _ = v11
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        var v13 int32 = -1
        _ = fi12
        _ = v13
    }
    var v14 uint64 = 42
    _ = v14
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 uint16, p1 string, p2 uint64) uint64 {
    for fi15 := 0; fi15 < 19; fi15++ {
        p0 = p0
        var v16 uint32 = 0
        _ = fi15
        _ = v16
    }
    p0 -= p0
    var v17 float64 = -2.25
    for fi18 := 0; fi18 < 8; fi18++ {
        fmt.Println(p1, p1, v17)
        p2 += (^(p2))
        _ = fi18
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return p2
}

func main() {
    var v1 uint = 126
    var v2 uint16 = 65535
    var v3 uint = (v1 - v1)
    var v4 int = 42
    var v5 float32 = 1.5
    var v6 int32 = 0
    if (3.14159 >= (0.0 + 1e10)) {
        var v19 uint32 = 62
        v20 := []int{v4, 4294967295, int(v19), int(v6)}
        _ = v19
        _ = v20
    }
    fmt.Println(v6, v4)
    v5 = (v5 + (v5 + v5))
    v21 := []int{(^(v4)), (v4 & v4), v4}
    _ = len(v21)
    v2 += (^(v2))
    var v22 string = "\n\t"
    switch (v4 ^ v4) {
    case 32767:
        var v23 uint8 = 149
        _ = v23
    default:
        v1 += 436
    }
    for fi24 := 0; fi24 < 1; fi24++ {
        fmt.Println(v4, v22)
        if true {
            switch v22 {
            case "hello":
                var v25 uint8 = 1
                _ = v25
            case "\n\t":
                if (0 <= 0) {
                    v21 = append(v21, fi24)
                    var v26 uint16 = v2
                    _ = v26
                }
            case "foo bar":
                if len(v21) > 0 {
                    v21[0] = (v4 * v4)
                }
            }
            var v27 uint16 = v2
            v21 = append(v21, (fi24 * 1))
            _ = v27
        }
        _ = fi24
    }
    _ = len(v21)
    fmt.Println(v4, v4, v2)
    var v28 float64 = (((3.14159 / 1.5) + 1e10) - 0.1)
    v29 := GenS2{F0: -808}
    if (!true) {
        if true {
            v30 := GenS3{F0: (v2 >> ((384) & 63)), F1: -408, F2: (v2 * v2)}
            _ = v30
        }
    } else {
        fmt.Println(v5, v2, v2)
        v31 := GenS4{F0: (v6 & 224), F1: 553}
        if false {
            v2 = (^(v2))
        } else {
            v32 := GenS1{F0: ((v22 + v22) + v22), F1: (v4 | 84)}
            if true {
                var v33 uint64 = genFn1((v2 % ((v2) | 1)), ((("foo bar" + "a") + (("" + (("hello" + ("foo bar" + (v22 + ((v22 + "hello") + v22)))) + v22)) + "hello")) + ""), 0)
                _ = v33
            }
            v5 = v5
            _ = v32
        }
        _ = v31
    }
    if false {
        v22 = (("a" + v22) + v22)
        v34 := GenS1{F0: "hello", F1: (v4 * 2147483647)}
        var v35 uint64 = 4294967295
        _ = v34
        _ = v35
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
