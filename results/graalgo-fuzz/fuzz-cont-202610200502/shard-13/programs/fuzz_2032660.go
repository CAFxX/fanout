package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 bool; F2 uint8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32; F1 string; F2 uint64 }
type GenS5 struct { F0 bool; F1 string }

func main() {
    var v1 int64 = 42
    var v2 int8 = 42
    var v3 uint = 4294967295
    v4 := []int{9223372036854775807}
    v1 += (v1 - (v1 << ((v3) & 63)))
    v5 := GenS1{F0: 128}
    var v6 int64 = (v1 % (((v1 / ((v1) | 1))) | 1))
    switch (("a" + ("a" + "")) + (("\n\t" + "") + "foo bar")) {
    case "foo bar":
        var v7 float64 = 1e10
        _ = v7
    case "\n\t":
        v8 := GenS2{F0: 128, F1: true, F2: 128}
        _ = v8
    case "hello":
        var v9 uint64 = 820
        _ = v9
    default:
        fmt.Println(v3, v6)
    }
    fmt.Println(v6, v2, v6)
    if len(v4) > 0 {
        v4[0] = -1
    }
    var v10 int16 = 678
    v2 = -1
    var v11 float32 = (0.1 - 3.14159)
    v11 += v11
    var v12 uint32 = 2147483647
    v1 = (v6 + -906)
    v13 := GenS3{F0: (-(v2))}
    switch -1 {
    case 127:
        if (((v12 >= 280) || false) || (v2 == (v2 % ((-19) | 1)))) {
            var v14 float64 = -2.25
            _ = v14
        } else {
            switch "hello" {
            case "a":
                var v15 uint16 = 2
                _ = v15
            case "\n\t":
                v16 := GenS2{F0: 51, F1: true, F2: 1}
                _ = v16
            }
            var v17 int16 = v10
            fmt.Println(v2)
            _ = v17
        }
    default:
        switch "" {
        case "hello":
            fmt.Println(v2, v12)
        case "foo bar":
            if (true || false) {
                v18 := GenS4{F0: 334, F1: "a", F2: 590}
                var v19 int64 = 379
                _ = v18
                _ = v19
            }
        case "a":
            switch 469 {
            case 32768:
                fmt.Println(v12)
            case 2147483647:
                switch "foo bar" {
                case "foo bar":
                    var v20 uint8 = 1
                    _ = v20
                }
            }
        default:
            fmt.Println(v10, v10, v11)
        }
    }
    var v21 bool = false
    v12 = 2147483648
    if false {
        for fi22 := 0; fi22 < 13; fi22++ {
            v23 := GenS5{F0: v21, F1: "foo bar"}
            v24 := GenS4{F0: -2147483648, F1: "hello", F2: 428}
            _ = fi22
            _ = v23
            _ = v24
        }
        v25 := GenS4{F0: -389, F1: "\n\t", F2: 1}
        _ = v25
    } else {
        if (1 < 1) {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        } else {
            var v26 int16 = v10
            for fi27 := 0; fi27 < 10; fi27++ {
                if true {
                    if v21 {
                        fmt.Println(v26, v12, v1)
                    }
                }
                switch (("a" + ("a" + ("a" + ""))) + ("" + "\n\t")) {
                case "\n\t":
                    v26 += -32768
                case "hello":
                    var v28 int32 = -1
                    _ = v28
                case "a":
                    var v29 uint32 = 1
                    _ = v29
                }
                _ = fi27
            }
            var v30 int = int(v12)
            _ = v26
            _ = v30
        }
        for fi31 := 0; fi31 < 2; fi31++ {
            v32 := GenS1{F0: 0}
            var v33 int = -9223372036854775808
            _ = fi31
            _ = v32
            _ = v33
        }
        fmt.Println(v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
