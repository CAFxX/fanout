package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 bool; F1 uint }
type GenS4 struct { F0 uint8; F1 uint32; F2 uint16 }
func genFn0() bool {
    var v6 uint8 = 255
    for fi7 := 0; fi7 < 4; fi7++ {
        v8 := []int{(fi7 ^ fi7)}
        _ = fi7
        _ = v8
    }
    v6 = v6
    _ = v6
    return (!false)
}
func genFn1(p0 uint64, p1 uint) int64 {
    p1 = (^(p1))
    switch 2147483647 {
    case -691:
        var v9 uint32 = 625
        _ = v9
    case 0:
        if false {
            v10 := GenS1{F0: 1}
            if false {
                var v11 bool = (p0 > (p0 * (^(p0))))
                fmt.Println(p0, p1, p0)
                _ = v11
            }
            _ = v10
        }
    case -2147483648:
        var v12 int = -2147483648
        _ = v12
    }
    for fi13 := 0; fi13 < 6; fi13++ {
        var v14 float32 = 3.14159
        v14 = v14
        _ = fi13
        _ = v14
    }
    var v15 int = -306
    _ = p0
    _ = p1
    _ = v15
    return -9223372036854775808
}

func main() {
    var v1 uint32 = 2147483647
    var v2 int32 = -2147483648
    var v3 string = "\n\t"
    var v4 uint8 = 219
    v5 := make(map[string]string)
    fmt.Println(v1)
    switch (v3 + "\n\t") {
    case "a":
        v2 = v2
    case "hello":
        v16 := []int{127}
        _ = v16
    }
    var v17 uint32 = (v1 & v1)
    v4 = v4
    switch v3 {
    case "\n\t":
        v18 := v5[v3]
        _ = v18
    case "hello":
        if false {
            if (42 == 1) {
                for fi19 := 0; fi19 < 15; fi19++ {
                    for fi20 := 0; fi20 < 6; fi20++ {
                        v21 := v5[v3]
                        _ = fi20
                        _ = v21
                    }
                    v22 := GenS2{F0: 42, F1: 127}
                    fmt.Println(v3, fi19, v17)
                    _ = fi19
                    _ = v22
                }
            }
            v2 = v2
        }
    case "":
        for fi23 := 0; fi23 < 11; fi23++ {
            var v24 bool = false
            v25 := GenS3{F0: true, F1: 4294967295}
            _ = fi23
            _ = v24
            _ = v25
        }
    default:
        var v26 int32 = v2
        _ = v26
    }
    v27 := GenS4{F0: (v4 | 255), F1: v1, F2: 458}
    v28 := v5[(v3 + v3)]
    v29 := []int{0}
    for fi30 := 0; fi30 < 5; fi30++ {
        if len(v29) >= 2 {
            v29 = v29[1:len(v29)]
        }
        _ = fi30
    }
    var v31 float32 = 1e10
    v32 := GenS2{F0: 2, F1: -98}
    v3 = "hello"
    switch -1 {
    case -2147483648:
        fmt.Println(v31, v1)
    default:
        v3 = v28
    }
    for fi33 := 0; fi33 < 6; fi33++ {
        fmt.Println(v1, v4)
        v28 = ("foo bar" + (v3 + v3))
        _ = fi33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
