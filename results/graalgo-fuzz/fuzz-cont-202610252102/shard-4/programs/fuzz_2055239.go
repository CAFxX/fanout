package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 uint16; F1 int16; F2 bool }
type GenS4 struct { F0 uint32; F1 uint16 }
type GenS5 struct { F0 string }
func genFn0(p0 uint8, p1 uint16) int32 {
    if false {
        var v7 string = ("\n\t" + "foo bar")
        p0 -= (p0 & p0)
        v8 := GenS1{F0: 65535}
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: 224}
    _ = p0
    _ = p1
    _ = v9
    return -127
}

func main() {
    var v1 uint16 = 65535
    var v2 uint16 = (v1 | (v1 - 1))
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]bool)
    v5 := []string{"foo bar", "a"}
    var v6 int64 = -21
    for fi10 := 0; fi10 < 15; fi10++ {
        fmt.Println(v6, fi10)
        v11 := GenS2{F0: -128, F1: 2147483648}
        v1 += v2
        _ = fi10
        _ = v11
    }
    v6 = v6
    if true {
        var v12 int8 = int8(v6)
        if true {
            var v13 int32 = 0
            _ = v13
        }
        if (-2.25 < ((0.1 + ((1e10 / (-2.25 / 3.14159)) - ((1e10 - -2.25) * 3.14159))) * 0.0)) {
            var v14 uint8 = 128
            _ = v14
        } else {
            var v15 float64 = 1.5
            v12 = v12
            _ = v15
        }
        _ = v12
    } else {
        v16 := v4[79]
        delete(v4, 1)
        for fi17 := 0; fi17 < 12; fi17++ {
            for fi18 := 0; fi18 < 16; fi18++ {
                v19 := v4[-105]
                fmt.Println(fi18)
                v19 = (1.5 == -2.25)
                v5 = append(v5, (((("" + "a") + "hello") + ("foo bar" + "")) + "foo bar"))
                _ = fi18
                _ = v19
            }
            v5 = append(v5, "a")
            v20 := GenS1{F0: 0}
            _ = fi17
            _ = v20
        }
        _ = v16
    }
    for fi21 := 0; fi21 < 9; fi21++ {
        v6 = 198
        _ = fi21
    }
    var v22 bool = ("" >= "\n\t")
    v23 := GenS3{F0: (v1 * (^(v2))), F1: 0, F2: false}
    v24 := GenS2{F0: 21, F1: uint32(v1)}
    for fi25 := 0; fi25 < 5; fi25++ {
        var v26 float64 = 1.5
        for fi27 := 0; fi27 < 14; fi27++ {
            v28 := GenS1{F0: 4294967295}
            var v29 bool = (float64(-2.25) > ((v26 / float64(1e10)) + v26))
            _ = fi27
            _ = v28
            _ = v29
        }
        _ = fi25
        _ = v26
    }
    v30 := GenS1{F0: 42}
    fmt.Println(v2, v2, v2)
    switch ("hello" + "a") {
    case "\n\t":
        v2 += (v1 * (v1 ^ v2))
    case "foo bar":
        var v31 int64 = v6
        _ = v31
    case "a":
        if (!false) {
            fmt.Println(v2, v1, v1)
            v1 -= v2
            var v32 bool = false
            _ = v32
        } else {
            v33 := GenS4{F0: 65535, F1: 0}
            _ = v33
        }
    default:
        v2 = 2
    }
    v34 := GenS3{F0: (v1 * v1), F1: -32768, F2: (1.5 != (0.1 + float64(0.1)))}
    for fi35 := 0; fi35 < 20; fi35++ {
        if ((-122 < int32(v1)) && (!(!v22))) {
            v36 := GenS5{F0: ("" + ("foo bar" + "foo bar"))}
            _ = v36
        }
        v37 := GenS4{F0: 4294967295, F1: (v1 ^ (^(v1)))}
        fmt.Println(fi35, v22, v1)
        _ = fi35
        _ = v37
    }
    if len(v5) > 0 {
        v5[0] = "\n\t"
    }
    var v38 int16 = -1
    switch "foo bar" {
    case "a":
        var v39 uint64 = 18446744073709551615
        _ = v39
    default:
        _ = len(v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v30
    _ = v34
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
