package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint16; F1 bool }
type GenS5 struct { F0 int16; F1 uint64 }

func main() {
    var v1 string = "foo bar"
    var v2 int = 4294967295
    v3 := []uint32{4294967295, 4294967295, 65535, 0}
    v4 := make(map[uint32]string)
    var v5 int = (v2 >> ((2) & 63))
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 int64 = -452
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v8 uint64 = 18446744073709551615
        _ = fi6
        _ = v7
        _ = v8
    }
    switch v1 {
    case "foo bar":
        for fi9 := 0; fi9 < 15; fi9++ {
            if (true && (961 == 146)) {
                v10 := GenS1{F0: 0}
                var v11 int8 = 0
                _ = v10
                _ = v11
            } else {
                if (-1 >= -128) {
                    _ = len(v3)
                    var v12 float32 = 3.14159
                    if false {
                        for fi13 := 0; fi13 < 18; fi13++ {
                            var v14 int8 = 127
                            _ = fi13
                            _ = v14
                        }
                    } else {
                        if false {
                            var v15 uint = 255
                            _ = v15
                        }
                        v16 := GenS1{F0: 247}
                        _ = v16
                    }
                    _ = v12
                }
                var v17 uint8 = 225
                _ = v17
            }
            _ = fi9
        }
    default:
        v18 := GenS1{F0: 127}
        _ = v18
    }
    v19 := GenS1{F0: 129}
    for fi20 := 0; fi20 < 20; fi20++ {
        switch (v2 << ((4294967295) & 63)) {
        case -9223372036854775808:
            for fi21 := 0; fi21 < 10; fi21++ {
                if (52 <= 2147483647) {
                    switch fi20 {
                    case 32767:
                        if true {
                            var v22 int8 = -83
                            if (!(4294967295 < 898)) {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                            _ = v22
                        } else {
                            v1 = v1
                        }
                    case 9223372036854775807:
                        if false {
                            v23 := GenS2{F0: ("a" + (v1 + ((v1 + "") + ("a" + v1)))), F1: (^(fi21))}
                            var v24 bool = false
                            var v25 uint16 = 255
                            _ = v23
                            _ = v24
                            _ = v25
                        } else {
                            if ((42 >= 191) && (255 > 316)) {
                                if ((v1 + (v1 + "foo bar")) >= v1) {
                                    v26 := GenS3{F0: -128}
                                    v27 := GenS4{F0: 42, F1: (!true)}
                                    _ = v26
                                    _ = v27
                                }
                            }
                        }
                    default:
                        if false {
                            v28 := GenS2{F0: v1, F1: (^(v5))}
                            var v29 int32 = -186
                            _ = v28
                            _ = v29
                        }
                    }
                    _ = len(v3)
                }
                _ = fi21
            }
        case 2147483647:
            _ = len(v4)
        }
        v30 := GenS3{F0: -128}
        _ = fi20
        _ = v30
    }
    var v31 float32 = (1.5 / 1.0)
    v5 += (v5 ^ v2)
    var v32 float32 = float32(1.5)
    for fi33 := 0; fi33 < 5; fi33++ {
        v1 = "\n\t"
        _ = fi33
    }
    fmt.Println(v2, v32)
    fmt.Println(v32, v2)
    v1 = "\n\t"
    var v34 int32 = 0
    var v35 uint = 53
    if ((((-2.25 / ((3.14159 - v31) + v32)) > (v32 - -2.25)) && false) && (9223372036854775807 >= -1)) {
        v36 := GenS5{F0: -32768, F1: 834}
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
