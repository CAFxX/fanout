package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 uint16 }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 uint8; F1 uint8; F2 string }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 string }
func genFn0(p0 int8, p1 int16, p2 int) int8 {
    v5 := GenS1{F0: p1, F1: 2147483647, F2: 0}
    v6 := []int{}
    v7 := GenS2{F0: -51, F1: 72}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return p0
}

func main() {
    var v1 float32 = 1e10
    var v2 uint = 255
    v3 := []uint32{}
    v4 := make(map[uint32]string)
    v3 = append(v3, 424)
    switch "foo bar" {
    case "a":
        var v8 uint = v2
        _ = v8
    case "foo bar":
        if false {
            if ((0 == 32767) || (42 >= -2147483648)) {
                v2 += (^(v2))
            } else {
                switch "\n\t" {
                case "foo bar":
                    v4[927] = (("a" + ("hello" + "foo bar")) + (("hello" + "hello") + "\n\t"))
                case "a":
                    if false {
                        v2 -= 1
                    } else {
                        for fi9 := 0; fi9 < 3; fi9++ {
                            v2 += (v2 + v2)
                            _ = fi9
                        }
                        for fi10 := 0; fi10 < 14; fi10++ {
                            var v11 uint8 = 0
                            switch (^(fi10)) {
                            case -983:
                                v12 := GenS3{F0: 42, F1: 235, F2: ("hello" + "\n\t")}
                                _ = v12
                            case -1:
                                _ = len(v4)
                            case 984:
                                if true {
                                    v13 := GenS2{F0: 0, F1: v2}
                                    var v14 int16 = 0
                                    v11 += v11
                                    _ = v13
                                    _ = v14
                                } else {
                                    var v15 int = fi10
                                    _ = v15
                                }
                            default:
                                v2 += 335
                            }
                            _ = fi10
                            _ = v11
                        }
                        v1 += (float32((1e10 + (float64(v1) + 1e10))) + v1)
                    }
                }
            }
            if (true || (255 > 32767)) {
                var v16 uint32 = 4294967295
                _ = v16
            } else {
                for fi17 := 0; fi17 < 19; fi17++ {
                    fmt.Println(v1)
                    switch 4294967295 {
                    case 2147483647:
                        v18 := GenS3{F0: 2, F1: 255, F2: (("a" + "hello") + (("" + "hello") + ("\n\t" + ("a" + ""))))}
                        _ = v18
                    case 4:
                        v2 = 255
                    }
                    v1 += ((((v1 - (v1 + 0.1)) - float32(1.5)) - (((v1 * (1e10 - float32(1e10))) + (0.1 / float32(1e10))) - (v1 / (float32(3.14159) / v1)))) / v1)
                    _ = fi17
                }
                delete(v4, 4)
            }
            v19 := GenS2{F0: 42, F1: v2}
            _ = v19
        }
    default:
        var v20 uint32 = 386
        _ = v20
    }
    for fi21 := 0; fi21 < 20; fi21++ {
        fmt.Println(v1, v2, v1)
        if (127 > 127) {
            var v22 int64 = 382
            _ = v22
        } else {
            if true {
                var v23 uint16 = 65535
                fmt.Println(fi21, v23, fi21)
                for fi24 := 0; fi24 < 7; fi24++ {
                    if ((^(v2)) == (^(v2))) {
                        v1 += 1.5
                    }
                    var v25 float32 = (0.0 - v1)
                    _ = fi24
                    _ = v25
                }
                _ = v23
            } else {
                var v26 int = -278
                _ = v26
            }
            v2 = v2
        }
        var v27 uint = (v2 % ((245) | 1))
        _ = fi21
        _ = v27
    }
    v28 := GenS4{F0: -693}
    for fi29 := 0; fi29 < 19; fi29++ {
        v1 -= v1
        _ = fi29
    }
    if len(v3) > 0 {
        v3[0] = 2
    }
    var v30 int8 = genFn0(genFn0(0, -899, 128), 42, 9223372036854775807)
    v2 -= (v2 + 65535)
    switch ("\n\t" + "") {
    case "a":
        if (!(true && false)) {
            v30 = int8(v2)
        }
    case "":
        for fi31 := 0; fi31 < 15; fi31++ {
            if len(v3) > 0 {
                v3[0] = 2147483647
            }
            for fi32 := 0; fi32 < 13; fi32++ {
                v33 := GenS5{F0: ("foo bar" + ("hello" + "hello"))}
                var v34 int32 = 2147483647
                var v35 string = "foo bar"
                _ = fi32
                _ = v33
                _ = v34
                _ = v35
            }
            v1 -= (1e10 * v1)
            _ = fi31
        }
    }
    var v36 int8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v36
    _ = v4
}
