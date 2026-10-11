package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16; F1 bool }

func main() {
    var v1 int32 = 1
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]int)
    v4 := []uint32{2147483648}
    switch "foo bar" {
    case "":
        for fi5 := 0; fi5 < 13; fi5++ {
            switch "" {
            case "hello":
                if (!true) {
                    v1 += (^(v1))
                    if false {
                        switch fi5 {
                        case 4294967295:
                            v6 := GenS1{F0: 42, F1: 9223372036854775807}
                            _ = v6
                        case 0:
                            v3[2147483648] = (fi5 + -274)
                        case 9223372036854775807:
                            v1 = (v1 / ((-255) | 1))
                        }
                    } else {
                        _ = len(v2)
                        v7 := GenS1{F0: -441, F1: 281}
                        var v8 int = 174
                        _ = v7
                        _ = v8
                    }
                } else {
                    if len(v4) > 0 {
                        v4[0] = 42
                    }
                    switch "" {
                    case "":
                        switch (fi5 * (fi5 - -9223372036854775808)) {
                        case 477:
                            for fi9 := 0; fi9 < 6; fi9++ {
                                var v10 uint64 = 93
                                _ = fi9
                                _ = v10
                            }
                        default:
                            var v11 string = "foo bar"
                            _ = v11
                        }
                    case "a":
                        v1 -= v1
                    case "foo bar":
                        var v12 float32 = 0.0
                        _ = v12
                    default:
                        if true {
                            for fi13 := 0; fi13 < 13; fi13++ {
                                if (65535 < 0) {
                                    var v14 int8 = 42
                                    _ = v14
                                } else {
                                    fmt.Println(fi13, fi5, v1)
                                }
                                var v15 uint = 0
                                _ = fi13
                                _ = v15
                            }
                        } else {
                            if false {
                                var v16 int32 = 1
                                _ = v16
                            } else {
                                var v17 uint = 18446744073709551615
                                for fi18 := 0; fi18 < 20; fi18++ {
                                    fmt.Println(fi18, v1, fi18)
                                    v17 += 4294967295
                                    _ = fi18
                                }
                                _ = v17
                            }
                            var v19 uint8 = 1
                            _ = v19
                        }
                    }
                }
            default:
                var v20 uint64 = 1
                _ = v20
            }
            _ = fi5
        }
    case "foo bar":
        for fi21 := 0; fi21 < 20; fi21++ {
            v22 := GenS2{F0: 1}
            v1 -= 2147483647
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi21
            _ = v22
        }
    }
    v23 := GenS3{F0: 42, F1: ((false && true) && ("" > "foo bar"))}
    var v24 int = 32767
    var v25 uint = 709
    fmt.Println(v1, v25, v24)
    var v26 int16 = 32767
    var v27 uint32 = 629
    fmt.Println(v1, v27)
    v24 += (v24 << ((42) & 63))
    fmt.Println(v1, v25)
    var v28 int8 = 127
    switch 1 {
    case 2147483647:
        if true {
            v24 = (v24 >> ((v25) & 63))
            var v29 uint16 = 0
            _ = v29
        }
    case 397:
        var v30 int8 = v28
        _ = v30
    }
    for fi31 := 0; fi31 < 7; fi31++ {
        switch ("hello" + "") {
        case "\n\t":
            v32 := GenS2{F0: -77}
            _ = v32
        default:
            var v33 int32 = (v1 / ((0) | 1))
            _ = v33
        }
        v26 = v26
        _ = fi31
    }
    var v34 bool = (0 < 1)
    var v35 int16 = v26
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v34
    _ = v35
    _ = v4
}
