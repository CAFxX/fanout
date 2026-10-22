package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 int32; F2 int }
func genFn0(p0 int16) uint8 {
    p0 = (p0 / ((p0) | 1))
    v4 := []int{}
    _ = p0
    _ = v4
    return 175
}
func genFn1(p0 bool, p1 uint8) bool {
    if p0 {
        var v5 bool = (9223372036854775807 < 0)
        switch 1 {
        case 4294967295:
            var v6 float64 = float64(0.1)
            _ = v6
        }
        _ = v5
    }
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    return (!p0)
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint64 = 125
    var v3 int8 = -128
    var v7 float32 = 3.14159
    var v8 int = -2147483648
    if (int8(v8) < v3) {
        for fi9 := 0; fi9 < 7; fi9++ {
            for fi10 := 0; fi10 < 17; fi10++ {
                var v11 uint32 = 2
                _ = fi10
                _ = v11
            }
            var v12 int = (fi9 + (fi9 & fi9))
            _ = fi9
            _ = v12
        }
    } else {
        for fi13 := 0; fi13 < 18; fi13++ {
            v14 := []int{v8, fi13}
            var v15 float32 = 0.0
            v16 := GenS1{F0: (^(v8))}
            _ = fi13
            _ = v14
            _ = v15
            _ = v16
        }
        if false {
            fmt.Println(v2)
            v17 := GenS1{F0: 9223372036854775807}
            var v18 uint8 = 128
            _ = v17
            _ = v18
        } else {
            for fi19 := 0; fi19 < 3; fi19++ {
                for fi20 := 0; fi20 < 15; fi20++ {
                    v3 -= (v3 - (v3 + (-(v3))))
                    var v21 uint16 = 255
                    var v22 string = "a"
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                v23 := GenS1{F0: fi19}
                _ = fi19
                _ = v23
            }
        }
    }
    v24 := []int{}
    fmt.Println(v8, v2)
    v2 += (v2 * v2)
    v3 = int8(v8)
    _ = len(v1)
    switch (v8 | v8) {
    case 4294967295:
        var v25 uint16 = 597
        _ = v25
    case -20:
        switch (v8 << ((65535) & 63)) {
        case 127:
            if true {
                if false {
                    if ("hello" != "foo bar") {
                        v26 := GenS1{F0: v8}
                        var v27 uint = 65535
                        _ = v26
                        _ = v27
                    } else {
                        for fi28 := 0; fi28 < 12; fi28++ {
                            if false {
                                var v29 int8 = (v3 % ((v3) | 1))
                                _ = v29
                            } else {
                                v7 -= v7
                            }
                            _ = len(v24)
                            _ = fi28
                        }
                    }
                    var v30 uint8 = 1
                    v31 := GenS2{F0: ("a" + ""), F1: 303, F2: v8}
                    _ = v30
                    _ = v31
                }
            } else {
                var v32 bool = (((1 < 2147483647) && true) || (-32768 >= 479))
                _ = v32
            }
        default:
            for fi33 := 0; fi33 < 10; fi33++ {
                v3 -= -128
                _ = fi33
            }
        }
    }
    if true {
        var v34 bool = (((false || (true || false)) && true) && true)
        _ = v34
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v7
    _ = v8
}
