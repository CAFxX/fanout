package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 string; F1 uint16; F2 int }
type GenS4 struct { F0 int64 }

func main() {
    v1 := make(map[int]string)
    var v2 int = 1
    v3 := []uint32{4294967295, 1, 2, 493}
    var v4 uint = 2
    v5 := make(map[int]uint32)
    var v6 int16 = 0
    var v7 bool = ((18446744073709551615 > 311) || false)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v8 int = 42
    v3 = append(v3, 42)
    var v9 int8 = 0
    v9 += 59
    if (-40 != v9) {
        for fi10 := 0; fi10 < 14; fi10++ {
            v11 := GenS1{F0: 1, F1: (599 == 654)}
            _ = fi10
            _ = v11
        }
    } else {
        fmt.Println(v2, v6)
        for fi12 := 0; fi12 < 11; fi12++ {
            v13 := GenS2{F0: 2147483648, F1: (fi12 ^ 1)}
            _ = fi12
            _ = v13
        }
        v7 = v7
    }
    var v14 uint64 = 299
    v9 -= (v9 >> ((v4) & 63))
    switch "hello" {
    case "a":
        if ((1e10 + 0.1) != float32(((1e10 * -2.25) + 0.1))) {
            v4 += v4
            var v15 string = (("" + "\n\t") + "a")
            _ = v15
        }
    case "hello":
        v9 = v9
    case "":
        v2 += int(v4)
    default:
        v9 = (v9 - 0)
    }
    switch (-(v2)) {
    case 9223372036854775807:
        var v16 int = (v8 - v8)
        _ = v16
    case -1:
        var v17 uint16 = 42
        _ = v17
    default:
        if v7 {
            v18 := GenS3{F0: ("hello" + ""), F1: 42, F2: (v8 - v8)}
            v3 = append(v3, 42)
            _ = v18
        } else {
            v7 = false
            if len(v3) > 0 {
                v3[0] = 2147483648
            }
        }
    }
    var v19 bool = false
    for fi20 := 0; fi20 < 20; fi20++ {
        for fi21 := 0; fi21 < 17; fi21++ {
            switch "foo bar" {
            case "":
                var v22 int = int(v9)
                _ = v22
            case "a":
                fmt.Println(v2, v14, v6)
            case "hello":
                if (v19 && false) {
                    switch (v2 * 128) {
                    case 770:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    default:
                        v7 = v7
                    }
                }
            }
            var v23 uint8 = 223
            _ = fi21
            _ = v23
        }
        if ((1.5 / float64(float32(((1e10 * ((0.1 * (1e10 + ((-2.25 * (((float64(1e10) - -2.25) + (-2.25 + ((3.14159 / (1e10 / 1.0)) + 1e10))) / 1e10)) + ((3.14159 / ((3.14159 + 0.1) - (0.0 / -2.25))) * -2.25)))) / (-2.25 / 1.0))) * 1e10)))) >= 0.1) {
            v6 = v6
            fmt.Println(fi20)
            v24 := GenS2{F0: 1, F1: (fi20 | v8)}
            _ = v24
        }
        _ = fi20
    }
    v6 = (v6 / ((32767) | 1))
    v25 := GenS4{F0: -9223372036854775808}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
