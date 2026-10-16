package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 uint; F1 int16; F2 uint16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint32 }
func genFn0(p0 int64) int16 {
    var v4 string = ("" + ((("hello" + ("a" + ("a" + "hello"))) + "\n\t") + ("a" + "")))
    var v5 int8 = 100
    v6 := GenS1{F0: 42, F1: (p0 / ((1) | 1))}
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 1000
}

func main() {
    var v1 uint16 = 32767
    v2 := []uint32{42, 2147483647}
    var v3 float64 = (1.5 + (((1.5 / 0.1) - 3.14159) + 0.1))
    if (((1 > 65535) || false) && ((v1 - (v1 - v1)) < (v1 ^ 0))) {
        v7 := GenS1{F0: -1, F1: 0}
        _ = v7
    }
    var v8 int = 32768
    _ = len(v2)
    switch 811 {
    case 4294967295:
        switch "foo bar" {
        case "":
            for fi9 := 0; fi9 < 2; fi9++ {
                switch (("\n\t" + "hello") + "\n\t") {
                case "":
                    for fi10 := 0; fi10 < 8; fi10++ {
                        fmt.Println(v3, fi9)
                        var v11 uint = 416
                        var v12 uint = (v11 & v11)
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                }
                _ = fi9
            }
        default:
            v1 += 0
        }
    case -621:
        var v13 string = "a"
        _ = v13
    case 0:
        for fi14 := 0; fi14 < 9; fi14++ {
            var v15 uint16 = 42
            if false {
                var v16 bool = ((false || (!(false && ((v15 + v1) == v15)))) || (!((int64(v1) != 0) || true)))
                for fi17 := 0; fi17 < 6; fi17++ {
                    v18 := GenS2{F0: 1, F1: 42, F2: (v15 >> ((2) & 63))}
                    _ = len(v2)
                    switch "hello" {
                    case "hello":
                        var v19 uint32 = 140
                        _ = v19
                    case "":
                        var v20 int8 = -113
                        _ = v20
                    }
                    _ = fi17
                    _ = v18
                }
                _ = len(v2)
                _ = v16
            } else {
                v3 += v3
                v15 = 813
            }
            _ = fi14
            _ = v15
        }
    default:
        v3 += 1.5
    }
    var v21 bool = true
    switch ("a" + "foo bar") {
    case "foo bar":
        fmt.Println(v1, v21, v1)
    case "a":
        var v22 float32 = 0.1
        _ = v22
    default:
        if len(v2) > 0 {
            v2[0] = 755
        }
    }
    fmt.Println(v1, v1, v1)
    fmt.Println(v21, v8)
    v23 := GenS3{F0: v1}
    v24 := GenS4{F0: 42}
    v2 = append(v2, 42)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v21
    _ = v23
    _ = v24
    _ = v3
    _ = v8
}
