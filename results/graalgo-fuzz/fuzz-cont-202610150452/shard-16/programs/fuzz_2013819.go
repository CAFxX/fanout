package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 int64; F2 int16 }
type GenS3 struct { F0 bool; F1 bool }
type GenS4 struct { F0 uint32 }

func main() {
    var v1 int16 = -32768
    v2 := make(map[uint32]uint32)
    var v3 float32 = (1e10 / 1.5)
    v4 := []int{127, 212}
    var v5 bool = (472 > 4294967295)
    var v6 float32 = ((float32(((((1e10 + -2.25) + 0.0) / ((0.1 / 1.5) - -2.25)) + float64(v3))) / (v3 * float32(float64((v3 - -2.25))))) + v3)
    fmt.Println(v6, v5)
    _ = len(v4)
    v1 = (^(v1))
    for fi7 := 0; fi7 < 8; fi7++ {
        if true {
            for fi8 := 0; fi8 < 5; fi8++ {
                v9 := v2[441]
                if (42 <= 42) {
                    v5 = (!v5)
                    v10 := GenS1{F0: (v5 && false)}
                    _ = v10
                }
                switch fi8 {
                case 0:
                    var v11 uint64 = 2
                    _ = v11
                default:
                    var v12 string = "hello"
                    _ = v12
                }
                _ = fi8
                _ = v9
            }
            switch (^(fi7)) {
            case 619:
                fmt.Println(v1)
            }
            var v13 uint32 = 4294967295
            _ = v13
        } else {
            v3 += 1e10
            var v14 float32 = (v3 - v3)
            _ = v14
        }
        fmt.Println(fi7, v5)
        _ = fi7
    }
    v1 += (v1 | -141)
    var v15 uint32 = 2
    switch -2147483648 {
    case -526:
        if (-9223372036854775808 < 32768) {
            if len(v4) > 0 {
                v4[0] = 127
            }
            for fi16 := 0; fi16 < 10; fi16++ {
                v1 -= (-(v1))
                v5 = (!(1 <= -580))
                switch fi16 {
                case 296:
                    switch fi16 {
                    case 127:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    default:
                        v17 := v2[v15]
                        _ = v17
                    }
                }
                _ = fi16
            }
        } else {
            v18 := GenS2{F0: ("" + "hello"), F1: 656, F2: (v1 + 42)}
            v19 := GenS3{F0: (v5 && v5), F1: v5}
            _ = v18
            _ = v19
        }
    case 32767:
        for fi20 := 0; fi20 < 8; fi20++ {
            switch ("\n\t" + "foo bar") {
            case "a":
                v5 = (1e10 >= 1.5)
            case "foo bar":
                if v5 {
                    var v21 string = "a"
                    v15 += (v15 + 65535)
                    switch v21 {
                    case "":
                        v22 := GenS1{F0: (-9223372036854775808 <= 9223372036854775807)}
                        _ = v22
                    case "a":
                        var v23 uint64 = 266
                        _ = v23
                    default:
                        fmt.Println(v21, v6, fi20)
                    }
                    _ = v21
                } else {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
            default:
                var v24 uint = 65535
                _ = v24
            }
            fmt.Println(v1, v15, fi20)
            _ = fi20
        }
    case 9223372036854775807:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    default:
        v6 = 3.14159
    }
    var v25 float32 = v3
    v26 := GenS4{F0: 693}
    var v27 int16 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
