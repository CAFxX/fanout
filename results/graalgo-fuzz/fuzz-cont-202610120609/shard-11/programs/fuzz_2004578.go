package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 uint32 }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 string; F1 string; F2 int32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint; F1 uint; F2 bool }
func genFn0(p0 bool, p1 int64) uint32 {
    if (507 != 4294967295) {
        for fi4 := 0; fi4 < 16; fi4++ {
            var v5 int = (^(fi4))
            if p0 {
                var v6 uint16 = 42
                var v7 uint8 = 2
                _ = v6
                _ = v7
            } else {
                var v8 int32 = 1
                var v9 float64 = 0.0
                if (!(1 != -110)) {
                    v10 := GenS1{F0: 65535, F1: 1, F2: 1}
                    _ = v10
                } else {
                    v8 += 2147483647
                    var v11 uint64 = 119
                    if (1 != 255) {
                        var v12 uint = 255
                        _ = v12
                    }
                    _ = v11
                }
                _ = v8
                _ = v9
            }
            _ = fi4
            _ = v5
        }
        p1 += p1
    } else {
        var v13 float32 = -2.25
        v14 := []int{}
        _ = v13
        _ = v14
    }
    var v15 string = "hello"
    _ = p0
    _ = p1
    _ = v15
    return 2147483647
}

func main() {
    var v1 int16 = 985
    var v2 float64 = 1.5
    var v3 int16 = v1
    for fi16 := 0; fi16 < 12; fi16++ {
        if true {
            var v17 int = 2147483647
            v3 += v1
            _ = v17
        } else {
            if false {
                if (32767 >= 492) {
                    v18 := GenS1{F0: 255, F1: 65535, F2: 512}
                    _ = v18
                }
                v19 := GenS2{F0: "hello", F1: (fi16 / (((^(fi16))) | 1))}
                _ = v19
            } else {
                v3 -= (v3 & 1000)
            }
            var v20 string = ("" + (("hello" + ("" + "hello")) + ("foo bar" + ("foo bar" + "foo bar"))))
            _ = v20
        }
        for fi21 := 0; fi21 < 5; fi21++ {
            var v22 int64 = 897
            fmt.Println(fi21, v3)
            var v23 string = ((("foo bar" + "hello") + (("\n\t" + ("hello" + ("hello" + "a"))) + "")) + "foo bar")
            _ = fi21
            _ = v22
            _ = v23
        }
        _ = fi16
    }
    switch "hello" {
    case "\n\t":
        v3 -= v1
    default:
        v1 += -153
    }
    v3 = v3
    var v24 int64 = -997
    if ((255 < 126) && ((false || (!(431 == genFn0((830 != 432), 445)))) && (227 > 1))) {
        v3 = (v1 - (v1 | -975))
        if false {
            v24 = (v24 % ((317) | 1))
            v25 := []int{}
            v2 = v2
            _ = v25
        }
        v1 = (-(v3))
    }
    v26 := []int{-575}
    fmt.Println(v1)
    v27 := GenS3{F0: "foo bar", F1: ("foo bar" + "hello"), F2: 42}
    var v28 int = 843
    for fi29 := 0; fi29 < 18; fi29++ {
        var v30 uint = 311
        v31 := GenS4{F0: false}
        var v32 int16 = v3
        _ = fi29
        _ = v30
        _ = v31
        _ = v32
    }
    for fi33 := 0; fi33 < 20; fi33++ {
        v3 = v1
        v26 = append(v26, (fi33 << ((255) & 63)))
        _ = fi33
    }
    for fi34 := 0; fi34 < 15; fi34++ {
        v35 := GenS5{F0: 4294967295, F1: 0, F2: false}
        _ = fi34
        _ = v35
    }
    v2 += v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v24
    _ = v26
    _ = v27
    _ = v28
    _ = v3
}
