package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16; F1 int8; F2 uint64 }
type GenS4 struct { F0 int; F1 int16 }
type GenS5 struct { F0 string }

func main() {
    var v1 float32 = float32(1e10)
    v2 := []int{}
    v3 := make(map[string]string)
    v4 := []int{}
    v5 := make(map[uint32]uint32)
    var v6 int64 = -201
    if len(v4) > 0 {
        v4[0] = -302
    }
    fmt.Println(v6)
    fmt.Println(v6)
    v7 := GenS1{F0: ("hello" + "foo bar"), F1: (!(250 > 128))}
    v3["\n\t"] = "a"
    v1 += (v1 / 1.5)
    v8 := GenS2{F0: 2}
    if true {
        switch 127 {
        case 128:
            var v9 float64 = -2.25
            _ = v9
        case 1:
            v10 := GenS3{F0: 280, F1: 0, F2: 1}
            _ = v10
        case 2147483647:
            _ = len(v5)
        }
        v11 := GenS2{F0: 1}
        _ = v11
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        switch (("" + "hello") + ("\n\t" + "\n\t")) {
        case "":
            for fi13 := 0; fi13 < 3; fi13++ {
                v14 := v3[("" + "foo bar")]
                fmt.Println(fi12, fi13)
                if true {
                    v6 -= 42
                } else {
                    switch "a" {
                    case "foo bar":
                        switch (v14 + v14) {
                        case "a":
                            fmt.Println(fi12)
                            v15 := GenS2{F0: 100}
                            _ = v15
                        }
                    case "hello":
                        v6 += v6
                    case "a":
                        v6 -= v6
                    }
                }
                _ = fi13
                _ = v14
            }
        }
        v16 := GenS4{F0: fi12, F1: 0}
        _ = fi12
        _ = v16
    }
    if ((true && true) || true) {
        v17 := GenS2{F0: 127}
        if (1 == 235) {
            v1 = 1e10
            v6 += (v6 / ((v6) | 1))
        }
        for fi18 := 0; fi18 < 7; fi18++ {
            v1 += (float32(-2.25) - 0.0)
            _ = fi18
        }
        _ = v17
    } else {
        var v19 uint16 = 65535
        fmt.Println(v19, v19)
        v1 = v1
        _ = v19
    }
    switch "\n\t" {
    case "a":
        v20 := GenS4{F0: -9223372036854775808, F1: 1000}
        _ = v20
    case "hello":
        var v21 int64 = (v6 * (v6 - (v6 ^ v6)))
        _ = v21
    }
    delete(v5, 1)
    v6 = 0
    for fi22 := 0; fi22 < 20; fi22++ {
        if true {
            var v23 uint64 = 609
            _ = v23
        }
        fmt.Println(v6, fi22)
        v5[0] = 485
        _ = fi22
    }
    var v24 uint32 = 4294967295
    for fi25 := 0; fi25 < 18; fi25++ {
        v26 := GenS5{F0: "hello"}
        v24 = v24
        v5[4294967295] = (v24 / ((v24) | 1))
        _ = fi25
        _ = v26
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
