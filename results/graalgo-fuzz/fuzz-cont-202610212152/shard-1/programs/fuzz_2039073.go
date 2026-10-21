package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 uint32; F1 int8 }

func main() {
    v1 := []int{}
    var v2 bool = (32767 > 65535)
    var v3 uint64 = 0
    var v4 int = 0
    v5 := make(map[string]string)
    v6 := make(map[uint32]string)
    var v7 uint32 = 2
    v2 = (-2.25 <= 0.1)
    var v8 uint32 = (^(v7))
    v5["a"] = (("" + ("a" + ("\n\t" + "a"))) + "")
    var v9 int32 = 0
    v10 := GenS1{F0: "hello", F1: 4294967295}
    v11 := GenS1{F0: "a", F1: 65535}
    v3 -= (v3 & (^(v3)))
    for fi12 := 0; fi12 < 14; fi12++ {
        _ = len(v1)
        switch fi12 {
        case 42:
            v7 = (^(v7))
        case -367:
            if v2 {
                var v13 string = ""
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v4 -= (v4 * (v4 + (fi12 + (fi12 + (-(v4))))))
                _ = v13
            } else {
                v14 := GenS1{F0: (("\n\t" + "hello") + "a"), F1: 255}
                fmt.Println(v8, v7)
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = v14
            }
        case 0:
            var v15 uint8 = 168
            _ = v15
        default:
            switch "\n\t" {
            case "":
                var v16 int = (v4 / ((v4) | 1))
                _ = v16
            }
        }
        for fi17 := 0; fi17 < 9; fi17++ {
            var v18 uint = 65535
            v19 := GenS1{F0: "hello", F1: v18}
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = fi12
    }
    v20 := v5["hello"]
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if false {
        for fi21 := 0; fi21 < 2; fi21++ {
            for fi22 := 0; fi22 < 1; fi22++ {
                var v23 uint16 = 902
                var v24 bool = (42 >= -551)
                v25 := GenS1{F0: (v20 + v20), F1: 255}
                _ = fi22
                _ = v23
                _ = v24
                _ = v25
            }
            fmt.Println(v2)
            _ = fi21
        }
    }
    v8 -= (^(v7))
    fmt.Println(v3, v4, v8)
    v26 := GenS1{F0: v20, F1: 1}
    v27 := GenS2{F0: v7, F1: -88}
    for fi28 := 0; fi28 < 12; fi28++ {
        var v29 int8 = 42
        _ = fi28
        _ = v29
    }
    var v30 float32 = 1e10
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
