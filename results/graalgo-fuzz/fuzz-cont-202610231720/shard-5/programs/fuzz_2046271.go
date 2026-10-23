package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 uint8; F1 int8; F2 uint32 }
type GenS3 struct { F0 int64; F1 int }

func main() {
    var v1 bool = (1 < 209)
    v2 := make(map[uint32]int)
    var v3 int16 = 1000
    v4 := make(map[string]bool)
    switch ("" + ("a" + ("hello" + "foo bar"))) {
    case "a":
        v3 = (v3 - v3)
    default:
        switch 32767 {
        case 2147483647:
            switch "" {
            case "a":
                v5 := GenS1{F0: 1, F1: 32767}
                _ = v5
            default:
                switch ("hello" + "foo bar") {
                case "\n\t":
                    var v6 uint = 2
                    _ = v6
                default:
                    fmt.Println(v3)
                }
            }
        }
    }
    fmt.Println(v1, v1)
    v3 = (v3 % ((v3) | 1))
    if true {
        if (1 == 42) {
            v7 := GenS2{F0: 128, F1: -36, F2: 2147483648}
            switch ("" + "hello") {
            case "\n\t":
                fmt.Println(v1, v1)
            case "foo bar":
                var v8 int = 127
                _ = v8
            case "hello":
                fmt.Println(v3, v3)
            default:
                v9 := v4["\n\t"]
                _ = v9
            }
            var v10 int32 = 42
            _ = v10
            _ = v7
        } else {
            var v11 uint = 65535
            if (int32(v11) > -502) {
                if false {
                    for fi12 := 0; fi12 < 3; fi12++ {
                        var v13 uint8 = 42
                        v14 := GenS2{F0: v13, F1: -128, F2: 1}
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                }
                var v15 string = ("" + "")
                v1 = v1
                _ = v15
            }
            if (!v1) {
                v3 -= (v3 % ((int16(v11)) | 1))
                if (int8(v3) < 0) {
                    for fi16 := 0; fi16 < 3; fi16++ {
                        v4[("" + "hello")] = (!(!(0 != 9223372036854775807)))
                        v17 := []int{fi16, fi16, fi16}
                        _ = fi16
                        _ = v17
                    }
                    fmt.Println(v1)
                    v3 -= v3
                }
            } else {
                v11 -= 4294967295
            }
            _ = v11
        }
        v1 = v1
    } else {
        v1 = ((("foo bar" + "foo bar") <= ("a" + (("a" + "") + "a"))) || v1)
        fmt.Println(v1)
    }
    var v18 uint8 = 170
    if false {
        v3 += (v3 % ((v3) | 1))
        for fi19 := 0; fi19 < 2; fi19++ {
            v3 = v3
            delete(v2, 1)
            if (false && (195 < 18446744073709551615)) {
                var v20 uint64 = 774
                v20 += 2
                _ = v20
            }
            _ = fi19
        }
        for fi21 := 0; fi21 < 18; fi21++ {
            v22 := GenS1{F0: 203, F1: 805}
            v1 = (1 == 42)
            switch "" {
            case "":
                var v23 int = fi21
                _ = v23
            case "hello":
                for fi24 := 0; fi24 < 20; fi24++ {
                    if (107 <= 0) {
                        var v25 uint16 = 804
                        if (414 >= 37) {
                            var v26 float64 = ((0.0 + 1e10) + float64(-2.25))
                            v27 := GenS3{F0: 1, F1: (^(fi21))}
                            _ = v26
                            _ = v27
                        }
                        _ = v25
                    }
                    v3 -= int16(fi24)
                    _ = fi24
                }
            default:
                var v28 int32 = 1
                _ = v28
            }
            _ = fi21
            _ = v22
        }
    }
    if true {
        for fi29 := 0; fi29 < 1; fi29++ {
            var v30 uint64 = 18446744073709551615
            _ = fi29
            _ = v30
        }
        var v31 int = -133
        _ = v31
    }
    v32 := []int{475}
    v3 = -491
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v32
    _ = v4
}
