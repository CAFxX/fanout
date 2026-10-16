package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 uint16 }
type GenS2 struct { F0 int16; F1 string; F2 uint64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint8 }
func genFn0(p0 uint) int64 {
    v5 := GenS1{F0: true, F1: 1000, F2: 32767}
    var v6 int32 = 1
    var v7 uint64 = 391
    var v8 string = (("hello" + "foo bar") + "a")
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return -1
}
func genFn1() int64 {
    var v9 int32 = -1
    for fi10 := 0; fi10 < 2; fi10++ {
        fmt.Println(fi10, v9)
        var v11 string = ""
        _ = fi10
        _ = v11
    }
    var v12 uint = 202
    for fi13 := 0; fi13 < 7; fi13++ {
        fmt.Println(fi13)
        if (!false) {
            var v14 float32 = (1.5 / 3.14159)
            var v15 float32 = float32((0.0 - 0.1))
            _ = v14
            _ = v15
        }
        v16 := []int{fi13, fi13, fi13, fi13}
        _ = fi13
        _ = v16
    }
    v12 -= (v12 << (((v12 | v12)) & 63))
    _ = v12
    _ = v9
    return -1
}

func main() {
    var v1 string = ""
    v2 := []int{0, 4294967295}
    v3 := make(map[int]string)
    var v4 int8 = -125
    v17 := GenS2{F0: -178, F1: v1, F2: 4294967295}
    var v18 uint32 = 65535
    v19 := GenS3{F0: -1}
    for fi20 := 0; fi20 < 6; fi20++ {
        v4 += v4
        _ = fi20
    }
    var v21 float64 = 3.14159
    if true {
        for fi22 := 0; fi22 < 10; fi22++ {
            for fi23 := 0; fi23 < 1; fi23++ {
                var v24 int64 = 1
                fmt.Println(v21, fi23, fi23)
                v21 -= 0.0
                if (v24 < genFn1()) {
                    var v25 int16 = -32768
                    var v26 bool = (!(!(!true)))
                    switch v1 {
                    case "hello":
                        var v27 string = v1
                        _ = v27
                    default:
                        if (!(v18 > (v18 - 863))) {
                            var v28 int16 = (v25 * 32767)
                            var v29 float32 = 0.1
                            _ = v28
                            _ = v29
                        }
                    }
                    if len(v2) > 0 {
                        v2[0] = (-(fi22))
                    }
                    _ = v25
                    _ = v26
                }
                v4 = (^(v4))
                if true {
                    v4 += 77
                    v30 := GenS4{F0: 49}
                    _ = v30
                }
                _ = fi23
                _ = v24
            }
            v21 = 0.1
            _ = fi22
        }
    }
    v18 = v18
    var v31 int32 = 0
    v4 -= 0
    v18 -= (v18 >> ((65535) & 63))
    fmt.Println(v4, v21, v18)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v31
    _ = v4
}
