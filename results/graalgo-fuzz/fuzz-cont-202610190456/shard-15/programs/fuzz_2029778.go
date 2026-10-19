package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int }
type GenS2 struct { F0 int32; F1 uint16; F2 int32 }

func main() {
    var v1 bool = false
    v2 := []int{-820}
    var v3 float32 = float32(1e10)
    var v4 uint = 255
    var v5 string = (("" + "") + "a")
    if (1.5 != -2.25) {
        if len(v2) > 0 {
            v2[0] = 665
        }
        v6 := GenS1{F0: 993, F1: ("hello" + "foo bar"), F2: 4294967295}
        if (!(!true)) {
            v7 := GenS1{F0: 1, F1: (v5 + v5), F2: 0}
            _ = v7
        } else {
            var v8 int16 = 1
            if (0.0 == ((3.14159 * 1.5) * -2.25)) {
                if len(v2) > 0 {
                    v2[0] = -632
                }
                v9 := GenS2{F0: 924, F1: 507, F2: -843}
                if false {
                    v10 := GenS1{F0: (v8 >> ((v4) & 63)), F1: v5, F2: -617}
                    v1 = v1
                    _ = v10
                } else {
                    for fi11 := 0; fi11 < 17; fi11++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        if (65535 >= 487) {
                            var v12 int32 = 42
                            _ = v12
                        } else {
                            var v13 float32 = v3
                            _ = v13
                        }
                        _ = fi11
                    }
                    fmt.Println(v4, v8, v8)
                }
                _ = v9
            } else {
                var v14 float64 = 3.14159
                _ = v14
            }
            _ = v8
        }
        _ = v6
    } else {
        for fi15 := 0; fi15 < 3; fi15++ {
            if v1 {
                var v16 float32 = v3
                _ = v16
            } else {
                v5 = "a"
            }
            var v17 uint32 = 42
            v17 = (v17 & 998)
            _ = fi15
            _ = v17
        }
        var v18 uint16 = 131
        _ = v18
    }
    if ((!(!false)) && false) {
        switch "" {
        case "a":
            fmt.Println(v5, v4)
        case "hello":
            fmt.Println(v1, v5, v5)
        case "foo bar":
            if v1 {
                var v19 int32 = 42
                var v20 int = 42
                v21 := GenS1{F0: 1, F1: v5, F2: v20}
                _ = v19
                _ = v20
                _ = v21
            }
        }
        fmt.Println(v4, v1, v1)
    }
    fmt.Println(v3)
    var v22 float32 = (v3 * ((v3 + float32(1e10)) / -2.25))
    var v23 int64 = 1
    v4 = v4
    v5 = ("foo bar" + ((v5 + (v5 + (v5 + "\n\t"))) + v5))
    if v1 {
        v3 = ((1.5 + float32((0.0 - float64(((v22 * v3) + ((float32(1.5) / v22) / 0.1)))))) / -2.25)
        for fi24 := 0; fi24 < 19; fi24++ {
            var v25 uint8 = 0
            v4 -= v4
            _ = fi24
            _ = v25
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
