package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint }
func genFn0(p0 uint, p1 uint64) uint32 {
    var v4 int32 = -2147483648
    var v5 int8 = -114
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 1
}
func genFn1(p0 string, p1 uint16, p2 uint32) int {
    v6 := GenS1{F0: 2}
    var v7 float64 = -2.25
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 42
}

func main() {
    var v1 uint = 345
    var v2 uint32 = 814
    var v3 int = 1
    if false {
        var v8 uint16 = 0
        var v9 float64 = float64(((3.14159 - float32((3.14159 + 1.5))) * float32(float64(-2.25))))
        _ = v8
        _ = v9
    }
    v3 = 167
    switch genFn1("", 255, genFn0(0, 450)) {
    case -2147483648:
        if true {
            v10 := []int{127, v3, -146}
            fmt.Println(v2)
            _ = v10
        }
    case 1:
        fmt.Println(v2, v1)
    }
    v1 = 255
    v1 += (v1 ^ (v1 >> (((^(v1))) & 63)))
    v11 := GenS2{F0: v1}
    for fi12 := 0; fi12 < 17; fi12++ {
        switch "foo bar" {
        case "foo bar":
            v1 = v1
        case "":
            var v13 uint8 = 210
            _ = v13
        }
        _ = fi12
    }
    fmt.Println(v3, v2)
    fmt.Println(v1, v3)
    v14 := []int{}
    var v15 int8 = 127
    v15 = 100
    var v16 int = (v3 << ((42) & 63))
    v17 := GenS2{F0: 18446744073709551615}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
