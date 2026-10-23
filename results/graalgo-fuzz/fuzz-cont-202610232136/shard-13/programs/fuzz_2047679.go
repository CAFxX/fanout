package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint; F1 int32; F2 uint64 }
type GenS4 struct { F0 uint8; F1 uint64 }
type GenS5 struct { F0 int16; F1 uint32 }
func genFn0(p0 string, p1 uint32, p2 int) bool {
    v5 := GenS1{F0: (^(p2)), F1: -128, F2: 42}
    switch (p2 % ((p2) | 1)) {
    case 42:
        v6 := []int{p2, 1}
        _ = v6
    default:
        p2 -= 128
    }
    if ((p0 + "foo bar") == "\n\t") {
        for fi7 := 0; fi7 < 6; fi7++ {
            p0 = p0
            v8 := []int{-996, (-(p2)), (-(fi7))}
            p2 = -733
            _ = fi7
            _ = v8
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return (662 >= 1)
}

func main() {
    v1 := make(map[int]string)
    var v2 uint64 = 164
    var v3 uint32 = 322
    var v4 int = -757
    if (!false) {
        v4 -= 9223372036854775807
    } else {
        for fi9 := 0; fi9 < 16; fi9++ {
            v3 = (v3 >> ((1) & 63))
            if ((947 != v2) && false) {
                var v10 int16 = 1000
                _ = v10
            }
            _ = fi9
        }
    }
    v11 := GenS2{F0: (^(v4))}
    fmt.Println(v3, v4)
    var v12 int8 = 127
    switch (^(v4)) {
    case 9223372036854775807:
        var v13 uint = uint(v2)
        _ = v13
    case 2:
        v14 := GenS2{F0: 128}
        _ = v14
    default:
        var v15 uint32 = (v3 & v3)
        _ = v15
    }
    if true {
        v16 := GenS3{F0: 635, F1: 1, F2: v2}
        _ = v16
    }
    v17 := GenS4{F0: 42, F1: (v2 ^ (v2 - (^(v2))))}
    v18 := GenS2{F0: v4}
    if (!true) {
        for fi19 := 0; fi19 < 6; fi19++ {
            var v20 float64 = 1.5
            v21 := []int{500, 127, -2147483648, -2147483648}
            _ = fi19
            _ = v20
            _ = v21
        }
    }
    if (true && false) {
        v22 := GenS5{F0: 0, F1: v3}
        _ = v22
    }
    v23 := v1[(v4 | (v4 << ((105) & 63)))]
    v24 := GenS2{F0: (v4 & (v4 << ((42) & 63)))}
    switch (("foo bar" + v23) + "") {
    case "foo bar":
        fmt.Println(v12, v23)
    case "hello":
        fmt.Println(v2, v3, v4)
    case "":
        v25 := GenS1{F0: (v4 - (v4 / ((-860) | 1))), F1: (v12 ^ (v12 + v12)), F2: -813}
        _ = v25
    }
    var v26 int64 = -980
    for fi27 := 0; fi27 < 16; fi27++ {
        if (421 >= -32768) {
            var v28 uint64 = v2
            var v29 uint16 = 42
            _ = v28
            _ = v29
        }
        _ = fi27
    }
    v3 = v3
    fmt.Println(v23)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v26
    _ = v3
    _ = v4
}
