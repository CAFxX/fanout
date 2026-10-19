package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 int32 }
func genFn0(p0 int8, p1 bool) int8 {
    for fi5 := 0; fi5 < 8; fi5++ {
        var v6 bool = true
        _ = fi5
        _ = v6
    }
    if p1 {
        var v7 uint64 = 608
        v7 += (^(v7))
        var v8 uint16 = 65535
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 20; fi9++ {
        switch 32768 {
        case -607:
            for fi10 := 0; fi10 < 9; fi10++ {
                p1 = (65535 >= 255)
                v11 := GenS1{F0: 2, F1: 255}
                _ = fi10
                _ = v11
            }
        default:
            v12 := []int{}
            _ = v12
        }
        fmt.Println(p0, p1, fi9)
        p1 = (true && p1)
        _ = fi9
    }
    for fi13 := 0; fi13 < 7; fi13++ {
        p1 = p1
        _ = fi13
    }
    _ = p0
    _ = p1
    return (^(p0))
}

func main() {
    var v1 uint = 178
    v2 := []uint32{}
    var v3 int64 = 9223372036854775807
    var v4 uint32 = 27
    var v14 uint8 = 128
    v15 := GenS1{F0: 32767, F1: 255}
    var v16 int8 = 100
    if false {
        v16 -= v16
        if false {
            var v17 bool = false
            var v18 uint8 = v14
            v3 -= v3
            _ = v17
            _ = v18
        } else {
            v1 = 462
            v14 = (v14 >> (((v1 - (v1 << ((v1) & 63)))) & 63))
        }
        if (515 == 508) {
            for fi19 := 0; fi19 < 5; fi19++ {
                v20 := GenS1{F0: 42, F1: 0}
                v21 := GenS2{F0: -2147483648}
                _ = fi19
                _ = v20
                _ = v21
            }
            if ((^(v1)) != v1) {
                if false {
                    v22 := GenS2{F0: -1}
                    v3 = (v3 * v3)
                    _ = v22
                } else {
                    for fi23 := 0; fi23 < 10; fi23++ {
                        _ = len(v2)
                        v4 += 438
                        v14 += (v14 | v14)
                        _ = fi23
                    }
                    v16 = (v16 * v16)
                }
                var v24 uint8 = 255
                _ = v24
            }
        }
    } else {
        fmt.Println(v14, v4)
        v25 := GenS2{F0: 2147483647}
        _ = v25
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v14 -= v14
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if (-1 <= -225) {
        v1 = 85
    }
    if true {
        switch "foo bar" {
        case "":
            var v26 string = ("a" + "hello")
            _ = v26
        case "foo bar":
            var v27 uint8 = (^(v14))
            _ = v27
        case "\n\t":
            for fi28 := 0; fi28 < 11; fi28++ {
                v14 = uint8(v4)
                _ = fi28
            }
        }
    }
    fmt.Println(v4)
    v29 := GenS1{F0: 1, F1: 32767}
    v2 = append(v2, (v4 ^ (v4 % (((^(v4))) | 1))))
    _ = len(v2)
    if true {
        v3 -= v3
        fmt.Println(v1)
        v1 += v1
    } else {
        var v30 bool = false
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v29
    _ = v3
    _ = v4
}
