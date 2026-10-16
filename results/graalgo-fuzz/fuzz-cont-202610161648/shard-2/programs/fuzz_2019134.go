package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 uint }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 int8; F1 uint32; F2 uint8 }
func genFn0(p0 int64, p1 int8, p2 int32) bool {
    v5 := []int{736, -1, 4294967295}
    switch "foo bar" {
    case "":
        p0 += 0
    case "foo bar":
        for fi6 := 0; fi6 < 20; fi6++ {
            _ = len(v5)
            _ = fi6
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return ((-(p0)) <= (p0 << ((0) & 63)))
}
func genFn1(p0 uint16) uint32 {
    switch "\n\t" {
    case "":
        for fi7 := 0; fi7 < 10; fi7++ {
            if false {
                var v8 int = fi7
                _ = v8
            }
            p0 = p0
            p0 = p0
            _ = fi7
        }
    }
    if (!true) {
        var v9 bool = true
        p0 = (p0 + p0)
        _ = v9
    } else {
        p0 += 2
    }
    var v10 float64 = 3.14159
    var v11 uint32 = 328
    _ = p0
    _ = v10
    _ = v11
    return v11
}

func main() {
    var v1 float32 = -2.25
    var v2 int = 127
    v3 := make(map[string]uint32)
    var v4 bool = false
    v12 := GenS1{F0: 4294967295, F1: v2, F2: 983}
    fmt.Println(v2, v2, v4)
    v4 = false
    for fi13 := 0; fi13 < 6; fi13++ {
        v4 = true
        var v14 int16 = 42
        var v15 float64 = 0.0
        _ = fi13
        _ = v14
        _ = v15
    }
    var v16 float32 = (v1 * v1)
    v16 = -2.25
    v16 -= v16
    if v4 {
        fmt.Println(v16)
        fmt.Println(v16, v1, v1)
    }
    var v17 int = 9223372036854775807
    for fi18 := 0; fi18 < 20; fi18++ {
        if v4 {
            if v4 {
                v16 = ((v16 * 1.5) - v1)
            } else {
                for fi19 := 0; fi19 < 10; fi19++ {
                    v20 := GenS2{F0: 688, F1: 0}
                    v21 := []int{v17}
                    _ = fi19
                    _ = v20
                    _ = v21
                }
                var v22 int = 42
                v23 := GenS1{F0: 2, F1: (^(v17)), F2: 0}
                _ = v22
                _ = v23
            }
        } else {
            if (190 != 42) {
                v24 := []int{32767, 1, v17, 9223372036854775807}
                _ = v24
            } else {
                v25 := v3["hello"]
                _ = v25
            }
            fmt.Println(v16)
            v26 := GenS3{F0: 0, F1: 42}
            _ = v26
        }
        switch 9223372036854775807 {
        case 2147483647:
            v27 := []int{921, (fi18 << ((67) & 63)), v2, (-(fi18))}
            _ = v27
        case -9223372036854775808:
            for fi28 := 0; fi28 < 11; fi28++ {
                var v29 float64 = float64((v16 * (v1 - v1)))
                _ = fi28
                _ = v29
            }
        case -254:
            v30 := GenS4{F0: -1, F1: 424, F2: 222}
            _ = v30
        default:
            v31 := []int{}
            _ = v31
        }
        v32 := GenS2{F0: 2, F1: 42}
        _ = fi18
        _ = v32
    }
    var v33 uint16 = 42
    switch (v2 * v17) {
    case 32768:
        for fi34 := 0; fi34 < 3; fi34++ {
            delete(v3, "hello")
            _ = fi34
        }
    case 0:
        var v35 uint64 = 276
        _ = v35
    case 1:
        for fi36 := 0; fi36 < 7; fi36++ {
            v3[("foo bar" + "foo bar")] = genFn1((^(v33)))
            _ = fi36
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v33
    _ = v4
}
