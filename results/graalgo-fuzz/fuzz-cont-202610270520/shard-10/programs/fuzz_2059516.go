package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 int16; F1 int64; F2 uint }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int64 }
func genFn0(p0 bool) int16 {
    p0 = (128 <= 127)
    v6 := []int{}
    v7 := GenS1{F0: 42, F1: 85}
    switch ("\n\t" + "a") {
    case "":
        var v8 float32 = 0.0
        _ = v8
    case "hello":
        var v9 int16 = 78
        _ = v9
    default:
        for fi10 := 0; fi10 < 16; fi10++ {
            switch "a" {
            case "hello":
                switch (fi10 + (fi10 ^ (fi10 - (fi10 >> ((482) & 63))))) {
                case -9223372036854775808:
                    fmt.Println(p0, fi10)
                }
            case "foo bar":
                for fi11 := 0; fi11 < 16; fi11++ {
                    switch "foo bar" {
                    case "a":
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    case "foo bar":
                        var v12 uint = 0
                        _ = v12
                    }
                    p0 = ((true || false) || p0)
                    if true {
                        fmt.Println(fi11, p0, fi10)
                        v13 := GenS1{F0: -316, F1: 30}
                        if p0 {
                            p0 = true
                            _ = len(v6)
                            _ = len(v6)
                        }
                        p0 = false
                        switch "foo bar" {
                        case "hello":
                            switch (fi10 % ((-789) | 1)) {
                            case 4294967295:
                                switch ("a" + (("a" + ("" + ("\n\t" + "hello"))) + "hello")) {
                                case "hello":
                                    p0 = (-2.25 < 1.5)
                                case "":
                                    if (127 != int8(fi10)) {
                                        fmt.Println(fi11, fi10)
                                    }
                                case "a":
                                    var v14 uint8 = 1
                                    _ = v14
                                }
                            }
                        default:
                            if (50 >= -128) {
                                var v15 uint = 18446744073709551615
                                v16 := GenS2{F0: -1, F1: int64(v15), F2: (v15 & (v15 << ((v15) & 63)))}
                                p0 = p0
                                _ = v15
                                _ = v16
                            } else {
                                var v17 uint8 = 2
                                _ = v17
                            }
                            var v18 uint8 = 53
                            _ = v18
                        }
                        _ = v13
                    }
                    _ = fi11
                }
            case "\n\t":
                p0 = ((!true) || (-2.25 > 0.1))
            }
            p0 = (2 <= 473)
            _ = fi10
        }
    }
    if false {
        v19 := GenS2{F0: -1, F1: -55, F2: 2}
        _ = v19
    }
    _ = p0
    _ = v6
    _ = v7
    return -773
}

func main() {
    var v1 uint32 = 1
    v2 := make(map[string]bool)
    v3 := []uint32{65535, (v1 - (v1 / (((v1 - 869)) | 1))), 42}
    v4 := make(map[int]int)
    v5 := make(map[uint32]uint32)
    var v20 int32 = -2147483648
    var v21 bool = (18446744073709551615 < 4294967295)
    v1 += 887
    if (2 < 255) {
        var v22 int64 = 46
        v23 := GenS2{F0: -1, F1: v22, F2: 0}
        _ = v22
        _ = v23
    }
    v24 := GenS3{F0: 94}
    v25 := GenS4{F0: -1}
    var v26 uint = 773
    var v27 int16 = genFn0(v21)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
