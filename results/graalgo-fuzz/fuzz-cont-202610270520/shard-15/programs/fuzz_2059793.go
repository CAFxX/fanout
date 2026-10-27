package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 int16; F2 uint16 }
func genFn0(p0 uint8, p1 bool) uint8 {
    fmt.Println(p1, p1)
    var v7 uint8 = p0
    _ = p0
    _ = p1
    _ = v7
    return p0
}

func main() {
    v1 := make(map[int]string)
    v2 := []string{}
    var v3 int16 = 42
    var v4 int64 = 1
    var v5 int = -2147483648
    var v6 float32 = 3.14159
    for fi8 := 0; fi8 < 20; fi8++ {
        v9 := GenS1{F0: 1}
        var v10 uint = 778
        var v11 uint8 = 1
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    v12 := GenS1{F0: 2147483647}
    var v13 int32 = -2147483648
    if (!(!("hello" == ("hello" + "\n\t")))) {
        _ = len(v1)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v14 float64 = -2.25
    switch v5 {
    case 9223372036854775807:
        v14 = ((v14 - (float64(v6) / 3.14159)) - v14)
    }
    v15 := GenS1{F0: 2147483648}
    for fi16 := 0; fi16 < 13; fi16++ {
        var v17 uint16 = 896
        for fi18 := 0; fi18 < 6; fi18++ {
            v19 := v1[fi16]
            _ = fi18
            _ = v19
        }
        for fi20 := 0; fi20 < 17; fi20++ {
            _ = len(v2)
            v6 -= v6
            _ = fi20
        }
        _ = fi16
        _ = v17
    }
    if false {
        v2 = append(v2, "hello")
        if (100 != 1) {
            v21 := GenS2{F0: v4, F1: 297, F2: 2}
            v22 := GenS2{F0: v4, F1: 32767, F2: 570}
            v6 += v6
            _ = v21
            _ = v22
        } else {
            _ = len(v2)
            switch v5 {
            case 32768:
                switch v5 {
                case -1:
                    v23 := GenS2{F0: v4, F1: v3, F2: 2}
                    _ = v23
                default:
                    delete(v1, (v5 * (v5 ^ v5)))
                }
            case 0:
                var v24 uint = 19
                _ = v24
            case 9223372036854775807:
                fmt.Println(v14, v3)
            }
            switch "a" {
            case "\n\t":
                fmt.Println(v6, v5)
            default:
                var v25 string = "foo bar"
                _ = v25
            }
        }
    }
    for fi26 := 0; fi26 < 16; fi26++ {
        var v27 int8 = 1
        _ = fi26
        _ = v27
    }
    fmt.Println(v14, v4, v14)
    var v28 uint32 = 2147483647
    if true {
        if (("foo bar" + ("foo bar" + "hello")) <= "a") {
            switch v5 {
            case 9223372036854775807:
                _ = len(v1)
            case 2147483647:
                var v29 string = ("hello" + "\n\t")
                _ = v29
            case 127:
                var v30 int16 = v3
                _ = v30
            }
            v31 := v1[-755]
            var v32 uint64 = 18446744073709551615
            _ = v31
            _ = v32
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        fmt.Println(v4, v14)
    } else {
        var v33 uint = 4294967295
        var v34 int = 127
        _ = v33
        _ = v34
    }
    for fi35 := 0; fi35 < 20; fi35++ {
        switch fi35 {
        case 4294967295:
            var v36 uint32 = v28
            _ = v36
        }
        v4 = v4
        _ = fi35
    }
    var v37 bool = (18446744073709551615 != uint64(v28))
    fmt.Println(v4, v6, v28)
    switch ("" + "") {
    case "a":
        var v38 float32 = float32(((v14 / (((float64(3.14159) + float64(-2.25)) + float64(((0.0 * 3.14159) + float32(-2.25)))) + v14)) + 3.14159))
        _ = v38
    default:
        _ = len(v2)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v28
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
