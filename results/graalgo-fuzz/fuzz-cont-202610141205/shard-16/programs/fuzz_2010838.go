package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint; F1 string; F2 uint16 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 int16) int8 {
    fmt.Println(p0)
    for fi5 := 0; fi5 < 13; fi5++ {
        var v6 uint16 = 32767
        p0 -= p0
        for fi7 := 0; fi7 < 4; fi7++ {
            if false {
                var v8 uint64 = 18446744073709551615
                _ = v8
            }
            v6 -= (v6 / (((v6 >> ((65535) & 63))) | 1))
            _ = fi7
        }
        _ = fi5
        _ = v6
    }
    _ = p0
    return 1
}

func main() {
    v1 := []string{"", "hello", ("a" + "foo bar"), "hello"}
    var v2 bool = (!true)
    var v3 int16 = 642
    var v4 uint32 = 4294967295
    for fi9 := 0; fi9 < 13; fi9++ {
        fmt.Println(fi9)
        for fi10 := 0; fi10 < 3; fi10++ {
            if (fi10 <= int(v3)) {
                var v11 float32 = 3.14159
                _ = v11
            }
            var v12 int32 = -1
            var v13 float32 = ((1e10 / 3.14159) - 1.5)
            _ = fi10
            _ = v12
            _ = v13
        }
        v3 -= (v3 - (v3 >> ((817) & 63)))
        _ = fi9
    }
    if ((v2 || v2) || v2) {
        v4 += v4
        v14 := GenS1{F0: -1}
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v14
    }
    if len(v1) > 0 {
        v1[0] = ""
    }
    switch "\n\t" {
    case "foo bar":
        v15 := GenS1{F0: 694}
        _ = v15
    case "\n\t":
        var v16 int = -2147483648
        _ = v16
    case "a":
        if ((!false) || false) {
            v3 = v3
            fmt.Println(v3)
        }
    default:
        var v17 uint32 = v4
        _ = v17
    }
    var v18 float64 = -2.25
    v19 := GenS2{F0: 255, F1: ("foo bar" + ("a" + ("hello" + "\n\t"))), F2: 255}
    v3 = (v3 >> ((857) & 63))
    if true {
        fmt.Println(v3, v18)
        if len(v1) > 0 {
            v1[0] = ("a" + (((("" + "hello") + "a") + "foo bar") + "foo bar"))
        }
        v20 := GenS3{F0: v4}
        _ = v20
    }
    var v21 uint32 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
