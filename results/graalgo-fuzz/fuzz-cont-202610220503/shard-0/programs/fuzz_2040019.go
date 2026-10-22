package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16; F1 int8 }
func genFn0(p0 int64) int64 {
    switch 4294967295 {
    case 477:
        var v7 float64 = (((0.1 + (3.14159 + (3.14159 - (1e10 - -2.25)))) + 3.14159) - ((0.1 + 0.0) * (-2.25 / (((float64(float32(1.5)) / float64((0.0 + (-2.25 + -2.25)))) + ((0.1 + ((0.1 - 1e10) * (((-2.25 + -2.25) - -2.25) + 0.0))) / 3.14159)) - float64((1e10 - 0.1))))))
        _ = v7
    case 718:
        var v8 uint8 = 255
        _ = v8
    case 232:
        p0 -= p0
    default:
        p0 -= p0
    }
    var v9 uint8 = 255
    if false {
        v10 := GenS1{F0: 0, F1: 982}
        v11 := []int{1, -662}
        for fi12 := 0; fi12 < 8; fi12++ {
            v13 := GenS1{F0: 953, F1: 1}
            _ = fi12
            _ = v13
        }
        _ = v10
        _ = v11
    }
    v14 := []int{-916}
    v15 := GenS1{F0: 255, F1: 2}
    _ = p0
    _ = v14
    _ = v15
    _ = v9
    return (p0 | (p0 / ((p0) | 1)))
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]bool)
    var v3 int8 = 100
    var v4 int8 = v3
    var v5 int8 = 0
    v6 := make(map[uint32]uint32)
    v16 := []int{127, 127, 1, -1}
    for fi17 := 0; fi17 < 5; fi17++ {
        var v18 bool = (!false)
        v3 -= int8(fi17)
        var v19 int = (fi17 * fi17)
        _ = fi17
        _ = v18
        _ = v19
    }
    v20 := GenS1{F0: 1, F1: 1}
    var v21 uint16 = 928
    fmt.Println(v3, v5)
    if (0 > 32767) {
        var v22 int16 = -1
        _ = v22
    } else {
        var v23 uint64 = 4294967295
        switch ("a" + "\n\t") {
        case "hello":
            if true {
                v23 += v23
                switch "hello" {
                case "foo bar":
                    v24 := GenS1{F0: (v21 << ((65535) & 63)), F1: 42}
                    _ = v24
                case "a":
                    for fi25 := 0; fi25 < 10; fi25++ {
                        var v26 uint = 42
                        _ = len(v6)
                        _ = fi25
                        _ = v26
                    }
                default:
                    if (1000 != 401) {
                        fmt.Println(v3)
                        var v27 uint = uint(v23)
                        _ = v27
                    }
                }
            }
        default:
            v28 := GenS2{F0: (v4 % (((v5 % ((1) | 1))) | 1))}
            _ = v28
        }
        _ = v23
    }
    v29 := GenS2{F0: 100}
    var v30 uint = 0
    delete(v1, "\n\t")
    var v31 int8 = (v5 % ((v3) | 1))
    for fi32 := 0; fi32 < 13; fi32++ {
        switch "hello" {
        case "\n\t":
            var v33 uint32 = 42
            _ = v33
        case "foo bar":
            fmt.Println(v3)
        case "hello":
            v34 := GenS3{F0: (v21 ^ (v21 - (v21 | 65535))), F1: v31}
            _ = v34
        }
        if (!false) {
            v31 = v31
            switch "" {
            case "":
                var v35 float32 = (0.0 * float32(1.5))
                _ = v35
            case "foo bar":
                var v36 int16 = -32768
                _ = v36
            default:
                if (!true) {
                    var v37 int64 = 0
                    if false {
                        v38 := GenS2{F0: v5}
                        _ = v38
                    }
                    _ = v37
                } else {
                    v21 = v21
                    if len(v16) >= 2 {
                        v16 = v16[1:len(v16)]
                    }
                    v31 = v4
                }
            }
        }
        _ = len(v16)
        _ = fi32
    }
    var v39 uint = (v30 / ((v30) | 1))
    v40 := GenS2{F0: (v3 % (((v3 << (((v39 & v39)) & 63))) | 1))}
    var v41 float32 = (((-2.25 / 1.5) * 3.14159) + (-2.25 + 3.14159))
    delete(v2, -1)
    fmt.Println(v5, v3, v5)
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v39, v30)
    case "foo bar":
        var v42 uint64 = 18446744073709551615
        _ = v42
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
