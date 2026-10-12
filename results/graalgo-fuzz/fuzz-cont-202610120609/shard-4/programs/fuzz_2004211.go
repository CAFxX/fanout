package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 int32; F2 int64 }
func genFn0(p0 string) bool {
    v6 := []int{-9223372036854775808}
    for fi7 := 0; fi7 < 3; fi7++ {
        p0 = "foo bar"
        var v8 uint16 = 0
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = v6
    return false
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]string)
    v4 := make(map[string]int)
    var v5 float64 = -2.25
    for fi9 := 0; fi9 < 17; fi9++ {
        var v10 uint32 = 65535
        v11 := v3[(v10 << ((18446744073709551615) & 63))]
        _ = fi9
        _ = v10
        _ = v11
    }
    v5 += v5
    v12 := GenS1{F0: 32767}
    v5 = v5
    if false {
        v5 = v5
    } else {
        fmt.Println(v5)
        var v13 uint32 = 65535
        switch 2147483647 {
        case 634:
            switch "" {
            case "hello":
                var v14 float64 = (v5 / ((v5 / 0.1) * 0.1))
                _ = v14
            case "\n\t":
                switch "a" {
                case "foo bar":
                    var v15 int = 4294967295
                    _ = v15
                default:
                    if false {
                        var v16 bool = (0 == 364)
                        switch 42 {
                        case 1:
                            fmt.Println(v13, v13, v16)
                        default:
                            var v17 float32 = 3.14159
                            _ = v17
                        }
                        v16 = v16
                        _ = v16
                    }
                }
            default:
                v18 := []int{}
                _ = v18
            }
        case 127:
            var v19 int32 = 947
            _ = v19
        }
        _ = v13
    }
    v20 := GenS2{F0: 1, F1: -75, F2: 0}
    v5 -= 3.14159
    if true {
        var v21 uint16 = 43
        var v22 int = 128
        var v23 uint32 = 271
        _ = v21
        _ = v22
        _ = v23
    } else {
        v5 += ((0.0 + (((float64(1.5) / (3.14159 / 1.5)) * v5) / v5)) - (v5 * 0.0))
        for fi24 := 0; fi24 < 1; fi24++ {
            v25 := []int{(fi24 + 42), fi24}
            fmt.Println(v5, v5)
            _ = fi24
            _ = v25
        }
        v26 := GenS1{F0: 604}
        _ = v26
    }
    var v27 int8 = 0
    v27 = v27
    var v28 bool = (255 <= 65535)
    v5 -= v5
    if (-2.25 < float32(v5)) {
        var v29 float32 = (-2.25 - 0.0)
        _ = v29
    } else {
        if v28 {
            v4["hello"] = 9223372036854775807
        } else {
            v30 := v2[("" + ("" + "foo bar"))]
            if (false && v28) {
                var v31 uint16 = 65535
                for fi32 := 0; fi32 < 15; fi32++ {
                    v27 -= int8(v31)
                    _ = fi32
                }
                _ = v31
            }
            for fi33 := 0; fi33 < 8; fi33++ {
                var v34 string = ("hello" + "a")
                var v35 int16 = 1
                if false {
                    v27 = v27
                    var v36 bool = false
                    _ = v36
                } else {
                    v34 = v34
                    switch v34 {
                    case "hello":
                        for fi37 := 0; fi37 < 12; fi37++ {
                            v30 = 2147483648
                            fmt.Println(v34, v5)
                            _ = fi37
                        }
                    }
                }
                _ = fi33
                _ = v34
                _ = v35
            }
            _ = v30
        }
        v5 = ((float64(0.1) - ((((1e10 * v5) / (float64(float32(float64(((-2.25 + 0.0) - 0.1)))) - v5)) / 1e10) / 3.14159)) + ((v5 - (1.5 / (v5 - ((v5 - v5) + (float64((0.0 * 1e10)) + 0.1))))) * (v5 + 0.0)))
    }
    if true {
        if (false || v28) {
            for fi38 := 0; fi38 < 16; fi38++ {
                v39 := []int{fi38}
                _ = fi38
                _ = v39
            }
        }
        var v40 uint32 = 2147483647
        v27 += (^(v27))
        _ = v40
    } else {
        if (!(0 == 0)) {
            fmt.Println(v28, v28, v5)
            var v41 uint = 42
            var v42 float64 = v5
            _ = v41
            _ = v42
        }
    }
    v27 -= v27
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
