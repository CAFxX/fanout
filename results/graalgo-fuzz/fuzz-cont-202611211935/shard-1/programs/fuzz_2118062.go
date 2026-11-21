package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int32 }
type GenS2 struct { F0 uint; F1 uint; F2 int64 }
type GenS3 struct { F0 string; F1 int8; F2 int16 }
func genFn0(p0 uint32, p1 int8, p2 uint16) uint16 {
    for fi7 := 0; fi7 < 11; fi7++ {
        fmt.Println(p0, p1, fi7)
        _ = fi7
    }
    for fi8 := 0; fi8 < 19; fi8++ {
        p1 = (-(p1))
        var v9 string = ""
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 9; fi10++ {
        for fi11 := 0; fi11 < 13; fi11++ {
            v12 := []int{fi11}
            _ = fi11
            _ = v12
        }
        p1 += p1
        var v13 float64 = 0.1
        _ = fi10
        _ = v13
    }
    p1 = p1
    if false {
        for fi14 := 0; fi14 < 14; fi14++ {
            fmt.Println(p1, p0, p0)
            p2 += p2
            v15 := GenS1{F0: true, F1: 65535, F2: 1}
            _ = fi14
            _ = v15
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return p2
}
func genFn1(p0 int32, p1 uint64, p2 bool) uint8 {
    if (2147483647 == 349) {
        if (p1 < p1) {
            v16 := []int{}
            _ = v16
        }
        v17 := []int{1, 9223372036854775807}
        _ = v17
    } else {
        var v18 int8 = 125
        var v19 bool = (!("hello" >= "foo bar"))
        _ = v18
        _ = v19
    }
    v20 := GenS2{F0: 255, F1: 255, F2: 42}
    switch -156 {
    case -600:
        v21 := GenS2{F0: 255, F1: 42, F2: 944}
        _ = v21
    case -890:
        v22 := GenS1{F0: (false || (((("a" + "a") + "") + "foo bar") < "hello")), F1: 1, F2: (p0 + p0)}
        _ = v22
    default:
        var v23 float32 = 0.0
        _ = v23
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    return 139
}

func main() {
    v1 := []uint32{}
    v2 := make(map[int]string)
    var v3 int8 = 0
    var v4 bool = ((!false) || false)
    var v5 uint32 = 2
    var v6 int32 = -1
    v6 = 0
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v5, v5, v5)
    v6 = v6
    var v24 float64 = 1.5
    v25 := GenS2{F0: 0, F1: 42, F2: int64(v5)}
    var v26 bool = v4
    var v27 string = ("foo bar" + "")
    switch v27 {
    case "":
        if ("" <= "a") {
            v24 += 0.0
        }
    case "a":
        _ = len(v2)
    }
    v5 = v5
    for fi28 := 0; fi28 < 10; fi28++ {
        for fi29 := 0; fi29 < 16; fi29++ {
            var v30 int32 = (^(v6))
            v31 := GenS2{F0: 1, F1: 2, F2: 9223372036854775807}
            _ = fi29
            _ = v30
            _ = v31
        }
        var v32 int64 = 9223372036854775807
        v2[fi28] = ("foo bar" + ("foo bar" + v27))
        _ = fi28
        _ = v32
    }
    v4 = v4
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v33 := GenS1{F0: true, F1: 13, F2: -2147483648}
    for fi34 := 0; fi34 < 17; fi34++ {
        if false {
            v35 := GenS3{F0: v27, F1: v3, F2: -707}
            v26 = v26
            _ = v35
        } else {
            fmt.Println(v24, v6, v3)
        }
        for fi36 := 0; fi36 < 15; fi36++ {
            v5 = 9
            _ = fi36
        }
        if v26 {
            switch (("\n\t" + (v27 + (v27 + v27))) + v27) {
            case "\n\t":
                if v4 {
                    _ = len(v1)
                    switch (v27 + v27) {
                    case "a":
                        delete(v2, 127)
                    case "":
                        v37 := v2[(fi34 - (fi34 % ((-2147483648) | 1)))]
                        _ = v37
                    case "\n\t":
                        v4 = (42 == 568)
                    default:
                        var v38 bool = ((v6 + 0) > v6)
                        _ = v38
                    }
                    var v39 uint64 = 4294967295
                    _ = v39
                }
            }
        } else {
            var v40 int8 = 42
            fmt.Println(v4, v6, v6)
            v41 := GenS3{F0: v27, F1: (v3 ^ 42), F2: 105}
            _ = v40
            _ = v41
        }
        _ = fi34
    }
    for fi42 := 0; fi42 < 17; fi42++ {
        delete(v2, fi42)
        _ = fi42
    }
    for fi43 := 0; fi43 < 8; fi43++ {
        if v4 {
            v5 += 954
            v6 -= 42
        } else {
            fmt.Println(v24)
            v24 = (v24 / v24)
            if true {
                var v44 uint8 = 1
                fmt.Println(v24, v5, v6)
                _ = v44
            }
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi43
    }
    v3 -= v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
