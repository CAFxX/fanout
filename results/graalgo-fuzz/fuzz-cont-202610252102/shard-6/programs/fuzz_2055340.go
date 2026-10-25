package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 string; F1 int32 }
func genFn0(p0 int, p1 int64) int64 {
    p0 = p0
    v4 := GenS1{F0: 0, F1: 1000}
    v5 := []int{p0}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return p1
}

func main() {
    v1 := make(map[string]string)
    var v2 int16 = 708
    var v3 int32 = 2147483647
    v3 = v3
    v2 -= v2
    switch -1 {
    case -1:
        for fi6 := 0; fi6 < 9; fi6++ {
            for fi7 := 0; fi7 < 16; fi7++ {
                v8 := GenS2{F0: fi7, F1: 0}
                _ = fi7
                _ = v8
            }
            if (1.5 >= 3.14159) {
                var v9 uint32 = 2147483648
                _ = len(v1)
                _ = v9
            }
            _ = fi6
        }
    case 311:
        for fi10 := 0; fi10 < 16; fi10++ {
            var v11 float32 = 0.0
            _ = fi10
            _ = v11
        }
    }
    fmt.Println(v2, v3)
    v12 := GenS3{F0: ("" + "a"), F1: -341}
    var v13 uint8 = 255
    v2 = (^(v2))
    for fi14 := 0; fi14 < 17; fi14++ {
        v15 := GenS1{F0: 42, F1: 1000}
        _ = fi14
        _ = v15
    }
    v3 = v3
    if false {
        for fi16 := 0; fi16 < 11; fi16++ {
            fmt.Println(v3, v13, fi16)
            _ = fi16
        }
        v3 += 42
        var v17 int64 = -1
        _ = v17
    } else {
        for fi18 := 0; fi18 < 17; fi18++ {
            v3 += v3
            _ = fi18
        }
    }
    if true {
        switch "hello" {
        case "":
            switch "" {
            case "hello":
                switch 0 {
                case 4294967295:
                    var v19 int64 = -1
                    _ = v19
                case 127:
                    v20 := []int{325}
                    _ = v20
                case 9223372036854775807:
                    if false {
                        v21 := v1["foo bar"]
                        _ = v21
                    }
                default:
                    var v22 int32 = (v3 << ((0) & 63))
                    _ = v22
                }
            case "a":
                var v23 float32 = 1e10
                _ = v23
            case "":
                var v24 int16 = v2
                _ = v24
            }
        }
    }
    v13 += 127
    if true {
        var v25 int = 729
        fmt.Println(v3, v2)
        var v26 int = v25
        _ = v25
        _ = v26
    } else {
        if true {
            var v27 int16 = (v2 ^ v2)
            _ = v27
        }
    }
    for fi28 := 0; fi28 < 2; fi28++ {
        if true {
            fmt.Println(v2)
        }
        v1[""] = ((("a" + "foo bar") + "\n\t") + "")
        _ = fi28
    }
    switch 9223372036854775807 {
    case -480:
        var v29 uint16 = 1
        _ = v29
    }
    v13 = (^(v13))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
}
