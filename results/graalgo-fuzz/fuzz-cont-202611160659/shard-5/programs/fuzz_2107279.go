package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int; F1 string }
type GenS5 struct { F0 bool; F1 bool }
func genFn0(p0 int, p1 uint32) int {
    v6 := []int{(p0 & (p0 / ((p0) | 1))), 196, p0, p0}
    if true {
        var v7 int32 = 0
        var v8 uint = 838
        _ = v7
        _ = v8
    } else {
        if (3.14159 <= (float32(1e10) + (3.14159 / 1.0))) {
            fmt.Println(p1)
            var v9 float32 = 0.1
            _ = v9
        }
    }
    fmt.Println(p0)
    var v10 int8 = 85
    _ = p0
    _ = p1
    _ = v10
    _ = v6
    return p0
}
func genFn1(p0 int, p1 uint32) uint8 {
    if (!false) {
        if (false || ("\n\t" >= "")) {
            var v11 int = (p0 << ((4294967295) & 63))
            p1 -= p1
            _ = v11
        } else {
            v12 := GenS1{F0: -838}
            v13 := []int{p0, 205}
            for fi14 := 0; fi14 < 3; fi14++ {
                _ = len(v13)
                fmt.Println(fi14, p1)
                _ = fi14
            }
            _ = v12
            _ = v13
        }
        if false {
            p0 += genFn0(-345, (p1 % ((42) | 1)))
            var v15 bool = true
            for fi16 := 0; fi16 < 15; fi16++ {
                v17 := GenS1{F0: 128}
                switch "" {
                case "hello":
                    for fi18 := 0; fi18 < 15; fi18++ {
                        var v19 uint64 = 633
                        var v20 uint64 = (^(v19))
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                case "":
                    switch p0 {
                    case 42:
                        v21 := GenS2{F0: 1, F1: p1}
                        _ = v21
                    case -27:
                        var v22 uint64 = 0
                        _ = v22
                    case -9223372036854775808:
                        if true {
                            for fi23 := 0; fi23 < 15; fi23++ {
                                var v24 uint64 = 225
                                _ = fi23
                                _ = v24
                            }
                        } else {
                            v25 := []int{(^(fi16)), (fi16 ^ fi16), p0}
                            _ = v25
                        }
                    }
                default:
                    fmt.Println(p1, v15, p0)
                }
                _ = fi16
                _ = v17
            }
            _ = v15
        } else {
            fmt.Println(p0)
        }
    }
    v26 := GenS1{F0: -586}
    var v27 int8 = 100
    v28 := []int{p0}
    _ = p0
    _ = p1
    _ = v26
    _ = v27
    _ = v28
    return 127
}

func main() {
    var v1 float32 = ((1e10 / -2.25) - (1e10 - (0.1 + float32(0.1))))
    v2 := make(map[uint32]string)
    var v3 int8 = -128
    var v4 uint8 = 255
    var v5 int32 = -119
    if (("hello" < (("hello" + "foo bar") + ("foo bar" + "a"))) && (true || (-1 == 565))) {
        v29 := GenS2{F0: 1, F1: 2}
        _ = v29
    } else {
        v30 := GenS1{F0: -1}
        v31 := GenS3{F0: -471}
        _ = v30
        _ = v31
    }
    fmt.Println(v4, v1)
    var v32 bool = false
    v5 = (v5 ^ (v5 % ((714) | 1)))
    switch -9223372036854775808 {
    case 877:
        for fi33 := 0; fi33 < 12; fi33++ {
            fmt.Println(v5)
            _ = fi33
        }
    case 490:
        if v32 {
            for fi34 := 0; fi34 < 11; fi34++ {
                v35 := []int{fi34}
                var v36 bool = v32
                v37 := GenS4{F0: genFn0(fi34, 626), F1: ("foo bar" + "\n\t")}
                _ = fi34
                _ = v35
                _ = v36
                _ = v37
            }
        }
    case 32767:
        fmt.Println(v3, v32, v1)
    }
    var v38 string = ""
    v39 := GenS5{F0: v32, F1: false}
    v40 := []int{int(v4), 2147483647, -990, -1}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
