package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 int32; F1 bool }
type GenS4 struct { F0 bool; F1 int }
type GenS5 struct { F0 int; F1 int16 }
type GenS6 struct { F0 uint }

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[string]uint32)
    var v3 float32 = (v1 - 3.14159)
    var v4 string = ("hello" + "")
    var v5 float64 = (1e10 - 0.1)
    v6 := GenS1{F0: -9223372036854775808, F1: -128}
    v7 := GenS2{F0: 2, F1: ("a" + ((("a" + "foo bar") + v4) + ("a" + (v4 + v4))))}
    var v8 uint16 = 240
    var v9 int = 933
    switch v9 {
    case -2147483648:
        for fi10 := 0; fi10 < 17; fi10++ {
            if (!false) {
                var v11 uint16 = (v8 + 65535)
                var v12 uint16 = (v8 | 1)
                _ = v11
                _ = v12
            } else {
                var v13 int64 = 9223372036854775807
                if true {
                    var v14 uint16 = v8
                    _ = v14
                } else {
                    v9 -= 4294967295
                }
                _ = v13
            }
            _ = fi10
        }
    }
    v5 = 3.14159
    _ = len(v2)
    if true {
        if true {
            v15 := GenS3{F0: -281, F1: (0 < -657)}
            v16 := GenS4{F0: (127 != 154), F1: 0}
            v17 := GenS5{F0: (v9 >> ((935) & 63)), F1: 32767}
            _ = v15
            _ = v16
            _ = v17
        }
        if false {
            delete(v2, (v4 + v4))
            switch v9 {
            case 42:
                var v18 int64 = 274
                _ = v18
            }
        }
    }
    v9 += (v9 - (^(v9)))
    if (false || true) {
        var v19 float32 = ((v3 / (v3 / 1.0)) - v3)
        _ = v19
    }
    v20 := []int{42, -520, v9}
    fmt.Println(v9, v5)
    switch v4 {
    case "\n\t":
        var v21 uint64 = uint64(v8)
        _ = v21
    case "hello":
        _ = len(v2)
    case "a":
        fmt.Println(v9)
    default:
        _ = len(v2)
    }
    v22 := v2[v4]
    for fi23 := 0; fi23 < 20; fi23++ {
        var v24 float32 = v3
        _ = fi23
        _ = v24
    }
    if false {
        var v25 int64 = int64(v22)
        switch "hello" {
        case "\n\t":
            var v26 uint32 = v22
            _ = v26
        }
        v3 -= v1
        _ = v25
    } else {
        switch v9 {
        case 32767:
            v9 = (v9 | -801)
        case 128:
            fmt.Println(v3, v9, v3)
        }
        v1 = ((v3 * 3.14159) + v3)
        v20 = append(v20, (v9 >> ((479) & 63)))
    }
    v1 = (v1 / 0.1)
    if (((((v3 - (v3 + v3)) * v1) / -2.25) >= v3) || true) {
        if false {
            delete(v2, "a")
            if len(v20) > 0 {
                v20[0] = v9
            }
        } else {
            var v27 float32 = v1
            var v28 float64 = (v5 * ((((1.5 / 1e10) / float64((1e10 / (v1 / (((((((1e10 * (3.14159 / 1e10)) - v1) / (v1 * (v3 * v27))) - 0.1) * float32(float64(v1))) / ((v3 + v3) - (float32(((v5 + v5) - (v5 / (v5 + v5)))) + v3))) / (v3 * v1)))))) - (v5 / (((1e10 / (v5 + 1.5)) * v5) + 0.1))) / (v5 / (v5 * float64(((1.5 + v1) / v3))))))
            if (1 > 4294967295) {
                v29 := v2[v4]
                for fi30 := 0; fi30 < 19; fi30++ {
                    switch (-(v9)) {
                    case 32768:
                        if len(v20) >= 2 {
                            v20 = v20[1:len(v20)]
                        }
                    }
                    _ = fi30
                }
                for fi31 := 0; fi31 < 16; fi31++ {
                    v32 := GenS6{F0: 65535}
                    _ = fi31
                    _ = v32
                }
                _ = v29
            }
            _ = v27
            _ = v28
        }
        for fi33 := 0; fi33 < 4; fi33++ {
            v4 = (v4 + v4)
            _ = fi33
        }
    } else {
        v3 += (-2.25 + v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
