package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 string }

func main() {
    v1 := []uint32{42, 2, 2}
    var v2 int32 = -864
    var v3 int = 128
    v4 := make(map[int]bool)
    var v5 int = (v3 >> ((2) & 63))
    fmt.Println(v5)
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 string = ""
        _ = fi6
        _ = v7
    }
    var v8 uint8 = 15
    for fi9 := 0; fi9 < 15; fi9++ {
        v8 = 1
        _ = fi9
    }
    if false {
        var v10 float64 = 0.1
        fmt.Println(v3, v3)
        if (!true) {
            _ = len(v1)
            switch ("hello" + ("a" + "")) {
            case "foo bar":
                fmt.Println(v5, v5)
            case "":
                if true {
                    v10 -= (v10 + (1.5 - 0.1))
                    v1 = append(v1, 4294967295)
                    var v11 int64 = -1
                    _ = v11
                }
            case "a":
                v5 += (v3 & v5)
            default:
                var v12 uint = 65535
                _ = v12
            }
            for fi13 := 0; fi13 < 2; fi13++ {
                for fi14 := 0; fi14 < 13; fi14++ {
                    v15 := GenS1{F0: 111}
                    _ = fi14
                    _ = v15
                }
                v16 := GenS1{F0: -1}
                v5 = fi13
                _ = fi13
                _ = v16
            }
        } else {
            v4[(v5 ^ (^(v5)))] = ((("\n\t" + "a") + ("hello" + (((("foo bar" + ("\n\t" + ("foo bar" + (("\n\t" + "hello") + (("\n\t" + "\n\t") + ("a" + "foo bar")))))) + "\n\t") + "a") + ""))) < "hello")
        }
        _ = v10
    } else {
        v17 := GenS2{F0: 4294967295, F1: ("foo bar" + "")}
        _ = v17
    }
    var v18 bool = true
    var v19 string = ("hello" + (("foo bar" + "") + "foo bar"))
    var v20 float64 = 1e10
    if (-1 == 127) {
        v19 = v19
        for fi21 := 0; fi21 < 7; fi21++ {
            v22 := GenS2{F0: 2, F1: (v19 + v19)}
            fmt.Println(v2, v19, fi21)
            _ = fi21
            _ = v22
        }
    } else {
        var v23 float64 = 1.5
        v8 = 255
        v2 += v2
        _ = v23
    }
    if (v18 && ((v18 && (175 != 246)) && ((250 >= 1) && false))) {
        for fi24 := 0; fi24 < 13; fi24++ {
            var v25 uint8 = v8
            var v26 string = (v19 + ("" + (v19 + "foo bar")))
            if (-423 != v2) {
                if (float32(((v20 / v20) + (v20 * (v20 / (-2.25 + (3.14159 - v20)))))) > 0.0) {
                    for fi27 := 0; fi27 < 15; fi27++ {
                        v28 := GenS1{F0: 42}
                        _ = fi27
                        _ = v28
                    }
                    fmt.Println(v5)
                }
            } else {
                v26 = v19
            }
            _ = fi24
            _ = v25
            _ = v26
        }
        v20 += v20
        v19 = v19
    } else {
        var v29 int32 = (v2 | (-(v2)))
        v29 = (v2 & v2)
        switch v19 {
        case "hello":
            v1 = append(v1, 2)
        case "\n\t":
            var v30 int64 = 1
            _ = v30
        default:
            v1 = append(v1, 65535)
        }
        _ = v29
    }
    _ = len(v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
