package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 int8 }
func genFn0() int64 {
    var v7 int8 = 100
    v8 := GenS1{F0: "\n\t"}
    fmt.Println(v7)
    _ = v7
    _ = v8
    return -31
}

func main() {
    v1 := []int{}
    var v2 int32 = 543
    v3 := make(map[uint32]string)
    var v4 bool = true
    var v5 uint32 = 42
    var v6 int32 = (v2 % ((v2) | 1))
    if (!true) {
        switch "hello" {
        case "\n\t":
            for fi9 := 0; fi9 < 14; fi9++ {
                v10 := GenS2{F0: (0 <= 1000), F1: 312}
                var v11 uint8 = 255
                _ = fi9
                _ = v10
                _ = v11
            }
        case "a":
            if len(v1) > 0 {
                v1[0] = 1
            }
        }
    }
    switch (("a" + "a") + "foo bar") {
    case "hello":
        var v12 uint64 = 1
        _ = v12
    case "":
        for fi13 := 0; fi13 < 18; fi13++ {
            v2 += v6
            _ = fi13
        }
    default:
        v14 := v3[(v5 & v5)]
        _ = v14
    }
    v5 = (v5 - 1)
    for fi15 := 0; fi15 < 10; fi15++ {
        switch (fi15 | (fi15 ^ fi15)) {
        case 639:
            v2 += 1
        }
        v16 := GenS3{F0: 1}
        v4 = (!false)
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 15; fi17++ {
        var v18 uint16 = 185
        var v19 uint16 = (v18 % ((65535) | 1))
        v4 = (v4 && true)
        _ = fi17
        _ = v18
        _ = v19
    }
    v5 += (v5 | 535)
    delete(v3, (^(v5)))
    for fi20 := 0; fi20 < 18; fi20++ {
        v4 = v4
        var v21 float64 = (1e10 / float64(0.1))
        _ = fi20
        _ = v21
    }
    var v22 uint16 = 309
    if true {
        switch "a" {
        case "a":
            var v23 uint = 255
            _ = v23
        case "foo bar":
            if v4 {
                v4 = (v4 && v4)
                v24 := GenS3{F0: 1}
                v2 += 2147483647
                _ = v24
            } else {
                var v25 uint = 2
                _ = v25
            }
        case "":
            var v26 float64 = 0.1
            _ = v26
        }
    } else {
        var v27 int8 = -1
        for fi28 := 0; fi28 < 11; fi28++ {
            if v4 {
                delete(v3, (v5 >> ((4294967295) & 63)))
            } else {
                var v29 int8 = -20
                var v30 bool = v4
                v27 -= 0
                _ = v29
                _ = v30
            }
            for fi31 := 0; fi31 < 19; fi31++ {
                switch fi31 {
                case 165:
                    var v32 int64 = -9223372036854775808
                    _ = v32
                case 42:
                    fmt.Println(v6)
                }
                _ = fi31
            }
            _ = fi28
        }
        _ = v27
    }
    v2 += (v2 * 42)
    var v33 uint64 = 677
    fmt.Println(v6, v4)
    v34 := GenS2{F0: ((v5 - 2) < 888), F1: (v22 << ((2) & 63))}
    v5 = v5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
