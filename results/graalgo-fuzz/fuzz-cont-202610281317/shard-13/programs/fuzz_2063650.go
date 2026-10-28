package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64; F1 string }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]string)
    var v3 int64 = 712
    v4 := make(map[string]uint32)
    var v5 float32 = 0.1
    v6 := GenS1{F0: -128}
    for fi7 := 0; fi7 < 19; fi7++ {
        v8 := GenS2{F0: (v3 - v3), F1: "hello"}
        _ = fi7
        _ = v8
    }
    switch 0 {
    case 320:
        if false {
            if false {
                v3 -= (v3 / ((v3) | 1))
                var v9 float32 = float32((float64((v5 + float32(float64(v5)))) / 1e10))
                _ = v9
            }
            var v10 int16 = 895
            v5 = v5
            _ = v10
        } else {
            v5 = 1e10
        }
    case -1:
        v11 := []int{32768, -545}
        _ = v11
    }
    if false {
        v3 -= (v3 & v3)
    } else {
        v3 += (v3 ^ (v3 | v3))
    }
    v3 = (^(v3))
    v5 += v5
    switch "foo bar" {
    case "\n\t":
        var v12 string = "foo bar"
        _ = v12
    }
    if true {
        v3 -= (v3 - -9223372036854775808)
    } else {
        v13 := []int{977, -106, 784}
        v5 = (float32(1.5) - v5)
        _ = v13
    }
    var v14 int32 = -2147483648
    var v15 uint32 = 809
    if ((v3 / ((v3) | 1)) != (-(v3))) {
        v16 := GenS1{F0: 42}
        var v17 uint16 = 65535
        _ = v16
        _ = v17
    } else {
        var v18 int8 = int8(v15)
        _ = v18
    }
    for fi19 := 0; fi19 < 5; fi19++ {
        v20 := []int{(fi19 - fi19), (fi19 - (fi19 - (fi19 << ((0) & 63)))), (fi19 << ((161) & 63)), fi19}
        v4["hello"] = v15
        _ = fi19
        _ = v20
    }
    if false {
        v15 = v15
        switch 699 {
        case -2147483648:
            v15 += v15
        case 42:
            for fi21 := 0; fi21 < 15; fi21++ {
                fmt.Println(v3, v15)
                v15 += v15
                _ = fi21
            }
        default:
            v22 := []int{}
            _ = v22
        }
    } else {
        for fi23 := 0; fi23 < 5; fi23++ {
            v5 += v5
            if (18446744073709551615 >= 494) {
                delete(v1, 938)
                v24 := []int{}
                switch fi23 {
                case 476:
                    var v25 string = "foo bar"
                    _ = v25
                default:
                    v3 = v3
                }
                _ = v24
            } else {
                v14 = v14
            }
            v3 = 9223372036854775807
            _ = fi23
        }
    }
    v26 := GenS1{F0: 0}
    v27 := []int{4294967295, 987, -666}
    v5 += 0.0
    v3 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
