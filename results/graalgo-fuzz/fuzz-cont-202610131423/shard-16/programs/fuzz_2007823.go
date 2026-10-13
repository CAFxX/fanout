package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 int16 }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 uint8; F1 int64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint64 }
func genFn0(p0 uint16) int8 {
    for fi4 := 0; fi4 < 20; fi4++ {
        v5 := GenS1{F0: -128, F1: (fi4 % ((fi4) | 1)), F2: 1}
        var v6 bool = (1 > 134)
        v7 := GenS2{F0: 1, F1: (255 != 2)}
        _ = fi4
        _ = v5
        _ = v6
        _ = v7
    }
    p0 = p0
    switch ("\n\t" + "") {
    case "":
        var v8 float32 = 0.0
        _ = v8
    case "\n\t":
        for fi9 := 0; fi9 < 17; fi9++ {
            for fi10 := 0; fi10 < 8; fi10++ {
                v11 := []int{fi10, fi9}
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                if len(v11) > 0 {
                    v11[0] = fi9
                }
                _ = fi10
                _ = v11
            }
            v12 := GenS1{F0: 127, F1: fi9, F2: 949}
            p0 -= (p0 << ((0) & 63))
            _ = fi9
            _ = v12
        }
    }
    switch "a" {
    case "foo bar":
        v13 := GenS3{F0: 1, F1: -1}
        _ = v13
    case "":
        v14 := []int{128}
        _ = v14
    case "a":
        var v15 bool = true
        _ = v15
    }
    var v16 float32 = -2.25
    _ = p0
    _ = v16
    return -78
}

func main() {
    var v1 uint = 255
    v2 := []int{}
    var v3 int64 = 779
    fmt.Println(v1)
    for fi17 := 0; fi17 < 5; fi17++ {
        var v18 int64 = v3
        _ = fi17
        _ = v18
    }
    var v19 uint32 = 146
    v3 = (v3 % ((242) | 1))
    v20 := GenS1{F0: -25, F1: int(v1), F2: -32768}
    switch 32768 {
    case 32768:
        v3 += 42
    case 128:
        var v21 uint64 = 2
        _ = v21
    }
    v19 -= 72
    for fi22 := 0; fi22 < 17; fi22++ {
        var v23 uint64 = 2
        v3 = v3
        _ = fi22
        _ = v23
    }
    fmt.Println(v19)
    v24 := GenS1{F0: -111, F1: 32768, F2: 908}
    v19 -= (^(v19))
    v2 = append(v2, int(v3))
    var v25 uint64 = 0
    v26 := GenS4{F0: int16(v19)}
    var v27 string = (("foo bar" + ("" + "hello")) + "hello")
    if (358 < 42) {
        switch 1 {
        case -9223372036854775808:
            v28 := GenS1{F0: -128, F1: -9223372036854775808, F2: 42}
            _ = v28
        case 4294967295:
            var v29 uint64 = v25
            _ = v29
        case 68:
            switch -973 {
            case 32768:
                var v30 bool = true
                _ = v30
            default:
                var v31 float32 = float32((0.1 - ((((3.14159 / 3.14159) - (1.5 / 1.5)) / 3.14159) + (0.1 + 3.14159))))
                _ = v31
            }
        default:
            var v32 uint64 = (v25 / (((v25 % ((v25) | 1))) | 1))
            _ = v32
        }
    } else {
        fmt.Println(v25)
        for fi33 := 0; fi33 < 16; fi33++ {
            var v34 int = 4294967295
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi33
            _ = v34
        }
        v35 := GenS5{F0: 42}
        _ = v35
    }
    for fi36 := 0; fi36 < 7; fi36++ {
        v1 = (v1 << ((v1) & 63))
        v37 := GenS3{F0: 2, F1: v3}
        _ = fi36
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
}
