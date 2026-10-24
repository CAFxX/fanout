package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 uint; F1 int16; F2 uint16 }
type GenS4 struct { F0 uint64 }
func genFn0(p0 int) uint {
    if true {
        v4 := []int{-729}
        _ = v4
    }
    var v5 int16 = 1000
    _ = p0
    _ = v5
    return 2
}

func main() {
    var v1 int8 = 42
    v2 := make(map[int]int)
    v3 := []string{"\n\t", "\n\t", "", ""}
    v1 = 100
    switch 1 {
    case 32768:
        if len(v3) > 0 {
            v3[0] = ((("\n\t" + ("" + ("hello" + ""))) + "\n\t") + ("" + ("\n\t" + "hello")))
        }
    }
    for fi6 := 0; fi6 < 6; fi6++ {
        var v7 uint32 = 0
        _ = fi6
        _ = v7
    }
    if ((1e10 - -2.25) != 1e10) {
        v8 := GenS1{F0: 4294967295, F1: v1}
        switch -750 {
        case 4294967295:
            if (42 >= 0) {
                switch "a" {
                case "hello":
                    delete(v2, 4294967295)
                case "a":
                    var v9 uint16 = 747
                    _ = v9
                case "":
                    var v10 uint32 = 420
                    _ = v10
                default:
                    var v11 uint64 = 4294967295
                    _ = v11
                }
                for fi12 := 0; fi12 < 16; fi12++ {
                    var v13 int16 = -242
                    fmt.Println(v13)
                    _ = fi12
                    _ = v13
                }
            } else {
                v14 := GenS2{F0: 1, F1: true}
                if len(v3) > 0 {
                    v3[0] = ""
                }
                if false {
                    for fi15 := 0; fi15 < 7; fi15++ {
                        if false {
                            v16 := GenS2{F0: 2, F1: true}
                            delete(v2, fi15)
                            v1 = v1
                            _ = v16
                        }
                        v17 := GenS3{F0: 763, F1: 866, F2: 257}
                        v1 = v1
                        _ = fi15
                        _ = v17
                    }
                    v18 := GenS2{F0: 0, F1: (!false)}
                    var v19 uint8 = 42
                    _ = v18
                    _ = v19
                } else {
                    v1 = v1
                }
                _ = v14
            }
        case -9223372036854775808:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        default:
            if false {
                for fi20 := 0; fi20 < 3; fi20++ {
                    v21 := v2[(fi20 & fi20)]
                    _ = fi20
                    _ = v21
                }
                if (1 != 65535) {
                    v1 -= (v1 / ((v1) | 1))
                    if true {
                        var v22 uint = 1
                        fmt.Println(v1, v1)
                        _ = v22
                    }
                }
            } else {
                for fi23 := 0; fi23 < 10; fi23++ {
                    if (!false) {
                        v1 = (v1 >> ((0) & 63))
                    }
                    _ = fi23
                }
                v2[-2147483648] = -166
            }
        }
        v1 = v1
        _ = v8
    } else {
        v24 := GenS3{F0: 42, F1: -390, F2: 42}
        _ = v24
    }
    v1 = v1
    var v25 int64 = 543
    fmt.Println(v1, v25)
    var v26 int = 541
    for fi27 := 0; fi27 < 2; fi27++ {
        if true {
            var v28 uint32 = 1
            var v29 uint16 = 1
            _ = v28
            _ = v29
        } else {
            v26 -= (fi27 + fi27)
            v30 := GenS4{F0: 4294967295}
            v31 := GenS4{F0: 0}
            _ = v30
            _ = v31
        }
        v25 -= (v25 >> ((18446744073709551615) & 63))
        _ = fi27
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
}
