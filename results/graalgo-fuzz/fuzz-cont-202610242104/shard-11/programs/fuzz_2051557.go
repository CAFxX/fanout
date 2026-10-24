package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool; F1 uint32 }
type GenS4 struct { F0 int16; F1 uint8; F2 bool }
func genFn0(p0 int64) string {
    if ((!true) && (!(42 < 681))) {
        v4 := []int{42}
        _ = v4
    }
    v5 := []int{409}
    var v6 uint64 = 941
    _ = p0
    _ = v5
    _ = v6
    return "foo bar"
}
func genFn1(p0 uint, p1 uint16, p2 int64) uint16 {
    for fi7 := 0; fi7 < 12; fi7++ {
        p2 += p2
        for fi8 := 0; fi8 < 15; fi8++ {
            var v9 uint = (p0 ^ (p0 / ((0) | 1)))
            v10 := GenS1{F0: (^(p1))}
            if (!false) {
                if false {
                    fmt.Println(p0, fi8)
                } else {
                    v11 := []int{}
                    if true {
                        v12 := GenS1{F0: (p1 * 850)}
                        _ = v12
                    }
                    _ = v11
                }
                var v13 uint16 = 815
                _ = v13
            }
            _ = fi8
            _ = v10
            _ = v9
        }
        v14 := GenS2{F0: (p2 & (-(p2)))}
        _ = fi7
        _ = v14
    }
    if (32767 != 761) {
        var v15 uint64 = 347
        for fi16 := 0; fi16 < 14; fi16++ {
            v15 += v15
            var v17 string = "a"
            v18 := GenS1{F0: p1}
            _ = fi16
            _ = v17
            _ = v18
        }
        fmt.Println(p2, p1, p1)
        _ = v15
    } else {
        fmt.Println(p2, p1, p2)
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := make(map[int]uint32)
    var v3 uint64 = 715
    switch -958 {
    case 240:
        v19 := v2[-2147483648]
        _ = v19
    }
    var v20 bool = (("" + "\n\t") == "\n\t")
    var v21 int64 = v1
    var v22 int = 4294967295
    for fi23 := 0; fi23 < 12; fi23++ {
        v24 := GenS1{F0: 255}
        v3 -= 4294967295
        _ = fi23
        _ = v24
    }
    var v25 uint64 = (v3 << ((2) & 63))
    if (-49 > -2147483648) {
        var v26 string = ((("hello" + "foo bar") + "a") + ("foo bar" + "hello"))
        var v27 string = ("a" + (v26 + "\n\t"))
        fmt.Println(v3, v25, v22)
        _ = v26
        _ = v27
    } else {
        fmt.Println(v22, v22, v3)
        fmt.Println(v3, v3, v22)
    }
    v25 -= (^(v25))
    switch (("\n\t" + "foo bar") + "hello") {
    case "\n\t":
        if v20 {
            var v28 int8 = -1
            v28 -= (^(v28))
            v25 = 923
            _ = v28
        } else {
            v29 := []int{v22, 1, 128}
            v30 := GenS3{F0: false, F1: 118}
            _ = v29
            _ = v30
        }
    case "a":
        var v31 float64 = float64(0.1)
        _ = v31
    case "foo bar":
        v32 := GenS4{F0: int16(v22), F1: 2, F2: (((float64(1e10) * (-2.25 * (((0.0 - 1e10) * (0.0 * float64(3.14159))) * 0.1))) - -2.25) > float64(0.0))}
        _ = v32
    default:
        var v33 uint64 = (^(v25))
        _ = v33
    }
    fmt.Println(v1, v21, v25)
    delete(v2, v22)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v25
    _ = v3
}
