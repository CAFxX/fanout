package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 int }
func genFn0(p0 int64, p1 uint8, p2 int8) int32 {
    for fi4 := 0; fi4 < 10; fi4++ {
        fmt.Println(p0, fi4)
        fmt.Println(fi4, p1, p0)
        _ = fi4
    }
    var v5 uint32 = 2147483648
    if (!true) {
        p2 += (p2 + (p2 % (((p2 & -128)) | 1)))
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return -129
}
func genFn1(p0 int8, p1 int, p2 uint) string {
    p0 -= -1
    var v6 int8 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return ("hello" + ("foo bar" + "hello"))
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]uint32)
    var v3 float32 = 3.14159
    _ = len(v2)
    for fi7 := 0; fi7 < 13; fi7++ {
        switch (fi7 & (fi7 >> ((42) & 63))) {
        case 0:
            v8 := []int{(fi7 * 32768), (fi7 >> ((18446744073709551615) & 63))}
            _ = v8
        default:
            v9 := GenS1{F0: -99, F1: 798}
            _ = v9
        }
        v3 -= ((v3 * ((v3 - v3) - 3.14159)) * v3)
        if false {
            var v10 int8 = 58
            if (826 > 0) {
                if false {
                    switch "\n\t" {
                    case "foo bar":
                        var v11 float64 = 1.5
                        _ = v11
                    case "a":
                        switch fi7 {
                        case 127:
                            v12 := GenS1{F0: v10, F1: -1}
                            _ = v12
                        default:
                            v3 = (v3 * 3.14159)
                        }
                    default:
                        v13 := GenS1{F0: (v10 * 28), F1: -61}
                        _ = v13
                    }
                    v14 := GenS2{F0: 2147483647}
                    _ = v14
                }
            }
            var v15 float32 = v3
            _ = v10
            _ = v15
        }
        _ = fi7
    }
    fmt.Println(v3)
    var v16 uint32 = 1
    var v17 bool = (genFn0(9223372036854775807, 2, 1) == 42)
    var v18 string = "\n\t"
    v18 = "hello"
    v3 = 3.14159
    if (-128 <= -128) {
        var v19 int8 = -128
        var v20 uint32 = 0
        if (924 > 981) {
            switch (v18 + v18) {
            case "":
                v21 := GenS1{F0: (v19 ^ (v19 % ((-44) | 1))), F1: 1}
                _ = v21
            default:
                switch 4294967295 {
                case 9223372036854775807:
                    var v22 uint64 = 309
                    _ = v22
                case 870:
                    if true {
                        v23 := []int{32767, -151, -340}
                        fmt.Println(v3, v20, v16)
                        _ = v23
                    } else {
                        var v24 uint16 = 32767
                        _ = v24
                    }
                case -1:
                    v25 := v2["\n\t"]
                    _ = v25
                }
            }
            v19 = v19
        } else {
            if (0 > 0) {
                v18 = v18
            }
            switch v18 {
            case "\n\t":
                v26 := GenS2{F0: 308}
                _ = v26
            }
            v27 := []int{483}
            _ = v27
        }
        _ = v19
        _ = v20
    } else {
        fmt.Println(v18, v16)
        v3 = ((v3 + (-2.25 / 1.5)) - (((v3 / 1.0) - ((0.1 / 1.0) - v3)) * v3))
        v17 = (v18 != ((("foo bar" + ("a" + "a")) + (v18 + "hello")) + v18))
    }
    v28 := GenS1{F0: 100, F1: -9223372036854775808}
    var v29 uint64 = 428
    if false {
        v30 := []int{-718, -1}
        v29 += (^(v29))
        _ = len(v30)
        _ = v30
    } else {
        v31 := v1[2147483647]
        _ = v31
    }
    if (0 >= 811) {
        v29 += (^(v29))
        if (!(2147483647 > 125)) {
            v3 -= (0.0 - (v3 + v3))
            var v32 int = 42
            switch 4294967295 {
            case 32768:
                var v33 int32 = int32(v16)
                _ = v33
            case 2147483647:
                v34 := GenS2{F0: v32}
                _ = v34
            default:
                for fi35 := 0; fi35 < 2; fi35++ {
                    if v17 {
                        var v36 uint = 1
                        v37 := []int{}
                        _ = v36
                        _ = v37
                    } else {
                        var v38 int16 = -1
                        _ = v38
                    }
                    v39 := GenS1{F0: -103, F1: 883}
                    _ = len(v1)
                    _ = fi35
                    _ = v39
                }
            }
            _ = v32
        }
        v40 := []int{-24, 0}
        _ = v40
    } else {
        _ = len(v1)
        fmt.Println(v18, v16, v29)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v28
    _ = v29
    _ = v3
}
