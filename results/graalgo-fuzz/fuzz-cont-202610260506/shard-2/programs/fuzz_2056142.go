package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 string }
func genFn0(p0 uint32, p1 string, p2 int32) uint64 {
    var v7 uint32 = (p0 ^ 0)
    v8 := GenS1{F0: 4294967295, F1: 4294967295, F2: p1}
    for fi9 := 0; fi9 < 19; fi9++ {
        var v10 uint32 = (^(v7))
        if (42 >= -1) {
            p2 += -943
            v11 := GenS1{F0: 42, F1: 646, F2: "a"}
            v12 := []int{fi9}
            _ = v11
            _ = v12
        } else {
            switch int(p2) {
            case 741:
                for fi13 := 0; fi13 < 17; fi13++ {
                    v7 += p0
                    _ = fi13
                }
            default:
                v14 := GenS1{F0: 42, F1: 0, F2: "hello"}
                _ = v14
            }
            var v15 float32 = 1.5
            var v16 float64 = -2.25
            _ = v15
            _ = v16
        }
        var v17 uint32 = v10
        _ = fi9
        _ = v10
        _ = v17
    }
    var v18 uint8 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v7
    _ = v8
    return 18446744073709551615
}

func main() {
    var v1 uint32 = 4294967295
    var v2 uint64 = 0
    v3 := []uint32{v1, 716, (v1 + v1)}
    var v4 int = 0
    v5 := make(map[uint32]bool)
    var v6 string = (("hello" + "foo bar") + "a")
    var v19 float32 = ((1.5 - 0.0) + 3.14159)
    v2 = genFn0(v1, "a", -60)
    if false {
        v20 := GenS1{F0: (v2 << ((399) & 63)), F1: (v2 & v2), F2: v6}
        var v21 uint16 = 386
        _ = v20
        _ = v21
    }
    var v22 uint64 = genFn0(v1, "\n\t", 42)
    v23 := GenS1{F0: v2, F1: 4294967295, F2: ""}
    var v24 uint32 = v1
    for fi25 := 0; fi25 < 3; fi25++ {
        v26 := GenS1{F0: (v2 >> ((0) & 63)), F1: v2, F2: "foo bar"}
        _ = fi25
        _ = v26
    }
    if (((v2 - 18446744073709551615) != genFn0(v24, v6, 95)) || false) {
        var v27 string = "a"
        var v28 uint64 = v22
        v29 := GenS1{F0: 18446744073709551615, F1: v28, F2: v6}
        _ = v27
        _ = v28
        _ = v29
    }
    var v30 float32 = (-2.25 * v19)
    fmt.Println(v22, v24, v1)
    v31 := GenS1{F0: 0, F1: (v22 / ((genFn0(411, "a", -306)) | 1)), F2: ((v6 + "a") + v6)}
    v5[998] = true
    var v32 uint = 0
    var v33 float32 = float32(-2.25)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
