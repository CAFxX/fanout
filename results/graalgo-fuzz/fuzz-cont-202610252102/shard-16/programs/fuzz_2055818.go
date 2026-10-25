package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 uint }
type GenS2 struct { F0 uint64; F1 int8; F2 bool }
type GenS3 struct { F0 int32; F1 uint64; F2 uint16 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint16 = 94
    v3 := []int{2147483647}
    v4 := make(map[int]bool)
    v5 := make(map[uint32]bool)
    switch 1 {
    case 1:
        if true {
            if (!(0.1 < (((0.0 / -2.25) * (1e10 + 0.1)) - 1.5))) {
                var v6 uint16 = (^(v2))
                _ = v6
            } else {
                if false {
                    for fi7 := 0; fi7 < 4; fi7++ {
                        v3 = append(v3, fi7)
                        _ = fi7
                    }
                } else {
                    var v8 float64 = 0.0
                    if (!true) {
                        v8 -= (v8 + v8)
                        for fi9 := 0; fi9 < 20; fi9++ {
                            var v10 string = "hello"
                            v2 = (v2 * 255)
                            _ = fi9
                            _ = v10
                        }
                    } else {
                        v8 = (((1e10 / float64((1.5 * 1e10))) * ((1e10 * ((((float64(float32((v8 * (v8 + ((0.0 * v8) - (v8 + v8)))))) - v8) * -2.25) + v8) * float64(0.0))) - 1.5)) / 0.1)
                    }
                    _ = v8
                }
                v2 = (^(v2))
            }
            for fi11 := 0; fi11 < 7; fi11++ {
                var v12 uint = 18446744073709551615
                _ = fi11
                _ = v12
            }
            if (false || (!false)) {
                v13 := GenS1{F0: 2, F1: 42, F2: 65535}
                _ = v13
            }
        } else {
            _ = len(v3)
        }
    case 655:
        var v14 uint = 2
        _ = v14
    }
    var v15 string = "foo bar"
    switch (("" + v15) + "hello") {
    case "a":
        var v16 int = -476
        _ = v16
    case "":
        v17 := GenS2{F0: 413, F1: -1, F2: true}
        _ = v17
    case "foo bar":
        v18 := GenS1{F0: 2, F1: -631, F2: 2}
        _ = v18
    default:
        if (2147483647 >= 0) {
            var v19 int16 = -654
            if ((-(v19)) > -580) {
                v15 = v15
            }
            _ = v19
        } else {
            v2 += (v2 / ((v2) | 1))
            v20 := GenS1{F0: 0, F1: 127, F2: 291}
            _ = v20
        }
    }
    var v21 int32 = 1
    var v22 uint32 = 116
    switch (("foo bar" + v15) + ("hello" + v15)) {
    case "\n\t":
        if false {
            v15 = ""
        } else {
            v23 := GenS2{F0: 847, F1: -1, F2: false}
            if len(v3) > 0 {
                v3[0] = -751
            }
            _ = v23
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v24 uint8 = 2
    v25 := GenS3{F0: 2147483647, F1: 1, F2: (v2 << ((524) & 63))}
    v24 = (^(v24))
    v1[399] = v22
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
