package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int32; F1 uint; F2 int32 }
type GenS5 struct { F0 int }

func main() {
    v1 := []string{"\n\t", "hello"}
    var v2 float64 = 1.5
    var v3 bool = (1 == -22)
    var v4 uint16 = 282
    v5 := GenS1{F0: 127}
    var v6 float64 = v2
    v3 = v3
    v4 += (v4 & v4)
    var v7 uint = 0
    v7 = (v7 ^ 42)
    switch 4294967295 {
    case 128:
        var v8 uint16 = v4
        _ = v8
    case -312:
        v7 = v7
    case 281:
        for fi9 := 0; fi9 < 11; fi9++ {
            if v3 {
                for fi10 := 0; fi10 < 1; fi10++ {
                    var v11 int32 = 2147483647
                    _ = fi10
                    _ = v11
                }
            }
            _ = fi9
        }
    }
    if false {
        v12 := GenS2{F0: 1}
        switch ("foo bar" + "hello") {
        case "a":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        default:
            _ = len(v1)
        }
        v13 := GenS3{F0: 42}
        _ = v12
        _ = v13
    }
    if (!(float64(0.1) > v2)) {
        if false {
            var v14 int8 = 100
            v6 -= v2
            _ = v14
        } else {
            switch -9223372036854775808 {
            case 127:
                var v15 uint64 = 4294967295
                _ = v15
            case 0:
                for fi16 := 0; fi16 < 6; fi16++ {
                    var v17 int64 = 0
                    for fi18 := 0; fi18 < 10; fi18++ {
                        v19 := GenS4{F0: 2147483647, F1: v7, F2: int32(v7)}
                        _ = fi18
                        _ = v19
                    }
                    v17 = 9223372036854775807
                    _ = fi16
                    _ = v17
                }
            }
            v3 = true
        }
        var v20 bool = v3
        var v21 float32 = (((0.1 - ((float32((v6 + 0.0)) * 0.0) / 1e10)) - (((float32(1e10) * (1e10 + -2.25)) + float32(((((0.1 * v6) - v2) * float64(1e10)) * 0.1))) - ((1e10 / (0.0 + 0.1)) + 1.5))) + 0.0)
        _ = v20
        _ = v21
    } else {
        if v3 {
            var v22 float32 = 0.1
            if len(v1) > 0 {
                v1[0] = "a"
            }
            switch "foo bar" {
            case "\n\t":
                var v23 int16 = 778
                _ = v23
            case "":
                fmt.Println(v3)
            case "a":
                v24 := GenS4{F0: 0, F1: (v7 >> (((^(v7))) & 63)), F2: -158}
                _ = v24
            }
            _ = v22
        }
    }
    v7 = v7
    if len(v1) > 0 {
        v1[0] = ("hello" + "a")
    }
    for fi25 := 0; fi25 < 13; fi25++ {
        _ = len(v1)
        fmt.Println(v4)
        _ = fi25
    }
    var v26 uint16 = v4
    v27 := GenS5{F0: 32767}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
