package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint16 }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 int; F1 int64; F2 int32 }
type GenS4 struct { F0 uint64; F1 uint16 }
type GenS5 struct { F0 string; F1 int64 }

func main() {
    var v1 int = 4294967295
    var v2 bool = false
    var v3 int8 = 42
    var v4 int32 = 0
    var v5 uint32 = 1
    if true {
        if false {
            v6 := GenS1{F0: 42, F1: v4, F2: 418}
            _ = v6
        } else {
            v7 := GenS2{F0: v3, F1: (v1 & v1)}
            _ = v7
        }
        v8 := []int{v1, (v1 | (v1 & 128)), v1, (v1 % (((v1 ^ (v1 >> ((423) & 63)))) | 1))}
        _ = v8
    } else {
        var v9 float32 = -2.25
        for fi10 := 0; fi10 < 3; fi10++ {
            v11 := GenS3{F0: (fi10 ^ (v1 << ((783) & 63))), F1: 1, F2: v4}
            v12 := []int{}
            if true {
                v13 := GenS4{F0: 18446744073709551615, F1: 619}
                _ = v13
            }
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = v9
    }
    switch 32768 {
    case 4294967295:
        var v14 int = 2147483647
        _ = v14
    }
    switch 9223372036854775807 {
    case 127:
        if true {
            fmt.Println(v2, v2, v2)
            var v15 int8 = (^(v3))
            var v16 int64 = -152
            _ = v15
            _ = v16
        } else {
            var v17 string = ""
            _ = v17
        }
    case 409:
        v18 := GenS3{F0: v1, F1: -9223372036854775808, F2: (v4 / ((144) | 1))}
        _ = v18
    case -9223372036854775808:
        if (v4 < 221) {
            var v19 int16 = 32767
            _ = v19
        }
    default:
        for fi20 := 0; fi20 < 10; fi20++ {
            v21 := []int{}
            _ = fi20
            _ = v21
        }
    }
    v22 := []int{}
    for fi23 := 0; fi23 < 20; fi23++ {
        v4 = (v4 << ((65535) & 63))
        _ = fi23
    }
    if v2 {
        var v24 int8 = v3
        if (v1 <= 9223372036854775807) {
            v2 = (2 > 933)
            v4 = (v4 ^ v4)
            if (!((((((!(((("" + "\n\t") != "foo bar") && (((-2.25 * 0.0) != 0.0) && false)) && (712 > 0))) && false) && v2) || (4294967295 >= 1)) && (v1 == 32768)) && ((v3 ^ (^(v24))) > v24))) {
                var v25 int8 = v24
                _ = v25
            } else {
                v26 := GenS1{F0: 1, F1: (v4 / ((v4) | 1)), F2: 948}
                _ = v26
            }
        } else {
            var v27 int = v1
            _ = v27
        }
        _ = v24
    } else {
        if v2 {
            _ = len(v22)
            if len(v22) >= 2 {
                v22 = v22[1:len(v22)]
            }
            v2 = ((-(v3)) == v3)
        } else {
            for fi28 := 0; fi28 < 4; fi28++ {
                v4 = -1
                var v29 int32 = -1
                _ = fi28
                _ = v29
            }
        }
    }
    v4 += v4
    switch ((("a" + "\n\t") + "foo bar") + "foo bar") {
    case "\n\t":
        v3 += v3
    }
    var v30 uint32 = (v5 - 42)
    for fi31 := 0; fi31 < 18; fi31++ {
        v1 = (fi31 + (-(fi31)))
        fmt.Println(v1, v30, v2)
        v32 := GenS5{F0: ("a" + ("foo bar" + "foo bar")), F1: 9223372036854775807}
        _ = fi31
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
