package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8; F1 int8 }
type GenS4 struct { F0 int16; F1 int }
type GenS5 struct { F0 uint32; F1 int8; F2 uint64 }
func genFn0(p0 int, p1 int8) uint16 {
    if (false && ((!(2 >= 128)) || false)) {
        v4 := GenS1{F0: 255}
        if true {
            switch "\n\t" {
            case "a":
                fmt.Println(p1, p1)
            case "":
                p1 += (p1 >> ((615) & 63))
            default:
                p1 -= p1
            }
        }
        var v5 string = ((("foo bar" + "hello") + "\n\t") + ("hello" + "a"))
        _ = v4
        _ = v5
    } else {
        if true {
            for fi6 := 0; fi6 < 2; fi6++ {
                fmt.Println(p0, p0, fi6)
                if true {
                    p0 -= p0
                    var v7 int32 = 1
                    v8 := []int{}
                    _ = v7
                    _ = v8
                }
                var v9 uint = 1
                _ = fi6
                _ = v9
            }
        } else {
            fmt.Println(p1)
            if (("\n\t" <= "a") && (false && false)) {
                p0 += (p0 >> ((65535) & 63))
            }
            var v10 uint64 = 42
            _ = v10
        }
        for fi11 := 0; fi11 < 2; fi11++ {
            if (2 > 42) {
                if ("foo bar" >= ("foo bar" + ("hello" + "foo bar"))) {
                    p1 = p1
                    fmt.Println(p0)
                    fmt.Println(p0)
                } else {
                    if false {
                        fmt.Println(p1, p0, fi11)
                    } else {
                        var v12 int16 = 1000
                        var v13 float64 = 0.0
                        v14 := []int{(fi11 ^ (p0 << ((18446744073709551615) & 63))), fi11}
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                    p1 -= (p1 % (((p1 + p1)) | 1))
                }
                for fi15 := 0; fi15 < 15; fi15++ {
                    v16 := GenS2{F0: 42}
                    p1 -= (p1 % ((p1) | 1))
                    _ = fi15
                    _ = v16
                }
            }
            _ = fi11
        }
    }
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    return 2
}
func genFn1(p0 uint16, p1 int8, p2 int) uint8 {
    v17 := GenS2{F0: 0}
    fmt.Println(p0)
    p0 = (p0 % (((p0 - (p0 + (p0 >> ((0) & 63))))) | 1))
    if (65535 < 1) {
        p2 += p2
        v18 := GenS3{F0: p1, F1: 0}
        v19 := GenS2{F0: 65535}
        _ = v18
        _ = v19
    } else {
        p1 += (p1 * p1)
        v20 := GenS4{F0: -1, F1: int(p0)}
        _ = v20
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return 0
}

func main() {
    var v1 int64 = -883
    v2 := make(map[string]string)
    var v3 int = 42
    delete(v2, "foo bar")
    var v21 uint16 = 2
    var v22 float64 = -2.25
    if (((32767 < 1000) || false) && false) {
        var v23 bool = (true && true)
        v21 = v21
        _ = v23
    }
    v22 -= v22
    fmt.Println(v22)
    v24 := GenS1{F0: 949}
    var v25 uint64 = 96
    v21 = v21
    fmt.Println(v1, v21, v21)
    switch 0 {
    case 0:
        v26 := GenS1{F0: 740}
        _ = v26
    default:
        v21 = (^(v21))
    }
    switch (^(v3)) {
    case 128:
        if true {
            v27 := GenS5{F0: 4294967295, F1: 42, F2: v25}
            v28 := v2["foo bar"]
            _ = v27
            _ = v28
        }
    case 42:
        v22 += v22
    default:
        for fi29 := 0; fi29 < 13; fi29++ {
            v1 -= v1
            fmt.Println(v21, v22, v21)
            _ = fi29
        }
    }
    v30 := v2["foo bar"]
    switch "a" {
    case "foo bar":
        for fi31 := 0; fi31 < 12; fi31++ {
            var v32 int8 = 1
            v33 := GenS2{F0: 18446744073709551615}
            _ = fi31
            _ = v32
            _ = v33
        }
    case "":
        v34 := GenS5{F0: 2, F1: 1, F2: (v25 * v25)}
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v24
    _ = v25
    _ = v3
    _ = v30
}
