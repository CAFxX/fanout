package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 uint32; F2 int16 }
type GenS4 struct { F0 int8 }

func main() {
    v1 := []uint32{}
    var v2 uint = 188
    var v3 uint8 = 0
    var v4 uint16 = 342
    if (0 <= uint64(v3)) {
        for fi5 := 0; fi5 < 8; fi5++ {
            for fi6 := 0; fi6 < 11; fi6++ {
                v1 = append(v1, 703)
                var v7 int = 128
                _ = fi6
                _ = v7
            }
            v3 = (v3 << ((v2) & 63))
            _ = fi5
        }
        v2 += (v2 ^ v2)
    } else {
        if len(v1) > 0 {
            v1[0] = 65535
        }
        fmt.Println(v2, v4, v3)
        var v8 uint64 = 73
        _ = v8
    }
    if true {
        v9 := GenS1{F0: 42, F1: 413, F2: -1}
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v9
    } else {
        for fi10 := 0; fi10 < 20; fi10++ {
            _ = len(v1)
            fmt.Println(v3, fi10, v4)
            fmt.Println(v4, v2)
            _ = fi10
        }
        v11 := GenS2{F0: 1}
        v3 = v3
        _ = v11
    }
    fmt.Println(v4, v2)
    if false {
        var v12 int8 = 100
        for fi13 := 0; fi13 < 14; fi13++ {
            for fi14 := 0; fi14 < 14; fi14++ {
                if false {
                    fmt.Println(fi13, v2)
                    v12 = v12
                }
                _ = fi14
            }
            if (false && (fi13 > fi13)) {
                var v15 uint = 0
                v16 := GenS3{F0: "foo bar", F1: 2147483648, F2: 1000}
                _ = v15
                _ = v16
            } else {
                v12 = (v12 * (v12 * -128))
                var v17 uint32 = 230
                v1 = append(v1, 2147483647)
                _ = v17
            }
            _ = len(v1)
            _ = fi13
        }
        var v18 int8 = v12
        _ = v12
        _ = v18
    } else {
        switch 42 {
        case -358:
            var v19 bool = (1.5 <= (1e10 / 1.0))
            _ = v19
        default:
            if len(v1) > 0 {
                v1[0] = 2
            }
        }
    }
    for fi20 := 0; fi20 < 19; fi20++ {
        v3 += v3
        _ = fi20
    }
    v21 := GenS2{F0: 1}
    var v22 float64 = 1e10
    switch -273 {
    case -9223372036854775808:
        v23 := GenS1{F0: -1, F1: 2147483647, F2: 2147483647}
        _ = v23
    case -2147483648:
        v3 += uint8(v2)
    case 4294967295:
        if (-9223372036854775808 > 128) {
            v24 := GenS4{F0: 0}
            switch "a" {
            case "\n\t":
                _ = len(v1)
            case "foo bar":
                if len(v1) > 0 {
                    v1[0] = 0
                }
            case "hello":
                v25 := GenS4{F0: 100}
                _ = v25
            }
            _ = v24
        } else {
            var v26 bool = true
            _ = v26
        }
    }
    v27 := GenS2{F0: 32767}
    var v28 int16 = 32767
    var v29 int16 = v28
    switch "foo bar" {
    case "a":
        var v30 int = -9223372036854775808
        _ = v30
    default:
        if true {
            var v31 uint64 = 18446744073709551615
            fmt.Println(v22)
            v31 += (v31 - (v31 * (v31 & v31)))
            _ = v31
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
