package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64 }
func genFn0() uint {
    var v6 string = "a"
    var v7 string = v6
    fmt.Println(v7)
    v8 := GenS1{F0: 167}
    _ = v6
    _ = v7
    _ = v8
    return 822
}

func main() {
    v1 := []uint32{1, 2147483648}
    var v2 int64 = 1
    var v3 uint32 = 811
    var v4 bool = (18446744073709551615 >= 4294967295)
    v5 := []int{42}
    _ = len(v5)
    var v9 uint16 = 2
    var v10 int = -2147483648
    for fi11 := 0; fi11 < 5; fi11++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        var v12 int32 = 696
        v13 := GenS2{F0: 1, F1: 725}
        _ = fi11
        _ = v12
        _ = v13
    }
    v14 := GenS3{F0: 1}
    for fi15 := 0; fi15 < 17; fi15++ {
        v9 -= 0
        v3 = v3
        _ = fi15
    }
    v3 = v3
    v16 := GenS4{F0: 530}
    if (-2.25 > 1e10) {
        if (1 == 127) {
            var v17 int64 = (v2 >> ((42) & 63))
            _ = v17
        }
        switch "\n\t" {
        case "":
            v18 := GenS4{F0: 508}
            _ = v18
        }
        fmt.Println(v9, v4)
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
