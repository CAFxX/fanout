package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 uint16; F2 string }
type GenS3 struct { F0 int64; F1 uint64 }
func genFn0(p0 int64, p1 int, p2 uint64) string {
    fmt.Println(p1, p1)
    v7 := GenS1{F0: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return ("" + ("a" + "a"))
}
func genFn1(p0 uint16, p1 int32, p2 int64) uint8 {
    v8 := GenS1{F0: 92}
    if (!(!true)) {
        v9 := []int{-355, 9223372036854775807, -519, -2147483648}
        for fi10 := 0; fi10 < 11; fi10++ {
            if false {
                v11 := GenS1{F0: 1}
                p0 = p0
                if (p0 != p0) {
                    var v12 float64 = ((((-2.25 * -2.25) / 1e10) / (3.14159 + 1.5)) * 3.14159)
                    if ((fi10 << ((2) & 63)) <= (fi10 + fi10)) {
                        v13 := GenS1{F0: 18446744073709551615}
                        _ = v13
                    }
                    if true {
                        var v14 int8 = 0
                        v15 := GenS1{F0: 854}
                        for fi16 := 0; fi16 < 6; fi16++ {
                            var v17 int32 = p1
                            var v18 uint = 707
                            v19 := GenS2{F0: v18, F1: p0, F2: "foo bar"}
                            _ = fi16
                            _ = v17
                            _ = v18
                            _ = v19
                        }
                        _ = v14
                        _ = v15
                    } else {
                        p1 += p1
                        p0 -= (p0 - p0)
                        for fi20 := 0; fi20 < 9; fi20++ {
                            var v21 uint64 = 18446744073709551615
                            _ = fi20
                            _ = v21
                        }
                    }
                    _ = v12
                }
                _ = v11
            }
            var v22 string = ("\n\t" + "hello")
            _ = fi10
            _ = v22
        }
        _ = v9
    } else {
        v23 := GenS3{F0: 42, F1: 42}
        _ = v23
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 195
}

func main() {
    v1 := make(map[int]string)
    var v2 float32 = 1.5
    var v3 int8 = -1
    var v4 uint16 = 32767
    var v5 bool = false
    var v6 int32 = 1
    var v24 uint64 = 18446744073709551615
    v1[0] = "a"
    v25 := GenS1{F0: v24}
    fmt.Println(v3, v24)
    if (v5 || true) {
        v26 := []int{127, 9223372036854775807, 424}
        for fi27 := 0; fi27 < 10; fi27++ {
            if len(v26) > 0 {
                v26[0] = (fi27 / ((0) | 1))
            }
            v6 -= (v6 % ((v6) | 1))
            _ = fi27
        }
        v6 = v6
        _ = v26
    }
    switch 9223372036854775807 {
    case 909:
        if v5 {
            v5 = true
        }
    case -252:
        v28 := GenS2{F0: 537, F1: 641, F2: "foo bar"}
        _ = v28
    case 42:
        var v29 uint = 197
        _ = v29
    default:
        switch 128 {
        case 9223372036854775807:
            var v30 string = ("" + "foo bar")
            _ = v30
        default:
            if (!v5) {
                delete(v1, 680)
            } else {
                v6 = (v6 * (v6 & (v6 * (v6 + v6))))
            }
        }
    }
    v31 := []int{}
    if v5 {
        v1[1] = "hello"
    }
    for fi32 := 0; fi32 < 19; fi32++ {
        if (v5 && ((("\n\t" + (("a" + "a") + "")) + "hello") != ("foo bar" + "a"))) {
            for fi33 := 0; fi33 < 6; fi33++ {
                v24 += (v24 + (v24 * v24))
                _ = fi33
            }
        }
        fmt.Println(v5, v4, v4)
        _ = fi32
    }
    var v34 float32 = ((((v2 - (v2 + (-2.25 / 1.0))) * v2) / 3.14159) / v2)
    v35 := GenS1{F0: v24}
    var v36 string = "\n\t"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v31
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
