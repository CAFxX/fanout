package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 uint32 }
type GenS2 struct { F0 int8; F1 uint; F2 string }
type GenS3 struct { F0 int64; F1 string; F2 int64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint32; F1 int; F2 int64 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int32; F1 int32; F2 uint64 }

func main() {
    var v1 int16 = 32767
    var v2 int64 = 329
    var v3 bool = (!((-(v2)) == v2))
    var v4 int8 = int8(v2)
    var v5 int = 9223372036854775807
    v6 := make(map[uint32]bool)
    v4 -= (v4 | v4)
    v4 += (-(v4))
    fmt.Println(v3)
    var v7 int32 = 2147483647
    switch (v5 * v5) {
    case 32767:
        v4 -= v4
    default:
        v7 = (-(v7))
    }
    for fi8 := 0; fi8 < 9; fi8++ {
        fmt.Println(v1)
        if false {
            v9 := GenS1{F0: 18446744073709551615, F1: 239, F2: 1}
            v7 -= (v7 * int32(v1))
            _ = v9
        } else {
            if true {
                switch "hello" {
                case "a":
                    v10 := GenS2{F0: 21, F1: 4294967295, F2: "\n\t"}
                    _ = v10
                default:
                    fmt.Println(v2)
                }
                switch (fi8 + 207) {
                case 633:
                    for fi11 := 0; fi11 < 4; fi11++ {
                        if (float32(-2.25) <= 3.14159) {
                            var v12 string = ""
                            v13 := GenS3{F0: (v2 * v2), F1: v12, F2: v2}
                            v4 -= (-(v4))
                            _ = v12
                            _ = v13
                        } else {
                            var v14 uint64 = 0
                            _ = v14
                        }
                        v15 := GenS4{F0: 255}
                        _ = fi11
                        _ = v15
                    }
                case 32768:
                    for fi16 := 0; fi16 < 20; fi16++ {
                        var v17 uint = 0
                        v6[4294967295] = true
                        v18 := []int{(fi16 << (((v17 - 18446744073709551615)) & 63)), fi8}
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                }
                v3 = (-1 > (v7 * (^(v7))))
            }
            for fi19 := 0; fi19 < 20; fi19++ {
                v5 = (^(fi19))
                v2 -= 398
                _ = fi19
            }
            var v20 int32 = 0
            _ = v20
        }
        v3 = (468 < 0)
        _ = fi8
    }
    for fi21 := 0; fi21 < 12; fi21++ {
        for fi22 := 0; fi22 < 6; fi22++ {
            var v23 int16 = (v1 & v1)
            if false {
                v24 := GenS5{F0: 65535, F1: fi22, F2: (v2 & v2)}
                switch (fi21 / (((^(fi21))) | 1)) {
                case 42:
                    if (42 > 2) {
                        _ = len(v6)
                        if (((v1 < v23) || ((v2 + 1) < (v2 | (v2 << ((861) & 63))))) || v3) {
                            v25 := []int{(fi22 & 202), int(v7), (v5 << ((255) & 63)), (v5 - fi22)}
                            v3 = (!(fi22 <= v5))
                            fmt.Println(v5, v2)
                            _ = v25
                        } else {
                            v23 -= v23
                            v26 := GenS6{F0: 42}
                            _ = v26
                        }
                    } else {
                        switch ("hello" + "\n\t") {
                        case "foo bar":
                            v1 += (v23 * (v1 >> ((18446744073709551615) & 63)))
                        case "":
                            v3 = ((v7 - (v7 / ((v7) | 1))) != (v7 >> ((2) & 63)))
                        case "\n\t":
                            v7 -= 1
                        default:
                            var v27 string = ("a" + "foo bar")
                            var v28 float64 = (-2.25 * float64(-2.25))
                            var v29 int32 = v7
                            _ = v27
                            _ = v28
                            _ = v29
                        }
                        v7 = (v7 * v7)
                        v23 = v1
                    }
                }
                v4 -= 127
                _ = v24
            } else {
                v2 += (v2 / ((42) | 1))
            }
            v30 := GenS7{F0: int32(v23), F1: 492, F2: 204}
            v31 := v6[38]
            _ = fi22
            _ = v23
            _ = v30
            _ = v31
        }
        v32 := []int{(v5 - 2147483647), (v5 | (fi21 - 42)), 0, (v5 + (v5 << ((308) & 63)))}
        _ = fi21
        _ = v32
    }
    switch "hello" {
    case "":
        for fi33 := 0; fi33 < 8; fi33++ {
            v2 = -1
            v34 := GenS2{F0: v4, F1: 18446744073709551615, F2: "hello"}
            _ = fi33
            _ = v34
        }
    default:
        v1 = (-(v1))
    }
    v35 := GenS6{F0: 426}
    var v36 int8 = (-(v4))
    v37 := GenS7{F0: 0, F1: v7, F2: 4294967295}
    v36 = -126
    var v38 uint16 = 32767
    fmt.Println(v3)
    v3 = v3
    v39 := []int{2147483647, (^(v5))}
    switch int(v2) {
    case -1:
        var v40 float64 = float64((-2.25 + 0.1))
        _ = v40
    default:
        v6[42] = ((v7 * (v7 >> ((42) & 63))) <= (v7 ^ v7))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
