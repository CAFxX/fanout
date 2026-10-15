package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 uint8; F2 int64 }
func genFn0(p0 uint16, p1 uint32, p2 uint) int {
    if (!(!(0 <= 1))) {
        p2 = (p2 * (p2 ^ p2))
        var v6 float64 = -2.25
        _ = v6
    } else {
        var v7 int16 = 0
        _ = v7
    }
    var v8 int = 4294967295
    for fi9 := 0; fi9 < 13; fi9++ {
        v10 := GenS1{F0: 42}
        _ = fi9
        _ = v10
    }
    var v11 int64 = 0
    fmt.Println(v11)
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v8
    return v8
}
func genFn1(p0 int16, p1 int16) bool {
    var v12 uint64 = 677
    for fi13 := 0; fi13 < 16; fi13++ {
        var v14 float32 = (((3.14159 * 0.0) + 0.0) + (((3.14159 * (float32((1e10 - ((((float64((1e10 - float32(-2.25))) / 0.1) + -2.25) - float64((-2.25 - 1e10))) - 3.14159))) + float32(0.0))) / -2.25) - (0.1 + 0.1)))
        if true {
            var v15 float64 = ((1e10 * -2.25) - ((0.0 + (1e10 * (0.1 / float64(v14)))) * 0.1))
            v16 := GenS1{F0: 93}
            var v17 bool = true
            _ = v15
            _ = v16
            _ = v17
        } else {
            if (0 <= 65535) {
                v18 := GenS1{F0: 0}
                v19 := GenS1{F0: 2}
                _ = v18
                _ = v19
            } else {
                var v20 int8 = 4
                var v21 uint64 = v12
                var v22 int16 = p0
                _ = v20
                _ = v21
                _ = v22
            }
            v23 := GenS2{F0: 107, F1: 1, F2: 0}
            if true {
                v24 := []int{fi13}
                v14 += (0.0 - v14)
                _ = v24
            } else {
                if true {
                    var v25 uint16 = 2
                    _ = v25
                }
            }
            _ = v23
        }
        p0 = p1
        _ = fi13
        _ = v14
    }
    v26 := GenS2{F0: -1, F1: 2, F2: 1}
    p1 = (p1 & 1)
    _ = p0
    _ = p1
    _ = v12
    _ = v26
    return (127 == -102)
}

func main() {
    var v1 uint64 = 2
    v2 := make(map[string]bool)
    var v3 uint16 = 65535
    var v4 uint32 = 4294967295
    v5 := make(map[int]bool)
    v1 = v1
    fmt.Println(v3, v3)
    switch -2147483648 {
    case 189:
        var v27 string = ("\n\t" + "foo bar")
        _ = v27
    }
    _ = len(v2)
    var v28 bool = (-9223372036854775808 == -1)
    var v29 uint = 0
    v30 := []int{128, 683, 777, 32767}
    for fi31 := 0; fi31 < 15; fi31++ {
        for fi32 := 0; fi32 < 18; fi32++ {
            var v33 int64 = 9223372036854775807
            v34 := GenS2{F0: v33, F1: 2, F2: v33}
            v35 := GenS2{F0: -9223372036854775808, F1: 0, F2: v33}
            _ = fi32
            _ = v33
            _ = v34
            _ = v35
        }
        _ = fi31
    }
    var v36 int32 = 2147483647
    fmt.Println(v36, v4)
    if len(v30) > 0 {
        v30[0] = 32767
    }
    var v37 int16 = 183
    v28 = (v28 || (true && v28))
    if (0.0 >= 1.5) {
        v38 := GenS2{F0: 0, F1: 42, F2: -807}
        _ = v38
    }
    var v39 bool = false
    var v40 uint64 = v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v36
    _ = v37
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
