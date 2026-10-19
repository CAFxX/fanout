package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 string; F2 int32 }
type GenS3 struct { F0 int; F1 int }
type GenS4 struct { F0 int16; F1 int }
type GenS5 struct { F0 string }
type GenS6 struct { F0 bool }
func genFn0(p0 uint64, p1 uint16) uint {
    var v6 uint64 = p0
    switch "\n\t" {
    case "hello":
        for fi7 := 0; fi7 < 19; fi7++ {
            var v8 string = (("hello" + "foo bar") + ("\n\t" + "foo bar"))
            for fi9 := 0; fi9 < 13; fi9++ {
                p1 = (p1 / ((32767) | 1))
                _ = fi9
            }
            var v10 int = fi7
            _ = fi7
            _ = v10
            _ = v8
        }
    }
    v11 := []int{}
    _ = p0
    _ = p1
    _ = v11
    _ = v6
    return 882
}

func main() {
    var v1 uint64 = 2
    v2 := []string{"", "\n\t", "\n\t"}
    v3 := []uint32{0, 126, 4294967295}
    var v4 uint64 = 1
    var v5 float64 = 0.1
    fmt.Println(v1, v4, v4)
    if true {
        var v12 uint8 = 1
        _ = v12
    }
    for fi13 := 0; fi13 < 11; fi13++ {
        v14 := GenS1{F0: (fi13 / (((fi13 >> ((617) & 63))) | 1))}
        _ = fi13
        _ = v14
    }
    v4 -= 617
    fmt.Println(v5)
    var v15 uint32 = 231
    if ((301 >= 334) && ("foo bar" >= "a")) {
        var v16 bool = false
        var v17 uint32 = v15
        _ = v16
        _ = v17
    } else {
        var v18 int16 = 1000
        if (false || (v18 == v18)) {
            if ((!true) || ((0.0 + float32((v5 + ((-2.25 - ((v5 * 0.0) * 0.0)) + 3.14159)))) > (-2.25 + (1.5 - 0.1)))) {
                for fi19 := 0; fi19 < 19; fi19++ {
                    var v20 uint16 = uint16(v1)
                    fmt.Println(v5, v15, v18)
                    switch 970 {
                    case 32767:
                        var v21 float64 = 3.14159
                        _ = v21
                    }
                    _ = fi19
                    _ = v20
                }
                var v22 bool = (!true)
                var v23 int32 = 1
                _ = v22
                _ = v23
            } else {
                fmt.Println(v18)
                v24 := GenS2{F0: -99, F1: "hello", F2: -213}
                _ = v24
            }
            v25 := GenS2{F0: -128, F1: "a", F2: -406}
            _ = v25
        }
        _ = v18
    }
    for fi26 := 0; fi26 < 13; fi26++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi26
    }
    for fi27 := 0; fi27 < 18; fi27++ {
        fmt.Println(v1, v4, fi27)
        var v28 uint64 = (v1 | (v4 / (((v4 & v1)) | 1)))
        var v29 uint32 = v15
        _ = fi27
        _ = v28
        _ = v29
    }
    for fi30 := 0; fi30 < 11; fi30++ {
        v4 = 18446744073709551615
        v1 += (v4 % (((v4 / ((0) | 1))) | 1))
        v31 := GenS3{F0: fi30, F1: fi30}
        _ = fi30
        _ = v31
    }
    v32 := GenS2{F0: -1, F1: (((("a" + "foo bar") + "") + "hello") + "foo bar"), F2: 2147483647}
    if false {
        v3 = append(v3, (v15 | 1))
        var v33 string = "hello"
        _ = v33
    } else {
        fmt.Println(v15, v4, v4)
        var v34 uint16 = 65535
        _ = v34
    }
    var v35 int = 413
    v36 := GenS4{F0: 42, F1: (v35 ^ v35)}
    if false {
        v37 := GenS5{F0: "a"}
        for fi38 := 0; fi38 < 9; fi38++ {
            v39 := GenS6{F0: ((v35 / (((v35 << ((255) & 63))) | 1)) > fi38)}
            _ = fi38
            _ = v39
        }
        _ = v37
    } else {
        for fi40 := 0; fi40 < 1; fi40++ {
            fmt.Println(v5)
            _ = fi40
        }
        v4 = (v1 + 42)
        v41 := GenS5{F0: (("" + ("hello" + "foo bar")) + (("hello" + ("hello" + ("\n\t" + ("" + "\n\t")))) + "hello"))}
        _ = v41
    }
    var v42 uint8 = 0
    var v43 float32 = float32(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v32
    _ = v35
    _ = v36
    _ = v4
    _ = v42
    _ = v43
    _ = v5
}
