package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 int16; F1 int8; F2 uint16 }
type GenS4 struct { F0 uint; F1 int8 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 bool; F1 int32 }
func genFn0() bool {
    v6 := []int{4294967295, -1}
    var v7 uint64 = 18446744073709551615
    v7 += 2
    v6 = append(v6, 127)
    _ = v6
    _ = v7
    return (3.14159 >= (0.1 + (1.5 - (-2.25 / (3.14159 / 1e10)))))
}
func genFn1() bool {
    var v8 uint16 = 2
    switch int(v8) {
    case -1:
        switch 32768 {
        case 92:
            v9 := GenS1{F0: 0, F1: -61}
            _ = v9
        }
    }
    _ = v8
    return true
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint = 560
    v3 := make(map[int]bool)
    var v4 int8 = -98
    var v5 bool = true
    if ((v5 || true) || v5) {
        var v10 uint16 = 65535
        fmt.Println(v4, v5)
        var v11 float64 = 3.14159
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 15; fi12++ {
        for fi13 := 0; fi13 < 8; fi13++ {
            switch (fi12 * 9223372036854775807) {
            case 0:
                v14 := GenS1{F0: v2, F1: 0}
                _ = v14
            case 2147483647:
                v15 := v1[895]
                _ = v15
            case -849:
                v5 = false
            }
            v2 += (v2 << (((v2 | (v2 ^ v2))) & 63))
            switch (fi12 | fi13) {
            case 32768:
                var v16 uint64 = 1
                _ = v16
            case 9223372036854775807:
                if (!v5) {
                    v17 := GenS2{F0: 0, F1: 42}
                    v18 := GenS3{F0: -32768, F1: -128, F2: 368}
                    v19 := GenS4{F0: 55, F1: v4}
                    _ = v17
                    _ = v18
                    _ = v19
                } else {
                    v2 -= v2
                    v20 := GenS5{F0: 1}
                    if false {
                        v21 := GenS5{F0: 1}
                        switch fi13 {
                        case 128:
                            var v22 int64 = 0
                            _ = v22
                        case 32768:
                            fmt.Println(v2)
                        }
                        _ = v21
                    } else {
                        if v5 {
                            if v5 {
                                var v23 uint8 = 18
                                _ = v23
                            }
                        }
                    }
                    _ = v20
                }
            case 441:
                if v5 {
                    switch -527 {
                    case 2147483647:
                        v2 = (v2 % (((v2 & 2)) | 1))
                    default:
                        v2 = (v2 | (v2 + 42))
                    }
                    fmt.Println(fi13, fi12, fi12)
                    v5 = (v2 <= 2)
                } else {
                    if ("a" > "a") {
                        v4 = (v4 / (((v4 | 74)) | 1))
                        var v24 int32 = 1
                        _ = v24
                    } else {
                        v25 := GenS2{F0: 2147483647, F1: 42}
                        v26 := GenS4{F0: v2, F1: v4}
                        _ = v25
                        _ = v26
                    }
                    if v5 {
                        v5 = v5
                        if ((-2.25 * -2.25) == 0.1) {
                            var v27 uint8 = 0
                            _ = v27
                        } else {
                            v4 = -1
                            v28 := GenS3{F0: 0, F1: (v4 << ((v2) & 63)), F2: 153}
                            fmt.Println(fi13, fi12)
                            _ = v28
                        }
                        if (65535 >= 919) {
                            v2 = v2
                            fmt.Println(v2)
                            var v29 uint = v2
                            v30 := []int{(fi13 + fi12)}
                            v4 += (v4 >> ((255) & 63))
                            _ = v29
                            _ = v30
                        } else {
                            v4 -= -1
                            v4 = v4
                        }
                    }
                    v5 = (!false)
                    v3[fi12] = true
                }
            }
            _ = fi13
        }
        v4 -= (v4 ^ (v4 / ((v4) | 1)))
        v31 := GenS6{F0: (!v5), F1: 1}
        _ = fi12
        _ = v31
    }
    if v5 {
        v4 = v4
    }
    var v32 int64 = -369
    var v33 int32 = 42
    v4 = 117
    v2 = v2
    v2 = (v2 - 1)
    if ((1 <= v4) && (18446744073709551615 <= 4294967295)) {
        if v5 {
            v33 += v33
        }
        v34 := GenS5{F0: v32}
        switch -9223372036854775808 {
        case 375:
            v35 := []int{}
            _ = v35
        default:
            if true {
                for fi36 := 0; fi36 < 11; fi36++ {
                    fmt.Println(v4)
                    for fi37 := 0; fi37 < 10; fi37++ {
                        fmt.Println(v5)
                        v2 = 741
                        v32 += -978
                        _ = fi37
                    }
                    _ = fi36
                }
                var v38 float32 = 1.5
                _ = v38
            }
        }
        _ = v34
    } else {
        var v39 int16 = -1
        _ = v39
    }
    switch 4294967295 {
    case 4294967295:
        v40 := []int{}
        _ = v40
    case 2147483647:
        var v41 uint16 = 1
        _ = v41
    case -1:
        v2 += v2
    default:
        v42 := GenS5{F0: v32}
        _ = v42
    }
    v32 += v32
    for fi43 := 0; fi43 < 5; fi43++ {
        var v44 uint64 = 1
        _ = fi43
        _ = v44
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
