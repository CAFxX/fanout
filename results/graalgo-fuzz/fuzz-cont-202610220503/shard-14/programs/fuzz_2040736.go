package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 string; F1 string }
func genFn0(p0 int32, p1 uint8, p2 int16) int {
    if (!false) {
        p1 -= p1
        p2 = p2
    }
    switch (("hello" + ("" + "a")) + "\n\t") {
    case "foo bar":
        var v4 bool = (!false)
        _ = v4
    case "":
        p1 += p1
    case "\n\t":
        switch "foo bar" {
        case "":
            v5 := []int{-2147483648, 26}
            _ = v5
        default:
            v6 := GenS1{F0: -128, F1: 42}
            _ = v6
        }
    }
    v7 := []int{-931, 21}
    if (!(float32(float64(0.1)) != float32(float64(1.5)))) {
        for fi8 := 0; fi8 < 12; fi8++ {
            for fi9 := 0; fi9 < 19; fi9++ {
                _ = len(v7)
                _ = fi9
            }
            _ = fi8
        }
        switch ("foo bar" + (("hello" + "\n\t") + "")) {
        case "\n\t":
            if (2 == 1) {
                fmt.Println(p0, p2)
                p1 += (p1 & p1)
            } else {
                fmt.Println(p2)
            }
        case "hello":
            for fi10 := 0; fi10 < 3; fi10++ {
                var v11 uint8 = p1
                var v12 float64 = 0.1
                _ = fi10
                _ = v11
                _ = v12
            }
        case "":
            p1 = p1
        }
        var v13 int32 = 2147483647
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 32768
}
func genFn1(p0 uint16) uint8 {
    switch 0 {
    case 42:
        switch "hello" {
        case "\n\t":
            v14 := GenS1{F0: 62, F1: 58}
            _ = v14
        case "a":
            var v15 uint = 4294967295
            _ = v15
        default:
            p0 += p0
        }
    }
    v16 := []int{}
    var v17 uint32 = 1
    v17 += 0
    _ = p0
    _ = v16
    _ = v17
    return 128
}

func main() {
    var v1 int16 = -300
    var v2 uint = 255
    var v3 uint = (v2 << ((18446744073709551615) & 63))
    if (-2147483648 != -2147483648) {
        if (true && (((!false) || ((^(v2)) > (v2 / ((v3) | 1)))) || ((2 > 23) || (!(255 <= 654))))) {
            for fi18 := 0; fi18 < 9; fi18++ {
                v19 := []int{32767, 127}
                v1 += v1
                _ = fi18
                _ = v19
            }
            fmt.Println(v2)
        }
        var v20 bool = false
        var v21 float64 = 0.1
        _ = v20
        _ = v21
    } else {
        v1 = 587
        v22 := GenS1{F0: -128, F1: 0}
        _ = v22
    }
    v1 = v1
    v2 = (^(v2))
    v2 = (v2 | (v3 * (^(v3))))
    var v23 bool = (128 == 0)
    v24 := GenS1{F0: -1, F1: 2147483647}
    var v25 int32 = 122
    var v26 int16 = (v1 + v1)
    switch 9223372036854775807 {
    case 128:
        var v27 int16 = (v26 << (((v3 << ((v3) & 63))) & 63))
        _ = v27
    }
    v1 += 1000
    var v28 int8 = 0
    v29 := []int{}
    for fi30 := 0; fi30 < 20; fi30++ {
        for fi31 := 0; fi31 < 13; fi31++ {
            v23 = (!(1 != 0))
            _ = fi31
        }
        var v32 string = (("" + ("hello" + "\n\t")) + ("" + ("\n\t" + "a")))
        for fi33 := 0; fi33 < 10; fi33++ {
            var v34 int32 = (v25 << (((v2 % ((v2) | 1))) & 63))
            _ = fi33
            _ = v34
        }
        _ = fi30
        _ = v32
    }
    var v35 float32 = ((0.0 * (0.1 * (3.14159 / (1.5 / float32(1.5))))) + 0.0)
    switch -892 {
    case 9223372036854775807:
        for fi36 := 0; fi36 < 6; fi36++ {
            for fi37 := 0; fi37 < 3; fi37++ {
                v28 = -52
                if len(v29) > 0 {
                    v29[0] = 32768
                }
                v29 = append(v29, fi36)
                _ = fi37
            }
            _ = fi36
        }
    case 0:
        v38 := GenS2{F0: (("" + "a") + ""), F1: "foo bar"}
        _ = v38
    case 128:
        v28 = -1
    default:
        var v39 uint16 = 849
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v35
}
