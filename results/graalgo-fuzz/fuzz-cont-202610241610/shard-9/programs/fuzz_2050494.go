package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 uint32; F1 int8; F2 int16 }
type GenS3 struct { F0 int; F1 int }
func genFn0(p0 int, p1 int64, p2 int16) uint16 {
    v4 := GenS1{F0: false, F1: (p1 >> ((255) & 63))}
    p0 = (p0 + (p0 * p0))
    var v5 string = "foo bar"
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 255
}
func genFn1() uint64 {
    var v6 string = ((("" + "hello") + "hello") + "a")
    v7 := []int{}
    v8 := GenS1{F0: (float64(1e10) >= 1.5), F1: 9223372036854775807}
    _ = v6
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    var v1 float32 = (3.14159 - float32(1.5))
    v2 := []uint32{}
    v3 := make(map[int]int)
    v1 = 1.5
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v1 = (v1 - float32(((0.1 - 3.14159) + 1e10)))
    for fi9 := 0; fi9 < 14; fi9++ {
        var v10 bool = true
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 6; fi11++ {
        v1 -= v1
        _ = fi11
    }
    v1 -= 0.0
    if true {
        for fi12 := 0; fi12 < 3; fi12++ {
            if false {
                v13 := GenS1{F0: true, F1: 544}
                v14 := GenS1{F0: true, F1: 42}
                _ = v13
                _ = v14
            }
            var v15 uint8 = 255
            _ = fi12
            _ = v15
        }
        switch ("foo bar" + "foo bar") {
        case "a":
            var v16 uint = 626
            _ = v16
        case "hello":
            v17 := GenS1{F0: (42 >= 130), F1: 42}
            _ = v17
        default:
            for fi18 := 0; fi18 < 6; fi18++ {
                if false {
                    v1 = 3.14159
                    var v19 int = (fi18 / (((fi18 ^ fi18)) | 1))
                    _ = v19
                }
                if false {
                    var v20 uint = 1
                    var v21 uint = 0
                    fmt.Println(v1, v21)
                    _ = v20
                    _ = v21
                } else {
                    if (!true) {
                        v1 = (float32(0.1) * ((-2.25 * v1) * (v1 - v1)))
                        v1 = v1
                        for fi22 := 0; fi22 < 1; fi22++ {
                            switch "hello" {
                            case "a":
                                v1 -= 3.14159
                            case "":
                                v23 := GenS1{F0: true, F1: -551}
                                _ = v23
                            case "hello":
                                _ = len(v2)
                            default:
                                var v24 string = ("\n\t" + "foo bar")
                                _ = v24
                            }
                            v1 += v1
                            _ = fi22
                        }
                    }
                    var v25 uint16 = 393
                    v26 := GenS2{F0: 1, F1: -72, F2: -32768}
                    _ = v25
                    _ = v26
                }
                v27 := GenS2{F0: 2147483647, F1: 78, F2: -1}
                _ = fi18
                _ = v27
            }
        }
    } else {
        v28 := GenS3{F0: -547, F1: 42}
        fmt.Println(v1)
        v29 := v3[2147483647]
        _ = v28
        _ = v29
    }
    var v30 uint = 2
    var v31 float32 = (v1 * float32(1.5))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
}
