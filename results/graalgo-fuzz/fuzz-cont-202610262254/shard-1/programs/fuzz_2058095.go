package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }

func main() {
    var v1 uint = 208
    v2 := make(map[uint32]int)
    var v3 string = ""
    v4 := []int{-549, 128, 42}
    v2[0] = 128
    fmt.Println(v1, v1)
    if len(v4) > 0 {
        v4[0] = 0
    }
    if ((65535 <= 65535) && (42 != 1)) {
        v1 = (v1 ^ 65535)
        for fi5 := 0; fi5 < 7; fi5++ {
            v1 += (v1 | 255)
            _ = fi5
        }
        v3 = v3
    } else {
        if true {
            if (191 < 1) {
                for fi6 := 0; fi6 < 7; fi6++ {
                    var v7 string = v3
                    _ = fi6
                    _ = v7
                }
            }
            for fi8 := 0; fi8 < 14; fi8++ {
                var v9 uint8 = 1
                v10 := GenS1{F0: "", F1: v9}
                _ = fi8
                _ = v10
                _ = v9
            }
            v3 = v3
        }
    }
    delete(v2, 291)
    _ = len(v4)
    var v11 int64 = -1
    v1 += (v1 & (^(v1)))
    var v12 int8 = 100
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
