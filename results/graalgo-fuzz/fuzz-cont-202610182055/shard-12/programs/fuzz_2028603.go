package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int32 }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 uint32; F1 int }
type GenS4 struct { F0 uint8; F1 bool }
type GenS5 struct { F0 int8; F1 bool }
type GenS6 struct { F0 uint8; F1 int8 }
func genFn0(p0 uint32, p1 int32) int {
    var v4 int32 = (p1 | (p1 & (^(p1))))
    var v5 uint16 = 255
    for fi6 := 0; fi6 < 5; fi6++ {
        p1 = p1
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return -430
}

func main() {
    v1 := make(map[string]bool)
    var v2 bool = false
    v3 := make(map[string]string)
    var v7 string = "hello"
    var v8 int32 = -1
    if (!false) {
        v7 = "a"
        switch v7 {
        case "":
            switch v7 {
            case "":
                v7 = "hello"
            case "a":
                switch 128 {
                case -2147483648:
                    v9 := GenS1{F0: 2, F1: 32767, F2: (v8 % ((v8) | 1))}
                    _ = v9
                case -924:
                    for fi10 := 0; fi10 < 4; fi10++ {
                        if (-1 > 0) {
                            v8 = v8
                            var v11 float32 = 0.1
                            var v12 uint16 = 45
                            _ = v11
                            _ = v12
                        } else {
                            v2 = v2
                            for fi13 := 0; fi13 < 6; fi13++ {
                                var v14 int32 = (v8 >> ((18446744073709551615) & 63))
                                var v15 string = (("\n\t" + "\n\t") + v7)
                                _ = fi13
                                _ = v14
                                _ = v15
                            }
                            v2 = (!v2)
                        }
                        for fi16 := 0; fi16 < 11; fi16++ {
                            var v17 float64 = (-2.25 - 0.1)
                            _ = len(v3)
                            if ((((!false) && ((float32(3.14159) != (((-2.25 * ((-2.25 * 1.5) * 0.1)) * ((-2.25 / float32(1e10)) - 0.1)) * 0.1)) || (3.14159 != 3.14159))) || (true || v2)) && v2) {
                                v8 = 1
                                var v18 int = fi10
                                _ = v18
                            } else {
                                v2 = v2
                                v17 = 0.1
                            }
                            _ = fi16
                            _ = v17
                        }
                        _ = fi10
                    }
                case -815:
                    switch (((v7 + "") + v7) + ("a" + v7)) {
                    case "":
                        var v19 bool = v2
                        _ = v19
                    case "hello":
                        var v20 int64 = -710
                        _ = v20
                    case "\n\t":
                        switch (v7 + v7) {
                        case "":
                            v21 := GenS2{F0: -9223372036854775808, F1: 0}
                            _ = v21
                        default:
                            if (!(128 != 116)) {
                                v22 := []int{genFn0(325, (v8 | v8))}
                                var v23 int64 = 42
                                var v24 int32 = (v8 + v8)
                                _ = v22
                                _ = v23
                                _ = v24
                            }
                        }
                    }
                default:
                    for fi25 := 0; fi25 < 2; fi25++ {
                        v2 = false
                        _ = fi25
                    }
                }
            }
        case "\n\t":
            v26 := GenS2{F0: -2147483648, F1: 2}
            _ = v26
        }
        switch 9223372036854775807 {
        case -852:
            var v27 uint32 = 0
            _ = v27
        default:
            v8 = v8
        }
    }
    v2 = (v2 && false)
    switch ("\n\t" + "") {
    case "\n\t":
        for fi28 := 0; fi28 < 5; fi28++ {
            v2 = (-1 == 0)
            _ = fi28
        }
    }
    var v29 uint64 = 4294967295
    for fi30 := 0; fi30 < 17; fi30++ {
        var v31 uint16 = 1
        v31 = v31
        v32 := GenS3{F0: 1, F1: (fi30 << ((uint(v31)) & 63))}
        _ = fi30
        _ = v31
        _ = v32
    }
    var v33 uint64 = v29
    switch "foo bar" {
    case "\n\t":
        v29 = v29
    case "hello":
        if v2 {
            v1[v7] = (825 <= 127)
        }
    default:
        for fi34 := 0; fi34 < 12; fi34++ {
            for fi35 := 0; fi35 < 20; fi35++ {
                v36 := GenS4{F0: 0, F1: false}
                _ = len(v3)
                v37 := GenS2{F0: 443, F1: 0}
                _ = fi35
                _ = v36
                _ = v37
            }
            _ = fi34
        }
    }
    v29 -= v33
    v2 = v2
    if v2 {
        v2 = (v2 && v2)
        if v2 {
            var v38 uint = 42
            if (false || v2) {
                v2 = (!(2 <= 127))
                var v39 int8 = -74
                var v40 uint16 = 65535
                _ = v39
                _ = v40
            }
            _ = v38
        } else {
            for fi41 := 0; fi41 < 14; fi41++ {
                for fi42 := 0; fi42 < 16; fi42++ {
                    v43 := []int{(fi41 & (-(fi42)))}
                    _ = fi42
                    _ = v43
                }
                _ = fi41
            }
        }
    } else {
        if v2 {
            v44 := GenS5{F0: 42, F1: v2}
            v45 := GenS4{F0: 1, F1: (42 <= 820)}
            fmt.Println(v29)
            _ = v44
            _ = v45
        }
    }
    var v46 uint16 = 0
    var v47 float32 = -2.25
    var v48 bool = (!(42 == 334))
    var v49 int = -381
    v48 = v48
    v50 := GenS6{F0: 42, F1: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v33
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v50
    _ = v7
    _ = v8
}
