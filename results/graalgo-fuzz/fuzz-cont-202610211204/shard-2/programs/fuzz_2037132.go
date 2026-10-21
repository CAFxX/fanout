package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 bool }
func genFn0(p0 int16) int {
    p0 = 1000
    for fi4 := 0; fi4 < 13; fi4++ {
        fmt.Println(p0, p0)
        v5 := []int{fi4}
        var v6 uint = 941
        _ = fi4
        _ = v5
        _ = v6
    }
    p0 = (p0 / ((p0) | 1))
    for fi7 := 0; fi7 < 13; fi7++ {
        if (1000 < p0) {
            var v8 int = (fi7 * fi7)
            _ = v8
        } else {
            p0 = p0
            var v9 uint8 = 154
            p0 = p0
            _ = v9
        }
        p0 += 81
        fmt.Println(p0, p0)
        _ = fi7
    }
    _ = p0
    return -801
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float32 = (0.0 + 0.0)
    var v3 int16 = -32768
    if (true && false) {
        fmt.Println(v3)
        v3 = v3
    } else {
        var v10 int32 = 42
        if (4294967295 == 4294967295) {
            for fi11 := 0; fi11 < 7; fi11++ {
                var v12 uint16 = 2
                _ = fi11
                _ = v12
            }
            if false {
                fmt.Println(v3, v10, v10)
                var v13 int = 2147483647
                v3 += (v3 / ((-1) | 1))
                _ = v13
            }
        }
        _ = v10
    }
    if (0 >= 42) {
        var v14 uint32 = 0
        var v15 int16 = v3
        _ = v14
        _ = v15
    } else {
        var v16 uint8 = 146
        var v17 float32 = (float32(((1e10 / (0.0 + -2.25)) * (((0.1 - 1e10) * -2.25) - 0.1))) * v2)
        _ = v16
        _ = v17
    }
    if (v3 < v3) {
        switch ("foo bar" + "") {
        case "":
            var v18 uint64 = 346
            _ = v18
        }
        for fi19 := 0; fi19 < 7; fi19++ {
            for fi20 := 0; fi20 < 1; fi20++ {
                v21 := []int{}
                fmt.Println(fi20, fi19)
                _ = fi20
                _ = v21
            }
            switch "a" {
            case "hello":
                v22 := GenS1{F0: 51, F1: -2147483648}
                _ = v22
            }
            _ = fi19
        }
    } else {
        for fi23 := 0; fi23 < 18; fi23++ {
            v2 += v2
            fmt.Println(v3)
            _ = fi23
        }
    }
    v1[903] = ((v2 < -2.25) && false)
    v24 := v1[707]
    v25 := GenS1{F0: 18446744073709551615, F1: -295}
    var v26 int32 = -519
    var v27 int32 = v26
    for fi28 := 0; fi28 < 8; fi28++ {
        switch (fi28 & -719) {
        case 127:
            switch 32767 {
            case -9223372036854775808:
                var v29 uint64 = 4294967295
                _ = v29
            case 4294967295:
                v2 -= v2
            default:
                var v30 uint = 910
                _ = v30
            }
        case -1:
            v31 := GenS1{F0: 1, F1: 2147483647}
            _ = v31
        default:
            v32 := GenS2{F0: v24}
            _ = v32
        }
        switch (^(fi28)) {
        case 128:
            for fi33 := 0; fi33 < 1; fi33++ {
                var v34 string = "foo bar"
                var v35 uint = 583
                _ = fi33
                _ = v34
                _ = v35
            }
        case 423:
            if v24 {
                var v36 int64 = 9223372036854775807
                _ = v36
            } else {
                switch genFn0(v3) {
                case 32767:
                    if (2147483647 > 929) {
                        var v37 uint8 = 128
                        v37 = 229
                        _ = v37
                    }
                case 742:
                    v2 += float32((float64(0.0) * 3.14159))
                case 9223372036854775807:
                    fmt.Println(v2, v2)
                }
                v3 -= (v3 + (^(v3)))
                for fi38 := 0; fi38 < 18; fi38++ {
                    v3 -= v3
                    _ = fi38
                }
            }
        default:
            for fi39 := 0; fi39 < 17; fi39++ {
                v40 := v1[42]
                var v41 uint8 = 2
                var v42 float32 = -2.25
                _ = fi39
                _ = v40
                _ = v41
                _ = v42
            }
        }
        _ = fi28
    }
    fmt.Println(v2, v26, v27)
    if true {
        var v43 uint16 = 472
        _ = v43
    }
    var v44 uint8 = 1
    v45 := GenS1{F0: 1, F1: (^(v26))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v44
    _ = v45
}
