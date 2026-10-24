package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 uint16; F1 bool }
func genFn0(p0 uint8) int {
    var v7 bool = false
    v8 := GenS1{F0: true, F1: (v7 && (42 <= 935))}
    var v9 uint = 255
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 513
}

func main() {
    var v1 int32 = -1
    var v2 int32 = v1
    var v3 int64 = 880
    v4 := make(map[uint32]bool)
    var v5 float64 = 3.14159
    var v6 int32 = v1
    switch 9223372036854775807 {
    case -9223372036854775808:
        v10 := []int{734}
        _ = v10
    case -1:
        v11 := []int{32768}
        _ = v11
    case 32767:
        fmt.Println(v2, v2, v5)
    default:
        for fi12 := 0; fi12 < 5; fi12++ {
            if true {
                v3 -= (v3 | (v3 >> ((4294967295) & 63)))
            } else {
                var v13 int = -173
                _ = v13
            }
            var v14 uint16 = 42
            v15 := GenS1{F0: false, F1: true}
            _ = fi12
            _ = v14
            _ = v15
        }
    }
    v16 := v4[55]
    for fi17 := 0; fi17 < 3; fi17++ {
        for fi18 := 0; fi18 < 12; fi18++ {
            var v19 uint8 = 128
            _ = len(v4)
            _ = fi18
            _ = v19
        }
        v20 := []int{genFn0(128)}
        v5 = float64((1e10 * 0.0))
        _ = fi17
        _ = v20
    }
    for fi21 := 0; fi21 < 3; fi21++ {
        fmt.Println(v1)
        v2 = (v2 / ((v1) | 1))
        _ = fi21
    }
    v16 = v16
    v1 += v1
    if (!((v6 * v1) >= v6)) {
        var v22 uint64 = 1
        v2 -= v6
        _ = v22
    }
    fmt.Println(v3, v6, v16)
    if (2 != 18446744073709551615) {
        fmt.Println(v2, v16, v16)
    }
    v23 := []int{genFn0(181), -9223372036854775808}
    if (242 != 900) {
        if v16 {
            for fi24 := 0; fi24 < 6; fi24++ {
                switch ("foo bar" + "\n\t") {
                case "":
                    for fi25 := 0; fi25 < 19; fi25++ {
                        fmt.Println(v16, v5)
                        _ = fi25
                    }
                default:
                    var v26 string = ("foo bar" + "foo bar")
                    _ = v26
                }
                var v27 float32 = (1.5 - 0.0)
                _ = fi24
                _ = v27
            }
            v28 := GenS1{F0: (42 <= 32767), F1: (4294967295 != 302)}
            if ((v6 + -2147483648) > (v6 >> ((255) & 63))) {
                var v29 string = "\n\t"
                for fi30 := 0; fi30 < 8; fi30++ {
                    var v31 float64 = (-2.25 / v5)
                    _ = fi30
                    _ = v31
                }
                _ = v29
            } else {
                for fi32 := 0; fi32 < 19; fi32++ {
                    v33 := GenS1{F0: v16, F1: true}
                    var v34 int = (-(fi32))
                    _ = fi32
                    _ = v33
                    _ = v34
                }
                v35 := GenS2{F0: 255, F1: 247}
                _ = v35
            }
            _ = v28
        } else {
            v23 = append(v23, 32768)
        }
        var v36 uint16 = 2
        v37 := GenS3{F0: v36, F1: (42 >= -4)}
        _ = v36
        _ = v37
    } else {
        var v38 uint16 = 2
        var v39 int32 = (v2 & 0)
        _ = v38
        _ = v39
    }
    v6 = (v6 >> ((2) & 63))
    var v40 int32 = (v2 << ((18446744073709551615) & 63))
    v40 -= 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
