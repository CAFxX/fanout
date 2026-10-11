package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 int32; F2 int16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64 }
func genFn0(p0 int32) uint16 {
    if false {
        var v5 float32 = (1.5 / 1.0)
        _ = v5
    }
    v6 := GenS1{F0: 2}
    v7 := GenS1{F0: 4294967295}
    _ = p0
    _ = v6
    _ = v7
    return 2
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint = 255
    var v3 string = "\n\t"
    var v4 float64 = 0.1
    v8 := v1[65535]
    v9 := []int{v8, v8, v8, v8}
    v9 = append(v9, 219)
    var v10 int64 = 1
    switch ("foo bar" + (v3 + "a")) {
    case "a":
        var v11 uint32 = 42
        _ = v11
    case "":
        fmt.Println(v3, v8, v10)
    }
    fmt.Println(v3)
    for fi12 := 0; fi12 < 12; fi12++ {
        v13 := GenS2{F0: 0, F1: -1, F2: -15}
        if len(v9) > 0 {
            v9[0] = 128
        }
        var v14 float32 = 3.14159
        _ = fi12
        _ = v13
        _ = v14
    }
    switch (^(v8)) {
    case -117:
        switch v8 {
        case 32768:
            v3 = v3
        case 950:
            fmt.Println(v4, v3, v3)
        case 588:
            v3 = v3
        }
    case 42:
        delete(v1, 0)
    case 2147483647:
        v10 = 9223372036854775807
    }
    v15 := GenS1{F0: 2147483648}
    var v16 uint8 = 68
    fmt.Println(v2, v3, v8)
    switch v8 {
    case 768:
        for fi17 := 0; fi17 < 11; fi17++ {
            _ = len(v9)
            v2 = 4294967295
            _ = fi17
        }
    case 42:
        var v18 uint32 = 959
        _ = v18
    case 32767:
        fmt.Println(v3, v10, v3)
    default:
        for fi19 := 0; fi19 < 17; fi19++ {
            if false {
                v4 = v4
            } else {
                if (42 == int8(v10)) {
                    if true {
                        for fi20 := 0; fi20 < 11; fi20++ {
                            if false {
                                v3 = v3
                                var v21 int16 = 32767
                                switch (v8 | 970) {
                                case -1:
                                    v10 -= 868
                                default:
                                    fmt.Println(v2, v21, v3)
                                    if len(v9) >= 2 {
                                        v9 = v9[1:len(v9)]
                                    }
                                }
                                _ = v21
                            }
                            _ = fi20
                        }
                    } else {
                        var v22 int16 = 32767
                        v16 += (^(v16))
                        if len(v9) > 0 {
                            v9[0] = v8
                        }
                        _ = v22
                    }
                    for fi23 := 0; fi23 < 6; fi23++ {
                        switch ("\n\t" + (v3 + v3)) {
                        case "a":
                            switch fi23 {
                            case 128:
                                if false {
                                    v9 = append(v9, 2147483647)
                                    v3 = v3
                                    if true {
                                        v16 += (^(v16))
                                        _ = len(v9)
                                    }
                                }
                            case 0:
                                v24 := GenS1{F0: 750}
                                _ = v24
                            default:
                                if true {
                                    switch v3 {
                                    case "a":
                                        if len(v9) >= 2 {
                                            v9 = v9[1:len(v9)]
                                        }
                                    case "":
                                        v3 = ("" + v3)
                                        v8 += -270
                                        var v25 int8 = 100
                                        _ = v25
                                    case "\n\t":
                                        v26 := GenS2{F0: 42, F1: 2147483647, F2: 1000}
                                        _ = v26
                                    }
                                    var v27 uint8 = v16
                                    _ = v27
                                }
                            }
                        case "":
                            if true {
                                v10 = (v10 % (((v10 + (v10 << ((18446744073709551615) & 63)))) | 1))
                            }
                        case "hello":
                            var v28 uint32 = uint32(v2)
                            _ = v28
                        }
                        _ = fi23
                    }
                } else {
                    for fi29 := 0; fi29 < 17; fi29++ {
                        var v30 uint32 = 4294967295
                        v31 := v1[(v30 | (v30 / ((v30) | 1)))]
                        v32 := GenS3{F0: -1}
                        _ = fi29
                        _ = v30
                        _ = v31
                        _ = v32
                    }
                }
                var v33 uint = (^(v2))
                _ = v33
            }
            _ = fi19
        }
    }
    var v34 uint8 = (^(v16))
    switch "\n\t" {
    case "foo bar":
        for fi35 := 0; fi35 < 12; fi35++ {
            v3 = ""
            var v36 int16 = 32767
            v37 := GenS4{F0: 42}
            _ = fi35
            _ = v36
            _ = v37
        }
    case "\n\t":
        var v38 uint8 = (^(v16))
        _ = v38
    default:
        var v39 int8 = -128
        _ = v39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v8
    _ = v9
}
