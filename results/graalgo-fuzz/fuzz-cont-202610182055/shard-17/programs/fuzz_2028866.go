package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 uint }

func main() {
    var v1 uint8 = 128
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]bool)
    var v4 int16 = -1
    var v5 uint64 = 2
    var v6 uint32 = 1
    v6 = (v6 & 513)
    var v7 int8 = 100
    for fi8 := 0; fi8 < 6; fi8++ {
        v5 += v5
        _ = fi8
    }
    var v9 string = "foo bar"
    if true {
        v10 := GenS1{F0: 1, F1: -1, F2: 4294967295}
        _ = v10
    }
    switch v9 {
    case "a":
        for fi11 := 0; fi11 < 4; fi11++ {
            delete(v2, v6)
            fmt.Println(v9, v9)
            if true {
                v5 -= (^(v5))
            }
            _ = fi11
        }
    default:
        switch v9 {
        case "":
            var v12 bool = false
            _ = v12
        case "foo bar":
            if (4294967295 == -1) {
                if (((v1 < v1) || true) || true) {
                    if (128 == 128) {
                        v13 := GenS1{F0: 0, F1: -888, F2: 42}
                        for fi14 := 0; fi14 < 8; fi14++ {
                            fmt.Println(fi14, v9, v5)
                            v15 := v3[(v6 - (v6 + 683))]
                            _ = fi14
                            _ = v15
                        }
                        _ = v13
                    }
                    var v16 bool = true
                    fmt.Println(v6, v5, v7)
                    _ = v16
                }
                var v17 uint16 = 1
                for fi18 := 0; fi18 < 2; fi18++ {
                    var v19 uint32 = v6
                    switch -474 {
                    case 42:
                        v9 = v9
                    }
                    _ = fi18
                    _ = v19
                }
                _ = v17
            } else {
                fmt.Println(v9, v5, v4)
                v20 := GenS1{F0: 9223372036854775807, F1: 1, F2: 255}
                for fi21 := 0; fi21 < 4; fi21++ {
                    var v22 int8 = 42
                    _ = fi21
                    _ = v22
                }
                _ = v20
            }
        default:
            var v23 string = v9
            _ = v23
        }
    }
    v7 += (v7 % ((-10) | 1))
    fmt.Println(v4, v9)
    fmt.Println(v1, v9)
    for fi24 := 0; fi24 < 13; fi24++ {
        var v25 int32 = 0
        fmt.Println(v25)
        fmt.Println(fi24)
        _ = fi24
        _ = v25
    }
    v5 = 260
    var v26 bool = true
    for fi27 := 0; fi27 < 1; fi27++ {
        v28 := GenS1{F0: int(v4), F1: 1, F2: uint(v1)}
        if (v26 || v26) {
            v29 := GenS1{F0: -800, F1: 42, F2: 4294967295}
            var v30 uint32 = 4294967295
            v9 = ""
            _ = v29
            _ = v30
        } else {
            var v31 string = (((v9 + v9) + (("hello" + v9) + v9)) + v9)
            switch (v9 + (("foo bar" + "") + v9)) {
            case "a":
                v5 -= v5
            }
            _ = v31
        }
        _ = fi27
        _ = v28
    }
    fmt.Println(v9, v7, v26)
    var v32 uint8 = (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
