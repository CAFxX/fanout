package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 int8; F1 bool; F2 int64 }
type GenS4 struct { F0 int; F1 uint8 }
type GenS5 struct { F0 int32; F1 uint8 }
func genFn0(p0 bool, p1 uint) uint {
    v7 := []int{616}
    if len(v7) > 0 {
        v7[0] = 128
    }
    _ = p0
    _ = p1
    _ = v7
    return 4294967295
}
func genFn1(p0 bool, p1 uint) int16 {
    if (273 != 695) {
        if false {
            if (255 < 30) {
                v8 := GenS1{F0: -684, F1: 9223372036854775807}
                if true {
                    var v9 uint64 = 1
                    p1 -= (^(p1))
                    var v10 uint = (p1 + uint(v9))
                    _ = v10
                    _ = v9
                } else {
                    p0 = (0.0 >= (0.1 * 1e10))
                    p0 = ((700 == 65535) || true)
                    var v11 uint32 = 252
                    _ = v11
                }
                var v12 int8 = -1
                _ = v12
                _ = v8
            }
            v13 := GenS1{F0: 4294967295, F1: -9223372036854775808}
            var v14 int16 = 0
            _ = v13
            _ = v14
        }
        fmt.Println(p1, p1)
        v15 := GenS2{F0: 480, F1: (((("\n\t" + "a") + "hello") + "foo bar") + ("" + ""))}
        _ = v15
    }
    switch ((("a" + "foo bar") + "a") + "foo bar") {
    case "\n\t":
        var v16 int32 = 2147483647
        _ = v16
    }
    if p0 {
        v17 := GenS3{F0: 100, F1: (!(p0 || false)), F2: -1}
        _ = v17
    }
    for fi18 := 0; fi18 < 17; fi18++ {
        v19 := GenS1{F0: (^(fi18)), F1: 18}
        _ = fi18
        _ = v19
    }
    _ = p0
    _ = p1
    return -1
}

func main() {
    var v1 uint8 = 1
    var v2 float32 = float32(1e10)
    var v3 float32 = v2
    var v4 float64 = 3.14159
    var v5 bool = (v4 < v4)
    var v6 int64 = -1
    v5 = (175 == 128)
    var v20 string = ("hello" + "foo bar")
    var v21 int16 = 173
    fmt.Println(v2, v6, v4)
    v1 -= (v1 / (((^(v1))) | 1))
    for fi22 := 0; fi22 < 11; fi22++ {
        if (("" + (v20 + "")) >= v20) {
            if v5 {
                v23 := GenS2{F0: v6, F1: ((v20 + ("\n\t" + (v20 + "hello"))) + "foo bar")}
                fmt.Println(v21)
                _ = v23
            } else {
                v24 := GenS4{F0: 6, F1: 255}
                v2 = v3
                _ = v24
            }
            v3 = v3
            v20 = ((("\n\t" + ("foo bar" + (v20 + v20))) + ((("" + "\n\t") + v20) + "a")) + ((("\n\t" + v20) + ("foo bar" + (v20 + ("hello" + v20)))) + v20))
        }
        _ = fi22
    }
    var v25 int32 = -866
    var v26 uint16 = 711
    for fi27 := 0; fi27 < 18; fi27++ {
        var v28 float64 = v4
        var v29 uint64 = 0
        _ = fi27
        _ = v28
        _ = v29
    }
    v6 -= v6
    for fi30 := 0; fi30 < 15; fi30++ {
        var v31 bool = true
        for fi32 := 0; fi32 < 20; fi32++ {
            if v5 {
                var v33 float64 = 1e10
                v20 = (v20 + "a")
                var v34 int16 = genFn1(true, 0)
                v35 := GenS5{F0: int32(v26), F1: v1}
                _ = v33
                _ = v34
                _ = v35
            }
            var v36 int16 = (-(v21))
            _ = fi32
            _ = v36
        }
        _ = fi30
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
