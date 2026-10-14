package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 uint64 }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 int32; F1 int8; F2 string }
func genFn0(p0 uint) uint {
    var v7 uint32 = 65535
    switch ("" + "\n\t") {
    case "\n\t":
        fmt.Println(p0, p0)
    case "foo bar":
        v8 := GenS1{F0: 1000, F1: -671, F2: 70}
        _ = v8
    case "hello":
        fmt.Println(p0, v7)
    default:
        v9 := GenS2{F0: false, F1: -526}
        _ = v9
    }
    var v10 float32 = 1.5
    _ = p0
    _ = v10
    _ = v7
    return 286
}
func genFn1(p0 int16, p1 int64) uint8 {
    var v11 float32 = 0.0
    v12 := GenS2{F0: (((!false) && (892 <= 567)) || true), F1: 2147483647}
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 255
}

func main() {
    var v1 uint64 = 379
    var v2 int32 = 927
    var v3 uint = 0
    var v4 uint = (v3 / (((v3 + v3)) | 1))
    v5 := make(map[uint32]bool)
    var v6 uint8 = 71
    v13 := GenS2{F0: (65535 != 2147483647), F1: -994}
    for fi14 := 0; fi14 < 11; fi14++ {
        v15 := []int{-196, 32767, fi14, fi14}
        v1 = (v1 * (v1 * v1))
        for fi16 := 0; fi16 < 16; fi16++ {
            v6 -= v6
            var v17 uint8 = genFn1(-62, 0)
            _ = fi16
            _ = v17
        }
        _ = fi14
        _ = v15
    }
    for fi18 := 0; fi18 < 7; fi18++ {
        switch (fi18 ^ fi18) {
        case 2147483647:
            var v19 string = "\n\t"
            _ = v19
        case -2147483648:
            if (uint32(v1) == 4294967295) {
                fmt.Println(fi18, fi18, fi18)
            }
        }
        v4 = genFn0((v4 - 255))
        if (v1 < uint64(v6)) {
            var v20 int32 = 1
            _ = v20
        }
        _ = fi18
    }
    switch -550 {
    case 32768:
        if true {
            for fi21 := 0; fi21 < 6; fi21++ {
                var v22 uint = genFn0((^(v4)))
                _ = fi21
                _ = v22
            }
            v23 := GenS3{F0: 0, F1: -94, F2: (("foo bar" + "") + "a")}
            _ = v23
        } else {
            if false {
                v24 := []int{}
                v25 := GenS1{F0: 1000, F1: (v2 + 0), F2: 18446744073709551615}
                _ = v24
                _ = v25
            } else {
                switch int(v1) {
                case -449:
                    fmt.Println(v2, v6, v2)
                case 178:
                    v26 := GenS3{F0: v2, F1: -17, F2: ("" + ("hello" + ("a" + "hello")))}
                    _ = v26
                }
            }
            v3 = genFn0(v3)
        }
    case 1:
        v3 -= genFn0(v3)
    }
    var v27 float64 = 0.0
    var v28 int64 = -67
    var v29 uint32 = uint32(v1)
    fmt.Println(v2, v29, v28)
    v4 = (v4 + (v4 << (((v4 - genFn0(genFn0((v3 & v3))))) & 63)))
    v1 += v1
    switch int(v1) {
    case 1:
        var v30 bool = (v28 == v28)
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
