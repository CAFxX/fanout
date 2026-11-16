package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 uint16 }
type GenS2 struct { F0 uint16; F1 string; F2 uint8 }
type GenS3 struct { F0 bool; F1 int16; F2 string }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]int)
    v3 := make(map[uint32]string)
    var v4 int32 = -727
    var v5 uint32 = 593
    v6 := []int{}
    v7 := GenS1{F0: 287, F1: 127, F2: 1}
    var v8 uint8 = 156
    var v9 string = "\n\t"
    if (42 > 355) {
        v8 = (v8 | v8)
        v10 := v2[("foo bar" + v9)]
        fmt.Println(v8)
        _ = v10
    } else {
        v9 = v9
        if true {
            v11 := GenS2{F0: 255, F1: v9, F2: v8}
            for fi12 := 0; fi12 < 14; fi12++ {
                var v13 uint8 = (v8 * (^(v8)))
                var v14 uint16 = 65535
                v9 = (v9 + v9)
                _ = fi12
                _ = v13
                _ = v14
            }
            if false {
                var v15 int = 32767
                var v16 uint32 = (^(v5))
                _ = v15
                _ = v16
            }
            _ = v11
        }
        var v17 uint = 255
        _ = v17
    }
    v18 := GenS2{F0: 883, F1: ("a" + (v9 + "\n\t")), F2: v8}
    v9 = ("hello" + v9)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    switch (v9 + "a") {
    case "a":
        if (!false) {
            var v19 uint8 = v8
            switch -2147483648 {
            case 2147483647:
                v2["hello"] = -497
            case -832:
                v5 += v5
            case 572:
                var v20 float32 = 1.5
                _ = v20
            }
            var v21 uint32 = v5
            _ = v19
            _ = v21
        }
    }
    v22 := GenS3{F0: true, F1: 356, F2: v9}
    var v23 string = v9
    var v24 int8 = -128
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
