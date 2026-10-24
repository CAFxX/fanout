package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 int16; F2 uint8 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 uint16, p1 int64) uint64 {
    if false {
        var v5 int8 = 127
        _ = v5
    }
    for fi6 := 0; fi6 < 6; fi6++ {
        for fi7 := 0; fi7 < 17; fi7++ {
            if false {
                v8 := []int{}
                _ = v8
            } else {
                if ("hello" > ("a" + (("\n\t" + "hello") + ("" + "hello")))) {
                    switch "hello" {
                    case "hello":
                        fmt.Println(p1)
                    case "foo bar":
                        p0 -= (p0 * 65535)
                    case "\n\t":
                        fmt.Println(p0, fi7)
                    }
                    v9 := []int{(fi6 / ((fi6) | 1)), fi7}
                    fmt.Println(p0, p1, p0)
                    fmt.Println(fi7, fi6, p1)
                    _ = v9
                } else {
                    if true {
                        if (42 != 0) {
                            v10 := []int{}
                            var v11 uint32 = 1
                            _ = v10
                            _ = v11
                        } else {
                            var v12 float32 = (((1.5 - 0.0) + 0.1) - 1.5)
                            _ = v12
                        }
                        v13 := []int{32767, fi6}
                        _ = v13
                    }
                    v14 := []int{fi6}
                    fmt.Println(p0)
                    switch (^(fi7)) {
                    case 32768:
                        if len(v14) > 0 {
                            v14[0] = (fi7 * -189)
                        }
                        if len(v14) > 0 {
                            v14[0] = -2147483648
                        }
                    default:
                        if (!false) {
                            v15 := GenS1{F0: (fi7 & fi7)}
                            _ = v15
                        }
                    }
                    var v16 string = (("foo bar" + "foo bar") + "\n\t")
                    _ = v14
                    _ = v16
                }
                var v17 int64 = p1
                v17 += v17
                _ = v17
            }
            var v18 bool = (!false)
            _ = fi7
            _ = v18
        }
        p1 += p1
        if false {
            if (-2147483648 < 2147483647) {
                switch fi6 {
                case -1:
                    p1 = p1
                case 2147483647:
                    fmt.Println(p0)
                default:
                    var v19 float32 = 1e10
                    _ = v19
                }
                var v20 float64 = ((0.0 * -2.25) / 3.14159)
                _ = v20
            }
            var v21 float32 = (1.5 / 1e10)
            p0 -= (p0 | (p0 | (p0 + (^(p0)))))
            _ = v21
        } else {
            var v22 int16 = 1
            if (0 == 2) {
                p1 = p1
            }
            fmt.Println(fi6, fi6, p1)
            _ = v22
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    return 42
}
func genFn1(p0 int16, p1 int8) int8 {
    p1 = (p1 | (p1 << ((4294967295) & 63)))
    p1 -= p1
    _ = p0
    _ = p1
    return (-(p1))
}

func main() {
    v1 := make(map[int]int)
    var v2 uint16 = 208
    var v3 uint16 = v2
    v4 := make(map[int]int)
    v23 := GenS2{F0: 42, F1: -32768, F2: 140}
    v24 := []int{32767, -9223372036854775808, -755, 32767}
    v25 := GenS2{F0: -1, F1: 42, F2: 0}
    switch ("a" + "") {
    case "\n\t":
        v26 := GenS1{F0: 0}
        _ = v26
    case "hello":
        var v27 bool = false
        _ = v27
    case "foo bar":
        var v28 bool = (2 >= 2)
        _ = v28
    default:
        v3 = (^(v3))
    }
    v29 := GenS3{F0: 1}
    v30 := v1[-9223372036854775808]
    if (0.1 < 0.1) {
        delete(v1, 128)
        v24 = append(v24, v30)
        v1[v30] = -59
    } else {
        var v31 string = "foo bar"
        var v32 uint = 906
        _ = v31
        _ = v32
    }
    v33 := GenS3{F0: 474}
    var v34 int32 = -2147483648
    v3 += v2
    switch v30 {
    case 49:
        _ = len(v4)
    default:
        v2 -= v2
    }
    v34 = (^(v34))
    _ = len(v4)
    switch 32767 {
    case 40:
        var v35 uint8 = 0
        _ = v35
    default:
        for fi36 := 0; fi36 < 13; fi36++ {
            var v37 int64 = -762
            v34 = -880
            if true {
                fmt.Println(v2, v30)
            } else {
                v30 = -9223372036854775808
                var v38 float32 = 3.14159
                if true {
                    v34 = v34
                }
                _ = v38
            }
            _ = fi36
            _ = v37
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v33
    _ = v34
    _ = v4
}
