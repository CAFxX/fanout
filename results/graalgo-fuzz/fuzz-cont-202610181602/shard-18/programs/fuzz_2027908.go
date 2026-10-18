package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 string }
func genFn0(p0 uint32, p1 int8, p2 int32) uint16 {
    v5 := GenS1{F0: 128, F1: -582}
    switch -925 {
    case 42:
        fmt.Println(p0, p0)
    case 31:
        if (2 >= 42) {
            v6 := []int{9223372036854775807, -2147483648, 128, 2147483647}
            p1 += p1
            if false {
                v7 := GenS1{F0: 255, F1: 285}
                if false {
                    for fi8 := 0; fi8 < 15; fi8++ {
                        v9 := GenS1{F0: 42, F1: 442}
                        p2 = p2
                        _ = fi8
                        _ = v9
                    }
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                }
                fmt.Println(p1)
                _ = v7
            }
            _ = v6
        } else {
            v10 := GenS1{F0: 2, F1: 0}
            _ = v10
        }
    case -133:
        switch "foo bar" {
        case "":
            p2 = 1
        case "foo bar":
            var v11 uint8 = 22
            _ = v11
        default:
            p2 = p2
        }
    default:
        p0 = 42
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 369
}
func genFn1(p0 uint32) uint64 {
    fmt.Println(p0)
    if true {
        fmt.Println(p0)
        var v12 uint16 = genFn0(p0, 33, 1)
        _ = v12
    }
    var v13 float32 = -2.25
    _ = p0
    _ = v13
    return 988
}

func main() {
    var v1 uint = 1
    v2 := make(map[int]string)
    v3 := []string{(("a" + "\n\t") + "\n\t"), "a"}
    var v4 uint32 = 15
    v14 := GenS2{F0: ((("\n\t" + "") + "hello") + "hello"), F1: 2}
    if (false && (!true)) {
        switch 189 {
        case 411:
            fmt.Println(v4, v4)
        case -9223372036854775808:
            delete(v2, -2147483648)
        }
        v4 = (^(v4))
    }
    v15 := GenS2{F0: (((("a" + "hello") + "") + "hello") + "foo bar"), F1: v4}
    switch 128 {
    case 4294967295:
        fmt.Println(v1)
    case 2147483647:
        for fi16 := 0; fi16 < 18; fi16++ {
            for fi17 := 0; fi17 < 2; fi17++ {
                switch fi16 {
                case -2147483648:
                    var v18 int32 = 433
                    _ = v18
                case -663:
                    v19 := GenS3{F0: ("a" + "\n\t")}
                    _ = v19
                case -9223372036854775808:
                    var v20 uint32 = v4
                    _ = v20
                }
                v21 := GenS1{F0: 203, F1: -693}
                _ = fi17
                _ = v21
            }
            if true {
                if true {
                    v1 = v1
                    v4 = v4
                    _ = len(v3)
                }
            } else {
                fmt.Println(v4, v1, v4)
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = len(v2)
            }
            _ = fi16
        }
    case 425:
        for fi22 := 0; fi22 < 10; fi22++ {
            if (uint(v4) < (v1 | v1)) {
                var v23 int = (fi22 - (-(fi22)))
                _ = v23
            } else {
                v1 = (v1 >> (((v1 >> (((^(v1))) & 63))) & 63))
            }
            var v24 int8 = -1
            for fi25 := 0; fi25 < 8; fi25++ {
                v1 = (v1 >> (((v1 * v1)) & 63))
                var v26 int = 32768
                _ = fi25
                _ = v26
            }
            _ = fi22
            _ = v24
        }
    default:
        var v27 uint32 = 541
        _ = v27
    }
    v1 = (v1 & v1)
    if false {
        v28 := GenS1{F0: 0, F1: -9223372036854775808}
        v29 := GenS2{F0: "hello", F1: v4}
        var v30 int16 = 42
        _ = v28
        _ = v29
        _ = v30
    }
    if false {
        var v31 uint16 = uint16(v1)
        v2[957] = "a"
        var v32 uint = 16
        _ = v31
        _ = v32
    }
    if true {
        var v33 float32 = 3.14159
        if len(v3) > 0 {
            v3[0] = "\n\t"
        }
        _ = v33
    } else {
        v4 = (v4 >> ((v1) & 63))
    }
    var v34 uint8 = 2
    var v35 float64 = ((1.5 - 1e10) + 3.14159)
    if len(v3) > 0 {
        v3[0] = "foo bar"
    }
    for fi36 := 0; fi36 < 8; fi36++ {
        v34 = v34
        var v37 uint8 = v34
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi36
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
}
