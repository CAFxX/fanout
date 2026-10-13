package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8; F1 uint64; F2 int64 }
func genFn0(p0 uint32, p1 uint, p2 uint) int64 {
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := GenS1{F0: 0}
        for fi9 := 0; fi9 < 19; fi9++ {
            v10 := GenS1{F0: 255}
            _ = fi9
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    if (false || (0 != 0)) {
        v11 := GenS1{F0: 32767}
        var v12 uint32 = 1
        var v13 uint = (^(p2))
        _ = v11
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 uint64 = 4294967295
    var v2 int64 = 9223372036854775807
    var v3 uint8 = 127
    var v4 uint32 = 2
    v5 := []string{}
    var v6 uint16 = 293
    var v14 int16 = 1000
    var v15 int16 = (v14 ^ v14)
    fmt.Println(v1)
    var v16 float64 = (1e10 / 0.1)
    v17 := GenS1{F0: 255}
    var v18 int8 = 42
    for fi19 := 0; fi19 < 16; fi19++ {
        v1 = (^(v1))
        _ = fi19
    }
    v2 -= v2
    if (false && (true || ((false && false) && true))) {
        switch 34 {
        case 2147483647:
            for fi20 := 0; fi20 < 20; fi20++ {
                v6 = (v6 / ((v6) | 1))
                for fi21 := 0; fi21 < 3; fi21++ {
                    var v22 string = "\n\t"
                    if len(v5) > 0 {
                        v5[0] = "\n\t"
                    }
                    _ = fi21
                    _ = v22
                }
                _ = fi20
            }
        case 42:
            for fi23 := 0; fi23 < 5; fi23++ {
                switch fi23 {
                case 0:
                    var v24 uint8 = (v3 << ((1) & 63))
                    _ = v24
                case 32768:
                    var v25 uint = uint(v3)
                    _ = v25
                default:
                    for fi26 := 0; fi26 < 4; fi26++ {
                        if (0.1 != v16) {
                            var v27 uint16 = (v6 - (^(v6)))
                            _ = v27
                        } else {
                            var v28 float64 = v16
                            var v29 int = fi23
                            fmt.Println(v28)
                            _ = v28
                            _ = v29
                        }
                        var v30 int16 = (v14 * v15)
                        if ((v4 % (((v4 - v4)) | 1)) >= (v4 >> ((2) & 63))) {
                            v16 = v16
                            fmt.Println(v16, v6)
                        } else {
                            if len(v5) > 0 {
                                v5[0] = "\n\t"
                            }
                        }
                        _ = fi26
                        _ = v30
                    }
                }
                v18 += (v18 * (v18 & v18))
                var v31 bool = false
                _ = fi23
                _ = v31
            }
        }
    }
    v32 := GenS2{F0: 0}
    if (1 <= -2147483648) {
        v15 += (v15 * (v14 << ((65535) & 63)))
    }
    v33 := GenS3{F0: (v3 * 255), F1: (v1 ^ (v1 ^ v1)), F2: (v2 << ((1) & 63))}
    var v34 uint8 = (^(v3))
    v15 -= v15
    var v35 uint = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
