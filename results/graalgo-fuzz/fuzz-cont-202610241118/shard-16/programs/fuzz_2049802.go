package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 uint8; F1 int16; F2 uint32 }
type GenS4 struct { F0 int16; F1 uint16 }

func main() {
    var v1 uint16 = 58
    v2 := make(map[int]uint32)
    var v3 uint64 = 18446744073709551615
    v4 := make(map[uint32]bool)
    var v5 bool = false
    switch -378 {
    case 32767:
        fmt.Println(v3)
    default:
        var v6 uint64 = (v3 >> ((42) & 63))
        _ = v6
    }
    fmt.Println(v1, v1, v3)
    v1 = 0
    for fi7 := 0; fi7 < 18; fi7++ {
        for fi8 := 0; fi8 < 1; fi8++ {
            v5 = v5
            if ((255 == v1) || ((^(v3)) < v3)) {
                delete(v2, 9223372036854775807)
            }
            _ = fi8
        }
        switch ("foo bar" + "a") {
        case "foo bar":
            var v9 float64 = 0.0
            _ = v9
        case "\n\t":
            var v10 string = ""
            _ = v10
        case "":
            for fi11 := 0; fi11 < 5; fi11++ {
                if true {
                    v5 = (!(42 != 2147483648))
                    v12 := v2[892]
                    _ = v12
                }
                if v5 {
                    switch fi7 {
                    case 42:
                        switch (fi11 >> ((2) & 63)) {
                        case 1:
                            v13 := []int{(fi11 | -860), fi7, (fi11 - fi7)}
                            _ = v13
                        case 0:
                            var v14 uint = 4294967295
                            _ = v14
                        case 4294967295:
                            v15 := GenS1{F0: -216, F1: (v1 % (((^(v1))) | 1))}
                            _ = v15
                        }
                    case 32767:
                        v1 = (v1 % ((32767) | 1))
                    case -1:
                        var v16 bool = v5
                        _ = v16
                    }
                    v17 := []int{-9223372036854775808, 4294967295}
                    _ = v17
                }
                _ = fi11
            }
        }
        _ = fi7
    }
    v1 = v1
    if v5 {
        for fi18 := 0; fi18 < 12; fi18++ {
            var v19 int64 = 1
            v19 = v19
            _ = fi18
            _ = v19
        }
    } else {
        v4[2] = (v5 || (-874 > -1))
    }
    fmt.Println(v1, v1, v5)
    v20 := GenS2{F0: "hello", F1: -28}
    var v21 bool = (v5 && (645 <= 2147483647))
    v22 := GenS2{F0: ("" + ("hello" + ("foo bar" + ""))), F1: -93}
    v23 := []int{}
    var v24 uint64 = 0
    var v25 float64 = (-2.25 * 0.0)
    var v26 uint64 = 619
    if (!false) {
        v27 := GenS3{F0: 220, F1: 443, F2: 0}
        v28 := GenS4{F0: -32768, F1: (v1 | 65535)}
        var v29 int32 = 42
        _ = v27
        _ = v28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
