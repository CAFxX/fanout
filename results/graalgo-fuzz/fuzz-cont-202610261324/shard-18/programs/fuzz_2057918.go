package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 uint32 }
func genFn0() int {
    var v7 float64 = (-2.25 - 1.5)
    var v8 int64 = -1
    var v9 uint = 403
    v10 := GenS1{F0: 2}
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 32767
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]uint32)
    var v3 bool = false
    var v4 uint32 = 2
    var v5 uint16 = 263
    var v6 int = 128
    var v11 uint8 = 60
    if (v3 || (42 != 880)) {
        v3 = v3
        v12 := GenS1{F0: 1}
        v13 := v2["\n\t"]
        _ = v12
        _ = v13
    } else {
        v11 -= (v11 + v11)
    }
    v14 := GenS2{F0: v3, F1: v6}
    v3 = (!true)
    v6 = v6
    fmt.Println(v11)
    for fi15 := 0; fi15 < 6; fi15++ {
        if v3 {
            if ((!(9223372036854775807 == -1)) && (int8(v11) >= 100)) {
                var v16 int32 = 1
                for fi17 := 0; fi17 < 9; fi17++ {
                    var v18 uint8 = (^(v11))
                    v19 := []int{}
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                _ = v16
            }
        }
        var v20 string = "\n\t"
        _ = fi15
        _ = v20
    }
    if (!v3) {
        switch genFn0() {
        case 42:
            fmt.Println(v4, v4, v6)
        }
    }
    v11 -= v11
    v11 -= 2
    switch "\n\t" {
    case "":
        v21 := GenS3{F0: 216}
        _ = v21
    }
    var v22 int32 = 0
    v23 := GenS2{F0: ("foo bar" > ""), F1: v6}
    v11 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
