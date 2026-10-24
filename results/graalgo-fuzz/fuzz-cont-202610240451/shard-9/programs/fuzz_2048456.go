package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int64) uint16 {
    p0 += p0
    switch -79 {
    case 32767:
        switch ("\n\t" + "a") {
        case "\n\t":
            v6 := GenS1{F0: ""}
            _ = v6
        }
    }
    v7 := GenS2{F0: 183}
    _ = p0
    _ = v7
    return 2
}

func main() {
    v1 := []uint32{2147483647, 357, 0}
    var v2 int8 = 38
    var v3 int64 = 862
    v4 := []uint32{2147483648}
    v5 := []int{32768, 1, 32768}
    if len(v4) > 0 {
        v4[0] = 41
    }
    v3 -= v3
    var v8 float64 = float64(1.5)
    var v9 uint8 = 128
    v9 -= (v9 * 213)
    fmt.Println(v9, v8, v9)
    fmt.Println(v8, v2)
    var v10 uint = 705
    v3 = (v3 - v3)
    var v11 uint64 = 18446744073709551615
    switch "" {
    case "hello":
        if false {
            if true {
                v12 := GenS1{F0: "a"}
                _ = v12
            }
            if len(v1) > 0 {
                v1[0] = 848
            }
            v11 += 373
        } else {
            var v13 int = 128
            v11 = v11
            _ = v13
        }
    case "a":
        var v14 uint = (v10 >> (((v10 + v10)) & 63))
        _ = v14
    }
    fmt.Println(v11)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
