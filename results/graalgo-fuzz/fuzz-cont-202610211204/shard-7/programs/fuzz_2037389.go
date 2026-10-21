package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int; F1 string; F2 int }
type GenS4 struct { F0 uint32; F1 int64 }
type GenS5 struct { F0 uint8; F1 int64 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int32 }
func genFn0(p0 int8) int64 {
    if (1e10 < 0.1) {
        if false {
            v6 := GenS1{F0: -269}
            if false {
                if false {
                    if false {
                        if false {
                            var v7 uint16 = 991
                            _ = v7
                        } else {
                            v8 := GenS1{F0: -710}
                            v9 := GenS2{F0: 0}
                            _ = v8
                            _ = v9
                        }
                    } else {
                        for fi10 := 0; fi10 < 5; fi10++ {
                            switch (fi10 - 32767) {
                            case -175:
                                p0 = (p0 + p0)
                            case 128:
                                fmt.Println(p0, fi10)
                            default:
                                var v11 uint32 = 42
                                _ = v11
                            }
                            _ = fi10
                        }
                        fmt.Println(p0)
                    }
                    fmt.Println(p0)
                }
                v12 := []int{127, -75, 197}
                for fi13 := 0; fi13 < 5; fi13++ {
                    var v14 uint16 = 0
                    if (-299 < -1) {
                        v15 := GenS3{F0: fi13, F1: ("" + "a"), F2: 32767}
                        _ = v15
                    } else {
                        fmt.Println(fi13, v14, p0)
                        var v16 int32 = 42
                        _ = v16
                    }
                    var v17 float64 = ((0.1 - (3.14159 * float64(3.14159))) - 1e10)
                    _ = fi13
                    _ = v14
                    _ = v17
                }
                _ = v12
            } else {
                p0 = (^(p0))
            }
            v18 := GenS3{F0: -1, F1: "a", F2: -1}
            _ = v18
            _ = v6
        } else {
            var v19 uint8 = 0
            _ = v19
        }
        for fi20 := 0; fi20 < 8; fi20++ {
            v21 := GenS4{F0: 1, F1: 247}
            v22 := GenS5{F0: 255, F1: 9223372036854775807}
            for fi23 := 0; fi23 < 12; fi23++ {
                var v24 int64 = 42
                _ = fi23
                _ = v24
            }
            _ = fi20
            _ = v21
            _ = v22
        }
    }
    v25 := []int{-9223372036854775808, -177, -2147483648, 32767}
    _ = p0
    _ = v25
    return 1
}

func main() {
    var v1 int8 = 28
    var v2 float32 = -2.25
    v3 := []int{206}
    var v4 float64 = (((0.1 / -2.25) * 0.1) - -2.25)
    var v5 uint = 65535
    var v26 int32 = 1
    for fi27 := 0; fi27 < 7; fi27++ {
        if false {
            for fi28 := 0; fi28 < 6; fi28++ {
                v29 := GenS6{F0: 42}
                v2 = v2
                _ = fi28
                _ = v29
            }
        } else {
            var v30 uint8 = uint8(v5)
            v1 -= v1
            _ = v30
        }
        fmt.Println(v26, v26)
        _ = fi27
    }
    if false {
        switch 2147483647 {
        case 573:
            for fi31 := 0; fi31 < 2; fi31++ {
                var v32 uint16 = 193
                _ = fi31
                _ = v32
            }
        }
    }
    var v33 uint64 = 93
    switch "\n\t" {
    case "foo bar":
        var v34 uint64 = 919
        _ = v34
    case "":
        for fi35 := 0; fi35 < 7; fi35++ {
            var v36 uint16 = 42
            fmt.Println(v26, v2, v33)
            _ = fi35
            _ = v36
        }
    default:
        v37 := GenS2{F0: 73}
        _ = v37
    }
    if (false || true) {
        v3 = append(v3, 42)
        if len(v3) > 0 {
            v3[0] = 0
        }
    } else {
        switch "hello" {
        case "\n\t":
            v26 = v26
        case "foo bar":
            if true {
                v38 := GenS7{F0: -2147483648}
                v5 += (^(v5))
                _ = v38
            }
        }
    }
    var v39 int16 = 111
    for fi40 := 0; fi40 < 12; fi40++ {
        var v41 string = ("hello" + "a")
        _ = fi40
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v33
    _ = v39
    _ = v4
    _ = v5
}
