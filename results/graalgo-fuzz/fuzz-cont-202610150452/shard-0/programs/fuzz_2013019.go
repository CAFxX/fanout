package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool; F1 int32 }

func main() {
    var v1 int = 128
    v2 := make(map[int]uint32)
    var v3 int32 = -1
    fmt.Println(v3, v1)
    v4 := GenS1{F0: 65535, F1: "\n\t"}
    v1 = (v1 * (v1 * (^(v1))))
    var v5 int64 = -335
    for fi6 := 0; fi6 < 16; fi6++ {
        v7 := GenS1{F0: 255, F1: ("hello" + "hello")}
        _ = fi6
        _ = v7
    }
    v1 = v1
    v3 -= v3
    if (true || true) {
        if true {
            v5 = v5
        } else {
            v8 := []int{(^(v1))}
            for fi9 := 0; fi9 < 12; fi9++ {
                for fi10 := 0; fi10 < 17; fi10++ {
                    v11 := GenS2{F0: v5}
                    var v12 float64 = 3.14159
                    var v13 uint8 = 95
                    _ = fi10
                    _ = v11
                    _ = v12
                    _ = v13
                }
                fmt.Println(v3, v5)
                _ = fi9
            }
            _ = v8
        }
        for fi14 := 0; fi14 < 1; fi14++ {
            switch v1 {
            case -567:
                v15 := GenS3{F0: ((1e10 >= (-2.25 * ((1.5 + 0.1) * (float32(1.5) + 1e10)))) && false), F1: (v3 % (((v3 % ((-670) | 1))) | 1))}
                _ = v15
            case -141:
                var v16 float32 = 3.14159
                _ = v16
            default:
                delete(v2, (fi14 - -742))
            }
            _ = fi14
        }
        v17 := GenS3{F0: false, F1: (v3 / ((-1) | 1))}
        _ = v17
    }
    var v18 uint8 = 103
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
