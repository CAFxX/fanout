package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 bool; F1 uint16; F2 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int8; F1 int8; F2 int32 }
type GenS5 struct { F0 int16; F1 bool; F2 bool }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint8 = 2
    v3 := make(map[int]string)
    v4 := make(map[uint32]uint32)
    v5 := []int{-2147483648, 9223372036854775807}
    v6 := []string{((("\n\t" + "\n\t") + "foo bar") + "hello"), "\n\t"}
    v7 := GenS1{F0: 9223372036854775807, F1: 65535}
    _ = len(v3)
    switch ("\n\t" + ("hello" + ("a" + ""))) {
    case "a":
        if false {
            fmt.Println(v2)
            var v8 uint = 2
            _ = v8
        }
    default:
        for fi9 := 0; fi9 < 2; fi9++ {
            switch -1 {
            case -565:
                var v10 uint64 = 0
                _ = v10
            case -999:
                v11 := GenS2{F0: ((32767 > 1000) && true), F1: 32767, F2: 42}
                _ = v11
            }
            _ = fi9
        }
    }
    v2 = (v2 % (((v2 & v2)) | 1))
    v12 := GenS3{F0: 0}
    switch "hello" {
    case "foo bar":
        if (0 <= 4294967295) {
            v5 = append(v5, int(v2))
            for fi13 := 0; fi13 < 20; fi13++ {
                v2 = v2
                v2 -= 42
                _ = fi13
            }
        }
    case "hello":
        if ("hello" == "foo bar") {
            var v14 int32 = 704
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 14; fi15++ {
        for fi16 := 0; fi16 < 2; fi16++ {
            switch fi16 {
            case 1:
                fmt.Println(fi15, fi15, fi15)
                _ = len(v1)
                if false {
                    if false {
                        fmt.Println(fi15, v2, fi15)
                        v2 += (v2 + (v2 + v2))
                        _ = len(v6)
                    }
                }
            case -284:
                v4[93] = 490
            default:
                v17 := GenS4{F0: 100, F1: 33, F2: -933}
                _ = v17
            }
            if true {
                fmt.Println(fi15)
            }
            v2 -= 0
            _ = fi16
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi15
    }
    for fi18 := 0; fi18 < 19; fi18++ {
        v19 := GenS2{F0: (false || (0.1 < 1e10)), F1: 2, F2: 42}
        switch "" {
        case "a":
            if (!((true && ("" != "a")) && (fi18 <= (fi18 + (^(fi18)))))) {
                v2 = v2
            }
        case "":
            var v20 uint16 = 32767
            _ = v20
        default:
            v21 := GenS5{F0: 357, F1: false, F2: (false || true)}
            _ = v21
        }
        _ = fi18
        _ = v19
    }
    if true {
        for fi22 := 0; fi22 < 9; fi22++ {
            if (float32(1e10) > 1.5) {
                v6 = append(v6, "a")
            }
            var v23 uint64 = 396
            var v24 uint16 = 797
            _ = fi22
            _ = v23
            _ = v24
        }
    }
    v25 := GenS5{F0: 1, F1: false, F2: (42 > 736)}
    v2 -= v2
    var v26 uint32 = 0
    v5 = append(v5, 4294967295)
    if true {
        v2 = 1
        if (1 < 42) {
            v26 = (v26 | 0)
        } else {
            for fi27 := 0; fi27 < 18; fi27++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                fmt.Println(v2, v2)
                switch "a" {
                case "foo bar":
                    v26 += (v26 | 749)
                default:
                    fmt.Println(v2)
                }
                _ = fi27
            }
            if false {
                var v28 uint32 = 2147483648
                _ = v28
            }
        }
        fmt.Println(v26)
    } else {
        v26 = v26
    }
    v26 = 0
    if len(v5) > 0 {
        v5[0] = -2147483648
    }
    var v29 uint64 = 18446744073709551615
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
