package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }

func main() {
    var v1 float64 = ((0.0 * 3.14159) * (3.14159 + (((3.14159 + (1.5 * 1e10)) / float64(0.1)) - 0.0)))
    var v2 int16 = 389
    v3 := []string{"a"}
    var v4 bool = false
    if (1 == 18446744073709551615) {
        v2 = v2
        for fi5 := 0; fi5 < 18; fi5++ {
            var v6 uint16 = 0
            var v7 uint32 = 510
            var v8 float64 = 1e10
            _ = fi5
            _ = v6
            _ = v7
            _ = v8
        }
        v1 -= (float64(0.0) * (float64(0.0) / v1))
    }
    var v9 string = "foo bar"
    v2 += -63
    var v10 string = v9
    if (v1 != (float64((-2.25 - -2.25)) - v1)) {
        v10 = ("\n\t" + v9)
        if false {
            _ = len(v3)
            var v11 bool = true
            _ = len(v3)
            _ = v11
        }
        switch 127 {
        case 4294967295:
            var v12 int32 = 42
            _ = v12
        case -459:
            v1 = v1
        case 9223372036854775807:
            var v13 int16 = 0
            _ = v13
        }
    }
    var v14 uint = 135
    var v15 int32 = 0
    v14 += (^(v14))
    fmt.Println(v10, v1)
    var v16 uint16 = 32767
    switch v10 {
    case "hello":
        if (9223372036854775807 > 128) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            v17 := GenS1{F0: 18446744073709551615, F1: 738}
            for fi18 := 0; fi18 < 1; fi18++ {
                if v4 {
                    if true {
                        v15 = (v15 << ((v14) & 63))
                        for fi19 := 0; fi19 < 19; fi19++ {
                            var v20 int16 = (v2 * (v2 % (((v2 + v2)) | 1)))
                            _ = fi19
                            _ = v20
                        }
                    } else {
                        var v21 int = fi18
                        v3 = append(v3, v9)
                        _ = v21
                    }
                } else {
                    if (1.5 < (v1 * v1)) {
                        var v22 uint16 = v16
                        _ = v22
                    } else {
                        if ((!v4) || (int(v14) != (-(fi18)))) {
                            v23 := GenS1{F0: 547, F1: 42}
                            v9 = v9
                            _ = v23
                        }
                        fmt.Println(v14, fi18, v10)
                        var v24 uint16 = 680
                        _ = v24
                    }
                    if (2 == 2) {
                        if (("a" >= v10) && (((316 == (v16 << (((v14 ^ (v14 + (v14 | (v14 ^ (v14 ^ 685)))))) & 63))) && false) || (0 == 6))) {
                            var v25 bool = v4
                            _ = v25
                        }
                    }
                }
                v15 = (v15 / ((-1) | 1))
                v26 := GenS1{F0: 158, F1: 284}
                _ = fi18
                _ = v26
            }
            _ = v17
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
