package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 uint }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool; F1 int64; F2 bool }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint32; F1 uint32 }
type GenS6 struct { F0 uint16; F1 bool }
type GenS7 struct { F0 int32; F1 uint32 }
type GenS8 struct { F0 uint; F1 uint8 }
func genFn0() uint {
    var v5 uint32 = 2147483647
    for fi6 := 0; fi6 < 3; fi6++ {
        switch fi6 {
        case 4294967295:
            for fi7 := 0; fi7 < 13; fi7++ {
                var v8 int64 = 0
                switch 128 {
                case 2147483647:
                    if true {
                        v9 := []int{537, 127}
                        v10 := GenS1{F0: -616, F1: 42, F2: 17}
                        v5 = 823
                        v11 := GenS2{F0: 0}
                        v12 := GenS3{F0: false, F1: v8, F2: true}
                        _ = v10
                        _ = v11
                        _ = v12
                        _ = v9
                    }
                case 411:
                    switch ((("\n\t" + "a") + (((("" + ((((("" + (("foo bar" + "") + "")) + "a") + (("hello" + "\n\t") + "hello")) + ("\n\t" + "\n\t")) + ("hello" + ("" + ("\n\t" + ""))))) + "foo bar") + (("a" + ("\n\t" + "hello")) + "")) + "\n\t")) + "hello") {
                    case "foo bar":
                        if (42 <= 219) {
                            v8 = -9223372036854775808
                        }
                        fmt.Println(v8)
                    case "":
                        fmt.Println(fi6, fi7, fi6)
                    case "hello":
                        var v13 uint32 = 42
                        _ = v13
                    default:
                        v5 += v5
                    }
                }
                v8 = v8
                v14 := GenS4{F0: 18446744073709551615}
                v8 -= -457
                v15 := GenS3{F0: (-617 >= -850), F1: 228, F2: true}
                if false {
                    fmt.Println(v5, v5, fi7)
                    v16 := []int{fi6, fi7}
                    _ = v16
                } else {
                    v17 := []int{}
                    fmt.Println(fi6, v5)
                    v18 := GenS3{F0: (true && false), F1: (v8 / ((v8) | 1)), F2: false}
                    v19 := GenS5{F0: 2, F1: v5}
                    v8 = v8
                    v20 := GenS6{F0: 154, F1: (true || false)}
                    var v21 int64 = v8
                    _ = v17
                    _ = v18
                    _ = v19
                    _ = v20
                    _ = v21
                }
                v22 := []int{}
                _ = fi7
                _ = v14
                _ = v15
                _ = v22
                _ = v8
            }
        }
        _ = fi6
    }
    if (-32768 == 494) {
        var v23 uint64 = 4294967295
        _ = v23
    }
    _ = v5
    return 2
}
func genFn1(p0 uint16) uint {
    if (!(!(888 != 4294967295))) {
        var v24 int32 = -652
        _ = v24
    }
    if true {
        var v25 string = "foo bar"
        if true {
            p0 -= p0
            v26 := []int{1, 0}
            var v27 string = v25
            _ = v26
            _ = v27
        }
        for fi28 := 0; fi28 < 13; fi28++ {
            v25 = v25
            _ = fi28
        }
        _ = v25
    } else {
        v29 := GenS7{F0: -660, F1: 578}
        _ = v29
    }
    _ = p0
    return 42
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 255
    var v3 uint16 = v2
    var v4 int32 = -848
    if false {
        v30 := GenS5{F0: 4294967295, F1: 2147483648}
        fmt.Println(v3, v2)
        _ = v30
    }
    v4 -= (v4 - (^(v4)))
    var v31 uint16 = v3
    v31 = (v3 % ((v31) | 1))
    v32 := GenS1{F0: -876, F1: 1, F2: genFn1(569)}
    var v33 uint32 = 0
    fmt.Println(v31)
    var v34 bool = false
    var v35 int64 = -706
    var v36 bool = (v34 || v34)
    switch 32768 {
    case -399:
        if ((!(!(!false))) && v34) {
            v2 -= (v3 + (v2 % (((v2 | 2)) | 1)))
        } else {
            if true {
                v37 := GenS2{F0: 450}
                _ = v37
            }
        }
    case 4294967295:
        if v36 {
            for fi38 := 0; fi38 < 14; fi38++ {
                v39 := GenS8{F0: 0, F1: 2}
                _ = fi38
                _ = v39
            }
            var v40 float64 = 1e10
            _ = v40
        }
    }
    v41 := GenS2{F0: 18446744073709551615}
    var v42 uint = 65535
    var v43 uint64 = 127
    for fi44 := 0; fi44 < 20; fi44++ {
        for fi45 := 0; fi45 < 16; fi45++ {
            var v46 int8 = 100
            _ = fi45
            _ = v46
        }
        _ = fi44
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v41
    _ = v42
    _ = v43
}
