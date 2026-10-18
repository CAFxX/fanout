package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 string; F1 string; F2 uint }
type GenS3 struct { F0 uint8; F1 uint64; F2 uint8 }
type GenS4 struct { F0 string; F1 bool; F2 uint64 }
type GenS5 struct { F0 uint64 }
func genFn0(p0 uint64) int {
    v7 := GenS1{F0: 42, F1: 150}
    var v8 uint16 = 65535
    fmt.Println(p0, p0)
    switch -1 {
    case 4294967295:
        v9 := GenS2{F0: "a", F1: "hello", F2: 65535}
        _ = v9
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        for fi11 := 0; fi11 < 3; fi11++ {
            if true {
                p0 -= p0
                v8 = (v8 >> ((2) & 63))
                p0 = 182
            } else {
                p0 -= p0
            }
            _ = fi11
        }
        _ = fi10
    }
    _ = p0
    _ = v7
    _ = v8
    return -2147483648
}

func main() {
    var v1 int32 = 42
    var v2 int64 = 42
    var v3 float32 = -2.25
    v4 := []uint32{514, 2147483648}
    v5 := make(map[string]int)
    var v6 uint16 = 433
    v3 += (0.1 * float32(3.14159))
    _ = len(v5)
    if true {
        var v12 int32 = -270
        for fi13 := 0; fi13 < 19; fi13++ {
            var v14 uint16 = (v6 % ((v6) | 1))
            _ = fi13
            _ = v14
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v12
    }
    if true {
        if false {
            for fi15 := 0; fi15 < 20; fi15++ {
                v16 := GenS3{F0: 42, F1: 42, F2: 1}
                _ = fi15
                _ = v16
            }
            var v17 bool = false
            _ = v17
        }
    } else {
        fmt.Println(v6, v3, v2)
    }
    v18 := GenS3{F0: 2, F1: 18446744073709551615, F2: 2}
    v2 -= (-(v2))
    v1 = (^(v1))
    switch "" {
    case "foo bar":
        switch ("foo bar" + ("a" + "hello")) {
        case "":
            fmt.Println(v3, v2, v2)
        }
    case "\n\t":
        var v19 uint64 = 18446744073709551615
        _ = v19
    case "hello":
        fmt.Println(v3)
    default:
        v20 := GenS4{F0: "", F1: false, F2: 0}
        _ = v20
    }
    if false {
        var v21 uint16 = v6
        var v22 float64 = ((-2.25 * (0.0 + -2.25)) - 1e10)
        _ = v21
        _ = v22
    } else {
        var v23 uint16 = v6
        v3 = (0.0 + 1.5)
        if (!false) {
            if (234 <= 918) {
                var v24 int = -478
                delete(v5, "\n\t")
                _ = v24
            }
        }
        _ = v23
    }
    v25 := GenS5{F0: 42}
    if true {
        if true {
            v2 -= 425
            if len(v4) > 0 {
                v4[0] = 65535
            }
            var v26 uint32 = 446
            _ = v26
        } else {
            switch ("foo bar" + ((("\n\t" + ("\n\t" + "foo bar")) + "a") + "a")) {
            case "\n\t":
                v2 += v2
            }
        }
    } else {
        var v27 int = -585
        fmt.Println(v3, v1)
        _ = v27
    }
    delete(v5, "\n\t")
    v1 -= v1
    for fi28 := 0; fi28 < 11; fi28++ {
        v2 = (v2 * v2)
        _ = fi28
    }
    v2 = -1
    v1 = 429
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
