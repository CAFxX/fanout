package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 uint32; F1 uint32; F2 string }
type GenS3 struct { F0 int8; F1 bool; F2 uint64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint }
func genFn0() int32 {
    var v7 uint = 2
    v8 := GenS1{F0: 1, F1: (32767 < 32767)}
    if (false && (-684 <= 32767)) {
        fmt.Println(v7)
    }
    _ = v7
    _ = v8
    return 42
}
func genFn1() bool {
    if (42 != 65535) {
        var v9 float32 = 1e10
        for fi10 := 0; fi10 < 8; fi10++ {
            var v11 int8 = 1
            var v12 uint = 18446744073709551615
            v13 := GenS2{F0: 903, F1: 4294967295, F2: "a"}
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
        _ = v9
    }
    v14 := []int{190, 9223372036854775807}
    _ = v14
    return (!((!(157 != 646)) && false))
}

func main() {
    var v1 int64 = -1
    var v2 string = ("a" + "a")
    var v3 int32 = 270
    var v4 int32 = v3
    var v5 bool = false
    var v6 uint8 = 42
    v15 := GenS3{F0: 0, F1: (!false), F2: 18446744073709551615}
    var v16 float64 = float64((3.14159 - 0.0))
    v17 := []int{41, -2147483648, -2147483648}
    for fi18 := 0; fi18 < 10; fi18++ {
        switch (fi18 % (((fi18 ^ (fi18 << ((65535) & 63)))) | 1)) {
        case -2147483648:
            v16 = (float64(0.1) * ((v16 * ((v16 * v16) + v16)) + (0.0 * ((float64((-2.25 / float32(float64(1.5)))) + float64(float32(1.5))) * v16))))
        case 0:
            switch 92 {
            case 681:
                v1 += v1
            case 32767:
                var v19 int32 = v3
                _ = v19
            default:
                var v20 int = int(v6)
                _ = v20
            }
        default:
            var v21 uint32 = 4294967295
            _ = v21
        }
        _ = fi18
    }
    v1 = (v1 >> ((uint(v6)) & 63))
    if ((!((0 > 218) || (!(((v2 + v2) == v2) || v5)))) && v5) {
        for fi22 := 0; fi22 < 14; fi22++ {
            var v23 uint16 = 32767
            var v24 int = -9223372036854775808
            v23 -= v23
            _ = fi22
            _ = v23
            _ = v24
        }
    }
    v25 := GenS4{F0: -2147483648}
    var v26 uint = 441
    for fi27 := 0; fi27 < 9; fi27++ {
        if (2147483648 >= 1) {
            for fi28 := 0; fi28 < 10; fi28++ {
                if (v4 != (v3 ^ (v4 * int32(v6)))) {
                    var v29 int8 = -40
                    _ = v29
                }
                switch fi28 {
                case -9223372036854775808:
                    _ = len(v17)
                case 897:
                    var v30 uint = (v26 % ((v26) | 1))
                    _ = v30
                case 42:
                    v31 := GenS5{F0: (^(v26))}
                    _ = v31
                default:
                    if (v5 || (!(fi27 >= fi28))) {
                        v2 = v2
                        v4 += (v4 + v3)
                        if v5 {
                            if len(v17) >= 2 {
                                v17 = v17[1:len(v17)]
                            }
                        } else {
                            fmt.Println(v16, fi27)
                            var v32 float64 = v16
                            var v33 int8 = 100
                            var v34 uint64 = 18446744073709551615
                            v5 = v5
                            var v35 uint32 = 42
                            v32 = float64(1e10)
                            _ = v32
                            _ = v33
                            _ = v34
                            _ = v35
                        }
                        v26 = v26
                    } else {
                        v3 = v4
                    }
                }
                _ = fi28
            }
            var v36 uint16 = 0
            _ = v36
        }
        _ = fi27
    }
    _ = len(v17)
    for fi37 := 0; fi37 < 9; fi37++ {
        for fi38 := 0; fi38 < 7; fi38++ {
            _ = len(v17)
            fmt.Println(fi37, v3, v3)
            v5 = (true && (!(652 > 0)))
            if len(v17) >= 2 {
                v17 = v17[1:len(v17)]
            }
            _ = fi38
        }
        _ = fi37
    }
    v1 = -967
    _ = len(v17)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
