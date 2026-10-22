package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint; F1 int32 }

func main() {
    var v1 uint = 921
    var v2 uint64 = 1
    var v3 uint8 = 128
    var v4 float32 = (((-2.25 + 0.1) / 1.0) - (1e10 * ((0.0 / 0.1) - 0.1)))
    if (-2147483648 > 9223372036854775807) {
        switch "foo bar" {
        case "hello":
            var v5 int16 = -895
            _ = v5
        case "\n\t":
            switch ("\n\t" + "a") {
            case "hello":
                v3 = (v3 * (v3 ^ v3))
            case "\n\t":
                fmt.Println(v2, v4, v4)
            }
        case "":
            var v6 uint = uint(v3)
            _ = v6
        default:
            v7 := GenS1{F0: 1, F1: true}
            _ = v7
        }
    }
    switch ("hello" + ("a" + "hello")) {
    case "foo bar":
        for fi8 := 0; fi8 < 14; fi8++ {
            fmt.Println(v4, fi8)
            _ = fi8
        }
    case "":
        var v9 int64 = 0
        _ = v9
    }
    for fi10 := 0; fi10 < 14; fi10++ {
        v11 := GenS1{F0: 1, F1: (65535 >= (^(v1)))}
        v12 := GenS1{F0: 0, F1: true}
        if true {
            v13 := GenS2{F0: 2147483647}
            switch ("" + "\n\t") {
            case "":
                switch ("" + "") {
                case "":
                    for fi14 := 0; fi14 < 10; fi14++ {
                        v1 += (v1 / (((v1 & (v1 * (v1 + v1)))) | 1))
                        _ = fi14
                    }
                case "hello":
                    v15 := GenS3{F0: 45}
                    _ = v15
                default:
                    if (3.14159 != (0.1 * float64(3.14159))) {
                        v2 += v2
                        var v16 float64 = float64((v4 - 1.5))
                        _ = v16
                    } else {
                        var v17 uint8 = v3
                        _ = v17
                    }
                }
            }
            v18 := GenS3{F0: 382}
            _ = v13
            _ = v18
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    if ("hello" > "a") {
        v19 := GenS3{F0: -2147483648}
        if (!true) {
            var v20 int = -538
            if false {
                var v21 bool = (!false)
                _ = v21
            } else {
                v22 := []int{9223372036854775807}
                _ = v22
            }
            v23 := []int{v20}
            _ = v20
            _ = v23
        }
        v24 := []int{0, 2147483647, 1}
        _ = v19
        _ = v24
    }
    fmt.Println(v4, v3)
    v25 := []int{-2147483648}
    if true {
        if false {
            if (false && (!false)) {
                switch "hello" {
                case "a":
                    if len(v25) > 0 {
                        v25[0] = 2147483647
                    }
                case "foo bar":
                    var v26 uint32 = 65535
                    _ = v26
                }
                if (2147483647 < 0) {
                    v4 -= (v4 * v4)
                } else {
                    var v27 float64 = float64((v4 + v4))
                    if (!false) {
                        v28 := GenS4{F0: (v1 & 0), F1: 2147483647}
                        _ = v28
                    }
                    _ = v27
                }
            }
        }
        switch -2147483648 {
        case 766:
            if true {
                v2 -= (v2 - (v2 & v2))
                for fi29 := 0; fi29 < 2; fi29++ {
                    var v30 uint16 = 42
                    _ = fi29
                    _ = v30
                }
            }
        case 9223372036854775807:
            v3 = (v3 ^ (^(v3)))
        }
    } else {
        v1 = v1
    }
    var v31 uint32 = 42
    var v32 bool = false
    switch "a" {
    case "":
        for fi33 := 0; fi33 < 20; fi33++ {
            v25 = append(v25, fi33)
            for fi34 := 0; fi34 < 13; fi34++ {
                var v35 uint32 = (v31 | (v31 - 1))
                _ = fi34
                _ = v35
            }
            var v36 int16 = 1
            _ = fi33
            _ = v36
        }
    case "a":
        if len(v25) >= 2 {
            v25 = v25[1:len(v25)]
        }
    default:
        for fi37 := 0; fi37 < 6; fi37++ {
            fmt.Println(v31)
            _ = fi37
        }
    }
    fmt.Println(v32, v3, v31)
    switch 1 {
    case 4294967295:
        var v38 int16 = -583
        _ = v38
    default:
        if len(v25) > 0 {
            v25[0] = -977
        }
    }
    switch ("\n\t" + "") {
    case "hello":
        if len(v25) >= 2 {
            v25 = v25[1:len(v25)]
        }
    case "a":
        if len(v25) >= 2 {
            v25 = v25[1:len(v25)]
        }
    case "":
        var v39 uint32 = 2147483648
        _ = v39
    }
    fmt.Println(v4, v2, v1)
    var v40 int16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v40
}
