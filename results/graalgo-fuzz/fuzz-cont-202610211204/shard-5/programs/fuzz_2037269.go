package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 int16 }
func genFn0() uint {
    var v6 uint64 = 18446744073709551615
    if false {
        v7 := []int{2147483647}
        if len(v7) > 0 {
            v7[0] = 975
        }
        _ = v7
    } else {
        v8 := []int{32768, 2147483647, 128, -981}
        if false {
            if false {
                v6 = v6
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                var v9 float32 = 0.0
                _ = v9
            }
        }
        v10 := GenS1{F0: ("hello" + "a"), F1: (!(88 > uint8(v6)))}
        _ = v10
        _ = v8
    }
    v11 := GenS1{F0: (("\n\t" + "a") + ("hello" + "hello")), F1: (true || true)}
    v12 := GenS2{F0: 32767}
    _ = v11
    _ = v12
    _ = v6
    return 3
}

func main() {
    var v1 int32 = 2147483647
    v2 := make(map[string]bool)
    var v3 uint64 = 803
    v4 := []uint32{}
    var v5 bool = false
    var v13 int64 = 750
    var v14 uint8 = 255
    var v15 float64 = 0.0
    v16 := GenS1{F0: "", F1: (42 == 255)}
    v5 = (42 >= 1)
    var v17 int64 = v13
    v18 := v2["hello"]
    v19 := GenS2{F0: 1000}
    if len(v4) > 0 {
        v4[0] = 2147483648
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
