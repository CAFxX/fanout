package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 int8; F1 int8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint16 }
func genFn0(p0 int64, p1 int8, p2 int32) uint32 {
    p2 = p2
    v6 := GenS1{F0: 255}
    v7 := GenS1{F0: 32767}
    p2 -= int32(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 uint32, p1 int, p2 int) uint {
    for fi8 := 0; fi8 < 2; fi8++ {
        switch p1 {
        case -276:
            var v9 uint16 = 65535
            _ = v9
        }
        switch fi8 {
        case 32767:
            fmt.Println(p1, p0)
        case -1:
            fmt.Println(p2, p1)
        }
        _ = fi8
    }
    v10 := []int{p2}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 683
}

func main() {
    var v1 int16 = 1000
    var v2 uint16 = 591
    var v3 uint = 18446744073709551615
    var v4 uint32 = 0
    var v5 int = 127
    v11 := []int{}
    v2 = (v2 % ((v2) | 1))
    v12 := GenS2{F0: -21}
    for fi13 := 0; fi13 < 12; fi13++ {
        if len(v11) > 0 {
            v11[0] = 4294967295
        }
        fmt.Println(fi13, v3, v2)
        v14 := GenS3{F0: 135, F1: true}
        _ = fi13
        _ = v14
    }
    if false {
        fmt.Println(v3, v2, v5)
    } else {
        for fi15 := 0; fi15 < 9; fi15++ {
            v16 := GenS4{F0: 25, F1: 1}
            _ = fi15
            _ = v16
        }
    }
    if (((0.1 - 1e10) > (3.14159 - 0.1)) && false) {
        if ((!(9223372036854775807 <= -252)) && ((793 > v1) || (1 >= 0))) {
            if len(v11) > 0 {
                v11[0] = 42
            }
        }
    }
    var v17 uint32 = v4
    switch (("hello" + (("\n\t" + ((((("hello" + "hello") + "") + "\n\t") + "") + "a")) + "")) + "\n\t") {
    case "":
        switch "hello" {
        case "foo bar":
            var v18 bool = ((v2 << (((^(v3))) & 63)) != 255)
            _ = v18
        case "":
            var v19 uint32 = (v4 >> (((v3 >> (((v3 % ((v3) | 1))) & 63))) & 63))
            _ = v19
        case "a":
            var v20 int32 = 0
            _ = v20
        }
    case "hello":
        var v21 uint8 = uint8(v3)
        _ = v21
    default:
        v2 -= (^(v2))
    }
    if len(v11) > 0 {
        v11[0] = (v5 * v5)
    }
    var v22 uint64 = 534
    var v23 float32 = -2.25
    for fi24 := 0; fi24 < 17; fi24++ {
        var v25 int16 = (v1 / ((v1) | 1))
        for fi26 := 0; fi26 < 1; fi26++ {
            v27 := GenS5{F0: 32767}
            var v28 int = fi24
            switch ("\n\t" + "\n\t") {
            case "a":
                switch (fi26 & fi24) {
                case -9223372036854775808:
                    v29 := GenS6{F0: v2}
                    _ = v29
                case 4294967295:
                    var v30 uint32 = (v4 - genFn0(-446, 58, -1))
                    _ = v30
                }
            }
            _ = fi26
            _ = v27
            _ = v28
        }
        _ = fi24
        _ = v25
    }
    for fi31 := 0; fi31 < 2; fi31++ {
        v4 -= v4
        if true {
            var v32 int = (v5 ^ (v5 % ((-644) | 1)))
            var v33 uint8 = 2
            _ = v32
            _ = v33
        } else {
            fmt.Println(v17, fi31)
            var v34 uint = (^(v3))
            _ = v34
        }
        if false {
            v22 = 4294967295
        }
        _ = fi31
    }
    fmt.Println(v17, v3, v22)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
