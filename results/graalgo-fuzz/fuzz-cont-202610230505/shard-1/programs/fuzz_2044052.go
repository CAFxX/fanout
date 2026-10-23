package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16; F1 bool; F2 uint64 }
func genFn0() int32 {
    v7 := []int{41, 4294967295, -1}
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 int16 = 1
        var v10 int = (-(fi8))
        for fi11 := 0; fi11 < 8; fi11++ {
            if (-102 != -38) {
                _ = len(v7)
                if (1 >= 102) {
                    v12 := GenS1{F0: v9, F1: 32767, F2: 2}
                    v13 := GenS1{F0: 695, F1: 65535, F2: 2}
                    _ = v12
                    _ = v13
                }
                v9 = 1000
                v10 -= (fi11 / ((-83) | 1))
            }
            v14 := GenS1{F0: v9, F1: 705, F2: 464}
            _ = fi11
            _ = v14
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    if false {
        for fi15 := 0; fi15 < 10; fi15++ {
            if false {
                v16 := GenS1{F0: 1000, F1: 42, F2: 0}
                _ = v16
            } else {
                var v17 int8 = -72
                _ = v17
            }
            var v18 int32 = 2147483647
            v18 += 1
            _ = fi15
            _ = v18
        }
        switch 32767 {
        case 127:
            var v19 uint = 186
            _ = v19
        case 859:
            var v20 int = 0
            _ = v20
        case 746:
            v21 := GenS1{F0: 32767, F1: 0, F2: 18446744073709551615}
            _ = v21
        }
        var v22 int64 = 0
        _ = v22
    }
    _ = v7
    return -2147483648
}

func main() {
    v1 := []string{""}
    v2 := make(map[uint32]int)
    var v3 bool = ((-1 == 127) && false)
    var v4 float64 = (0.0 * 0.0)
    v5 := []uint32{2147483648, 2147483648, 1}
    v6 := make(map[uint32]uint32)
    fmt.Println(v3, v3)
    if (32767 == 763) {
        v2[899] = 9223372036854775807
    }
    var v23 uint = 18446744073709551615
    v24 := GenS1{F0: -32768, F1: 1, F2: 2}
    var v25 float64 = (((v4 * 0.1) - v4) * v4)
    if false {
        switch "a" {
        case "":
            for fi26 := 0; fi26 < 18; fi26++ {
                fmt.Println(v3)
                var v27 string = ("\n\t" + "")
                _ = fi26
                _ = v27
            }
        case "a":
            if ("" < "a") {
                v2[228] = 2147483647
                var v28 int32 = -2147483648
                v3 = true
                _ = v28
            } else {
                fmt.Println(v25, v23, v3)
                var v29 uint = 65535
                _ = v29
            }
        }
    }
    v25 -= (float64(0.1) - (1.5 * ((float64(0.1) - v25) - -2.25)))
    fmt.Println(v23, v3)
    if (true && (true && true)) {
        v6[2147483648] = 42
        v30 := GenS2{F0: 4294967295}
        if true {
            var v31 int32 = 512
            for fi32 := 0; fi32 < 12; fi32++ {
                v31 -= 1
                _ = fi32
            }
            v3 = (216 < 42)
            _ = v31
        } else {
            switch "a" {
            case "a":
                v33 := GenS3{F0: -32768, F1: ((!(1e10 <= 3.14159)) || (true || true)), F2: 895}
                _ = v33
            case "":
                for fi34 := 0; fi34 < 3; fi34++ {
                    _ = len(v1)
                    v25 = (v25 / v4)
                    _ = fi34
                }
            }
        }
        _ = v30
    } else {
        fmt.Println(v25, v3)
        v4 = v25
    }
    fmt.Println(v3)
    var v35 bool = ((42 < 32767) || v3)
    var v36 int16 = -1
    var v37 int = -466
    var v38 int8 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
