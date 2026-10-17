package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 string; F1 string; F2 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int16; F1 bool }

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = -128
    var v3 uint64 = 2
    v4 := []uint32{2147483648, 185}
    if (false && (65535 < 2147483647)) {
        if (!false) {
            v2 += 42
            var v5 uint64 = (v3 - v3)
            for fi6 := 0; fi6 < 6; fi6++ {
                fmt.Println(fi6, v2)
                var v7 bool = true
                if false {
                    switch (-(fi6)) {
                    case 4294967295:
                        if v7 {
                            fmt.Println(v3, v2)
                            if true {
                                _ = len(v4)
                            } else {
                                switch 4294967295 {
                                case 0:
                                    v8 := GenS1{F0: "", F1: 42}
                                    _ = v8
                                case 4294967295:
                                    fmt.Println(fi6, v7)
                                default:
                                    var v9 float32 = 1e10
                                    _ = v9
                                }
                                v2 = v2
                            }
                            _ = len(v4)
                        }
                    default:
                        v3 -= v5
                    }
                } else {
                    _ = len(v1)
                }
                _ = fi6
                _ = v7
            }
            _ = v5
        }
        for fi10 := 0; fi10 < 9; fi10++ {
            v11 := GenS1{F0: "", F1: 4294967295}
            for fi12 := 0; fi12 < 13; fi12++ {
                if true {
                    var v13 uint8 = 42
                    v2 = -83
                    if ((42 == 596) || (false || false)) {
                        switch fi12 {
                        case 128:
                            var v14 string = "hello"
                            _ = v14
                        case 1:
                            var v15 uint64 = v3
                            _ = v15
                        }
                        v2 = 44
                    }
                    _ = v13
                }
                var v16 int = fi10
                var v17 int64 = 0
                _ = len(v1)
                _ = fi12
                _ = v16
                _ = v17
            }
            _ = fi10
            _ = v11
        }
    }
    switch "foo bar" {
    case "foo bar":
        fmt.Println(v2)
    default:
        v3 = (v3 >> ((42) & 63))
    }
    for fi18 := 0; fi18 < 1; fi18++ {
        if (!true) {
            if false {
                for fi19 := 0; fi19 < 5; fi19++ {
                    var v20 uint8 = 2
                    _ = fi19
                    _ = v20
                }
            }
            if (("a" + "hello") == "") {
                v21 := GenS2{F0: "hello", F1: "", F2: (v2 | (^(v2)))}
                var v22 int8 = (v2 << ((255) & 63))
                _ = v21
                _ = v22
            } else {
                fmt.Println(v2, v2, v3)
                if (0.0 < 0.1) {
                    v2 -= v2
                } else {
                    var v23 int16 = -32768
                    var v24 uint = 255
                    _ = v23
                    _ = v24
                }
                var v25 float32 = 3.14159
                _ = v25
            }
        }
        var v26 uint32 = 700
        _ = fi18
        _ = v26
    }
    if false {
        v2 += -25
    }
    fmt.Println(v3, v3)
    v2 = v2
    v27 := GenS3{F0: 42}
    var v28 uint32 = 910
    for fi29 := 0; fi29 < 2; fi29++ {
        v30 := GenS4{F0: 771, F1: (!(4294967295 != (v3 & (v3 / ((v3) | 1)))))}
        switch ("" + "\n\t") {
        case "":
            var v31 uint16 = 0
            _ = v31
        default:
            var v32 uint32 = v28
            _ = v32
        }
        _ = fi29
        _ = v30
    }
    var v33 int32 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v33
    _ = v4
}
