package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
func genFn0(p0 bool, p1 int8, p2 int) string {
    for fi7 := 0; fi7 < 20; fi7++ {
        switch "\n\t" {
        case "a":
            var v8 int16 = -490
            _ = v8
        default:
            v9 := GenS1{F0: 42, F1: (p1 ^ 0)}
            _ = v9
        }
        p1 = -102
        for fi10 := 0; fi10 < 19; fi10++ {
            var v11 string = ("" + "a")
            var v12 uint32 = 65535
            var v13 float32 = float32((1.5 - ((3.14159 / 0.1) / 1e10)))
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
        _ = fi7
    }
    if (18446744073709551615 < 2) {
        v14 := GenS1{F0: 314, F1: -128}
        _ = v14
    }
    p1 += p1
    var v15 uint32 = 0
    v16 := GenS1{F0: -1, F1: p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    return "a"
}

func main() {
    var v1 int = -844
    v2 := make(map[string]bool)
    var v3 int8 = 114
    v4 := []int{v1, v1, v1}
    v5 := make(map[uint32]int)
    v6 := make(map[int]bool)
    v4 = append(v4, v1)
    if (!false) {
        fmt.Println(v3, v1)
        for fi17 := 0; fi17 < 3; fi17++ {
            var v18 uint32 = 700
            v18 = v18
            _ = fi17
            _ = v18
        }
        v3 = (v3 + (v3 & v3))
    }
    switch "a" {
    case "a":
        var v19 uint = 42
        _ = v19
    case "\n\t":
        var v20 float32 = 1.5
        _ = v20
    default:
        _ = len(v4)
    }
    for fi21 := 0; fi21 < 11; fi21++ {
        fmt.Println(v1, v3)
        if true {
            for fi22 := 0; fi22 < 13; fi22++ {
                fmt.Println(fi22, v1, v1)
                var v23 uint16 = 643
                var v24 uint8 = 2
                _ = fi22
                _ = v23
                _ = v24
            }
        } else {
            for fi25 := 0; fi25 < 11; fi25++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                switch ((("foo bar" + ("" + "foo bar")) + "") + "hello") {
                case "\n\t":
                    var v26 int8 = -128
                    _ = v26
                case "hello":
                    fmt.Println(v1)
                case "a":
                    switch -450 {
                    case -577:
                        var v27 int16 = -32768
                        _ = v27
                    case 745:
                        var v28 uint64 = 42
                        _ = v28
                    case 42:
                        v6[fi21] = false
                    default:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                }
                _ = fi25
            }
            switch 32767 {
            case 4294967295:
                v3 -= 42
            }
        }
        _ = fi21
    }
    v1 += v1
    var v29 bool = (true || (0 == 0))
    delete(v6, v1)
    fmt.Println(v1)
    var v30 float32 = (3.14159 / -2.25)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v31 uint = 410
    if (-1 <= -32768) {
        var v32 uint32 = 2147483648
        for fi33 := 0; fi33 < 4; fi33++ {
            v1 = v1
            _ = fi33
        }
        _ = v32
    } else {
        v1 -= v1
    }
    switch ("hello" + "hello") {
    case "hello":
        var v34 uint32 = 543
        _ = v34
    case "":
        var v35 uint32 = 65535
        _ = v35
    case "a":
        v36 := GenS1{F0: 1, F1: (v3 / (((v3 ^ v3)) | 1))}
        _ = v36
    default:
        v30 += ((1.5 * v30) + v30)
    }
    switch 2147483647 {
    case -445:
        v3 -= -31
    case -2147483648:
        if len(v4) > 0 {
            v4[0] = (v1 % ((v1) | 1))
        }
    default:
        var v37 float32 = 0.0
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
