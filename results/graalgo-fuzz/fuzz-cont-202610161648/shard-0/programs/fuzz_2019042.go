package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 uint16 }
type GenS2 struct { F0 int; F1 int32 }
type GenS3 struct { F0 int16; F1 string }
type GenS4 struct { F0 uint32; F1 bool; F2 int64 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint32; F1 int8 }
func genFn0(p0 int64, p1 uint16) int8 {
    if (!false) {
        for fi5 := 0; fi5 < 10; fi5++ {
            var v6 uint32 = 2
            switch "a" {
            case "foo bar":
                var v7 float32 = (((3.14159 * ((0.0 / 0.1) - 0.0)) + (0.1 - float32(0.1))) + 0.1)
                _ = v7
            case "a":
                p1 -= 255
            }
            _ = fi5
            _ = v6
        }
        v8 := GenS1{F0: 0, F1: 209, F2: (p1 * (p1 - (p1 / ((65535) | 1))))}
        _ = v8
    }
    var v9 uint8 = 128
    switch -433 {
    case 2147483647:
        for fi10 := 0; fi10 < 13; fi10++ {
            v11 := GenS1{F0: 0, F1: 2, F2: p1}
            for fi12 := 0; fi12 < 16; fi12++ {
                switch fi10 {
                case 482:
                    fmt.Println(fi10, p1)
                case 42:
                    fmt.Println(p0, p1, p0)
                case 32767:
                    v13 := GenS1{F0: 589, F1: v9, F2: (p1 << ((255) & 63))}
                    _ = v13
                }
                _ = fi12
            }
            v14 := GenS2{F0: (fi10 << ((255) & 63)), F1: 42}
            _ = fi10
            _ = v11
            _ = v14
        }
    case 156:
        var v15 string = "hello"
        _ = v15
    case -702:
        var v16 int16 = int16(v9)
        _ = v16
    default:
        var v17 int16 = -618
        _ = v17
    }
    if ((^(p1)) <= p1) {
        fmt.Println(v9)
        v9 -= (^(v9))
        v18 := GenS3{F0: 32767, F1: ""}
        _ = v18
    } else {
        if false {
            v19 := []int{}
            _ = v19
        }
        var v20 int16 = 42
        v20 = 760
        _ = v20
    }
    _ = p0
    _ = p1
    _ = v9
    return -128
}

func main() {
    v1 := []string{}
    v2 := make(map[string]int)
    var v3 uint8 = 127
    var v4 uint8 = 128
    var v21 int64 = 1
    for fi22 := 0; fi22 < 5; fi22++ {
        for fi23 := 0; fi23 < 4; fi23++ {
            v4 = 69
            _ = fi23
        }
        for fi24 := 0; fi24 < 3; fi24++ {
            v25 := GenS4{F0: 2147483647, F1: false, F2: (v21 + (-(v21)))}
            _ = fi24
            _ = v25
        }
        var v26 uint = 42
        _ = fi22
        _ = v26
    }
    fmt.Println(v4, v3)
    fmt.Println(v3, v4, v4)
    var v27 string = "a"
    v3 -= 189
    v27 = ((v27 + v27) + v27)
    v27 = v27
    v3 = v4
    v27 = ((v27 + "foo bar") + "")
    switch 670 {
    case 846:
        switch v27 {
        case "a":
            for fi28 := 0; fi28 < 19; fi28++ {
                if true {
                    v29 := GenS5{F0: 32767}
                    fmt.Println(v4)
                    _ = v29
                } else {
                    for fi30 := 0; fi30 < 13; fi30++ {
                        v31 := v2[("foo bar" + "")]
                        _ = fi30
                        _ = v31
                    }
                    var v32 int = 128
                    v33 := GenS6{F0: 65535, F1: 108}
                    _ = v32
                    _ = v33
                }
                v34 := GenS6{F0: 65535, F1: 60}
                _ = fi28
                _ = v34
            }
        case "\n\t":
            _ = len(v2)
        default:
            v3 += v3
        }
    case 128:
        v4 = v4
    default:
        var v35 bool = false
        _ = v35
    }
    var v36 uint8 = 0
    v37 := GenS4{F0: 2, F1: (103 > 42), F2: v21}
    v27 = v27
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v27
    _ = v3
    _ = v36
    _ = v37
    _ = v4
}
