package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 int64 }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint; F1 uint16 }
type GenS5 struct { F0 string; F1 uint64; F2 int32 }
func genFn0(p0 int64, p1 int8) uint16 {
    var v7 bool = (float32(0.0) != 0.0)
    v8 := []int{0, -9223372036854775808}
    if (v7 || true) {
        fmt.Println(p1)
    } else {
        var v9 int16 = 1000
        var v10 uint16 = 1
        var v11 uint16 = 0
        _ = v10
        _ = v11
        _ = v9
    }
    var v12 uint32 = 65535
    p0 -= (p0 % ((p0) | 1))
    _ = p0
    _ = p1
    _ = v12
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = -75
    var v3 int8 = (^(v2))
    var v4 int32 = 2147483647
    v5 := []string{"", ("\n\t" + "a"), ""}
    v6 := []int{119}
    v13 := GenS1{F0: 1, F1: 2, F2: -173}
    for fi14 := 0; fi14 < 5; fi14++ {
        v15 := GenS1{F0: (fi14 * 0), F1: 42, F2: 565}
        if len(v6) > 0 {
            v6[0] = (fi14 << ((0) & 63))
        }
        _ = fi14
        _ = v15
    }
    var v16 uint8 = 42
    var v17 int8 = 42
    var v18 uint16 = 435
    switch "a" {
    case "":
        var v19 uint64 = 42
        _ = v19
    case "hello":
        if (!true) {
            switch "hello" {
            case "a":
                if (!true) {
                    if (false && false) {
                        if (1 == uint32(v16)) {
                            v4 = (-(v4))
                            var v20 int = 0
                            v20 = v20
                            _ = v20
                        }
                        v21 := GenS1{F0: 1, F1: (v16 ^ 216), F2: 42}
                        _ = v21
                    }
                    switch ("a" + (("foo bar" + "a") + (("a" + "foo bar") + (("a" + "\n\t") + "\n\t")))) {
                    case "hello":
                        v16 = (v16 * v16)
                    case "\n\t":
                        v22 := GenS1{F0: -1, F1: v16, F2: 13}
                        _ = v22
                    default:
                        fmt.Println(v16, v2)
                    }
                    v23 := GenS2{F0: v18, F1: 55}
                    _ = v23
                }
            }
        } else {
            fmt.Println(v3, v16)
            if true {
                v24 := GenS2{F0: v18, F1: 60}
                fmt.Println(v3, v3, v4)
                if false {
                    var v25 uint16 = (v18 ^ (^(v18)))
                    v16 -= (v16 * 128)
                    v4 -= 1
                    _ = v25
                } else {
                    for fi26 := 0; fi26 < 17; fi26++ {
                        fmt.Println(fi26, v16)
                        var v27 int16 = -1
                        _ = fi26
                        _ = v27
                    }
                    v2 = -14
                }
                _ = v24
            }
        }
    case "\n\t":
        if false {
            fmt.Println(v16)
            var v28 uint16 = (v18 * (^(v18)))
            fmt.Println(v17, v28)
            _ = v28
        } else {
            if ((true && true) || false) {
                v17 -= 10
            }
            for fi29 := 0; fi29 < 15; fi29++ {
                fmt.Println(v18, v17, v2)
                fmt.Println(v2, v4)
                var v30 uint8 = (^(v16))
                _ = fi29
                _ = v30
            }
            fmt.Println(v2, v18)
        }
    default:
        if true {
            var v31 string = "a"
            if ((v16 >= v16) || ((v16 % ((v16) | 1)) != v16)) {
                fmt.Println(v2, v4)
                v16 = v16
                if false {
                    v2 = (-(v17))
                    if false {
                        v6 = append(v6, 1)
                        if false {
                            var v32 int8 = (v2 - 16)
                            _ = v32
                        } else {
                            var v33 uint = 2
                            _ = v33
                        }
                    } else {
                        for fi34 := 0; fi34 < 6; fi34++ {
                            var v35 int64 = 670
                            _ = fi34
                            _ = v35
                        }
                    }
                }
            } else {
                var v36 uint64 = 42
                for fi37 := 0; fi37 < 14; fi37++ {
                    v2 -= v2
                    var v38 uint = 1
                    v38 += 2
                    _ = fi37
                    _ = v38
                }
                if false {
                    for fi39 := 0; fi39 < 1; fi39++ {
                        if ((((-2.25 / 0.1) - (0.1 + float32(-2.25))) != float32((-2.25 / -2.25))) && false) {
                            var v40 bool = (-446 != 1)
                            _ = v40
                        } else {
                            _ = len(v1)
                            var v41 int = 128
                            _ = v41
                        }
                        v16 = (v16 << ((1) & 63))
                        _ = fi39
                    }
                } else {
                    v31 = ("a" + v31)
                }
                _ = v36
            }
            var v42 int32 = (v4 << ((804) & 63))
            _ = v31
            _ = v42
        } else {
            v43 := GenS3{F0: 622}
            v16 += (v16 | 128)
            v1[2147483647] = (((v2 / (((v3 % (((-(v17))) | 1))) | 1)) > (v17 | (^(v3)))) || false)
            _ = v43
        }
    }
    fmt.Println(v2)
    v16 = (v16 & v16)
    _ = len(v1)
    v44 := GenS4{F0: 0, F1: v18}
    v45 := GenS5{F0: "\n\t", F1: 755, F2: -2147483648}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
