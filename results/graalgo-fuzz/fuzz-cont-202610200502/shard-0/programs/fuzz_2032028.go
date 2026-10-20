package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int8 }
func genFn0(p0 uint16) int {
    if (!(!true)) {
        for fi5 := 0; fi5 < 5; fi5++ {
            fmt.Println(p0)
            for fi6 := 0; fi6 < 10; fi6++ {
                p0 -= (p0 ^ p0)
                var v7 float32 = float32(float64(0.0))
                if (-32768 != 1000) {
                    var v8 int16 = 1
                    _ = v8
                }
                var v9 uint64 = 2
                _ = fi6
                _ = v7
                _ = v9
            }
            _ = fi5
        }
        p0 = 32767
    } else {
        var v10 int8 = 100
        var v11 bool = false
        _ = v10
        _ = v11
    }
    fmt.Println(p0)
    var v12 uint16 = p0
    switch 32768 {
    case -1:
        if false {
            p0 -= p0
        }
    case 0:
        for fi13 := 0; fi13 < 7; fi13++ {
            if ((fi13 | fi13) <= -121) {
                fmt.Println(v12, p0, fi13)
            } else {
                v14 := []int{(fi13 & 2147483647), (fi13 % (((fi13 / ((fi13) | 1))) | 1))}
                v12 = (p0 - v12)
                _ = v14
            }
            if (0.0 >= (0.0 * ((-2.25 * 3.14159) * 1.5))) {
                v15 := []int{}
                p0 = 0
                p0 = (^(v12))
                _ = v15
            } else {
                for fi16 := 0; fi16 < 5; fi16++ {
                    switch "" {
                    case "":
                        if false {
                            var v17 int16 = 0
                            _ = v17
                        } else {
                            fmt.Println(fi16)
                            v18 := GenS1{F0: 4294967295, F1: 965, F2: -128}
                            _ = v18
                        }
                    }
                    _ = fi16
                }
                v19 := GenS1{F0: 0, F1: 0, F2: 1}
                if (!true) {
                    var v20 uint16 = p0
                    v21 := GenS1{F0: 255, F1: 1, F2: -14}
                    for fi22 := 0; fi22 < 5; fi22++ {
                        v23 := []int{(fi22 + (fi22 + -9223372036854775808))}
                        v20 += (v12 / ((p0) | 1))
                        var v24 bool = false
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                    _ = v20
                    _ = v21
                }
                _ = v19
            }
            v25 := GenS1{F0: 0, F1: 0, F2: 100}
            _ = fi13
            _ = v25
        }
    default:
        switch 4294967295 {
        case 9223372036854775807:
            for fi26 := 0; fi26 < 8; fi26++ {
                var v27 int = fi26
                v12 = v12
                _ = fi26
                _ = v27
            }
        default:
            var v28 float32 = 3.14159
            _ = v28
        }
    }
    _ = p0
    _ = v12
    return 32767
}

func main() {
    var v1 int64 = 42
    v2 := []string{"\n\t", ((((("" + "\n\t") + "foo bar") + "") + (("foo bar" + "\n\t") + "foo bar")) + "hello"), "a"}
    var v3 string = "\n\t"
    var v4 uint8 = 255
    v29 := GenS1{F0: 42, F1: -573, F2: 101}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch "\n\t" {
    case "":
        fmt.Println(v1, v4, v3)
    }
    fmt.Println(v1)
    if true {
        for fi30 := 0; fi30 < 18; fi30++ {
            var v31 int8 = 100
            _ = fi30
            _ = v31
        }
        v1 = 1
        v4 = v4
    }
    var v32 int8 = 1
    switch -2147483648 {
    case 42:
        var v33 int8 = (v32 / ((v32) | 1))
        _ = v33
    case 806:
        for fi34 := 0; fi34 < 5; fi34++ {
            var v35 int32 = -982
            v35 = (^(v35))
            v36 := GenS1{F0: 0, F1: (v1 >> ((2) & 63)), F2: v32}
            _ = fi34
            _ = v35
            _ = v36
        }
    case 2147483647:
        var v37 int = 9223372036854775807
        _ = v37
    default:
        v3 = v3
    }
    v32 = -11
    if true {
        v38 := GenS1{F0: 987, F1: (^(v1)), F2: (v32 - 0)}
        _ = v38
    }
    var v39 int8 = int8(v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v32
    _ = v39
    _ = v4
}
