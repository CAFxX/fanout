package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 int = 127
    v2 := make(map[string]uint32)
    var v3 uint8 = 42
    v4 := []string{}
    v5 := make(map[uint32]bool)
    var v6 uint = 46
    if true {
        v7 := GenS1{F0: "hello"}
        if (true || (-2147483648 != -1)) {
            for fi8 := 0; fi8 < 7; fi8++ {
                var v9 uint32 = 24
                _ = fi8
                _ = v9
            }
        } else {
            fmt.Println(v1)
        }
        var v10 int32 = 0
        _ = v10
        _ = v7
    } else {
        v3 = 127
        v11 := GenS1{F0: "foo bar"}
        delete(v5, 0)
        _ = v11
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        switch -2147483648 {
        case 9223372036854775807:
            v5[583] = false
        case 32767:
            fmt.Println(v3, fi12, v3)
        }
        _ = fi12
    }
    var v13 uint32 = 1
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi14 := 0; fi14 < 19; fi14++ {
        v1 -= v1
        v15 := GenS1{F0: (("" + "foo bar") + "hello")}
        if len(v4) > 0 {
            v4[0] = "hello"
        }
        _ = fi14
        _ = v15
    }
    switch "\n\t" {
    case "":
        switch "\n\t" {
        case "foo bar":
            var v16 uint32 = uint32(v6)
            _ = v16
        case "a":
            fmt.Println(v3)
        case "hello":
            var v17 float64 = 1e10
            _ = v17
        default:
            fmt.Println(v13, v13)
        }
    case "foo bar":
        for fi18 := 0; fi18 < 15; fi18++ {
            v19 := v5[(v13 << (((v6 | 1)) & 63))]
            _ = fi18
            _ = v19
        }
    case "\n\t":
        for fi20 := 0; fi20 < 14; fi20++ {
            fmt.Println(v3, v13, fi20)
            v21 := GenS1{F0: "a"}
            _ = fi20
            _ = v21
        }
    default:
        var v22 uint = v6
        _ = v22
    }
    v23 := GenS1{F0: "hello"}
    var v24 bool = false
    var v25 int16 = 42
    fmt.Println(v25, v13)
    fmt.Println(v13, v3, v1)
    for fi26 := 0; fi26 < 18; fi26++ {
        v6 = 255
        switch (fi26 / (((v1 % (((fi26 ^ fi26)) | 1))) | 1)) {
        case 738:
            v6 = uint(v3)
        case 0:
            v27 := GenS1{F0: "\n\t"}
            _ = v27
        default:
            v28 := GenS1{F0: ("" + "a")}
            _ = v28
        }
        for fi29 := 0; fi29 < 19; fi29++ {
            fmt.Println(v13, v24)
            _ = fi29
        }
        _ = fi26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
