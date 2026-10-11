package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 int32; F1 int16 }
type GenS4 struct { F0 uint; F1 int32 }
func genFn0(p0 int) uint16 {
    v6 := GenS1{F0: (0 <= 0)}
    for fi7 := 0; fi7 < 6; fi7++ {
        p0 = fi7
        _ = fi7
    }
    var v8 uint8 = 2
    var v9 uint8 = (v8 * v8)
    _ = p0
    _ = v6
    _ = v8
    _ = v9
    return 32767
}
func genFn1(p0 uint, p1 uint64, p2 uint32) int64 {
    var v10 string = ("foo bar" + "")
    v11 := []int{}
    v12 := GenS2{F0: 1, F1: -128}
    p1 = (p1 % ((uint64(p0)) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    return -1
}

func main() {
    var v1 uint64 = 0
    v2 := []uint32{798, 687}
    var v3 int32 = -2147483648
    var v4 uint32 = 175
    v5 := []int{}
    _ = len(v2)
    v13 := GenS3{F0: v3, F1: int16(v4)}
    var v14 uint64 = v1
    for fi15 := 0; fi15 < 11; fi15++ {
        var v16 int32 = v3
        if len(v2) > 0 {
            v2[0] = 326
        }
        fmt.Println(v14, v3)
        _ = fi15
        _ = v16
    }
    if (!(!false)) {
        if false {
            v14 += (v1 + 589)
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            var v17 uint8 = 221
            _ = v17
        } else {
            var v18 int16 = -1
            _ = v18
        }
    } else {
        v14 = v1
        switch "hello" {
        case "foo bar":
            v19 := GenS1{F0: false}
            _ = v19
        case "":
            fmt.Println(v4)
        }
        v20 := GenS3{F0: v3, F1: -1}
        _ = v20
    }
    v3 = v3
    fmt.Println(v3)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v21 uint = 18446744073709551615
    if true {
        v22 := GenS4{F0: 4294967295, F1: (v3 ^ v3)}
        _ = v22
    }
    if false {
        switch 188 {
        case 32768:
            var v23 uint = v21
            _ = v23
        case 9223372036854775807:
            if false {
                v24 := GenS1{F0: true}
                _ = v24
            }
        }
    } else {
        var v25 uint = v21
        v21 = 0
        if true {
            v26 := GenS3{F0: v3, F1: 1}
            _ = v26
        } else {
            var v27 float64 = 3.14159
            switch 677 {
            case -557:
                var v28 uint64 = v14
                _ = v28
            }
            var v29 int64 = genFn1(v25, 4294967295, (v4 ^ (^(v4))))
            _ = v27
            _ = v29
        }
        _ = v25
    }
    var v30 uint32 = (v4 % ((65535) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v21
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
