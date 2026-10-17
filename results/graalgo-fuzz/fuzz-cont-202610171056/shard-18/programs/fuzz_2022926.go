package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 int8 }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 string = "\n\t"
    v2 := []uint32{2147483648, 0}
    v3 := make(map[uint32]string)
    var v4 uint64 = 4294967295
    var v5 bool = ((v4 ^ 18446744073709551615) != v4)
    v4 = v4
    fmt.Println(v1)
    var v6 float32 = float32(3.14159)
    var v7 int = -1
    var v8 int32 = -2147483648
    fmt.Println(v4, v6, v1)
    if true {
        fmt.Println(v6)
        fmt.Println(v8, v8, v7)
    }
    if v5 {
        fmt.Println(v5)
    }
    v9 := GenS1{F0: false, F1: false, F2: 115}
    for fi10 := 0; fi10 < 12; fi10++ {
        switch "foo bar" {
        case "\n\t":
            fmt.Println(v6, v7)
        case "a":
            v11 := GenS1{F0: false, F1: (!(!v5)), F2: 0}
            _ = v11
        default:
            fmt.Println(v6, v7, v7)
        }
        var v12 int8 = 42
        _ = fi10
        _ = v12
    }
    switch v1 {
    case "a":
        for fi13 := 0; fi13 < 14; fi13++ {
            v8 -= v8
            _ = fi13
        }
    case "foo bar":
        for fi14 := 0; fi14 < 2; fi14++ {
            v15 := GenS1{F0: v5, F1: (!v5), F2: 0}
            _ = fi14
            _ = v15
        }
    case "\n\t":
        for fi16 := 0; fi16 < 3; fi16++ {
            _ = len(v2)
            _ = fi16
        }
    default:
        v6 = (v6 / 1.5)
    }
    v1 = "\n\t"
    var v17 uint64 = 0
    for fi18 := 0; fi18 < 15; fi18++ {
        v19 := GenS1{F0: (2147483648 != 65535), F1: (!v5), F2: -90}
        var v20 uint = 723
        _ = fi18
        _ = v19
        _ = v20
    }
    if true {
        v21 := GenS2{F0: 65535}
        _ = v21
    }
    if (1 != -26) {
        switch v7 {
        case 4294967295:
            for fi22 := 0; fi22 < 9; fi22++ {
                if (!("\n\t" != (v1 + ("foo bar" + "\n\t")))) {
                    var v23 float64 = 3.14159
                    v17 = (v4 << ((611) & 63))
                    _ = v23
                }
                _ = fi22
            }
        case -872:
            var v24 int8 = 0
            _ = v24
        }
        _ = len(v2)
    } else {
        var v25 bool = (963 > -1)
        if ((-2.25 + v6) == (v6 * float32(-2.25))) {
            switch v7 {
            case 127:
                if len(v2) > 0 {
                    v2[0] = 1
                }
            case -271:
                v6 = (v6 / (v6 / v6))
            case 9223372036854775807:
                var v26 uint = 2
                _ = v26
            }
            var v27 int16 = 1000
            _ = v27
        }
        var v28 int8 = 0
        _ = v25
        _ = v28
    }
    for fi29 := 0; fi29 < 4; fi29++ {
        var v30 uint32 = 65535
        if ((v4 - 177) >= 368) {
            fmt.Println(fi29, v30)
        }
        _ = fi29
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
