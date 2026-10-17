package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 int16; F1 int64 }
func genFn0(p0 int64, p1 string, p2 int32) int8 {
    v6 := []int{9223372036854775807, 273, 127, 128}
    var v7 uint16 = 756
    v8 := GenS1{F0: v7, F1: (4294967295 >= 1)}
    for fi9 := 0; fi9 < 17; fi9++ {
        var v10 int64 = p0
        var v11 uint32 = 722
        v10 = p0
        _ = fi9
        _ = v10
        _ = v11
    }
    if (!(!false)) {
        v12 := GenS1{F0: 1, F1: true}
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return -1
}

func main() {
    var v1 uint16 = 791
    var v2 uint = 255
    var v3 string = "foo bar"
    var v4 uint32 = 2147483648
    var v5 uint32 = v4
    var v13 int16 = 981
    v14 := []int{1}
    for fi15 := 0; fi15 < 4; fi15++ {
        var v16 int = 305
        _ = fi15
        _ = v16
    }
    if true {
        var v17 uint32 = v5
        if ((v2 & v2) == v2) {
            fmt.Println(v13)
            if len(v14) > 0 {
                v14[0] = -8
            }
        } else {
            var v18 int8 = 0
            _ = v18
        }
        _ = v17
    } else {
        switch 261 {
        case 1:
            var v19 uint16 = (v1 * 32767)
            _ = v19
        default:
            v1 += (v1 + 536)
        }
        if (-421 <= 0) {
            v13 += 1
        } else {
            if false {
                v5 += v4
                v20 := GenS2{F0: v13, F1: -9223372036854775808}
                _ = v20
            } else {
                if false {
                    v14 = append(v14, 32767)
                    if (!false) {
                        v4 -= 2
                        v1 -= v1
                        if len(v14) >= 2 {
                            v14 = v14[1:len(v14)]
                        }
                    }
                } else {
                    v1 = v1
                    if len(v14) >= 2 {
                        v14 = v14[1:len(v14)]
                    }
                }
            }
        }
    }
    v2 += v2
    v2 = (v2 + (v2 >> (((v2 ^ 4294967295)) & 63)))
    var v21 uint32 = 65535
    _ = len(v14)
    v22 := GenS2{F0: (v13 & v13), F1: 42}
    _ = len(v14)
    fmt.Println(v5)
    v5 = v5
    var v23 int32 = -633
    v4 -= v5
    for fi24 := 0; fi24 < 1; fi24++ {
        switch (fi24 + (fi24 % (((^(fi24))) | 1))) {
        case 128:
            var v25 bool = false
            _ = v25
        default:
            var v26 int32 = (v23 >> ((uint(v5)) & 63))
            _ = v26
        }
        switch ("foo bar" + v3) {
        case "hello":
            v23 += (v23 + (v23 & v23))
        case "a":
            if len(v14) > 0 {
                v14[0] = int(v5)
            }
        default:
            v27 := GenS2{F0: v13, F1: 1}
            _ = v27
        }
        _ = fi24
    }
    v21 -= 89
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
