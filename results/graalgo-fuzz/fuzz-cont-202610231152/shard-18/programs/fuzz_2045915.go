package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 int32; F2 uint8 }
func genFn0(p0 bool, p1 int, p2 int8) bool {
    var v4 bool = (1e10 == 0.0)
    if ((v4 || (!false)) || (-1 != 762)) {
        p0 = (18446744073709551615 == 42)
    } else {
        p2 = (p2 << ((168) & 63))
        switch ("\n\t" + "hello") {
        case "a":
            v5 := []int{}
            _ = v5
        default:
            p2 -= (p2 << ((1) & 63))
        }
    }
    var v6 float64 = (3.14159 + 1e10)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v6
    return p0
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint16 = 0
    var v3 uint = 0
    for fi7 := 0; fi7 < 5; fi7++ {
        if (false && (!true)) {
            var v8 int = fi7
            fmt.Println(v8, fi7, v3)
            v8 = fi7
            _ = v8
        }
        if (9223372036854775807 != 1) {
            v3 = (v3 * 255)
            for fi9 := 0; fi9 < 14; fi9++ {
                v2 = 42
                v3 = (v3 & 42)
                v2 -= (v2 >> ((v3) & 63))
                _ = fi9
            }
            for fi10 := 0; fi10 < 5; fi10++ {
                v3 += 18446744073709551615
                var v11 int = 1
                var v12 uint = uint(v2)
                _ = fi10
                _ = v11
                _ = v12
            }
        }
        _ = fi7
    }
    var v13 bool = true
    v3 = v3
    var v14 int32 = -569
    for fi15 := 0; fi15 < 11; fi15++ {
        for fi16 := 0; fi16 < 9; fi16++ {
            v17 := GenS1{F0: v14}
            if false {
                v2 = (v2 + 1)
                switch ("\n\t" + "hello") {
                case "a":
                    var v18 uint8 = 128
                    _ = v18
                case "":
                    v2 += (v2 << ((171) & 63))
                case "foo bar":
                    v19 := []int{(fi15 | fi16), fi15, fi16}
                    _ = v19
                default:
                    if (2 == 249) {
                        switch ("" + (("\n\t" + "\n\t") + "\n\t")) {
                        case "a":
                            var v20 uint = (v3 & (v3 / (((v3 - 733)) | 1)))
                            _ = v20
                        case "hello":
                            v21 := []int{32767, fi15, fi16, 9223372036854775807}
                            _ = v21
                        case "foo bar":
                            v14 = v14
                        default:
                            v3 -= v3
                        }
                    } else {
                        fmt.Println(v3)
                    }
                }
                v14 = (-(v14))
            }
            _ = fi16
            _ = v17
        }
        v1[("" + "")] = v13
        v22 := GenS2{F0: (v2 << ((v3) & 63)), F1: -2147483648, F2: 144}
        _ = fi15
        _ = v22
    }
    v13 = (!(1.5 <= -2.25))
    delete(v1, "\n\t")
    v14 = v14
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
}
