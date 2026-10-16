package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 int64, p1 int, p2 uint) uint16 {
    switch 2147483647 {
    case -9223372036854775808:
        switch "hello" {
        case "hello":
            var v4 int8 = 1
            _ = v4
        case "foo bar":
            fmt.Println(p0)
        case "a":
            p2 -= (p2 | 1)
        default:
            for fi5 := 0; fi5 < 9; fi5++ {
                p1 += -9223372036854775808
                _ = fi5
            }
        }
    case 4294967295:
        v6 := GenS1{F0: 0}
        _ = v6
    case -501:
        v7 := []int{(p1 << ((p2) & 63)), p1}
        _ = v7
    }
    for fi8 := 0; fi8 < 16; fi8++ {
        p1 = -2147483648
        _ = fi8
    }
    var v9 uint = p2
    if false {
        if (!(0 < 1)) {
            var v10 int64 = -1
            v11 := []int{(p1 * (p1 / ((128) | 1))), 1, (p1 << (((p2 | (v9 ^ v9))) & 63))}
            _ = v10
            _ = v11
        }
        var v12 uint16 = 0
        _ = v12
    }
    var v13 uint16 = 32767
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v9
    return v13
}
func genFn1(p0 int64, p1 uint) uint16 {
    v14 := GenS1{F0: 80}
    p1 = p1
    p0 += 9223372036854775807
    for fi15 := 0; fi15 < 8; fi15++ {
        var v16 float64 = (0.1 * ((((((1.5 * (float64((3.14159 / 1.5)) + 0.0)) / (1.5 + -2.25)) + 1.5) + 1e10) - float64((0.1 * (-2.25 + 1e10)))) * (0.1 + (1.5 - (float64(3.14159) - (0.1 / 1e10))))))
        _ = fi15
        _ = v16
    }
    v17 := []int{}
    _ = p0
    _ = p1
    _ = v14
    _ = v17
    return 255
}

func main() {
    var v1 bool = true
    var v2 int32 = -2147483648
    var v3 int32 = v2
    var v18 uint16 = 32767
    v2 -= v2
    v19 := GenS2{F0: 42}
    v18 += 2
    var v20 int64 = -266
    v21 := []int{0, 32768}
    var v22 int32 = v2
    v21 = append(v21, -1)
    fmt.Println(v3)
    fmt.Println(v22)
    for fi23 := 0; fi23 < 4; fi23++ {
        var v24 uint = 4294967295
        v25 := GenS2{F0: 4294967295}
        for fi26 := 0; fi26 < 18; fi26++ {
            switch (fi23 | fi23) {
            case -1:
                var v27 uint8 = 2
                _ = v27
            default:
                switch (-(fi23)) {
                case 2147483647:
                    v22 = (v2 ^ (v22 + (v3 | int32(v20))))
                default:
                    fmt.Println(fi23)
                }
            }
            switch ((("foo bar" + "\n\t") + "a") + "foo bar") {
            case "hello":
                if ((((false || (629 < fi26)) && (!(v22 < -539))) && v1) && (!true)) {
                    if true {
                        var v28 uint8 = 42
                        var v29 uint16 = v18
                        _ = v28
                        _ = v29
                    } else {
                        fmt.Println(v18, v24)
                        fmt.Println(v18)
                    }
                }
                v21 = append(v21, fi23)
            case "a":
                var v30 float64 = (3.14159 - 3.14159)
                _ = v30
            case "\n\t":
                var v31 bool = true
                _ = v31
            }
            v18 = v18
            _ = fi26
        }
        _ = fi23
        _ = v24
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
