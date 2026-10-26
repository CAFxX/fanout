package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 int; F2 int64 }
type GenS3 struct { F0 int; F1 int64 }
func genFn0(p0 uint64, p1 string, p2 int32) int8 {
    for fi6 := 0; fi6 < 16; fi6++ {
        var v7 int8 = -25
        v7 = 127
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 18; fi8++ {
        for fi9 := 0; fi9 < 11; fi9++ {
            v10 := GenS1{F0: p2}
            fmt.Println(fi9)
            if false {
                if (false || (-1 > -1)) {
                    var v11 float32 = 0.1
                    var v12 uint16 = 868
                    if (!(!false)) {
                        v13 := []int{(fi8 - (fi8 | fi9)), fi9, fi9, fi8}
                        _ = v13
                    } else {
                        v14 := []int{}
                        _ = v14
                    }
                    _ = v11
                    _ = v12
                }
            }
            _ = fi9
            _ = v10
        }
        v15 := []int{-27, (^(fi8))}
        _ = fi8
        _ = v15
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        v17 := GenS1{F0: 42}
        var v18 int8 = 100
        _ = fi16
        _ = v17
        _ = v18
    }
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    return -1
}
func genFn1() uint16 {
    var v19 int16 = -638
    var v20 int8 = 1
    if true {
        for fi21 := 0; fi21 < 3; fi21++ {
            var v22 string = ""
            for fi23 := 0; fi23 < 9; fi23++ {
                if (0 >= int32(v19)) {
                    v24 := []int{}
                    v22 = (("hello" + v22) + (v22 + v22))
                    _ = v24
                }
                v19 = 721
                _ = fi23
            }
            _ = fi21
            _ = v22
        }
        v25 := []int{-9223372036854775808, int(v19), 1}
        for fi26 := 0; fi26 < 1; fi26++ {
            v19 = v19
            var v27 uint8 = 103
            v20 += v20
            _ = fi26
            _ = v27
        }
        _ = v25
    } else {
        v28 := GenS1{F0: -622}
        v29 := GenS2{F0: (-(v19)), F1: int(v20), F2: 1}
        v20 -= (v20 - (-(v20)))
        _ = v28
        _ = v29
    }
    for fi30 := 0; fi30 < 14; fi30++ {
        fmt.Println(fi30)
        _ = fi30
    }
    _ = v19
    _ = v20
    return 0
}

func main() {
    var v1 int32 = -2147483648
    var v2 int16 = -32768
    v3 := []string{}
    var v4 uint32 = 65535
    var v5 float64 = (0.1 / 1e10)
    switch "a" {
    case "\n\t":
        v2 -= (v2 - (v2 / ((v2) | 1)))
    case "a":
        v5 += v5
    default:
        fmt.Println(v1, v1, v5)
    }
    if len(v3) > 0 {
        v3[0] = ""
    }
    var v31 float64 = (v5 * float64((1e10 / 3.14159)))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch (("foo bar" + (("hello" + ("hello" + ("" + "foo bar"))) + "")) + "\n\t") {
    case "foo bar":
        v4 += 4294967295
    case "a":
        _ = len(v3)
    case "":
        v32 := GenS2{F0: (^(v2)), F1: -431, F2: -1}
        _ = v32
    default:
        var v33 uint32 = (^(v4))
        _ = v33
    }
    if (genFn1() != 1) {
        v34 := GenS3{F0: 4294967295, F1: 42}
        _ = v34
    }
    v3 = append(v3, ("" + ""))
    fmt.Println(v2, v5, v1)
    var v35 bool = true
    switch ("hello" + "foo bar") {
    case "a":
        v36 := GenS3{F0: 4294967295, F1: -162}
        _ = v36
    case "hello":
        v37 := GenS2{F0: v2, F1: 42, F2: -1}
        _ = v37
    case "\n\t":
        switch 32768 {
        case -9223372036854775808:
            var v38 int8 = 1
            _ = v38
        }
    }
    var v39 bool = true
    if (v39 && v39) {
        switch -421 {
        case 32767:
            v2 -= (-(v2))
        case -772:
            v3 = append(v3, "\n\t")
        default:
            v5 = (v5 / (v31 * float64(3.14159)))
        }
    }
    for fi40 := 0; fi40 < 14; fi40++ {
        if (127 <= 255) {
            if v35 {
                v39 = false
            }
        }
        v35 = v39
        if v39 {
            v4 = v4
            v31 = float64(float32(((0.0 / 1.0) * float64((-2.25 - -2.25)))))
        } else {
            v2 -= (v2 % ((-32768) | 1))
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        _ = fi40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v35
    _ = v39
    _ = v4
    _ = v5
}
