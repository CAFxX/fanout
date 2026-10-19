package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 uint32; F2 int64 }
type GenS3 struct { F0 int16; F1 string; F2 int }

func main() {
    var v1 float64 = 3.14159
    var v2 uint8 = 40
    var v3 int8 = -104
    v4 := make(map[int]uint32)
    v5 := []int{32768, 2147483647}
    var v6 uint8 = v2
    if false {
        v7 := GenS1{F0: v3}
        var v8 int = 127
        switch v8 {
        case 0:
            for fi9 := 0; fi9 < 16; fi9++ {
                fmt.Println(fi9)
                v8 = (-(fi9))
                _ = fi9
            }
        case -1:
            var v10 uint = 828
            _ = v10
        }
        _ = v7
        _ = v8
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        v2 = v6
        v12 := GenS1{F0: v3}
        v13 := GenS1{F0: 0}
        _ = fi11
        _ = v12
        _ = v13
    }
    v14 := GenS1{F0: (v3 % ((1) | 1))}
    if true {
        if (!true) {
            v15 := GenS1{F0: v3}
            v1 = v1
            var v16 int8 = v3
            _ = v15
            _ = v16
        }
        v17 := GenS1{F0: (-(v3))}
        _ = v17
    } else {
        fmt.Println(v1)
    }
    v3 = 93
    if false {
        v6 += 0
        v2 += 0
    }
    for fi18 := 0; fi18 < 8; fi18++ {
        var v19 int16 = int16(fi18)
        if (2 >= 255) {
            for fi20 := 0; fi20 < 6; fi20++ {
                fmt.Println(v2, v1)
                _ = fi20
            }
            for fi21 := 0; fi21 < 12; fi21++ {
                var v22 int = (fi21 + fi21)
                v23 := GenS2{F0: 0, F1: 1, F2: -9223372036854775808}
                _ = fi21
                _ = v22
                _ = v23
            }
            _ = len(v4)
        } else {
            v6 -= v6
        }
        if (true && false) {
            fmt.Println(v2)
        }
        _ = fi18
        _ = v19
    }
    switch 42 {
    case 2147483647:
        v24 := GenS3{F0: -780, F1: ((("hello" + ("\n\t" + "\n\t")) + "hello") + "foo bar"), F2: 32767}
        _ = v24
    default:
        var v25 float32 = 1e10
        _ = v25
    }
    v1 = (v1 + 0.0)
    switch (("foo bar" + "\n\t") + ("foo bar" + ("hello" + ""))) {
    case "foo bar":
        v3 -= v3
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
