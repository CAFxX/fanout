package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 uint16 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int16 = -245
    v3 := make(map[string]string)
    v4 := []uint32{2147483647, 0}
    v5 := []int{0, 241}
    var v6 int16 = v2
    for fi7 := 0; fi7 < 11; fi7++ {
        fmt.Println(v2)
        switch (-(fi7)) {
        case -491:
            fmt.Println(fi7)
        case 602:
            switch ("hello" + ("\n\t" + ("a" + "a"))) {
            case "hello":
                if (-2.25 < 1e10) {
                    for fi8 := 0; fi8 < 14; fi8++ {
                        var v9 int32 = -1
                        delete(v1, 231)
                        v9 = v9
                        _ = fi8
                        _ = v9
                    }
                    var v10 int = (fi7 - (fi7 ^ (fi7 % ((fi7) | 1))))
                    _ = v10
                } else {
                    v11 := v3[("hello" + ("" + (("hello" + "\n\t") + "\n\t")))]
                    _ = v11
                }
            case "":
                switch "hello" {
                case "\n\t":
                    v6 = 536
                case "a":
                    v6 = (v2 + (v6 | v6))
                }
            default:
                v2 = v2
            }
        case 4294967295:
            if (42 <= 21) {
                v2 += v2
            } else {
                v6 -= (v6 ^ v6)
                for fi12 := 0; fi12 < 5; fi12++ {
                    var v13 int32 = -421
                    var v14 string = ""
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            }
        default:
            if true {
                switch (fi7 & fi7) {
                case 2147483647:
                    var v15 bool = false
                    _ = v15
                case 32767:
                    v2 += v2
                case -2147483648:
                    var v16 int32 = 42
                    _ = v16
                default:
                    v4 = append(v4, 946)
                }
                fmt.Println(v6, fi7, v2)
            }
        }
        _ = fi7
    }
    fmt.Println(v2, v6)
    fmt.Println(v2)
    var v17 string = ""
    var v18 int32 = 0
    if false {
        v6 = (v2 + v6)
        for fi19 := 0; fi19 < 14; fi19++ {
            v2 = v6
            _ = fi19
        }
        v20 := GenS1{F0: -123, F1: -562}
        _ = v20
    } else {
        if false {
            v21 := v3[v17]
            _ = v21
        }
        v22 := GenS2{F0: v17, F1: (v6 | v6)}
        fmt.Println(v17, v6, v2)
        _ = v22
    }
    switch 1 {
    case 32768:
        for fi23 := 0; fi23 < 14; fi23++ {
            v1[0] = v17
            v24 := GenS1{F0: 32767, F1: 0}
            v18 -= -1
            _ = fi23
            _ = v24
        }
    }
    v25 := GenS1{F0: 862, F1: 42}
    switch v17 {
    case "a":
        var v26 uint64 = 1
        _ = v26
    case "hello":
        switch -9223372036854775808 {
        case 42:
            _ = len(v5)
        default:
            var v27 int16 = (v2 >> ((18446744073709551615) & 63))
            _ = v27
        }
    }
    v4 = append(v4, 2147483647)
    var v28 float64 = float64((0.1 + 1e10))
    v5 = append(v5, -2147483648)
    var v29 uint64 = 1
    for fi30 := 0; fi30 < 10; fi30++ {
        var v31 uint16 = 2
        v32 := GenS2{F0: v17, F1: v2}
        _ = fi30
        _ = v31
        _ = v32
    }
    var v33 int = 32767
    switch ("a" + v17) {
    case "\n\t":
        var v34 int64 = 9223372036854775807
        _ = v34
    default:
        v3[(v17 + v17)] = v17
    }
    v35 := GenS3{F0: 1}
    if false {
        if true {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = len(v1)
        }
        if (true || (255 != 65535)) {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v2 = (v6 & v6)
        }
        var v36 int32 = v18
        _ = v36
    } else {
        var v37 uint16 = 162
        var v38 uint16 = (^(v37))
        _ = v37
        _ = v38
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v33
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
