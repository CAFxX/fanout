package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 uint64; F1 int64; F2 uint16 }
type GenS3 struct { F0 uint8; F1 string; F2 int64 }
type GenS4 struct { F0 int8 }

func main() {
    var v1 uint = 18446744073709551615
    var v2 int8 = -1
    v3 := make(map[string]uint32)
    v4 := make(map[int]bool)
    var v5 uint32 = 65535
    var v6 uint64 = 0
    v3[("hello" + "")] = 2147483648
    for fi7 := 0; fi7 < 16; fi7++ {
        for fi8 := 0; fi8 < 3; fi8++ {
            if true {
                v9 := GenS1{F0: (-(v2)), F1: 20}
                if false {
                    var v10 int16 = 32767
                    _ = v10
                }
                var v11 uint16 = 1
                _ = v11
                _ = v9
            } else {
                fmt.Println(v1, v5)
            }
            _ = fi8
        }
        if ((121 != v2) && ((v5 * v5) <= v5)) {
            v5 = uint32(v1)
            v12 := []int{fi7, fi7}
            _ = v12
        }
        _ = fi7
    }
    v13 := GenS2{F0: (^(v6)), F1: 9223372036854775807, F2: 42}
    switch -638 {
    case 0:
        v5 = v5
    default:
        v1 -= (^(v1))
    }
    switch int(v2) {
    case 4294967295:
        var v14 string = "\n\t"
        _ = v14
    default:
        v15 := v3[(("hello" + "\n\t") + "a")]
        _ = v15
    }
    for fi16 := 0; fi16 < 12; fi16++ {
        switch 128 {
        case 9223372036854775807:
            var v17 int32 = -727
            _ = v17
        case 42:
            fmt.Println(v5, v2)
        case -492:
            v18 := GenS3{F0: 127, F1: ("" + ("" + "a")), F2: 1}
            _ = v18
        }
        _ = fi16
    }
    fmt.Println(v6, v2)
    if true {
        if ((v2 <= (v2 / (((v2 - (v2 & (v2 - (v2 / ((100) | 1)))))) | 1))) || false) {
            if true {
                v6 += 42
                switch "a" {
                case "\n\t":
                    var v19 uint64 = (v6 | (v6 + 279))
                    _ = v19
                case "foo bar":
                    var v20 int64 = 42
                    _ = v20
                case "":
                    v1 = 277
                }
            } else {
                v2 += (v2 % ((42) | 1))
            }
            switch 42 {
            case 127:
                var v21 uint64 = (v6 * (^(v6)))
                _ = v21
            default:
                v6 = (^(v6))
            }
        } else {
            v6 = 1
            if (v6 != (v6 << (((v1 * (v1 >> ((v1) & 63)))) & 63))) {
                v5 = v5
            } else {
                v22 := GenS1{F0: v2, F1: -128}
                v1 = v1
                _ = v22
            }
        }
        v23 := []int{923}
        _ = v23
    } else {
        v1 += 1
    }
    for fi24 := 0; fi24 < 11; fi24++ {
        var v25 uint32 = 2
        v26 := v4[int(v2)]
        v6 = 18446744073709551615
        _ = fi24
        _ = v25
        _ = v26
    }
    v27 := []int{127}
    for fi28 := 0; fi28 < 4; fi28++ {
        fmt.Println(v5, v1, fi28)
        switch ("\n\t" + (("foo bar" + "a") + "a")) {
        case "":
            var v29 bool = false
            _ = v29
        case "foo bar":
            fmt.Println(v1)
        }
        _ = fi28
    }
    for fi30 := 0; fi30 < 20; fi30++ {
        var v31 int64 = 42
        var v32 bool = (32767 == 32767)
        _ = fi30
        _ = v31
        _ = v32
    }
    for fi33 := 0; fi33 < 13; fi33++ {
        v34 := GenS4{F0: 1}
        var v35 uint16 = 665
        _ = fi33
        _ = v34
        _ = v35
    }
    var v36 int64 = 789
    fmt.Println(v1, v2, v1)
    for fi37 := 0; fi37 < 10; fi37++ {
        v1 = v1
        _ = len(v4)
        v38 := v3[""]
        _ = fi37
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v27
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
