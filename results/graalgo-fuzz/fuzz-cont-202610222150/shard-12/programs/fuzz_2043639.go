package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int; F1 bool }

func main() {
    var v1 int32 = -1
    var v2 bool = false
    v3 := []int{127, -944, -1, 2147483647}
    var v4 float32 = 0.0
    v2 = (true || (!(!((-1 < 1) && (1 > -945)))))
    fmt.Println(v2, v2)
    _ = len(v3)
    if v2 {
        v3 = append(v3, 4294967295)
        v2 = (!((-1 < 50) || v2))
        if (("foo bar" <= "\n\t") || v2) {
            v5 := GenS1{F0: (v1 / (((v1 >> ((476) & 63))) | 1)), F1: -378}
            _ = v5
        } else {
            v3 = append(v3, -369)
            v6 := GenS2{F0: 164}
            for fi7 := 0; fi7 < 12; fi7++ {
                var v8 int32 = (v1 | v1)
                v8 = v1
                v4 += v4
                _ = fi7
                _ = v8
            }
            _ = v6
        }
    } else {
        for fi9 := 0; fi9 < 15; fi9++ {
            v2 = (86 == 127)
            v10 := GenS3{F0: (fi9 >= -438)}
            var v11 uint = 1
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    v2 = true
    switch 9223372036854775807 {
    case -856:
        var v12 int32 = v1
        _ = v12
    case 127:
        for fi13 := 0; fi13 < 14; fi13++ {
            for fi14 := 0; fi14 < 10; fi14++ {
                if v2 {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v15 string = (("foo bar" + "foo bar") + "foo bar")
                    _ = v15
                }
                v4 -= (3.14159 + v4)
                if (-32768 > 989) {
                    v3 = append(v3, (fi14 | (fi13 % ((fi14) | 1))))
                    v4 = v4
                }
                switch -9223372036854775808 {
                case 128:
                    v16 := GenS2{F0: 127}
                    _ = v16
                case -387:
                    v3 = append(v3, 128)
                default:
                    var v17 uint8 = 128
                    _ = v17
                }
                _ = fi14
            }
            v1 = v1
            var v18 int16 = 0
            _ = fi13
            _ = v18
        }
    default:
        if v2 {
            fmt.Println(v4)
        }
    }
    v1 += (v1 % ((v1) | 1))
    v2 = ("foo bar" != "")
    var v19 string = "foo bar"
    for fi20 := 0; fi20 < 1; fi20++ {
        fmt.Println(v1, v2)
        if v2 {
            switch (fi20 | 1) {
            case 1:
                var v21 int = fi20
                _ = v21
            case 4294967295:
                switch (-(fi20)) {
                case 612:
                    v19 = ""
                case -422:
                    v22 := GenS4{F0: (fi20 + fi20), F1: (42 == 1)}
                    _ = v22
                default:
                    var v23 bool = v2
                    _ = v23
                }
            default:
                fmt.Println(v2)
            }
            var v24 bool = v2
            _ = v24
        }
        var v25 bool = (-2.25 < (0.1 * (3.14159 / 3.14159)))
        _ = fi20
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
