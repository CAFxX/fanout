package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 uint8; F2 int32 }
type GenS3 struct { F0 uint32; F1 int8 }
type GenS4 struct { F0 int64; F1 uint64 }

func main() {
    var v1 float64 = 3.14159
    var v2 int32 = 42
    var v3 string = "a"
    v4 := []uint32{45, 712}
    var v5 uint16 = 149
    v6 := make(map[string]bool)
    if (42 < v2) {
        fmt.Println(v2)
        v7 := GenS1{F0: 4294967295}
        _ = v7
    }
    for fi8 := 0; fi8 < 4; fi8++ {
        _ = len(v4)
        fmt.Println(v2, fi8, v3)
        _ = fi8
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        var v10 uint64 = 1
        _ = fi9
        _ = v10
    }
    switch v3 {
    case "\n\t":
        v6["a"] = true
    case "a":
        if (((v2 << ((1) & 63)) != (v2 % ((v2) | 1))) || false) {
            if ((0.0 * float32(v1)) >= 3.14159) {
                for fi11 := 0; fi11 < 2; fi11++ {
                    v6[v3] = false
                    for fi12 := 0; fi12 < 2; fi12++ {
                        v13 := GenS1{F0: 2}
                        var v14 int8 = 100
                        switch fi12 {
                        case 128:
                            v5 = v5
                        case 1:
                            var v15 uint64 = 42
                            _ = v15
                        case -9223372036854775808:
                            var v16 int64 = -1
                            _ = v16
                        default:
                            if (4294967295 >= 1) {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                        }
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                    for fi17 := 0; fi17 < 20; fi17++ {
                        v3 = ((v3 + v3) + v3)
                        var v18 uint64 = 42
                        if ((32767 > 32767) || (-1 <= 649)) {
                            var v19 bool = ((v1 == 0.1) || false)
                            v2 += (v2 - (-(v2)))
                            _ = v19
                        }
                        _ = fi17
                        _ = v18
                    }
                    _ = fi11
                }
                switch "foo bar" {
                case "a":
                    v20 := GenS1{F0: 162}
                    _ = v20
                case "":
                    var v21 uint8 = 255
                    _ = v21
                }
            }
            var v22 uint64 = 588
            _ = v22
        } else {
            v2 = (-(v2))
            fmt.Println(v5, v1, v2)
        }
    case "hello":
        v23 := GenS2{F0: v3, F1: 157, F2: v2}
        _ = v23
    }
    if false {
        fmt.Println(v2, v2, v1)
    } else {
        if false {
            if (107 < (v2 / ((v2) | 1))) {
                switch -1 {
                case 296:
                    v1 += ((0.0 + 0.0) - float64(1.5))
                case 127:
                    if len(v4) > 0 {
                        v4[0] = 576
                    }
                default:
                    if (true || false) {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        if true {
                            v24 := GenS3{F0: 0, F1: 127}
                            fmt.Println(v5, v1, v2)
                            _ = v24
                        } else {
                            delete(v6, (v3 + v3))
                            v3 = v3
                        }
                    } else {
                        _ = len(v4)
                        switch "a" {
                        case "a":
                            v1 += v1
                        }
                    }
                }
                v25 := GenS1{F0: 819}
                _ = v25
            } else {
                v2 = 0
            }
            for fi26 := 0; fi26 < 12; fi26++ {
                var v27 uint64 = 1
                _ = fi26
                _ = v27
            }
        }
    }
    fmt.Println(v3)
    var v28 int64 = 836
    v29 := GenS4{F0: (v28 & v28), F1: 4294967295}
    var v30 uint8 = 104
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
