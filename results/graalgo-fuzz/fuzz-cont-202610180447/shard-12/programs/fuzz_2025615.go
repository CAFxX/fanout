package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 int8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64; F1 int }
func genFn0() int16 {
    var v6 float64 = -2.25
    var v7 int64 = -1
    _ = v6
    _ = v7
    return -275
}
func genFn1() uint16 {
    switch -438 {
    case -6:
        var v8 string = "hello"
        _ = v8
    case 128:
        var v9 float32 = -2.25
        _ = v9
    case -2147483648:
        var v10 uint8 = 255
        _ = v10
    default:
        v11 := GenS1{F0: 32767, F1: 2147483647, F2: 42}
        _ = v11
    }
    if (1.5 >= 0.1) {
        var v12 int32 = 1
        _ = v12
    } else {
        v13 := GenS2{F0: 1}
        var v14 float32 = 3.14159
        var v15 int32 = 0
        _ = v13
        _ = v14
        _ = v15
    }
    var v16 bool = false
    v16 = v16
    _ = v16
    return 903
}

func main() {
    var v1 int64 = -1
    var v2 int16 = 0
    var v3 int8 = 124
    var v4 int32 = -32
    v5 := []string{("foo bar" + "hello"), ""}
    if (-719 < (v4 * v4)) {
        var v17 bool = (!false)
        _ = v17
    } else {
        v3 = (-(v3))
        v2 -= v2
        for fi18 := 0; fi18 < 4; fi18++ {
            var v19 float32 = 0.0
            fmt.Println(v2, v1, v3)
            switch "\n\t" {
            case "":
                v20 := GenS3{F0: 971}
                _ = v20
            case "\n\t":
                v1 -= -1
            case "hello":
                fmt.Println(v1, fi18, v4)
            }
            _ = fi18
            _ = v19
        }
    }
    switch "hello" {
    case "hello":
        v3 = (v3 << ((1) & 63))
    case "\n\t":
        if false {
            var v21 int = int(v1)
            _ = v21
        }
    }
    switch ("hello" + "a") {
    case "\n\t":
        v22 := GenS2{F0: 2}
        _ = v22
    case "":
        v4 += (v4 | 2147483647)
    case "a":
        for fi23 := 0; fi23 < 6; fi23++ {
            var v24 float64 = -2.25
            if (3.14159 > (v24 + v24)) {
                fmt.Println(v2)
                v25 := GenS4{F0: 4294967295, F1: fi23}
                for fi26 := 0; fi26 < 4; fi26++ {
                    v4 = (v4 << ((587) & 63))
                    v1 -= -1
                    _ = fi26
                }
                _ = v25
            }
            _ = fi23
            _ = v24
        }
    }
    fmt.Println(v1, v2, v3)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v5 = append(v5, ("a" + ""))
    v1 -= (^(v1))
    switch "foo bar" {
    case "a":
        v2 = int16(v4)
    default:
        v1 = v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
