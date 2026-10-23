package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 bool }

func main() {
    v1 := make(map[string]string)
    v2 := []string{(("" + "a") + "hello")}
    v3 := make(map[int]bool)
    var v4 int16 = 440
    if true {
        var v5 int32 = 2147483647
        if true {
            v5 += (v5 << ((255) & 63))
            v6 := GenS1{F0: 9223372036854775807, F1: false}
            for fi7 := 0; fi7 < 20; fi7++ {
                fmt.Println(fi7)
                delete(v3, -573)
                _ = fi7
            }
            _ = v6
        }
        v4 = v4
        _ = v5
    }
    v8 := v3[-9223372036854775808]
    switch "hello" {
    case "a":
        v8 = v8
    default:
        var v9 float32 = -2.25
        _ = v9
    }
    _ = len(v3)
    v10 := GenS2{F0: (^(v4)), F1: ("" + "")}
    v8 = true
    for fi11 := 0; fi11 < 7; fi11++ {
        switch "foo bar" {
        case "a":
            var v12 bool = (301 != 42)
            _ = v12
        case "\n\t":
            if v8 {
                v4 = v4
                v13 := GenS2{F0: v4, F1: "\n\t"}
                if len(v2) > 0 {
                    v2[0] = "foo bar"
                }
                _ = v13
            }
        }
        var v14 bool = v8
        v14 = v8
        _ = fi11
        _ = v14
    }
    if v8 {
        var v15 int16 = (v4 << ((392) & 63))
        for fi16 := 0; fi16 < 19; fi16++ {
            v17 := GenS2{F0: (v4 / (((v15 * v15)) | 1)), F1: ((("a" + "\n\t") + "hello") + "")}
            _ = fi16
            _ = v17
        }
        var v18 int = 32767
        _ = v15
        _ = v18
    }
    v19 := GenS1{F0: 427, F1: false}
    if v8 {
        v8 = v8
        for fi20 := 0; fi20 < 14; fi20++ {
            v21 := GenS2{F0: v4, F1: "\n\t"}
            var v22 int64 = 613
            _ = fi20
            _ = v21
            _ = v22
        }
        _ = len(v2)
    } else {
        var v23 float64 = (float64(float32(0.0)) + (3.14159 + float64(1.5)))
        var v24 float32 = (float32((v23 + -2.25)) + 0.0)
        fmt.Println(v24)
        _ = v23
        _ = v24
    }
    var v25 uint = 1
    fmt.Println(v25, v4, v25)
    var v26 uint16 = 32767
    switch "hello" {
    case "hello":
        fmt.Println(v26, v25)
    case "foo bar":
        switch 128 {
        case -423:
            v26 = v26
        case 2147483647:
            v26 = (v26 % (((v26 - v26)) | 1))
        case 176:
            v27 := v1["foo bar"]
            _ = v27
        default:
            v8 = (-9223372036854775808 <= -804)
        }
    case "a":
        if ((!true) && false) {
            switch 32767 {
            case 930:
                var v28 uint32 = 42
                _ = v28
            default:
                v29 := GenS3{F0: (int32(v4) > -965)}
                _ = v29
            }
        } else {
            if (-2147483648 > -1) {
                v25 -= v25
            }
        }
    }
    var v30 uint8 = 44
    switch "a" {
    case "a":
        v31 := GenS2{F0: v4, F1: ""}
        _ = v31
    case "foo bar":
        for fi32 := 0; fi32 < 15; fi32++ {
            var v33 float32 = 1e10
            if len(v2) > 0 {
                v2[0] = ("\n\t" + "\n\t")
            }
            v2 = append(v2, "hello")
            _ = fi32
            _ = v33
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v4
    _ = v8
}
