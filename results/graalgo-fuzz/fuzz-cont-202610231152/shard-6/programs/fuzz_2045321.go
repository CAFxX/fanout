package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
func genFn0(p0 uint64) int16 {
    var v5 bool = (!(42 >= p0))
    var v6 float64 = 1e10
    switch 450 {
    case -933:
        v7 := []int{-2147483648, -9223372036854775808}
        _ = v7
    default:
        for fi8 := 0; fi8 < 11; fi8++ {
            v5 = v5
            v9 := []int{}
            _ = fi8
            _ = v9
        }
    }
    var v10 int32 = -2147483648
    _ = p0
    _ = v10
    _ = v5
    _ = v6
    return 1
}

func main() {
    var v1 int8 = -1
    var v2 uint64 = 18446744073709551615
    v3 := []string{"foo bar", ""}
    var v4 int8 = -128
    v1 += v1
    v1 = v1
    if false {
        if len(v3) > 0 {
            v3[0] = "hello"
        }
        if (false && true) {
            var v11 uint64 = (^(v2))
            _ = v11
        } else {
            v1 = (v4 - (v4 << ((806) & 63)))
            for fi12 := 0; fi12 < 8; fi12++ {
                switch (fi12 ^ (fi12 / ((fi12) | 1))) {
                case -1:
                    v13 := GenS1{F0: fi12, F1: fi12}
                    _ = v13
                case 128:
                    var v14 int32 = -2147483648
                    _ = v14
                default:
                    switch fi12 {
                    case 4294967295:
                        var v15 uint = 4294967295
                        _ = v15
                    case -9223372036854775808:
                        var v16 uint64 = 1
                        _ = v16
                    }
                }
                _ = fi12
            }
            switch "\n\t" {
            case "\n\t":
                if (127 <= -2147483648) {
                    var v17 int = -690
                    v2 -= 4294967295
                    _ = v17
                }
            case "foo bar":
                if (0 > 1) {
                    fmt.Println(v2)
                    v1 += -128
                    var v18 float64 = 1e10
                    _ = v18
                } else {
                    for fi19 := 0; fi19 < 17; fi19++ {
                        v4 = v4
                        if false {
                            v4 = (-(v1))
                            v1 += (v4 | v4)
                        } else {
                            var v20 float32 = float32(0.0)
                            _ = v20
                        }
                        _ = fi19
                    }
                    v21 := GenS1{F0: 127, F1: 4294967295}
                    v2 -= (v2 / (((v2 + v2)) | 1))
                    _ = v21
                }
            case "":
                if (v4 >= v4) {
                    if false {
                        fmt.Println(v2)
                        v2 += v2
                    } else {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                }
            default:
                var v22 uint = 255
                _ = v22
            }
        }
    } else {
        switch -846 {
        case -905:
            fmt.Println(v1)
        }
    }
    var v23 int16 = -965
    v23 -= 1
    v23 = 0
    fmt.Println(v2, v2)
    fmt.Println(v1, v23, v1)
    var v24 uint = 157
    var v25 int8 = v1
    if (1e10 != (((3.14159 - 1e10) - float64((0.0 * (((3.14159 / (3.14159 * (-2.25 - 1.5))) + ((-2.25 * 3.14159) * 1.5)) - ((3.14159 - 3.14159) / 3.14159))))) + 0.1)) {
        var v26 string = "hello"
        v25 = (v25 >> (((v24 >> ((168) & 63))) & 63))
        if len(v3) > 0 {
            v3[0] = "\n\t"
        }
        _ = v26
    } else {
        v27 := GenS1{F0: 4294967295, F1: -558}
        _ = v27
    }
    var v28 uint32 = 2147483648
    v29 := GenS1{F0: 4294967295, F1: int(v24)}
    switch ("" + "") {
    case "a":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "foo bar":
        switch 32768 {
        case 1:
            v4 += -1
        default:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    case "":
        v2 = 1
    default:
        var v30 uint32 = v28
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
