package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 uint; F1 uint8 }
type GenS4 struct { F0 string; F1 uint16; F2 int32 }
func genFn0() uint8 {
    if false {
        for fi6 := 0; fi6 < 4; fi6++ {
            if true {
                v7 := []int{(fi6 & -332), 1}
                switch (^(fi6)) {
                case -2147483648:
                    v8 := GenS1{F0: 18446744073709551615}
                    _ = v8
                case -9223372036854775808:
                    var v9 uint16 = 255
                    _ = v9
                case 128:
                    var v10 int64 = 42
                    _ = v10
                }
                var v11 int8 = 0
                _ = v11
                _ = v7
            }
            var v12 int64 = -9223372036854775808
            _ = fi6
            _ = v12
        }
        var v13 int64 = -9223372036854775808
        _ = v13
    }
    switch "" {
    case "":
        for fi14 := 0; fi14 < 14; fi14++ {
            var v15 uint8 = 181
            v15 += v15
            if false {
                fmt.Println(v15, v15)
            } else {
                var v16 float64 = 0.0
                for fi17 := 0; fi17 < 4; fi17++ {
                    if false {
                        var v18 string = ((("hello" + ("a" + "a")) + "\n\t") + "\n\t")
                        _ = v18
                    } else {
                        fmt.Println(fi14)
                        var v19 string = ""
                        v16 = (v16 - (-2.25 - float64(float32(0.0))))
                        _ = v19
                    }
                    _ = fi17
                }
                _ = v16
            }
            _ = fi14
            _ = v15
        }
    case "a":
        if true {
            var v20 int64 = 0
            var v21 int8 = -99
            switch -1 {
            case 32767:
                var v22 int = -9223372036854775808
                _ = v22
            case -9223372036854775808:
                fmt.Println(v21, v20)
            case 721:
                v23 := GenS2{F0: 1, F1: v21}
                _ = v23
            default:
                fmt.Println(v20, v21)
            }
            _ = v20
            _ = v21
        } else {
            v24 := []int{4294967295}
            _ = v24
        }
    default:
        v25 := GenS1{F0: 4294967295}
        _ = v25
    }
    var v26 float32 = ((0.0 / 3.14159) * 1.5)
    for fi27 := 0; fi27 < 4; fi27++ {
        v28 := GenS3{F0: 0, F1: 60}
        _ = fi27
        _ = v28
    }
    var v29 int64 = -509
    _ = v26
    _ = v29
    return 42
}

func main() {
    var v1 int8 = 72
    var v2 uint32 = 2
    var v3 uint8 = 214
    var v4 int = 127
    v5 := []uint32{v2}
    fmt.Println(v2)
    if true {
        if len(v5) > 0 {
            v5[0] = 1
        }
        v30 := GenS2{F0: 65535, F1: (v1 % (((v1 * v1)) | 1))}
        _ = v30
    } else {
        var v31 bool = true
        _ = v31
    }
    _ = len(v5)
    if false {
        var v32 float64 = 1.5
        _ = v32
    }
    if (v3 <= 128) {
        var v33 uint64 = 595
        var v34 int = v4
        v35 := GenS4{F0: "foo bar", F1: 1, F2: 382}
        _ = v33
        _ = v34
        _ = v35
    }
    v36 := GenS1{F0: 255}
    fmt.Println(v4, v1, v2)
    var v37 uint = 42
    v2 -= (^(v2))
    v38 := GenS4{F0: "", F1: 255, F2: int32(v1)}
    var v39 float32 = (3.14159 + 1e10)
    var v40 uint = v37
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
