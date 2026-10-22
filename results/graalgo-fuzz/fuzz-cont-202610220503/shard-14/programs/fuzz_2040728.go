package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0() bool {
    v7 := GenS1{F0: 124}
    var v8 bool = true
    var v9 uint = 0
    v9 = v9
    _ = v7
    _ = v8
    _ = v9
    return v8
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]uint32)
    var v3 uint16 = 2
    var v4 int8 = 42
    var v5 int32 = 285
    var v6 int = -948
    fmt.Println(v3, v5, v5)
    v10 := []int{v6}
    var v11 uint16 = v3
    var v12 uint64 = 42
    var v13 uint64 = v12
    v14 := GenS1{F0: 10}
    for fi15 := 0; fi15 < 19; fi15++ {
        v12 += (v12 ^ (v12 % (((v12 | v13)) | 1)))
        _ = fi15
    }
    fmt.Println(v6)
    var v16 uint = 2
    for fi17 := 0; fi17 < 1; fi17++ {
        for fi18 := 0; fi18 < 10; fi18++ {
            if true {
                _ = len(v10)
                _ = len(v2)
                if false {
                    v19 := GenS1{F0: (v4 % ((v4) | 1))}
                    _ = v19
                } else {
                    v13 -= v13
                    fmt.Println(v6, v5)
                    var v20 uint8 = 99
                    _ = v20
                }
            }
            _ = fi18
        }
        v4 += (v4 >> (((v16 >> ((v16) & 63))) & 63))
        _ = fi17
    }
    switch (-(v6)) {
    case -61:
        v5 = (v5 + (v5 % ((v5) | 1)))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
