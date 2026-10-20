package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8; F1 bool }
func genFn0() int64 {
    v6 := GenS1{F0: 975, F1: (("\n\t" + "\n\t") + "a")}
    var v7 uint = 42
    v8 := []int{32767, 0}
    _ = v6
    _ = v7
    _ = v8
    return 1
}
func genFn1(p0 uint32) uint16 {
    v9 := GenS2{F0: -835}
    p0 = p0
    fmt.Println(p0)
    _ = p0
    _ = v9
    return 42
}

func main() {
    v1 := []uint32{2147483648}
    var v2 uint32 = 1
    var v3 float64 = -2.25
    var v4 uint = 42
    var v5 int32 = -481
    if (true && (false && true)) {
        if true {
            v3 += ((1e10 * ((v3 * (-2.25 / (float64(float32(v3)) * v3))) + ((v3 + (0.0 + 1.5)) / 1.0))) - v3)
            if true {
                var v10 uint16 = 2
                var v11 int32 = v5
                var v12 uint16 = genFn1(v2)
                _ = v10
                _ = v11
                _ = v12
            }
            var v13 int32 = v5
            _ = v13
        } else {
            if ("a" > "a") {
                fmt.Println(v2)
            }
        }
    } else {
        _ = len(v1)
    }
    var v14 int64 = 0
    v15 := GenS1{F0: 4294967295, F1: ("" + "a")}
    v16 := GenS1{F0: 1, F1: "foo bar"}
    fmt.Println(v5)
    if true {
        var v17 int64 = (v14 << ((v4) & 63))
        fmt.Println(v5)
        var v18 int8 = 42
        _ = v17
        _ = v18
    }
    v5 = v5
    var v19 int8 = 100
    switch "a" {
    case "hello":
        var v20 uint = (v4 ^ v4)
        _ = v20
    }
    var v21 int32 = v5
    switch -736 {
    case 127:
        var v22 int = 964
        _ = v22
    case -523:
        switch int(v14) {
        case 2147483647:
            v23 := GenS3{F0: 104, F1: (42 >= 25)}
            _ = v23
        case 0:
            var v24 bool = ((!false) && (!("foo bar" <= ((("a" + "\n\t") + ("foo bar" + "foo bar")) + "foo bar"))))
            _ = v24
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
