package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 bool }
type GenS2 struct { F0 bool; F1 int64; F2 uint32 }
type GenS3 struct { F0 int }

func main() {
    var v1 bool = (!false)
    var v2 uint16 = 1
    v3 := make(map[string]uint32)
    _ = len(v3)
    for fi4 := 0; fi4 < 5; fi4++ {
        v5 := GenS1{F0: ((-2.25 + float32(1.5)) < 0.1), F1: -1, F2: true}
        _ = fi4
        _ = v5
    }
    var v6 uint8 = 0
    if ((true || (65535 != 255)) && v1) {
        v7 := []int{-148, -2147483648, -1, 0}
        fmt.Println(v2, v2)
        v1 = false
        _ = v7
    } else {
        var v8 int8 = -128
        switch 127 {
        case 2147483647:
            if v1 {
                for fi9 := 0; fi9 < 13; fi9++ {
                    fmt.Println(v6, v8)
                    fmt.Println(v6, fi9, fi9)
                    _ = fi9
                }
            }
        case 376:
            v2 += v2
        default:
            _ = len(v3)
        }
        _ = v8
    }
    _ = len(v3)
    if (!v1) {
        var v10 int16 = -32768
        for fi11 := 0; fi11 < 6; fi11++ {
            v1 = v1
            _ = fi11
        }
        fmt.Println(v6, v6, v2)
        _ = v10
    }
    v12 := GenS1{F0: v1, F1: -128, F2: v1}
    v1 = (v6 < 2)
    fmt.Println(v1, v6)
    _ = len(v3)
    var v13 uint8 = v6
    var v14 bool = (((v13 >> ((725) & 63)) < 81) && (438 != 65535))
    switch "hello" {
    case "a":
        var v15 int32 = 2147483647
        _ = v15
    case "":
        fmt.Println(v14, v1, v14)
    default:
        for fi16 := 0; fi16 < 12; fi16++ {
            v17 := GenS2{F0: false, F1: -9223372036854775808, F2: 0}
            var v18 int = (fi16 / ((2147483647) | 1))
            _ = fi16
            _ = v17
            _ = v18
        }
    }
    if v1 {
        v1 = v1
        v19 := GenS3{F0: 127}
        switch -2147483648 {
        case 1:
            for fi20 := 0; fi20 < 10; fi20++ {
                var v21 int = (^(fi20))
                _ = fi20
                _ = v21
            }
        case -149:
            if ("a" == "foo bar") {
                var v22 int8 = -128
                v23 := []int{32767, -400, 0}
                v3["\n\t"] = 65535
                _ = v22
                _ = v23
            } else {
                if (1 <= 147) {
                    switch ("\n\t" + "hello") {
                    case "":
                        v24 := v3["\n\t"]
                        _ = v24
                    case "hello":
                        var v25 int = 1
                        _ = v25
                    }
                    v26 := []int{220, -1, 2147483647, 1}
                    v27 := GenS3{F0: 128}
                    _ = v26
                    _ = v27
                }
                var v28 uint16 = v2
                _ = v28
            }
        case 9223372036854775807:
            var v29 int64 = 9223372036854775807
            _ = v29
        }
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v6
}
