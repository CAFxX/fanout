package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint8 }
func genFn0() int16 {
    var v6 int = -9223372036854775808
    if false {
        v7 := GenS1{F0: 42}
        v8 := []int{(^(v6)), v6}
        v9 := GenS2{F0: 255, F1: 4294967295}
        _ = v7
        _ = v8
        _ = v9
    }
    if true {
        v6 = 128
        for fi10 := 0; fi10 < 17; fi10++ {
            for fi11 := 0; fi11 < 20; fi11++ {
                fmt.Println(fi10, v6)
                v12 := []int{-438, (fi10 - (v6 - 128))}
                var v13 int64 = -433
                _ = fi11
                _ = v12
                _ = v13
            }
            for fi14 := 0; fi14 < 2; fi14++ {
                switch "" {
                case "\n\t":
                    if false {
                        var v15 string = ""
                        _ = v15
                    }
                }
                fmt.Println(fi10, fi10, v6)
                switch fi10 {
                case -451:
                    v16 := GenS1{F0: 1}
                    _ = v16
                case 2147483647:
                    if true {
                        v17 := []int{fi14}
                        v17 = append(v17, (-(fi10)))
                        var v18 string = ""
                        _ = v17
                        _ = v18
                    } else {
                        v6 += fi14
                        fmt.Println(fi14, fi10)
                        if (48 <= 100) {
                            v19 := GenS3{F0: (true || true)}
                            _ = v19
                        } else {
                            fmt.Println(fi10, v6, v6)
                            v6 = fi10
                        }
                    }
                default:
                    if (255 == 128) {
                        var v20 int = 127
                        _ = v20
                    }
                }
                _ = fi14
            }
            _ = fi10
        }
    }
    v21 := []int{220, (^(v6)), v6, (-(v6))}
    _ = v21
    _ = v6
    return 42
}

func main() {
    var v1 float32 = -2.25
    v2 := []uint32{42, 65535}
    v3 := []int{127, 2147483647}
    var v4 int16 = 0
    v5 := []uint32{42}
    for fi22 := 0; fi22 < 14; fi22++ {
        switch ("a" + "hello") {
        case "":
            for fi23 := 0; fi23 < 7; fi23++ {
                var v24 uint16 = 409
                v24 -= (v24 >> ((581) & 63))
                fmt.Println(v4, v24)
                var v25 uint8 = 42
                _ = fi23
                _ = v24
                _ = v25
            }
        case "foo bar":
            fmt.Println(fi22, v4)
        case "hello":
            fmt.Println(fi22, fi22, fi22)
        default:
            v1 += v1
        }
        var v26 uint8 = 127
        _ = fi22
        _ = v26
    }
    v27 := GenS4{F0: ("a" + "foo bar")}
    var v28 int64 = 993
    fmt.Println(v1, v28)
    var v29 uint16 = 938
    v30 := GenS5{F0: 0}
    v28 += v28
    var v31 int = 9223372036854775807
    fmt.Println(v29, v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
