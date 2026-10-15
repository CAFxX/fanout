package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 int32; F1 bool; F2 uint16 }
type GenS3 struct { F0 int; F1 int; F2 string }
func genFn0(p0 uint64, p1 string) string {
    v4 := []int{-9223372036854775808}
    v5 := GenS1{F0: 42, F1: 556}
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 bool = (((2 <= 139) || false) && ((0.0 < 1.5) || (true || (!(2147483647 == 987)))))
        var v8 int16 = 42
        var v9 int64 = 1
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    var v10 uint64 = p0
    _ = p0
    _ = p1
    _ = v10
    _ = v4
    _ = v5
    return ((p1 + p1) + p1)
}
func genFn1(p0 int64) bool {
    var v11 int16 = 32767
    var v12 uint64 = 4294967295
    for fi13 := 0; fi13 < 6; fi13++ {
        p0 += p0
        _ = fi13
    }
    if (true || true) {
        v14 := []int{}
        var v15 int16 = (v11 + (v11 - 1))
        v11 += v15
        _ = v14
        _ = v15
    } else {
        v16 := []int{-9223372036854775808, 9223372036854775807, -2147483648, 128}
        p0 += 308
        _ = v16
    }
    _ = p0
    _ = v11
    _ = v12
    return ((((3.14159 / 3.14159) / 1.5) * (-2.25 * (1e10 - 1e10))) <= 0.0)
}

func main() {
    var v1 bool = (!(!false))
    var v2 string = ((("foo bar" + "a") + "\n\t") + "hello")
    var v3 int32 = -431
    v2 = "foo bar"
    if v1 {
        var v17 bool = (false || (!(431 != 422)))
        _ = v17
    }
    v18 := GenS1{F0: 911, F1: 2147483647}
    var v19 uint8 = 0
    if v1 {
        var v20 int64 = 757
        v21 := []int{}
        _ = v20
        _ = v21
    }
    switch 2147483647 {
    case -490:
        var v22 float64 = float64(0.1)
        _ = v22
    }
    for fi23 := 0; fi23 < 9; fi23++ {
        v1 = v1
        switch (v2 + v2) {
        case "\n\t":
            v24 := []int{fi23}
            _ = v24
        case "":
            v25 := GenS2{F0: (v3 / ((v3) | 1)), F1: v1, F2: 0}
            _ = v25
        }
        _ = fi23
    }
    v26 := GenS3{F0: 128, F1: 997, F2: "a"}
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v26
    _ = v3
}
