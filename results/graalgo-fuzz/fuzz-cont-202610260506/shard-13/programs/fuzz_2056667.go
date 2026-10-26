package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int32; F1 string; F2 int64 }
func genFn0(p0 uint64, p1 uint64) uint {
    v4 := GenS1{F0: 127}
    if true {
        var v5 string = ""
        _ = v5
    } else {
        v6 := []int{127, 796}
        if (3.14159 > 0.1) {
            v7 := GenS1{F0: 127}
            if true {
                v8 := GenS1{F0: -1}
                _ = v8
            }
            _ = v7
        } else {
            if (127 != 25) {
                var v9 int16 = 42
                v10 := GenS2{F0: 0}
                _ = v10
                _ = v9
            }
        }
        _ = v6
    }
    var v11 int32 = 686
    var v12 int8 = 1
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v4
    return 42
}
func genFn1(p0 string, p1 bool) uint32 {
    p0 = ((("hello" + p0) + p0) + "")
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 uint16 = 2
    var v2 int = 32768
    var v3 uint8 = 1
    if false {
        fmt.Println(v2, v2, v3)
        for fi13 := 0; fi13 < 14; fi13++ {
            var v14 float64 = 1.5
            v15 := []int{-2147483648}
            v1 = 42
            _ = fi13
            _ = v14
            _ = v15
        }
        var v16 int = (v2 - v2)
        _ = v16
    }
    var v17 int64 = 9223372036854775807
    for fi18 := 0; fi18 < 4; fi18++ {
        if false {
            fmt.Println(v2, fi18, v1)
            var v19 float32 = 0.0
            _ = v19
        } else {
            for fi20 := 0; fi20 < 3; fi20++ {
                v17 = 9223372036854775807
                var v21 uint32 = 0
                _ = fi20
                _ = v21
            }
        }
        var v22 uint8 = v3
        var v23 int8 = -16
        _ = fi18
        _ = v22
        _ = v23
    }
    v1 -= v1
    v24 := GenS1{F0: 100}
    var v25 uint64 = 0
    var v26 uint = uint(v25)
    var v27 bool = (true || false)
    if (1 == -128) {
        for fi28 := 0; fi28 < 12; fi28++ {
            v29 := GenS3{F0: (4294967295 <= 0)}
            var v30 float64 = ((3.14159 + float64(1.5)) / (0.1 - float64(((((1e10 + ((((((1e10 * ((1e10 + 0.0) * (1.5 + 0.1))) * ((((1.5 * (((-2.25 - 1e10) * 1e10) / 1e10)) - 3.14159) + 1e10) + ((0.0 - 1.5) - -2.25))) + 0.0) + ((((((0.1 + 0.1) / 1.5) - -2.25) * (-2.25 - float32(0.0))) / 1.0) - 0.0)) + -2.25) + 0.1)) - float32(0.1)) * float32(1e10)) / -2.25))))
            _ = fi28
            _ = v29
            _ = v30
        }
        if ((v17 ^ v17) != (v17 >> ((v26) & 63))) {
            var v31 uint16 = 32767
            v32 := []int{42, v2}
            _ = v31
            _ = v32
        } else {
            v33 := []int{42}
            if false {
                for fi34 := 0; fi34 < 15; fi34++ {
                    v35 := GenS3{F0: v27}
                    var v36 bool = (1 >= -208)
                    _ = fi34
                    _ = v35
                    _ = v36
                }
            }
            var v37 uint8 = 35
            _ = v33
            _ = v37
        }
        var v38 uint16 = 1
        _ = v38
    } else {
        for fi39 := 0; fi39 < 9; fi39++ {
            fmt.Println(fi39, v2)
            var v40 uint16 = uint16(v26)
            _ = fi39
            _ = v40
        }
    }
    v41 := GenS4{F0: -1, F1: "a", F2: v17}
    var v42 int32 = -1
    if false {
        var v43 float32 = 0.0
        _ = v43
    }
    fmt.Println(v3, v26)
    var v44 float32 = float32(-2.25)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v17
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v41
    _ = v42
    _ = v44
}
