package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int16 }
type GenS2 struct { F0 int; F1 int16; F2 bool }
type GenS3 struct { F0 int8; F1 uint32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint }
func genFn0(p0 int8, p1 uint8) uint64 {
    switch "\n\t" {
    case "a":
        p0 -= (^(p0))
    case "":
        for fi6 := 0; fi6 < 16; fi6++ {
            p1 += p1
            switch "\n\t" {
            case "hello":
                p1 = p1
            }
            _ = fi6
        }
    default:
        var v7 uint = 1
        _ = v7
    }
    fmt.Println(p1)
    v8 := GenS1{F0: 4294967295, F1: -939, F2: -1}
    for fi9 := 0; fi9 < 9; fi9++ {
        for fi10 := 0; fi10 < 19; fi10++ {
            var v11 uint = 0
            _ = fi10
            _ = v11
        }
        var v12 bool = (fi9 >= fi9)
        _ = fi9
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v8
    return 580
}

func main() {
    var v1 uint8 = 42
    v2 := []uint32{2, 4294967295, 2, 659}
    v3 := []string{("\n\t" + "\n\t"), ("foo bar" + ("hello" + ("hello" + "a")))}
    var v4 int32 = -272
    var v5 string = "a"
    var v13 float64 = (0.1 / (1e10 + 1e10))
    v3 = append(v3, v5)
    v2 = append(v2, 528)
    if (1 >= 9223372036854775807) {
        v1 += 1
        for fi14 := 0; fi14 < 13; fi14++ {
            v13 = v13
            _ = fi14
        }
    } else {
        v1 = (v1 + 255)
        for fi15 := 0; fi15 < 5; fi15++ {
            v2 = append(v2, 2)
            v16 := GenS2{F0: 32767, F1: 42, F2: true}
            for fi17 := 0; fi17 < 20; fi17++ {
                var v18 int = fi17
                if true {
                    if true {
                        var v19 int16 = -241
                        v13 += float64(1.5)
                        v19 += (v19 - (v19 * v19))
                        var v20 string = ("a" + "foo bar")
                        v4 = v4
                        _ = v19
                        _ = v20
                    }
                    if true {
                        var v21 int64 = 227
                        fmt.Println(v18, fi17, v1)
                        _ = v21
                    }
                    if true {
                        v18 = 128
                        v22 := GenS3{F0: 42, F1: 959}
                        v2 = append(v2, 431)
                        _ = v22
                    } else {
                        var v23 string = ("a" + v5)
                        fmt.Println(v13, fi17)
                        _ = v23
                    }
                }
                _ = fi17
                _ = v18
            }
            _ = fi15
            _ = v16
        }
        v24 := GenS1{F0: 1, F1: 2147483647, F2: 0}
        _ = v24
    }
    v1 = (v1 ^ v1)
    var v25 uint16 = 255
    var v26 uint64 = 42
    v27 := GenS4{F0: (v4 << ((18446744073709551615) & 63))}
    fmt.Println(v26, v5)
    switch "hello" {
    case "hello":
        v26 = 4294967295
    case "\n\t":
        v28 := GenS1{F0: 4294967295, F1: 2147483647, F2: 863}
        _ = v28
    }
    v13 += v13
    fmt.Println(v13, v5)
    for fi29 := 0; fi29 < 18; fi29++ {
        v5 = (v5 + v5)
        var v30 int = fi29
        _ = fi29
        _ = v30
    }
    var v31 int16 = int16(v25)
    switch 32768 {
    case 4294967295:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    _ = len(v2)
    v32 := GenS5{F0: 65535}
    v5 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
