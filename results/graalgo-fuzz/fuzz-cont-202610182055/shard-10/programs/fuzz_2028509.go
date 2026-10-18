package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 bool; F1 int; F2 int32 }
type GenS4 struct { F0 int16; F1 int8 }

func main() {
    v1 := []int{-495, 220, 4294967295}
    v2 := make(map[uint32]bool)
    var v3 uint32 = 42
    var v4 int32 = 42
    switch 2147483647 {
    case 32768:
        v4 = v4
    case -2147483648:
        for fi5 := 0; fi5 < 1; fi5++ {
            v4 = 0
            _ = fi5
        }
    case 574:
        var v6 float64 = float64(3.14159)
        _ = v6
    }
    if false {
        if ((v4 | v4) == -2147483648) {
            v4 = v4
            var v7 int8 = 1
            _ = v7
        } else {
            var v8 uint32 = v3
            v4 = 42
            var v9 int8 = 1
            _ = v8
            _ = v9
        }
        if (!(48 > 1)) {
            for fi10 := 0; fi10 < 10; fi10++ {
                fmt.Println(fi10, v4, v3)
                v3 += v3
                v11 := GenS1{F0: v3}
                _ = fi10
                _ = v11
            }
            if false {
                var v12 float32 = (3.14159 + 0.1)
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                for fi13 := 0; fi13 < 16; fi13++ {
                    delete(v2, 2)
                    v1 = append(v1, fi13)
                    v3 += (v3 << ((2) & 63))
                    _ = fi13
                }
                _ = v12
            }
            var v14 int32 = v4
            _ = v14
        } else {
            var v15 int = 32767
            if true {
                v15 = (-(v15))
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v3 = 473
            } else {
                v16 := GenS2{F0: 18446744073709551615, F1: 9223372036854775807}
                var v17 float64 = ((3.14159 / 3.14159) / 1.5)
                _ = v16
                _ = v17
            }
            if false {
                switch "a" {
                case "\n\t":
                    v4 += (v4 + 921)
                default:
                    v15 = v15
                }
                fmt.Println(v3)
                _ = len(v1)
            }
            _ = v15
        }
        v3 = (v3 ^ 65535)
    } else {
        v4 = v4
        v18 := GenS3{F0: true, F1: -614, F2: (v4 % ((v4) | 1))}
        _ = v18
    }
    v19 := GenS3{F0: ((1 < 42) && true), F1: 2147483647, F2: (v4 + 42)}
    v4 = v4
    for fi20 := 0; fi20 < 7; fi20++ {
        fmt.Println(v3)
        for fi21 := 0; fi21 < 19; fi21++ {
            v1 = append(v1, fi21)
            if (true && true) {
                var v22 uint32 = 0
                _ = v22
            } else {
                if ((828 >= -1) || true) {
                    if true {
                        if false {
                            v4 -= v4
                        }
                        var v23 uint8 = 127
                        _ = v23
                    } else {
                        v4 = 549
                    }
                    delete(v2, v3)
                    fmt.Println(fi21)
                } else {
                    if (18446744073709551615 <= 1) {
                        var v24 string = ("\n\t" + "hello")
                        v25 := GenS4{F0: 353, F1: 0}
                        _ = v24
                        _ = v25
                    } else {
                        fmt.Println(fi20, fi20, fi20)
                        fmt.Println(fi20, fi20)
                        var v26 int32 = (v4 & (v4 & -1))
                        _ = v26
                    }
                    var v27 bool = (true && false)
                    v27 = true
                    _ = v27
                }
            }
            v4 += v4
            _ = fi21
        }
        _ = fi20
    }
    if (((("foo bar" + "") + "hello") + (("hello" + "a") + "a")) != "hello") {
        v3 = 2147483647
        fmt.Println(v4, v4)
        _ = len(v1)
    }
    v4 += v4
    var v28 float64 = float64(float32(1e10))
    if true {
        var v29 int8 = 108
        v30 := GenS3{F0: false, F1: 4294967295, F2: -1}
        _ = v29
        _ = v30
    } else {
        if false {
            v3 = 2147483647
            for fi31 := 0; fi31 < 20; fi31++ {
                fmt.Println(fi31)
                if false {
                    v28 = (v28 / v28)
                }
                var v32 uint64 = 2
                _ = fi31
                _ = v32
            }
            for fi33 := 0; fi33 < 20; fi33++ {
                var v34 int8 = 0
                v34 -= -1
                _ = fi33
                _ = v34
            }
        }
        for fi35 := 0; fi35 < 15; fi35++ {
            if len(v1) > 0 {
                v1[0] = fi35
            }
            var v36 bool = (false || ((^(fi35)) >= fi35))
            _ = fi35
            _ = v36
        }
        fmt.Println(v3)
    }
    for fi37 := 0; fi37 < 14; fi37++ {
        fmt.Println(v4)
        _ = fi37
    }
    for fi38 := 0; fi38 < 17; fi38++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v4 = (v4 & v4)
        fmt.Println(v28)
        _ = fi38
    }
    fmt.Println(v3, v4, v28)
    if true {
        var v39 int16 = -32768
        v4 -= (v4 >> ((uint(v3)) & 63))
        _ = v39
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v40 := GenS4{F0: -32768, F1: 1}
        fmt.Println(v3)
        _ = v40
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v28
    _ = v3
    _ = v4
}
