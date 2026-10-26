package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 uint8; F2 int8 }
type GenS4 struct { F0 uint8; F1 int8; F2 string }

func main() {
    var v1 uint8 = 106
    var v2 float64 = float64(1.5)
    var v3 float32 = (3.14159 + (1.5 / 3.14159))
    var v4 int64 = 963
    var v5 uint = 977
    var v6 uint16 = 255
    v7 := []int{}
    switch -1 {
    case -9223372036854775808:
        if true {
            if len(v7) > 0 {
                v7[0] = 127
            }
            var v8 uint64 = 42
            v8 -= (^(v8))
            _ = v8
        }
    }
    if false {
        var v9 float64 = 1e10
        var v10 int64 = v4
        _ = v10
        _ = v9
    } else {
        var v11 int8 = 1
        _ = v11
    }
    v12 := GenS1{F0: "", F1: uint32(v1), F2: "hello"}
    for fi13 := 0; fi13 < 11; fi13++ {
        var v14 uint64 = uint64(v1)
        var v15 uint = uint(v6)
        _ = fi13
        _ = v14
        _ = v15
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    var v16 float32 = ((-2.25 / 1.5) + 1.5)
    var v17 int16 = 32767
    if (18446744073709551615 < 0) {
        v17 = v17
        var v18 float32 = -2.25
        fmt.Println(v16, v6, v4)
        _ = v18
    }
    var v19 int16 = 32767
    var v20 uint = v5
    var v21 bool = false
    for fi22 := 0; fi22 < 11; fi22++ {
        if ((128 < (v1 / ((v1) | 1))) || v21) {
            if (!(!v21)) {
                var v23 int16 = (v17 % (((v19 << ((18446744073709551615) & 63))) | 1))
                v24 := GenS1{F0: "foo bar", F1: 0, F2: "a"}
                if v21 {
                    var v25 float32 = v16
                    _ = v25
                }
                _ = v23
                _ = v24
            } else {
                for fi26 := 0; fi26 < 18; fi26++ {
                    fmt.Println(v3, v3)
                    _ = fi26
                }
                if v21 {
                    var v27 int8 = -42
                    _ = v27
                }
            }
            for fi28 := 0; fi28 < 18; fi28++ {
                if v21 {
                    v29 := GenS2{F0: (v5 - 42)}
                    fmt.Println(v6, v20, v21)
                    v1 = (v1 & (v1 % ((v1) | 1)))
                    _ = v29
                }
                _ = fi28
            }
        } else {
            v3 += v16
            for fi30 := 0; fi30 < 19; fi30++ {
                v20 = (v20 << ((v20) & 63))
                _ = fi30
            }
        }
        _ = fi22
    }
    v31 := GenS3{F0: v20, F1: v1, F2: 100}
    v32 := GenS3{F0: v5, F1: (v1 + (v1 ^ v1)), F2: -117}
    switch ("foo bar" + "") {
    case "":
        v16 = 3.14159
    }
    switch -838 {
    case -859:
        if v21 {
            v33 := GenS4{F0: v1, F1: 127, F2: ((("foo bar" + (((("foo bar" + ("a" + "hello")) + ("a" + (("foo bar" + "hello") + "\n\t"))) + (("\n\t" + "a") + (("hello" + ("a" + ((("hello" + "foo bar") + "hello") + ""))) + ""))) + "\n\t")) + "foo bar") + "a")}
            fmt.Println(v4, v20)
            var v34 bool = ((v2 + (-2.25 * float64(v3))) > ((v2 * v2) - -2.25))
            _ = v33
            _ = v34
        }
    case 127:
        v35 := GenS4{F0: 127, F1: -128, F2: ("hello" + ("a" + ((("" + "foo bar") + ("hello" + "hello")) + "a")))}
        _ = v35
    case 822:
        var v36 int64 = v4
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
