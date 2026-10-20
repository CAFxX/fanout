package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 int }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 uint; F1 int8; F2 uint32 }
type GenS4 struct { F0 int8; F1 string }
type GenS5 struct { F0 uint16; F1 int; F2 int8 }
type GenS6 struct { F0 int32 }
func genFn0(p0 int) uint8 {
    if true {
        if (902 == 323) {
            p0 += p0
            p0 = -9223372036854775808
        } else {
            if ((715 <= -2147483648) || ((618 > 826) || ("a" == "a"))) {
                for fi4 := 0; fi4 < 4; fi4++ {
                    v5 := []int{fi4}
                    for fi6 := 0; fi6 < 4; fi6++ {
                        v7 := GenS1{F0: fi4, F1: 42, F2: 817}
                        p0 = p0
                        switch 128 {
                        case -1:
                            p0 -= (p0 / (((fi4 % (((p0 / ((p0) | 1))) | 1))) | 1))
                        case 438:
                            v8 := GenS2{F0: 100, F1: 2}
                            _ = v8
                        }
                        _ = fi6
                        _ = v7
                    }
                    _ = fi4
                    _ = v5
                }
                for fi9 := 0; fi9 < 15; fi9++ {
                    var v10 uint = 673
                    switch "a" {
                    case "a":
                        v10 = (v10 << ((315) & 63))
                    default:
                        var v11 uint8 = 127
                        _ = v11
                    }
                    _ = fi9
                    _ = v10
                }
                p0 = (p0 >> ((64) & 63))
            } else {
                var v12 uint32 = 60
                var v13 int32 = 1
                var v14 uint8 = 71
                _ = v12
                _ = v13
                _ = v14
            }
            p0 -= -377
        }
        var v15 uint8 = 255
        var v16 uint8 = 127
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 4; fi17++ {
        fmt.Println(fi17)
        v18 := GenS3{F0: 124, F1: 42, F2: 2}
        _ = fi17
        _ = v18
    }
    var v19 uint8 = 144
    _ = p0
    _ = v19
    return v19
}

func main() {
    var v1 uint32 = 632
    v2 := make(map[int]uint32)
    var v3 int32 = -685
    for fi20 := 0; fi20 < 6; fi20++ {
        v21 := GenS4{F0: 127, F1: (("" + ("a" + "hello")) + ("" + "hello"))}
        var v22 int16 = 32767
        _ = fi20
        _ = v21
        _ = v22
    }
    var v23 int16 = -32768
    v1 = v1
    var v24 string = "\n\t"
    delete(v2, 2147483647)
    for fi25 := 0; fi25 < 2; fi25++ {
        v26 := []int{}
        var v27 uint32 = (v1 + (v1 + v1))
        v23 = v23
        _ = fi25
        _ = v26
        _ = v27
    }
    if true {
        v23 = (v23 % (((v23 % ((v23) | 1))) | 1))
        var v28 string = v24
        _ = v28
    }
    v23 -= v23
    var v29 int32 = (v3 & 2147483647)
    v24 = "hello"
    for fi30 := 0; fi30 < 8; fi30++ {
        for fi31 := 0; fi31 < 7; fi31++ {
            if (0 == 65535) {
                fmt.Println(v3)
            }
            _ = fi31
        }
        v23 = v23
        _ = fi30
    }
    for fi32 := 0; fi32 < 17; fi32++ {
        v3 = v29
        switch 42 {
        case 481:
            var v33 float64 = 3.14159
            _ = v33
        }
        switch "hello" {
        case "hello":
            var v34 int64 = -1
            _ = v34
        case "":
            if false {
                v35 := GenS5{F0: 2, F1: fi32, F2: 84}
                if (!(!((9223372036854775807 == 4294967295) || true))) {
                    v36 := v2[(fi32 - fi32)]
                    fmt.Println(v24, v29)
                    _ = v36
                }
                _ = v35
            }
        case "a":
            var v37 int16 = -32768
            _ = v37
        }
        _ = fi32
    }
    var v38 float32 = 1.5
    var v39 uint16 = 32767
    delete(v2, -9223372036854775808)
    v1 = 2147483647
    switch 2147483647 {
    case 0:
        for fi40 := 0; fi40 < 10; fi40++ {
            v41 := GenS3{F0: 255, F1: 1, F2: (v1 % (((v1 % ((v1) | 1))) | 1))}
            v42 := GenS6{F0: 2147483647}
            fmt.Println(v38)
            _ = fi40
            _ = v41
            _ = v42
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v29
    _ = v3
    _ = v38
    _ = v39
}
