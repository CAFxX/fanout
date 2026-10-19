package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 uint64; F2 int8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int32; F1 string; F2 int32 }
type GenS5 struct { F0 uint32; F1 bool }

func main() {
    var v1 uint64 = 742
    var v2 uint64 = (v1 >> ((2) & 63))
    var v3 int16 = 1000
    v4 := make(map[uint32]bool)
    var v5 uint = 469
    v5 += (v5 / (((v5 ^ 18446744073709551615)) | 1))
    switch 32768 {
    case 866:
        switch 1 {
        case 128:
            fmt.Println(v1, v3)
        case 32767:
            if false {
                v2 += v2
                var v6 float64 = (0.1 / 1.5)
                v2 = (v2 << (((v5 | 255)) & 63))
                _ = v6
            } else {
                var v7 int64 = -93
                _ = v7
            }
        }
    case 444:
        v8 := []int{}
        _ = v8
    case 9223372036854775807:
        if false {
            v9 := GenS1{F0: ("foo bar" + ("hello" + (("hello" + "foo bar") + "a")))}
            v3 = (v3 << ((v5) & 63))
            _ = v9
        } else {
            v10 := v4[65535]
            v11 := GenS1{F0: "hello"}
            _ = v10
            _ = v11
        }
    default:
        var v12 uint16 = 255
        _ = v12
    }
    for fi13 := 0; fi13 < 13; fi13++ {
        if false {
            v14 := []int{fi13, (fi13 & (fi13 % ((fi13) | 1)))}
            if (false && (!false)) {
                var v15 string = "a"
                _ = v15
            }
            _ = len(v4)
            _ = v14
        }
        _ = fi13
    }
    v3 -= v3
    for fi16 := 0; fi16 < 13; fi16++ {
        for fi17 := 0; fi17 < 1; fi17++ {
            v3 = (^(v3))
            var v18 float64 = 1.5
            v5 += (v5 >> ((v5) & 63))
            _ = fi17
            _ = v18
        }
        _ = fi16
    }
    if ((v1 - v2) < 0) {
        v19 := []int{853, -1}
        switch "" {
        case "a":
            for fi20 := 0; fi20 < 5; fi20++ {
                v21 := GenS2{F0: "a", F1: (v2 - 4294967295), F2: 127}
                if (!(!(3.14159 >= 0.1))) {
                    switch ("\n\t" + "") {
                    case "":
                        var v22 int16 = (-(v3))
                        _ = v22
                    default:
                        v2 = (v1 >> (((v5 ^ v5)) & 63))
                    }
                    var v23 int16 = v3
                    fmt.Println(v23, v3, v23)
                    _ = v23
                }
                for fi24 := 0; fi24 < 8; fi24++ {
                    if false {
                        if true {
                            if (!true) {
                                v1 = v1
                            }
                            if true {
                                if len(v19) >= 2 {
                                    v19 = v19[1:len(v19)]
                                }
                            }
                            fmt.Println(v5, v1)
                        } else {
                            switch "\n\t" {
                            case "foo bar":
                                if true {
                                    if true {
                                        if (v5 > (^(v5))) {
                                            fmt.Println(v5, v1)
                                        }
                                        var v25 int16 = 139
                                        var v26 int8 = 127
                                        _ = v25
                                        _ = v26
                                    } else {
                                        var v27 int8 = 1
                                        switch "a" {
                                        case "":
                                            v27 += v27
                                        case "foo bar":
                                            if false {
                                                v28 := GenS3{F0: v1}
                                                var v29 float32 = (((3.14159 * 1e10) + ((float32(1.5) - 1.5) + (((float32(0.0) * ((((0.0 + 1.5) / (float32((1.5 - 3.14159)) + 0.0)) / (3.14159 + ((1e10 * 1.5) + 1e10))) - (1.5 / 1.5))) / -2.25) - 0.0))) - 3.14159)
                                                v1 -= v2
                                                _ = v28
                                                _ = v29
                                            }
                                        default:
                                            switch "\n\t" {
                                            case "\n\t":
                                                var v30 float32 = 1e10
                                                _ = v30
                                            case "a":
                                                if len(v19) >= 2 {
                                                    v19 = v19[1:len(v19)]
                                                }
                                            }
                                        }
                                        _ = v27
                                    }
                                    _ = len(v4)
                                    v4[823] = (!(!(0 >= -2147483648)))
                                    if (81 < -128) {
                                        var v31 int32 = 42
                                        _ = v31
                                    } else {
                                        v5 += (v5 | (v5 / ((v5) | 1)))
                                    }
                                    delete(v4, uint32(v1))
                                } else {
                                    if false {
                                        var v32 float32 = 1.5
                                        v33 := GenS4{F0: 851, F1: (("a" + (("hello" + "hello") + "\n\t")) + ("a" + "hello")), F2: 1}
                                        if true {
                                            v3 = (^(v3))
                                        }
                                        var v34 uint64 = v2
                                        v5 -= v5
                                        _ = v32
                                        _ = v33
                                        _ = v34
                                    }
                                    v4[2147483647] = (false && false)
                                    switch "hello" {
                                    case "hello":
                                        var v35 uint16 = uint16(v2)
                                        _ = v35
                                    }
                                }
                            case "":
                                v1 = v1
                            case "hello":
                                v5 -= (v5 % ((v5) | 1))
                            }
                        }
                        var v36 string = "foo bar"
                        v37 := GenS2{F0: v36, F1: (v1 + v2), F2: 9}
                        _ = v36
                        _ = v37
                    } else {
                        var v38 uint32 = 4294967295
                        v39 := GenS5{F0: (v38 + 0), F1: (("hello" + ("a" + "foo bar")) > "\n\t")}
                        _ = v38
                        _ = v39
                    }
                    v40 := GenS1{F0: "foo bar"}
                    _ = fi24
                    _ = v40
                }
                _ = fi20
                _ = v21
            }
        case "\n\t":
            v1 -= 18446744073709551615
        }
        fmt.Println(v1)
        _ = v19
    } else {
        for fi41 := 0; fi41 < 5; fi41++ {
            v2 = (v2 >> ((v5) & 63))
            _ = fi41
        }
        v4[1] = (false && (v5 < v5))
    }
    v42 := []int{127, 941}
    var v43 int = 730
    fmt.Println(v43, v2)
    if ((v3 - (v3 / (((-(v3))) | 1))) != (v3 % (((^(v3))) | 1))) {
        _ = len(v4)
        fmt.Println(v43, v3)
        v44 := v4[65535]
        _ = v44
    }
    var v45 int = 42
    v1 = v2
    var v46 string = ("\n\t" + ("" + ("a" + "foo bar")))
    delete(v4, 1)
    v1 -= (v2 & v1)
    v5 = 4294967295
    v4[2] = (v46 < v46)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v45
    _ = v46
    _ = v5
}
