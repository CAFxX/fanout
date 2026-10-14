package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool; F1 uint }
func genFn0(p0 int, p1 uint, p2 uint32) int64 {
    if ((!true) || true) {
        p1 = p1
    } else {
        for fi5 := 0; fi5 < 18; fi5++ {
            p1 = 255
            _ = fi5
        }
        p2 += 42
    }
    if (true || (1 < -251)) {
        v6 := GenS1{F0: -1}
        _ = v6
    }
    var v7 uint = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 399
}

func main() {
    var v1 int8 = -83
    v2 := make(map[uint32]string)
    var v3 int64 = -9223372036854775808
    v4 := []string{}
    if ((0.0 + (0.1 * ((3.14159 + 1.5) + 1e10))) >= float64((1.5 / ((((1e10 + -2.25) * 0.0) + 0.1) + (0.1 - -2.25))))) {
        _ = len(v4)
    } else {
        for fi8 := 0; fi8 < 18; fi8++ {
            v2[2147483647] = ("hello" + "\n\t")
            for fi9 := 0; fi9 < 17; fi9++ {
                switch "a" {
                case "\n\t":
                    _ = len(v4)
                default:
                    var v10 int16 = -238
                    _ = v10
                }
                if true {
                    var v11 uint16 = 65535
                    _ = v11
                }
                if len(v4) > 0 {
                    v4[0] = ("\n\t" + "foo bar")
                }
                _ = fi9
            }
            _ = fi8
        }
        v12 := GenS1{F0: (-(v3))}
        v13 := GenS1{F0: v3}
        _ = v12
        _ = v13
    }
    if (2 < 1) {
        for fi14 := 0; fi14 < 13; fi14++ {
            var v15 string = "foo bar"
            var v16 float32 = 0.1
            _ = fi14
            _ = v15
            _ = v16
        }
        delete(v2, 1)
        v17 := GenS2{F0: 42}
        _ = v17
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi18 := 0; fi18 < 4; fi18++ {
        for fi19 := 0; fi19 < 1; fi19++ {
            fmt.Println(fi18, fi19)
            _ = fi19
        }
        v3 -= v3
        v1 = -128
        _ = fi18
    }
    var v20 string = (("hello" + "") + "\n\t")
    if true {
        fmt.Println(v3, v3, v3)
    } else {
        if ((v3 % ((v3) | 1)) >= 9223372036854775807) {
            v3 = (-(v3))
        }
        if len(v4) > 0 {
            v4[0] = v20
        }
        var v21 int16 = -1
        _ = v21
    }
    switch "a" {
    case "hello":
        if true {
            var v22 int32 = 42
            _ = v22
        } else {
            var v23 int32 = 0
            var v24 uint32 = 329
            _ = v23
            _ = v24
        }
    case "":
        var v25 uint32 = 4294967295
        _ = v25
    case "a":
        v2[838] = "foo bar"
    default:
        v3 += v3
    }
    v26 := GenS3{F0: (0 == 2147483648), F1: 272}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v26
    _ = v3
    _ = v4
}
