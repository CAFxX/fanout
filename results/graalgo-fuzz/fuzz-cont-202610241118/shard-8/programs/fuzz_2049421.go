package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 uint64; F1 int64 }

func main() {
    v1 := []string{(("" + "hello") + "")}
    var v2 int64 = 964
    v3 := []uint32{1}
    var v4 int = 9223372036854775807
    v5 := make(map[string]int)
    for fi6 := 0; fi6 < 7; fi6++ {
        if (255 != 32767) {
            v4 += v4
            var v7 uint64 = 0
            for fi8 := 0; fi8 < 11; fi8++ {
                v4 = (fi8 - (fi8 >> ((18446744073709551615) & 63)))
                if false {
                    v2 += (v2 & v2)
                    v9 := GenS1{F0: 42}
                    _ = v9
                }
                _ = fi8
            }
            _ = v7
        }
        var v10 uint64 = 302
        var v11 int64 = v2
        _ = fi6
        _ = v10
        _ = v11
    }
    switch (v4 & v4) {
    case -9223372036854775808:
        v12 := GenS2{F0: ("a" + "foo bar"), F1: ""}
        _ = v12
    default:
        _ = len(v5)
    }
    for fi13 := 0; fi13 < 4; fi13++ {
        var v14 uint16 = 0
        _ = fi13
        _ = v14
    }
    if ((v4 | (-(v4))) == (v4 % (((v4 | -257)) | 1))) {
        var v15 uint64 = 42
        _ = v15
    }
    v16 := GenS1{F0: 0}
    _ = len(v1)
    v5["hello"] = (-(v4))
    if false {
        v2 = (v2 ^ 359)
        var v17 bool = (!false)
        for fi18 := 0; fi18 < 3; fi18++ {
            var v19 float64 = 0.1
            _ = fi18
            _ = v19
        }
        _ = v17
    }
    for fi20 := 0; fi20 < 18; fi20++ {
        v21 := GenS1{F0: 0}
        var v22 int64 = v2
        _ = fi20
        _ = v21
        _ = v22
    }
    v23 := GenS2{F0: "\n\t", F1: "\n\t"}
    v4 += v4
    v4 -= v4
    v4 = v4
    switch "foo bar" {
    case "":
        if (v4 == v4) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            switch "foo bar" {
            case "a":
                var v24 uint64 = 1
                _ = v24
            case "foo bar":
                v4 += -587
            case "hello":
                var v25 int64 = (-(v2))
                _ = v25
            }
            v2 += v2
        }
    default:
        var v26 uint = 255
        _ = v26
    }
    var v27 string = "foo bar"
    v28 := GenS3{F0: 0, F1: v2}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v2
    _ = v23
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
