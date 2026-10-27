package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 uint; F1 uint32 }
type GenS4 struct { F0 uint32; F1 int64 }
func genFn0(p0 string, p1 int16, p2 bool) int16 {
    v6 := []int{-916, -9223372036854775808, 128, -951}
    switch "foo bar" {
    case "hello":
        p0 = p0
    case "":
        fmt.Println(p2)
    }
    switch p0 {
    case "":
        var v7 float64 = -2.25
        _ = v7
    }
    var v8 uint32 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v8
    return (p1 | (p1 / ((p1) | 1)))
}
func genFn1(p0 int64, p1 uint) int16 {
    v9 := []int{32767, -170, 675, -1}
    for fi10 := 0; fi10 < 10; fi10++ {
        p1 += (p1 - (p1 * p1))
        p0 += (p0 << ((p1) & 63))
        _ = fi10
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        switch int(p0) {
        case 128:
            if len(v9) > 0 {
                v9[0] = -1
            }
        default:
            var v12 int64 = (^(p0))
            _ = v12
        }
        _ = fi11
    }
    p0 += 0
    if (p0 > -1) {
        _ = len(v9)
    }
    _ = p0
    _ = p1
    _ = v9
    return -1
}

func main() {
    var v1 int64 = 983
    var v2 int8 = 111
    v3 := []uint32{65535, 530}
    var v4 uint64 = 18446744073709551615
    v5 := []uint32{2}
    var v13 float64 = (((1.5 - ((1e10 * -2.25) + 1.5)) - 1e10) + 0.0)
    for fi14 := 0; fi14 < 10; fi14++ {
        if true {
            var v15 int8 = (v2 / ((v2) | 1))
            v2 += v2
            _ = v15
        } else {
            v13 += v13
            for fi16 := 0; fi16 < 14; fi16++ {
                var v17 int64 = v1
                v18 := GenS1{F0: (v1 + v1)}
                _ = fi16
                _ = v17
                _ = v18
            }
        }
        _ = fi14
    }
    if true {
        v13 = (1e10 + v13)
        switch ("foo bar" + ("hello" + ("a" + "\n\t"))) {
        case "foo bar":
            v19 := GenS1{F0: v1}
            _ = v19
        case "":
            var v20 int16 = 1000
            _ = v20
        case "hello":
            var v21 int16 = 1000
            _ = v21
        default:
            var v22 int = -34
            _ = v22
        }
        v1 += v1
    } else {
        var v23 int32 = -788
        var v24 uint16 = 1
        _ = v23
        _ = v24
    }
    var v25 uint = 65535
    if true {
        v26 := GenS2{F0: (v4 | v4), F1: -2147483648}
        v27 := GenS3{F0: (v25 >> (((^(v25))) & 63)), F1: 4294967295}
        var v28 uint16 = 882
        _ = v26
        _ = v27
        _ = v28
    } else {
        for fi29 := 0; fi29 < 17; fi29++ {
            var v30 uint = (^(v25))
            switch fi29 {
            case 9223372036854775807:
                v1 += v1
            }
            _ = fi29
            _ = v30
        }
    }
    var v31 int8 = (^(v2))
    var v32 uint64 = v4
    v25 = v25
    for fi33 := 0; fi33 < 4; fi33++ {
        if (449 == 32767) {
            var v34 uint8 = 128
            if len(v5) > 0 {
                v5[0] = 1
            }
            _ = v34
        } else {
            v25 += uint(v32)
        }
        switch 0 {
        case 1:
            var v35 int16 = 1
            _ = v35
        case -267:
            if true {
                fmt.Println(v25, v13)
                var v36 int32 = 534
                var v37 int32 = v36
                _ = v36
                _ = v37
            }
        }
        _ = fi33
    }
    var v38 int32 = -1
    v31 += v31
    v39 := GenS4{F0: 65535, F1: 179}
    if len(v3) > 0 {
        v3[0] = 4294967295
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v25
    _ = v3
    _ = v31
    _ = v32
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
