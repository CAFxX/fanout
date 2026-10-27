package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 int }
type GenS2 struct { F0 bool; F1 int; F2 uint64 }
type GenS3 struct { F0 uint }
func genFn0(p0 bool) int16 {
    if (-9223372036854775808 > 9223372036854775807) {
        var v5 uint8 = 0
        fmt.Println(p0, p0)
        _ = v5
    }
    v6 := []int{678, 9223372036854775807, -506}
    _ = p0
    _ = v6
    return -382
}
func genFn1(p0 uint64) bool {
    fmt.Println(p0)
    fmt.Println(p0)
    _ = p0
    return (!false)
}

func main() {
    var v1 uint32 = 4294967295
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]int)
    for fi7 := 0; fi7 < 18; fi7++ {
        fmt.Println(fi7, v1)
        var v8 bool = true
        _ = fi7
        _ = v8
    }
    switch "hello" {
    case "foo bar":
        v1 -= v1
    default:
        v9 := GenS1{F0: 0, F1: true, F2: 1}
        _ = v9
    }
    if (!(true && ("\n\t" != "foo bar"))) {
        fmt.Println(v1)
    } else {
        v1 -= (v1 - v1)
    }
    if (true && (true || (false || ((true || false) || true)))) {
        delete(v4, 0)
        v10 := GenS2{F0: (3.14159 >= ((0.0 / 1.0) / 1.0)), F1: 623, F2: 18446744073709551615}
        v1 -= v1
        _ = v10
    } else {
        _ = len(v4)
        v11 := []int{4294967295, 128}
        var v12 string = "a"
        _ = v11
        _ = v12
    }
    fmt.Println(v1)
    v2[(v1 << ((1) & 63))] = (v1 < 188)
    switch 128 {
    case 769:
        v1 -= v1
    case 4294967295:
        for fi13 := 0; fi13 < 6; fi13++ {
            if ((4294967295 <= 2) || false) {
                for fi14 := 0; fi14 < 18; fi14++ {
                    delete(v3, v1)
                    _ = fi14
                }
                var v15 int16 = 1
                v16 := v4[128]
                _ = v15
                _ = v16
            } else {
                v1 = (v1 - (v1 * (v1 & (v1 % ((42) | 1)))))
            }
            for fi17 := 0; fi17 < 1; fi17++ {
                v1 = (v1 ^ 65535)
                _ = fi17
            }
            _ = fi13
        }
    case 9223372036854775807:
        var v18 bool = (!((((float64((1e10 + 3.14159)) * 0.0) * 0.1) == ((1.5 / 1.0) / 1.0)) || ((0.0 / ((0.0 - (0.0 - 1.5)) - 0.1)) >= (3.14159 * float32((1e10 - 1e10))))))
        _ = v18
    default:
        for fi19 := 0; fi19 < 13; fi19++ {
            v20 := []int{}
            _ = fi19
            _ = v20
        }
    }
    var v21 string = ("foo bar" + (("hello" + (("hello" + "hello") + "hello")) + (("\n\t" + "hello") + ("" + ("\n\t" + "")))))
    v22 := v3[42]
    for fi23 := 0; fi23 < 1; fi23++ {
        v24 := GenS3{F0: 839}
        v22 = (^(v22))
        _ = fi23
        _ = v24
    }
    switch (("\n\t" + (((v21 + ("hello" + v21)) + ("hello" + v21)) + v21)) + v21) {
    case "a":
        v21 = v21
    case "":
        if true {
            v1 -= 533
        }
    default:
        v21 = (v21 + v21)
    }
    for fi25 := 0; fi25 < 7; fi25++ {
        var v26 bool = false
        for fi27 := 0; fi27 < 4; fi27++ {
            switch "\n\t" {
            case "hello":
                var v28 int8 = -128
                _ = v28
            }
            var v29 int64 = 42
            v30 := GenS3{F0: 631}
            var v31 uint16 = 255
            _ = fi27
            _ = v29
            _ = v30
            _ = v31
        }
        _ = fi25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
