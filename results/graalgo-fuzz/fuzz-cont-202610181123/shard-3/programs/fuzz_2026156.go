package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 int8; F1 uint64; F2 int8 }
func genFn0(p0 uint32, p1 uint) int {
    var v6 uint8 = 128
    p1 = p1
    v7 := []int{32767, 2147483647}
    var v8 string = "a"
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 4294967295
}
func genFn1(p0 uint64, p1 string) int {
    p0 = p0
    for fi9 := 0; fi9 < 11; fi9++ {
        for fi10 := 0; fi10 < 6; fi10++ {
            p1 = ""
            v11 := GenS1{F0: 593, F1: (42 == 541)}
            var v12 float32 = 0.0
            _ = fi10
            _ = v11
            _ = v12
        }
        v13 := []int{genFn0(597, uint(p0))}
        var v14 int64 = -383
        _ = fi9
        _ = v13
        _ = v14
    }
    var v15 bool = ((float64(1.5) >= 1.5) && (!false))
    _ = p0
    _ = p1
    _ = v15
    return 0
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]string)
    var v3 uint16 = 2
    var v4 int32 = 0
    v5 := []int{}
    v3 = v3
    if true {
        for fi16 := 0; fi16 < 3; fi16++ {
            var v17 string = ("hello" + "foo bar")
            fmt.Println(v3, v3, v3)
            v17 = "a"
            _ = fi16
            _ = v17
        }
        var v18 int16 = 42
        if false {
            if (4 >= 42) {
                v19 := GenS2{F0: 1, F1: 4, F2: 127}
                var v20 uint16 = v3
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = v19
                _ = v20
            }
            var v21 int64 = 260
            _ = v21
        }
        _ = v18
    }
    v4 = v4
    if (0 == -51) {
        switch 865 {
        case -8:
            switch ("hello" + "a") {
            case "a":
                var v22 float64 = float64(float32((0.0 * 1.5)))
                _ = v22
            case "\n\t":
                var v23 uint16 = (v3 - v3)
                _ = v23
            case "":
                for fi24 := 0; fi24 < 18; fi24++ {
                    if true {
                        v25 := GenS1{F0: uint64(v3), F1: ((v4 / ((v4) | 1)) < (v4 & 2147483647))}
                        var v26 string = ""
                        _ = v25
                        _ = v26
                    } else {
                        var v27 int16 = 1000
                        switch (fi24 + fi24) {
                        case -480:
                            v28 := GenS2{F0: 56, F1: 1, F2: 0}
                            _ = v28
                        default:
                            v3 = (v3 % (((^(v3))) | 1))
                        }
                        if len(v5) > 0 {
                            v5[0] = fi24
                        }
                        _ = v27
                    }
                    v4 = -2147483648
                    if len(v5) > 0 {
                        v5[0] = (fi24 | fi24)
                    }
                    _ = fi24
                }
            }
        case 766:
            v3 += v3
        default:
            v3 -= (v3 + 1)
        }
        v4 += v4
    }
    var v29 int32 = (v4 | 111)
    v1[(("\n\t" + ("" + ("hello" + "foo bar"))) + (("\n\t" + "hello") + ("a" + "hello")))] = "foo bar"
    for fi30 := 0; fi30 < 7; fi30++ {
        v2[228] = "\n\t"
        v3 = v3
        _ = fi30
    }
    v4 -= -768
    v4 -= v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
