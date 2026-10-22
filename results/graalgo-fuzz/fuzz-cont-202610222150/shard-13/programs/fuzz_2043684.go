package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 int8 }
type GenS2 struct { F0 uint32; F1 int8; F2 int }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint; F1 uint32; F2 uint8 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int16; F1 uint8 }
func genFn0(p0 uint, p1 uint32, p2 uint32) string {
    v4 := GenS1{F0: 183, F1: 42, F2: 100}
    for fi5 := 0; fi5 < 8; fi5++ {
        var v6 uint16 = 114
        v7 := []int{(fi5 + fi5)}
        _ = fi5
        _ = v6
        _ = v7
    }
    p1 = 1
    var v8 uint64 = 0
    var v9 string = "a"
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v8
    _ = v9
    return "\n\t"
}

func main() {
    var v1 uint16 = 312
    v2 := make(map[int]string)
    var v3 float64 = 0.1
    var v10 int32 = 0
    v11 := GenS2{F0: 1, F1: 109, F2: 2147483647}
    var v12 float64 = v3
    v12 += v12
    v3 = float64(3.14159)
    if (!(!true)) {
        if true {
            v13 := GenS3{F0: 9223372036854775807}
            var v14 uint8 = 182
            _ = v13
            _ = v14
        } else {
            v10 += -1
            var v15 float32 = 1e10
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 8; fi16++ {
        for fi17 := 0; fi17 < 17; fi17++ {
            v12 = 1.5
            v18 := GenS4{F0: 316, F1: 2147483647, F2: 84}
            var v19 float64 = v12
            if (718 < 711) {
                fmt.Println(v3, v3)
                v1 = v1
                v20 := GenS5{F0: 1000}
                _ = v20
            } else {
                v2[fi17] = ("foo bar" + "foo bar")
                if false {
                    var v21 int8 = -128
                    if false {
                        var v22 string = ""
                        v21 = (v21 & (v21 & 85))
                        v10 -= -1
                        _ = v22
                    } else {
                        var v23 uint = 42
                        switch "" {
                        case "hello":
                            fmt.Println(v10)
                        case "":
                            v24 := GenS3{F0: 430}
                            _ = v24
                        }
                        _ = v23
                    }
                    _ = v21
                }
            }
            v25 := GenS6{F0: 693, F1: 42}
            _ = fi17
            _ = v18
            _ = v19
            _ = v25
        }
        var v26 int = fi16
        var v27 float32 = -2.25
        _ = fi16
        _ = v26
        _ = v27
    }
    v12 = 1.5
    if true {
        _ = len(v2)
    } else {
        for fi28 := 0; fi28 < 15; fi28++ {
            v1 += (^(v1))
            _ = fi28
        }
    }
    _ = len(v2)
    for fi29 := 0; fi29 < 7; fi29++ {
        var v30 bool = (!(!("foo bar" == "")))
        v31 := []int{fi29}
        var v32 uint16 = v1
        _ = fi29
        _ = v30
        _ = v31
        _ = v32
    }
    fmt.Println(v10, v10)
    v33 := GenS1{F0: 9223372036854775807, F1: 254, F2: 1}
    fmt.Println(v1, v10, v3)
    var v34 uint8 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v33
    _ = v34
}
