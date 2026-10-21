package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 int16, p1 int32) bool {
    v5 := GenS1{F0: 42}
    var v6 uint64 = 42
    fmt.Println(p1, v6)
    switch -2147483648 {
    case 4294967295:
        var v7 int16 = (p0 - (p0 | p0))
        _ = v7
    case 787:
        switch "\n\t" {
        case "hello":
            for fi8 := 0; fi8 < 13; fi8++ {
                switch 9223372036854775807 {
                case 9223372036854775807:
                    var v9 int32 = p1
                    _ = v9
                case 32768:
                    v10 := GenS1{F0: 0}
                    _ = v10
                }
                _ = fi8
            }
        }
    }
    var v11 bool = true
    _ = p0
    _ = p1
    _ = v11
    _ = v5
    _ = v6
    return v11
}
func genFn1(p0 uint8, p1 uint64) int16 {
    p1 += 839
    v12 := []int{128, 9223372036854775807, 32768}
    var v13 float32 = (1.5 / 1e10)
    if false {
        p0 = (p0 >> ((4294967295) & 63))
        var v14 uint32 = 65535
        switch -9223372036854775808 {
        case 195:
            var v15 float64 = 1e10
            _ = v15
        case -9223372036854775808:
            var v16 uint64 = p1
            _ = v16
        case 42:
            v17 := GenS1{F0: v14}
            _ = v17
        default:
            switch -831 {
            case 127:
                v14 -= 645
            default:
                switch ("\n\t" + "foo bar") {
                case "\n\t":
                    v13 = v13
                default:
                    if len(v12) >= 2 {
                        v12 = v12[1:len(v12)]
                    }
                }
            }
        }
        _ = v14
    }
    if true {
        var v18 uint = 734
        v18 += (^(v18))
        var v19 uint64 = (p1 * p1)
        _ = v18
        _ = v19
    }
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return 42
}

func main() {
    var v1 int16 = 563
    v2 := make(map[string]int)
    v3 := make(map[int]bool)
    var v4 uint64 = 524
    var v20 bool = true
    for fi21 := 0; fi21 < 6; fi21++ {
        fmt.Println(v20)
        v22 := GenS1{F0: 2}
        fmt.Println(v1, fi21)
        _ = fi21
        _ = v22
    }
    var v23 float64 = 1e10
    if (0 != v4) {
        v24 := []int{2147483647, -886, 0}
        switch 2147483647 {
        case -9223372036854775808:
            switch -799 {
            case -888:
                var v25 float64 = (v23 - 0.1)
                _ = v25
            }
        case -2147483648:
            for fi26 := 0; fi26 < 15; fi26++ {
                _ = len(v24)
                _ = fi26
            }
        case -893:
            v23 = -2.25
        }
        _ = v24
    }
    v23 += v23
    v27 := []int{1, -9223372036854775808, -2147483648, -9223372036854775808}
    v1 = (v1 % ((-1) | 1))
    var v28 uint64 = v4
    for fi29 := 0; fi29 < 16; fi29++ {
        var v30 int16 = v1
        _ = fi29
        _ = v30
    }
    v23 = float64(float32(v23))
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
