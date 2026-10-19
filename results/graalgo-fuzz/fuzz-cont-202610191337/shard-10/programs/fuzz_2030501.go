package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 uint64 = 54
    var v2 uint64 = (^(v1))
    var v3 int16 = -215
    v4 := []uint32{}
    v5 := []string{"\n\t"}
    v6 := make(map[int]string)
    v6[32768] = "hello"
    v1 = 458
    v7 := GenS1{F0: 0}
    var v8 int = -1
    v9 := GenS1{F0: -480}
    if true {
        var v10 int32 = 267
        _ = v10
    } else {
        for fi11 := 0; fi11 < 14; fi11++ {
            if (566 <= 42) {
                switch 777 {
                case 42:
                    if len(v4) > 0 {
                        v4[0] = 1
                    }
                case 528:
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
            } else {
                fmt.Println(fi11, v8, v8)
                v12 := GenS1{F0: -97}
                _ = v12
            }
            _ = fi11
        }
    }
    v13 := v6[v8]
    v3 -= v3
    v14 := GenS1{F0: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
