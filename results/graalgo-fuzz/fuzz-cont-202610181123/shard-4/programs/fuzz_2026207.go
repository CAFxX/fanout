package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 bool; F1 uint; F2 uint8 }
type GenS4 struct { F0 uint8; F1 int64; F2 int8 }
type GenS5 struct { F0 uint32 }
func genFn0() int64 {
    if false {
        v6 := []int{}
        _ = v6
    }
    var v7 float64 = (float64((-2.25 / -2.25)) + 1.5)
    v7 += float64(float32((v7 + 1e10)))
    _ = v7
    return 1
}
func genFn1(p0 uint16, p1 int8) int32 {
    v8 := GenS1{F0: 2147483647}
    if (-1 > 180) {
        var v9 int16 = 1
        _ = v9
    } else {
        v10 := []int{1, 42}
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v8
    return 42
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 bool = true
    v3 := make(map[string]bool)
    var v4 int = 128
    v5 := make(map[uint32]int)
    switch 0 {
    case -9223372036854775808:
        v4 = v4
    default:
        fmt.Println(v2, v2)
    }
    v11 := GenS1{F0: -193}
    fmt.Println(v4)
    switch v4 {
    case 578:
        for fi12 := 0; fi12 < 3; fi12++ {
            var v13 float32 = (1e10 * float32(3.14159))
            _ = fi12
            _ = v13
        }
    }
    if v2 {
        v14 := []int{v4, (v4 >> ((18446744073709551615) & 63))}
        v4 = v4
        if (903 != -395) {
            v2 = v2
            v15 := GenS2{F0: -128, F1: 461}
            _ = v15
        }
        _ = v14
    }
    for fi16 := 0; fi16 < 13; fi16++ {
        var v17 uint64 = 669
        v17 = (^(v17))
        v2 = (262 < 18446744073709551615)
        _ = fi16
        _ = v17
    }
    v2 = v2
    v18 := []int{42, v4, -1, v4}
    for fi19 := 0; fi19 < 15; fi19++ {
        if v2 {
            switch "\n\t" {
            case "hello":
                v2 = (((!v2) || v2) || (255 < 2))
            case "":
                if v2 {
                    v4 = (fi19 | (v4 ^ 9223372036854775807))
                    var v20 uint = 371
                    fmt.Println(v4, fi19)
                    _ = v20
                }
            }
            if v2 {
                fmt.Println(v2, fi19)
            }
        }
        if ((genFn1(42, 126) != 2147483647) && v2) {
            if false {
                fmt.Println(fi19, v4)
            } else {
                if v2 {
                    v2 = true
                    if len(v18) >= 2 {
                        v18 = v18[1:len(v18)]
                    }
                } else {
                    _ = len(v18)
                    v2 = v2
                }
                fmt.Println(fi19, v2)
            }
            var v21 int = (fi19 | v4)
            if false {
                switch ("a" + "foo bar") {
                case "":
                    if true {
                        fmt.Println(v21, fi19, v4)
                        v4 -= fi19
                    }
                }
                var v22 int = (fi19 ^ v4)
                v23 := GenS1{F0: 2147483647}
                _ = v22
                _ = v23
            } else {
                fmt.Println(v4)
            }
            _ = v21
        } else {
            var v24 bool = (-1 > -1)
            delete(v1, 2147483647)
            if v24 {
                switch v4 {
                case 2147483647:
                    if (0.1 != 0.0) {
                        v2 = (v4 > 4294967295)
                    }
                default:
                    fmt.Println(fi19)
                }
                var v25 float64 = float64(0.1)
                _ = v25
            }
            _ = v24
        }
        var v26 uint = 2
        _ = fi19
        _ = v26
    }
    switch 32767 {
    case 127:
        v27 := GenS3{F0: v2, F1: 552, F2: 128}
        _ = v27
    case 68:
        switch (v4 & v4) {
        case 127:
            var v28 int32 = 2147483647
            _ = v28
        case 4294967295:
            v29 := GenS2{F0: 42, F1: 1}
            _ = v29
        default:
            fmt.Println(v4, v2)
        }
    }
    v2 = (416 == 367)
    v30 := GenS4{F0: 0, F1: -1, F2: -1}
    v31 := GenS1{F0: 2147483647}
    for fi32 := 0; fi32 < 16; fi32++ {
        var v33 int = 704
        var v34 string = ("a" + "hello")
        var v35 uint16 = 42
        _ = fi32
        _ = v33
        _ = v34
        _ = v35
    }
    var v36 float64 = 0.0
    if (v2 || v2) {
        for fi37 := 0; fi37 < 4; fi37++ {
            var v38 int16 = 1000
            _ = fi37
            _ = v38
        }
        v39 := GenS4{F0: 128, F1: 42, F2: 42}
        var v40 uint16 = 32767
        _ = v39
        _ = v40
    }
    for fi41 := 0; fi41 < 18; fi41++ {
        fmt.Println(v2, v2)
        _ = len(v18)
        v42 := GenS5{F0: 0}
        _ = fi41
        _ = v42
    }
    if v2 {
        for fi43 := 0; fi43 < 17; fi43++ {
            v2 = (false && v2)
            _ = fi43
        }
        for fi44 := 0; fi44 < 13; fi44++ {
            if v2 {
                var v45 uint = 1
                v18 = append(v18, fi44)
                _ = v45
            } else {
                var v46 int64 = 1
                if v2 {
                    fmt.Println(fi44)
                    if v2 {
                        fmt.Println(v4)
                        var v47 bool = (!(v36 > v36))
                        if len(v18) >= 2 {
                            v18 = v18[1:len(v18)]
                        }
                        _ = v47
                    }
                }
                v48 := GenS5{F0: 65535}
                _ = v46
                _ = v48
            }
            _ = fi44
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v18
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v36
    _ = v4
    _ = v5
}
