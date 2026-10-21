package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 int32 }
func genFn0(p0 int8, p1 uint32) int64 {
    var v6 int = 128
    if (false || ((!false) && (p1 <= (p1 - (p1 % ((p1) | 1)))))) {
        if false {
            fmt.Println(p0)
            switch "" {
            case "\n\t":
                var v7 int16 = 776
                _ = v7
            }
        }
    } else {
        for fi8 := 0; fi8 < 16; fi8++ {
            v9 := GenS1{F0: 2, F1: (p0 ^ p0)}
            for fi10 := 0; fi10 < 14; fi10++ {
                if false {
                    switch ("foo bar" + "foo bar") {
                    case "":
                        v11 := GenS1{F0: 644, F1: (p0 ^ (p0 / ((p0) | 1)))}
                        _ = v11
                    }
                    fmt.Println(p0, p1, fi8)
                    v12 := GenS2{F0: 131, F1: 153}
                    _ = v12
                }
                _ = fi10
            }
            _ = fi8
            _ = v9
        }
        fmt.Println(p0, p1)
    }
    _ = p0
    _ = p1
    _ = v6
    return -736
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []uint32{2147483648, 129}
    var v3 int16 = 1
    var v4 int = 0
    v5 := []uint32{239, 4294967295}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch "" {
    case "a":
        switch 9223372036854775807 {
        case 42:
            v4 = v4
        }
    }
    for fi13 := 0; fi13 < 20; fi13++ {
        if (-918 > 9223372036854775807) {
            v14 := GenS3{F0: 1}
            _ = v14
        } else {
            fmt.Println(fi13)
            switch "\n\t" {
            case "foo bar":
                var v15 int8 = -114
                _ = v15
            case "hello":
                fmt.Println(fi13, v3)
            case "a":
                switch ("a" + "\n\t") {
                case "":
                    v4 = fi13
                case "\n\t":
                    for fi16 := 0; fi16 < 2; fi16++ {
                        var v17 float32 = 1e10
                        v3 = (v3 * (^(v3)))
                        _ = fi16
                        _ = v17
                    }
                case "a":
                    if len(v2) > 0 {
                        v2[0] = 65535
                    }
                }
            default:
                if (3.14159 != -2.25) {
                    switch "hello" {
                    case "hello":
                        v18 := GenS1{F0: 42, F1: -13}
                        _ = v18
                    case "foo bar":
                        for fi19 := 0; fi19 < 4; fi19++ {
                            if true {
                                var v20 float32 = 3.14159
                                _ = v20
                            }
                            var v21 uint32 = 0
                            delete(v1, v21)
                            _ = fi19
                            _ = v21
                        }
                    default:
                        v4 += v4
                    }
                    v4 = (v4 - fi13)
                }
            }
        }
        var v22 bool = false
        v4 = (fi13 % ((v4) | 1))
        _ = fi13
        _ = v22
    }
    if true {
        v4 -= v4
        var v23 int32 = 2147483647
        var v24 int16 = (^(v3))
        _ = v23
        _ = v24
    }
    v4 -= (v4 / (((v4 & (v4 & (v4 << ((255) & 63))))) | 1))
    var v25 bool = (2147483647 <= 307)
    if v25 {
        var v26 uint32 = 2
        v25 = v25
        fmt.Println(v26, v3)
        _ = v26
    } else {
        var v27 int = (v4 >> ((1) & 63))
        _ = v27
    }
    v25 = false
    v28 := v1[2]
    v29 := GenS3{F0: -1}
    var v30 bool = (-128 == 1)
    fmt.Println(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
