package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint; F1 bool; F2 uint32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int; F1 int }
type GenS7 struct { F0 int8; F1 int32 }

func main() {
    v1 := []int{169}
    var v2 int64 = 529
    v3 := make(map[uint32]uint32)
    var v4 int16 = 0
    v2 = -9223372036854775808
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v4 += (v4 ^ int16(v2))
    v5 := GenS1{F0: 16, F1: 2147483647}
    if (!false) {
        fmt.Println(v2, v4)
        _ = len(v1)
    } else {
        v6 := GenS1{F0: 0, F1: 107}
        _ = v6
    }
    if (false || ((3.14159 != (-2.25 + (0.1 - float32(-2.25)))) && false)) {
        var v7 int8 = 0
        v4 = v4
        var v8 uint = 825
        _ = v7
        _ = v8
    } else {
        v4 = (-(v4))
        var v9 string = "foo bar"
        var v10 int32 = 1
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 17; fi11++ {
        if (-922 >= -1) {
            for fi12 := 0; fi12 < 14; fi12++ {
                v2 -= (v2 * 439)
                if (!(((("foo bar" + "foo bar") + "a") > "a") || (!(!true)))) {
                    if true {
                        switch "hello" {
                        case "hello":
                            v2 = v2
                        case "a":
                            v2 -= (v2 << ((0) & 63))
                        case "foo bar":
                            var v13 int32 = -67
                            _ = v13
                        }
                    }
                }
                v4 = v4
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                if (!(true || (1 >= 42))) {
                    v14 := GenS1{F0: 1, F1: 2}
                    v15 := GenS2{F0: -970}
                    if true {
                        v4 -= v4
                    } else {
                        v4 += v4
                        v4 += v4
                        if false {
                            v2 = -1
                            var v16 int64 = v2
                            _ = v16
                        } else {
                            v17 := GenS3{F0: true}
                            var v18 float32 = (-2.25 - (1e10 + -2.25))
                            v3[2147483647] = 65535
                            _ = v17
                            _ = v18
                        }
                    }
                    _ = v14
                    _ = v15
                } else {
                    var v19 uint32 = 42
                    if false {
                        if (-866 != 2147483647) {
                            fmt.Println(v2, fi11)
                            delete(v3, 2)
                            v20 := GenS2{F0: (-(v4))}
                            _ = v20
                        }
                    } else {
                        v21 := GenS4{F0: 1, F1: true, F2: v19}
                        _ = v21
                    }
                    _ = v19
                }
                _ = fi12
            }
            for fi22 := 0; fi22 < 20; fi22++ {
                switch "a" {
                case "foo bar":
                    if (!("a" >= "hello")) {
                        v23 := GenS3{F0: true}
                        v1 = append(v1, (-(fi11)))
                        _ = v23
                    }
                case "a":
                    var v24 uint32 = 881
                    _ = v24
                }
                v2 -= (-(v2))
                var v25 uint64 = 0
                _ = fi22
                _ = v25
            }
            if (263 >= 2) {
                v2 = (v2 * (^(v2)))
                fmt.Println(v4, v4, v2)
                for fi26 := 0; fi26 < 9; fi26++ {
                    fmt.Println(v2, v4)
                    v27 := GenS5{F0: 0}
                    switch -18 {
                    case 1:
                        var v28 uint = 1
                        _ = v28
                    case 127:
                        v29 := GenS6{F0: fi26, F1: 573}
                        _ = v29
                    }
                    _ = fi26
                    _ = v27
                }
            }
        }
        fmt.Println(fi11, fi11)
        _ = fi11
    }
    v2 += v2
    switch 32768 {
    case -391:
        fmt.Println(v4, v4)
    }
    fmt.Println(v2)
    if (!(!(1 <= -1))) {
        v4 += (v4 | (-(v4)))
    } else {
        fmt.Println(v2)
    }
    for fi30 := 0; fi30 < 19; fi30++ {
        v2 = -1
        fmt.Println(v2, fi30, v2)
        _ = fi30
    }
    fmt.Println(v2, v2)
    v31 := GenS7{F0: -30, F1: -678}
    _ = len(v3)
    var v32 int = 9223372036854775807
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
