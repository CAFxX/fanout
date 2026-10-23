package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32 }
func genFn0() uint64 {
    v7 := GenS1{F0: -65}
    var v8 uint64 = 4294967295
    _ = v7
    _ = v8
    return v8
}
func genFn1(p0 uint64) int64 {
    switch -948 {
    case 32767:
        var v9 int8 = 127
        _ = v9
    }
    var v10 string = "a"
    var v11 int = 9223372036854775807
    _ = p0
    _ = v10
    _ = v11
    return 0
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int32 = 351
    var v3 string = ("" + "hello")
    v4 := make(map[string]string)
    var v5 uint64 = 18446744073709551615
    var v6 int16 = -273
    var v12 uint32 = 0
    v5 -= v5
    v13 := GenS2{F0: (v12 >> ((42) & 63))}
    delete(v1, "hello")
    fmt.Println(v2)
    var v14 int32 = v2
    for fi15 := 0; fi15 < 6; fi15++ {
        var v16 float64 = 0.1
        v2 += v2
        switch fi15 {
        case 1:
            for fi17 := 0; fi17 < 19; fi17++ {
                if (1 != 42) {
                    var v18 uint16 = 42
                    v19 := v4[("\n\t" + v3)]
                    v20 := []int{fi15, (-(fi17)), (fi17 - fi17)}
                    _ = v18
                    _ = v19
                    _ = v20
                } else {
                    switch fi15 {
                    case -9223372036854775808:
                        var v21 int16 = (v6 | 904)
                        _ = v21
                    case 32768:
                        v22 := []int{fi17}
                        _ = v22
                    case 42:
                        fmt.Println(v2, v3)
                    }
                }
                _ = fi17
            }
        case 127:
            var v23 uint = 869
            _ = v23
        }
        _ = fi15
        _ = v16
    }
    v6 = v6
    fmt.Println(v3)
    v3 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
