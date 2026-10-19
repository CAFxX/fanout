package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 uint8 }
type GenS2 struct { F0 int; F1 int64; F2 int }

func main() {
    v1 := make(map[string]int)
    var v2 uint16 = 65535
    var v3 bool = (0 != 2147483647)
    var v4 float32 = 0.1
    var v5 int8 = -87
    var v6 int8 = (v5 + v5)
    v2 = v2
    fmt.Println(v2)
    if v3 {
        v7 := GenS1{F0: true, F1: 32768, F2: 101}
        _ = v7
    }
    _ = len(v1)
    v8 := GenS1{F0: (!v3), F1: 32767, F2: 1}
    var v9 int8 = v5
    fmt.Println(v3, v4, v3)
    if v3 {
        if true {
            if (4294967295 != 2) {
                delete(v1, "a")
                v10 := []int{32767, 42}
                _ = v10
            } else {
                if false {
                    v1["foo bar"] = -147
                    v11 := GenS2{F0: 1, F1: -758, F2: 702}
                    _ = v11
                }
                v9 = 100
                var v12 bool = (-2147483648 <= 167)
                _ = v12
            }
        }
        v13 := GenS2{F0: 2147483647, F1: -671, F2: 0}
        _ = v13
    }
    if v3 {
        var v14 int16 = -203
        v15 := GenS1{F0: (true && v3), F1: -878, F2: 0}
        var v16 uint8 = 255
        _ = v14
        _ = v15
        _ = v16
    }
    if (v3 && v3) {
        delete(v1, "\n\t")
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
