package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 int }
func genFn0(p0 uint32, p1 int, p2 bool) uint8 {
    switch p1 {
    case -1:
        if false {
            for fi7 := 0; fi7 < 10; fi7++ {
                if (!p2) {
                    var v8 int64 = 42
                    var v9 uint16 = 191
                    switch (fi7 >> ((1) & 63)) {
                    case 32768:
                        fmt.Println(p0)
                    }
                    _ = v8
                    _ = v9
                }
                var v10 uint16 = 532
                v11 := GenS1{F0: (100 > -1), F1: -1}
                _ = fi7
                _ = v10
                _ = v11
            }
            switch "foo bar" {
            case "foo bar":
                v12 := []int{128, (p1 << ((42) & 63)), p1, (-(p1))}
                _ = v12
            default:
                var v13 uint32 = 65535
                _ = v13
            }
        }
    case 2147483647:
        p0 += p0
    case 660:
        v14 := GenS2{F0: 1, F1: 4294967295}
        _ = v14
    }
    var v15 int32 = 0
    v16 := GenS2{F0: 127, F1: 255}
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    return 255
}

func main() {
    v1 := []string{("\n\t" + "\n\t"), ((("hello" + "") + ((("\n\t" + ("a" + "")) + (("a" + "foo bar") + "")) + "hello")) + ("a" + (((((("foo bar" + "a") + "") + "foo bar") + "") + "hello") + "foo bar"))), "foo bar", (("hello" + "hello") + "foo bar")}
    v2 := make(map[string]int)
    var v3 float32 = (0.0 / 1.5)
    v4 := make(map[uint32]bool)
    v5 := make(map[uint32]int)
    v6 := make(map[int]string)
    var v17 uint64 = 18446744073709551615
    v18 := v6[0]
    v19 := v2[(((("" + "") + v18) + v18) + ("hello" + (((("foo bar" + ("" + v18)) + (v18 + ("hello" + "foo bar"))) + "hello") + v18)))]
    var v20 uint16 = uint16(v17)
    v18 = v18
    v19 = (v19 >> ((255) & 63))
    fmt.Println(v3)
    fmt.Println(v19, v17)
    v19 = (v19 * (v19 & 4294967295))
    var v21 int = v19
    var v22 int32 = -2147483648
    for fi23 := 0; fi23 < 4; fi23++ {
        if false {
            var v24 float64 = 1e10
            v20 += (v20 << ((675) & 63))
            _ = v24
        } else {
            switch 42 {
            case 468:
                if false {
                    var v25 int16 = 42
                    v26 := GenS3{F0: (v19 ^ (^(v21)))}
                    var v27 uint64 = (^(v17))
                    _ = v25
                    _ = v26
                    _ = v27
                }
            }
            v3 -= v3
            _ = len(v1)
        }
        _ = fi23
    }
    switch (v21 << ((0) & 63)) {
    case 32767:
        var v28 uint = 765
        _ = v28
    case 0:
        v22 = (-(v22))
    default:
        fmt.Println(v3, v18)
    }
    v1 = append(v1, (v18 + ""))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
