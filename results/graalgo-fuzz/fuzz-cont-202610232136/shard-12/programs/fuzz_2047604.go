package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 uint32; F1 uint64; F2 int64 }
type GenS4 struct { F0 int16; F1 int; F2 uint }
func genFn0(p0 int64, p1 int32) uint32 {
    p0 = p0
    switch ("foo bar" + (("hello" + ((("foo bar" + "hello") + "\n\t") + "foo bar")) + "hello")) {
    case "a":
        fmt.Println(p0, p0)
    case "hello":
        for fi4 := 0; fi4 < 2; fi4++ {
            v5 := GenS1{F0: fi4}
            v6 := []int{4294967295}
            for fi7 := 0; fi7 < 16; fi7++ {
                if len(v6) > 0 {
                    v6[0] = fi7
                }
                _ = fi7
            }
            _ = fi4
            _ = v5
            _ = v6
        }
    }
    _ = p0
    _ = p1
    return 2147483648
}

func main() {
    var v1 uint32 = 2147483647
    v2 := []uint32{}
    v3 := make(map[string]string)
    if true {
        v1 += v1
        if (false && (!true)) {
            v1 += (v1 << ((1) & 63))
            var v8 uint16 = 42
            v1 += (v1 | v1)
            _ = v8
        } else {
            switch "" {
            case "a":
                if (("" > "") && (!false)) {
                    v9 := GenS1{F0: 42}
                    _ = v9
                }
            case "":
                for fi10 := 0; fi10 < 16; fi10++ {
                    switch fi10 {
                    case 32768:
                        var v11 int32 = -859
                        _ = v11
                    case -9223372036854775808:
                        for fi12 := 0; fi12 < 20; fi12++ {
                            if false {
                                v1 = (v1 & (v1 / (((v1 - 2147483648)) | 1)))
                            } else {
                                fmt.Println(fi10)
                            }
                            v1 += v1
                            _ = fi12
                        }
                    default:
                        v1 += (v1 * 65535)
                    }
                    var v13 uint32 = (v1 + (v1 | v1))
                    _ = len(v2)
                    _ = fi10
                    _ = v13
                }
            }
            if (true || (!((!true) || (true && true)))) {
                switch "" {
                case "\n\t":
                    if (!(0 > 84)) {
                        for fi14 := 0; fi14 < 6; fi14++ {
                            var v15 string = (("hello" + ("a" + ("foo bar" + "hello"))) + "")
                            v16 := GenS2{F0: 127, F1: 42}
                            fmt.Println(fi14)
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                        var v17 string = "a"
                        _ = v17
                    } else {
                        v2 = append(v2, 0)
                        for fi18 := 0; fi18 < 5; fi18++ {
                            for fi19 := 0; fi19 < 2; fi19++ {
                                if ((425 > 0) || (!(2147483647 > 461))) {
                                    if true {
                                        if len(v2) > 0 {
                                            v2[0] = v1
                                        }
                                        var v20 int8 = 16
                                        _ = v20
                                    } else {
                                        var v21 int8 = 79
                                        _ = v21
                                    }
                                    fmt.Println(v1, v1)
                                    v1 -= 42
                                }
                                if true {
                                    v1 -= v1
                                    var v22 int = (fi19 + (fi19 * 128))
                                    _ = v22
                                }
                                _ = fi19
                            }
                            v2 = append(v2, genFn0(9223372036854775807, int32(v1)))
                            fmt.Println(fi18)
                            _ = fi18
                        }
                    }
                case "foo bar":
                    fmt.Println(v1)
                case "hello":
                    if len(v2) > 0 {
                        v2[0] = 772
                    }
                }
                var v23 float32 = 1.5
                _ = v23
            }
        }
        v24 := GenS3{F0: 1, F1: 20, F2: -661}
        _ = v24
    }
    v1 = (v1 * (^(v1)))
    if len(v2) > 0 {
        v2[0] = genFn0(-824, -981)
    }
    v1 = v1
    var v25 uint8 = 239
    v25 += v25
    for fi26 := 0; fi26 < 11; fi26++ {
        var v27 uint8 = (v25 | (v25 * 255))
        if true {
            var v28 string = "hello"
            _ = v28
        } else {
            var v29 uint64 = 42
            v2 = append(v2, v1)
            if (2 < 470) {
                var v30 uint16 = 42
                _ = v30
            } else {
                v3[""] = ("foo bar" + "foo bar")
                var v31 float64 = 1.5
                _ = v31
            }
            _ = v29
        }
        if len(v2) > 0 {
            v2[0] = (v1 & 2147483647)
        }
        _ = fi26
        _ = v27
    }
    v32 := GenS4{F0: int16(v1), F1: -1, F2: uint(v1)}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v32
}
