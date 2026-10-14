package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 bool }
func genFn0(p0 string, p1 uint16, p2 uint32) uint {
    p0 = ("" + p0)
    for fi5 := 0; fi5 < 11; fi5++ {
        fmt.Println(fi5)
        _ = fi5
    }
    for fi6 := 0; fi6 < 16; fi6++ {
        if false {
            var v7 float32 = 0.1
            var v8 uint8 = 42
            for fi9 := 0; fi9 < 19; fi9++ {
                fmt.Println(v8, v8)
                _ = fi9
            }
            _ = v7
            _ = v8
        } else {
            v10 := []int{}
            p2 = p2
            var v11 uint32 = (p2 - p2)
            _ = v10
            _ = v11
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = p2
    return 65535
}
func genFn1(p0 int64) bool {
    for fi12 := 0; fi12 < 17; fi12++ {
        switch "foo bar" {
        case "hello":
            var v13 int8 = 100
            _ = v13
        }
        v14 := GenS1{F0: (("a" + ("" + (("" + "foo bar") + "foo bar"))) < "foo bar"), F1: 0}
        switch (^(fi12)) {
        case 32767:
            fmt.Println(fi12)
        }
        _ = fi12
        _ = v14
    }
    fmt.Println(p0)
    _ = p0
    return (128 == 128)
}

func main() {
    var v1 uint32 = 2
    v2 := []int{127, 314, -1}
    var v3 uint = 255
    v4 := make(map[string]uint32)
    var v15 int8 = -79
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if (-9223372036854775808 < 0) {
        v2 = append(v2, 128)
        v16 := GenS1{F0: (1 <= 1), F1: v1}
        v15 += 74
        _ = v16
    } else {
        var v17 int = 127
        _ = v17
    }
    for fi18 := 0; fi18 < 9; fi18++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v1 += (v1 + (v1 << (((v3 | (v3 + genFn0("foo bar", 32767, 4294967295)))) & 63)))
        if ((v1 >> ((v3) & 63)) > 557) {
            var v19 uint32 = (v1 << (((^(v3))) & 63))
            _ = v19
        } else {
            var v20 bool = false
            if len(v2) > 0 {
                v2[0] = fi18
            }
            v4["a"] = 0
            _ = v20
        }
        _ = fi18
    }
    v4[("foo bar" + "a")] = 277
    v21 := GenS1{F0: false, F1: (v1 ^ v1)}
    v22 := GenS2{F0: -32768, F1: 1}
    v23 := GenS3{F0: false}
    v24 := GenS4{F0: true}
    switch "\n\t" {
    case "\n\t":
        var v25 uint16 = 32767
        _ = v25
    case "foo bar":
        var v26 bool = ("" > "foo bar")
        _ = v26
    case "a":
        v1 = v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
