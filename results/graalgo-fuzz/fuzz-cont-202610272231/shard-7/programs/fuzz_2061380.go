package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int }
func genFn0(p0 uint64, p1 int, p2 uint64) int {
    switch (p1 % ((2147483647) | 1)) {
    case 825:
        for fi7 := 0; fi7 < 16; fi7++ {
            if false {
                var v8 int8 = 100
                fmt.Println(p0)
                _ = v8
            }
            p2 = 498
            _ = fi7
        }
    case 32767:
        var v9 uint32 = 70
        _ = v9
    default:
        v10 := GenS1{F0: 2147483647, F1: uint32(p0)}
        _ = v10
    }
    v11 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 873
}

func main() {
    var v1 int16 = 1
    var v2 uint16 = 65535
    var v3 float64 = 3.14159
    var v4 int64 = -1
    var v5 float32 = 0.1
    var v6 uint16 = (^(v2))
    v2 = (v6 - v6)
    v12 := GenS1{F0: 0, F1: 42}
    var v13 int = 2147483647
    var v14 float64 = (3.14159 + 3.14159)
    v15 := []int{(v13 / (((^(v13))) | 1))}
    if true {
        switch v13 {
        case 127:
            v6 = v2
        default:
            for fi16 := 0; fi16 < 16; fi16++ {
                v3 -= (-2.25 + 1e10)
                v1 = 873
                _ = fi16
            }
        }
    } else {
        for fi17 := 0; fi17 < 5; fi17++ {
            fmt.Println(v14, v6, v5)
            _ = fi17
        }
        v18 := GenS1{F0: 42, F1: 2147483648}
        for fi19 := 0; fi19 < 3; fi19++ {
            var v20 float64 = float64(v5)
            v21 := GenS2{F0: 317}
            v22 := GenS2{F0: 617}
            _ = fi19
            _ = v20
            _ = v21
            _ = v22
        }
        _ = v18
    }
    v6 -= (v2 - (^(v6)))
    if false {
        for fi23 := 0; fi23 < 9; fi23++ {
            if true {
                if ((!false) && (true || false)) {
                    switch (^(fi23)) {
                    case 128:
                        if ("\n\t" < "foo bar") {
                            v4 = (^(v4))
                            for fi24 := 0; fi24 < 20; fi24++ {
                                v13 -= (-(v13))
                                v13 = genFn0(2, (fi24 % ((genFn0(185, genFn0(1, 0, 2), 4294967295)) | 1)), 30)
                                _ = fi24
                            }
                        }
                    case -1:
                        var v25 uint = 4294967295
                        _ = v25
                    }
                    var v26 float64 = ((v14 * ((float64(((1.5 * float32((0.1 / (v3 / (1.5 + v14))))) / v5)) / 1.5) * (0.1 * float64(v5)))) / ((v3 + 1e10) + 1.5))
                    _ = v26
                }
            }
            _ = fi23
        }
    } else {
        switch "" {
        case "":
            switch "hello" {
            case "hello":
                var v27 uint8 = 127
                _ = v27
            case "":
                switch 32767 {
                case 32768:
                    if (true && false) {
                        var v28 int32 = int32(v13)
                        v6 -= (^(v2))
                        if len(v15) > 0 {
                            v15[0] = genFn0(18446744073709551615, (v13 | (v13 * 0)), 4294967295)
                        }
                        _ = v28
                    } else {
                        for fi29 := 0; fi29 < 6; fi29++ {
                            v2 = v2
                            v30 := GenS3{F0: (v13 << ((18446744073709551615) & 63))}
                            _ = fi29
                            _ = v30
                        }
                        v14 = v14
                    }
                case 1:
                    v31 := GenS1{F0: 847, F1: 65535}
                    _ = v31
                }
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
