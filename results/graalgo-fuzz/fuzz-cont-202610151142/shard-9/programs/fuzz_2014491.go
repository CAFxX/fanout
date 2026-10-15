package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 int }
type GenS2 struct { F0 bool }

func main() {
    v1 := []string{"hello"}
    v2 := make(map[string]bool)
    var v3 uint32 = 784
    var v4 uint64 = 2
    v5 := make(map[string]uint32)
    var v6 int32 = 1
    var v7 float32 = float32((1.5 + 0.1))
    if false {
        if true {
            switch "hello" {
            case "a":
                var v8 int = -350
                _ = v8
            case "foo bar":
                v9 := GenS1{F0: true, F1: 0, F2: -460}
                _ = v9
            }
        }
        var v10 uint16 = 1
        var v11 int64 = 0
        _ = v10
        _ = v11
    }
    var v12 float64 = ((1e10 / (float64(1e10) - -2.25)) * 3.14159)
    switch 686 {
    case 1:
        v5["hello"] = (v3 ^ v3)
    }
    if true {
        var v13 int8 = int8(v6)
        var v14 uint64 = 68
        _ = v13
        _ = v14
    } else {
        var v15 int = 32767
        v16 := GenS1{F0: false, F1: 549, F2: v15}
        fmt.Println(v4, v15, v7)
        _ = v15
        _ = v16
    }
    switch "foo bar" {
    case "a":
        v3 -= (v3 - v3)
    case "\n\t":
        switch "" {
        case "foo bar":
            var v17 float64 = v12
            _ = v17
        default:
            v3 -= 1
        }
    default:
        if (-1 != 9223372036854775807) {
            v7 = (v7 + v7)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    }
    var v18 int8 = -11
    switch int(v18) {
    case 128:
        var v19 string = ("a" + ("" + ("a" + "hello")))
        _ = v19
    default:
        if true {
            for fi20 := 0; fi20 < 11; fi20++ {
                fmt.Println(v6)
                fmt.Println(v18, v3, v18)
                _ = fi20
            }
        } else {
            v21 := v5[("" + ((("hello" + "a") + "foo bar") + "\n\t"))]
            _ = v21
        }
    }
    switch 498 {
    case 2147483647:
        var v22 uint64 = 4294967295
        _ = v22
    case -1:
        v23 := GenS2{F0: (false || false)}
        _ = v23
    case -511:
        var v24 int8 = 100
        _ = v24
    default:
        v4 -= (v4 - v4)
    }
    if (-2.25 >= 0.1) {
        v25 := GenS2{F0: (0.1 < v7)}
        v26 := GenS1{F0: (4294967295 >= 42), F1: 65535, F2: 86}
        v6 = v6
        _ = v25
        _ = v26
    }
    switch 4294967295 {
    case 562:
        var v27 uint32 = 696
        _ = v27
    case 128:
        var v28 uint8 = 158
        _ = v28
    case 0:
        var v29 float32 = (3.14159 - v7)
        _ = v29
    default:
        switch "hello" {
        case "\n\t":
            v3 = (v3 / ((v3) | 1))
        case "a":
            if (!(false || true)) {
                switch "" {
                case "a":
                    if (!true) {
                        var v30 int16 = 0
                        var v31 int8 = v18
                        _ = v30
                        _ = v31
                    }
                case "foo bar":
                    var v32 float32 = float32(0.0)
                    _ = v32
                case "hello":
                    switch 128 {
                    case 128:
                        var v33 float64 = (v12 + v12)
                        _ = v33
                    case 127:
                        v18 += (v18 - v18)
                    }
                }
                var v34 int64 = -345
                v12 += (((v12 / (v12 * (0.1 * v12))) * 1.5) - v12)
                _ = v34
            }
        case "":
            if len(v1) > 0 {
                v1[0] = ""
            }
        default:
            var v35 int16 = 32767
            _ = v35
        }
    }
    var v36 bool = (-344 >= 420)
    for fi37 := 0; fi37 < 9; fi37++ {
        switch 0 {
        case 32767:
            fmt.Println(v12, v36)
        case 42:
            var v38 int = 32768
            _ = v38
        default:
            var v39 int16 = 32767
            _ = v39
        }
        _ = fi37
    }
    v4 -= 238
    if v36 {
        v6 += (v6 + v6)
        fmt.Println(v18, v12, v36)
        v4 -= (^(v4))
    }
    v7 = v7
    var v40 float32 = 3.14159
    v3 = (v3 ^ (^(v3)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v18
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v7
}
