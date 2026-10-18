package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint }
type GenS2 struct { F0 uint8; F1 uint64; F2 int8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint64; F1 bool; F2 uint }
type GenS5 struct { F0 int; F1 int32 }
func genFn0() int16 {
    var v5 int64 = 42
    fmt.Println(v5)
    _ = v5
    return -17
}

func main() {
    var v1 int64 = 1
    v2 := make(map[int]uint32)
    var v3 uint16 = 705
    var v4 int = 283
    if true {
        v6 := v2[820]
        if true {
            for fi7 := 0; fi7 < 13; fi7++ {
                var v8 uint64 = 777
                v6 = uint32(v8)
                _ = fi7
                _ = v8
            }
            v9 := GenS1{F0: 2147483648, F1: 1, F2: 811}
            v1 -= int64(v6)
            _ = v9
        }
        for fi10 := 0; fi10 < 19; fi10++ {
            var v11 int16 = genFn0()
            v4 = v4
            for fi12 := 0; fi12 < 13; fi12++ {
                switch v4 {
                case 184:
                    v1 -= (v1 / (((v1 | 1)) | 1))
                case -9223372036854775808:
                    v11 -= int16(v4)
                default:
                    v4 = fi12
                }
                _ = fi12
            }
            _ = fi10
            _ = v11
        }
        _ = v6
    }
    for fi13 := 0; fi13 < 1; fi13++ {
        fmt.Println(v1)
        switch "foo bar" {
        case "a":
            v14 := GenS1{F0: 559, F1: -1, F2: 465}
            _ = v14
        }
        if (false || (65535 >= 18446744073709551615)) {
            var v15 int = int(v3)
            if (v15 > fi13) {
                fmt.Println(v4)
                v1 = (v1 ^ (v1 & (v1 / (((v1 + v1)) | 1))))
            }
            var v16 uint32 = 0
            _ = v15
            _ = v16
        } else {
            v17 := GenS2{F0: 236, F1: 2, F2: -1}
            var v18 bool = false
            for fi19 := 0; fi19 < 13; fi19++ {
                if v18 {
                    var v20 int = 32768
                    _ = v20
                }
                v21 := []int{fi13}
                v22 := GenS3{F0: v3}
                _ = fi19
                _ = v21
                _ = v22
            }
            _ = v17
            _ = v18
        }
        _ = fi13
    }
    v3 = (v3 * 42)
    for fi23 := 0; fi23 < 6; fi23++ {
        switch (-(v4)) {
        case 42:
            var v24 uint = 255
            _ = v24
        default:
            v3 += (v3 / (((^(v3))) | 1))
        }
        v1 = 42
        if true {
            var v25 int = 32768
            _ = v25
        }
        _ = fi23
    }
    fmt.Println(v3, v1, v3)
    switch v4 {
    case 0:
        switch ("" + (("a" + "hello") + ("foo bar" + ""))) {
        case "foo bar":
            fmt.Println(v3)
        default:
            for fi26 := 0; fi26 < 7; fi26++ {
                fmt.Println(fi26)
                _ = fi26
            }
        }
    default:
        v27 := GenS1{F0: 1, F1: 0, F2: 380}
        _ = v27
    }
    v2[v4] = 0
    v28 := []int{(v4 * v4)}
    for fi29 := 0; fi29 < 1; fi29++ {
        switch 44 {
        case 32768:
            var v30 int16 = 0
            _ = v30
        }
        switch "hello" {
        case "\n\t":
            for fi31 := 0; fi31 < 15; fi31++ {
                v32 := GenS3{F0: v3}
                var v33 uint16 = 65535
                _ = fi31
                _ = v32
                _ = v33
            }
        case "hello":
            v1 = v1
        case "foo bar":
            v1 += v1
        }
        _ = fi29
    }
    delete(v2, v4)
    v1 = (v1 % ((-724) | 1))
    delete(v2, v4)
    for fi34 := 0; fi34 < 12; fi34++ {
        fmt.Println(v3)
        _ = fi34
    }
    v35 := GenS4{F0: 18446744073709551615, F1: (v1 == (v1 * (v1 * (^(v1))))), F2: 65535}
    _ = len(v2)
    var v36 int = 0
    v37 := GenS5{F0: v4, F1: 2147483647}
    var v38 uint32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
}
