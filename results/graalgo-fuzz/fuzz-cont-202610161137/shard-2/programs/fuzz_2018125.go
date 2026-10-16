package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 uint8; F1 uint8; F2 int32 }

func main() {
    var v1 uint32 = 129
    var v2 int16 = -478
    var v3 int64 = -9223372036854775808
    var v4 float32 = float32(0.1)
    v5 := []int{-2147483648, 142, 2147483647, 42}
    if (-128 <= -1) {
        var v6 uint32 = 4294967295
        _ = v6
    } else {
        v7 := GenS1{F0: uint64(v1), F1: "\n\t"}
        _ = v7
    }
    var v8 int16 = v2
    switch ("hello" + "foo bar") {
    case "a":
        var v9 uint32 = v1
        _ = v9
    case "hello":
        v10 := GenS1{F0: 4294967295, F1: ("\n\t" + "")}
        _ = v10
    case "\n\t":
        if true {
            v8 -= (v8 ^ v2)
        } else {
            fmt.Println(v3, v8)
            v2 = 1
        }
    }
    for fi11 := 0; fi11 < 2; fi11++ {
        if len(v5) > 0 {
            v5[0] = (fi11 / ((fi11) | 1))
        }
        var v12 uint64 = 18446744073709551615
        v13 := GenS1{F0: (v12 | 2), F1: ("\n\t" + "foo bar")}
        _ = fi11
        _ = v12
        _ = v13
    }
    v14 := GenS1{F0: 749, F1: ("foo bar" + "foo bar")}
    for fi15 := 0; fi15 < 17; fi15++ {
        var v16 uint32 = v1
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi15
        _ = v16
    }
    if len(v5) > 0 {
        v5[0] = 0
    }
    v17 := GenS2{F0: 127, F1: 42, F2: -1}
    for fi18 := 0; fi18 < 4; fi18++ {
        for fi19 := 0; fi19 < 20; fi19++ {
            switch fi19 {
            case 2147483647:
                v20 := GenS1{F0: 1, F1: ""}
                _ = v20
            }
            fmt.Println(v1, fi18, v1)
            v5 = append(v5, 32768)
            _ = fi19
        }
        var v21 bool = false
        _ = fi18
        _ = v21
    }
    switch -2147483648 {
    case 42:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case 1:
        for fi22 := 0; fi22 < 18; fi22++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            for fi23 := 0; fi23 < 2; fi23++ {
                v24 := GenS1{F0: 691, F1: "foo bar"}
                fmt.Println(v2, v1)
                v25 := GenS1{F0: 0, F1: "a"}
                _ = fi23
                _ = v24
                _ = v25
            }
            _ = fi22
        }
    case 4294967295:
        fmt.Println(v2, v2)
    default:
        v3 = (v3 | (v3 / (((v3 * 9223372036854775807)) | 1)))
    }
    if (v1 != v1) {
        if false {
            if len(v5) > 0 {
                v5[0] = 32767
            }
        }
    }
    var v26 bool = (4294967295 == 1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
