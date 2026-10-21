package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 int64; F1 string }
func genFn0(p0 int64, p1 uint16, p2 uint) uint32 {
    v4 := GenS1{F0: p0}
    fmt.Println(p1, p0, p2)
    var v5 int16 = -644
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 711
}
func genFn1(p0 int64, p1 int16, p2 int64) string {
    v6 := []int{42, 544, -1}
    var v7 int = -2147483648
    for fi8 := 0; fi8 < 5; fi8++ {
        _ = len(v6)
        _ = fi8
    }
    var v9 int32 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v9
    return "hello"
}

func main() {
    v1 := []int{}
    var v2 int = 589
    var v3 string = ""
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi10 := 0; fi10 < 7; fi10++ {
        v2 += (v2 * 1)
        v11 := GenS2{F0: genFn0(9223372036854775807, 1, 42), F1: 17}
        var v12 int32 = 2147483647
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 uint = 4294967295
    switch 42 {
    case 700:
        v2 = v2
    case 128:
        fmt.Println(v3, v13, v3)
    case 1:
        var v14 uint8 = 2
        _ = v14
    default:
        if (true || false) {
            v3 = v3
            var v15 int = (v2 - -2147483648)
            switch "" {
            case "":
                for fi16 := 0; fi16 < 14; fi16++ {
                    fmt.Println(v13, v15)
                    for fi17 := 0; fi17 < 12; fi17++ {
                        v13 = (^(v13))
                        var v18 uint64 = 0
                        if (0 != -9223372036854775808) {
                            v2 -= v2
                            var v19 uint = uint(v18)
                            _ = v19
                        } else {
                            v2 = fi17
                            switch "\n\t" {
                            case "foo bar":
                                var v20 int16 = int16(v15)
                                _ = v20
                            }
                            v21 := GenS1{F0: 1}
                            _ = v21
                        }
                        _ = fi17
                        _ = v18
                    }
                    switch v2 {
                    case 32767:
                        for fi22 := 0; fi22 < 20; fi22++ {
                            var v23 int16 = -1
                            var v24 int64 = 9223372036854775807
                            v2 += v15
                            _ = fi22
                            _ = v23
                            _ = v24
                        }
                    case 319:
                        var v25 int32 = -159
                        _ = v25
                    default:
                        for fi26 := 0; fi26 < 18; fi26++ {
                            v13 = (v13 - v13)
                            _ = fi26
                        }
                    }
                    _ = fi16
                }
            case "a":
                for fi27 := 0; fi27 < 12; fi27++ {
                    v3 = v3
                    v28 := GenS3{F0: 692, F1: ((("a" + "hello") + "\n\t") + "foo bar")}
                    v1 = append(v1, (^(fi27)))
                    _ = fi27
                    _ = v28
                }
            }
            _ = v15
        }
    }
    for fi29 := 0; fi29 < 16; fi29++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v30 uint = 160
        v2 = fi29
        _ = fi29
        _ = v30
    }
    var v31 uint16 = 255
    v13 = v13
    switch v3 {
    case "foo bar":
        var v32 uint32 = 762
        _ = v32
    }
    fmt.Println(v13)
    v1 = append(v1, -22)
    if len(v1) > 0 {
        v1[0] = -9223372036854775808
    }
    v1 = append(v1, (-(v2)))
    _ = len(v1)
    for fi33 := 0; fi33 < 11; fi33++ {
        var v34 int = v2
        v31 = 21
        _ = fi33
        _ = v34
    }
    var v35 int32 = 42
    if true {
        _ = len(v1)
        var v36 int16 = 42
        _ = v36
    }
    var v37 int = v2
    v31 = (v31 ^ v31)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v31
    _ = v35
    _ = v37
}
