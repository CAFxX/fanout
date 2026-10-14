package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 int8; F1 bool; F2 uint16 }
type GenS3 struct { F0 uint32; F1 int }
type GenS4 struct { F0 int; F1 uint16 }
type GenS5 struct { F0 uint16; F1 uint8; F2 uint32 }
func genFn0() int16 {
    var v4 int32 = -177
    for fi5 := 0; fi5 < 18; fi5++ {
        v6 := GenS1{F0: (!((-128 >= -128) || false)), F1: v4}
        v7 := GenS2{F0: -1, F1: true, F2: 255}
        _ = fi5
        _ = v6
        _ = v7
    }
    if true {
        switch 32767 {
        case 4294967295:
            var v8 string = "hello"
            _ = v8
        case -2147483648:
            v4 -= -1
        }
        v4 += 2147483647
        v4 -= -748
    } else {
        for fi9 := 0; fi9 < 2; fi9++ {
            v4 -= (v4 << ((18446744073709551615) & 63))
            v4 += 2147483647
            v10 := GenS1{F0: true, F1: 0}
            _ = fi9
            _ = v10
        }
        v11 := GenS3{F0: 797, F1: -856}
        _ = v11
    }
    fmt.Println(v4)
    fmt.Println(v4)
    _ = v4
    return -973
}
func genFn1(p0 int32, p1 int8) uint64 {
    for fi12 := 0; fi12 < 19; fi12++ {
        var v13 string = "foo bar"
        v13 = (v13 + v13)
        v14 := []int{}
        _ = fi12
        _ = v13
        _ = v14
    }
    p0 = -485
    _ = p0
    _ = p1
    return 1
}

func main() {
    v1 := []uint32{116, 4294967295, 2}
    var v2 int8 = 127
    v3 := make(map[uint32]string)
    v15 := GenS4{F0: -9223372036854775808, F1: 65535}
    var v16 uint64 = 42
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi17 := 0; fi17 < 6; fi17++ {
        var v18 uint = 0
        var v19 int16 = -762
        _ = len(v1)
        _ = fi17
        _ = v18
        _ = v19
    }
    _ = len(v3)
    delete(v3, 301)
    for fi20 := 0; fi20 < 11; fi20++ {
        v2 = 123
        switch "a" {
        case "a":
            v21 := GenS3{F0: 42, F1: (^(fi20))}
            _ = v21
        default:
            switch (fi20 << ((123) & 63)) {
            case 2147483647:
                for fi22 := 0; fi22 < 15; fi22++ {
                    v2 = (v2 >> ((0) & 63))
                    _ = fi22
                }
            case 0:
                if false {
                    v23 := GenS5{F0: 42, F1: 1, F2: 8}
                    _ = v23
                }
            case 800:
                if false {
                    v2 = (v2 - v2)
                    var v24 uint = 255
                    var v25 int32 = 42
                    _ = v24
                    _ = v25
                }
            default:
                fmt.Println(fi20)
            }
        }
        _ = fi20
    }
    if (float64(1.5) != ((3.14159 / -2.25) * (-2.25 - 0.1))) {
        if true {
            switch 127 {
            case -9223372036854775808:
                if false {
                    var v26 int8 = v2
                    _ = v26
                }
            case 81:
                for fi27 := 0; fi27 < 2; fi27++ {
                    var v28 uint = 634
                    fmt.Println(v16)
                    _ = fi27
                    _ = v28
                }
            default:
                v2 -= -33
            }
            v16 = v16
            switch 4294967295 {
            case 661:
                v16 = 4294967295
            case -2147483648:
                var v29 int8 = (v2 & v2)
                _ = v29
            default:
                for fi30 := 0; fi30 < 15; fi30++ {
                    for fi31 := 0; fi31 < 11; fi31++ {
                        fmt.Println(fi31, fi30)
                        var v32 uint16 = 695
                        _ = fi31
                        _ = v32
                    }
                    _ = fi30
                }
            }
        } else {
            if ((4294967295 <= 261) && false) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        }
        for fi33 := 0; fi33 < 1; fi33++ {
            v16 = v16
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi33
        }
        if false {
            v16 = genFn1(-1, v2)
        }
    }
    v2 -= 42
    var v34 int8 = v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v34
}
