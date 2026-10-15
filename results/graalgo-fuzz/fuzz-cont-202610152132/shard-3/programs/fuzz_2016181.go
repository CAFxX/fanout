package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]bool)
    var v3 float64 = 0.0
    var v4 float32 = 0.0
    var v5 int32 = 0
    v6 := make(map[int]uint32)
    v5 -= 2147483647
    v7 := v1["hello"]
    v5 = v5
    if (("\n\t" <= ("" + (("\n\t" + "") + ""))) && v7) {
        switch -312 {
        case -1:
            switch ("\n\t" + "hello") {
            case "a":
                var v8 uint16 = 45
                _ = v8
            case "hello":
                _ = len(v2)
            case "\n\t":
                for fi9 := 0; fi9 < 20; fi9++ {
                    v5 += v5
                    var v10 uint = 255
                    var v11 string = "\n\t"
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            }
        }
        var v12 uint32 = 1
        v5 = v5
        _ = v12
    }
    if false {
        for fi13 := 0; fi13 < 1; fi13++ {
            v14 := []int{(^(fi13)), 32767, fi13}
            _ = fi13
            _ = v14
        }
    }
    switch -554 {
    case 0:
        var v15 uint = 255
        _ = v15
    case 285:
        v16 := GenS1{F0: 395, F1: 1000}
        _ = v16
    case 112:
        for fi17 := 0; fi17 < 12; fi17++ {
            fmt.Println(v7)
            v3 = (v3 + (v3 * v3))
            _ = fi17
        }
    }
    var v18 bool = (1e10 > 1e10)
    var v19 uint64 = 0
    var v20 float32 = 3.14159
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
