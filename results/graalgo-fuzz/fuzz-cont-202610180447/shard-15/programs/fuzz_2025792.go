package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool; F1 int; F2 int16 }

func main() {
    var v1 int = -1
    var v2 string = "hello"
    v3 := []string{v2, (v2 + v2), (v2 + v2)}
    var v4 uint32 = 716
    var v5 float64 = -2.25
    v6 := []uint32{(v4 << ((18446744073709551615) & 63)), 837, 2147483647, 0}
    if (1 <= -9223372036854775808) {
        var v7 uint64 = 774
        _ = v7
    }
    for fi8 := 0; fi8 < 8; fi8++ {
        if false {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            if (!false) {
                var v9 string = (v2 + ((v2 + v2) + "\n\t"))
                var v10 float64 = v5
                _ = v10
                _ = v9
            }
            v11 := GenS1{F0: "a", F1: 2, F2: 2}
            _ = v11
        } else {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            var v12 float64 = v5
            _ = v12
        }
        _ = fi8
    }
    if false {
        fmt.Println(v2, v5)
        switch "\n\t" {
        case "a":
            fmt.Println(v5)
        case "foo bar":
            if false {
                for fi13 := 0; fi13 < 15; fi13++ {
                    var v14 string = v2
                    _ = fi13
                    _ = v14
                }
                for fi15 := 0; fi15 < 18; fi15++ {
                    v1 = (v1 % (((-(fi15))) | 1))
                    if true {
                        var v16 int32 = -1
                        _ = v16
                    }
                    for fi17 := 0; fi17 < 11; fi17++ {
                        v3 = append(v3, v2)
                        var v18 int32 = -352
                        var v19 uint8 = 8
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                    _ = fi15
                }
                v1 -= (v1 * v1)
            } else {
                fmt.Println(v1, v2)
                v20 := GenS2{F0: (0 < 1)}
                _ = v20
            }
        default:
            fmt.Println(v2)
        }
    }
    v2 = v2
    var v21 uint64 = 916
    if (1 <= 42) {
        v1 -= -2147483648
        for fi22 := 0; fi22 < 14; fi22++ {
            if (0 >= 2) {
                v4 = (v4 / ((v4) | 1))
            } else {
                if len(v6) > 0 {
                    v6[0] = v4
                }
            }
            _ = fi22
        }
        if ((!false) && (340 < 127)) {
            if false {
                v4 = 921
            } else {
                v5 = (1e10 / v5)
                var v23 bool = (0 >= -1)
                _ = v23
            }
            var v24 uint32 = (^(v4))
            _ = v24
        } else {
            v25 := GenS2{F0: false}
            var v26 int32 = -1
            v27 := GenS3{F0: ((!true) && false), F1: v1, F2: -1}
            _ = v25
            _ = v26
            _ = v27
        }
    } else {
        var v28 int64 = -1
        v1 = (v1 >> ((42) & 63))
        var v29 bool = (v4 > 65535)
        _ = v28
        _ = v29
    }
    v21 += (v21 % ((456) | 1))
    var v30 uint64 = v21
    if false {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    var v31 uint16 = 42
    v30 = (v30 << ((4294967295) & 63))
    var v32 uint16 = v31
    if ((!(v30 >= (v21 & (^(v30))))) && false) {
        fmt.Println(v21)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
