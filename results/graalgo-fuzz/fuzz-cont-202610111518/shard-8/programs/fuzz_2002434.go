package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
func genFn0(p0 int16) int64 {
    switch 128 {
    case 20:
        p0 -= p0
    default:
        v7 := []int{128, 869, 42, 903}
        _ = v7
    }
    switch -840 {
    case 1:
        switch 331 {
        case 42:
            var v8 string = ""
            _ = v8
        case -9223372036854775808:
            var v9 int = -1
            _ = v9
        default:
            p0 = p0
        }
    }
    _ = p0
    return 9223372036854775807
}
func genFn1() int {
    if (0 <= 42) {
        for fi10 := 0; fi10 < 14; fi10++ {
            v11 := []int{fi10, fi10, fi10, fi10}
            var v12 int32 = -2147483648
            for fi13 := 0; fi13 < 5; fi13++ {
                v12 = v12
                fmt.Println(fi10, fi10)
                var v14 uint = 65535
                if false {
                    switch (-(fi13)) {
                    case 721:
                        var v15 int8 = -128
                        _ = v15
                    case 2147483647:
                        fmt.Println(v14)
                    case 0:
                        if true {
                            var v16 uint16 = uint16(v14)
                            v17 := GenS1{F0: (!false), F1: ((("\n\t" + "") + ("hello" + "")) + "a")}
                            var v18 string = ("foo bar" + "foo bar")
                            _ = v16
                            _ = v17
                            _ = v18
                        }
                        var v19 int32 = (-(v12))
                        v20 := GenS1{F0: true, F1: "\n\t"}
                        if false {
                            var v21 uint64 = 261
                            fmt.Println(v12, v12, v12)
                            _ = v21
                        }
                        if len(v11) > 0 {
                            v11[0] = fi10
                        }
                        _ = v19
                        _ = v20
                    default:
                        v14 = 785
                    }
                    fmt.Println(fi13, v12, v12)
                } else {
                    fmt.Println(v12, fi10, fi13)
                }
                _ = fi13
                _ = v14
            }
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    if (-9223372036854775808 >= genFn0(0)) {
        var v22 int = 4294967295
        v23 := []int{v22, 1, v22, v22}
        fmt.Println(v22)
        _ = v22
        _ = v23
    }
    switch 9223372036854775807 {
    case 32767:
        v24 := GenS1{F0: (true && true), F1: ""}
        _ = v24
    case 32768:
        if (false && true) {
            v25 := GenS1{F0: false, F1: "hello"}
            var v26 uint8 = 168
            _ = v25
            _ = v26
        }
    }
    var v27 int16 = 1000
    _ = v27
    return 536
}

func main() {
    v1 := make(map[string]uint32)
    var v2 string = ("" + "")
    v3 := []int{32767, 127, 128}
    var v4 int32 = -1
    v5 := []string{(v2 + v2), v2, ((v2 + v2) + v2), ""}
    var v6 float64 = 3.14159
    var v28 bool = true
    v4 = 2147483647
    for fi29 := 0; fi29 < 19; fi29++ {
        var v30 uint64 = 591
        for fi31 := 0; fi31 < 1; fi31++ {
            v6 = float64((1.5 - 3.14159))
            _ = len(v5)
            _ = fi31
        }
        var v32 uint64 = v30
        _ = fi29
        _ = v30
        _ = v32
    }
    var v33 int8 = 100
    fmt.Println(v2, v28, v6)
    var v34 bool = v28
    var v35 int64 = 0
    v36 := GenS1{F0: (v28 || (!((!(v28 || true)) && false))), F1: (v2 + (v2 + v2))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
