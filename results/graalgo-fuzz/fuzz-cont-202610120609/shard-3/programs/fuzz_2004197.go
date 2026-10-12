package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16 }
func genFn0() uint64 {
    var v6 uint = 255
    switch ((("\n\t" + "hello") + (("foo bar" + "hello") + ("hello" + (("hello" + "foo bar") + "hello")))) + ("\n\t" + "\n\t")) {
    case "":
        for fi7 := 0; fi7 < 1; fi7++ {
            v8 := GenS1{F0: 763, F1: 1, F2: 32767}
            v6 += (v6 & (v6 & (v6 / (((v6 & 4294967295)) | 1))))
            _ = fi7
            _ = v8
        }
    case "\n\t":
        var v9 bool = false
        _ = v9
    }
    fmt.Println(v6)
    _ = v6
    return 4294967295
}
func genFn1() uint64 {
    v10 := []int{32768, 615, -377, 127}
    if true {
        var v11 int16 = 32767
        var v12 uint = 42
        switch (("" + "") + "foo bar") {
        case "":
            v12 += (v12 % (((v12 & (v12 | v12))) | 1))
        case "a":
            var v13 int8 = -80
            _ = v13
        }
        _ = v11
        _ = v12
    }
    _ = v10
    return 167
}

func main() {
    var v1 int16 = -850
    var v2 uint16 = 422
    var v3 float64 = (float64(((0.1 / 0.1) * 0.0)) + -2.25)
    var v4 int64 = -1
    var v5 float32 = -2.25
    v3 = (v3 / (((((3.14159 - v3) * -2.25) - ((0.0 * v3) / (((v3 + -2.25) - v3) * (float64((v5 * 0.0)) / 1.5)))) * v3) - (v3 + v3)))
    v14 := GenS1{F0: 1, F1: 42, F2: v2}
    if (((false || true) && true) || (41 != 102)) {
        v15 := GenS2{F0: 42}
        v16 := []int{32768, -1}
        _ = v15
        _ = v16
    } else {
        fmt.Println(v5)
        fmt.Println(v4, v5)
    }
    if false {
        var v17 int16 = -338
        _ = v17
    } else {
        fmt.Println(v2, v2)
        v18 := []int{0, 32767}
        _ = v18
    }
    switch ("foo bar" + ("a" + "hello")) {
    case "a":
        v3 = (v3 / v3)
    case "foo bar":
        var v19 float64 = v3
        _ = v19
    case "hello":
        var v20 int = -1
        _ = v20
    }
    fmt.Println(v3, v5)
    fmt.Println(v2, v3)
    var v21 uint32 = 309
    v22 := []int{-1, 2147483647, -1}
    _ = len(v22)
    v23 := GenS3{F0: v1}
    v5 = (((v5 * v5) / float32(((0.0 / (float64((float32((-2.25 * v3)) + 1e10)) - (v3 + v3))) / v3))) / v5)
    v2 = v2
    fmt.Println(v3)
    var v24 int8 = 127
    v4 = (v4 + (v4 >> ((4294967295) & 63)))
    var v25 uint8 = 0
    var v26 int32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
