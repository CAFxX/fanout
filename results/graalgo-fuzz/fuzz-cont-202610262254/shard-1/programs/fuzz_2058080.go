package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 int8 }
func genFn0(p0 bool, p1 uint16) uint16 {
    p1 = 255
    for fi5 := 0; fi5 < 12; fi5++ {
        for fi6 := 0; fi6 < 20; fi6++ {
            switch (fi6 >> ((18446744073709551615) & 63)) {
            case 2147483647:
                v7 := GenS1{F0: 1, F1: 277}
                _ = v7
            case 0:
                var v8 bool = ((0.1 + 1.5) != 0.1)
                _ = v8
            default:
                var v9 float32 = (1e10 * 1.5)
                _ = v9
            }
            v10 := GenS1{F0: -1, F1: 694}
            if p0 {
                var v11 int64 = -1
                switch "" {
                case "":
                    fmt.Println(p1, v11, fi5)
                case "foo bar":
                    v12 := []int{fi5, -2147483648}
                    _ = v12
                case "hello":
                    v13 := GenS1{F0: 0, F1: 33}
                    _ = v13
                default:
                    if p0 {
                        var v14 int32 = -1
                        _ = v14
                    } else {
                        if (!true) {
                            v15 := []int{fi6, (^(fi5)), fi6}
                            var v16 uint = 807
                            switch "\n\t" {
                            case "":
                                var v17 uint16 = p1
                                _ = v17
                            }
                            _ = v15
                            _ = v16
                        } else {
                            fmt.Println(fi5)
                        }
                        var v18 uint8 = 42
                        var v19 uint16 = p1
                        _ = v18
                        _ = v19
                    }
                }
                switch "hello" {
                case "\n\t":
                    v20 := GenS1{F0: 1000, F1: 18446744073709551615}
                    _ = v20
                case "":
                    v21 := GenS1{F0: 524, F1: 387}
                    _ = v21
                case "foo bar":
                    p1 += (p1 * (p1 * 849))
                }
                _ = v11
            }
            var v22 int32 = -243
            _ = fi6
            _ = v10
            _ = v22
        }
        if true {
            switch fi5 {
            case 1:
                var v23 string = "hello"
                _ = v23
            default:
                var v24 float64 = 3.14159
                _ = v24
            }
            fmt.Println(p0, fi5, fi5)
        }
        p1 -= 769
        _ = fi5
    }
    switch "" {
    case "foo bar":
        for fi25 := 0; fi25 < 9; fi25++ {
            var v26 uint16 = (p1 >> ((373) & 63))
            for fi27 := 0; fi27 < 9; fi27++ {
                v28 := []int{(^(fi27)), -944}
                _ = fi27
                _ = v28
            }
            v29 := GenS1{F0: 42, F1: 965}
            _ = fi25
            _ = v26
            _ = v29
        }
    default:
        switch ("foo bar" + ("\n\t" + ("hello" + "\n\t"))) {
        case "a":
            for fi30 := 0; fi30 < 9; fi30++ {
                p0 = false
                _ = fi30
            }
        case "hello":
            v31 := GenS2{F0: 102}
            _ = v31
        }
    }
    _ = p0
    _ = p1
    return (^(p1))
}
func genFn1(p0 uint32, p1 bool) int32 {
    if p1 {
        var v32 uint32 = (p0 | p0)
        _ = v32
    } else {
        var v33 uint32 = (p0 / ((p0) | 1))
        var v34 int = -1
        v34 = v34
        _ = v33
        _ = v34
    }
    if false {
        p1 = p1
        p0 = p0
        p1 = true
    }
    fmt.Println(p0, p0)
    p1 = ("a" != ("\n\t" + "hello"))
    _ = p0
    _ = p1
    return -660
}

func main() {
    var v1 int = 9223372036854775807
    v2 := make(map[int]int)
    v3 := []uint32{4294967295, 2147483647, 594, 4294967295}
    var v4 int8 = -128
    var v35 int8 = v4
    var v36 uint64 = 260
    v2[100] = v1
    v36 += v36
    if (v4 < (v4 / ((v35) | 1))) {
        var v37 int = v1
        _ = v37
    } else {
        var v38 string = "hello"
        _ = v38
    }
    v39 := v2[v1]
    v40 := GenS2{F0: (v35 + v35)}
    if false {
        v35 = v35
    } else {
        var v41 int64 = -9223372036854775808
        fmt.Println(v1, v1)
        if (-2.25 == (1.5 - ((1.5 * 0.0) + 1.5))) {
            if len(v3) > 0 {
                v3[0] = 843
            }
            var v42 float32 = 1.5
            var v43 uint = 981
            _ = v42
            _ = v43
        }
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v39
    _ = v4
    _ = v40
}
