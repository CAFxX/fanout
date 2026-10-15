package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 uint64 }
type GenS2 struct { F0 int8; F1 int; F2 int16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int64 }
func genFn0() int64 {
    for fi5 := 0; fi5 < 8; fi5++ {
        var v6 float64 = 0.1
        fmt.Println(fi5)
        var v7 int = (fi5 << ((42) & 63))
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 float32 = 1.5
    _ = v8
    return -971
}
func genFn1(p0 uint8) string {
    fmt.Println(p0)
    p0 = p0
    v9 := GenS1{F0: 338, F1: 255, F2: 845}
    _ = p0
    _ = v9
    return "a"
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int64 = 0
    var v3 int8 = 87
    v4 := []uint32{65535, 0}
    if (!(!false)) {
        if false {
            var v10 float64 = 3.14159
            v3 += -82
            _ = v10
        } else {
            v2 = (v2 ^ 9223372036854775807)
            for fi11 := 0; fi11 < 3; fi11++ {
                if (true && false) {
                    for fi12 := 0; fi12 < 6; fi12++ {
                        var v13 string = ("\n\t" + (("" + "a") + "foo bar"))
                        v3 += 100
                        var v14 int8 = (v3 * v3)
                        v15 := GenS2{F0: v3, F1: fi11, F2: 1000}
                        _ = fi12
                        _ = v13
                        _ = v14
                        _ = v15
                    }
                    _ = len(v4)
                }
                var v16 uint16 = 32767
                for fi17 := 0; fi17 < 11; fi17++ {
                    _ = len(v4)
                    switch (("foo bar" + ("\n\t" + ("foo bar" + "foo bar"))) + "foo bar") {
                    case "foo bar":
                        var v18 uint64 = 2
                        _ = v18
                    default:
                        var v19 bool = (921 == -96)
                        _ = v19
                    }
                    v16 = v16
                    _ = fi17
                }
                _ = fi11
                _ = v16
            }
        }
    }
    v1[2147483647] = 0
    var v20 int16 = 32767
    v2 += v2
    fmt.Println(v3, v3, v3)
    v20 = v20
    switch "\n\t" {
    case "":
        for fi21 := 0; fi21 < 17; fi21++ {
            if (!(true && (((0 >= -1) && ((true && (65535 < 42)) && (((4294967295 >= 0) && true) || false))) || (int8(v20) <= -128)))) {
                v4 = append(v4, 2147483648)
                switch fi21 {
                case -718:
                    if (!(666 < 146)) {
                        fmt.Println(v2, v20, v20)
                        if true {
                            v3 = (v3 + -128)
                            if true {
                                if (false || false) {
                                    fmt.Println(v2)
                                    v22 := GenS3{F0: 727}
                                    var v23 int8 = v3
                                    _ = v22
                                    _ = v23
                                } else {
                                    switch -1 {
                                    case 224:
                                        v20 = 0
                                    case 0:
                                        switch fi21 {
                                        case -369:
                                            switch (fi21 ^ fi21) {
                                            case 544:
                                                fmt.Println(v20, fi21)
                                            case 32767:
                                                v24 := GenS2{F0: v3, F1: (fi21 * 4294967295), F2: (v20 * 42)}
                                                _ = v24
                                            default:
                                                for fi25 := 0; fi25 < 13; fi25++ {
                                                    v2 -= (v2 % ((1) | 1))
                                                    v20 -= (v20 / ((v20) | 1))
                                                    _ = fi25
                                                }
                                            }
                                        case -2147483648:
                                            v2 -= (v2 << ((1) & 63))
                                        }
                                    case 367:
                                        for fi26 := 0; fi26 < 12; fi26++ {
                                            v3 = (-(v3))
                                            v27 := GenS4{F0: v2}
                                            switch -772 {
                                            case 127:
                                                var v28 uint64 = 263
                                                _ = v28
                                            case 457:
                                                fmt.Println(fi26, fi26, v20)
                                            case 605:
                                                fmt.Println(v3, fi21)
                                            }
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            _ = fi26
                                            _ = v27
                                        }
                                    }
                                }
                            }
                            fmt.Println(fi21, v2, v2)
                        }
                    }
                default:
                    var v29 int64 = v2
                    _ = v29
                }
                _ = len(v4)
            }
            _ = fi21
        }
    case "a":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case "hello":
        if true {
            var v30 float32 = 0.0
            _ = v30
        } else {
            var v31 uint32 = 1
            if ("" > (("a" + "\n\t") + (("foo bar" + "foo bar") + "hello"))) {
                fmt.Println(v2, v2, v20)
            }
            v32 := GenS3{F0: 4294967295}
            _ = v31
            _ = v32
        }
    }
    if (!false) {
        v20 += 42
    }
    fmt.Println(v3, v20, v20)
    v1[9223372036854775807] = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
