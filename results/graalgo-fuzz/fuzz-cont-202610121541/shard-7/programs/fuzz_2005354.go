package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 int32 }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int64 }
func genFn0(p0 int32, p1 uint, p2 int16) uint32 {
    var v4 float32 = 1e10
    if (!true) {
        var v5 int64 = 42
        v6 := []int{611, -1, 674}
        _ = v5
        _ = v6
    }
    var v7 uint16 = 812
    var v8 int = 53
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 float64 = ((3.14159 / 1e10) + 1.5)
    var v2 float32 = 0.0
    v3 := []int{192, -2147483648}
    fmt.Println(v2)
    v1 = 3.14159
    if false {
        v9 := GenS1{F0: 1, F1: -2147483648, F2: 0}
        v1 -= 1e10
        _ = v9
    } else {
        for fi10 := 0; fi10 < 13; fi10++ {
            _ = len(v3)
            _ = fi10
        }
        v11 := GenS1{F0: -9223372036854775808, F1: -1, F2: -567}
        v2 = v2
        _ = v11
    }
    var v12 int = 32767
    v13 := GenS2{F0: 4294967295, F1: 978}
    var v14 int8 = 1
    var v15 string = ""
    for fi16 := 0; fi16 < 13; fi16++ {
        v15 = v15
        if (!(!false)) {
            v17 := GenS3{F0: v15}
            if true {
                fmt.Println(v1)
                fmt.Println(v14, fi16, fi16)
                if false {
                    if len(v3) > 0 {
                        v3[0] = v12
                    }
                    v14 = v14
                    v18 := GenS4{F0: 1}
                    _ = v18
                }
            }
            if false {
                if len(v3) > 0 {
                    v3[0] = 41
                }
                for fi19 := 0; fi19 < 6; fi19++ {
                    v14 += v14
                    var v20 uint64 = 2
                    v1 += (v1 * -2.25)
                    _ = fi19
                    _ = v20
                }
            }
            _ = v17
        } else {
            fmt.Println(v12, v14, v14)
        }
        _ = len(v3)
        _ = fi16
    }
    v21 := GenS3{F0: "\n\t"}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v3
}
