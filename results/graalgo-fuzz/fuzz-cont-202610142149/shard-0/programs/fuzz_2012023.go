package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32; F1 uint16 }

func main() {
    var v1 int16 = 42
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]int)
    v4 := GenS1{F0: v1, F1: -1, F2: 42}
    v5 := []int{}
    v1 += 831
    var v6 int = 4294967295
    v1 = (^(v1))
    var v7 int = v6
    v7 += v6
    var v8 int64 = -197
    var v9 uint16 = 65535
    var v10 uint8 = 55
    v10 = (v10 % ((255) | 1))
    if (1 > 866) {
        for fi11 := 0; fi11 < 16; fi11++ {
            v8 = (v8 - 1)
            _ = fi11
        }
        v1 += (v1 >> ((42) & 63))
        if ((!(438 <= -1)) && (!true)) {
            if len(v5) > 0 {
                v5[0] = 0
            }
            if false {
                if true {
                    switch 453 {
                    case 0:
                        v12 := GenS2{F0: false}
                        _ = v12
                    case 2147483647:
                        fmt.Println(v6)
                    }
                    if true {
                        if false {
                            var v13 uint32 = 65535
                            v1 += v1
                            _ = v13
                        }
                    }
                }
            } else {
                switch "hello" {
                case "foo bar":
                    switch -411 {
                    case 1:
                        fmt.Println(v9, v10, v10)
                    }
                case "a":
                    if (uint16(v10) != v9) {
                        _ = len(v5)
                        v14 := GenS1{F0: v1, F1: -2147483648, F2: v8}
                        _ = v14
                    }
                case "\n\t":
                    for fi15 := 0; fi15 < 3; fi15++ {
                        delete(v3, "\n\t")
                        if false {
                            switch ("\n\t" + "a") {
                            case "\n\t":
                                v5 = append(v5, fi15)
                            }
                        } else {
                            v16 := GenS1{F0: v1, F1: fi15, F2: v8}
                            v9 -= (v9 + (v9 + (v9 >> ((998) & 63))))
                            _ = v16
                        }
                        fmt.Println(v10, v10)
                        _ = fi15
                    }
                }
                v17 := GenS3{F0: 127}
                v1 = (v1 >> ((42) & 63))
                _ = v17
            }
            v1 = 960
        }
    }
    v7 = 42
    v1 += (v1 >> ((964) & 63))
    v7 = -221
    v18 := GenS1{F0: v1, F1: (-(v7)), F2: 206}
    v19 := GenS4{F0: 42, F1: (^(v9))}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
