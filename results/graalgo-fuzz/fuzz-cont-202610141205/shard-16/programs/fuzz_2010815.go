package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint8 }

func main() {
    v1 := make(map[int]bool)
    v2 := []int{-9223372036854775808, 128, 2147483647}
    var v3 float64 = ((0.0 + ((3.14159 - float64((0.1 + 0.1))) / float64(3.14159))) + ((float64(0.1) + 0.0) + 3.14159))
    v4 := make(map[string]string)
    v5 := []uint32{4294967295, 2147483648}
    v6 := []int{4294967295}
    fmt.Println(v3)
    var v7 int8 = 0
    switch ("" + "hello") {
    case "":
        if ((false || (-943 == 2147483647)) && (!(false && (true || true)))) {
            var v8 uint16 = 758
            v9 := GenS1{F0: 1, F1: 5}
            for fi10 := 0; fi10 < 10; fi10++ {
                for fi11 := 0; fi11 < 8; fi11++ {
                    v3 = (v3 / v3)
                    v3 = (1.5 - (v3 + -2.25))
                    _ = fi11
                }
                var v12 int8 = v7
                if true {
                    for fi13 := 0; fi13 < 10; fi13++ {
                        v8 += v8
                        v3 += 1.5
                        var v14 uint32 = 2147483647
                        _ = fi13
                        _ = v14
                    }
                    v15 := GenS2{F0: 1, F1: 42}
                    _ = v15
                }
                _ = fi10
                _ = v12
            }
            _ = v8
            _ = v9
        } else {
            if ((v7 / ((v7) | 1)) <= v7) {
                for fi16 := 0; fi16 < 16; fi16++ {
                    switch fi16 {
                    case -9223372036854775808:
                        var v17 uint = 4294967295
                        _ = v17
                    case 32767:
                        v3 = ((1.5 + v3) * 1.5)
                    default:
                        for fi18 := 0; fi18 < 11; fi18++ {
                            switch fi16 {
                            case 127:
                                var v19 int32 = 0
                                _ = v19
                            case 128:
                                v7 = 1
                            default:
                                if false {
                                    v7 -= (^(v7))
                                    v7 = v7
                                    v3 = ((-2.25 / (float64(1.5) + v3)) + 0.0)
                                }
                            }
                            var v20 uint8 = 255
                            _ = fi18
                            _ = v20
                        }
                    }
                    var v21 uint16 = 1
                    _ = fi16
                    _ = v21
                }
            }
            var v22 bool = ((v7 ^ v7) > v7)
            v22 = false
            _ = v22
        }
    }
    for fi23 := 0; fi23 < 12; fi23++ {
        for fi24 := 0; fi24 < 2; fi24++ {
            if (32767 > 255) {
                v25 := GenS3{F0: 1}
                switch "" {
                case "a":
                    if (true && true) {
                        fmt.Println(v7)
                        fmt.Println(fi24, v7)
                    }
                case "":
                    fmt.Println(v3, fi23, fi24)
                }
                _ = len(v6)
                _ = v25
            }
            v3 += 1e10
            _ = fi24
        }
        _ = fi23
    }
    for fi26 := 0; fi26 < 12; fi26++ {
        fmt.Println(v7)
        for fi27 := 0; fi27 < 6; fi27++ {
            if false {
                v7 -= (^(v7))
                v28 := GenS4{F0: "\n\t"}
                _ = v28
            } else {
                v29 := GenS5{F0: 114}
                if len(v6) > 0 {
                    v6[0] = (fi27 / (((-(fi26))) | 1))
                }
                _ = v29
            }
            switch ("\n\t" + "hello") {
            case "":
                v6 = append(v6, -56)
            case "hello":
                v4["a"] = "hello"
            }
            switch (-(fi26)) {
            case -985:
                v7 -= v7
                var v30 int = fi26
                fmt.Println(v30)
                _ = v30
            default:
                if false {
                    v7 += (v7 << ((1) & 63))
                    _ = len(v1)
                    var v31 uint16 = 0
                    _ = v31
                }
            }
            _ = fi27
        }
        _ = fi26
    }
    fmt.Println(v3)
    if ((!(false && ((true || true) || ("hello" != ("" + ("a" + "\n\t")))))) || false) {
        switch -2147483648 {
        case 128:
            var v32 uint = 255
            _ = v32
        case 42:
            if len(v5) > 0 {
                v5[0] = 283
            }
        }
        v33 := GenS1{F0: -1, F1: 2}
        _ = v33
    }
    v3 += (v3 + 1.5)
    switch 9223372036854775807 {
    case 2147483647:
        v3 = (v3 / (float64(0.1) + (v3 + 3.14159)))
    default:
        var v34 int16 = 244
        _ = v34
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
