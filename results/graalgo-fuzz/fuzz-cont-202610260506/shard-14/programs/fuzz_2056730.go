package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 string }
type GenS2 struct { F0 uint32; F1 int64; F2 uint }
type GenS3 struct { F0 uint; F1 uint64; F2 int64 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]string)
    v3 := make(map[string]int)
    var v4 uint64 = 18446744073709551615
    var v5 string = ""
    v6 := make(map[uint32]int)
    var v7 int16 = -32768
    fmt.Println(v7, v5, v5)
    v7 -= 42
    switch v5 {
    case "a":
        _ = len(v3)
    case "\n\t":
        var v8 float64 = 3.14159
        _ = v8
    case "":
        if true {
            v7 += 1000
            fmt.Println(v4)
            v5 = ((((("a" + "foo bar") + (("" + v5) + "a")) + v5) + v5) + (v5 + v5))
        } else {
            if (!false) {
                var v9 bool = true
                v10 := GenS1{F0: 255, F1: 2, F2: "hello"}
                _ = v10
                _ = v9
            }
            v4 -= (v4 / ((v4) | 1))
        }
    }
    v4 = v4
    for fi11 := 0; fi11 < 6; fi11++ {
        var v12 uint8 = 206
        _ = fi11
        _ = v12
    }
    if (!true) {
        var v13 uint64 = v4
        for fi14 := 0; fi14 < 7; fi14++ {
            if true {
                var v15 uint16 = 744
                var v16 uint32 = 4294967295
                _ = v15
                _ = v16
            } else {
                switch 32767 {
                case -9223372036854775808:
                    var v17 uint8 = 127
                    _ = v17
                case 128:
                    if false {
                        var v18 int32 = 42
                        v19 := []int{fi14, (-(fi14))}
                        delete(v3, "")
                        _ = v18
                        _ = v19
                    }
                case 2147483647:
                    for fi20 := 0; fi20 < 3; fi20++ {
                        v21 := GenS2{F0: 943, F1: -9223372036854775808, F2: 510}
                        _ = fi20
                        _ = v21
                    }
                default:
                    v22 := GenS3{F0: 1, F1: v4, F2: 452}
                    _ = v22
                }
                var v23 float32 = 0.1
                v23 = v23
                _ = v23
            }
            _ = fi14
        }
        v24 := GenS1{F0: 0, F1: uint16(v4), F2: v5}
        _ = v13
        _ = v24
    }
    v25 := []int{4294967295, 0, 127, 32768}
    switch v5 {
    case "a":
        var v26 uint8 = 161
        _ = v26
    default:
        if (false || true) {
            var v27 string = (v5 + "a")
            v4 -= (^(v4))
            if false {
                var v28 int32 = 754
                fmt.Println(v27)
                _ = v28
            } else {
                if true {
                    var v29 uint16 = 443
                    var v30 uint16 = v29
                    _ = v29
                    _ = v30
                }
            }
            _ = v27
        }
    }
    for fi31 := 0; fi31 < 11; fi31++ {
        v3[v5] = fi31
        if false {
            var v32 float32 = 1e10
            fmt.Println(fi31, v4, v7)
            var v33 int64 = 42
            _ = v32
            _ = v33
        }
        if len(v25) > 0 {
            v25[0] = (fi31 >> ((4294967295) & 63))
        }
        _ = fi31
    }
    if true {
        v34 := GenS1{F0: 255, F1: 358, F2: (v5 + "a")}
        delete(v3, ("" + v5))
        _ = v34
    } else {
        v7 = 301
        var v35 uint32 = 776
        switch "\n\t" {
        case "":
            for fi36 := 0; fi36 < 15; fi36++ {
                for fi37 := 0; fi37 < 20; fi37++ {
                    if true {
                        v38 := GenS1{F0: 258, F1: 0, F2: v5}
                        if false {
                            var v39 uint64 = v4
                            v2[("hello" + "\n\t")] = "\n\t"
                            _ = v39
                        } else {
                            var v40 int = (fi36 & fi36)
                            _ = v40
                        }
                        var v41 int16 = v7
                        _ = v38
                        _ = v41
                    }
                    _ = fi37
                }
                for fi42 := 0; fi42 < 3; fi42++ {
                    fmt.Println(v35, fi36)
                    _ = fi42
                }
                _ = fi36
            }
        default:
            fmt.Println(v4, v35)
        }
        _ = v35
    }
    if len(v25) >= 2 {
        v25 = v25[1:len(v25)]
    }
    if len(v25) > 0 {
        v25[0] = -1
    }
    fmt.Println(v4)
    var v43 int = 32768
    fmt.Println(v4, v43, v7)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
    _ = v7
}
