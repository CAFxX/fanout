package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64; F1 uint32; F2 int }
func genFn0(p0 uint8, p1 int) int64 {
    v6 := GenS1{F0: 1000}
    v7 := GenS1{F0: 42}
    switch 128 {
    case 32767:
        for fi8 := 0; fi8 < 4; fi8++ {
            switch "a" {
            case "foo bar":
                p1 -= (fi8 & 0)
            case "a":
                fmt.Println(p0)
            default:
                var v9 uint = 255
                _ = v9
            }
            v10 := GenS1{F0: 32767}
            _ = fi8
            _ = v10
        }
    case 32768:
        var v11 uint16 = 1000
        _ = v11
    case -2147483648:
        v12 := []int{p1}
        _ = v12
    default:
        for fi13 := 0; fi13 < 3; fi13++ {
            v14 := []int{}
            _ = fi13
            _ = v14
        }
    }
    if false {
        var v15 bool = false
        var v16 bool = true
        _ = v15
        _ = v16
    }
    if (0 == 1) {
        switch (-(p1)) {
        case -609:
            for fi17 := 0; fi17 < 16; fi17++ {
                for fi18 := 0; fi18 < 4; fi18++ {
                    if false {
                        switch 84 {
                        case -2147483648:
                            p1 -= fi17
                        default:
                            v19 := GenS1{F0: 908}
                            var v20 int64 = 9223372036854775807
                            _ = v19
                            _ = v20
                        }
                        v21 := []int{p1, p1, 225}
                        _ = v21
                    }
                    v22 := []int{146}
                    p0 = 88
                    _ = fi18
                    _ = v22
                }
                _ = fi17
            }
        case -2147483648:
            var v23 int = (p1 | -9223372036854775808)
            _ = v23
        case 53:
            var v24 uint8 = p0
            _ = v24
        default:
            v25 := []int{p1, 32767}
            _ = v25
        }
        fmt.Println(p0, p0)
    } else {
        v26 := []int{(-(p1)), p1, 32768, -979}
        if (128 != (p0 - 42)) {
            if false {
                p1 = -2147483648
                p1 += 128
            } else {
                var v27 int64 = 9223372036854775807
                if len(v26) > 0 {
                    v26[0] = p1
                }
                _ = v27
            }
            fmt.Println(p1)
        }
        _ = v26
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -226
}

func main() {
    var v1 int16 = 0
    var v2 uint16 = 430
    var v3 int8 = -72
    v4 := make(map[string]uint32)
    v5 := make(map[uint32]uint32)
    v2 += (v2 & (v2 / ((v2) | 1)))
    if true {
        var v28 int16 = 42
        _ = v28
    } else {
        var v29 int8 = 42
        var v30 int8 = v3
        var v31 string = "foo bar"
        _ = v29
        _ = v30
        _ = v31
    }
    fmt.Println(v3, v3, v2)
    if (true || ((!false) || (1e10 <= -2.25))) {
        var v32 float64 = 3.14159
        var v33 uint16 = v2
        var v34 int = -876
        _ = v32
        _ = v33
        _ = v34
    } else {
        if (4294967295 > 1) {
            v1 = (v1 / (((v1 << ((2) & 63))) | 1))
        } else {
            v3 += v3
            switch 4294967295 {
            case 9223372036854775807:
                switch "foo bar" {
                case "\n\t":
                    delete(v5, 2147483647)
                }
            }
            v35 := GenS2{F0: v2}
            _ = v35
        }
    }
    v36 := []int{32768, 2147483647}
    v37 := GenS3{F0: 1, F1: 65535, F2: -148}
    v5[2147483648] = 1
    if true {
        v38 := GenS1{F0: (v1 << ((18446744073709551615) & 63))}
        v3 = v3
        _ = v38
    } else {
        var v39 int = 9223372036854775807
        v3 = v3
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
