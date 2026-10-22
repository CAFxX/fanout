package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint }

func main() {
    v1 := []int{4294967295, -1, -569}
    v2 := []uint32{102, 687}
    var v3 int16 = -1
    var v4 float64 = -2.25
    var v5 uint32 = 65535
    var v6 int32 = 843
    v4 = v4
    for fi7 := 0; fi7 < 16; fi7++ {
        for fi8 := 0; fi8 < 2; fi8++ {
            var v9 int = (fi8 - (fi8 | fi8))
            var v10 float32 = 1e10
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    fmt.Println(v5, v4)
    v11 := GenS1{F0: (v6 - (^(v6))), F1: -37, F2: 1}
    for fi12 := 0; fi12 < 5; fi12++ {
        fmt.Println(v6)
        if ((42 <= 42) || (!false)) {
            switch -2147483648 {
            case 724:
                fmt.Println(v3)
            case 32768:
                if false {
                    for fi13 := 0; fi13 < 13; fi13++ {
                        var v14 string = ""
                        _ = fi13
                        _ = v14
                    }
                    v3 += 32767
                    var v15 int = (fi12 ^ fi12)
                    _ = v15
                } else {
                    var v16 uint32 = v5
                    v6 = 0
                    _ = v16
                }
            default:
                var v17 float64 = v4
                _ = v17
            }
        } else {
            v6 += 2147483647
            switch "\n\t" {
            case "hello":
                fmt.Println(v3, v4, fi12)
            }
        }
        _ = fi12
    }
    v5 += 1
    if true {
        var v18 int = -2147483648
        fmt.Println(v4, v4, v18)
        _ = v18
    }
    switch "a" {
    case "foo bar":
        v3 = -1
    default:
        v4 = v4
    }
    if (!false) {
        v6 = v6
        switch -276 {
        case -644:
            if false {
                if false {
                    v19 := GenS2{F0: 581}
                    fmt.Println(v4, v5)
                    v20 := GenS3{F0: 0}
                    _ = v19
                    _ = v20
                }
                v21 := GenS1{F0: (v6 % (((v6 & (v6 | v6))) | 1)), F1: 100, F2: 144}
                _ = v21
            } else {
                switch 127 {
                case 0:
                    if true {
                        v4 += ((1e10 - v4) + -2.25)
                    } else {
                        switch 42 {
                        case 9223372036854775807:
                            fmt.Println(v3, v5)
                        case 32767:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        default:
                            var v22 float32 = ((1.5 + ((3.14159 + 1.5) * (3.14159 + 1.5))) + (0.0 / -2.25))
                            _ = v22
                        }
                    }
                case 2147483647:
                    fmt.Println(v3)
                case -1:
                    v4 = ((v4 + (v4 - v4)) * (v4 / 1.5))
                }
                for fi23 := 0; fi23 < 14; fi23++ {
                    var v24 float32 = 1.5
                    if len(v2) > 0 {
                        v2[0] = (v5 | v5)
                    }
                    _ = fi23
                    _ = v24
                }
                var v25 int32 = 0
                _ = v25
            }
        case 13:
            v26 := GenS1{F0: v6, F1: -21, F2: 255}
            _ = v26
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v27 uint = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
