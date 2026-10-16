package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 int32; F1 int32; F2 string }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64; F1 string }

func main() {
    var v1 uint16 = 1
    var v2 bool = ((v1 + 761) != (^(v1)))
    var v3 uint64 = 2
    var v4 uint64 = (^(v3))
    var v5 int8 = -128
    switch -2147483648 {
    case 127:
        for fi6 := 0; fi6 < 13; fi6++ {
            for fi7 := 0; fi7 < 13; fi7++ {
                if v2 {
                    switch (fi6 & (fi6 & 342)) {
                    case 304:
                        switch "a" {
                        case "hello":
                            fmt.Println(v5, v1, v1)
                        case "a":
                            if true {
                                var v8 uint = 4294967295
                                _ = v8
                            }
                        case "\n\t":
                            v1 = (v1 ^ v1)
                        default:
                            if (127 >= -121) {
                                v9 := GenS1{F0: uint8(v1), F1: ((v2 && false) && (42 != 32767))}
                                v5 = v5
                                _ = v9
                            }
                        }
                    case 4294967295:
                        if (v2 && (false || v2)) {
                            v10 := GenS2{F0: 0, F1: 1, F2: ""}
                            _ = v10
                        }
                    default:
                        var v11 uint = 1
                        var v12 uint8 = 128
                        _ = v11
                        _ = v12
                    }
                    var v13 int64 = -9223372036854775808
                    _ = v13
                }
                switch (fi7 * (fi6 & 128)) {
                case -1:
                    v5 += v5
                case 922:
                    v14 := []int{9223372036854775807}
                    fmt.Println(v5)
                    _ = v14
                default:
                    var v15 uint = 806
                    _ = v15
                }
                _ = fi7
            }
            _ = fi6
        }
    case 968:
        var v16 int16 = 1
        _ = v16
    }
    v4 = v3
    v4 += (v4 ^ (v3 - (^(v4))))
    if false {
        var v17 int64 = 0
        _ = v17
    }
    if (!false) {
        v5 = (v5 - (v5 + 127))
    }
    if (!v2) {
        switch 588 {
        case 0:
            for fi18 := 0; fi18 < 8; fi18++ {
                v19 := GenS1{F0: 128, F1: v2}
                var v20 uint32 = uint32(v1)
                v21 := GenS2{F0: 2147483647, F1: 311, F2: ("foo bar" + "a")}
                _ = fi18
                _ = v19
                _ = v20
                _ = v21
            }
        default:
            if (false && (!((v3 * (v3 / ((v4) | 1))) == v3))) {
                for fi22 := 0; fi22 < 3; fi22++ {
                    v23 := GenS3{F0: -473}
                    _ = fi22
                    _ = v23
                }
            }
        }
    }
    for fi24 := 0; fi24 < 11; fi24++ {
        v3 += 4294967295
        fmt.Println(v3, v1)
        v25 := []int{}
        _ = fi24
        _ = v25
    }
    v26 := GenS4{F0: -475, F1: "a"}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
