package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64; F1 int64; F2 string }
type GenS5 struct { F0 uint; F1 bool; F2 uint32 }

func main() {
    var v1 string = ("hello" + "\n\t")
    var v2 int16 = 0
    var v3 int8 = 1
    var v4 float64 = 3.14159
    var v5 uint8 = 10
    var v6 int64 = 42
    v7 := GenS1{F0: 4294967295}
    for fi8 := 0; fi8 < 18; fi8++ {
        v9 := GenS2{F0: (v2 >= v2)}
        v10 := GenS2{F0: false}
        v11 := []int{(-(fi8)), int(v3), fi8, fi8}
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    v4 += 1e10
    v1 = "a"
    for fi12 := 0; fi12 < 12; fi12++ {
        var v13 int16 = 1000
        var v14 uint32 = 4294967295
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 int64 = (v6 << ((18446744073709551615) & 63))
    switch 32768 {
    case 47:
        if (int64(v3) == v6) {
            if true {
                var v16 uint16 = 65535
                v2 -= (-(v2))
                v17 := []int{32768, -1}
                _ = v16
                _ = v17
            } else {
                v1 = (v1 + ((("hello" + "") + "\n\t") + "a"))
                switch v1 {
                case "":
                    if ((false || ((-524 != -426) && (false && (0 >= 2)))) && (-32 != v3)) {
                        for fi18 := 0; fi18 < 17; fi18++ {
                            switch fi18 {
                            case 0:
                                v4 += 0.0
                            }
                            var v19 uint64 = 970
                            _ = fi18
                            _ = v19
                        }
                        v15 = 9223372036854775807
                        v3 += (^(v3))
                    }
                case "\n\t":
                    switch v1 {
                    case "hello":
                        for fi20 := 0; fi20 < 18; fi20++ {
                            v15 = (-(v6))
                            v4 = 3.14159
                            switch 32768 {
                            case -43:
                                var v21 int8 = v3
                                _ = v21
                            case 724:
                                v22 := GenS2{F0: true}
                                _ = v22
                            default:
                                var v23 int64 = (-(v15))
                                _ = v23
                            }
                            _ = fi20
                        }
                    default:
                        var v24 uint = 2
                        _ = v24
                    }
                case "foo bar":
                    v25 := []int{}
                    _ = v25
                }
            }
            v1 = v1
        } else {
            if (!true) {
                v26 := GenS3{F0: v2}
                _ = v26
            }
            var v27 int8 = (v3 & v3)
            _ = v27
        }
    }
    v28 := GenS4{F0: -815, F1: (-(v6)), F2: v1}
    if (v15 == (v6 % (((^(v6))) | 1))) {
        v15 += -115
        var v29 bool = true
        _ = v29
    } else {
        for fi30 := 0; fi30 < 16; fi30++ {
            v31 := GenS5{F0: 2, F1: false, F2: 42}
            _ = fi30
            _ = v31
        }
        var v32 float64 = float64(3.14159)
        var v33 int16 = v2
        _ = v32
        _ = v33
    }
    for fi34 := 0; fi34 < 6; fi34++ {
        fmt.Println(fi34)
        fmt.Println(v5)
        v1 = (((v1 + v1) + v1) + ((v1 + v1) + "foo bar"))
        _ = fi34
    }
    v5 += (^(v5))
    fmt.Println(v5, v1, v6)
    v35 := GenS5{F0: 2, F1: false, F2: 2147483647}
    v36 := GenS2{F0: false}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v28
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
