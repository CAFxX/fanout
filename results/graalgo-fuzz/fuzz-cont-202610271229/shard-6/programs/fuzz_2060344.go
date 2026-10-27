package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 int) int16 {
    p0 = p0
    v5 := []int{(p0 / (((p0 ^ p0)) | 1)), p0, (p0 - 69)}
    _ = p0
    _ = v5
    return 465
}
func genFn1(p0 uint) uint8 {
    var v6 string = (("" + ("\n\t" + "hello")) + "foo bar")
    v7 := []int{128}
    v8 := GenS1{F0: 1}
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 135
}

func main() {
    var v1 uint = 2
    v2 := []int{}
    var v3 float32 = 0.1
    v4 := make(map[uint32]string)
    v1 -= v1
    for fi9 := 0; fi9 < 6; fi9++ {
        v10 := v4[353]
        var v11 int16 = 1000
        for fi12 := 0; fi12 < 19; fi12++ {
            _ = len(v2)
            var v13 int = int(v1)
            _ = fi12
            _ = v13
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    v14 := GenS1{F0: 1}
    v3 = 0.1
    if false {
        v3 += (v3 * -2.25)
        v3 = v3
    }
    var v15 uint8 = 255
    if false {
        var v16 int64 = 375
        _ = v16
    } else {
        if (-16 == -128) {
            v17 := GenS1{F0: uint32(v15)}
            _ = v17
        } else {
            var v18 float32 = v3
            if ((!true) && false) {
                _ = len(v2)
            } else {
                var v19 uint = v1
                v20 := GenS2{F0: 54}
                _ = v19
                _ = v20
            }
            _ = v18
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v21 := GenS1{F0: 222}
    var v22 uint16 = 105
    var v23 string = ""
    v23 = (v23 + v23)
    fmt.Println(v1)
    v15 -= v15
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
