package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 int64; F1 int64; F2 uint32 }
type GenS3 struct { F0 uint64; F1 uint32 }
type GenS4 struct { F0 uint64; F1 int; F2 uint }
type GenS5 struct { F0 uint32; F1 uint }

func main() {
    v1 := []string{"\n\t", "hello", ("" + "a"), "\n\t"}
    var v2 int32 = 42
    v3 := make(map[string]uint32)
    if (!false) {
        if true {
            var v4 uint32 = 161
            v5 := GenS1{F0: 9223372036854775807, F1: 1000}
            v4 = (^(v4))
            _ = v4
            _ = v5
        }
        switch -672 {
        case 44:
            _ = len(v3)
        case 0:
            var v6 uint32 = 134
            _ = v6
        }
        if false {
            switch "hello" {
            case "foo bar":
                v2 -= 0
            }
            fmt.Println(v2)
        } else {
            v7 := GenS2{F0: -9223372036854775808, F1: 1, F2: 2147483647}
            v8 := GenS1{F0: 1, F1: 1}
            _ = v7
            _ = v8
        }
    } else {
        var v9 float64 = float64((-2.25 * ((1.5 + (1e10 - -2.25)) * 3.14159)))
        v2 = (-(v2))
        _ = v9
    }
    v2 = 42
    v3[(("foo bar" + "") + ("a" + (("foo bar" + "foo bar") + "hello")))] = 0
    for fi10 := 0; fi10 < 8; fi10++ {
        v11 := GenS1{F0: 213, F1: -32768}
        _ = fi10
        _ = v11
    }
    if (992 < 817) {
        if true {
            v12 := v3["\n\t"]
            _ = v12
        } else {
            var v13 uint = 255
            _ = v13
        }
        for fi14 := 0; fi14 < 10; fi14++ {
            v15 := GenS1{F0: 107, F1: 317}
            fmt.Println(v2, fi14)
            _ = fi14
            _ = v15
        }
    } else {
        fmt.Println(v2)
    }
    if false {
        var v16 int32 = (v2 >> ((0) & 63))
        _ = v16
    } else {
        var v17 string = "hello"
        _ = v17
    }
    for fi18 := 0; fi18 < 9; fi18++ {
        v2 -= (v2 - (-(v2)))
        if false {
            var v19 float64 = (0.1 + 1e10)
            var v20 string = "\n\t"
            _ = v19
            _ = v20
        }
        v21 := GenS3{F0: 0, F1: 26}
        _ = fi18
        _ = v21
    }
    var v22 int32 = 0
    var v23 int8 = 127
    for fi24 := 0; fi24 < 16; fi24++ {
        for fi25 := 0; fi25 < 6; fi25++ {
            if false {
                var v26 uint32 = 935
                _ = v26
            }
            var v27 int = int(v2)
            v23 = 1
            _ = fi25
            _ = v27
        }
        switch "hello" {
        case "":
            fmt.Println(v2, v23)
        }
        _ = fi24
    }
    v28 := GenS2{F0: -9223372036854775808, F1: 42, F2: 24}
    v23 += (-(v23))
    v29 := GenS4{F0: 2, F1: -977, F2: 774}
    var v30 int8 = v23
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    _ = len(v3)
    if (1 > 1) {
        v31 := GenS5{F0: 2147483647, F1: 4294967295}
        var v32 uint64 = 578
        _ = v31
        _ = v32
    }
    fmt.Println(v22, v30)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v30
}
