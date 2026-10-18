package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 uint32; F1 uint32 }
func genFn0(p0 uint32, p1 int32) int8 {
    v6 := []int{}
    _ = len(v6)
    var v7 int64 = -476
    var v8 int64 = (-(v7))
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return -1
}
func genFn1(p0 int16, p1 int8, p2 uint64) uint32 {
    p2 += (p2 | p2)
    v9 := GenS1{F0: 494, F1: p0}
    p1 += p1
    v10 := []int{}
    var v11 int8 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v9
    return 386
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]string)
    var v3 bool = (0 == -2147483648)
    var v4 string = ("a" + ("foo bar" + ""))
    var v5 int32 = 1
    for fi12 := 0; fi12 < 18; fi12++ {
        fmt.Println(v5, v4)
        if (0 <= 54) {
            switch 1 {
            case 127:
                for fi13 := 0; fi13 < 17; fi13++ {
                    switch ((v4 + v4) + v4) {
                    case "a":
                        var v14 float64 = 0.1
                        _ = v14
                    }
                    _ = fi13
                }
            }
            v3 = (v3 && (!(-32768 > 0)))
        }
        for fi15 := 0; fi15 < 19; fi15++ {
            var v16 uint = 381
            switch (fi15 * (fi15 | (fi15 & fi12))) {
            case -953:
                v5 += (v5 * -963)
            default:
                var v17 int32 = (^(v5))
                _ = v17
            }
            switch fi12 {
            case -82:
                fmt.Println(v4)
            default:
                v3 = v3
                v18 := GenS1{F0: 129, F1: -32768}
                v19 := GenS1{F0: 32767, F1: 177}
                fmt.Println(fi15, v16, fi12)
                v5 += (v5 << (((v16 << ((v16) & 63))) & 63))
                _ = v18
                _ = v19
            }
            v4 = v4
            v20 := []int{(fi12 - fi12), fi12}
            var v21 uint16 = 42
            v22 := GenS1{F0: 42, F1: 1}
            v23 := v1[(-(fi12))]
            v5 -= v5
            switch (^(fi12)) {
            case 32768:
                if v3 {
                    fmt.Println(v21, fi12, fi15)
                    var v24 int64 = 557
                    _ = v24
                }
            case 500:
                delete(v1, fi12)
                var v25 float32 = 0.1
                var v26 int64 = -786
                _ = v25
                _ = v26
            default:
                v27 := GenS2{F0: (v23 + 2147483648), F1: (v23 ^ v23)}
                _ = v27
            }
            var v28 uint8 = 2
            v4 = "\n\t"
            var v29 int64 = 0
            var v30 uint64 = 1
            _ = fi15
            _ = v16
            _ = v20
            _ = v21
            _ = v22
            _ = v23
            _ = v28
            _ = v29
            _ = v30
        }
        _ = fi12
    }
    for fi31 := 0; fi31 < 6; fi31++ {
        var v32 uint64 = 1
        fmt.Println(v32, v4, fi31)
        _ = fi31
        _ = v32
    }
    var v33 int8 = 1
    switch "hello" {
    case "foo bar":
        v34 := GenS1{F0: 1, F1: -405}
        _ = v34
    case "\n\t":
        switch ("" + ("a" + "hello")) {
        case "hello":
            delete(v1, 32767)
        case "\n\t":
            delete(v1, -66)
        case "":
            for fi35 := 0; fi35 < 13; fi35++ {
                var v36 string = "\n\t"
                _ = fi35
                _ = v36
            }
        }
    case "":
        v2["foo bar"] = "a"
    }
    fmt.Println(v5, v5, v33)
    fmt.Println(v4, v5, v33)
    v33 += 1
    var v37 int = 128
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v37
    _ = v4
    _ = v5
}
