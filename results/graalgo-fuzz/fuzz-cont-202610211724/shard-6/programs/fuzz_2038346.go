package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 uint8; F1 string }
func genFn0(p0 string, p1 int64, p2 uint32) int {
    if false {
        p1 = -1
        var v5 uint16 = 807
        _ = v5
    } else {
        if false {
            fmt.Println(p1)
        }
        v6 := GenS1{F0: 32768, F1: 0}
        _ = v6
    }
    v7 := []int{259, 42, 4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 128
}

func main() {
    var v1 float32 = 1.5
    v2 := make(map[string]int)
    var v3 uint32 = 441
    var v4 int8 = 1
    var v8 uint = 0
    for fi9 := 0; fi9 < 16; fi9++ {
        var v10 int8 = v4
        _ = fi9
        _ = v10
    }
    if false {
        v8 = (v8 & 0)
    } else {
        if (!(true || (((v1 + v1) / v1) >= v1))) {
            v1 = v1
            v8 = (v8 * v8)
        } else {
            v11 := []int{-2147483648}
            var v12 string = ("hello" + ("\n\t" + "\n\t"))
            var v13 float64 = float64((v1 + ((v1 / (v1 * (v1 - v1))) + (v1 - ((3.14159 * v1) * v1)))))
            _ = v11
            _ = v12
            _ = v13
        }
        switch 32768 {
        case -9223372036854775808:
            v1 = v1
        default:
            v14 := GenS1{F0: int(v4), F1: (v8 - v8)}
            _ = v14
        }
        v3 -= (v3 & v3)
    }
    var v15 int = -302
    if ((v4 | v4) != v4) {
        var v16 int = v15
        for fi17 := 0; fi17 < 17; fi17++ {
            v1 = 3.14159
            v18 := v2["foo bar"]
            _ = fi17
            _ = v18
        }
        for fi19 := 0; fi19 < 19; fi19++ {
            switch 1 {
            case 0:
                if ((v8 & v8) > (^(v8))) {
                    v20 := GenS1{F0: 128, F1: (v8 / (((v8 / (((v8 + (v8 & v8))) | 1))) | 1))}
                    _ = v20
                }
            default:
                var v21 int = (fi19 / (((v16 * 0)) | 1))
                _ = v21
            }
            fmt.Println(v16, v8, v15)
            _ = fi19
        }
        _ = v16
    } else {
        if ((!((v3 / ((42) | 1)) > (^(v3)))) || false) {
            fmt.Println(v8, v3, v4)
        }
    }
    fmt.Println(v4, v1)
    v3 = v3
    for fi22 := 0; fi22 < 19; fi22++ {
        if true {
            if false {
                for fi23 := 0; fi23 < 16; fi23++ {
                    fmt.Println(v15, fi22)
                    fmt.Println(v1, v3, v3)
                    var v24 float32 = (((v1 * v1) - (v1 * (v1 - 0.0))) * float32((-2.25 - (1e10 + 1e10))))
                    _ = fi23
                    _ = v24
                }
                var v25 uint16 = 145
                v8 = (^(v8))
                _ = v25
            } else {
                v8 -= v8
            }
            var v26 uint16 = 981
            _ = v26
        }
        _ = fi22
    }
    fmt.Println(v8, v1)
    for fi27 := 0; fi27 < 6; fi27++ {
        if false {
            if (42 >= 1) {
                v28 := GenS2{F0: (2 < 18446744073709551615), F1: (fi27 + genFn0("hello", 0, (v3 ^ (v3 & 1))))}
                fmt.Println(v3, v4, v8)
                _ = v28
            }
        }
        if (950 > 42) {
            if true {
                v29 := GenS1{F0: fi27, F1: (^(v8))}
                for fi30 := 0; fi30 < 10; fi30++ {
                    v31 := GenS3{F0: 255, F1: ("" + "a")}
                    _ = fi30
                    _ = v31
                }
                _ = v29
            } else {
                v8 += v8
            }
        }
        v1 += 3.14159
        _ = fi27
    }
    switch "" {
    case "\n\t":
        v8 = (^(v8))
    case "a":
        v32 := []int{v15, v15, (v15 / (((v15 / ((v15) | 1))) | 1)), v15}
        _ = v32
    default:
        fmt.Println(v3)
    }
    for fi33 := 0; fi33 < 16; fi33++ {
        for fi34 := 0; fi34 < 18; fi34++ {
            v3 = 2147483647
            v4 += (v4 / ((-128) | 1))
            var v35 uint = 0
            switch (fi34 / ((fi34) | 1)) {
            case -2147483648:
                if false {
                    v3 = v3
                }
            case 667:
                if false {
                    v8 = v8
                }
            case 9223372036854775807:
                v36 := []int{(-(fi33)), int(v4), (fi34 << (((v35 + (^(v35)))) & 63)), (fi33 ^ -9223372036854775808)}
                _ = v36
            }
            _ = fi34
            _ = v35
        }
        v2["a"] = (fi33 % (((v15 * genFn0((("hello" + (("\n\t" + "a") + "\n\t")) + ""), 9223372036854775807, (v3 & v3)))) | 1))
        _ = fi33
    }
    v15 = v15
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
