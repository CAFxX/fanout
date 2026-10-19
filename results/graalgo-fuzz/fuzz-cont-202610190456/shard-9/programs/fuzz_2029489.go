package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }

func main() {
    v1 := make(map[int]string)
    var v2 int8 = 73
    var v3 int32 = 0
    var v4 int16 = 153
    v5 := []uint32{65535}
    var v6 uint64 = 4294967295
    v7 := GenS1{F0: 9223372036854775807, F1: v4}
    v2 -= (^(v2))
    fmt.Println(v4, v6, v2)
    if (!((v6 * v6) > (v6 >> ((182) & 63)))) {
        var v8 float64 = float64(float32(0.0))
        v8 -= v8
        _ = v8
    } else {
        v4 += 32767
    }
    for fi9 := 0; fi9 < 16; fi9++ {
        var v10 float32 = float32(((3.14159 - -2.25) / 1e10))
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 7; fi11++ {
        v6 = v6
        fmt.Println(v2)
        _ = fi11
    }
    if true {
        v4 += (v4 | -745)
    } else {
        for fi12 := 0; fi12 < 12; fi12++ {
            v3 += (v3 ^ (-(v3)))
            _ = fi12
        }
        var v13 uint = 42
        _ = v13
    }
    v4 += (v4 << ((18446744073709551615) & 63))
    var v14 float64 = 1e10
    for fi15 := 0; fi15 < 5; fi15++ {
        var v16 int16 = (v4 * v4)
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        switch ("hello" + (("a" + "foo bar") + "a")) {
        case "":
            var v18 int32 = (v3 | v3)
            _ = v18
        case "foo bar":
            fmt.Println(v3, v2)
        case "\n\t":
            if (177 > 42) {
                v2 = (v2 % ((v2) | 1))
                var v19 uint8 = 255
                _ = v19
            } else {
                v3 += (v3 ^ 42)
                v3 = (-(v3))
                if len(v5) > 0 {
                    v5[0] = 1
                }
            }
        }
        if ((!(1 == 42)) && false) {
            var v20 string = ("\n\t" + ("hello" + ("a" + ("hello" + ("foo bar" + "")))))
            var v21 int = (fi17 - (fi17 & -929))
            _ = v20
            _ = v21
        } else {
            var v22 uint32 = 701
            v1[fi17] = ("" + "foo bar")
            _ = v22
        }
        var v23 string = "a"
        _ = fi17
        _ = v23
    }
    v5 = append(v5, 143)
    var v24 int32 = v3
    if (!(2 <= 44)) {
        switch 9223372036854775807 {
        case 4294967295:
            v25 := GenS1{F0: 42, F1: (-(v4))}
            _ = v25
        case 2147483647:
            var v26 int32 = 67
            _ = v26
        case 22:
            if (v2 >= 127) {
                v27 := GenS1{F0: int(v4), F1: -1}
                _ = v27
            }
        }
        v3 += (v24 >> ((2) & 63))
    } else {
        var v28 int16 = -201
        v29 := GenS1{F0: -128, F1: (v28 & v4)}
        fmt.Println(v6, v14)
        _ = v28
        _ = v29
    }
    v14 -= (0.0 * v14)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
