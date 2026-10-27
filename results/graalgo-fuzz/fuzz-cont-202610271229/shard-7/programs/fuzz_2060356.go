package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 int64 }
type GenS2 struct { F0 string }

func main() {
    var v1 float64 = 3.14159
    v2 := []int{32768}
    var v3 uint64 = 1
    v3 += v3
    v4 := GenS1{F0: -2147483648, F1: 1, F2: -1}
    if true {
        v5 := GenS1{F0: -1, F1: 42, F2: 0}
        var v6 float64 = 3.14159
        v3 = v3
        _ = v5
        _ = v6
    } else {
        switch "hello" {
        case "":
            v3 = v3
        }
    }
    var v7 bool = true
    var v8 int32 = 1
    if (382 != 0) {
        if (v3 > (v3 >> ((995) & 63))) {
            v9 := GenS1{F0: v8, F1: 255, F2: 1}
            _ = len(v2)
            v10 := GenS2{F0: "foo bar"}
            _ = v10
            _ = v9
        }
    }
    for fi11 := 0; fi11 < 1; fi11++ {
        switch "" {
        case "a":
            switch "hello" {
            case "hello":
                for fi12 := 0; fi12 < 19; fi12++ {
                    v3 -= (v3 + 42)
                    v7 = ((42 >= int8(fi11)) || v7)
                    if len(v2) > 0 {
                        v2[0] = fi12
                    }
                    _ = fi12
                }
            case "":
                v3 = 1
            }
        case "\n\t":
            if (0 < -1) {
                var v13 uint8 = 152
                switch fi11 {
                case 1:
                    switch -1 {
                    case 127:
                        v1 = v1
                    case 4294967295:
                        v13 += (^(v13))
                    case 379:
                        v7 = ((v13 >> ((65535) & 63)) >= v13)
                    default:
                        fmt.Println(fi11)
                    }
                case 128:
                    var v14 int = fi11
                    _ = v14
                case 458:
                    _ = len(v2)
                }
                v13 = 127
                _ = v13
            }
        }
        v3 = 914
        v3 = v3
        _ = fi11
    }
    var v15 float64 = (float64(float32(v1)) * v1)
    fmt.Println(v1, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
