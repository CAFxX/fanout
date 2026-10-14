package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 string; F1 int32; F2 string }
func genFn0(p0 int16, p1 bool, p2 uint32) int64 {
    v6 := GenS1{F0: (p0 / ((-631) | 1)), F1: -128}
    v7 := GenS2{F0: 0, F1: 2}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return -1
}

func main() {
    var v1 uint64 = 945
    v2 := []string{}
    v3 := []string{("" + "hello")}
    v4 := []uint32{2}
    v5 := make(map[uint32]string)
    if true {
        v1 = v1
    }
    for fi8 := 0; fi8 < 7; fi8++ {
        v9 := GenS3{F0: -128}
        delete(v5, 2)
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        fmt.Println(fi10, v1)
        delete(v5, 0)
        v11 := GenS2{F0: v1, F1: 128}
        _ = fi10
        _ = v11
    }
    var v12 int = -2147483648
    var v13 string = "hello"
    v12 -= (v12 + 4294967295)
    if (true || false) {
        switch "" {
        case "hello":
            v13 = "hello"
        }
        _ = len(v3)
        for fi14 := 0; fi14 < 17; fi14++ {
            if true {
                _ = len(v5)
                fmt.Println(v13, v13, v12)
                fmt.Println(v1, fi14, fi14)
            } else {
                v15 := v5[255]
                _ = v15
            }
            for fi16 := 0; fi16 < 8; fi16++ {
                v12 += fi16
                _ = fi16
            }
            fmt.Println(v13, fi14)
            _ = fi14
        }
    } else {
        if true {
            switch -360 {
            case 43:
                fmt.Println(v13, v12, v12)
            default:
                v2 = append(v2, v13)
            }
        } else {
            if len(v3) > 0 {
                v3[0] = "a"
            }
            switch v12 {
            case 9223372036854775807:
                for fi17 := 0; fi17 < 13; fi17++ {
                    v1 += (v1 >> ((0) & 63))
                    v13 = "hello"
                    if (631 <= -873) {
                        v18 := GenS3{F0: 1}
                        for fi19 := 0; fi19 < 2; fi19++ {
                            v1 += 18446744073709551615
                            _ = fi19
                        }
                        for fi20 := 0; fi20 < 1; fi20++ {
                            switch "hello" {
                            case "":
                                v12 -= fi20
                            case "foo bar":
                                var v21 int16 = 1000
                                _ = v21
                            }
                            v22 := GenS4{F0: (v13 + "foo bar"), F1: 42, F2: v13}
                            v13 = (("foo bar" + v13) + "a")
                            _ = fi20
                            _ = v22
                        }
                        _ = v18
                    }
                    _ = fi17
                }
            case -9223372036854775808:
                fmt.Println(v12, v13)
            }
            if true {
                v23 := GenS2{F0: v1, F1: 128}
                fmt.Println(v13)
                var v24 int16 = 454
                _ = v23
                _ = v24
            }
        }
    }
    var v25 uint16 = 370
    var v26 uint64 = (v1 << ((65535) & 63))
    for fi27 := 0; fi27 < 7; fi27++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
