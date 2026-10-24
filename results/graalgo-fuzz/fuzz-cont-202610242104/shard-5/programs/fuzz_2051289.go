package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint64; F1 int32; F2 string }
func genFn0(p0 uint8, p1 uint) string {
    var v4 uint8 = (p0 << (((p1 << (((p1 + 1)) & 63))) & 63))
    var v5 float32 = ((1.5 + 1.5) * 0.1)
    v6 := GenS1{F0: 0, F1: 17}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return "foo bar"
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]int)
    v3 := make(map[int]uint32)
    switch 267 {
    case -691:
        var v7 uint16 = 42
        _ = v7
    case -935:
        for fi8 := 0; fi8 < 5; fi8++ {
            switch "hello" {
            case "\n\t":
                var v9 int16 = 789
                _ = v9
            case "":
                var v10 uint8 = 1
                _ = v10
            case "foo bar":
                var v11 uint8 = 189
                _ = v11
            }
            _ = fi8
        }
    default:
        if (!true) {
            var v12 int64 = 42
            v12 = 9223372036854775807
            var v13 int64 = v12
            _ = v12
            _ = v13
        } else {
            v14 := GenS2{F0: 32767}
            var v15 int = 32767
            fmt.Println(v15)
            _ = v14
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        v17 := GenS2{F0: fi16}
        switch 4294967295 {
        case 0:
            for fi18 := 0; fi18 < 20; fi18++ {
                v19 := GenS3{F0: "\n\t"}
                var v20 bool = (fi18 < fi16)
                v20 = true
                delete(v2, 2147483647)
                _ = fi18
                _ = v19
                _ = v20
            }
        case 32767:
            fmt.Println(fi16)
        default:
            var v21 string = "hello"
            _ = v21
        }
        _ = fi16
        _ = v17
    }
    var v22 int64 = -9223372036854775808
    v23 := GenS4{F0: 2, F1: 2147483647, F2: ""}
    v24 := GenS1{F0: 1, F1: 1}
    v25 := []int{}
    fmt.Println(v22)
    switch 4294967295 {
    case 128:
        v26 := v3[9223372036854775807]
        _ = v26
    }
    fmt.Println(v22)
    var v27 uint = 255
    var v28 float64 = (0.0 * (-2.25 / 1e10))
    for fi29 := 0; fi29 < 11; fi29++ {
        switch ("\n\t" + ("\n\t" + "")) {
        case "a":
            var v30 bool = true
            _ = v30
        }
        _ = fi29
    }
    v28 = v28
    for fi31 := 0; fi31 < 15; fi31++ {
        v22 -= (v22 >> (((^(v27))) & 63))
        _ = fi31
    }
    if (42 != 599) {
        switch (("" + ("\n\t" + ("foo bar" + (("foo bar" + "hello") + "foo bar")))) + ("hello" + (("" + "foo bar") + "a"))) {
        case "foo bar":
            v22 += -792
        }
    } else {
        var v32 int8 = 1
        _ = v32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v27
    _ = v28
    _ = v3
}
