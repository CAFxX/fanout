package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16; F1 uint32; F2 int8 }
type GenS4 struct { F0 uint8; F1 int64; F2 uint64 }
type GenS5 struct { F0 int16; F1 int8; F2 uint8 }
type GenS6 struct { F0 uint32; F1 uint8 }
func genFn0() int32 {
    v5 := GenS1{F0: 255, F1: 50}
    var v6 uint8 = 170
    fmt.Println(v6)
    _ = v5
    _ = v6
    return 0
}

func main() {
    var v1 int = 393
    var v2 uint = 695
    v3 := make(map[string]int)
    var v4 uint8 = 1
    var v7 int16 = 1000
    for fi8 := 0; fi8 < 8; fi8++ {
        if true {
            delete(v3, ((("\n\t" + "\n\t") + "a") + "a"))
            v2 = v2
        }
        _ = fi8
    }
    var v9 int8 = 1
    v10 := []int{(v1 * v1), v1, v1, v1}
    fmt.Println(v1)
    for fi11 := 0; fi11 < 13; fi11++ {
        var v12 bool = (!true)
        _ = fi11
        _ = v12
    }
    v7 -= (^(v7))
    fmt.Println(v1, v9, v2)
    var v13 bool = false
    if v13 {
        for fi14 := 0; fi14 < 14; fi14++ {
            v2 += v2
            v4 = (v4 | (v4 / (((v4 + v4)) | 1)))
            if ("" > "foo bar") {
                var v15 int = (-(fi14))
                _ = v15
            }
            _ = fi14
        }
        var v16 string = "foo bar"
        _ = v16
    } else {
        v13 = ((3.14159 - (3.14159 / (float64(((3.14159 / 0.1) * 1.5)) - ((float64((((1.5 + (1.5 - -2.25)) - (1.5 - 0.0)) - -2.25)) / (float64(1.5) / 1e10)) - ((0.1 + float64(-2.25)) - float64((float32((float64((((((0.0 + ((3.14159 + float32(0.1)) + 0.1)) + 3.14159) / 3.14159) - (1.5 / 3.14159)) * float32((-2.25 + (1e10 - -2.25))))) / 1.5)) - 1e10))))))) > (0.0 + (-2.25 - 0.0)))
        fmt.Println(v2)
        if (2147483648 == 2147483647) {
            var v17 int8 = (v9 / ((101) | 1))
            var v18 uint16 = 65535
            v19 := GenS2{F0: (!false)}
            _ = v17
            _ = v18
            _ = v19
        } else {
            var v20 bool = v13
            v7 += (v7 * (v7 - (v7 - (v7 * (v7 << ((v2) & 63))))))
            _ = v20
        }
    }
    var v21 float64 = 3.14159
    v10 = append(v10, v1)
    v2 = (v2 + v2)
    v22 := GenS3{F0: v7, F1: uint32(v4), F2: (v9 % ((int8(v1)) | 1))}
    for fi23 := 0; fi23 < 2; fi23++ {
        for fi24 := 0; fi24 < 3; fi24++ {
            v25 := GenS1{F0: 128, F1: v4}
            _ = fi24
            _ = v25
        }
        if len(v10) > 0 {
            v10[0] = (v1 * (v1 >> ((v2) & 63)))
        }
        if (!v13) {
            if len(v10) > 0 {
                v10[0] = (-(v1))
            }
            for fi26 := 0; fi26 < 13; fi26++ {
                if true {
                    var v27 uint16 = 2
                    var v28 bool = true
                    _ = v27
                    _ = v28
                }
                _ = fi26
            }
        }
        _ = fi23
    }
    if (!((-(v9)) > (-(v9)))) {
        for fi29 := 0; fi29 < 19; fi29++ {
            for fi30 := 0; fi30 < 2; fi30++ {
                fmt.Println(v2)
                switch v1 {
                case -9223372036854775808:
                    v31 := GenS4{F0: (v4 << ((121) & 63)), F1: 1, F2: 42}
                    _ = v31
                case 2147483647:
                    if ((0 < -1) || ((v2 >> ((v2) & 63)) == (v2 * (v2 - v2)))) {
                        v32 := GenS5{F0: (-(v7)), F1: (v9 << ((v2) & 63)), F2: (v4 & uint8(v2))}
                        _ = v32
                    } else {
                        fmt.Println(v2, v13)
                        v10 = append(v10, v1)
                        v33 := GenS5{F0: v7, F1: 13, F2: (v4 ^ (v4 * (^(v4))))}
                        _ = v33
                    }
                case -2147483648:
                    v1 = fi30
                    fmt.Println(fi29)
                    v9 = 0
                }
                _ = len(v3)
                _ = fi30
            }
            v34 := GenS6{F0: 0, F1: 127}
            var v35 int64 = -121
            _ = fi29
            _ = v34
            _ = v35
        }
    }
    v2 -= 65535
    for fi36 := 0; fi36 < 9; fi36++ {
        v13 = (!(((((0.1 + 3.14159) + 0.1) == 1e10) && (!false)) && (1 <= 42)))
        switch ((("\n\t" + "a") + "") + "foo bar") {
        case "foo bar":
            v2 -= 330
        case "hello":
            v2 = v2
        case "\n\t":
            v1 -= 2147483647
        }
        _ = fi36
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v7
    _ = v9
}
