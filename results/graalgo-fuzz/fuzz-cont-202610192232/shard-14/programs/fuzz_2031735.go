package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 uint32; F1 int8; F2 int8 }
type GenS3 struct { F0 uint8; F1 uint8 }
func genFn0(p0 int) uint8 {
    v4 := GenS1{F0: "foo bar", F1: ((((("hello" + "a") + ("foo bar" + "\n\t")) + ("a" + "foo bar")) >= ("\n\t" + ("" + ("" + "foo bar")))) && true)}
    var v5 int8 = 1
    v6 := []int{}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if (p0 < 176) {
        var v7 int16 = -32768
        var v8 int64 = 1
        v7 = (-(v7))
        _ = v7
        _ = v8
    } else {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        if (p0 > (p0 - 2147483647)) {
            for fi9 := 0; fi9 < 16; fi9++ {
                var v10 uint8 = 255
                fmt.Println(p0)
                switch "hello" {
                case "":
                    if len(v6) > 0 {
                        v6[0] = (p0 / ((fi9) | 1))
                    }
                case "a":
                    if (162 != 32767) {
                        v6 = append(v6, 417)
                        var v11 int = fi9
                        _ = v11
                    }
                case "\n\t":
                    p0 = (-(fi9))
                }
                _ = fi9
                _ = v10
            }
        }
    }
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 128
}

func main() {
    v1 := make(map[int]string)
    v2 := []uint32{158, 715, 741}
    v3 := make(map[uint32]bool)
    var v12 uint64 = 828
    if false {
        var v13 uint64 = v12
        _ = v13
    } else {
        switch "" {
        case "foo bar":
            var v14 int64 = 0
            _ = v14
        case "hello":
            v12 = 1
        case "\n\t":
            v12 = v12
        }
    }
    for fi15 := 0; fi15 < 13; fi15++ {
        for fi16 := 0; fi16 < 7; fi16++ {
            _ = len(v2)
            v17 := GenS1{F0: "", F1: false}
            _ = fi16
            _ = v17
        }
        _ = fi15
    }
    switch ("hello" + "a") {
    case "\n\t":
        var v18 float32 = 1e10
        _ = v18
    case "hello":
        for fi19 := 0; fi19 < 8; fi19++ {
            var v20 uint64 = (v12 - (v12 ^ 198))
            _ = fi19
            _ = v20
        }
    default:
        v12 = (v12 << ((255) & 63))
    }
    fmt.Println(v12)
    v12 = (^(v12))
    v21 := GenS1{F0: "hello", F1: true}
    for fi22 := 0; fi22 < 11; fi22++ {
        if (!false) {
            fmt.Println(fi22, v12)
            if (!(62 <= 128)) {
                v2 = append(v2, 4294967295)
                v12 = v12
            }
            v23 := GenS2{F0: 417, F1: -128, F2: 46}
            _ = v23
        } else {
            switch fi22 {
            case 1:
                var v24 float64 = 1.5
                _ = v24
            case -138:
                for fi25 := 0; fi25 < 4; fi25++ {
                    var v26 uint16 = 42
                    var v27 bool = false
                    _ = fi25
                    _ = v26
                    _ = v27
                }
            }
        }
        _ = fi22
    }
    v28 := GenS3{F0: 120, F1: 127}
    if false {
        switch "\n\t" {
        case "\n\t":
            fmt.Println(v12)
        default:
            if true {
                v12 -= 42
                if ((42 < 1) && true) {
                    for fi29 := 0; fi29 < 11; fi29++ {
                        switch (fi29 ^ (fi29 << ((318) & 63))) {
                        case 42:
                            if false {
                                var v30 int8 = 42
                                var v31 uint16 = 42
                                var v32 uint = 4294967295
                                _ = v30
                                _ = v31
                                _ = v32
                            }
                        case -9223372036854775808:
                            v33 := GenS3{F0: 11, F1: 158}
                            _ = v33
                        }
                        _ = fi29
                    }
                    for fi34 := 0; fi34 < 18; fi34++ {
                        v2 = append(v2, 523)
                        _ = fi34
                    }
                    v12 -= (v12 + (v12 * 729))
                } else {
                    v12 = 695
                }
            }
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v35 := GenS2{F0: 42, F1: 7, F2: 79}
        _ = v35
    }
    v12 = 1
    if (!(4294967295 <= (v12 / (((v12 & 2)) | 1)))) {
        v2 = append(v2, 4294967295)
    } else {
        v12 = (v12 >> ((154) & 63))
        v12 += v12
    }
    for fi36 := 0; fi36 < 17; fi36++ {
        v12 += 42
        _ = fi36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v2
    _ = v21
    _ = v28
    _ = v3
}
