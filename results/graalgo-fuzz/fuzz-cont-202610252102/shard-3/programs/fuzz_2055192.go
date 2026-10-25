package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 bool; F1 int; F2 int8 }
type GenS4 struct { F0 int32 }
func genFn0(p0 int, p1 int, p2 uint32) uint8 {
    v5 := GenS1{F0: 1}
    switch (p1 | p1) {
    case 2147483647:
        var v6 uint8 = 255
        _ = v6
    case 4294967295:
        if false {
            var v7 float32 = 0.1
            _ = v7
        } else {
            fmt.Println(p0)
            p2 -= 544
            if ((-1 >= 0) && false) {
                p0 -= -1
                if false {
                    var v8 uint16 = 373
                    v9 := GenS1{F0: 255}
                    p1 -= -461
                    _ = v8
                    _ = v9
                }
            }
        }
    case 32767:
        fmt.Println(p1, p2, p2)
    default:
        v10 := []int{(p0 & -1000), int(p2), (p1 | p1)}
        _ = v10
    }
    if (!false) {
        v11 := GenS2{F0: true, F1: (p1 << ((0) & 63))}
        fmt.Println(p1, p0)
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 118
}

func main() {
    var v1 float32 = 0.0
    var v2 float64 = (float64(1e10) / 0.1)
    v3 := make(map[int]int)
    v4 := make(map[int]uint32)
    var v12 uint32 = 4294967295
    v13 := GenS2{F0: true, F1: -1}
    v12 -= 65535
    if (false && (!(65535 <= (^(v12))))) {
        v12 = v12
        v12 = (v12 << ((4294967295) & 63))
    }
    v2 = v2
    fmt.Println(v12)
    v1 -= v1
    var v14 int32 = 2147483647
    v12 -= 639
    v15 := GenS3{F0: (!false), F1: 0, F2: -88}
    switch (("foo bar" + "") + "") {
    case "foo bar":
        switch 9223372036854775807 {
        case 32768:
            v16 := GenS4{F0: v14}
            _ = v16
        case 275:
            var v17 int32 = v14
            _ = v17
        default:
            if true {
                v18 := []int{-9223372036854775808, int(v14)}
                _ = v18
            } else {
                v19 := GenS4{F0: (^(v14))}
                var v20 bool = false
                for fi21 := 0; fi21 < 8; fi21++ {
                    v22 := GenS4{F0: v14}
                    v2 = v2
                    var v23 uint = 715
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                _ = v19
                _ = v20
            }
        }
    case "hello":
        var v24 int16 = 32767
        _ = v24
    }
    fmt.Println(v12, v1, v14)
    var v25 int32 = v14
    if ((!true) && false) {
        switch "a" {
        case "foo bar":
            if false {
                for fi26 := 0; fi26 < 11; fi26++ {
                    var v27 uint16 = 0
                    var v28 uint8 = genFn0(fi26, fi26, (v12 << ((65535) & 63)))
                    v1 += v1
                    _ = fi26
                    _ = v27
                    _ = v28
                }
            }
        }
    }
    var v29 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v4
}
