package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 bool; F2 int32 }
type GenS3 struct { F0 int8; F1 uint; F2 uint16 }

func main() {
    var v1 float32 = float32(1e10)
    var v2 uint32 = 174
    var v3 int32 = 829
    var v4 uint16 = 32767
    if false {
        v1 -= (v1 * 0.1)
        if (false || (((0 < -128) || true) && (!(("a" != "foo bar") && true)))) {
            switch ("a" + (("" + "hello") + "foo bar")) {
            case "a":
                v2 -= v2
            case "":
                if false {
                    var v5 string = "hello"
                    v5 = v5
                    var v6 int = 0
                    _ = v5
                    _ = v6
                } else {
                    v3 += (v3 & (-(v3)))
                }
            case "hello":
                if false {
                    fmt.Println(v4)
                    if true {
                        v7 := []int{804}
                        for fi8 := 0; fi8 < 6; fi8++ {
                            if false {
                                v9 := GenS1{F0: 128}
                                switch "hello" {
                                case "a":
                                    fmt.Println(v2, v2, fi8)
                                }
                                v2 = v2
                                _ = v9
                            }
                            var v10 int64 = -9223372036854775808
                            for fi11 := 0; fi11 < 17; fi11++ {
                                fmt.Println(v1)
                                var v12 uint8 = 127
                                _ = fi11
                                _ = v12
                            }
                            _ = fi8
                            _ = v10
                        }
                        v13 := GenS2{F0: ("hello" + "foo bar"), F1: (128 < -2147483648), F2: 2147483647}
                        _ = v13
                        _ = v7
                    } else {
                        fmt.Println(v3, v2, v3)
                    }
                    v14 := GenS3{F0: int8(v4), F1: 1, F2: (^(v4))}
                    _ = v14
                } else {
                    switch int(v3) {
                    case 2147483647:
                        var v15 bool = false
                        _ = v15
                    }
                    var v16 float32 = v1
                    if (true || false) {
                        fmt.Println(v4, v3, v16)
                    } else {
                        for fi17 := 0; fi17 < 15; fi17++ {
                            var v18 int16 = 308
                            _ = fi17
                            _ = v18
                        }
                        var v19 int64 = 766
                        _ = v19
                    }
                    _ = v16
                }
            }
        } else {
            switch -9223372036854775808 {
            case -245:
                var v20 bool = false
                _ = v20
            case -485:
                v21 := []int{2147483647}
                _ = v21
            default:
                v2 += uint32(v4)
            }
        }
        if true {
            var v22 float32 = float32(3.14159)
            v23 := GenS2{F0: "\n\t", F1: false, F2: v3}
            _ = v22
            _ = v23
        }
    }
    v24 := GenS2{F0: "foo bar", F1: false, F2: v3}
    v3 = v3
    var v25 uint = 2
    var v26 string = ("a" + "foo bar")
    if (float64(float32(0.1)) == (3.14159 - -2.25)) {
        v27 := []int{471, -1}
        _ = v27
    }
    v28 := []int{}
    v26 = (((v26 + v26) + ("" + "hello")) + ("hello" + ("" + v26)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v28
    _ = v3
    _ = v4
}
