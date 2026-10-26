package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 uint64; F1 uint64; F2 string }
type GenS3 struct { F0 int16; F1 int8 }
func genFn0(p0 uint, p1 uint32) int16 {
    p1 -= 65535
    var v5 uint = (p0 & p0)
    for fi6 := 0; fi6 < 19; fi6++ {
        v7 := []int{}
        switch -856 {
        case 127:
            switch "\n\t" {
            case "a":
                var v8 uint64 = 370
                _ = v8
            }
        case 9223372036854775807:
            v9 := GenS1{F0: true, F1: -1}
            _ = v9
        }
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v5
    return -1
}
func genFn1() uint32 {
    var v10 uint64 = 2
    fmt.Println(v10)
    for fi11 := 0; fi11 < 17; fi11++ {
        var v12 uint8 = 42
        for fi13 := 0; fi13 < 6; fi13++ {
            v12 = v12
            v10 -= uint64(v12)
            _ = fi13
        }
        _ = fi11
        _ = v12
    }
    fmt.Println(v10)
    fmt.Println(v10)
    _ = v10
    return 2
}

func main() {
    var v1 uint32 = 2147483647
    v2 := []string{"\n\t", "foo bar"}
    v3 := []string{"foo bar", "foo bar", "a", ("a" + "")}
    var v4 int64 = 892
    v1 -= (v1 ^ (v1 + genFn1()))
    v1 -= v1
    for fi14 := 0; fi14 < 20; fi14++ {
        v4 = v4
        switch (("foo bar" + ("" + "a")) + "\n\t") {
        case "hello":
            var v15 int16 = 1
            _ = v15
        }
        _ = fi14
    }
    v4 = 9223372036854775807
    switch (((("" + ((("a" + "hello") + "hello") + "")) + "a") + "a") + "a") {
    case "":
        fmt.Println(v4)
    case "a":
        for fi16 := 0; fi16 < 13; fi16++ {
            var v17 int16 = -1
            _ = fi16
            _ = v17
        }
    case "foo bar":
        var v18 string = ("" + "hello")
        _ = v18
    }
    v1 += (v1 << ((2) & 63))
    for fi19 := 0; fi19 < 11; fi19++ {
        v20 := GenS2{F0: 691, F1: 1, F2: "foo bar"}
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 18; fi21++ {
        var v22 int64 = 602
        _ = fi21
        _ = v22
    }
    for fi23 := 0; fi23 < 12; fi23++ {
        fmt.Println(v1)
        if ((1e10 / 0.1) < (1e10 - float64(3.14159))) {
            v1 -= v1
            v24 := GenS3{F0: 880, F1: -26}
            fmt.Println(v4, fi23, v1)
            _ = v24
        } else {
            v25 := GenS1{F0: true, F1: 1000}
            if (3.14159 < (((-2.25 + (0.1 + (((1.5 * 0.0) + ((0.0 - (0.1 - (3.14159 / (1e10 - 0.1)))) * 1.5)) - 3.14159))) / 0.1) / -2.25)) {
                var v26 float32 = float32(0.0)
                _ = v26
            } else {
                v4 -= (v4 ^ 42)
                var v27 uint16 = 32767
                _ = v27
            }
            _ = v25
        }
        _ = fi23
    }
    if false {
        v4 = v4
    }
    if false {
        v1 -= 2147483647
        v1 = genFn1()
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
