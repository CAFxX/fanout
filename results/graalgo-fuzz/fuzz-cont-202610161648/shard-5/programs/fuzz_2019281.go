package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 bool; F1 bool; F2 int8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool; F1 uint8 }
func genFn0(p0 uint16, p1 int16) uint32 {
    v7 := GenS1{F0: 1, F1: p0}
    var v8 int64 = 23
    var v9 int8 = -128
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 int8 = -52
    v2 := []int{}
    var v3 int8 = 19
    var v4 int8 = (v1 ^ 42)
    var v5 uint64 = 90
    var v6 int32 = -27
    switch 32767 {
    case 4294967295:
        fmt.Println(v3, v4)
    case -1:
        switch 1 {
        case 128:
            var v10 float32 = 3.14159
            _ = v10
        case 554:
            _ = len(v2)
        case -421:
            var v11 string = "a"
            _ = v11
        }
    default:
        for fi12 := 0; fi12 < 11; fi12++ {
            v3 -= -1
            var v13 int = int(v5)
            v5 = v5
            _ = fi12
            _ = v13
        }
    }
    v3 = v3
    if ((v5 << ((525) & 63)) == v5) {
        fmt.Println(v1, v1, v6)
        var v14 int32 = -2147483648
        switch 9223372036854775807 {
        case 740:
            v15 := GenS2{F0: (true || true), F1: false, F2: (v4 << ((255) & 63))}
            _ = v15
        case -9223372036854775808:
            if true {
                var v16 int16 = 42
                _ = v16
            }
        case 40:
            v17 := GenS2{F0: (!("foo bar" > "\n\t")), F1: true, F2: (v4 << ((953) & 63))}
            _ = v17
        default:
            v4 = (v4 * v4)
        }
        _ = v14
    } else {
        var v18 int16 = -1
        var v19 string = "foo bar"
        _ = v18
        _ = v19
    }
    if true {
        v4 += 127
        if true {
            v6 = (v6 << ((uint(v5)) & 63))
            fmt.Println(v3, v5, v5)
        } else {
            v6 += (v6 >> ((255) & 63))
            var v20 uint32 = 2147483647
            v21 := GenS1{F0: 781, F1: 700}
            _ = v20
            _ = v21
        }
        for fi22 := 0; fi22 < 7; fi22++ {
            for fi23 := 0; fi23 < 15; fi23++ {
                v3 += v3
                v4 += v1
                _ = fi23
            }
            _ = fi22
        }
    }
    v24 := GenS3{F0: (v4 - v1)}
    var v25 int8 = v1
    v6 = (v6 | (v6 ^ (^(v6))))
    _ = len(v2)
    switch ("hello" + "") {
    case "\n\t":
        v26 := GenS4{F0: (!((32767 == 1) || ((((("" + ("\n\t" + "foo bar")) + "\n\t") + "a") != (("foo bar" + "foo bar") + "\n\t")) || true))), F1: 255}
        _ = v26
    case "foo bar":
        switch -18 {
        case -692:
            v4 += 127
        case 127:
            v27 := GenS1{F0: 539, F1: 626}
            _ = v27
        }
    case "":
        v25 = (^(v4))
    default:
        fmt.Println(v25, v1, v3)
    }
    v28 := GenS3{F0: v1}
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
