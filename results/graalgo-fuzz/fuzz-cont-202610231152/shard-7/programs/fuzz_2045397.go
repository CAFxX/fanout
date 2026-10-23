package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int; F1 uint32 }
func genFn0(p0 uint64) uint8 {
    p0 -= (^(p0))
    var v7 uint = 350
    var v8 int16 = -341
    _ = p0
    _ = v7
    _ = v8
    return 69
}
func genFn1(p0 uint8, p1 uint) uint8 {
    v9 := GenS1{F0: 32767, F1: p1}
    var v10 int8 = 100
    switch -276 {
    case 620:
        if (926 <= 18446744073709551615) {
            fmt.Println(v10, p1)
        } else {
            for fi11 := 0; fi11 < 17; fi11++ {
                v12 := GenS1{F0: 127, F1: p1}
                _ = fi11
                _ = v12
            }
        }
    case -87:
        v10 = v10
    default:
        var v13 uint32 = 0
        _ = v13
    }
    p1 = p1
    var v14 uint16 = 101
    _ = p0
    _ = p1
    _ = v10
    _ = v14
    _ = v9
    return (^(p0))
}

func main() {
    var v1 int8 = 95
    var v2 uint64 = 168
    var v3 string = ("\n\t" + "foo bar")
    var v4 uint64 = v2
    var v5 uint32 = 2
    v6 := make(map[string]bool)
    v5 += (v5 % ((v5) | 1))
    v3 = v3
    for fi15 := 0; fi15 < 18; fi15++ {
        v16 := GenS2{F0: true}
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 8; fi17++ {
        var v18 float32 = 0.0
        if false {
            v1 = (-(v1))
        }
        var v19 uint32 = v5
        _ = fi17
        _ = v18
        _ = v19
    }
    fmt.Println(v2, v3)
    v20 := v6[v3]
    for fi21 := 0; fi21 < 6; fi21++ {
        v20 = v20
        v22 := []int{fi21}
        for fi23 := 0; fi23 < 20; fi23++ {
            fmt.Println(v20)
            var v24 uint = 458
            _ = fi23
            _ = v24
        }
        _ = fi21
        _ = v22
    }
    for fi25 := 0; fi25 < 2; fi25++ {
        var v26 uint = 2
        v27 := []int{(fi25 ^ int(v26)), (fi25 & fi25), (-(fi25))}
        for fi28 := 0; fi28 < 16; fi28++ {
            fmt.Println(v3, v2)
            _ = fi28
        }
        _ = fi25
        _ = v26
        _ = v27
    }
    if true {
        var v29 uint16 = 255
        v29 = (v29 - v29)
        _ = v29
    } else {
        if false {
            fmt.Println(v4, v5, v5)
            var v30 int8 = (v1 + (v1 ^ (v1 * v1)))
            _ = v30
        } else {
            v4 = 882
            v31 := GenS3{F0: (v5 * v5)}
            _ = v31
        }
        fmt.Println(v3, v2, v1)
    }
    v20 = v20
    v2 -= v4
    var v32 bool = v20
    for fi33 := 0; fi33 < 2; fi33++ {
        v3 = (v3 + ((((v3 + (v3 + ("a" + "hello"))) + "hello") + ("" + "\n\t")) + (v3 + "\n\t")))
        _ = fi33
    }
    v34 := GenS4{F0: -2147483648, F1: (v5 / ((v5) | 1))}
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v32
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
