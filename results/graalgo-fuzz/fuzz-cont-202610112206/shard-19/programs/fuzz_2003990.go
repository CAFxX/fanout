package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int; F1 int64; F2 uint }
func genFn0(p0 string, p1 int16) uint64 {
    p0 = "hello"
    if true {
        fmt.Println(p1)
        v4 := []int{127, 2147483647, 211}
        _ = v4
    }
    var v5 float64 = (((1.5 / 1.0) + (-2.25 / -2.25)) + (-2.25 + (1.5 / ((-2.25 + 1.5) + float64(1.5)))))
    p1 = p1
    _ = p0
    _ = p1
    _ = v5
    return 1
}
func genFn1(p0 uint64) int8 {
    v6 := GenS1{F0: (!(42 >= 0)), F1: 2147483647}
    v7 := GenS2{F0: 932, F1: 127}
    v8 := GenS2{F0: 2147483647, F1: 100}
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 29
}

func main() {
    v1 := []uint32{110, 218, 0}
    var v2 int8 = -111
    var v3 int16 = -32768
    if (2147483647 > 0) {
        v9 := GenS2{F0: 2, F1: 27}
        if false {
            for fi10 := 0; fi10 < 3; fi10++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                for fi11 := 0; fi11 < 10; fi11++ {
                    if (!((!true) && false)) {
                        var v12 int16 = 1000
                        v1 = append(v1, 0)
                        fmt.Println(v3)
                        v13 := GenS3{F0: true}
                        _ = v12
                        _ = v13
                    } else {
                        if len(v1) > 0 {
                            v1[0] = 42
                        }
                        _ = len(v1)
                        var v14 int8 = 1
                        fmt.Println(fi11)
                        _ = len(v1)
                        _ = v14
                    }
                    v1 = append(v1, 2147483648)
                    _ = fi11
                }
                v2 -= (v2 + 1)
                _ = fi10
            }
            var v15 uint64 = 506
            v16 := GenS4{F0: 788}
            _ = v15
            _ = v16
        } else {
            v17 := GenS5{F0: -276}
            v1 = append(v1, 42)
            _ = v17
        }
        _ = v9
    } else {
        var v18 uint32 = 0
        for fi19 := 0; fi19 < 17; fi19++ {
            fmt.Println(v3)
            _ = fi19
        }
        var v20 uint64 = 18446744073709551615
        _ = v18
        _ = v20
    }
    var v21 int = 2147483647
    var v22 int16 = (v3 & v3)
    var v23 int = 32768
    if (2147483648 < 0) {
        v1 = append(v1, 2)
        for fi24 := 0; fi24 < 7; fi24++ {
            v21 = (v23 | v23)
            fmt.Println(v22, v21, v21)
            _ = fi24
        }
        switch -632 {
        case 4294967295:
            var v25 int16 = 42
            _ = v25
        }
    } else {
        var v26 uint16 = 65535
        for fi27 := 0; fi27 < 1; fi27++ {
            fmt.Println(v22)
            v28 := GenS4{F0: 255}
            _ = fi27
            _ = v28
        }
        var v29 bool = true
        _ = v26
        _ = v29
    }
    for fi30 := 0; fi30 < 19; fi30++ {
        if ((-1 == (v2 / (((v2 & 100)) | 1))) || (!true)) {
            for fi31 := 0; fi31 < 6; fi31++ {
                v21 += fi30
                v32 := GenS6{F0: v21, F1: -421, F2: 131}
                _ = fi31
                _ = v32
            }
        }
        if (!false) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        v33 := GenS3{F0: true}
        _ = fi30
        _ = v33
    }
    var v34 uint64 = 0
    for fi35 := 0; fi35 < 15; fi35++ {
        switch v21 {
        case -9223372036854775808:
            var v36 int64 = 9223372036854775807
            _ = v36
        case 2147483647:
            v1 = append(v1, 2)
        }
        for fi37 := 0; fi37 < 9; fi37++ {
            if true {
                var v38 int8 = -83
                _ = v38
            } else {
                v34 += v34
                v22 -= v22
            }
            v21 += (fi35 + fi37)
            _ = fi37
        }
        switch (^(v21)) {
        case 0:
            var v39 int32 = 2147483647
            _ = v39
        case 32768:
            var v40 uint64 = (v34 >> ((876) & 63))
            _ = v40
        default:
            var v41 float32 = 3.14159
            _ = v41
        }
        _ = fi35
    }
    var v42 bool = true
    v3 = (v22 + 423)
    if len(v1) > 0 {
        v1[0] = 324
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v34
    _ = v42
}
