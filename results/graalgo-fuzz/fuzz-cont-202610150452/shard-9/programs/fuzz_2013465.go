package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 bool }
type GenS2 struct { F0 string; F1 int16; F2 int64 }
type GenS3 struct { F0 int16 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]string)
    var v3 int16 = -1
    v3 = 12
    if false {
        _ = len(v1)
    }
    if true {
        var v4 int = 438
        v5 := GenS1{F0: false, F1: (1 == -465), F2: false}
        v6 := []int{127}
        _ = v4
        _ = v5
        _ = v6
    }
    if (-546 == -2147483648) {
        v3 = v3
        if false {
            var v7 uint = 255
            if false {
                for fi8 := 0; fi8 < 12; fi8++ {
                    for fi9 := 0; fi9 < 1; fi9++ {
                        if true {
                            v7 = v7
                            switch -750 {
                            case 434:
                                if (!((v7 >> ((v7) & 63)) == v7)) {
                                    var v10 bool = false
                                    var v11 uint64 = 18446744073709551615
                                    var v12 uint32 = uint32(v7)
                                    v1[v12] = "a"
                                    _ = v10
                                    _ = v11
                                    _ = v12
                                }
                            case -9223372036854775808:
                                switch -2147483648 {
                                case -393:
                                    fmt.Println(v7, v3)
                                }
                            }
                        }
                        var v13 string = "foo bar"
                        _ = fi9
                        _ = v13
                    }
                    var v14 int8 = 42
                    _ = fi8
                    _ = v14
                }
                v7 -= 18446744073709551615
                v15 := GenS1{F0: (1e10 == 1.5), F1: (true || (378 >= v7)), F2: (72 >= 380)}
                _ = v15
            }
            _ = v7
        } else {
            v16 := GenS1{F0: (true || false), F1: true, F2: (0 > 0)}
            fmt.Println(v3)
            fmt.Println(v3)
            _ = v16
        }
    } else {
        if false {
            var v17 uint8 = 42
            v3 += -632
            _ = v17
        } else {
            v18 := GenS1{F0: ((true || true) || ((("" + "") + (("foo bar" + "foo bar") + "hello")) < "hello")), F1: false, F2: ((false && ((!true) || false)) || true)}
            _ = v18
        }
        v19 := GenS1{F0: true, F1: false, F2: ((false || true) && (-134 >= -9223372036854775808))}
        _ = v19
    }
    v3 -= v3
    if (515 >= 255) {
        var v20 int16 = (v3 ^ (v3 | 0))
        var v21 float64 = -2.25
        v22 := v1[279]
        _ = v20
        _ = v21
        _ = v22
    } else {
        v3 = 1000
    }
    if true {
        v23 := []int{1, 32768, 2147483647}
        _ = v23
    } else {
        v24 := GenS2{F0: ("hello" + (("a" + "hello") + ("" + "hello"))), F1: (v3 / ((v3) | 1)), F2: -1}
        for fi25 := 0; fi25 < 20; fi25++ {
            v2["\n\t"] = (("hello" + "a") + "")
            _ = fi25
        }
        v3 = v3
        _ = v24
    }
    for fi26 := 0; fi26 < 16; fi26++ {
        fmt.Println(fi26, v3)
        fmt.Println(v3, fi26)
        fmt.Println(v3, v3)
        _ = fi26
    }
    var v27 string = ""
    v27 = v27
    delete(v2, (("a" + (v27 + "")) + ((v27 + (v27 + v27)) + "hello")))
    v28 := []int{127, -9223372036854775808, 4294967295}
    v29 := GenS3{F0: v3}
    var v30 uint16 = 2
    fmt.Println(v27, v27)
    v31 := GenS2{F0: ("" + ((v27 + v27) + v27)), F1: v3, F2: 9223372036854775807}
    var v32 string = "a"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
}
