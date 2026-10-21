package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 int; F2 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int8; F1 int64 }
func genFn0() string {
    var v5 uint16 = 284
    v6 := GenS1{F0: 127}
    v5 = v5
    v7 := []int{684, -578, 726, 1}
    var v8 uint64 = 4294967295
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return ("\n\t" + ("\n\t" + "a"))
}

func main() {
    var v1 int32 = 554
    v2 := make(map[uint32]int)
    v3 := []string{"\n\t", ("hello" + ""), "foo bar", "a"}
    var v4 int16 = 0
    for fi9 := 0; fi9 < 14; fi9++ {
        v1 = 0
        v10 := GenS2{F0: -791, F1: (fi9 | (fi9 >> ((449) & 63))), F2: -47}
        _ = fi9
        _ = v10
    }
    _ = len(v3)
    var v11 int16 = (v4 & v4)
    for fi12 := 0; fi12 < 19; fi12++ {
        fmt.Println(fi12, fi12)
        _ = fi12
    }
    for fi13 := 0; fi13 < 2; fi13++ {
        var v14 int8 = 100
        if (true || (fi13 == 878)) {
            var v15 int16 = (v11 & v11)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v15
        }
        _ = fi13
        _ = v14
    }
    var v16 uint16 = 0
    var v17 uint32 = uint32(v16)
    fmt.Println(v17, v16)
    var v18 string = "\n\t"
    v19 := GenS3{F0: v18}
    for fi20 := 0; fi20 < 13; fi20++ {
        v21 := GenS4{F0: 0, F1: 6}
        _ = fi20
        _ = v21
    }
    _ = len(v3)
    v18 = v18
    var v22 float64 = -2.25
    v4 = (v4 ^ 42)
    var v23 float64 = v22
    var v24 uint16 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
