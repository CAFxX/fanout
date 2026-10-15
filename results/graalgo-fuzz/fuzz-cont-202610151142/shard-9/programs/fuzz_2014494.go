package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 uint16; F2 string }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 int; F1 bool }
func genFn0() uint64 {
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 int = (-(fi6))
        v8 := GenS1{F0: "a"}
        for fi9 := 0; fi9 < 8; fi9++ {
            var v10 int = (^(v7))
            v7 = (fi6 / (((fi9 + fi9)) | 1))
            _ = fi9
            _ = v10
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    var v11 bool = (127 <= -92)
    _ = v11
    return 18446744073709551615
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 int = 128
    var v3 uint64 = 903
    v4 := []string{"hello", "foo bar", "a", ("\n\t" + "foo bar")}
    var v5 uint16 = 0
    switch (v2 >> ((uint(v5)) & 63)) {
    case 9223372036854775807:
        for fi12 := 0; fi12 < 17; fi12++ {
            v13 := GenS2{F0: ((1.5 - 1e10) >= 1.5), F1: (v5 & v5), F2: "hello"}
            _ = fi12
            _ = v13
        }
    case -2147483648:
        if false {
            v1 += v1
        } else {
            switch v2 {
            case 2147483647:
                var v14 int32 = -1
                _ = v14
            case 1:
                v5 = v5
            }
        }
    default:
        v1 = v1
    }
    var v15 int8 = -101
    for fi16 := 0; fi16 < 19; fi16++ {
        v5 += (v5 << ((255) & 63))
        var v17 int = fi16
        _ = fi16
        _ = v17
    }
    if (false && true) {
        var v18 uint64 = (v3 >> ((255) & 63))
        _ = v18
    }
    if true {
        v3 += v3
        v19 := GenS3{F0: v1, F1: (v3 + v3)}
        v5 += v5
        _ = v19
    } else {
        if (!true) {
            fmt.Println(v2, v2)
        }
    }
    if true {
        for fi20 := 0; fi20 < 3; fi20++ {
            switch ("\n\t" + "") {
            case "\n\t":
                v15 = (v15 - v15)
            }
            if (false || (9223372036854775807 > -303)) {
                var v21 int32 = -2147483648
                if len(v4) > 0 {
                    v4[0] = "foo bar"
                }
                fmt.Println(fi20, v15, v3)
                _ = v21
            }
            if len(v4) > 0 {
                v4[0] = "a"
            }
            _ = fi20
        }
    }
    v22 := GenS4{F0: v2, F1: true}
    fmt.Println(v5)
    fmt.Println(v5)
    var v23 int8 = (v15 - (v15 / ((v15) | 1)))
    for fi24 := 0; fi24 < 4; fi24++ {
        var v25 int8 = v23
        _ = fi24
        _ = v25
    }
    v15 -= (v15 % (((v23 * v15)) | 1))
    fmt.Println(v3, v1)
    switch -9223372036854775808 {
    case 128:
        switch "a" {
        case "":
            v2 += (v2 | v2)
        case "a":
            if true {
                v3 = uint64(v5)
            } else {
                var v26 bool = true
                var v27 uint64 = v3
                _ = v26
                _ = v27
            }
        case "hello":
            var v28 int16 = 1
            _ = v28
        default:
            v4 = append(v4, "")
        }
    default:
        for fi29 := 0; fi29 < 9; fi29++ {
            v2 = (fi29 >> ((65535) & 63))
            fmt.Println(v5)
            var v30 uint32 = 65535
            _ = fi29
            _ = v30
        }
    }
    if len(v4) > 0 {
        v4[0] = "a"
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
