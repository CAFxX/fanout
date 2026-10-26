package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 bool }
type GenS2 struct { F0 bool; F1 int; F2 uint64 }

func main() {
    var v1 string = "foo bar"
    var v2 int8 = -128
    v3 := make(map[int]string)
    v4 := []string{(((v1 + v1) + "\n\t") + v1), "hello", ("" + v1)}
    var v5 uint32 = 495
    var v6 float32 = 3.14159
    v7 := v3[212]
    v8 := GenS1{F0: (v2 * v2), F1: -32768, F2: (18446744073709551615 >= 1)}
    switch v1 {
    case "a":
        fmt.Println(v1)
    case "":
        fmt.Println(v2)
    }
    if (!(false && (2 >= 42))) {
        switch ("" + "a") {
        case "a":
            if false {
                var v9 int16 = 1000
                for fi10 := 0; fi10 < 19; fi10++ {
                    for fi11 := 0; fi11 < 2; fi11++ {
                        v9 = v9
                        if (0 >= 269) {
                            var v12 uint32 = 2147483648
                            _ = v12
                        } else {
                            v7 = v1
                            v9 += (v9 << ((18446744073709551615) & 63))
                            v13 := v3[fi10]
                            _ = v13
                        }
                        _ = fi11
                    }
                    fmt.Println(v1, fi10, v9)
                    _ = fi10
                }
                v9 = 159
                _ = v9
            }
        }
    }
    v5 -= (v5 * (^(v5)))
    if (235 <= 127) {
        var v14 bool = false
        for fi15 := 0; fi15 < 11; fi15++ {
            fmt.Println(v2, fi15)
            _ = fi15
        }
        _ = v14
    } else {
        var v16 float64 = 0.0
        v16 += float64(((-2.25 * -2.25) + -2.25))
        fmt.Println(v1, v6)
        _ = v16
    }
    v6 = ((3.14159 + (3.14159 / ((v6 + float32(-2.25)) * 0.0))) / v6)
    var v17 uint16 = 32767
    switch -9223372036854775808 {
    case 887:
        v17 -= (^(v17))
    case 127:
        v17 += v17
    case 0:
        if (1 != -1) {
            v18 := v3[-1]
            _ = v18
        } else {
            v19 := GenS1{F0: 100, F1: -392, F2: ((42 == 465) || false)}
            if (int32(v5) < -1) {
                v4 = append(v4, ((v7 + ("" + v7)) + v1))
                v6 -= ((v6 / 0.1) - 0.0)
            } else {
                v20 := GenS2{F0: (1 >= 255), F1: 842, F2: 583}
                _ = v20
            }
            if (-2.25 >= (-2.25 - -2.25)) {
                v2 -= (v2 + v2)
            }
            _ = v19
        }
    }
    switch (v7 + "\n\t") {
    case "\n\t":
        for fi21 := 0; fi21 < 16; fi21++ {
            fmt.Println(v7, v5, v7)
            _ = len(v3)
            _ = fi21
        }
    default:
        for fi22 := 0; fi22 < 9; fi22++ {
            if false {
                v23 := v3[-327]
                _ = v23
            } else {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v24 := GenS1{F0: v2, F1: 42, F2: (!false)}
                _ = v24
            }
            var v25 uint64 = uint64(v17)
            _ = fi22
            _ = v25
        }
    }
    fmt.Println(v5, v17)
    v1 = "foo bar"
    if (!(32768 != 945)) {
        v26 := GenS2{F0: true, F1: -1, F2: 18446744073709551615}
        v27 := v3[2147483647]
        _ = v26
        _ = v27
    }
    v17 -= (v17 * v17)
    v2 += (v2 << ((839) & 63))
    fmt.Println(v5, v7)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
