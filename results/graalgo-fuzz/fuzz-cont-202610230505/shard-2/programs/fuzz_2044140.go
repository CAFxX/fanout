package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 int64 }
type GenS2 struct { F0 int32; F1 string; F2 uint16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64; F1 uint; F2 int }
type GenS5 struct { F0 uint16; F1 int16 }
type GenS6 struct { F0 uint32; F1 uint8; F2 uint8 }
func genFn0(p0 uint8, p1 int32, p2 uint8) int64 {
    var v6 int32 = -1
    if (0 != 312) {
        v6 += (p1 * 119)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 731
}
func genFn1() int {
    var v7 int = 448
    var v8 int32 = -2147483648
    if ((false || true) || true) {
        v8 += v8
    }
    v7 -= 128
    _ = v7
    _ = v8
    return v7
}

func main() {
    var v1 uint = 386
    var v2 float32 = 0.1
    v3 := make(map[uint32]bool)
    v4 := make(map[string]int)
    var v5 int8 = 100
    if false {
        if (!false) {
            var v9 uint64 = 0
            var v10 uint = (v1 / ((255) | 1))
            switch 1 {
            case 1:
                v9 += v9
            default:
                v1 -= (^(v1))
            }
            _ = v10
            _ = v9
        } else {
            v11 := v3[4294967295]
            var v12 int = 32767
            v13 := GenS1{F0: 2, F1: 128, F2: 42}
            _ = v11
            _ = v12
            _ = v13
        }
        v5 = v5
        var v14 int64 = 42
        _ = v14
    } else {
        var v15 uint32 = 2147483648
        var v16 uint64 = 1
        _ = v15
        _ = v16
    }
    switch 412 {
    case 1:
        delete(v3, 4294967295)
    }
    switch -853 {
    case -1:
        switch "a" {
        case "foo bar":
            var v17 float64 = (3.14159 * 1e10)
            _ = v17
        case "":
            if (v2 < v2) {
                for fi18 := 0; fi18 < 6; fi18++ {
                    v19 := []int{fi18, genFn1(), fi18, -2147483648}
                    fmt.Println(fi18)
                    _ = fi18
                    _ = v19
                }
                fmt.Println(v2)
                v20 := GenS2{F0: -589, F1: "hello", F2: 0}
                _ = v20
            } else {
                switch "\n\t" {
                case "\n\t":
                    v21 := GenS2{F0: -2147483648, F1: "\n\t", F2: 65535}
                    _ = v21
                case "":
                    v2 = v2
                default:
                    v22 := []int{127, genFn1(), 32767, -736}
                    _ = v22
                }
            }
        case "hello":
            var v23 int64 = 1
            _ = v23
        default:
            fmt.Println(v5)
        }
    default:
        v5 += (v5 - v5)
    }
    var v24 string = ("\n\t" + ("foo bar" + "\n\t"))
    v24 = (v24 + (v24 + "hello"))
    for fi25 := 0; fi25 < 4; fi25++ {
        v26 := GenS2{F0: -1, F1: v24, F2: 255}
        for fi27 := 0; fi27 < 8; fi27++ {
            v24 = (v24 + v24)
            if (true || (true && (false && true))) {
                switch (fi27 << ((2) & 63)) {
                case 0:
                    var v28 uint32 = 4294967295
                    _ = v28
                }
                fmt.Println(v24)
                fmt.Println(v2, v2)
            }
            delete(v4, "\n\t")
            _ = fi27
        }
        _ = fi25
        _ = v26
    }
    v1 = 65535
    switch "foo bar" {
    case "a":
        switch v24 {
        case "foo bar":
            var v29 uint64 = 0
            _ = v29
        default:
            for fi30 := 0; fi30 < 8; fi30++ {
                switch "\n\t" {
                case "\n\t":
                    v2 += v2
                default:
                    _ = len(v3)
                }
                var v31 int32 = -815
                _ = fi30
                _ = v31
            }
        }
    case "hello":
        v2 = v2
    default:
        v5 = 42
    }
    v2 = (v2 + (v2 * v2))
    v32 := v3[2]
    v33 := GenS3{F0: (^(v1))}
    if (2147483648 <= 2) {
        v34 := GenS1{F0: 92, F1: 37, F2: -9223372036854775808}
        var v35 uint = v1
        fmt.Println(v5, v32, v5)
        _ = v34
        _ = v35
    }
    if (v32 || v32) {
        if ((v32 && (!(false && (false || true)))) || (!true)) {
            v36 := GenS4{F0: 971, F1: v1, F2: -2147483648}
            v24 = v24
            if v32 {
                var v37 int64 = 42
                _ = v37
            }
            _ = v36
        } else {
            v38 := []int{-1, 32768}
            var v39 int16 = 1
            var v40 string = "foo bar"
            _ = v38
            _ = v39
            _ = v40
        }
        switch "\n\t" {
        case "foo bar":
            v41 := GenS5{F0: 65535, F1: 42}
            _ = v41
        case "":
            var v42 int32 = 2147483647
            _ = v42
        case "a":
            for fi43 := 0; fi43 < 18; fi43++ {
                v44 := GenS6{F0: 0, F1: 188, F2: 127}
                delete(v4, (v24 + v24))
                _ = fi43
                _ = v44
            }
        }
        v1 = v1
    } else {
        for fi45 := 0; fi45 < 6; fi45++ {
            fmt.Println(v1, fi45, v5)
            var v46 uint8 = 2
            if v32 {
                v32 = false
                var v47 int8 = -86
                _ = v47
            } else {
                if v32 {
                    for fi48 := 0; fi48 < 19; fi48++ {
                        var v49 int32 = 1
                        _ = fi48
                        _ = v49
                    }
                    v50 := GenS2{F0: -49, F1: v24, F2: 986}
                    _ = v50
                } else {
                    v24 = v24
                    fmt.Println(v1, v1, v46)
                }
            }
            _ = fi45
            _ = v46
        }
    }
    var v51 int8 = (v5 << ((v1) & 63))
    var v52 float32 = v2
    var v53 int16 = -32768
    v52 = ((v52 - (v2 / float32((float64(v2) + 1e10)))) - (v52 / v2))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v53
}
