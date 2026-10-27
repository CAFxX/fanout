package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int; F1 int64; F2 uint }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint64; F1 uint16 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint32; F1 int; F2 int32 }

func main() {
    var v1 uint16 = 0
    v2 := make(map[string]uint32)
    var v3 int64 = -907
    if (944 < 811) {
        for fi4 := 0; fi4 < 3; fi4++ {
            v3 -= -1
            _ = fi4
        }
        v3 += 873
        fmt.Println(v1)
    } else {
        for fi5 := 0; fi5 < 15; fi5++ {
            v3 = (v3 ^ v3)
            v6 := GenS1{F0: 283, F1: (!true)}
            var v7 int32 = 255
            _ = fi5
            _ = v6
            _ = v7
        }
        if false {
            v8 := v2[("hello" + "")]
            var v9 int32 = 1
            v3 -= 42
            _ = v8
            _ = v9
        }
        v1 = v1
    }
    switch ("hello" + "") {
    case "foo bar":
        fmt.Println(v3, v1)
    case "":
        for fi10 := 0; fi10 < 15; fi10++ {
            v11 := []int{(fi10 >> ((42) & 63))}
            _ = fi10
            _ = v11
        }
    case "a":
        var v12 int8 = 100
        _ = v12
    default:
        v1 -= (v1 ^ (v1 >> ((uint(v1)) & 63)))
    }
    v13 := GenS2{F0: v3}
    var v14 int8 = 1
    v15 := GenS3{F0: 0, F1: 976, F2: 255}
    var v16 uint64 = 748
    fmt.Println(v16, v14, v14)
    for fi17 := 0; fi17 < 12; fi17++ {
        v18 := GenS3{F0: fi17, F1: 42, F2: 1}
        var v19 uint16 = (^(v1))
        delete(v2, "foo bar")
        _ = fi17
        _ = v18
        _ = v19
    }
    switch 0 {
    case -1:
        v14 -= v14
    case 9223372036854775807:
        var v20 uint64 = v16
        _ = v20
    case -509:
        v1 = (v1 % ((v1) | 1))
    }
    v14 += (v14 % ((100) | 1))
    var v21 uint = 366
    var v22 float64 = 3.14159
    var v23 int64 = (-(v3))
    v24 := GenS4{F0: 32767}
    switch 1 {
    case 42:
        var v25 string = ""
        _ = v25
    case 128:
        if true {
            fmt.Println(v23, v23, v1)
            _ = len(v2)
            v26 := GenS5{F0: v16, F1: (v1 & v1)}
            _ = v26
        } else {
            v27 := GenS3{F0: 32768, F1: v3, F2: 65535}
            _ = v27
        }
    case -868:
        var v28 uint8 = 42
        _ = v28
    }
    for fi29 := 0; fi29 < 13; fi29++ {
        v30 := GenS6{F0: 1}
        for fi31 := 0; fi31 < 18; fi31++ {
            if (-1 <= 94) {
                v14 -= -1
                v16 = (v16 - v16)
            }
            fmt.Println(v16, v14, v22)
            var v32 int16 = 1000
            _ = fi31
            _ = v32
        }
        _ = fi29
        _ = v30
    }
    if false {
        v3 = -1
        v33 := GenS7{F0: 2147483648, F1: -642, F2: -979}
        v34 := GenS7{F0: 261, F1: -225, F2: 0}
        _ = v33
        _ = v34
    }
    var v35 int16 = -127
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v35
}
