package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 uint16 }
func genFn0(p0 bool, p1 bool) uint16 {
    var v7 uint32 = 936
    if p1 {
        var v8 int32 = 2147483647
        _ = v8
    } else {
        var v9 float64 = float64(0.0)
        if p1 {
            v10 := []int{}
            _ = v10
        } else {
            var v11 int32 = -183
            _ = v11
        }
        _ = v9
    }
    for fi12 := 0; fi12 < 18; fi12++ {
        v13 := GenS1{F0: (-(fi12))}
        if (1 < 227) {
            switch "hello" {
            case "foo bar":
                if p1 {
                    var v14 bool = false
                    v15 := GenS1{F0: fi12}
                    _ = v14
                    _ = v15
                } else {
                    p0 = p1
                    for fi16 := 0; fi16 < 2; fi16++ {
                        if true {
                            v17 := GenS2{F0: "\n\t", F1: 2}
                            switch "hello" {
                            case "a":
                                p1 = (-2.25 < (((0.1 * 1e10) * (-2.25 / 1e10)) / float64(((0.0 - (1.5 / -2.25)) + 3.14159))))
                            case "":
                                if p0 {
                                    v18 := GenS2{F0: "a", F1: 60}
                                    p1 = (-654 >= 0)
                                    p0 = (126 >= 128)
                                    switch fi12 {
                                    case -436:
                                        v19 := []int{(fi16 * 32767), (-(fi16)), fi16}
                                        _ = v19
                                    case 9223372036854775807:
                                        var v20 uint = 144
                                        _ = v20
                                    }
                                    _ = v18
                                }
                            }
                            _ = v17
                        }
                        _ = fi16
                    }
                    var v21 uint16 = 65535
                    _ = v21
                }
            case "\n\t":
                v22 := GenS1{F0: fi12}
                _ = v22
            }
        }
        fmt.Println(p0, p0, p1)
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v7
    return 1
}

func main() {
    var v1 string = ("a" + "hello")
    v2 := make(map[string]bool)
    var v3 int8 = 127
    var v4 uint64 = 2
    var v5 int = -2147483648
    var v6 int32 = -743
    if true {
        for fi23 := 0; fi23 < 14; fi23++ {
            var v24 uint8 = 2
            _ = fi23
            _ = v24
        }
        switch v5 {
        case -790:
            var v25 int64 = 9223372036854775807
            _ = v25
        case -959:
            var v26 int8 = 1
            _ = v26
        case 32767:
            v6 = v6
        }
        v3 += (v3 * v3)
    } else {
        v4 -= v4
    }
    var v27 uint = 65535
    switch v1 {
    case "a":
        v28 := []int{}
        _ = v28
    default:
        fmt.Println(v4, v4)
    }
    v27 += (v27 & 183)
    v29 := GenS1{F0: -291}
    var v30 uint32 = 4294967295
    v3 += (v3 ^ (v3 - -1))
    v2[v1] = false
    delete(v2, v1)
    fmt.Println(v30, v5)
    switch (v1 + v1) {
    case "\n\t":
        v4 -= (v4 + (v4 / (((v4 * 208)) | 1)))
    case "":
        delete(v2, v1)
    case "hello":
        v30 = (^(v30))
    }
    var v31 string = "\n\t"
    switch "foo bar" {
    case "hello":
        v30 = (v30 % ((v30) | 1))
    default:
        v32 := v2[(("a" + v1) + "a")]
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
