package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 bool; F2 int }
type GenS3 struct { F0 uint16; F1 uint; F2 uint8 }
func genFn0(p0 uint8, p1 uint8) int32 {
    if true {
        switch 357 {
        case 589:
            v5 := GenS1{F0: p0}
            _ = v5
        default:
            p1 -= p0
        }
    }
    v6 := GenS1{F0: 2}
    v7 := []int{}
    fmt.Println(p1)
    p1 += (^(p0))
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -1
}

func main() {
    var v1 uint64 = 244
    v2 := make(map[int]int)
    var v3 int8 = 1
    v4 := make(map[uint32]uint32)
    if false {
        v1 = (v1 >> ((255) & 63))
        var v8 float32 = 0.0
        for fi9 := 0; fi9 < 16; fi9++ {
            switch fi9 {
            case 32767:
                v1 = (^(v1))
            case -98:
                v8 = float32(1e10)
            case 0:
                fmt.Println(v1)
            default:
                switch 615 {
                case 0:
                    v10 := []int{(-(fi9))}
                    _ = v10
                }
            }
            var v11 uint = 787
            _ = fi9
            _ = v11
        }
        _ = v8
    }
    var v12 bool = false
    if ((v3 | v3) >= (v3 | (v3 & v3))) {
        v13 := GenS1{F0: 229}
        _ = v13
    }
    var v14 string = "hello"
    v15 := []int{2147483647, 246}
    var v16 uint16 = 722
    v17 := GenS1{F0: 1}
    v18 := GenS1{F0: uint8(v1)}
    for fi19 := 0; fi19 < 19; fi19++ {
        for fi20 := 0; fi20 < 12; fi20++ {
            if true {
                v12 = v12
                fmt.Println(fi19, v16)
            } else {
                if (1 > 1) {
                    var v21 float32 = 1.5
                    v3 -= int8(v1)
                    var v22 int = (fi19 >> ((42) & 63))
                    if v12 {
                        var v23 int = v22
                        delete(v2, (^(fi20)))
                        _ = v23
                    } else {
                        v14 = "a"
                        if false {
                            var v24 int8 = (v3 - v3)
                            if len(v15) >= 2 {
                                v15 = v15[1:len(v15)]
                            }
                            var v25 int64 = -9223372036854775808
                            _ = v24
                            _ = v25
                        }
                        var v26 uint64 = (v1 / (((v1 - (v1 * v1))) | 1))
                        _ = v26
                    }
                    _ = v21
                    _ = v22
                } else {
                    fmt.Println(v1, v12, fi20)
                }
                var v27 float64 = float64(3.14159)
                _ = len(v2)
                _ = v27
            }
            fmt.Println(v16, v3)
            v16 = v16
            _ = fi20
        }
        var v28 float32 = 0.0
        var v29 bool = (fi19 <= 32768)
        _ = fi19
        _ = v28
        _ = v29
    }
    v2[0] = 480
    v30 := GenS1{F0: 0}
    _ = len(v15)
    v14 = v14
    if ((v14 != v14) || (0.1 == (1e10 / (0.1 * 0.1)))) {
        switch -9223372036854775808 {
        case 42:
            var v31 float64 = 0.1
            _ = v31
        default:
            var v32 int16 = 1
            _ = v32
        }
    }
    v14 = "foo bar"
    switch -663 {
    case 5:
        for fi33 := 0; fi33 < 20; fi33++ {
            v34 := GenS1{F0: 128}
            v35 := GenS1{F0: 42}
            _ = fi33
            _ = v34
            _ = v35
        }
    case 32768:
        for fi36 := 0; fi36 < 20; fi36++ {
            var v37 float64 = (0.0 - 0.0)
            v38 := GenS2{F0: int16(v3), F1: true, F2: fi36}
            fmt.Println(v12, v14)
            _ = fi36
            _ = v37
            _ = v38
        }
    case 42:
        if true {
            v39 := GenS3{F0: v16, F1: 18446744073709551615, F2: 128}
            _ = v39
        } else {
            v3 = 1
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v30
    _ = v4
}
