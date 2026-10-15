package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 string }

func main() {
    var v1 uint64 = 2
    v2 := []string{"a", ("foo bar" + "foo bar")}
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]string)
    v4[(("" + "\n\t") + "")] = "foo bar"
    if true {
        if true {
            var v5 uint8 = 220
            v6 := GenS1{F0: "hello"}
            _ = v5
            _ = v6
        }
        v1 += (v1 % ((v1) | 1))
    } else {
        for fi7 := 0; fi7 < 14; fi7++ {
            if (1 > 104) {
                v1 -= 1
                v1 -= v1
            } else {
                switch ("" + ("\n\t" + "\n\t")) {
                case "\n\t":
                    v1 -= (v1 >> ((0) & 63))
                case "a":
                    var v8 uint = 4294967295
                    _ = v8
                case "foo bar":
                    if (-32768 != 0) {
                        v1 = 4294967295
                        v9 := GenS2{F0: 32767, F1: -1}
                        _ = v9
                    } else {
                        v10 := GenS1{F0: ("foo bar" + "a")}
                        if true {
                            var v11 uint64 = 1
                            _ = v11
                        } else {
                            switch "a" {
                            case "hello":
                                v1 = 18446744073709551615
                            case "foo bar":
                                if false {
                                    if false {
                                        v1 = 4294967295
                                    } else {
                                        v12 := GenS1{F0: "a"}
                                        if (!(2 == 2147483648)) {
                                            v13 := GenS2{F0: 1, F1: 839}
                                            delete(v3, 41)
                                            fmt.Println(fi7, v1)
                                            _ = v13
                                        } else {
                                            var v14 uint16 = 44
                                            _ = v14
                                        }
                                        if (-1 < -751) {
                                            var v15 string = "hello"
                                            var v16 uint8 = 0
                                            _ = v15
                                            _ = v16
                                        } else {
                                            for fi17 := 0; fi17 < 5; fi17++ {
                                                var v18 float32 = -2.25
                                                switch "" {
                                                case "\n\t":
                                                    var v19 int16 = -373
                                                    _ = v19
                                                case "":
                                                    v20 := GenS1{F0: ""}
                                                    _ = v20
                                                case "a":
                                                    v21 := GenS1{F0: ""}
                                                    _ = v21
                                                default:
                                                    var v22 int32 = int32(v1)
                                                    _ = v22
                                                }
                                                var v23 float64 = 1e10
                                                _ = fi17
                                                _ = v18
                                                _ = v23
                                            }
                                        }
                                        _ = v12
                                    }
                                    v24 := GenS2{F0: 255, F1: 1000}
                                    _ = v24
                                } else {
                                    var v25 int32 = 0
                                    v26 := GenS2{F0: 2, F1: 42}
                                    for fi27 := 0; fi27 < 16; fi27++ {
                                        v25 = v25
                                        v28 := GenS3{F0: 72}
                                        _ = fi27
                                        _ = v28
                                    }
                                    _ = v25
                                    _ = v26
                                }
                            }
                        }
                        _ = v10
                    }
                }
                for fi29 := 0; fi29 < 8; fi29++ {
                    var v30 bool = true
                    _ = fi29
                    _ = v30
                }
            }
            _ = fi7
        }
    }
    for fi31 := 0; fi31 < 17; fi31++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi31
    }
    var v32 int8 = 53
    var v33 string = "\n\t"
    delete(v3, 725)
    switch -1 {
    case -983:
        fmt.Println(v1, v33, v1)
    }
    var v34 uint32 = 4294967295
    v34 = v34
    for fi35 := 0; fi35 < 12; fi35++ {
        var v36 int32 = -1
        v1 = (v1 & (v1 / ((v1) | 1)))
        _ = fi35
        _ = v36
    }
    fmt.Println(v33)
    switch 714 {
    case -9223372036854775808:
        _ = len(v2)
    case 32768:
        v1 -= (^(v1))
    default:
        v37 := GenS4{F0: ("foo bar" + "foo bar")}
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
}
