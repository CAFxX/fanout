package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 int; F1 int64; F2 int32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint8; F1 uint8 }

func main() {
    var v1 uint32 = 4294967295
    var v2 int32 = 0
    var v3 float64 = 1.5
    v4 := make(map[uint32]uint32)
    v5 := make(map[uint32]string)
    v1 -= (v1 >> ((1) & 63))
    var v6 int16 = 1
    switch ("\n\t" + "hello") {
    case "foo bar":
        v7 := v4[(v1 >> ((0) & 63))]
        _ = v7
    }
    if (1 == v1) {
        v6 = (v6 / ((v6) | 1))
        v2 = v2
    } else {
        var v8 int64 = 9223372036854775807
        if true {
            v9 := []int{-9223372036854775808}
            for fi10 := 0; fi10 < 16; fi10++ {
                fmt.Println(v6, v8, v1)
                fmt.Println(v8, fi10, v6)
                _ = fi10
            }
            _ = v9
        } else {
            v8 = v8
            switch 1 {
            case -9223372036854775808:
                v8 -= v8
            default:
                v11 := GenS1{F0: 1}
                _ = v11
            }
            v3 += v3
        }
        _ = v8
    }
    v12 := GenS2{F0: -9223372036854775808, F1: 2147483647}
    for fi13 := 0; fi13 < 6; fi13++ {
        switch (fi13 % ((fi13) | 1)) {
        case 913:
            if (((((42 <= 18446744073709551615) && ((255 < 196) && false)) || false) && (0 < 255)) || (2147483647 >= -191)) {
                v14 := GenS1{F0: 18446744073709551615}
                v15 := GenS3{F0: fi13, F1: 42, F2: v2}
                v16 := GenS4{F0: int32(fi13)}
                _ = v14
                _ = v15
                _ = v16
            } else {
                switch fi13 {
                case 0:
                    v2 = -954
                default:
                    if true {
                        switch ("\n\t" + (("" + ("foo bar" + ("hello" + ("hello" + "\n\t")))) + "hello")) {
                        case "\n\t":
                            v17 := []int{}
                            _ = v17
                        case "hello":
                            fmt.Println(fi13, v6, v3)
                        default:
                            v18 := GenS3{F0: 127, F1: 0, F2: (v2 % (((^(v2))) | 1))}
                            _ = v18
                        }
                        v6 = 0
                        v6 -= v6
                    }
                }
                v3 += 0.1
            }
        }
        var v19 int64 = -1
        _ = fi13
        _ = v19
    }
    v20 := GenS1{F0: 18446744073709551615}
    if (-1 > -128) {
        var v21 int8 = -1
        _ = v21
    } else {
        if (-74 > -1) {
            for fi22 := 0; fi22 < 1; fi22++ {
                for fi23 := 0; fi23 < 16; fi23++ {
                    var v24 string = ""
                    v6 = (v6 * v6)
                    if (!((v2 % ((1) | 1)) <= 441)) {
                        v25 := GenS5{F0: v24}
                        fmt.Println(v2, v24)
                        _ = v25
                    } else {
                        v26 := GenS4{F0: -2147483648}
                        _ = v26
                    }
                    _ = fi23
                    _ = v24
                }
                _ = fi22
            }
        }
        fmt.Println(v2, v3, v2)
    }
    var v27 uint32 = v1
    delete(v4, (v1 >> ((18446744073709551615) & 63)))
    var v28 uint16 = 65535
    if false {
        if true {
            v29 := []int{-37, -10}
            _ = v29
        } else {
            v6 -= v6
            v30 := []int{}
            if (((32767 != 2147483647) || (-729 >= 0)) || false) {
                if false {
                    var v31 uint8 = 1
                    fmt.Println(v3, v6, v27)
                    v4[uint32(v31)] = 135
                    _ = v31
                } else {
                    var v32 int32 = v2
                    _ = v32
                }
            } else {
                if true {
                    v33 := GenS6{F0: 127, F1: 128}
                    _ = v33
                }
            }
            _ = v30
        }
        if false {
            var v34 int32 = 1
            _ = v34
        }
    }
    var v35 float64 = (float64(0.1) * v3)
    fmt.Println(v6, v35, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v27
    _ = v28
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
