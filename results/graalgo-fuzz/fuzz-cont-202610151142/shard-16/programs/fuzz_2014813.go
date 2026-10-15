package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8 }
func genFn0() uint8 {
    v5 := []int{-70, 32768, 32768}
    var v6 float64 = 1.5
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 uint16 = 135
    var v2 float64 = (0.1 - ((0.1 - 0.0) / 1.0))
    var v3 float64 = v2
    v4 := []int{1}
    var v7 int = int(v1)
    var v8 int64 = 42
    switch int(v1) {
    case 0:
        for fi9 := 0; fi9 < 4; fi9++ {
            for fi10 := 0; fi10 < 1; fi10++ {
                var v11 uint = 1
                _ = fi10
                _ = v11
            }
            _ = fi9
        }
    case 962:
        if (!(true && true)) {
            fmt.Println(v7, v2, v7)
            v4 = append(v4, (v7 / ((32767) | 1)))
        } else {
            v4 = append(v4, v7)
        }
    case 2147483647:
        v2 -= float64(1.5)
    default:
        for fi12 := 0; fi12 < 12; fi12++ {
            var v13 int32 = 2147483647
            for fi14 := 0; fi14 < 1; fi14++ {
                switch -129 {
                case 0:
                    switch v7 {
                    case 0:
                        v8 = (v8 >> ((1) & 63))
                    }
                case 875:
                    v4 = append(v4, v7)
                case 2147483647:
                    if true {
                        v13 += v13
                        var v15 int64 = -1
                        _ = v15
                    }
                default:
                    v16 := GenS1{F0: 1}
                    _ = v16
                }
                if ((3.14159 <= float32(3.14159)) && true) {
                    var v17 uint = 4294967295
                    _ = len(v4)
                    _ = v17
                }
                var v18 float64 = ((float64(3.14159) * v3) + v3)
                _ = fi14
                _ = v18
            }
            if false {
                switch "" {
                case "foo bar":
                    v4 = append(v4, fi12)
                case "\n\t":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                default:
                    var v19 int8 = 108
                    _ = v19
                }
                var v20 float64 = (3.14159 + (((v2 / 1.0) + float64((1e10 * 1.5))) / v3))
                v21 := GenS2{F0: -128}
                _ = v20
                _ = v21
            }
            _ = fi12
            _ = v13
        }
    }
    fmt.Println(v1, v7, v7)
    var v22 int = -95
    for fi23 := 0; fi23 < 5; fi23++ {
        _ = len(v4)
        _ = fi23
    }
    var v24 int = (v22 ^ 197)
    fmt.Println(v3, v3)
    var v25 int16 = -328
    var v26 string = ("" + "foo bar")
    fmt.Println(v8, v22, v24)
    switch (v24 + (v22 % ((v7) | 1))) {
    case 1:
        fmt.Println(v25, v24)
    case -599:
        v27 := GenS1{F0: 93}
        _ = v27
    case -1:
        _ = len(v4)
    default:
        v28 := GenS1{F0: uint8(v1)}
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
