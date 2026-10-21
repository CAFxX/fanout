package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 int16 }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 string }
func genFn0() uint8 {
    var v6 int64 = 1
    switch "\n\t" {
    case "foo bar":
        fmt.Println(v6)
    default:
        var v7 int32 = 361
        _ = v7
    }
    v8 := []int{149, 180, 32768}
    _ = v6
    _ = v8
    return 251
}
func genFn1() uint32 {
    v9 := []int{824, 42, 2147483647}
    var v10 uint64 = 4294967295
    var v11 string = "foo bar"
    _ = v10
    _ = v11
    _ = v9
    return 0
}

func main() {
    var v1 string = ("foo bar" + ("a" + "foo bar"))
    var v2 float32 = -2.25
    var v3 uint16 = 42
    var v4 uint32 = 1
    var v5 bool = ((^(v4)) != v4)
    fmt.Println(v1, v2)
    for fi12 := 0; fi12 < 20; fi12++ {
        for fi13 := 0; fi13 < 19; fi13++ {
            if v5 {
                v2 = 0.0
                switch (fi13 >> ((0) & 63)) {
                case -223:
                    v14 := []int{}
                    _ = v14
                case 312:
                    fmt.Println(v3)
                    v5 = true
                case 93:
                    v15 := []int{(-(fi13))}
                    _ = v15
                default:
                    var v16 string = ""
                    _ = v16
                }
                v1 = v1
            }
            _ = fi13
        }
        _ = fi12
    }
    var v17 float32 = v2
    v1 = v1
    for fi18 := 0; fi18 < 19; fi18++ {
        fmt.Println(v2)
        if v5 {
            var v19 int16 = 42
            fmt.Println(v3, v1, v19)
            _ = v19
        } else {
            v17 = v2
            switch (v1 + v1) {
            case "a":
                v5 = true
            case "":
                var v20 int64 = 9223372036854775807
                _ = v20
            }
            var v21 int32 = 0
            _ = v21
        }
        v2 = (-2.25 * (-2.25 / (1.5 / v2)))
        _ = fi18
    }
    var v22 int8 = -1
    if v5 {
        v22 += -128
    } else {
        if (32767 <= -461) {
            var v23 int16 = 475
            _ = v23
        } else {
            v24 := []int{}
            v25 := GenS1{F0: -1, F1: 457, F2: 1000}
            v5 = v5
            _ = v24
            _ = v25
        }
        v26 := GenS2{F0: 1, F1: 32767}
        _ = v26
    }
    v27 := GenS3{F0: "a"}
    var v28 float32 = 1.5
    var v29 uint16 = v3
    var v30 int32 = 0
    v31 := []int{-9223372036854775808}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v22
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
