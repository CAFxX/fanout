package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8 }
func genFn0(p0 int) uint64 {
    var v5 uint8 = 255
    switch "foo bar" {
    case "":
        v5 += (v5 >> ((0) & 63))
    case "a":
        var v6 uint = 954
        _ = v6
    case "hello":
        for fi7 := 0; fi7 < 5; fi7++ {
            var v8 int64 = 830
            var v9 uint64 = 18446744073709551615
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    _ = p0
    _ = v5
    return 614
}

func main() {
    var v1 uint = 0
    var v2 uint = v1
    var v3 uint32 = 65535
    var v4 bool = true
    for fi10 := 0; fi10 < 3; fi10++ {
        v11 := GenS1{F0: 100, F1: 0, F2: 0}
        var v12 float32 = 0.0
        v13 := GenS2{F0: v4}
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    var v14 uint8 = 128
    fmt.Println(v1)
    var v15 int8 = 42
    var v16 int8 = -128
    for fi17 := 0; fi17 < 8; fi17++ {
        v3 = (v3 / (((v3 % ((4294967295) | 1))) | 1))
        _ = fi17
    }
    var v18 int16 = 537
    for fi19 := 0; fi19 < 7; fi19++ {
        switch "foo bar" {
        case "":
            var v20 uint8 = (^(v14))
            _ = v20
        case "hello":
            if (2 <= 4294967295) {
                v21 := GenS1{F0: v14, F1: 712, F2: v3}
                v22 := []int{2147483647, (-(fi19)), -267, -115}
                _ = v21
                _ = v22
            }
        case "\n\t":
            for fi23 := 0; fi23 < 6; fi23++ {
                v14 = (v14 % (((v14 ^ (v14 - (^(v14))))) | 1))
                v24 := GenS3{F0: (^(v14))}
                var v25 float32 = 3.14159
                _ = fi23
                _ = v24
                _ = v25
            }
        }
        _ = fi19
    }
    if v4 {
        v26 := GenS1{F0: (v14 + (v14 - (v14 * 127))), F1: uint16(v14), F2: (v3 % (((v3 % ((v3) | 1))) | 1))}
        _ = v26
    }
    var v27 int16 = -210
    switch "\n\t" {
    case "hello":
        if v4 {
            var v28 int = 4294967295
            _ = v28
        } else {
            v3 -= 1
        }
    case "\n\t":
        for fi29 := 0; fi29 < 18; fi29++ {
            v30 := []int{fi29, fi29, (fi29 & (^(fi29))), (fi29 & fi29)}
            _ = fi29
            _ = v30
        }
    case "foo bar":
        var v31 int = 2147483647
        _ = v31
    default:
        v3 = 65535
    }
    var v32 bool = (0 > -772)
    switch -2147483648 {
    case 9223372036854775807:
        v27 -= 1
    case 485:
        switch "hello" {
        case "":
            v33 := []int{628, 2147483647}
            _ = v33
        }
    default:
        for fi34 := 0; fi34 < 9; fi34++ {
            var v35 int16 = (^(v27))
            _ = fi34
            _ = v35
        }
    }
    v36 := []int{263}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v27
    _ = v3
    _ = v32
    _ = v36
    _ = v4
}
