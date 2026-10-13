package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 uint }
type GenS2 struct { F0 int32; F1 uint64; F2 uint64 }
type GenS3 struct { F0 int16; F1 int32; F2 uint32 }
func genFn0(p0 bool, p1 uint16) string {
    if (65535 != 2147483647) {
        v7 := []int{-523}
        v8 := GenS1{F0: 1, F1: -1, F2: 40}
        _ = v7
        _ = v8
    }
    var v9 string = "\n\t"
    var v10 uint8 = 255
    v11 := GenS2{F0: -176, F1: 153, F2: 382}
    v12 := []int{1}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    _ = v9
    return "a"
}

func main() {
    var v1 uint64 = 0
    var v2 int16 = 42
    var v3 int64 = 674
    var v4 int64 = -1
    v5 := []string{"", ("a" + (("" + "hello") + "\n\t")), ""}
    var v6 int16 = v2
    if (!((-65 >= 3) || false)) {
        for fi13 := 0; fi13 < 14; fi13++ {
            for fi14 := 0; fi14 < 1; fi14++ {
                if (!(3.14159 >= (1e10 / 3.14159))) {
                    if (232 != 302) {
                        v3 = -9223372036854775808
                    }
                    v3 = v3
                    switch (fi14 ^ (-(fi13))) {
                    case -7:
                        v15 := GenS1{F0: 2147483647, F1: 1, F2: 18446744073709551615}
                        _ = v15
                    case 387:
                        var v16 uint8 = 255
                        _ = v16
                    }
                    _ = len(v5)
                    var v17 int32 = 2147483647
                    _ = v17
                } else {
                    switch fi14 {
                    case 309:
                        v3 = v3
                    }
                    switch "" {
                    case "foo bar":
                        var v18 string = ("foo bar" + ((("" + "a") + "\n\t") + "\n\t"))
                        _ = v18
                    case "a":
                        var v19 int = 42
                        _ = v19
                    }
                }
                switch (fi13 - (fi14 & 1)) {
                case 32767:
                    v6 = (v6 / (((-(v2))) | 1))
                case 42:
                    v4 = (-(v3))
                case -952:
                    v5 = append(v5, ("a" + "foo bar"))
                default:
                    fmt.Println(v4, v3, v4)
                }
                _ = fi14
            }
            for fi20 := 0; fi20 < 5; fi20++ {
                switch "" {
                case "hello":
                    fmt.Println(v2, v6)
                }
                v5 = append(v5, ("" + ("foo bar" + "foo bar")))
                switch fi13 {
                case 756:
                    fmt.Println(v6)
                case -445:
                    v21 := GenS1{F0: 2147483647, F1: 42, F2: 382}
                    _ = v21
                case 32768:
                    v3 = 42
                default:
                    var v22 uint64 = v1
                    _ = v22
                }
                _ = fi20
            }
            _ = fi13
        }
        fmt.Println(v6)
        if len(v5) > 0 {
            v5[0] = ""
        }
    } else {
        v23 := GenS1{F0: uint32(v1), F1: 128, F2: 255}
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v23
    }
    switch 104 {
    case 42:
        var v24 int64 = 0
        _ = v24
    case -936:
        fmt.Println(v2)
    }
    var v25 int32 = -1
    if len(v5) > 0 {
        v5[0] = ("hello" + (((("" + "\n\t") + "\n\t") + ("a" + "a")) + ("a" + "\n\t")))
    }
    switch int(v4) {
    case 1:
        var v26 uint8 = 2
        _ = v26
    case -355:
        fmt.Println(v6, v25, v25)
    case 127:
        for fi27 := 0; fi27 < 4; fi27++ {
            _ = len(v5)
            var v28 uint16 = 0
            _ = fi27
            _ = v28
        }
    }
    fmt.Println(v4, v6, v3)
    v29 := GenS2{F0: 1, F1: 0, F2: v1}
    v30 := GenS3{F0: (v6 << ((1) & 63)), F1: v25, F2: 178}
    var v31 int8 = 0
    v31 = v31
    v32 := GenS1{F0: 2147483647, F1: 0, F2: 65535}
    for fi33 := 0; fi33 < 20; fi33++ {
        v3 -= (v3 & v4)
        var v34 float32 = float32(1e10)
        fmt.Println(v34, v3, v1)
        _ = fi33
        _ = v34
    }
    fmt.Println(v3, v3)
    for fi35 := 0; fi35 < 9; fi35++ {
        _ = len(v5)
        switch (fi35 * 127) {
        case 0:
            v36 := GenS2{F0: v25, F1: (v1 - v1), F2: v1}
            _ = v36
        }
        if (!false) {
            for fi37 := 0; fi37 < 13; fi37++ {
                v31 -= v31
                var v38 int16 = -32768
                var v39 string = "a"
                _ = fi37
                _ = v38
                _ = v39
            }
        } else {
            v31 = int8(v2)
            var v40 string = "a"
            _ = v40
        }
        _ = fi35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
