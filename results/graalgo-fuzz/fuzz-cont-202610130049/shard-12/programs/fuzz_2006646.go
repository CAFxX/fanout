package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool }
func genFn0(p0 int8) int8 {
    var v6 uint = 969
    v7 := GenS1{F0: 892}
    _ = p0
    _ = v6
    _ = v7
    return (p0 / ((0) | 1))
}
func genFn1() int32 {
    for fi8 := 0; fi8 < 18; fi8++ {
        fmt.Println(fi8)
        var v9 uint16 = 412
        _ = fi8
        _ = v9
    }
    var v10 int = -838
    v10 = v10
    _ = v10
    return 0
}

func main() {
    v1 := []int{42, 32767}
    v2 := []string{("" + "foo bar")}
    v3 := []int{-153, 9223372036854775807}
    v4 := make(map[uint32]int)
    v5 := make(map[uint32]string)
    var v11 uint = 729
    switch -909 {
    case 605:
        v11 += v11
    case 128:
        var v12 uint = 255
        _ = v12
    case 0:
        var v13 int16 = 1
        _ = v13
    default:
        v14 := v4[4294967295]
        _ = v14
    }
    fmt.Println(v11)
    var v15 uint32 = 65535
    var v16 int = -9223372036854775808
    v15 += v15
    switch 2147483647 {
    case -9223372036854775808:
        v17 := GenS1{F0: 0}
        _ = v17
    case -875:
        if true {
            for fi18 := 0; fi18 < 18; fi18++ {
                if len(v3) > 0 {
                    v3[0] = v16
                }
                _ = fi18
            }
            v19 := GenS2{F0: 0}
            _ = v19
        }
    case 32767:
        switch ("\n\t" + ("foo bar" + "a")) {
        case "hello":
            for fi20 := 0; fi20 < 11; fi20++ {
                var v21 uint16 = 51
                _ = fi20
                _ = v21
            }
        case "foo bar":
            var v22 uint8 = 42
            _ = v22
        case "a":
            switch "hello" {
            case "":
                fmt.Println(v16, v11)
            case "hello":
                if (803 != 415) {
                    if (42 < 1) {
                        if ((false || true) && ((genFn0(127) <= 42) && true)) {
                            for fi23 := 0; fi23 < 12; fi23++ {
                                if (true || (v15 >= v15)) {
                                    var v24 uint32 = v15
                                    _ = v24
                                }
                                var v25 int8 = genFn0(-128)
                                _ = fi23
                                _ = v25
                            }
                            v15 += v15
                        }
                    } else {
                        var v26 int8 = -1
                        _ = v26
                    }
                    if ((true || ((((1e10 * float32(0.0)) - (0.1 / ((1e10 / -2.25) + ((3.14159 * ((0.1 + 1.5) + 0.0)) / float32(0.1))))) > -2.25) || (!(1 >= int16(v16))))) || false) {
                        for fi27 := 0; fi27 < 8; fi27++ {
                            v28 := GenS2{F0: 42}
                            _ = len(v5)
                            _ = fi27
                            _ = v28
                        }
                        v1 = append(v1, v16)
                    }
                    if ((1 == (v16 | v16)) && false) {
                        v11 += uint(v15)
                        delete(v5, v15)
                    } else {
                        var v29 uint = v11
                        _ = v29
                    }
                } else {
                    switch "hello" {
                    case "foo bar":
                        v11 = (v11 + 255)
                    default:
                        var v30 uint8 = 243
                        _ = v30
                    }
                    v31 := GenS1{F0: 42}
                    _ = v31
                }
            case "foo bar":
                v32 := GenS2{F0: -128}
                _ = v32
            }
        default:
            fmt.Println(v15, v16, v11)
        }
    default:
        v15 += (v15 | v15)
    }
    v33 := GenS3{F0: (!true)}
    var v34 int64 = 9223372036854775807
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
