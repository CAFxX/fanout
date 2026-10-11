package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int8 }
type GenS2 struct { F0 uint16; F1 uint16; F2 uint8 }
type GenS3 struct { F0 uint64; F1 int8; F2 bool }
type GenS4 struct { F0 bool }

func main() {
    var v1 uint = 65535
    v2 := make(map[int]int)
    v3 := make(map[int]string)
    var v4 int8 = -1
    v5 := GenS1{F0: 42, F1: 0, F2: v4}
    var v6 uint = 4294967295
    fmt.Println(v1, v1, v1)
    var v7 int = -2147483648
    for fi8 := 0; fi8 < 6; fi8++ {
        var v9 int64 = -1
        v10 := GenS2{F0: 783, F1: 230, F2: 35}
        _ = fi8
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        var v12 uint = (v1 << ((1) & 63))
        v13 := GenS1{F0: (v12 | v6), F1: 0, F2: (v4 / ((v4) | 1))}
        v14 := []int{}
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    switch v7 {
    case -2147483648:
        var v15 float64 = float64((3.14159 / 1e10))
        _ = v15
    case 160:
        v16 := v2[v7]
        _ = v16
    default:
        if false {
            if true {
                if (v6 <= v1) {
                    if ("" < "hello") {
                        v17 := GenS3{F0: 2, F1: (-(v4)), F2: false}
                        for fi18 := 0; fi18 < 1; fi18++ {
                            v1 = v1
                            _ = fi18
                        }
                        _ = v17
                    }
                } else {
                    fmt.Println(v6)
                }
                if true {
                    v19 := GenS4{F0: ("hello" > "hello")}
                    v20 := []int{}
                    _ = v19
                    _ = v20
                } else {
                    fmt.Println(v6, v1, v4)
                    v4 += 0
                }
            }
        }
    }
    delete(v2, (v7 - (v7 & (^(v7)))))
    for fi21 := 0; fi21 < 2; fi21++ {
        if false {
            fmt.Println(v7)
            v22 := []int{v7, 128}
            _ = v22
        }
        v4 = v4
        _ = fi21
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
