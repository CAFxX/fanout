package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 int16 }
type GenS2 struct { F0 uint8; F1 int; F2 uint8 }
func genFn0(p0 bool, p1 bool, p2 uint32) int32 {
    for fi5 := 0; fi5 < 14; fi5++ {
        switch (fi5 << ((2) & 63)) {
        case 127:
            for fi6 := 0; fi6 < 1; fi6++ {
                if true {
                    v7 := []int{}
                    if p0 {
                        var v8 bool = p0
                        _ = v8
                    } else {
                        switch (fi6 >> ((65535) & 63)) {
                        case 863:
                            var v9 uint64 = 42
                            _ = v9
                        default:
                            var v10 int16 = 266
                            _ = v10
                        }
                        p0 = p1
                        p0 = p0
                        var v11 uint16 = 32767
                        _ = v11
                    }
                    p0 = p1
                    _ = v7
                }
                _ = fi6
            }
        case 166:
            v12 := GenS1{F0: 52, F1: 0, F2: 1}
            _ = v12
        case 128:
            for fi13 := 0; fi13 < 9; fi13++ {
                var v14 uint16 = 0
                var v15 uint32 = p2
                v16 := GenS1{F0: 255, F1: 0, F2: 1}
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
            }
        }
        if (!p1) {
            p2 = 194
            v17 := []int{}
            _ = v17
        }
        if (((1.5 + ((0.1 + ((0.0 * (1.5 - 1.5)) + 1.5)) - (0.0 / 1e10))) + 0.0) > -2.25) {
            v18 := []int{238, fi5, fi5}
            _ = v18
        }
        _ = fi5
    }
    switch (("\n\t" + "foo bar") + ("hello" + "\n\t")) {
    case "\n\t":
        var v19 float64 = -2.25
        _ = v19
    case "foo bar":
        v20 := GenS1{F0: 255, F1: 0, F2: 42}
        _ = v20
    case "a":
        v21 := GenS1{F0: 255, F1: -503, F2: -282}
        _ = v21
    default:
        fmt.Println(p1, p1)
    }
    v22 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    return -2147483648
}

func main() {
    v1 := []string{"", "", ""}
    var v2 int16 = -241
    var v3 uint = 856
    var v4 uint16 = 2
    for fi23 := 0; fi23 < 2; fi23++ {
        if len(v1) > 0 {
            v1[0] = "hello"
        }
        var v24 string = "foo bar"
        _ = fi23
        _ = v24
    }
    v25 := GenS2{F0: 2, F1: 32767, F2: 128}
    switch "" {
    case "foo bar":
        v4 = (v4 & (v4 - v4))
    case "hello":
        v3 += v3
    case "\n\t":
        fmt.Println(v2, v2)
    default:
        if false {
            if (!true) {
                fmt.Println(v3)
                var v26 uint16 = v4
                v2 = v2
                _ = v26
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        } else {
            var v27 int16 = v2
            _ = v27
        }
    }
    for fi28 := 0; fi28 < 4; fi28++ {
        var v29 float64 = 0.1
        v29 += v29
        _ = fi28
        _ = v29
    }
    v2 += v2
    v3 = 640
    var v30 float64 = float64(((0.1 * 0.0) * ((0.0 - ((1.5 + float32(float64(1e10))) + (1.5 / 1.0))) / 1.0)))
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    var v31 int8 = 47
    if ((!(false || false)) && (false || true)) {
        _ = len(v1)
    } else {
        var v32 float64 = (float64(3.14159) - (v30 / (v30 - (float64(3.14159) - v30))))
        _ = v32
    }
    var v33 int8 = (-(v31))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v30
    _ = v31
    _ = v33
    _ = v4
}
