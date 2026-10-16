package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 uint8; F1 uint8; F2 uint32 }

func main() {
    var v1 int16 = -32768
    var v2 float64 = 0.0
    var v3 uint16 = 0
    v4 := make(map[string]uint32)
    v1 = (-(v1))
    v4[("foo bar" + "\n\t")] = 0
    var v5 string = "a"
    for fi6 := 0; fi6 < 15; fi6++ {
        v7 := GenS1{F0: 4294967295, F1: (-1 != 0)}
        var v8 float64 = 0.1
        _ = fi6
        _ = v7
        _ = v8
    }
    v3 -= (v3 / ((722) | 1))
    switch (v5 + (v5 + "a")) {
    case "\n\t":
        v5 = (v5 + "\n\t")
    default:
        if true {
            v9 := GenS2{F0: 2, F1: 208, F2: 65535}
            for fi10 := 0; fi10 < 15; fi10++ {
                var v11 bool = (-464 <= 9223372036854775807)
                _ = fi10
                _ = v11
            }
            var v12 uint64 = 42
            _ = v12
            _ = v9
        } else {
            v5 = v5
            v13 := GenS1{F0: 325, F1: true}
            _ = v13
        }
    }
    if (false && true) {
        if false {
            switch v5 {
            case "a":
                for fi14 := 0; fi14 < 8; fi14++ {
                    var v15 int8 = -1
                    _ = fi14
                    _ = v15
                }
            }
            var v16 int32 = 1
            _ = v16
        }
        if false {
            _ = len(v4)
        }
        if false {
            fmt.Println(v1)
        }
    }
    if false {
        v17 := GenS2{F0: 128, F1: 42, F2: 4294967295}
        v5 = v5
        _ = v17
    }
    for fi18 := 0; fi18 < 4; fi18++ {
        for fi19 := 0; fi19 < 8; fi19++ {
            var v20 uint64 = 42
            v5 = ""
            var v21 uint32 = 598
            _ = fi19
            _ = v20
            _ = v21
        }
        if (false && (9223372036854775807 < -344)) {
            var v22 uint64 = 4294967295
            _ = v22
        }
        _ = fi18
    }
    fmt.Println(v5, v5, v3)
    for fi23 := 0; fi23 < 3; fi23++ {
        var v24 int8 = 42
        switch (fi23 / (((^(fi23))) | 1)) {
        case 557:
            var v25 string = "\n\t"
            _ = v25
        case -2147483648:
            for fi26 := 0; fi26 < 20; fi26++ {
                fmt.Println(fi26)
                _ = fi26
            }
        case 2147483647:
            v27 := []int{(fi23 - -1), (fi23 << ((2) & 63)), fi23, fi23}
            _ = v27
        }
        _ = fi23
        _ = v24
    }
    var v28 uint16 = 820
    var v29 float32 = 0.0
    v30 := GenS1{F0: 217, F1: true}
    v2 = 1.5
    var v31 int8 = 119
    fmt.Println(v28, v29, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
