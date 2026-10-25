package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64; F1 int32; F2 bool }
type GenS4 struct { F0 uint16; F1 int8; F2 bool }
type GenS5 struct { F0 uint16; F1 uint8; F2 int32 }
type GenS6 struct { F0 int8; F1 int; F2 uint16 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 bool; F1 uint32; F2 string }
func genFn0(p0 uint32) int16 {
    v6 := []int{0}
    var v7 uint64 = 1
    v8 := GenS1{F0: 122}
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 int = -1
    var v2 int16 = 854
    v3 := make(map[int]string)
    var v4 uint = 1
    v5 := make(map[string]int)
    v4 = (v4 >> ((v4) & 63))
    fmt.Println(v4)
    v4 -= v4
    v2 = v2
    var v9 bool = (((-2.25 + 1.5) == 3.14159) || false)
    v10 := GenS1{F0: 238}
    _ = len(v5)
    fmt.Println(v2)
    if v9 {
        if v9 {
            switch v1 {
            case 42:
                var v11 uint16 = 65535
                _ = v11
            case 284:
                fmt.Println(v9, v4, v1)
            case -9223372036854775808:
                v12 := GenS1{F0: 131}
                _ = v12
            default:
                v13 := GenS2{F0: 0}
                _ = v13
            }
            v1 = (v1 & v1)
        } else {
            fmt.Println(v9, v4, v4)
            var v14 bool = (0 >= 100)
            if false {
                v2 -= (-(v2))
                var v15 uint16 = 65535
                _ = v15
            } else {
                if true {
                    v16 := GenS3{F0: 18446744073709551615, F1: 738, F2: v9}
                    v2 += -861
                    _ = len(v5)
                    _ = v16
                } else {
                    switch (v1 * v1) {
                    case -890:
                        v17 := GenS4{F0: 255, F1: 0, F2: v9}
                        _ = v17
                    case 1:
                        var v18 string = "\n\t"
                        _ = v18
                    default:
                        v19 := GenS5{F0: 255, F1: 128, F2: 2147483647}
                        _ = v19
                    }
                    var v20 uint8 = 106
                    _ = v20
                }
            }
            _ = v14
        }
        for fi21 := 0; fi21 < 10; fi21++ {
            v22 := GenS6{F0: 42, F1: (v1 | v1), F2: 597}
            _ = fi21
            _ = v22
        }
        for fi23 := 0; fi23 < 5; fi23++ {
            v24 := []int{-118, (v1 - v1), 32767}
            v24 = append(v24, -2147483648)
            var v25 bool = false
            _ = fi23
            _ = v24
            _ = v25
        }
    } else {
        for fi26 := 0; fi26 < 16; fi26++ {
            v27 := GenS7{F0: 118}
            _ = fi26
            _ = v27
        }
    }
    v28 := GenS7{F0: -121}
    v29 := GenS8{F0: "\n\t"}
    var v30 float32 = 0.1
    var v31 int8 = 42
    var v32 float64 = 0.0
    if v9 {
        var v33 float64 = 3.14159
        if (2147483647 < 4294967295) {
            if ((42 < 255) || v9) {
                v34 := GenS9{F0: 235}
                v9 = ((v9 && ((v9 || v9) && false)) || v9)
                _ = v34
            } else {
                v9 = true
                var v35 uint64 = 838
                _ = v35
            }
        } else {
            if (-2147483648 != 2147483647) {
                for fi36 := 0; fi36 < 7; fi36++ {
                    v37 := []int{32768}
                    var v38 int16 = (v2 & int16(v4))
                    _ = fi36
                    _ = v37
                    _ = v38
                }
            }
        }
        v39 := []int{(v1 * v1), 9223372036854775807, (v1 * v1), v1}
        _ = v33
        _ = v39
    } else {
        var v40 uint64 = 816
        if ((v31 & -1) <= v31) {
            var v41 uint = 640
            var v42 uint = v41
            if (42 >= 2) {
                for fi43 := 0; fi43 < 20; fi43++ {
                    for fi44 := 0; fi44 < 19; fi44++ {
                        var v45 float32 = 0.0
                        fmt.Println(v45)
                        v41 += v41
                        _ = fi44
                        _ = v45
                    }
                    _ = fi43
                }
            } else {
                v46 := GenS10{F0: v9, F1: 135, F2: ""}
                v1 -= v1
                _ = v46
            }
            _ = v41
            _ = v42
        }
        _ = v40
    }
    v32 += (v32 + ((3.14159 / v32) / float64(v30)))
    if v9 {
        var v47 float64 = v32
        v48 := []int{-215, (-(v1)), v1, v1}
        v49 := GenS1{F0: 1}
        _ = v47
        _ = v48
        _ = v49
    } else {
        for fi50 := 0; fi50 < 20; fi50++ {
            v51 := GenS3{F0: 0, F1: 1, F2: false}
            _ = fi50
            _ = v51
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v9
}
