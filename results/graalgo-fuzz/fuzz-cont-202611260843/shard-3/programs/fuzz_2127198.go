package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 int = -1
    var v2 int64 = 91
    var v3 uint64 = 45
    var v4 int32 = 42
    v5 := []int{(v1 + 0), 42, v1}
    if (3.14159 >= float32(float64((((0.0 * (0.0 - 1e10)) - (float32((float64((0.1 - 0.1)) * -2.25)) + ((0.1 - (1e10 * 0.1)) - 3.14159))) - float32(1e10))))) {
        var v6 int8 = 127
        v1 = 4294967295
        _ = v6
    } else {
        if (false && true) {
            for fi7 := 0; fi7 < 14; fi7++ {
                if false {
                    var v8 float32 = 1e10
                    for fi9 := 0; fi9 < 5; fi9++ {
                        if (v2 != 9223372036854775807) {
                            var v10 float32 = (float32((0.0 + ((-2.25 * 1e10) + 1e10))) - (1e10 / v8))
                            var v11 uint8 = 255
                            _ = v10
                            _ = v11
                        }
                        var v12 float64 = ((1e10 - ((1.5 * 1e10) + 1.5)) + (3.14159 + (float64((v8 / -2.25)) / -2.25)))
                        if (2 < 42) {
                            fmt.Println(v12, v4, fi7)
                        }
                        _ = fi9
                        _ = v12
                    }
                    v13 := GenS1{F0: v3}
                    _ = v13
                    _ = v8
                } else {
                    var v14 int64 = (v2 | (v2 % ((v2) | 1)))
                    var v15 float64 = float64(-2.25)
                    _ = v14
                    _ = v15
                }
                _ = fi7
            }
        }
        var v16 int16 = 1
        fmt.Println(v4, v3, v3)
        _ = v16
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v17 bool = (v3 < v3)
    fmt.Println(v17)
    var v18 uint64 = v3
    v17 = (0 >= -1)
    v2 -= (v2 ^ v2)
    v18 += 1
    fmt.Println(v3, v1)
    v2 -= (v2 & -9223372036854775808)
    var v19 uint8 = 1
    var v20 int32 = (v4 * v4)
    if (!v17) {
        var v21 bool = ((!(("" + "\n\t") >= "\n\t")) && (2 > 2))
        _ = v21
    }
    if (!v17) {
        var v22 string = "foo bar"
        v5 = append(v5, v1)
        if (int(v19) != (-(v1))) {
            v23 := GenS1{F0: (v3 / ((1) | 1))}
            var v24 uint = 65535
            var v25 uint32 = 2
            _ = v23
            _ = v24
            _ = v25
        }
        _ = v22
    }
    v2 = -86
    for fi26 := 0; fi26 < 2; fi26++ {
        switch "a" {
        case "\n\t":
            v17 = false
        case "hello":
            switch (^(v1)) {
            case 32768:
                for fi27 := 0; fi27 < 17; fi27++ {
                    if ((false || false) || false) {
                        v28 := GenS1{F0: v18}
                        _ = v28
                    } else {
                        var v29 int = (fi27 % ((v1) | 1))
                        _ = v29
                    }
                    _ = fi27
                }
            case 531:
                v4 = v20
            default:
                v5 = append(v5, v1)
            }
        default:
            for fi30 := 0; fi30 < 6; fi30++ {
                v18 = (v3 / ((v3) | 1))
                var v31 float64 = -2.25
                _ = fi30
                _ = v31
            }
        }
        if v17 {
            var v32 uint16 = 32767
            _ = v32
        } else {
            v3 = v3
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        v4 = (v4 | -2147483648)
        _ = fi26
    }
    var v33 string = ""
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
