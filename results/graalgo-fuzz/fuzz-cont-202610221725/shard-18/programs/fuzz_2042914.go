package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 uint64; F2 int8 }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 uint64; F1 uint; F2 uint8 }
type GenS5 struct { F0 int32; F1 uint8 }
func genFn0(p0 uint32, p1 int64, p2 int8) int {
    for fi4 := 0; fi4 < 20; fi4++ {
        if (0 >= 0) {
            var v5 int8 = (p2 | p2)
            for fi6 := 0; fi6 < 16; fi6++ {
                var v7 uint16 = 1
                p2 += (p2 + (v5 + (p2 >> ((636) & 63))))
                v8 := GenS1{F0: 18}
                _ = fi6
                _ = v7
                _ = v8
            }
            _ = v5
        }
        v9 := GenS2{F0: (("hello" + "hello") + "a"), F1: uint64(p0), F2: p2}
        p1 += 9223372036854775807
        _ = fi4
        _ = v9
    }
    if false {
        var v10 float32 = (3.14159 + 1e10)
        if (0 < 65535) {
            switch "\n\t" {
            case "a":
                v11 := []int{}
                _ = v11
            case "":
                p0 = p0
            case "foo bar":
                if false {
                    v12 := []int{}
                    for fi13 := 0; fi13 < 11; fi13++ {
                        if len(v12) > 0 {
                            v12[0] = fi13
                        }
                        p0 += (p0 + 1)
                        _ = fi13
                    }
                    v12 = append(v12, int(p1))
                    _ = v12
                }
            default:
                p2 = p2
            }
            v14 := GenS3{F0: (-697 < -1), F1: 1}
            _ = v14
        } else {
            v15 := []int{-2147483648, 42, 42}
            _ = v15
        }
        p2 += -110
        _ = v10
    } else {
        if true {
            var v16 int32 = 2147483647
            switch "a" {
            case "hello":
                var v17 uint32 = 2147483647
                _ = v17
            }
            _ = v16
        }
    }
    p2 += -123
    for fi18 := 0; fi18 < 5; fi18++ {
        v19 := GenS1{F0: 128}
        var v20 int32 = -1
        _ = fi18
        _ = v19
        _ = v20
    }
    var v21 uint8 = 64
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    return -300
}
func genFn1() int64 {
    v22 := GenS4{F0: 0, F1: 255, F2: 242}
    var v23 uint8 = 188
    _ = v22
    _ = v23
    return 42
}

func main() {
    var v1 int64 = 1
    var v2 float64 = (1e10 / (((-2.25 + 1e10) * ((0.1 + 3.14159) / 1.0)) * 1e10))
    v3 := make(map[uint32]string)
    var v24 uint32 = 631
    v2 = v2
    switch ("a" + "\n\t") {
    case "hello":
        for fi25 := 0; fi25 < 20; fi25++ {
            v1 = genFn1()
            v26 := []int{genFn0(288, (v1 + (v1 - (v1 << ((0) & 63)))), -128), (fi25 % ((-190) | 1))}
            _ = fi25
            _ = v26
        }
    default:
        var v27 int32 = 967
        _ = v27
    }
    v28 := GenS3{F0: (!false), F1: 852}
    v29 := GenS1{F0: 31}
    for fi30 := 0; fi30 < 17; fi30++ {
        v31 := []int{(fi30 + -650), genFn0(v24, v1, -2), genFn0((v24 % (((v24 / ((v24) | 1))) | 1)), (-(v1)), 0), fi30}
        _ = fi30
        _ = v31
    }
    var v32 uint64 = 1
    var v33 string = "a"
    fmt.Println(v1)
    switch int(v24) {
    case 2147483647:
        if true {
            v1 = v1
            v34 := GenS5{F0: -2147483648, F1: 93}
            var v35 uint32 = (v24 % ((v24) | 1))
            _ = v34
            _ = v35
        }
    case 32767:
        v36 := []int{127, -2147483648}
        _ = v36
    }
    var v37 int64 = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v37
}
