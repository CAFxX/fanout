package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 int16 }

func main() {
    var v1 int64 = -9223372036854775808
    var v2 int32 = 2147483647
    v3 := make(map[string]int)
    v4 := make(map[int]string)
    v5 := make(map[uint32]int)
    if ((42 <= 1) || (4294967295 >= 65535)) {
        v6 := GenS1{F0: 104, F1: 65535}
        if (!(42 >= 181)) {
            if (255 != 65535) {
                var v7 uint8 = 160
                _ = v7
            }
            for fi8 := 0; fi8 < 15; fi8++ {
                v2 += 42
                _ = fi8
            }
            v1 = 842
        } else {
            v2 = v2
        }
        for fi9 := 0; fi9 < 18; fi9++ {
            v2 -= (v2 ^ v2)
            _ = fi9
        }
        _ = v6
    } else {
        switch -1 {
        case 127:
            switch 361 {
            case 975:
                var v10 float32 = 0.0
                _ = v10
            case 2147483647:
                for fi11 := 0; fi11 < 13; fi11++ {
                    fmt.Println(v2, fi11, fi11)
                    _ = fi11
                }
            default:
                v12 := GenS1{F0: -970, F1: 2}
                _ = v12
            }
        case 369:
            var v13 uint8 = 128
            _ = v13
        case -2147483648:
            if (386 != 32767) {
                v14 := []int{128, -2147483648}
                var v15 uint32 = 4294967295
                _ = v14
                _ = v15
            }
        }
    }
    for fi16 := 0; fi16 < 11; fi16++ {
        for fi17 := 0; fi17 < 4; fi17++ {
            v1 += v1
            var v18 float64 = float64(1.5)
            delete(v5, 487)
            _ = fi17
            _ = v18
        }
        v1 = v1
        _ = fi16
    }
    v1 += (v1 >> ((255) & 63))
    v1 = v1
    var v19 int16 = 286
    v1 += v1
    var v20 uint16 = 255
    if ((v20 | (^(v20))) <= v20) {
        v21 := GenS2{F0: 870, F1: 711}
        _ = v21
    } else {
        v19 -= -32768
        for fi22 := 0; fi22 < 18; fi22++ {
            if false {
                v20 += 700
            } else {
                var v23 int8 = 0
                _ = v23
            }
            v24 := GenS1{F0: fi22, F1: 2}
            for fi25 := 0; fi25 < 10; fi25++ {
                v26 := []int{-434, (fi25 & fi22), (-(fi22)), 32768}
                if len(v26) >= 2 {
                    v26 = v26[1:len(v26)]
                }
                v2 += -1
                if false {
                    v19 = (^(v19))
                    v27 := GenS3{F0: -1}
                    v28 := GenS2{F0: 42, F1: 0}
                    _ = v27
                    _ = v28
                }
                _ = fi25
                _ = v26
            }
            _ = fi22
            _ = v24
        }
    }
    v29 := []int{-308, 32767, 9223372036854775807}
    for fi30 := 0; fi30 < 6; fi30++ {
        switch -55 {
        case 32767:
            v29 = append(v29, (fi30 ^ fi30))
        case -1:
            if len(v29) >= 2 {
                v29 = v29[1:len(v29)]
            }
        default:
            for fi31 := 0; fi31 < 15; fi31++ {
                _ = len(v29)
                _ = fi31
            }
        }
        _ = fi30
    }
    fmt.Println(v1, v20)
    v20 = v20
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
