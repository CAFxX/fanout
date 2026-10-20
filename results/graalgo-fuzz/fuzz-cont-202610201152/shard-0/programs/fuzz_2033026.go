package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 uint16; F1 uint16 }
func genFn0(p0 uint64, p1 uint8) uint64 {
    var v6 float64 = 0.0
    if (false && ((!(1 == 551)) || false)) {
        p1 -= p1
        p0 += 679
        v7 := GenS1{F0: -1}
        _ = v7
    } else {
        if ("\n\t" <= (("a" + "foo bar") + "foo bar")) {
            var v8 float32 = (float32(v6) + float32(v6))
            if (!false) {
                var v9 uint = 42
                fmt.Println(p1, v6)
                var v10 uint16 = 560
                _ = v10
                _ = v9
            } else {
                var v11 float64 = v6
                v12 := []int{32768, 128}
                _ = v11
                _ = v12
            }
            _ = v8
        }
    }
    var v13 string = "hello"
    p0 += 0
    _ = p0
    _ = p1
    _ = v13
    _ = v6
    return p0
}
func genFn1(p0 int16, p1 uint8, p2 uint8) uint64 {
    fmt.Println(p1, p1)
    var v14 float64 = float64(-2.25)
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 511
}

func main() {
    v1 := make(map[int]string)
    var v2 float64 = float64(0.0)
    var v3 int64 = 0
    v4 := make(map[int]uint32)
    var v5 uint16 = 42
    v15 := GenS2{F0: 0, F1: 129}
    var v16 uint32 = 2
    var v17 int = 128
    v18 := v1[157]
    var v19 int = (v17 & v17)
    switch 9223372036854775807 {
    case 387:
        v20 := GenS3{F0: v5, F1: 2}
        _ = v20
    }
    for fi21 := 0; fi21 < 12; fi21++ {
        for fi22 := 0; fi22 < 19; fi22++ {
            fmt.Println(v5, v3, v3)
            switch ("hello" + v18) {
            case "\n\t":
                fmt.Println(v3, fi21)
            case "hello":
                v23 := GenS3{F0: v5, F1: 759}
                _ = v23
            }
            _ = fi22
        }
        if (((v18 + v18) + v18) >= ("" + v18)) {
            switch v19 {
            case -295:
                v19 = (v17 | 1)
            case 0:
                fmt.Println(v2)
            case -2147483648:
                switch (("\n\t" + ((v18 + "foo bar") + ("\n\t" + (v18 + ("\n\t" + v18))))) + "") {
                case "a":
                    v24 := []int{(v17 & (fi21 ^ (v17 * 4294967295))), v17, 4294967295}
                    _ = v24
                case "\n\t":
                    var v25 int32 = -2147483648
                    _ = v25
                default:
                    var v26 uint32 = v16
                    _ = v26
                }
            default:
                var v27 int16 = int16(v5)
                _ = v27
            }
        } else {
            switch fi21 {
            case -630:
                var v28 uint64 = 0
                _ = v28
            }
            var v29 bool = false
            fmt.Println(fi21)
            _ = v29
        }
        _ = fi21
    }
    var v30 int32 = 1
    fmt.Println(v17, v19)
    v31 := []int{(v19 | (v17 * v19)), (v19 + -9223372036854775808)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
