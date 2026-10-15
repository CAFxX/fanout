package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16; F1 int; F2 int }
func genFn0() uint64 {
    if true {
        for fi5 := 0; fi5 < 11; fi5++ {
            var v6 string = (("" + "\n\t") + "")
            _ = fi5
            _ = v6
        }
        for fi7 := 0; fi7 < 17; fi7++ {
            var v8 uint8 = 0
            v8 = v8
            var v9 uint64 = 2
            _ = fi7
            _ = v8
            _ = v9
        }
    } else {
        for fi10 := 0; fi10 < 18; fi10++ {
            v11 := []int{}
            _ = fi10
            _ = v11
        }
        var v12 int8 = 127
        _ = v12
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        switch "foo bar" {
        case "hello":
            v14 := GenS1{F0: 65535, F1: 42}
            _ = v14
        case "a":
            v15 := []int{(fi13 ^ 4294967295)}
            _ = v15
        }
        var v16 int64 = 1
        var v17 uint32 = 885
        _ = fi13
        _ = v16
        _ = v17
    }
    var v18 string = "foo bar"
    var v19 int8 = 127
    v20 := GenS1{F0: 255, F1: 2}
    _ = v18
    _ = v19
    _ = v20
    return 0
}

func main() {
    var v1 float64 = (-2.25 - float64(1e10))
    var v2 float32 = -2.25
    var v3 float64 = float64(1e10)
    var v4 float64 = v1
    v3 = (v4 + v3)
    v21 := GenS2{F0: -1}
    fmt.Println(v4)
    v22 := []int{0, -9223372036854775808, -562}
    if len(v22) >= 2 {
        v22 = v22[1:len(v22)]
    }
    v23 := GenS3{F0: 255, F1: 0, F2: 506}
    v24 := GenS1{F0: 2, F1: 0}
    v25 := GenS1{F0: 169, F1: 7}
    _ = len(v22)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
