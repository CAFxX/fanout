package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }

func main() {
    var v1 int16 = -32768
    var v2 string = "\n\t"
    v3 := make(map[int]string)
    var v4 uint64 = 2
    fmt.Println(v2, v4, v4)
    v5 := []int{42, 27, -106, 128}
    for fi6 := 0; fi6 < 7; fi6++ {
        if true {
            var v7 uint8 = 255
            var v8 bool = false
            for fi9 := 0; fi9 < 1; fi9++ {
                fmt.Println(fi9)
                var v10 int8 = 100
                _ = len(v3)
                _ = fi9
                _ = v10
            }
            _ = v7
            _ = v8
        }
        var v11 uint8 = 128
        switch (fi6 | 32767) {
        case 0:
            _ = len(v5)
        }
        _ = fi6
        _ = v11
    }
    v12 := GenS1{F0: 0, F1: (3.14159 >= float32(1.5))}
    for fi13 := 0; fi13 < 5; fi13++ {
        v14 := GenS1{F0: 672, F1: true}
        _ = fi13
        _ = v14
    }
    var v15 uint8 = 1
    for fi16 := 0; fi16 < 6; fi16++ {
        switch v2 {
        case "hello":
            v17 := GenS1{F0: -41, F1: (!(1.5 >= 0.1))}
            _ = v17
        case "":
            v3[109] = v2
        }
        delete(v3, fi16)
        _ = fi16
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
