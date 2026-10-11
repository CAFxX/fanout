package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 int16 }
type GenS2 struct { F0 int32; F1 int32; F2 uint32 }

func main() {
    v1 := make(map[int]string)
    var v2 float32 = -2.25
    v3 := make(map[uint32]int)
    var v4 int64 = -9223372036854775808
    var v5 int8 = 42
    v6 := make(map[int]int)
    var v7 bool = false
    v5 -= (v5 | (v5 * v5))
    v7 = v7
    var v8 uint = 1
    v2 -= v2
    if ((229 != 2) && false) {
        v2 += (3.14159 + float32((3.14159 / (1e10 / float64((v2 * v2))))))
        v7 = v7
        var v9 int16 = 1000
        _ = v9
    } else {
        var v10 int32 = 137
        switch 4294967295 {
        case -2147483648:
            for fi11 := 0; fi11 < 15; fi11++ {
                fmt.Println(v2)
                fmt.Println(v5, v10, v5)
                _ = fi11
            }
        case 761:
            var v12 int16 = -432
            _ = v12
        case 682:
            var v13 uint64 = 4294967295
            _ = v13
        default:
            var v14 int64 = v4
            _ = v14
        }
        v7 = false
        _ = v10
    }
    switch -1 {
    case -139:
        v2 += (float32(0.0) + float32(0.1))
    case 4294967295:
        for fi15 := 0; fi15 < 20; fi15++ {
            v16 := GenS1{F0: 221, F1: 4294967295, F2: -1}
            _ = fi15
            _ = v16
        }
    case -78:
        switch "\n\t" {
        case "":
            for fi17 := 0; fi17 < 11; fi17++ {
                var v18 uint32 = 4294967295
                v8 -= v8
                _ = fi17
                _ = v18
            }
        case "a":
            _ = len(v6)
        }
    default:
        var v19 float64 = 0.1
        _ = v19
    }
    var v20 uint64 = 4294967295
    for fi21 := 0; fi21 < 11; fi21++ {
        v2 = (v2 + v2)
        switch ("" + (("" + "\n\t") + "foo bar")) {
        case "hello":
            v5 -= (v5 & 127)
        }
        delete(v1, fi21)
        _ = fi21
    }
    v20 -= v20
    fmt.Println(v8, v2)
    for fi22 := 0; fi22 < 12; fi22++ {
        v23 := v1[(fi22 + fi22)]
        if v7 {
            var v24 float64 = -2.25
            if false {
                var v25 int32 = -2147483648
                fmt.Println(v5)
                fmt.Println(v24, v4)
                _ = v25
            }
            for fi26 := 0; fi26 < 18; fi26++ {
                _ = len(v1)
                var v27 int32 = int32(v8)
                v28 := []int{}
                _ = fi26
                _ = v27
                _ = v28
            }
            _ = v24
        } else {
            if (527 <= -2147483648) {
                v4 = 95
                if v7 {
                    v7 = (2147483647 < 434)
                    for fi29 := 0; fi29 < 16; fi29++ {
                        v2 = v2
                        var v30 uint32 = 4294967295
                        _ = fi29
                        _ = v30
                    }
                    var v31 uint64 = 927
                    _ = v31
                }
                v32 := []int{int(v5), (fi22 - fi22), (^(fi22)), fi22}
                _ = v32
            }
            v33 := GenS1{F0: 127, F1: (v20 >> ((42) & 63)), F2: 0}
            v34 := GenS2{F0: 139, F1: -1, F2: 65535}
            _ = v33
            _ = v34
        }
        v8 = 65535
        _ = fi22
        _ = v23
    }
    if v7 {
        v8 = 1
        var v35 uint8 = 255
        _ = v35
    }
    if true {
        for fi36 := 0; fi36 < 3; fi36++ {
            v5 = 85
            for fi37 := 0; fi37 < 11; fi37++ {
                if true {
                    v38 := GenS2{F0: 0, F1: 491, F2: 2147483648}
                    v6[(fi37 >> ((0) & 63))] = fi37
                    _ = v38
                } else {
                    var v39 uint8 = 0
                    v8 = (^(v8))
                    v40 := GenS2{F0: -2147483648, F1: -1, F2: 741}
                    v4 -= v4
                    _ = v39
                    _ = v40
                }
                v4 -= v4
                _ = fi37
            }
            _ = fi36
        }
        var v41 float64 = 3.14159
        v20 += (v20 & (v20 << ((42) & 63)))
        _ = v41
    }
    for fi42 := 0; fi42 < 3; fi42++ {
        var v43 uint = 553
        var v44 int32 = 457
        _ = fi42
        _ = v43
        _ = v44
    }
    v20 = v20
    v45 := v3[0]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
