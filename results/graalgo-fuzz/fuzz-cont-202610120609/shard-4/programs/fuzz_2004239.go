package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 uint; F1 int8; F2 int64 }

func main() {
    var v1 int = -822
    v2 := []uint32{}
    var v3 bool = false
    var v4 int64 = 0
    v2 = append(v2, 4294967295)
    if (4294967295 < 26) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v2 = append(v2, 2)
    fmt.Println(v4, v1)
    if (float64(-2.25) < 3.14159) {
        v3 = (844 != 262)
        v3 = (-128 <= -115)
        var v5 float32 = 1.5
        _ = v5
    }
    if len(v2) > 0 {
        v2[0] = 50
    }
    v3 = (195 == 2)
    switch "foo bar" {
    case "foo bar":
        var v6 int16 = -1
        _ = v6
    case "":
        switch (v1 | (v1 + v1)) {
        case -1:
            switch ("" + "") {
            case "hello":
                var v7 uint8 = 1
                _ = v7
            case "foo bar":
                var v8 uint8 = 255
                _ = v8
            case "\n\t":
                var v9 int = (v1 << ((1) & 63))
                _ = v9
            default:
                fmt.Println(v4)
            }
        case -573:
            v4 -= (v4 % ((-9223372036854775808) | 1))
        default:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    case "\n\t":
        for fi10 := 0; fi10 < 10; fi10++ {
            v11 := GenS1{F0: 1, F1: 1}
            _ = fi10
            _ = v11
        }
    }
    _ = len(v2)
    v1 += -39
    var v12 uint = 65535
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi13 := 0; fi13 < 17; fi13++ {
        v14 := GenS2{F0: 4294967295, F1: -128, F2: v4}
        var v15 int32 = 1
        _ = fi13
        _ = v14
        _ = v15
    }
    if (int(v4) == (-(v1))) {
        var v16 bool = (30 != 255)
        switch 128 {
        case 32768:
            if len(v2) > 0 {
                v2[0] = 774
            }
        default:
            fmt.Println(v4)
        }
        for fi17 := 0; fi17 < 4; fi17++ {
            for fi18 := 0; fi18 < 7; fi18++ {
                v1 -= v1
                fmt.Println(fi18, v3, v1)
                _ = fi18
            }
            var v19 uint = v12
            for fi20 := 0; fi20 < 3; fi20++ {
                _ = len(v2)
                _ = fi20
            }
            _ = fi17
            _ = v19
        }
        _ = v16
    }
    for fi21 := 0; fi21 < 2; fi21++ {
        var v22 uint16 = 0
        if len(v2) > 0 {
            v2[0] = 2147483648
        }
        v4 = -9223372036854775808
        _ = fi21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
