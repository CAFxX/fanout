package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 uint; F1 uint32 }
type GenS5 struct { F0 uint64; F1 uint64; F2 int8 }
type GenS6 struct { F0 int }
func genFn0() int32 {
    var v7 uint32 = 2147483648
    if (0 > 862) {
        var v8 string = "hello"
        _ = v8
    }
    v9 := GenS1{F0: 2}
    _ = v7
    _ = v9
    return 706
}
func genFn1(p0 uint64, p1 int64) uint32 {
    for fi10 := 0; fi10 < 5; fi10++ {
        v11 := GenS2{F0: 65535, F1: 42}
        v12 := []int{fi10, (fi10 | fi10)}
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 string = "foo bar"
    _ = p0
    _ = p1
    _ = v13
    return 438
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]bool)
    v3 := []string{("" + ("a" + ("\n\t" + ("foo bar" + "\n\t")))), (("foo bar" + "foo bar") + "hello")}
    var v4 uint8 = 240
    var v5 float32 = 3.14159
    v6 := make(map[int]int)
    v4 += v4
    v5 -= 1.5
    var v14 uint = 89
    delete(v1, 2)
    v3 = append(v3, (("foo bar" + "") + "\n\t"))
    v15 := GenS3{F0: 32767, F1: 0}
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    v5 -= v5
    for fi16 := 0; fi16 < 2; fi16++ {
        v5 = v5
        if true {
            v17 := v2[genFn1(2, 441)]
            var v18 string = "\n\t"
            _ = v17
            _ = v18
        } else {
            var v19 uint64 = 4294967295
            v4 = (v4 + v4)
            v19 = (v19 >> ((65535) & 63))
            _ = v19
        }
        _ = len(v2)
        _ = fi16
    }
    switch "a" {
    case "foo bar":
        switch 1 {
        case 32767:
            var v20 int8 = -1
            _ = v20
        case 0:
            for fi21 := 0; fi21 < 1; fi21++ {
                v22 := GenS4{F0: v14, F1: 823}
                _ = fi21
                _ = v22
            }
        case -407:
            v23 := GenS4{F0: v14, F1: 893}
            _ = v23
        default:
            switch -2147483648 {
            case 42:
                var v24 uint8 = v4
                _ = v24
            case 0:
                v1[545] = true
            default:
                v25 := GenS5{F0: 42, F1: 0, F2: -7}
                _ = v25
            }
        }
    }
    _ = len(v2)
    fmt.Println(v4, v5, v4)
    v26 := GenS3{F0: 52, F1: 18446744073709551615}
    v14 = (v14 & v14)
    switch "" {
    case "\n\t":
        v27 := GenS6{F0: 543}
        _ = v27
    case "hello":
        fmt.Println(v14, v14, v14)
    case "":
        var v28 int32 = -2147483648
        _ = v28
    }
    var v29 uint64 = 2
    v30 := GenS5{F0: (v29 >> ((49) & 63)), F1: 2, F2: 0}
    if (825 > 32767) {
        fmt.Println(v4, v14, v4)
        v31 := GenS5{F0: 1, F1: (v29 * uint64(v14)), F2: 127}
        _ = v31
    } else {
        if ((true && false) || (false || false)) {
            v4 += (v4 + 0)
            if (((-1 > -621) || (0 == 1)) || true) {
                var v32 float64 = 0.1
                if ((v14 >> ((v14) & 63)) >= v14) {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                } else {
                    switch 2147483647 {
                    case 127:
                        for fi33 := 0; fi33 < 7; fi33++ {
                            v14 -= v14
                            v32 = v32
                            _ = fi33
                        }
                    default:
                        v34 := GenS4{F0: v14, F1: genFn1(uint64(v14), 0)}
                        _ = v34
                    }
                }
                _ = v32
            } else {
                var v35 float64 = 1.5
                v36 := GenS6{F0: -2147483648}
                fmt.Println(v29)
                _ = v35
                _ = v36
            }
            v37 := GenS2{F0: 255, F1: 937}
            _ = v37
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
