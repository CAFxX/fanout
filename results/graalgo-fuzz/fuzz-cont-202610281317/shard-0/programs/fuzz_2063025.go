package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8 }
func genFn0(p0 bool, p1 int8) int16 {
    p0 = true
    if (-919 <= -9223372036854775808) {
        var v7 uint32 = 0
        _ = v7
    }
    var v8 string = ""
    _ = p0
    _ = p1
    _ = v8
    return 168
}
func genFn1(p0 int32, p1 uint16, p2 uint64) int8 {
    var v9 int16 = -913
    switch -357 {
    case 4294967295:
        var v10 int = 954
        _ = v10
    case 322:
        var v11 uint32 = 2
        _ = v11
    case 127:
        fmt.Println(p1, p2)
    default:
        v12 := GenS1{F0: -1}
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 29
}

func main() {
    v1 := []uint32{2147483648, 65535}
    var v2 float32 = 1e10
    var v3 uint32 = 2147483647
    v4 := []int{127}
    v5 := make(map[uint32]bool)
    v6 := make(map[string]bool)
    v3 = v3
    var v13 uint64 = 4294967295
    switch -9223372036854775808 {
    case -192:
        for fi14 := 0; fi14 < 12; fi14++ {
            for fi15 := 0; fi15 < 20; fi15++ {
                var v16 bool = (true && false)
                _ = fi15
                _ = v16
            }
            _ = fi14
        }
    case 149:
        v17 := v5[65535]
        _ = v17
    case 4294967295:
        fmt.Println(v3)
    default:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    v18 := GenS2{F0: 128}
    if true {
        v19 := v5[v3]
        switch "\n\t" {
        case "hello":
            v2 = (v2 + 0.1)
        case "":
            v20 := GenS1{F0: -106}
            _ = v20
        case "a":
            switch "\n\t" {
            case "":
                v21 := GenS3{F0: 42}
                _ = v21
            case "foo bar":
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            case "\n\t":
                var v22 int8 = 127
                _ = v22
            default:
                for fi23 := 0; fi23 < 16; fi23++ {
                    if ((v13 / (((v13 - v13)) | 1)) <= 2) {
                        fmt.Println(v19, v13, v3)
                        switch -831 {
                        case 0:
                            v13 -= 1
                        }
                    }
                    _ = fi23
                }
            }
        }
        v19 = v19
        _ = v19
    } else {
        v24 := GenS3{F0: 108}
        if (!true) {
            for fi25 := 0; fi25 < 2; fi25++ {
                for fi26 := 0; fi26 < 9; fi26++ {
                    fmt.Println(v2, fi26, fi26)
                    _ = fi26
                }
                for fi27 := 0; fi27 < 8; fi27++ {
                    var v28 int16 = -1
                    _ = fi27
                    _ = v28
                }
                v3 = 38
                _ = fi25
            }
            for fi29 := 0; fi29 < 1; fi29++ {
                fmt.Println(v3, fi29, fi29)
                _ = fi29
            }
            for fi30 := 0; fi30 < 16; fi30++ {
                v2 += v2
                _ = fi30
            }
        }
        _ = v24
    }
    var v31 float64 = (1e10 / 0.1)
    v2 -= v2
    if true {
        switch "foo bar" {
        case "\n\t":
            fmt.Println(v13, v2, v2)
        default:
            v3 = (v3 & 2147483648)
        }
        var v32 string = (("\n\t" + "a") + "a")
        _ = v32
    } else {
        var v33 int = 9223372036854775807
        var v34 int8 = -57
        _ = v33
        _ = v34
    }
    v35 := GenS3{F0: 127}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v31
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
