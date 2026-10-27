package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 int8; F1 uint64; F2 string }
type GenS4 struct { F0 uint8; F1 string; F2 int }
type GenS5 struct { F0 int16 }
func genFn0(p0 int32, p1 uint) int64 {
    for fi5 := 0; fi5 < 18; fi5++ {
        var v6 int16 = 32767
        v7 := []int{fi5}
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := []int{32768, 716}
    var v9 int = 32767
    if false {
        var v10 float64 = (3.14159 - 3.14159)
        var v11 float32 = 1e10
        _ = v10
        _ = v11
    }
    var v12 uint16 = 255
    _ = p0
    _ = p1
    _ = v12
    _ = v8
    _ = v9
    return -1
}

func main() {
    v1 := make(map[string]int)
    var v2 uint16 = 35
    var v3 uint64 = 414
    var v4 uint64 = v3
    v4 = v4
    var v13 int16 = -293
    v13 = (v13 & v13)
    v13 += -32768
    var v14 float32 = 1.5
    if true {
        fmt.Println(v3, v13)
    } else {
        switch "" {
        case "a":
            v15 := GenS1{F0: -2147483648, F1: 255}
            _ = v15
        case "\n\t":
            v16 := []int{2147483647, 32767}
            _ = v16
        }
    }
    v17 := []int{910}
    v14 -= 0.0
    if (!true) {
        v14 += v14
        v18 := v1[(("a" + "hello") + "hello")]
        _ = v18
    } else {
        v4 += v4
        delete(v1, "hello")
        v19 := GenS1{F0: -2147483648, F1: 501}
        _ = v19
    }
    fmt.Println(v4, v13)
    _ = len(v17)
    fmt.Println(v4, v14, v3)
    v20 := GenS1{F0: 2147483647, F1: 255}
    v14 = float32(1.5)
    switch -2147483648 {
    case 694:
        for fi21 := 0; fi21 < 12; fi21++ {
            if true {
                v22 := GenS1{F0: 42, F1: 937}
                v14 = (float32(0.0) * 0.0)
                var v23 bool = true
                _ = v22
                _ = v23
            } else {
                v24 := v1["a"]
                if ((v3 & v3) < v3) {
                    v1[("\n\t" + (("" + ("hello" + "\n\t")) + "\n\t"))] = v24
                } else {
                    fmt.Println(v14, v14, v2)
                    switch 2147483647 {
                    case 88:
                        v25 := GenS2{F0: true, F1: 1}
                        _ = v25
                    case 0:
                        v13 = v13
                    default:
                        if false {
                            fmt.Println(v2, v2, v24)
                            var v26 float64 = 1.5
                            var v27 bool = false
                            _ = v26
                            _ = v27
                        } else {
                            var v28 int64 = 42
                            fmt.Println(v13, v4)
                            if len(v17) > 0 {
                                v17[0] = (v24 - v24)
                            }
                            _ = v28
                        }
                    }
                    v29 := GenS3{F0: -100, F1: v4, F2: ("" + ("hello" + "foo bar"))}
                    _ = v29
                }
                _ = v24
            }
            var v30 float32 = v14
            _ = fi21
            _ = v30
        }
    case -1:
        switch "a" {
        case "hello":
            var v31 uint64 = v3
            _ = v31
        case "foo bar":
            var v32 int64 = genFn0(2147483647, 2)
            _ = v32
        case "":
            fmt.Println(v14)
        }
    case 128:
        if true {
            v33 := GenS4{F0: 117, F1: ("" + "a"), F2: 9223372036854775807}
            v34 := GenS5{F0: v13}
            _ = v33
            _ = v34
        }
    }
    v1[(("\n\t" + "\n\t") + (("a" + "\n\t") + "a"))] = -2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
