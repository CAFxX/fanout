package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 int16; F2 uint }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int32; F1 bool }
func genFn0(p0 int, p1 uint16, p2 int32) uint {
    v5 := GenS1{F0: 138}
    if true {
        if (18446744073709551615 > 2) {
            for fi6 := 0; fi6 < 10; fi6++ {
                v7 := []int{p0, (fi6 >> ((682) & 63)), (p0 + fi6), p0}
                v8 := GenS2{F0: -128, F1: 465, F2: 255}
                _ = fi6
                _ = v7
                _ = v8
            }
            p0 -= (p0 << ((1) & 63))
            if (false || true) {
                p2 += (p2 / (((p2 << ((18446744073709551615) & 63))) | 1))
                v9 := GenS3{F0: (^(p2))}
                _ = v9
            } else {
                var v10 string = ("foo bar" + "")
                p1 = p1
                _ = v10
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 592
}

func main() {
    var v1 int16 = 1000
    v2 := make(map[int]int)
    var v3 int32 = 455
    var v4 uint16 = 443
    switch 2147483647 {
    case 128:
        var v11 float32 = (0.0 * 3.14159)
        _ = v11
    case -266:
        fmt.Println(v4, v3, v3)
    case -766:
        var v12 int = 664
        _ = v12
    default:
        var v13 uint8 = 255
        _ = v13
    }
    var v14 int = 127
    v3 += (v3 | 223)
    fmt.Println(v3, v3)
    switch 870 {
    case 0:
        for fi15 := 0; fi15 < 8; fi15++ {
            v3 += 42
            for fi16 := 0; fi16 < 15; fi16++ {
                switch v14 {
                case 562:
                    v17 := []int{fi15, fi16, 1}
                    _ = v17
                }
                if (true || true) {
                    var v18 uint32 = 65535
                    var v19 uint64 = 0
                    fmt.Println(v19)
                    _ = v18
                    _ = v19
                }
                v4 -= (v4 >> ((862) & 63))
                _ = fi16
            }
            _ = fi15
        }
    case 752:
        _ = len(v2)
    case 1:
        var v20 int = v14
        _ = v20
    }
    if false {
        for fi21 := 0; fi21 < 1; fi21++ {
            for fi22 := 0; fi22 < 2; fi22++ {
                v23 := GenS3{F0: -2147483648}
                var v24 int8 = 100
                var v25 int8 = v24
                _ = fi22
                _ = v23
                _ = v24
                _ = v25
            }
            for fi26 := 0; fi26 < 4; fi26++ {
                v3 += -1
                if true {
                    var v27 uint64 = 1
                    var v28 int64 = int64(v3)
                    _ = v27
                    _ = v28
                } else {
                    v29 := []int{1}
                    _ = v29
                }
                var v30 float64 = 3.14159
                _ = fi26
                _ = v30
            }
            if true {
                var v31 uint8 = 1
                var v32 uint8 = v31
                _ = v31
                _ = v32
            }
            _ = fi21
        }
        v33 := GenS3{F0: v3}
        _ = v33
    } else {
        fmt.Println(v1, v4)
        v34 := GenS4{F0: v3, F1: true}
        var v35 uint = 0
        _ = v34
        _ = v35
    }
    if false {
        var v36 int32 = (v3 * v3)
        _ = v36
    }
    if (!(!(v4 < (v4 ^ 65535)))) {
        fmt.Println(v1, v4, v3)
        if false {
            v37 := GenS1{F0: 890}
            for fi38 := 0; fi38 < 6; fi38++ {
                var v39 uint64 = 0
                v4 = v4
                _ = fi38
                _ = v39
            }
            v40 := GenS1{F0: 982}
            _ = v37
            _ = v40
        } else {
            var v41 uint = uint(v4)
            _ = v41
        }
    } else {
        fmt.Println(v14, v14)
        var v42 uint = 1
        var v43 uint8 = 127
        _ = v42
        _ = v43
    }
    fmt.Println(v3)
    var v44 uint16 = (v4 << ((65535) & 63))
    fmt.Println(v14, v4, v44)
    var v45 uint = 4294967295
    var v46 uint32 = 42
    v2[v14] = 32768
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
}
