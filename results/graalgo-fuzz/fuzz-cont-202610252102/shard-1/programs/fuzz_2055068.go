package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int8; F1 int32; F2 uint16 }
type GenS5 struct { F0 uint8 }
func genFn0(p0 int64, p1 uint32) int64 {
    switch 127 {
    case 9223372036854775807:
        var v5 uint = 580
        _ = v5
    default:
        for fi6 := 0; fi6 < 19; fi6++ {
            fmt.Println(fi6, p1)
            _ = fi6
        }
    }
    v7 := []int{1, -1}
    if ((0.1 + ((1.5 * 0.1) + (((float32(float64((((-2.25 * float32(0.0)) * (0.0 + 3.14159)) * -2.25))) / 3.14159) + 1.5) - (1.5 + (1e10 / ((1.5 * 1.5) + 0.0)))))) < 1.5) {
        for fi8 := 0; fi8 < 7; fi8++ {
            _ = len(v7)
            _ = fi8
        }
        v9 := GenS1{F0: -1}
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    return (p0 ^ p0)
}

func main() {
    var v1 int8 = 55
    var v2 float32 = 3.14159
    v3 := make(map[int]bool)
    var v4 int8 = (v1 & -75)
    v1 = -63
    if false {
        if false {
            v10 := []int{127, 2147483647, 2147483647, -758}
            _ = v10
        }
        fmt.Println(v1, v1)
    }
    v2 = v2
    v11 := GenS1{F0: 42}
    for fi12 := 0; fi12 < 11; fi12++ {
        var v13 int32 = -2147483648
        if (((255 > 18446744073709551615) && (42 > v13)) || (0 > 1)) {
            v14 := GenS1{F0: 1000}
            switch -2147483648 {
            case -66:
                v2 = float32((-2.25 - 1e10))
            case 861:
                fmt.Println(v13, v13)
            case -1:
                for fi15 := 0; fi15 < 12; fi15++ {
                    var v16 uint16 = 255
                    v17 := GenS2{F0: 18446744073709551615}
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            }
            var v18 uint8 = 68
            _ = v14
            _ = v18
        }
        for fi19 := 0; fi19 < 14; fi19++ {
            if false {
                v20 := GenS1{F0: 42}
                switch (fi19 - fi12) {
                case 128:
                    var v21 uint64 = 42
                    _ = v21
                case 4294967295:
                    if false {
                        switch (fi12 & -2147483648) {
                        case 0:
                            var v22 int = (fi19 % ((648) | 1))
                            _ = v22
                        }
                    } else {
                        v13 = (-(v13))
                    }
                }
                _ = v20
            }
            if (0.0 >= (((1e10 - (1.5 - 1e10)) - (-2.25 * ((-2.25 / 1e10) - 0.0))) - (-2.25 - (3.14159 - float64(v2))))) {
                v2 -= v2
            }
            _ = fi19
        }
        _ = fi12
        _ = v13
    }
    if (v4 == (v4 + -1)) {
        if true {
            v4 = v1
            for fi23 := 0; fi23 < 16; fi23++ {
                var v24 uint = 1
                v25 := GenS3{F0: 121}
                _ = fi23
                _ = v24
                _ = v25
            }
            if false {
                v26 := GenS1{F0: 0}
                _ = v26
            } else {
                var v27 int8 = (v4 >> ((4294967295) & 63))
                for fi28 := 0; fi28 < 17; fi28++ {
                    if (3.14159 != -2.25) {
                        v29 := GenS4{F0: v4, F1: -2147483648, F2: 65535}
                        _ = v29
                    }
                    _ = fi28
                }
                _ = v27
            }
        } else {
            var v30 float32 = float32(1e10)
            fmt.Println(v30, v4)
            _ = v30
        }
    } else {
        for fi31 := 0; fi31 < 18; fi31++ {
            var v32 int64 = 42
            v4 = v4
            fmt.Println(v2, v1, fi31)
            _ = fi31
            _ = v32
        }
        fmt.Println(v1)
    }
    v33 := GenS4{F0: (^(v1)), F1: -1, F2: 781}
    if false {
        if (true && (194 > 128)) {
            if false {
                if (false || true) {
                    v1 += 42
                    v34 := GenS5{F0: 127}
                    var v35 uint = 255
                    _ = v34
                    _ = v35
                }
                v36 := GenS4{F0: (v1 * -31), F1: 2147483647, F2: 50}
                _ = v36
            }
            var v37 int64 = 1
            _ = v37
        } else {
            v38 := GenS1{F0: 0}
            fmt.Println(v1, v4, v4)
            fmt.Println(v1, v4)
            _ = v38
        }
    } else {
        v2 = v2
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v33
    _ = v4
}
