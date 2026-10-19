package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 string }
type GenS2 struct { F0 uint32; F1 uint32; F2 uint16 }

func main() {
    v1 := []uint32{775}
    v2 := []string{"\n\t", "hello"}
    var v3 int32 = -1
    var v4 int32 = -2147483648
    var v5 string = ""
    var v6 uint8 = 127
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 uint = 255
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 7; fi9++ {
        v10 := GenS1{F0: (fi9 | (fi9 & fi9)), F1: 945, F2: v5}
        if false {
            v1 = append(v1, 42)
            v11 := GenS1{F0: fi9, F1: -246, F2: ("a" + v5)}
            _ = v11
        }
        for fi12 := 0; fi12 < 7; fi12++ {
            var v13 bool = true
            switch (^(fi9)) {
            case -842:
                v5 = (v5 + v5)
            case -175:
                var v14 int16 = 32767
                _ = v14
            case -919:
                fmt.Println(v6)
            }
            v15 := GenS1{F0: (fi9 / ((fi9) | 1)), F1: (^(v4)), F2: ""}
            _ = fi12
            _ = v13
            _ = v15
        }
        _ = fi9
        _ = v10
    }
    for fi16 := 0; fi16 < 9; fi16++ {
        fmt.Println(v6)
        var v17 int8 = 100
        _ = fi16
        _ = v17
    }
    var v18 uint8 = 204
    v19 := GenS2{F0: 65535, F1: 0, F2: 2}
    var v20 int16 = -1
    var v21 float32 = (((-2.25 - 1e10) / 1.0) - 1e10)
    v18 += v6
    for fi22 := 0; fi22 < 17; fi22++ {
        for fi23 := 0; fi23 < 20; fi23++ {
            if len(v2) > 0 {
                v2[0] = (v5 + "")
            }
            v24 := GenS1{F0: fi22, F1: (^(v4)), F2: v5}
            _ = fi23
            _ = v24
        }
        v21 += v21
        _ = fi22
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
