package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool; F1 uint; F2 bool }

func main() {
    var v1 uint16 = 45
    v2 := make(map[int]int)
    v3 := make(map[uint32]string)
    var v4 int8 = 18
    fmt.Println(v4, v4)
    for fi5 := 0; fi5 < 16; fi5++ {
        switch fi5 {
        case 0:
            var v6 string = ("\n\t" + "")
            _ = v6
        case 128:
            v7 := GenS1{F0: true, F1: false, F2: -1}
            _ = v7
        default:
            if false {
                v1 = (v1 | 42)
            }
        }
        _ = fi5
    }
    v8 := GenS1{F0: (!true), F1: (!false), F2: 992}
    v9 := GenS1{F0: false, F1: (false && true), F2: 1}
    switch "" {
    case "":
        if (false || true) {
            v1 = 42
            var v10 int8 = (v4 * v4)
            _ = v10
        } else {
            v11 := GenS2{F0: v1}
            v12 := GenS1{F0: (!(int64(v4) >= -544)), F1: (-2147483648 >= 213), F2: 42}
            if false {
                switch 0 {
                case 1:
                    if (-142 <= -1) {
                        fmt.Println(v1, v4)
                        v13 := GenS3{F0: (1 >= 2), F1: 651, F2: (636 <= (v1 * (^(v1))))}
                        v14 := []int{0}
                        _ = v13
                        _ = v14
                    } else {
                        _ = len(v2)
                        fmt.Println(v4, v1)
                    }
                case -2147483648:
                    v1 += v1
                }
                fmt.Println(v1, v4)
            } else {
                var v15 uint = 4294967295
                _ = v15
            }
            _ = v11
            _ = v12
        }
    case "foo bar":
        var v16 float64 = -2.25
        _ = v16
    }
    fmt.Println(v1)
    v17 := []int{}
    var v18 bool = (0 == 631)
    v1 = v1
    var v19 int8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
