package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 int; F2 int64 }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 int16; F1 int64; F2 int64 }
func genFn0(p0 uint32, p1 uint16, p2 uint32) int8 {
    fmt.Println(p0, p1)
    fmt.Println(p1, p1, p0)
    v5 := GenS1{F0: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 87
}
func genFn1(p0 int64, p1 int16) uint64 {
    v6 := []int{42, 9223372036854775807, -289, -684}
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 int8 = -128
        var v9 int8 = (v8 % ((v8) | 1))
        for fi10 := 0; fi10 < 10; fi10++ {
            if false {
                if (1e10 > 0.0) {
                    var v11 int32 = 2147483647
                    var v12 int8 = v8
                    v6 = append(v6, -9223372036854775808)
                    _ = v11
                    _ = v12
                }
            } else {
                if (!(79 >= v8)) {
                    v13 := GenS1{F0: (p0 ^ p0)}
                    var v14 uint8 = 127
                    _ = v13
                    _ = v14
                }
                v8 = int8(p0)
                var v15 int64 = 42
                _ = v15
            }
            fmt.Println(fi7, fi7)
            v16 := GenS1{F0: (p0 % (((p0 - -9223372036854775808)) | 1))}
            _ = fi10
            _ = v16
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    p1 += (p1 + (p1 ^ p1))
    _ = p0
    _ = p1
    _ = v6
    return 42
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]int)
    v3 := make(map[int]uint32)
    v4 := make(map[int]bool)
    var v17 uint8 = 1
    fmt.Println(v17)
    if (true && ((v17 & (^(v17))) >= (v17 * (^(v17))))) {
        delete(v3, 32768)
    }
    v17 += 255
    var v18 float64 = 0.0
    var v19 uint64 = 0
    fmt.Println(v18)
    switch 4294967295 {
    case -2147483648:
        var v20 float32 = 3.14159
        _ = v20
    default:
        v17 -= 255
    }
    v21 := []int{-94, -9223372036854775808, 0}
    switch ("" + "foo bar") {
    case "foo bar":
        if true {
            fmt.Println(v19, v17)
            var v22 float64 = v18
            v23 := GenS2{F0: genFn1(42, 42), F1: -9223372036854775808, F2: 0}
            _ = v22
            _ = v23
        } else {
            if ((-32768 != -41) && (!(-1 > 781))) {
                v24 := GenS3{F0: v17, F1: 255}
                var v25 int16 = 82
                for fi26 := 0; fi26 < 19; fi26++ {
                    var v27 bool = false
                    fmt.Println(v27, v19, v25)
                    v27 = v27
                    _ = fi26
                    _ = v27
                }
                _ = v24
                _ = v25
            } else {
                switch "\n\t" {
                case "":
                    v28 := GenS2{F0: (^(v19)), F1: 52, F2: 0}
                    _ = v28
                case "\n\t":
                    v17 = 127
                case "foo bar":
                    switch 4294967295 {
                    case 979:
                        var v29 uint16 = 2
                        _ = v29
                    }
                default:
                    var v30 bool = true
                    _ = v30
                }
                switch ("foo bar" + "foo bar") {
                case "\n\t":
                    fmt.Println(v17, v18, v19)
                case "":
                    for fi31 := 0; fi31 < 8; fi31++ {
                        _ = len(v21)
                        _ = fi31
                    }
                default:
                    fmt.Println(v17)
                }
            }
            var v32 uint = 389
            _ = v32
        }
    case "a":
        switch 9223372036854775807 {
        case 1:
            var v33 float32 = 0.0
            _ = v33
        default:
            switch 9223372036854775807 {
            case 42:
                var v34 float64 = (0.0 / float64(float32(0.1)))
                _ = v34
            default:
                var v35 uint8 = 127
                _ = v35
            }
        }
    case "":
        for fi36 := 0; fi36 < 19; fi36++ {
            var v37 uint16 = 256
            _ = fi36
            _ = v37
        }
    default:
        if ((!(128 > -2147483648)) && true) {
            fmt.Println(v18, v17, v17)
            var v38 uint16 = 255
            v39 := GenS4{F0: 32767, F1: 42, F2: -9223372036854775808}
            _ = v38
            _ = v39
        }
    }
    v40 := v1[0]
    if len(v21) >= 2 {
        v21 = v21[1:len(v21)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v40
}
