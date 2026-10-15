package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 int64; F1 int16; F2 int64 }
type GenS3 struct { F0 uint; F1 int32; F2 uint32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int16 }
func genFn0(p0 int8) string {
    v5 := GenS1{F0: 42, F1: 1}
    var v6 float64 = 0.1
    var v7 string = "foo bar"
    if true {
        v8 := []int{2147483647, 0}
        switch v7 {
        case "\n\t":
            fmt.Println(v6, p0)
        case "a":
            switch -20 {
            case 32768:
                if false {
                    fmt.Println(v6)
                    switch -856 {
                    case 32767:
                        if (-323 <= -9223372036854775808) {
                            v9 := GenS1{F0: 18446744073709551615, F1: 0}
                            _ = v9
                        }
                    }
                    v8 = append(v8, 127)
                }
            case 9223372036854775807:
                p0 -= p0
            case 127:
                if len(v8) > 0 {
                    v8[0] = 32767
                }
            }
        case "hello":
            v10 := GenS1{F0: 18446744073709551615, F1: 684}
            _ = v10
        default:
            switch "a" {
            case "\n\t":
                v8 = append(v8, 2147483647)
            case "":
                var v11 uint8 = 0
                _ = v11
            case "a":
                switch -1 {
                case -854:
                    v12 := GenS1{F0: 940, F1: 627}
                    _ = v12
                case 9223372036854775807:
                    v6 = v6
                case 2147483647:
                    for fi13 := 0; fi13 < 16; fi13++ {
                        var v14 float64 = v6
                        p0 += (p0 >> ((42) & 63))
                        _ = fi13
                        _ = v14
                    }
                default:
                    for fi15 := 0; fi15 < 18; fi15++ {
                        var v16 uint64 = 42
                        for fi17 := 0; fi17 < 14; fi17++ {
                            v18 := GenS2{F0: 26, F1: -929, F2: -9223372036854775808}
                            _ = fi17
                            _ = v18
                        }
                        v19 := GenS3{F0: 18446744073709551615, F1: 1, F2: 42}
                        _ = fi15
                        _ = v16
                        _ = v19
                    }
                }
            }
        }
        _ = v8
    }
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return v7
}

func main() {
    var v1 int = -66
    var v2 uint32 = 2
    var v3 bool = (!false)
    var v4 uint16 = 1
    fmt.Println(v4, v4, v1)
    fmt.Println(v2)
    for fi20 := 0; fi20 < 10; fi20++ {
        if (1 <= -1) {
            v3 = (18446744073709551615 > 4294967295)
        } else {
            for fi21 := 0; fi21 < 12; fi21++ {
                v22 := GenS1{F0: 4294967295, F1: 398}
                v23 := []int{fi21, (^(fi21))}
                _ = fi21
                _ = v22
                _ = v23
            }
            v4 = 814
            v2 = v2
        }
        _ = fi20
    }
    v24 := []int{}
    v25 := GenS4{F0: 42}
    v4 += v4
    v2 -= v2
    v26 := GenS5{F0: -1}
    fmt.Println(v2, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
