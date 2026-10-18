package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int32 }
type GenS2 struct { F0 int64 }

func main() {
    var v1 uint16 = 65535
    var v2 uint16 = (v1 ^ v1)
    var v3 uint32 = 4294967295
    var v4 uint64 = uint64(v3)
    var v5 uint8 = 42
    var v6 uint64 = v4
    if (1.5 < (-2.25 + 3.14159)) {
        v3 = 42
        fmt.Println(v1, v3, v4)
        fmt.Println(v6, v2)
    } else {
        var v7 uint32 = v3
        for fi8 := 0; fi8 < 12; fi8++ {
            if true {
                if true {
                    v9 := []int{(fi8 << ((0) & 63)), fi8, (fi8 / ((fi8) | 1))}
                    fmt.Println(v1)
                    _ = v9
                }
                var v10 int64 = 0
                var v11 uint64 = 4294967295
                _ = v10
                _ = v11
            }
            _ = fi8
        }
        _ = v7
    }
    v12 := []int{330, 4294967295}
    v13 := GenS1{F0: (v1 / ((v1) | 1)), F1: (^(v5)), F2: 645}
    for fi14 := 0; fi14 < 14; fi14++ {
        switch ((("\n\t" + "\n\t") + "\n\t") + "") {
        case "foo bar":
            v6 -= v4
        case "\n\t":
            if true {
                v15 := GenS1{F0: v1, F1: uint8(v3), F2: 935}
                if false {
                    v4 = (v6 % (((v4 >> ((980) & 63))) | 1))
                    fmt.Println(v3, v4, v2)
                }
                fmt.Println(v5, v3)
                _ = v15
            }
        }
        _ = fi14
    }
    for fi16 := 0; fi16 < 12; fi16++ {
        v12 = append(v12, 32768)
        var v17 string = ""
        _ = fi16
        _ = v17
    }
    v1 = uint16(v3)
    switch -559 {
    case 491:
        for fi18 := 0; fi18 < 2; fi18++ {
            v3 -= (^(v3))
            _ = fi18
        }
    case -558:
        v1 = (v1 + 0)
    default:
        var v19 uint32 = v3
        _ = v19
    }
    v12 = append(v12, 9223372036854775807)
    if len(v12) > 0 {
        v12[0] = 853
    }
    if (22 >= 2147483648) {
        var v20 uint16 = (v1 / ((0) | 1))
        for fi21 := 0; fi21 < 18; fi21++ {
            v22 := GenS2{F0: 42}
            _ = fi21
            _ = v22
        }
        _ = v20
    }
    fmt.Println(v3)
    var v23 int16 = 42
    var v24 float32 = 0.1
    var v25 int64 = -1
    var v26 uint64 = 577
    var v27 int = 32768
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    for fi28 := 0; fi28 < 1; fi28++ {
        fmt.Println(v5)
        fmt.Println(v27, v5)
        _ = fi28
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
