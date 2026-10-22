package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 bool }
type GenS2 struct { F0 uint }

func main() {
    v1 := make(map[int]uint32)
    v2 := []uint32{0, 2147483648, 65535, 0}
    v3 := []uint32{37, 216, 2147483647, 0}
    var v4 string = ("foo bar" + ((("\n\t" + "a") + ((("foo bar" + "\n\t") + "foo bar") + (("" + ("hello" + ("a" + (("foo bar" + "") + "\n\t")))) + (("" + "") + "")))) + "foo bar"))
    if (true && true) {
        if false {
            v4 = v4
            var v5 int64 = 42
            _ = v5
        }
    }
    switch 42 {
    case 9223372036854775807:
        fmt.Println(v4)
    case 32767:
        v4 = "a"
    case -2147483648:
        for fi6 := 0; fi6 < 3; fi6++ {
            fmt.Println(fi6)
            _ = fi6
        }
    }
    if (255 <= 910) {
        v7 := v1[4294967295]
        if true {
            if (!true) {
                switch 484 {
                case -491:
                    var v8 string = v4
                    _ = v8
                case 184:
                    v7 = v7
                }
                var v9 int = 121
                _ = v9
            } else {
                switch -822 {
                case 32768:
                    for fi10 := 0; fi10 < 7; fi10++ {
                        v4 = ("hello" + v4)
                        _ = fi10
                    }
                default:
                    delete(v1, -9223372036854775808)
                }
                var v11 uint8 = 1
                _ = v11
            }
        }
        _ = v7
    } else {
        v12 := GenS1{F0: -1, F1: 1000, F2: (265 >= 1)}
        switch 32768 {
        case 127:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case 2147483647:
            var v13 int = -566
            _ = v13
        case 1:
            var v14 float64 = 3.14159
            _ = v14
        }
        _ = v12
    }
    for fi15 := 0; fi15 < 4; fi15++ {
        v4 = v4
        if true {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v16 int8 = -128
            var v17 int16 = 612
            _ = v16
            _ = v17
        } else {
            for fi18 := 0; fi18 < 12; fi18++ {
                v4 = v4
                _ = fi18
            }
        }
        _ = fi15
    }
    delete(v1, -9223372036854775808)
    v4 = "hello"
    switch 9223372036854775807 {
    case 4294967295:
        var v19 float32 = (-2.25 - -2.25)
        _ = v19
    }
    var v20 int32 = -1
    v21 := GenS1{F0: 0, F1: -32768, F2: true}
    switch 1 {
    case 128:
        if (757 != -2147483648) {
            for fi22 := 0; fi22 < 16; fi22++ {
                for fi23 := 0; fi23 < 16; fi23++ {
                    switch fi23 {
                    case 0:
                        v24 := GenS2{F0: 255}
                        _ = v24
                    }
                    v20 -= (v20 % ((v20) | 1))
                    _ = fi23
                }
                fmt.Println(v4, fi22, v20)
                _ = fi22
            }
            v4 = v4
            var v25 uint8 = 118
            _ = v25
        }
    default:
        var v26 string = v4
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
