package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint32; F1 bool; F2 uint16 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int16; F1 int32; F2 int64 }
type GenS8 struct { F0 string; F1 int64; F2 bool }
func genFn0(p0 uint16) string {
    p0 -= 255
    v7 := []int{707, 4294967295}
    p0 = p0
    var v8 uint16 = (p0 / (((p0 ^ p0)) | 1))
    _ = p0
    _ = v7
    _ = v8
    return ("foo bar" + ("" + ("" + "hello")))
}

func main() {
    v1 := []int{9223372036854775807, 1, 32767}
    var v2 int16 = -32768
    var v3 int64 = 0
    var v4 string = ((("hello" + "hello") + (("foo bar" + "hello") + "")) + "a")
    var v5 int64 = v3
    v6 := make(map[string]bool)
    v9 := GenS1{F0: 0, F1: (1.5 > (1.5 * (((3.14159 - (-2.25 * -2.25)) - 1e10) / float32(1e10))))}
    v10 := v6[((("\n\t" + "foo bar") + (((v4 + ((v4 + ((("hello" + v4) + "\n\t") + "\n\t")) + (v4 + v4))) + ("a" + "foo bar")) + "\n\t")) + v4)]
    v11 := GenS2{F0: 4294967295}
    fmt.Println(v3, v10, v3)
    var v12 string = (v4 + v4)
    if v10 {
        v1 = append(v1, 127)
        var v13 float64 = 0.0
        _ = v13
    } else {
        if false {
            fmt.Println(v2, v2)
            v14 := GenS3{F0: 2}
            _ = v14
        }
    }
    var v15 int64 = v3
    for fi16 := 0; fi16 < 8; fi16++ {
        switch (fi16 >> ((0) & 63)) {
        case -9223372036854775808:
            for fi17 := 0; fi17 < 2; fi17++ {
                var v18 uint8 = 2
                v15 = -138
                _ = fi17
                _ = v18
            }
        }
        v10 = v10
        fmt.Println(fi16, v12, v10)
        _ = fi16
    }
    var v19 int8 = 0
    _ = len(v6)
    switch 17 {
    case 128:
        v3 += v3
    case -1:
        v20 := GenS4{F0: 2147483647, F1: true, F2: 255}
        _ = v20
    case 4294967295:
        v21 := GenS3{F0: 422}
        _ = v21
    default:
        if ((-699 <= -481) && false) {
            v22 := GenS4{F0: 2147483648, F1: false, F2: 768}
            v19 = v19
            _ = v22
        }
    }
    if v10 {
        v23 := GenS3{F0: 2147483648}
        var v24 int32 = 772
        v3 -= v5
        _ = v23
        _ = v24
    } else {
        _ = len(v1)
        var v25 uint16 = 255
        if v10 {
            v12 = "\n\t"
            v26 := GenS5{F0: 983}
            switch 946 {
            case 128:
                if ((v4 > v12) || false) {
                    var v27 int32 = 763
                    v10 = v10
                    switch -148 {
                    case 4294967295:
                        v28 := GenS6{F0: 64}
                        _ = v28
                    }
                    _ = v27
                } else {
                    v4 = ((("a" + (v4 + (v4 + (v4 + "\n\t")))) + "foo bar") + v4)
                }
            }
            _ = v26
        }
        _ = v25
    }
    v29 := GenS1{F0: -2147483648, F1: v10}
    v30 := GenS7{F0: -788, F1: -1, F2: -1}
    var v31 int8 = (v19 % (((v19 / (((v19 << ((65535) & 63))) | 1))) | 1))
    switch "hello" {
    case "foo bar":
        switch "" {
        case "foo bar":
            v2 -= -32768
        default:
            if (2 < 331) {
                switch "" {
                case "\n\t":
                    v15 += (v5 + 970)
                }
            }
        }
    case "a":
        v32 := GenS8{F0: "hello", F1: v15, F2: (false || v10)}
        _ = v32
    case "hello":
        fmt.Println(v10, v12)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v19
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
