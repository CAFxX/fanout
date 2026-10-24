package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 int32; F1 uint64; F2 uint32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int; F1 string; F2 int16 }
type GenS5 struct { F0 int64 }
func genFn0(p0 uint8) int64 {
    if (3.14159 != float32((1e10 - (0.1 / 1.0)))) {
        var v6 int32 = 253
        _ = v6
    } else {
        if (836 <= -32768) {
            var v7 int8 = -1
            v8 := []int{}
            v8 = append(v8, -2147483648)
            _ = v7
            _ = v8
        }
    }
    p0 -= 0
    _ = p0
    return 777
}
func genFn1() int8 {
    var v9 int32 = 42
    switch ("\n\t" + ("hello" + "")) {
    case "hello":
        v10 := []int{-9223372036854775808, 2147483647}
        _ = v10
    default:
        switch 32768 {
        case 32768:
            var v11 float64 = 1.5
            _ = v11
        case 128:
            v9 -= (v9 & 663)
        }
    }
    _ = v9
    return 0
}

func main() {
    var v1 int8 = -60
    v2 := make(map[uint32]bool)
    var v3 uint32 = 821
    var v4 uint8 = 12
    var v5 uint8 = 80
    v12 := GenS1{F0: ("\n\t" + "foo bar"), F1: 65535}
    var v13 uint16 = uint16(v5)
    var v14 uint32 = (v3 + v3)
    fmt.Println(v14, v13)
    switch 159 {
    case 323:
        v15 := GenS2{F0: 1, F1: 974, F2: 65535}
        _ = v15
    case -2147483648:
        v16 := GenS1{F0: "a", F1: 4294967295}
        _ = v16
    case 4294967295:
        var v17 bool = (!true)
        _ = v17
    }
    var v18 int16 = -32768
    v19 := v2[v3]
    v20 := v2[4294967295]
    var v21 uint32 = 430
    v18 = v18
    var v22 uint64 = 42
    v2[v14] = true
    for fi23 := 0; fi23 < 3; fi23++ {
        switch "a" {
        case "\n\t":
            v19 = ((v13 / (((v13 & v13)) | 1)) == uint16(v22))
        default:
            var v24 float64 = 0.1
            _ = v24
        }
        var v25 uint64 = (v22 & v22)
        _ = fi23
        _ = v25
    }
    for fi26 := 0; fi26 < 16; fi26++ {
        v5 += v4
        var v27 string = ""
        for fi28 := 0; fi28 < 8; fi28++ {
            switch "" {
            case "":
                var v29 int16 = int16(v1)
                _ = v29
            case "\n\t":
                v30 := GenS3{F0: (v22 >> ((0) & 63))}
                _ = v30
            case "a":
                var v31 uint64 = uint64(v3)
                v32 := GenS3{F0: v31}
                switch "" {
                case "a":
                    var v33 uint64 = v31
                    _ = v33
                case "":
                    var v34 uint8 = (v5 >> ((uint(v3)) & 63))
                    _ = v34
                case "foo bar":
                    fmt.Println(v3)
                default:
                    fmt.Println(v20, v18, v21)
                }
                if (!(!(0 >= 32767))) {
                    if (9223372036854775807 != -9223372036854775808) {
                        _ = len(v2)
                        v35 := []int{fi26, fi26, 610, -858}
                        if (-128 != (v1 * (^(v1)))) {
                            var v36 int = int(v13)
                            _ = v36
                        } else {
                            fmt.Println(v18)
                            v37 := GenS4{F0: fi26, F1: ((v27 + "foo bar") + "a"), F2: v18}
                            _ = v37
                        }
                        _ = v35
                    } else {
                        if true {
                            v18 = v18
                            fmt.Println(v13, v21, v18)
                            v19 = (-2.25 < 1.5)
                            if v20 {
                                fmt.Println(v14, fi28)
                                fmt.Println(fi28, v1)
                                v27 = v27
                            } else {
                                v2[42] = (v18 >= (v18 ^ v18))
                                v19 = v20
                                var v38 float32 = float32(((0.0 * (float64(1.5) - 3.14159)) + 1.5))
                                _ = v38
                            }
                        }
                        if v20 {
                            v39 := GenS1{F0: v27, F1: uint32(v22)}
                            v20 = true
                            _ = v39
                        } else {
                            if (-1 == -410) {
                                v21 = 515
                            } else {
                                delete(v2, (v14 << ((278) & 63)))
                                v14 -= (v21 << ((42) & 63))
                            }
                            v18 = (v18 >> ((214) & 63))
                        }
                        v40 := GenS2{F0: 2147483647, F1: (v31 / (((v31 * v31)) | 1)), F2: v21}
                        _ = v40
                    }
                    v41 := GenS5{F0: genFn0((v5 / ((v4) | 1)))}
                    _ = v41
                }
                _ = v31
                _ = v32
            }
            _ = fi28
        }
        _ = fi26
        _ = v27
    }
    fmt.Println(v1, v19, v19)
    var v42 int64 = -1
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v42
    _ = v5
}
