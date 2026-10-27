package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
func genFn0(p0 uint16, p1 uint32, p2 int) bool {
    v7 := GenS1{F0: 766, F1: 18446744073709551615}
    var v8 uint8 = 2
    var v9 int = 926
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return false
}
func genFn1(p0 int8) int {
    fmt.Println(p0)
    p0 -= (-(p0))
    _ = p0
    return 4294967295
}

func main() {
    var v1 bool = false
    var v2 uint32 = 224
    var v3 uint32 = (v2 | (^(v2)))
    var v4 bool = (255 <= 2)
    v5 := make(map[string]int)
    v6 := make(map[string]uint32)
    var v10 int = 128
    for fi11 := 0; fi11 < 9; fi11++ {
        v12 := v5["\n\t"]
        for fi13 := 0; fi13 < 2; fi13++ {
            if v1 {
                v12 = (^(v12))
            } else {
                fmt.Println(v2, v3)
                fmt.Println(v10, v2)
                var v14 float32 = (-2.25 * 3.14159)
                _ = v14
            }
            var v15 int64 = 1
            var v16 int8 = 1
            _ = fi13
            _ = v15
            _ = v16
        }
        v1 = false
        _ = fi11
        _ = v12
    }
    v10 = -941
    switch ("a" + ("" + "foo bar")) {
    case "foo bar":
        v17 := GenS1{F0: 14, F1: 18446744073709551615}
        _ = v17
    case "\n\t":
        var v18 uint16 = 255
        _ = v18
    case "hello":
        v19 := []int{(-(v10)), (^(v10)), genFn1(72), (-(v10))}
        _ = v19
    }
    for fi20 := 0; fi20 < 5; fi20++ {
        switch ("a" + "hello") {
        case "":
            v21 := []int{v10}
            _ = v21
        case "foo bar":
            v1 = v1
        case "hello":
            v4 = false
        default:
            if v1 {
                switch "foo bar" {
                case "a":
                    if v4 {
                        v1 = v1
                    }
                case "\n\t":
                    if true {
                        v2 -= 0
                        var v22 int32 = -2147483648
                        if (127 == 250) {
                            v4 = ((v3 / ((65535) | 1)) < 2147483648)
                            v4 = v1
                        } else {
                            v6[((("hello" + ("a" + ("foo bar" + ("" + "a")))) + "a") + (("" + "foo bar") + ""))] = 557
                        }
                        _ = v22
                    } else {
                        var v23 float64 = (1.5 * 1e10)
                        _ = v23
                    }
                case "foo bar":
                    v24 := []int{v10}
                    _ = v24
                }
                v2 = 546
            } else {
                if (18446744073709551615 >= 425) {
                    var v25 uint32 = (^(v2))
                    for fi26 := 0; fi26 < 5; fi26++ {
                        delete(v6, "\n\t")
                        v2 = v25
                        _ = fi26
                    }
                    _ = v25
                } else {
                    if (0 >= 605) {
                        v27 := GenS1{F0: 114, F1: 2}
                        var v28 int16 = 32767
                        _ = v27
                        _ = v28
                    } else {
                        v10 = fi20
                    }
                }
            }
        }
        v1 = true
        _ = fi20
    }
    var v29 uint8 = 62
    var v30 string = "\n\t"
    v10 = v10
    v31 := GenS1{F0: uint64(v2), F1: 917}
    if (v10 < (-(v10))) {
        v32 := GenS1{F0: 75, F1: 739}
        for fi33 := 0; fi33 < 9; fi33++ {
            v34 := GenS1{F0: 814, F1: 42}
            v30 = v30
            _ = fi33
            _ = v34
        }
        v1 = (true || (!v1))
        _ = v32
    } else {
        fmt.Println(v3, v10)
        if (0 <= -128) {
            fmt.Println(v4, v3, v30)
        } else {
            var v35 uint8 = (v29 | v29)
            _ = v35
        }
    }
    v3 = v2
    v36 := GenS1{F0: 4294967295, F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
