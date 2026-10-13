package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool; F1 bool }
type GenS4 struct { F0 uint16 }

func main() {
    var v1 int8 = -11
    v2 := []uint32{469}
    var v3 uint32 = 1
    var v4 float32 = float32(float64(-2.25))
    switch "" {
    case "a":
        fmt.Println(v1, v4, v3)
    case "":
        v5 := GenS1{F0: 0, F1: -381}
        _ = v5
    case "foo bar":
        if true {
            v6 := GenS1{F0: -9223372036854775808, F1: 32767}
            _ = v6
        }
    default:
        if len(v2) > 0 {
            v2[0] = v3
        }
    }
    var v7 uint32 = v3
    switch ("" + ("" + "foo bar")) {
    case "\n\t":
        for fi8 := 0; fi8 < 19; fi8++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            fmt.Println(v7, v3, v1)
            _ = fi8
        }
    case "foo bar":
        var v9 int = 0
        _ = v9
    case "hello":
        if false {
            if (0 >= 42) {
                var v10 uint16 = 2
                fmt.Println(v1, v3, v4)
                _ = v10
            }
        }
    default:
        v2 = append(v2, (v7 ^ (v7 ^ v3)))
    }
    if (65535 < (^(v7))) {
        v2 = append(v2, (v3 & v7))
        switch ((("\n\t" + ("a" + "a")) + "a") + "hello") {
        case "foo bar":
            var v11 int64 = int64(v1)
            _ = v11
        case "\n\t":
            switch (("\n\t" + "foo bar") + "hello") {
            case "foo bar":
                var v12 uint64 = 2
                _ = v12
            }
        default:
            if false {
                var v13 int32 = 1
                v2 = append(v2, (v3 + (v3 + v7)))
                _ = v13
            } else {
                var v14 string = (("\n\t" + "a") + "foo bar")
                _ = v14
            }
        }
        var v15 uint = 488
        _ = v15
    }
    for fi16 := 0; fi16 < 13; fi16++ {
        for fi17 := 0; fi17 < 12; fi17++ {
            v18 := GenS2{F0: int64(fi17)}
            _ = fi17
            _ = v18
        }
        _ = fi16
    }
    var v19 int16 = -69
    var v20 uint32 = (^(v3))
    v21 := GenS1{F0: 1, F1: v19}
    for fi22 := 0; fi22 < 3; fi22++ {
        if false {
            v23 := GenS3{F0: (v19 >= v19), F1: false}
            switch 0 {
            case 32768:
                fmt.Println(fi22, v19, fi22)
            }
            v24 := GenS4{F0: 1}
            _ = v23
            _ = v24
        } else {
            fmt.Println(v20, fi22, v7)
            v25 := GenS4{F0: 1}
            _ = v25
        }
        _ = fi22
    }
    if (2 > 586) {
        for fi26 := 0; fi26 < 8; fi26++ {
            var v27 uint64 = 718
            _ = fi26
            _ = v27
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v7
}
