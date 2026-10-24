package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 uint64; F1 uint32 }
func genFn0(p0 int32, p1 int32, p2 uint32) bool {
    p0 = p0
    if (127 != 32767) {
        if (-1 > 127) {
            fmt.Println(p2, p1, p1)
            for fi6 := 0; fi6 < 9; fi6++ {
                switch (("foo bar" + "hello") + "hello") {
                case "foo bar":
                    v7 := []int{(fi6 / ((4294967295) | 1)), fi6, (fi6 | (fi6 >> ((1) & 63))), fi6}
                    _ = v7
                default:
                    if false {
                        if (1 != (-(p0))) {
                            v8 := GenS1{F0: 426, F1: 1}
                            _ = v8
                        }
                    }
                }
                v9 := []int{fi6, (-(fi6))}
                _ = fi6
                _ = v9
            }
        } else {
            if false {
                var v10 int32 = 42
                p2 += p2
                _ = v10
            }
            p2 = p2
        }
        switch 9223372036854775807 {
        case 128:
            var v11 bool = true
            _ = v11
        }
        switch -2147483648 {
        case -71:
            p1 += (^(p0))
        case 9223372036854775807:
            for fi12 := 0; fi12 < 4; fi12++ {
                p0 = p1
                p2 = (p2 + 4294967295)
                v13 := GenS2{F0: 0}
                _ = fi12
                _ = v13
            }
        case 4294967295:
            v14 := GenS3{F0: 42, F1: -79}
            _ = v14
        }
    } else {
        v15 := []int{-1, -716}
        _ = v15
    }
    switch "hello" {
    case "\n\t":
        var v16 int64 = -9223372036854775808
        _ = v16
    }
    fmt.Println(p1, p2)
    p0 = (p1 % ((p0) | 1))
    _ = p0
    _ = p1
    _ = p2
    return false
}

func main() {
    var v1 uint = 4294967295
    var v2 bool = (((-1 > 42) && false) || (1 != -482))
    var v3 bool = (-97 <= 0)
    var v4 int8 = -128
    var v5 float32 = 3.14159
    v17 := []int{-959}
    if true {
        var v18 int32 = 2147483647
        _ = len(v17)
        var v19 bool = (-462 == -638)
        _ = v18
        _ = v19
    } else {
        v17 = append(v17, -1)
    }
    var v20 uint = v1
    var v21 int8 = 127
    if false {
        for fi22 := 0; fi22 < 14; fi22++ {
            for fi23 := 0; fi23 < 20; fi23++ {
                var v24 uint8 = 1
                fmt.Println(v1, v21, v4)
                _ = fi23
                _ = v24
            }
            _ = fi22
        }
        v3 = v3
    } else {
        if true {
            fmt.Println(v3)
            fmt.Println(v4)
            if false {
                v3 = (-800 == 934)
            } else {
                var v25 string = "a"
                var v26 int8 = (v21 / (((v4 - (^(v4)))) | 1))
                _ = v25
                _ = v26
            }
        }
        fmt.Println(v4, v4)
    }
    for fi27 := 0; fi27 < 1; fi27++ {
        v28 := GenS4{F0: 1, F1: 623}
        _ = fi27
        _ = v28
    }
    _ = len(v17)
    switch 236 {
    case 128:
        if len(v17) >= 2 {
            v17 = v17[1:len(v17)]
        }
    default:
        v4 -= 1
    }
    if len(v17) > 0 {
        v17[0] = 4294967295
    }
    var v29 int = 127
    var v30 uint64 = 0
    var v31 int8 = -1
    v31 += (^(v21))
    switch 2147483647 {
    case -1:
        v20 = uint(v30)
    default:
        v32 := GenS4{F0: v30, F1: 1}
        _ = v32
    }
    for fi33 := 0; fi33 < 15; fi33++ {
        var v34 uint8 = 0
        for fi35 := 0; fi35 < 8; fi35++ {
            var v36 int16 = 0
            _ = fi35
            _ = v36
        }
        _ = fi33
        _ = v34
    }
    for fi37 := 0; fi37 < 13; fi37++ {
        var v38 int8 = (v31 & (v21 & (v21 << ((v20) & 63))))
        _ = fi37
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
