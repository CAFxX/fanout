package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 int32; F1 string }
type GenS4 struct { F0 int64 }
func genFn0(p0 uint32) int {
    var v6 uint8 = 235
    var v7 float32 = (3.14159 + 1.5)
    v8 := []int{9223372036854775807}
    for fi9 := 0; fi9 < 5; fi9++ {
        var v10 string = "hello"
        if ((v7 + (v7 - float32((float64((v7 * v7)) / float64(1.5))))) < v7) {
            var v11 uint64 = 42
            if (v7 < float32((((1e10 - 0.0) * (0.0 + float64(v7))) - 1.5))) {
                v12 := GenS1{F0: (v10 + v10)}
                _ = v12
            }
            _ = v11
        }
        var v13 uint8 = (^(v6))
        _ = fi9
        _ = v10
        _ = v13
    }
    var v14 uint32 = (p0 * 2)
    _ = p0
    _ = v14
    _ = v6
    _ = v7
    _ = v8
    return 128
}
func genFn1(p0 int) int {
    p0 += p0
    var v15 int8 = 42
    _ = p0
    _ = v15
    return 9223372036854775807
}

func main() {
    v1 := make(map[string]int)
    var v2 int8 = -88
    var v3 bool = true
    v4 := make(map[string]string)
    v5 := make(map[string]int)
    v2 += v2
    v16 := GenS1{F0: "a"}
    for fi17 := 0; fi17 < 19; fi17++ {
        if (!v3) {
            var v18 bool = false
            _ = v18
        }
        if ((v3 && v3) || (!v3)) {
            v19 := []int{genFn0(816), fi17, genFn1(-990), fi17}
            fmt.Println(v2)
            _ = v19
        } else {
            if ("" < "foo bar") {
                var v20 uint8 = 1
                _ = v20
            }
        }
        for fi21 := 0; fi21 < 2; fi21++ {
            v22 := []int{}
            fmt.Println(fi21)
            _ = fi21
            _ = v22
        }
        _ = fi17
    }
    v23 := []int{}
    v24 := GenS1{F0: "hello"}
    fmt.Println(v3, v3)
    v25 := GenS1{F0: "hello"}
    var v26 float32 = 0.1
    for fi27 := 0; fi27 < 14; fi27++ {
        v2 -= v2
        fmt.Println(v3)
        _ = fi27
    }
    if (!(!v3)) {
        v28 := GenS2{F0: (!true), F1: 65535}
        var v29 int = -9223372036854775808
        switch (v29 ^ 32768) {
        case -232:
            v30 := GenS3{F0: 0, F1: "hello"}
            _ = v30
        case 0:
            v4["a"] = (("hello" + "foo bar") + "\n\t")
        default:
            var v31 int = (v29 * v29)
            _ = v31
        }
        _ = v28
        _ = v29
    }
    fmt.Println(v26, v26, v26)
    var v32 uint16 = 40
    var v33 uint8 = 127
    fmt.Println(v3, v26)
    if len(v23) > 0 {
        v23[0] = 4294967295
    }
    var v34 float32 = v26
    switch 1 {
    case -715:
        v35 := GenS3{F0: -1, F1: "a"}
        _ = v35
    case -174:
        v23 = append(v23, -272)
    case 128:
        if true {
            var v36 uint8 = v33
            var v37 int8 = (v2 | v2)
            _ = v36
            _ = v37
        } else {
            var v38 bool = false
            v26 -= (-2.25 - v26)
            _ = v38
        }
    }
    v39 := GenS4{F0: -1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v39
    _ = v4
    _ = v5
}
