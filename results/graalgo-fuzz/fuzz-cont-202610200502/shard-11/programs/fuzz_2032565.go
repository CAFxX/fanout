package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 uint32; F1 string; F2 int8 }
type GenS4 struct { F0 uint32; F1 uint32; F2 uint32 }

func main() {
    var v1 float64 = (float64(0.1) - (1.5 + (1.5 + 1e10)))
    var v2 bool = false
    var v3 float32 = ((3.14159 * (1.5 * -2.25)) * float32(0.0))
    if v2 {
        if true {
            v2 = (433 <= 531)
            var v4 float32 = ((1.5 + v3) + v3)
            v5 := []int{32767, 0, 2147483647, 42}
            _ = v4
            _ = v5
        } else {
            var v6 float32 = v3
            _ = v6
        }
        if v2 {
            var v7 int8 = -128
            v3 = 0.1
            _ = v7
        } else {
            if (!(!v2)) {
                v8 := GenS1{F0: -1, F1: 1000}
                fmt.Println(v3)
                v9 := GenS1{F0: 2147483647, F1: -602}
                _ = v8
                _ = v9
            } else {
                if (4294967295 == 2147483647) {
                    var v10 int64 = 9223372036854775807
                    for fi11 := 0; fi11 < 9; fi11++ {
                        var v12 uint = 2
                        v10 += (v10 * (-(v10)))
                        _ = fi11
                        _ = v12
                    }
                    _ = v10
                } else {
                    v13 := []int{169, 127}
                    var v14 float32 = float32(1.5)
                    _ = v13
                    _ = v14
                }
            }
            v15 := GenS2{F0: 1, F1: 2}
            _ = v15
        }
    } else {
        v16 := []int{-337, 32767, -2147483648}
        var v17 float64 = v1
        v18 := GenS3{F0: 258, F1: "hello", F2: 13}
        _ = v16
        _ = v17
        _ = v18
    }
    var v19 uint16 = 510
    fmt.Println(v1)
    v3 = 1e10
    v20 := GenS4{F0: 65535, F1: 0, F2: 92}
    fmt.Println(v19)
    v21 := GenS4{F0: 42, F1: 84, F2: 0}
    fmt.Println(v1)
    v3 = v3
    var v22 int = 122
    if false {
        v23 := GenS1{F0: -1, F1: 32767}
        _ = v23
    } else {
        v24 := []int{v22}
        _ = v24
    }
    fmt.Println(v1, v3)
    v25 := []int{0, 9223372036854775807, 2147483647, v22}
    var v26 int64 = -9223372036854775808
    for fi27 := 0; fi27 < 5; fi27++ {
        var v28 uint16 = v19
        fmt.Println(v28)
        _ = fi27
        _ = v28
    }
    v29 := GenS2{F0: 255, F1: 672}
    var v30 string = ("\n\t" + "")
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v30
}
