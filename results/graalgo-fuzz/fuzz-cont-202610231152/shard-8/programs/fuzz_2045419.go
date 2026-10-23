package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 uint64; F1 string; F2 uint16 }
type GenS4 struct { F0 int16; F1 string }
func genFn0(p0 uint64, p1 uint32, p2 uint32) uint8 {
    v5 := []int{int(p0), -575, 42}
    var v6 uint8 = 255
    v6 = v6
    for fi7 := 0; fi7 < 10; fi7++ {
        for fi8 := 0; fi8 < 4; fi8++ {
            fmt.Println(v6, fi8)
            v9 := GenS1{F0: (p2 / (((p2 * (p2 - p1))) | 1))}
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        var v11 bool = (!(!(false || false)))
        switch -2147483648 {
        case -2147483648:
            var v12 uint32 = p1
            _ = v12
        case -1:
            if len(v5) > 0 {
                v5[0] = 9223372036854775807
            }
        default:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        if (!(p1 > (p2 ^ 851))) {
            v13 := GenS1{F0: (^(p1))}
            fmt.Println(v6, fi10)
            p1 = uint32(p0)
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return v6
}
func genFn1(p0 uint16, p1 uint) uint64 {
    var v14 int8 = 42
    switch "foo bar" {
    case "\n\t":
        v15 := GenS2{F0: 0, F1: false}
        _ = v15
    case "a":
        for fi16 := 0; fi16 < 6; fi16++ {
            var v17 uint64 = 315
            for fi18 := 0; fi18 < 17; fi18++ {
                v19 := []int{}
                fmt.Println(p1)
                _ = fi18
                _ = v19
            }
            if false {
                v20 := []int{(fi16 << (((^(p1))) & 63))}
                v21 := GenS3{F0: (^(v17)), F1: "\n\t", F2: p0}
                _ = v20
                _ = v21
            } else {
                var v22 string = (("a" + "foo bar") + ("" + "foo bar"))
                _ = v22
            }
            _ = fi16
            _ = v17
        }
    }
    _ = p0
    _ = p1
    _ = v14
    return 18446744073709551615
}

func main() {
    var v1 int = 9223372036854775807
    var v2 uint16 = 469
    var v3 uint = 255
    var v4 uint64 = 2
    v23 := []int{int(v3)}
    var v24 uint = v3
    v25 := GenS4{F0: 0, F1: "foo bar"}
    _ = len(v23)
    _ = len(v23)
    var v26 bool = false
    switch (v1 >> ((v24) & 63)) {
    case 4294967295:
        _ = len(v23)
    case 127:
        fmt.Println(v2, v2)
    }
    fmt.Println(v3)
    for fi27 := 0; fi27 < 4; fi27++ {
        fmt.Println(v4, v2)
        var v28 int = 9223372036854775807
        for fi29 := 0; fi29 < 3; fi29++ {
            v28 = fi27
            var v30 int64 = -1
            fmt.Println(v30)
            _ = fi29
            _ = v30
        }
        _ = fi27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
