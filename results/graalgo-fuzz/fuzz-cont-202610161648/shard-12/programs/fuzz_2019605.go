package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint; F1 int32 }
func genFn0(p0 string) int16 {
    var v6 int16 = 995
    v7 := GenS1{F0: 1}
    if false {
        var v8 bool = true
        if (!v8) {
            fmt.Println(p0)
        } else {
            v9 := []int{213}
            switch 32767 {
            case 4294967295:
                if (((1.5 * float64(3.14159)) + 0.1) != 0.1) {
                    switch 9223372036854775807 {
                    case 32768:
                        v10 := GenS1{F0: 2}
                        _ = v10
                    default:
                        var v11 uint16 = 141
                        _ = v11
                    }
                }
            }
            if (3.14159 <= 1.5) {
                v12 := GenS1{F0: 65535}
                _ = v12
            }
            _ = v9
        }
        _ = v8
    }
    _ = p0
    _ = v6
    _ = v7
    return (v6 - (^(v6)))
}

func main() {
    var v1 float64 = 1e10
    v2 := []int{}
    var v3 int64 = -9223372036854775808
    var v4 uint16 = 2
    var v5 string = "foo bar"
    for fi13 := 0; fi13 < 19; fi13++ {
        fmt.Println(v5, v1, v3)
        _ = fi13
    }
    var v14 bool = false
    var v15 int = -216
    v4 = v4
    v16 := GenS2{F0: 848}
    var v17 float64 = 0.0
    for fi18 := 0; fi18 < 15; fi18++ {
        v19 := GenS3{F0: 42}
        var v20 bool = v14
        v4 = v4
        _ = fi18
        _ = v19
        _ = v20
    }
    v21 := GenS4{F0: 255, F1: 622}
    _ = len(v2)
    var v22 uint = 777
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
