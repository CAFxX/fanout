package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64; F1 int8; F2 uint8 }

func main() {
    var v1 float64 = 0.1
    v2 := []int{780}
    v3 := []int{-538, 128, 32768}
    var v4 uint16 = 255
    var v5 uint64 = 42
    var v6 uint16 = v4
    if true {
        for fi7 := 0; fi7 < 17; fi7++ {
            switch "hello" {
            case "foo bar":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case "\n\t":
                v1 = v1
            case "a":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            default:
                for fi8 := 0; fi8 < 5; fi8++ {
                    fmt.Println(v4, fi8, v6)
                    v9 := GenS1{F0: -1}
                    _ = fi8
                    _ = v9
                }
            }
            fmt.Println(v4, v1, v1)
            _ = fi7
        }
        var v10 int32 = 0
        _ = v10
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi11 := 0; fi11 < 3; fi11++ {
        for fi12 := 0; fi12 < 13; fi12++ {
            if (v5 != v5) {
                fmt.Println(fi12)
                v13 := GenS2{F0: "foo bar"}
                _ = v13
            } else {
                var v14 int64 = -9223372036854775808
                var v15 int64 = (v14 / ((v14) | 1))
                if false {
                    fmt.Println(fi12)
                }
                var v16 uint8 = 50
                _ = v14
                _ = v15
                _ = v16
            }
            v6 += v4
            _ = fi12
        }
        v6 -= (^(v6))
        _ = fi11
    }
    v4 = (v6 | (^(v4)))
    var v17 int32 = -1
    if true {
        var v18 string = ""
        v3 = append(v3, 0)
        _ = v18
    }
    v4 += v4
    for fi19 := 0; fi19 < 16; fi19++ {
        for fi20 := 0; fi20 < 14; fi20++ {
            v5 = 4294967295
            var v21 int = fi19
            if (uint(v6) < 281) {
                if false {
                    v6 -= v6
                    v22 := GenS1{F0: -9223372036854775808}
                    fmt.Println(v21, v17)
                    _ = v22
                } else {
                    fmt.Println(v21, fi19, v6)
                    fmt.Println(v4, v17)
                    v5 -= v5
                    var v23 uint32 = 271
                    _ = v23
                }
            }
            v2 = append(v2, (v21 << ((42) & 63)))
            _ = fi20
            _ = v21
        }
        _ = fi19
    }
    if true {
        switch "foo bar" {
        case "\n\t":
            var v24 int64 = 42
            _ = v24
        case "hello":
            fmt.Println(v1)
        case "":
            var v25 int32 = (v17 % ((969) | 1))
            _ = v25
        }
        v26 := GenS1{F0: int64(v6)}
        for fi27 := 0; fi27 < 6; fi27++ {
            v28 := GenS3{F0: (v5 - 4294967295), F1: 42, F2: 128}
            _ = fi27
            _ = v28
        }
        _ = v26
    }
    v29 := GenS3{F0: (v5 + (v5 | v5)), F1: -1, F2: 2}
    v5 -= 451
    var v30 string = "a"
    fmt.Println(v4, v4, v1)
    v6 = 42
    var v31 uint64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
