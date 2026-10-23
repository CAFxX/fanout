package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 uint32; F2 int16 }
type GenS3 struct { F0 uint64; F1 string; F2 int8 }
func genFn0(p0 uint16) uint32 {
    var v5 bool = true
    var v6 uint = 0
    v7 := GenS1{F0: "a"}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 616
}

func main() {
    var v1 float32 = 0.0
    var v2 uint16 = 32767
    var v3 float64 = 1.5
    var v4 uint8 = 103
    var v8 uint16 = v2
    if false {
        switch "\n\t" {
        case "hello":
            var v9 uint = 866
            _ = v9
        case "foo bar":
            if true {
                v10 := GenS2{F0: 1, F1: 843, F2: 42}
                _ = v10
            } else {
                if false {
                    var v11 int64 = 9223372036854775807
                    _ = v11
                }
            }
        }
        var v12 int32 = -2147483648
        var v13 int64 = 42
        _ = v12
        _ = v13
    }
    if (false || ((false && true) || false)) {
        fmt.Println(v3, v4, v2)
    } else {
        for fi14 := 0; fi14 < 16; fi14++ {
            fmt.Println(fi14, v1, v4)
            v2 += v8
            _ = fi14
        }
    }
    v1 -= v1
    for fi15 := 0; fi15 < 20; fi15++ {
        v8 -= 0
        _ = fi15
    }
    for fi16 := 0; fi16 < 1; fi16++ {
        if (v8 < (v2 / (((v8 & (v8 & v8))) | 1))) {
            fmt.Println(v8, fi16)
        }
        v2 -= (v8 ^ (v8 + (v2 / ((v8) | 1))))
        var v17 bool = false
        _ = fi16
        _ = v17
    }
    if false {
        fmt.Println(v1, v2)
    } else {
        if true {
            for fi18 := 0; fi18 < 11; fi18++ {
                switch "\n\t" {
                case "foo bar":
                    if ((fi18 * (-(fi18))) <= fi18) {
                        var v19 int = -9223372036854775808
                        _ = v19
                    }
                case "a":
                    var v20 int64 = -824
                    _ = v20
                case "hello":
                    var v21 bool = true
                    _ = v21
                }
                v22 := GenS3{F0: 42, F1: "", F2: 1}
                _ = fi18
                _ = v22
            }
            fmt.Println(v3, v8)
        }
        switch "hello" {
        case "\n\t":
            var v23 int8 = -128
            _ = v23
        case "foo bar":
            fmt.Println(v3, v4)
        }
    }
    for fi24 := 0; fi24 < 14; fi24++ {
        var v25 uint64 = 42
        _ = fi24
        _ = v25
    }
    var v26 uint8 = (v4 & v4)
    fmt.Println(v2)
    switch 32768 {
    case -890:
        var v27 int32 = -2147483648
        _ = v27
    }
    v28 := GenS1{F0: ("" + ("" + ""))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v26
    _ = v28
    _ = v3
    _ = v4
    _ = v8
}
