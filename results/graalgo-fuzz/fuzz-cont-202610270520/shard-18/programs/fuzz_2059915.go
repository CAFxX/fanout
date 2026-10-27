package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 uint64; F1 int16; F2 int16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int8 }
func genFn0(p0 int32, p1 uint, p2 int) uint64 {
    fmt.Println(p2, p0)
    if false {
        fmt.Println(p0, p1)
    } else {
        p0 -= -214
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    v1 := make(map[string]string)
    var v2 uint16 = 65535
    v3 := make(map[uint32]int)
    var v4 uint8 = 206
    switch 945 {
    case 9223372036854775807:
        var v5 uint = 65535
        _ = v5
    default:
        v6 := GenS1{F0: -128}
        _ = v6
    }
    fmt.Println(v2, v2)
    switch 32768 {
    case 0:
        v7 := []int{}
        _ = v7
    case 32768:
        var v8 int64 = 707
        _ = v8
    case 2147483647:
        if (true && ((true || true) || (!false))) {
            for fi9 := 0; fi9 < 17; fi9++ {
                var v10 int = 4294967295
                if true {
                    v11 := GenS2{F0: 0, F1: -128}
                    var v12 uint = 65535
                    v4 -= 1
                    _ = v11
                    _ = v12
                } else {
                    var v13 uint64 = 2
                    var v14 int = 4294967295
                    var v15 int32 = 2147483647
                    _ = v13
                    _ = v14
                    _ = v15
                }
                _ = fi9
                _ = v10
            }
        } else {
            switch 912 {
            case 335:
                v16 := GenS3{F0: 2, F1: 1000, F2: -706}
                _ = v16
            case -1:
                var v17 uint8 = (v4 >> ((65535) & 63))
                _ = v17
            default:
                var v18 float32 = (1e10 / -2.25)
                _ = v18
            }
            for fi19 := 0; fi19 < 3; fi19++ {
                v20 := []int{fi19, (fi19 >> ((560) & 63)), (-(fi19))}
                _ = fi19
                _ = v20
            }
            if true {
                var v21 int32 = 1
                _ = v21
            }
        }
    }
    v22 := GenS1{F0: -108}
    var v23 uint16 = v2
    var v24 uint32 = 4294967295
    v23 = (^(v2))
    for fi25 := 0; fi25 < 8; fi25++ {
        if true {
            if false {
                switch "hello" {
                case "":
                    switch "foo bar" {
                    case "":
                        v26 := []int{32767, 2147483647, -9223372036854775808, -133}
                        _ = v26
                    default:
                        var v27 float32 = 3.14159
                        _ = v27
                    }
                case "a":
                    v4 -= v4
                case "\n\t":
                    if false {
                        for fi28 := 0; fi28 < 8; fi28++ {
                            var v29 int64 = 856
                            _ = fi28
                            _ = v29
                        }
                    } else {
                        v30 := GenS4{F0: false}
                        v23 += (v2 * v2)
                        _ = v30
                    }
                default:
                    v31 := GenS5{F0: (!(true && (false && true)))}
                    _ = v31
                }
                fmt.Println(v4, v2)
                var v32 int64 = 9223372036854775807
                _ = v32
            } else {
                v23 -= (^(v23))
            }
            if false {
                var v33 uint32 = (^(v24))
                _ = v33
            }
        }
        for fi34 := 0; fi34 < 20; fi34++ {
            var v35 uint8 = (v4 % ((221) | 1))
            var v36 string = ""
            v4 += (v4 << ((65535) & 63))
            _ = fi34
            _ = v35
            _ = v36
        }
        if ("" >= "a") {
            var v37 float64 = float64((1e10 - (1e10 * 0.1)))
            _ = v37
        } else {
            fmt.Println(v24)
        }
        _ = fi25
    }
    fmt.Println(v2, v23, v23)
    if true {
        var v38 uint32 = v24
        _ = v38
    } else {
        var v39 uint32 = v24
        v40 := GenS6{F0: 100}
        _ = v39
        _ = v40
    }
    v41 := []int{int(v2), 128}
    if len(v41) >= 2 {
        v41 = v41[1:len(v41)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v41:", len(v41), v41)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v41
}
