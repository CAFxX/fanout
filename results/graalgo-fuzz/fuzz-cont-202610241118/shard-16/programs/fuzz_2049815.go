package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 uint8; F2 int }
func genFn0(p0 uint16, p1 uint8, p2 string) uint {
    if true {
        if (true && true) {
            for fi6 := 0; fi6 < 16; fi6++ {
                v7 := []int{(fi6 >> ((4294967295) & 63)), -2147483648, (fi6 ^ (fi6 * (^(fi6))))}
                _ = fi6
                _ = v7
            }
            var v8 float64 = 0.1
            _ = v8
        }
    }
    var v9 int64 = 9223372036854775807
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 632
}

func main() {
    var v1 bool = (!(-2147483648 == 1))
    var v2 float32 = (0.0 * (((3.14159 * -2.25) - (0.0 + (float32(0.0) * 0.0))) * 3.14159))
    v3 := []uint32{}
    var v4 uint32 = 2
    var v5 uint64 = 633
    v2 = (1e10 - v2)
    v2 -= (1.5 + (float32(1e10) / (1.5 + (v2 / (((1e10 + 0.1) * float32(0.0)) + (v2 * (v2 - v2)))))))
    for fi10 := 0; fi10 < 7; fi10++ {
        v1 = true
        var v11 bool = (((65535 >= 18446744073709551615) && (0 == -1)) || (88 != 253))
        _ = fi10
        _ = v11
    }
    switch "foo bar" {
    case "\n\t":
        if (!(!false)) {
            var v12 uint8 = uint8(v5)
            if false {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v3 = append(v3, (v4 / ((uint32(v5)) | 1)))
            } else {
                if (!(v12 > v12)) {
                    _ = len(v3)
                }
                var v13 uint32 = 49
                var v14 string = "foo bar"
                _ = v13
                _ = v14
            }
            for fi15 := 0; fi15 < 1; fi15++ {
                fmt.Println(v2, v12)
                switch "" {
                case "\n\t":
                    v16 := GenS1{F0: -1}
                    _ = v16
                case "":
                    fmt.Println(v12, fi15)
                case "foo bar":
                    v12 -= (v12 - v12)
                }
                _ = fi15
            }
            _ = v12
        } else {
            if v1 {
                fmt.Println(v4, v2)
                for fi17 := 0; fi17 < 5; fi17++ {
                    var v18 bool = ("hello" > "")
                    fmt.Println(v18)
                    if (fi17 > (fi17 >> ((65535) & 63))) {
                        v1 = v18
                        var v19 float64 = 0.1
                        _ = v19
                    }
                    _ = fi17
                    _ = v18
                }
                for fi20 := 0; fi20 < 10; fi20++ {
                    for fi21 := 0; fi21 < 17; fi21++ {
                        if len(v3) > 0 {
                            v3[0] = v4
                        }
                        if false {
                            var v22 uint = 716
                            var v23 uint8 = 127
                            fmt.Println(v5, v5)
                            _ = v22
                            _ = v23
                        }
                        v5 = 42
                        _ = fi21
                    }
                    v5 = (v5 << ((65535) & 63))
                    _ = fi20
                }
            }
        }
    default:
        v5 = 1
    }
    v24 := GenS2{F0: 0, F1: 2, F2: 128}
    switch -351 {
    case -1:
        var v25 bool = v1
        _ = v25
    }
    v4 = 901
    v3 = append(v3, v4)
    var v26 uint8 = 26
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
