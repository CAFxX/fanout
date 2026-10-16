package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 bool; F1 uint; F2 uint64 }
type GenS4 struct { F0 uint32; F1 uint64; F2 uint8 }
func genFn0(p0 uint, p1 int, p2 uint8) int16 {
    var v5 int8 = -128
    v6 := GenS1{F0: 42, F1: 1}
    var v7 int32 = -6
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 42
}

func main() {
    v1 := make(map[int]string)
    var v2 bool = true
    var v3 int8 = 42
    var v4 string = "\n\t"
    if v2 {
        v2 = (!((!(!v2)) && ("foo bar" > "")))
        v4 = "hello"
    }
    fmt.Println(v3, v2, v4)
    var v8 uint32 = 1
    switch 846 {
    case 9223372036854775807:
        v9 := GenS1{F0: -237, F1: 42}
        _ = v9
    }
    if (!(656 == 42)) {
        if v2 {
            for fi10 := 0; fi10 < 17; fi10++ {
                var v11 string = v4
                if (!v2) {
                    var v12 uint = 0
                    for fi13 := 0; fi13 < 18; fi13++ {
                        var v14 int = (fi13 | fi10)
                        _ = fi13
                        _ = v14
                    }
                    _ = v12
                }
                _ = fi10
                _ = v11
            }
            v3 -= v3
        }
        if (0 == 255) {
            var v15 uint32 = v8
            v16 := v1[4294967295]
            _ = v15
            _ = v16
        }
    }
    for fi17 := 0; fi17 < 2; fi17++ {
        v18 := GenS2{F0: 1, F1: 42}
        fmt.Println(v2, v4)
        _ = fi17
        _ = v18
    }
    switch -232 {
    case 4294967295:
        v19 := GenS1{F0: 1000, F1: 9223372036854775807}
        _ = v19
    case 916:
        for fi20 := 0; fi20 < 16; fi20++ {
            v21 := []int{fi20}
            _ = fi20
            _ = v21
        }
    }
    var v22 string = v4
    var v23 uint64 = 0
    v22 = v22
    switch "a" {
    case "foo bar":
        var v24 int = 0
        _ = v24
    case "":
        var v25 uint16 = 810
        _ = v25
    case "hello":
        if ((255 != 254) || v2) {
            fmt.Println(v22)
        } else {
            v26 := GenS3{F0: v2, F1: 2, F2: 2}
            _ = v26
        }
    }
    if v2 {
        if v2 {
            switch -1 {
            case 929:
                v1[538] = "a"
            case 4294967295:
                for fi27 := 0; fi27 < 19; fi27++ {
                    _ = len(v1)
                    if (((v3 + v3) == (v3 + (v3 & (v3 / ((100) | 1))))) && false) {
                        var v28 int8 = (v3 * (v3 ^ (v3 + -128)))
                        fmt.Println(v4, v3, v22)
                        _ = v28
                    } else {
                        v29 := GenS3{F0: v2, F1: 736, F2: v23}
                        v22 = v22
                        var v30 int = fi27
                        _ = v29
                        _ = v30
                    }
                    for fi31 := 0; fi31 < 19; fi31++ {
                        var v32 uint64 = (v23 + 991)
                        if (v32 == 697) {
                            _ = len(v1)
                            fmt.Println(v23, v2, fi27)
                        }
                        var v33 bool = v2
                        _ = fi31
                        _ = v32
                        _ = v33
                    }
                    _ = fi27
                }
            default:
                var v34 float32 = (1.5 * -2.25)
                _ = v34
            }
            v22 = (v22 + v4)
            if v2 {
                if v2 {
                    if (307 == 2147483647) {
                        v22 = "\n\t"
                        var v35 uint8 = 231
                        _ = v35
                    } else {
                        v23 = (v23 ^ 18446744073709551615)
                        v36 := []int{-623, 42, 2147483647}
                        if v2 {
                            delete(v1, 9223372036854775807)
                        }
                        _ = v36
                    }
                } else {
                    v2 = (v2 && false)
                }
                v37 := GenS3{F0: v2, F1: 1, F2: v23}
                var v38 int64 = -346
                _ = v37
                _ = v38
            } else {
                for fi39 := 0; fi39 < 18; fi39++ {
                    var v40 uint8 = 102
                    fmt.Println(v8)
                    v41 := GenS1{F0: int16(v8), F1: -937}
                    _ = fi39
                    _ = v40
                    _ = v41
                }
                switch -1 {
                case 42:
                    v42 := []int{}
                    _ = v42
                case -485:
                    v22 = v4
                default:
                    v3 -= (v3 % ((v3) | 1))
                }
            }
        }
        v3 = (v3 << ((0) & 63))
        v43 := GenS4{F0: v8, F1: (v23 / (((v23 - 18446744073709551615)) | 1)), F2: 191}
        _ = v43
    } else {
        fmt.Println(v8, v3)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v8
}
