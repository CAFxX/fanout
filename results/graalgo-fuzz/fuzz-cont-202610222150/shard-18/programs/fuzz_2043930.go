package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint32; F1 uint8; F2 int8 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float32 = (-2.25 - 3.14159)
    var v3 uint32 = 2
    var v4 float32 = v2
    var v5 int16 = 32767
    var v6 uint8 = 2
    var v7 int16 = (v5 & -1)
    if (952 <= 65535) {
        if true {
            v8 := GenS1{F0: (v6 * v6), F1: 100}
            v9 := GenS2{F0: "", F1: 255}
            switch -759 {
            case -9223372036854775808:
                for fi10 := 0; fi10 < 12; fi10++ {
                    v11 := GenS3{F0: true}
                    if true {
                        var v12 int = int(v3)
                        _ = v12
                    } else {
                        v2 = (1e10 - 3.14159)
                        fmt.Println(v3, v6, v6)
                    }
                    _ = fi10
                    _ = v11
                }
            case -1:
                var v13 int = 0
                _ = v13
            }
            _ = v8
            _ = v9
        }
        switch "hello" {
        case "foo bar":
            var v14 float32 = (1.5 + float32(1.5))
            _ = v14
        }
        if true {
            v15 := []int{2147483647, 503, -66}
            _ = v15
        }
    } else {
        var v16 float64 = ((1e10 + (0.1 * (0.0 / (-2.25 - (0.1 * -2.25))))) * 0.0)
        _ = v16
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        v18 := GenS4{F0: (^(v3)), F1: v6, F2: -128}
        fmt.Println(fi17, v2)
        _ = fi17
        _ = v18
    }
    fmt.Println(v2)
    v3 += 0
    var v19 uint64 = 511
    switch (("hello" + "foo bar") + "\n\t") {
    case "hello":
        var v20 uint = 486
        _ = v20
    default:
        var v21 bool = false
        _ = v21
    }
    var v22 uint64 = 916
    v5 += 1
    delete(v1, (v3 * 65535))
    var v23 float64 = 1.5
    v24 := GenS3{F0: true}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
