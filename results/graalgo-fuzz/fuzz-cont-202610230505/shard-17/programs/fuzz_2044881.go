package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 int32 }
func genFn0(p0 int64, p1 uint) uint32 {
    switch -503 {
    case 42:
        for fi4 := 0; fi4 < 20; fi4++ {
            v5 := GenS1{F0: 2, F1: int64(fi4)}
            v6 := GenS1{F0: 32767, F1: 9223372036854775807}
            _ = fi4
            _ = v5
            _ = v6
        }
    }
    v7 := []int{19, 128, 901}
    _ = p0
    _ = p1
    _ = v7
    return 640
}
func genFn1(p0 uint16, p1 string, p2 int8) string {
    if false {
        fmt.Println(p2, p2)
        for fi8 := 0; fi8 < 14; fi8++ {
            v9 := []int{}
            for fi10 := 0; fi10 < 6; fi10++ {
                var v11 int64 = 0
                var v12 float32 = (-2.25 + 1.5)
                var v13 int32 = -133
                _ = fi10
                _ = v11
                _ = v12
                _ = v13
            }
            _ = fi8
            _ = v9
        }
        if false {
            var v14 uint32 = 42
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 12; fi15++ {
        var v16 int = (fi15 * (fi15 << ((42) & 63)))
        _ = fi15
        _ = v16
    }
    fmt.Println(p2, p2)
    for fi17 := 0; fi17 < 8; fi17++ {
        v18 := []int{(fi17 + fi17), 383, 32767}
        var v19 int64 = 42
        _ = fi17
        _ = v18
        _ = v19
    }
    _ = p0
    _ = p1
    _ = p2
    return p1
}

func main() {
    var v1 int64 = -588
    v2 := make(map[int]int)
    var v3 int32 = 42
    switch -634 {
    case -2147483648:
        if ((0 < 454) && true) {
            if (false || (true || false)) {
                if (-32768 != -32768) {
                    v1 = (v1 >> ((49) & 63))
                    for fi20 := 0; fi20 < 13; fi20++ {
                        v1 = v1
                        for fi21 := 0; fi21 < 7; fi21++ {
                            if ("" != "") {
                                v1 = 0
                                v1 -= int64(v3)
                                v3 = v3
                            } else {
                                fmt.Println(v3)
                            }
                            v3 = -2147483648
                            _ = fi21
                        }
                        fmt.Println(v3, v3)
                        _ = fi20
                    }
                }
                v1 = v1
                v3 -= 485
            }
            var v22 int64 = (v1 * (v1 + v1))
            v3 = (v3 - (v3 >> ((1) & 63)))
            _ = v22
        } else {
            if true {
                var v23 uint = 96
                _ = v23
            } else {
                v24 := GenS1{F0: 731, F1: (v1 << ((42) & 63))}
                _ = v24
            }
        }
    case -869:
        fmt.Println(v3, v3)
    case -9223372036854775808:
        v3 = v3
    }
    switch 2147483647 {
    case 128:
        v25 := GenS1{F0: 32767, F1: v1}
        _ = v25
    case -357:
        delete(v2, -2147483648)
    default:
        if true {
            switch 269 {
            case 526:
                switch "hello" {
                case "foo bar":
                    v3 = v3
                case "":
                    var v26 int = 0
                    _ = v26
                }
            }
            var v27 float64 = (-2.25 * 0.0)
            _ = v27
        } else {
            v1 = v1
            if (0.0 != 0.1) {
                var v28 float64 = float64(3.14159)
                _ = v28
            } else {
                var v29 uint32 = genFn0(v1, 18446744073709551615)
                if false {
                    var v30 int64 = 0
                    v1 = 1
                    _ = v30
                } else {
                    if true {
                        var v31 int32 = 1
                        _ = v31
                    }
                }
                _ = v29
            }
            var v32 string = (("\n\t" + ("foo bar" + "")) + "foo bar")
            _ = v32
        }
    }
    if (true || (!(!(18446744073709551615 <= 18446744073709551615)))) {
        v33 := GenS2{F0: -2147483648}
        for fi34 := 0; fi34 < 8; fi34++ {
            v1 += v1
            _ = fi34
        }
        _ = v33
    }
    v35 := GenS2{F0: 447}
    v36 := GenS1{F0: 255, F1: 42}
    var v37 float32 = -2.25
    var v38 float32 = v37
    var v39 uint = 18446744073709551615
    v40 := []int{2147483647}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v40:", len(v40), v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v40
}
