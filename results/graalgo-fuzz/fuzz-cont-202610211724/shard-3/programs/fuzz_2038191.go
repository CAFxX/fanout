package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
func genFn0() bool {
    v4 := []int{}
    var v5 float32 = 0.0
    switch ("" + "foo bar") {
    case "a":
        v6 := GenS1{F0: 0, F1: 1000}
        _ = v6
    case "foo bar":
        fmt.Println(v5)
    default:
        var v7 bool = (("hello" + ((("hello" + "\n\t") + ("foo bar" + (("a" + "") + ("hello" + "foo bar")))) + "a")) > "\n\t")
        _ = v7
    }
    var v8 bool = true
    v8 = (false && (!("a" > "a")))
    _ = v4
    _ = v5
    _ = v8
    return true
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]uint32)
    v3 := []int{32768, -645}
    for fi9 := 0; fi9 < 13; fi9++ {
        var v10 float64 = 0.0
        v10 -= v10
        _ = fi9
        _ = v10
    }
    if true {
        var v11 int16 = 1000
        _ = v11
    }
    var v12 int32 = 1
    for fi13 := 0; fi13 < 10; fi13++ {
        var v14 int = fi13
        _ = fi13
        _ = v14
    }
    v15 := GenS1{F0: 0, F1: 1}
    v16 := GenS1{F0: 650, F1: 1000}
    v12 -= v12
    delete(v2, 128)
    if false {
        v12 = v12
        for fi17 := 0; fi17 < 12; fi17++ {
            var v18 float32 = -2.25
            v1[(fi17 / (((fi17 / ((42) | 1))) | 1))] = 428
            _ = fi17
            _ = v18
        }
        var v19 float32 = 1.5
        _ = v19
    }
    for fi20 := 0; fi20 < 13; fi20++ {
        v21 := GenS1{F0: 1, F1: 42}
        switch (fi20 * fi20) {
        case 4294967295:
            switch (fi20 / ((fi20) | 1)) {
            case -805:
                var v22 float64 = -2.25
                _ = v22
            case 32767:
                v23 := GenS1{F0: 467, F1: 32767}
                _ = v23
            }
        case 1:
            if true {
                for fi24 := 0; fi24 < 11; fi24++ {
                    v25 := GenS1{F0: 4294967295, F1: -295}
                    _ = fi24
                    _ = v25
                }
                var v26 int16 = 1000
                var v27 uint32 = 65
                _ = v26
                _ = v27
            }
        }
        _ = fi20
        _ = v21
    }
    var v28 uint8 = 188
    var v29 uint8 = 0
    var v30 uint16 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
}
