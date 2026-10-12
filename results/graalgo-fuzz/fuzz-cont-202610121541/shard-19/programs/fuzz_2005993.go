package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int16; F1 int32 }

func main() {
    var v1 uint64 = 0
    var v2 int32 = -942
    var v3 bool = false
    var v4 int64 = 1
    v1 -= v1
    switch "\n\t" {
    case "foo bar":
        v5 := []int{}
        _ = v5
    case "a":
        fmt.Println(v3)
    case "\n\t":
        if ((v2 & (v2 ^ -522)) == v2) {
            v4 -= (v4 ^ 899)
        } else {
            fmt.Println(v4)
            v3 = (423 > 1)
            if (1.5 > 1e10) {
                var v6 uint8 = 201
                v7 := []int{2147483647, 937, 890, -522}
                _ = v6
                _ = v7
            } else {
                v3 = (42 == 42)
                var v8 uint16 = 1
                for fi9 := 0; fi9 < 6; fi9++ {
                    v10 := GenS1{F0: (v4 >> ((18446744073709551615) & 63)), F1: 940}
                    _ = fi9
                    _ = v10
                }
                _ = v8
            }
        }
    }
    switch 42 {
    case -9223372036854775808:
        v3 = false
    default:
        v11 := GenS2{F0: -472}
        _ = v11
    }
    var v12 uint64 = v1
    switch 722 {
    case 32768:
        var v13 uint32 = uint32(v1)
        _ = v13
    case -1:
        v14 := []int{}
        _ = v14
    case 128:
        for fi15 := 0; fi15 < 1; fi15++ {
            v16 := []int{-34, fi15, (fi15 ^ (fi15 % ((fi15) | 1))), (fi15 - (fi15 % ((9223372036854775807) | 1)))}
            if (true && v3) {
                if ((v3 && true) || (v3 || v3)) {
                    if len(v16) > 0 {
                        v16[0] = 686
                    }
                } else {
                    v2 -= (v2 / (((v2 ^ v2)) | 1))
                    var v17 uint16 = 32767
                    _ = v17
                }
                var v18 uint = 633
                _ = v18
            } else {
                var v19 uint16 = 32767
                var v20 int64 = v4
                for fi21 := 0; fi21 < 19; fi21++ {
                    v16 = append(v16, 137)
                    v22 := GenS1{F0: v20, F1: (v19 - 111)}
                    _ = fi21
                    _ = v22
                }
                _ = v19
                _ = v20
            }
            var v23 string = ("foo bar" + "foo bar")
            _ = fi15
            _ = v16
            _ = v23
        }
    default:
        v24 := []int{}
        _ = v24
    }
    v25 := []int{}
    if (("\n\t" + ((("a" + "a") + "a") + "\n\t")) > "hello") {
        if (255 != 616) {
            var v26 float64 = (1e10 / 1.5)
            switch "foo bar" {
            case "\n\t":
                fmt.Println(v12, v2)
            case "hello":
                var v27 uint8 = 42
                _ = v27
            case "foo bar":
                var v28 int = 894
                _ = v28
            default:
                v29 := GenS3{F0: (v4 & 9223372036854775807)}
                _ = v29
            }
            _ = v26
        } else {
            fmt.Println(v2, v1, v12)
            v4 = v4
            fmt.Println(v12, v12)
        }
        var v30 uint16 = 65535
        v31 := GenS4{F0: 42, F1: (v2 + v2)}
        _ = v30
        _ = v31
    }
    var v32 uint16 = 0
    var v33 float32 = 0.0
    if len(v25) > 0 {
        v25[0] = 4294967295
    }
    if (v2 <= (v2 | (v2 << ((1) & 63)))) {
        fmt.Println(v3, v1, v4)
    }
    v34 := GenS4{F0: 0, F1: (v2 >> ((553) & 63))}
    if false {
        fmt.Println(v2, v12)
        fmt.Println(v32)
        v33 = v33
    }
    for fi35 := 0; fi35 < 11; fi35++ {
        v32 += v32
        var v36 int64 = 42
        _ = fi35
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v25
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
}
