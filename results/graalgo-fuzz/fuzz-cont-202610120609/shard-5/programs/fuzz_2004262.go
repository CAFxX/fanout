package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 bool }
type GenS2 struct { F0 string; F1 uint }
func genFn0(p0 uint32, p1 int8, p2 string) int16 {
    v4 := GenS1{F0: true, F1: 65535, F2: false}
    if false {
        var v5 bool = (9223372036854775807 != 1)
        _ = v5
    }
    v6 := GenS1{F0: true, F1: p0, F2: (!false)}
    v7 := []int{-1}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v6
    _ = v7
    return 1000
}

func main() {
    var v1 bool = false
    var v2 float32 = float32(((((1e10 * -2.25) / float64(1e10)) - float64(1e10)) - float64(((1.5 + (0.1 - 1e10)) + (float32(3.14159) - (1e10 * -2.25))))))
    var v3 string = (("foo bar" + "foo bar") + "foo bar")
    var v8 uint8 = 2
    fmt.Println(v1)
    fmt.Println(v1, v2)
    if v1 {
        var v9 float64 = (float64(3.14159) * 0.0)
        var v10 float32 = v2
        v8 = (v8 * (v8 / ((v8) | 1)))
        _ = v10
        _ = v9
    }
    v8 += (v8 % ((2) | 1))
    v11 := []int{972, -820, -1, 257}
    var v12 uint8 = (v8 % (((v8 ^ v8)) | 1))
    v12 = v12
    for fi13 := 0; fi13 < 8; fi13++ {
        v14 := GenS2{F0: v3, F1: 2}
        _ = fi13
        _ = v14
    }
    if (-2.25 != v2) {
        for fi15 := 0; fi15 < 3; fi15++ {
            if (!(v1 || v1)) {
                v12 = 68
                v16 := GenS1{F0: true, F1: 2147483647, F2: (42 <= 4294967295)}
                _ = v16
            }
            for fi17 := 0; fi17 < 3; fi17++ {
                switch 128 {
                case -603:
                    if len(v11) > 0 {
                        v11[0] = (^(fi17))
                    }
                }
                v18 := GenS1{F0: v1, F1: 42, F2: (-2147483648 != -1)}
                _ = fi17
                _ = v18
            }
            var v19 float64 = ((1.5 + 0.0) / 1.0)
            _ = fi15
            _ = v19
        }
    }
    for fi20 := 0; fi20 < 12; fi20++ {
        v3 = v3
        v3 = (v3 + ("hello" + v3))
        _ = fi20
    }
    if (v8 == (v8 + v8)) {
        for fi21 := 0; fi21 < 8; fi21++ {
            fmt.Println(v2)
            var v22 int8 = 42
            switch v3 {
            case "foo bar":
                for fi23 := 0; fi23 < 11; fi23++ {
                    if (-981 > 86) {
                        switch "hello" {
                        case "":
                            var v24 int32 = 0
                            _ = v24
                        case "foo bar":
                            var v25 int16 = -1
                            _ = v25
                        case "hello":
                            var v26 string = (v3 + v3)
                            _ = v26
                        }
                        fmt.Println(fi23, v2, v1)
                        var v27 int32 = 1
                        _ = v27
                    }
                    v11 = append(v11, fi23)
                    _ = fi23
                }
            case "\n\t":
                var v28 int = fi21
                _ = v28
            default:
                var v29 int16 = 708
                _ = v29
            }
            _ = fi21
            _ = v22
        }
    }
    fmt.Println(v8, v2, v3)
    v12 = v8
    for fi30 := 0; fi30 < 18; fi30++ {
        var v31 uint = 255
        if len(v11) > 0 {
            v11[0] = (fi30 << ((65535) & 63))
        }
        if (0 < -32768) {
            v32 := GenS1{F0: (!v1), F1: 262, F2: ((^(fi30)) == (fi30 << ((1) & 63)))}
            _ = v32
        } else {
            fmt.Println(v8)
            switch v3 {
            case "foo bar":
                for fi33 := 0; fi33 < 5; fi33++ {
                    v11 = append(v11, fi30)
                    v34 := GenS1{F0: (1 == 2), F1: 4294967295, F2: v1}
                    _ = fi33
                    _ = v34
                }
            case "":
                if true {
                    v2 = v2
                }
            case "\n\t":
                var v35 int8 = 100
                _ = v35
            }
            v2 = (((3.14159 * -2.25) - v2) - float32(((1e10 * 1e10) + 3.14159)))
        }
        _ = fi30
        _ = v31
    }
    _ = len(v11)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v8
}
