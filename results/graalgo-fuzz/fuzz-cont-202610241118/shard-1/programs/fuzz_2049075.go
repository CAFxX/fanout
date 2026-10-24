package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8; F1 uint32 }
type GenS4 struct { F0 int32 }
func genFn0(p0 uint32, p1 int) uint64 {
    for fi6 := 0; fi6 < 7; fi6++ {
        if false {
            var v7 int8 = 72
            _ = v7
        }
        _ = fi6
    }
    if (p0 > p0) {
        if (4294967295 >= 42) {
            p1 += p1
            var v8 uint32 = p0
            var v9 uint8 = 127
            _ = v8
            _ = v9
        }
    }
    var v10 uint16 = 434
    var v11 string = (("hello" + ("a" + ((("" + ("foo bar" + ("hello" + ("a" + "")))) + "\n\t") + ("foo bar" + "foo bar")))) + "hello")
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 18446744073709551615
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]bool)
    v3 := make(map[string]string)
    v4 := []uint32{}
    v5 := make(map[int]string)
    for fi12 := 0; fi12 < 6; fi12++ {
        var v13 bool = false
        switch fi12 {
        case 42:
            var v14 uint16 = 32767
            _ = v14
        }
        for fi15 := 0; fi15 < 8; fi15++ {
            var v16 uint = 2
            var v17 float32 = float32((1e10 + (3.14159 / (3.14159 + 1e10))))
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = fi12
        _ = v13
    }
    var v18 uint = 658
    var v19 int32 = -887
    switch ("hello" + "\n\t") {
    case "":
        v18 -= 88
    }
    v5[9223372036854775807] = "a"
    v19 = (v19 | (v19 >> ((v18) & 63)))
    for fi20 := 0; fi20 < 7; fi20++ {
        v19 -= (^(v19))
        _ = fi20
    }
    v18 = v18
    if false {
        if len(v4) > 0 {
            v4[0] = 219
        }
        if false {
            var v21 uint16 = 32767
            _ = v21
        }
        if (1e10 >= 1.5) {
            v22 := GenS1{F0: 86}
            var v23 int8 = -1
            v1[569] = 2147483648
            _ = v22
            _ = v23
        }
    }
    fmt.Println(v18, v19)
    for fi24 := 0; fi24 < 16; fi24++ {
        fmt.Println(fi24)
        fmt.Println(v18, v18, v19)
        v25 := GenS2{F0: 93}
        _ = fi24
        _ = v25
    }
    var v26 int16 = 0
    if ((88 >= 0) || false) {
        if false {
            v4 = append(v4, 1)
        }
    }
    if (651 <= 0) {
        v27 := GenS1{F0: 0}
        var v28 float64 = float64(((((1.5 - 0.1) + 1e10) / 1.0) - ((((0.0 * -2.25) + 0.1) + (0.1 - 0.1)) / 1.0)))
        fmt.Println(v18, v19, v26)
        _ = v27
        _ = v28
    } else {
        switch "foo bar" {
        case "a":
            var v29 uint32 = 42
            _ = v29
        case "\n\t":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case "foo bar":
            var v30 uint64 = 42
            _ = v30
        default:
            switch "foo bar" {
            case "\n\t":
                for fi31 := 0; fi31 < 18; fi31++ {
                    fmt.Println(v26, fi31)
                    var v32 int8 = 127
                    var v33 uint16 = 255
                    _ = fi31
                    _ = v32
                    _ = v33
                }
            case "foo bar":
                v4 = append(v4, 2147483647)
            case "hello":
                v34 := GenS3{F0: 1, F1: 2}
                _ = v34
            default:
                var v35 uint64 = 2
                _ = v35
            }
        }
        for fi36 := 0; fi36 < 14; fi36++ {
            v37 := GenS4{F0: 2147483647}
            _ = fi36
            _ = v37
        }
    }
    _ = len(v5)
    switch 906 {
    case 582:
        delete(v3, "foo bar")
    }
    for fi38 := 0; fi38 < 9; fi38++ {
        fmt.Println(v18, v18)
        switch fi38 {
        case 917:
            v39 := GenS2{F0: 42}
            _ = v39
        case 0:
            fmt.Println(fi38, fi38, v26)
        default:
            v26 -= (^(v26))
        }
        var v40 int32 = v19
        _ = fi38
        _ = v40
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
