package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 int16 }
type GenS2 struct { F0 uint64; F1 uint; F2 uint16 }
type GenS3 struct { F0 uint8; F1 uint }
type GenS4 struct { F0 int64; F1 uint16 }
func genFn0() int64 {
    var v5 float64 = 1.5
    v6 := []int{32767}
    for fi7 := 0; fi7 < 14; fi7++ {
        for fi8 := 0; fi8 < 8; fi8++ {
            var v9 uint16 = 2
            _ = fi8
            _ = v9
        }
        v5 += -2.25
        _ = fi7
    }
    v5 += float64(-2.25)
    _ = v5
    _ = v6
    return 0
}
func genFn1() int64 {
    v10 := []int{}
    switch "" {
    case "a":
        var v11 uint32 = 2
        _ = v11
    default:
        v10 = append(v10, 700)
    }
    var v12 float32 = 1e10
    _ = v10
    _ = v12
    return -684
}

func main() {
    v1 := make(map[int]int)
    var v2 int = 32768
    var v3 uint8 = 199
    var v4 bool = (!(0 == 42))
    if v4 {
        var v13 uint16 = 255
        _ = v13
    }
    v4 = v4
    v4 = v4
    var v14 int16 = -615
    for fi15 := 0; fi15 < 10; fi15++ {
        for fi16 := 0; fi16 < 17; fi16++ {
            var v17 string = ("hello" + "")
            switch 2147483647 {
            case 42:
                var v18 int64 = 1
                _ = v18
            case -430:
                v19 := []int{fi16, (fi15 ^ fi16), fi15, (^(v2))}
                _ = v19
            default:
                var v20 uint = 18446744073709551615
                _ = v20
            }
            _ = fi16
            _ = v17
        }
        if false {
            v2 -= v2
            for fi21 := 0; fi21 < 3; fi21++ {
                v1[42] = fi15
                _ = fi21
            }
            switch "\n\t" {
            case "foo bar":
                v22 := GenS1{F0: v4, F1: 0, F2: (v14 + (v14 << ((uint(v3)) & 63)))}
                _ = v22
            case "hello":
                for fi23 := 0; fi23 < 1; fi23++ {
                    v24 := []int{(-(fi15)), (v2 / ((42) | 1)), (fi23 % ((2147483647) | 1))}
                    v25 := GenS1{F0: v4, F1: 42, F2: 371}
                    v3 = 204
                    _ = fi23
                    _ = v24
                    _ = v25
                }
            case "a":
                v26 := GenS1{F0: false, F1: 4294967295, F2: (v14 * v14)}
                _ = v26
            default:
                v2 += 1
            }
        } else {
            fmt.Println(v14)
            var v27 uint16 = 1
            v28 := GenS2{F0: 18446744073709551615, F1: 255, F2: 630}
            _ = v27
            _ = v28
        }
        v2 -= fi15
        _ = fi15
    }
    for fi29 := 0; fi29 < 16; fi29++ {
        var v30 float64 = (-2.25 / 1.0)
        _ = fi29
        _ = v30
    }
    for fi31 := 0; fi31 < 9; fi31++ {
        v14 += 42
        fmt.Println(v4)
        v2 += (fi31 + (^(v2)))
        _ = fi31
    }
    fmt.Println(v4, v14, v4)
    if ("foo bar" == "foo bar") {
        if v4 {
            v2 += (v2 - v2)
        }
        var v32 int16 = (v14 + v14)
        v14 = (v32 + (^(v14)))
        _ = v32
    } else {
        v4 = (false && v4)
        if false {
            v4 = (v4 && (42 > 0))
        } else {
            v33 := GenS3{F0: v3, F1: 255}
            _ = v33
        }
        if false {
            if (-1 < 0) {
                var v34 uint = 677
                if false {
                    for fi35 := 0; fi35 < 3; fi35++ {
                        for fi36 := 0; fi36 < 4; fi36++ {
                            var v37 int8 = -1
                            switch (v2 | 32768) {
                            case 32767:
                                fmt.Println(v14, fi36)
                                var v38 uint16 = 42
                                _ = v38
                            case -245:
                                if false {
                                    v39 := GenS2{F0: 2, F1: v34, F2: 32767}
                                    _ = v39
                                } else {
                                    var v40 int16 = (v14 ^ -543)
                                    _ = v40
                                }
                            case 2147483647:
                                var v41 int32 = 1
                                var v42 int = 32768
                                _ = v41
                                _ = v42
                            }
                            fmt.Println(v37)
                            _ = fi36
                            _ = v37
                        }
                        v43 := GenS4{F0: 9223372036854775807, F1: 484}
                        _ = fi35
                        _ = v43
                    }
                } else {
                    fmt.Println(v14, v4, v3)
                }
                fmt.Println(v3, v34)
                _ = v34
            }
            switch ("foo bar" + "\n\t") {
            case "a":
                delete(v1, (v2 << ((42) & 63)))
            case "\n\t":
                v2 -= 32768
            case "":
                v3 += 0
            default:
                switch "hello" {
                case "":
                    delete(v1, v2)
                }
            }
            switch "a" {
            case "":
                fmt.Println(v4, v3)
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
