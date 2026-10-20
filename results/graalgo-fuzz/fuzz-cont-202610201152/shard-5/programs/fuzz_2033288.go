package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 uint8; F1 int }
func genFn0(p0 uint32) int32 {
    v6 := GenS1{F0: 100, F1: 0}
    fmt.Println(p0)
    if false {
        var v7 int64 = int64(p0)
        switch 2147483647 {
        case -816:
            for fi8 := 0; fi8 < 3; fi8++ {
                for fi9 := 0; fi9 < 9; fi9++ {
                    var v10 string = "hello"
                    var v11 bool = false
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                fmt.Println(fi8)
                _ = fi8
            }
        }
        _ = v7
    } else {
        if false {
            var v12 string = "foo bar"
            _ = v12
        }
        p0 -= (p0 % (((^(p0))) | 1))
    }
    _ = p0
    _ = v6
    return -1
}

func main() {
    v1 := []int{4294967295, -1}
    var v2 uint = 727
    var v3 float64 = (3.14159 + (1.5 - (0.0 / 1.0)))
    v4 := make(map[string]bool)
    var v5 uint64 = 4294967295
    switch "hello" {
    case "foo bar":
        for fi13 := 0; fi13 < 6; fi13++ {
            var v14 uint = v2
            switch "foo bar" {
            case "":
                switch (^(fi13)) {
                case 4294967295:
                    v2 = v2
                default:
                    if false {
                        for fi15 := 0; fi15 < 8; fi15++ {
                            if true {
                                v5 = uint64(v2)
                                _ = len(v1)
                                var v16 float64 = ((v3 * v3) - ((float64((float32(v3) / float32(v3))) - 1e10) / (((float64(0.1) * (float64((float32(1.5) - 1.5)) + v3)) * v3) * v3)))
                                _ = v16
                            } else {
                                var v17 bool = true
                                _ = v17
                            }
                            _ = fi15
                        }
                        fmt.Println(v14, fi13)
                    }
                }
            case "foo bar":
                v18 := GenS1{F0: -1, F1: 0}
                _ = v18
            }
            _ = fi13
            _ = v14
        }
    case "hello":
        if false {
            v19 := GenS1{F0: 0, F1: 32767}
            v3 = 1.5
            v5 -= v5
            _ = v19
        } else {
            var v20 bool = (1 > 1)
            _ = v20
        }
    default:
        var v21 int32 = 924
        _ = v21
    }
    switch -976 {
    case -1:
        _ = len(v1)
    case 1:
        v3 += float64(float32(v3))
    case 2147483647:
        v3 = (v3 / -2.25)
    default:
        for fi22 := 0; fi22 < 6; fi22++ {
            v3 = v3
            var v23 int32 = 0
            var v24 uint8 = 2
            _ = fi22
            _ = v23
            _ = v24
        }
    }
    switch 42 {
    case 679:
        v25 := GenS1{F0: -128, F1: uint16(v5)}
        _ = v25
    case 9223372036854775807:
        switch -1 {
        case 4294967295:
            for fi26 := 0; fi26 < 18; fi26++ {
                v27 := GenS2{F0: 161, F1: ("" + "")}
                _ = fi26
                _ = v27
            }
        case 45:
            v3 -= (v3 + v3)
        case 435:
            v28 := GenS1{F0: -67, F1: 1}
            _ = v28
        }
    case 555:
        if len(v1) > 0 {
            v1[0] = 42
        }
    default:
        var v29 uint = v2
        _ = v29
    }
    _ = len(v4)
    var v30 int8 = 1
    var v31 string = ("foo bar" + "hello")
    var v32 uint16 = 615
    for fi33 := 0; fi33 < 1; fi33++ {
        v2 = v2
        v32 += v32
        _ = fi33
    }
    if true {
        fmt.Println(v31)
        v32 = v32
        for fi34 := 0; fi34 < 11; fi34++ {
            fmt.Println(v30, v3, v3)
            v32 += (v32 % ((v32) | 1))
            _ = fi34
        }
    } else {
        var v35 uint64 = 1
        _ = v35
    }
    v31 = (v31 + "foo bar")
    if (!false) {
        switch -2147483648 {
        case 32768:
            v5 -= 600
        case 2147483647:
            for fi36 := 0; fi36 < 20; fi36++ {
                for fi37 := 0; fi37 < 19; fi37++ {
                    fmt.Println(v30)
                    _ = fi37
                }
                _ = fi36
            }
        default:
            var v38 float32 = 1e10
            _ = v38
        }
    }
    fmt.Println(v32, v32)
    if true {
        var v39 int8 = (v30 ^ v30)
        _ = v39
    }
    if (1e10 > (1e10 + 3.14159)) {
        var v40 uint = 4294967295
        for fi41 := 0; fi41 < 9; fi41++ {
            var v42 int16 = 42
            if true {
                var v43 int16 = v42
                v30 = 0
                _ = v43
            }
            v44 := GenS3{F0: 127, F1: -9223372036854775808}
            _ = fi41
            _ = v42
            _ = v44
        }
        v45 := v4[v31]
        _ = v40
        _ = v45
    } else {
        var v46 uint8 = uint8(v32)
        _ = v46
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
