package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 int64 }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 string; F1 int32; F2 bool }
type GenS4 struct { F0 string; F1 int16 }
func genFn0(p0 bool) uint {
    if p0 {
        v6 := GenS1{F0: "\n\t", F1: 0, F2: -9223372036854775808}
        _ = v6
    } else {
        v7 := []int{-1, -201, -766}
        var v8 uint8 = 51
        _ = v7
        _ = v8
    }
    v9 := []int{9223372036854775807}
    _ = len(v9)
    for fi10 := 0; fi10 < 13; fi10++ {
        v11 := GenS2{F0: -1, F1: -1}
        p0 = true
        _ = fi10
        _ = v11
    }
    if p0 {
        var v12 int8 = 40
        if p0 {
            v9 = append(v9, 32768)
            v13 := GenS2{F0: v12, F1: 2147483647}
            _ = v13
        }
        p0 = (587 <= 42)
        _ = v12
    } else {
        p0 = p0
        p0 = (3.14159 > (0.1 + 1e10))
    }
    _ = p0
    _ = v9
    return 65535
}

func main() {
    v1 := []string{("hello" + ("\n\t" + "foo bar"))}
    v2 := []int{}
    var v3 uint8 = 167
    v4 := make(map[int]int)
    var v5 int16 = 1
    for fi14 := 0; fi14 < 8; fi14++ {
        v3 = 98
        _ = fi14
    }
    switch ("hello" + "foo bar") {
    case "a":
        v4[-409] = -747
    case "":
        fmt.Println(v3, v3)
    }
    v3 += (v3 - 60)
    for fi15 := 0; fi15 < 20; fi15++ {
        for fi16 := 0; fi16 < 11; fi16++ {
            v5 = (v5 & v5)
            v17 := GenS3{F0: "", F1: 0, F2: true}
            _ = fi16
            _ = v17
        }
        _ = fi15
    }
    for fi18 := 0; fi18 < 8; fi18++ {
        switch 631 {
        case -1:
            switch 42 {
            case 9223372036854775807:
                var v19 uint8 = v3
                _ = v19
            case 32767:
                v3 = 2
            default:
                if (133 <= 18446744073709551615) {
                    v5 = -32768
                } else {
                    fmt.Println(fi18)
                    v5 -= (v5 % ((42) | 1))
                    switch 1 {
                    case 42:
                        v3 -= 128
                    case 128:
                        for fi20 := 0; fi20 < 18; fi20++ {
                            if true {
                                v21 := GenS3{F0: "foo bar", F1: -437, F2: (42 == int8(v5))}
                                _ = v21
                            } else {
                                _ = len(v2)
                            }
                            var v22 int64 = -690
                            var v23 string = "hello"
                            _ = fi20
                            _ = v22
                            _ = v23
                        }
                    case 127:
                        v24 := GenS4{F0: "hello", F1: (-(v5))}
                        _ = v24
                    }
                }
            }
        default:
            fmt.Println(v3, v3)
        }
        _ = fi18
    }
    fmt.Println(v3, v3)
    for fi25 := 0; fi25 < 19; fi25++ {
        var v26 uint16 = 1
        _ = fi25
        _ = v26
    }
    for fi27 := 0; fi27 < 4; fi27++ {
        if false {
            if true {
                v3 += 96
                var v28 int64 = -9223372036854775808
                v28 += v28
                _ = v28
            } else {
                v29 := GenS3{F0: "foo bar", F1: 738, F2: true}
                _ = v29
            }
        } else {
            var v30 int = 127
            _ = len(v1)
            _ = v30
        }
        _ = len(v1)
        _ = fi27
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
