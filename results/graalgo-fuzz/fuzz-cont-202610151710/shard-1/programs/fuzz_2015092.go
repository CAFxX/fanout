package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint8 }
type GenS2 struct { F0 int8; F1 string; F2 uint }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 int32; F1 uint8 }
type GenS5 struct { F0 uint16 }
func genFn0() int8 {
    var v5 int32 = 405
    if (false && (-916 != 0)) {
        for fi6 := 0; fi6 < 16; fi6++ {
            for fi7 := 0; fi7 < 17; fi7++ {
                if (!true) {
                    v8 := GenS1{F0: "a", F1: 2, F2: 128}
                    _ = v8
                } else {
                    v9 := []int{fi7, -2147483648}
                    var v10 uint32 = 515
                    switch ("foo bar" + "hello") {
                    case "a":
                        v11 := GenS2{F0: -1, F1: "\n\t", F2: 255}
                        _ = v11
                    case "\n\t":
                        var v12 uint16 = 32767
                        _ = v12
                    default:
                        v10 -= v10
                    }
                    _ = v10
                    _ = v9
                }
                fmt.Println(fi6, fi7)
                v13 := []int{fi6, fi6, (fi6 * (-(fi6)))}
                _ = fi7
                _ = v13
            }
            _ = fi6
        }
        v14 := GenS2{F0: 1, F1: "hello", F2: 0}
        _ = v14
    } else {
        v15 := []int{-2147483648, -672, -249}
        v16 := GenS2{F0: 0, F1: (("hello" + "\n\t") + ""), F2: 255}
        v15 = append(v15, 42)
        _ = v15
        _ = v16
    }
    v17 := GenS2{F0: 42, F1: ((("foo bar" + "hello") + ("foo bar" + ("foo bar" + "a"))) + ""), F2: 450}
    v5 = 2147483647
    _ = v17
    _ = v5
    return 100
}
func genFn1(p0 string, p1 string, p2 int32) string {
    var v18 int16 = -390
    v19 := GenS3{F0: 42, F1: false}
    p0 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v19
    return p1
}

func main() {
    var v1 uint = 65535
    var v2 int8 = -1
    var v3 int64 = 9223372036854775807
    var v4 float64 = -2.25
    fmt.Println(v3, v4, v2)
    if ((v2 - v2) >= (v2 << (((v1 + v1)) & 63))) {
        var v20 int = 42
        v4 -= (3.14159 - v4)
        if (!(2147483647 >= 0)) {
            v20 = v20
            fmt.Println(v20, v20)
        }
        _ = v20
    } else {
        var v21 uint8 = 255
        var v22 uint64 = 760
        _ = v21
        _ = v22
    }
    v1 = v1
    v3 = 1
    var v23 bool = true
    var v24 int = 42
    if false {
        var v25 int8 = genFn0()
        var v26 int8 = (v2 & 107)
        for fi27 := 0; fi27 < 1; fi27++ {
            v28 := GenS3{F0: (^(v3)), F1: v23}
            _ = fi27
            _ = v28
        }
        _ = v25
        _ = v26
    }
    if (0 <= 975) {
        switch ((("a" + "\n\t") + "a") + ("a" + ((("hello" + "\n\t") + (("a" + "\n\t") + "hello")) + "a"))) {
        case "\n\t":
            v29 := GenS4{F0: 42, F1: 128}
            _ = v29
        default:
            v2 = -1
        }
        if true {
            for fi30 := 0; fi30 < 5; fi30++ {
                for fi31 := 0; fi31 < 20; fi31++ {
                    fmt.Println(v1, v23, v4)
                    v1 -= (^(v1))
                    v1 -= 18446744073709551615
                    _ = fi31
                }
                v1 = v1
                v32 := []int{4294967295, v24}
                _ = fi30
                _ = v32
            }
        }
        v33 := []int{(v24 - 128), 2147483647}
        _ = v33
    } else {
        var v34 uint32 = 4294967295
        for fi35 := 0; fi35 < 8; fi35++ {
            v23 = true
            var v36 string = ""
            var v37 int64 = int64(v34)
            _ = fi35
            _ = v36
            _ = v37
        }
        _ = v34
    }
    switch "a" {
    case "hello":
        v24 -= int(v2)
    default:
        v38 := []int{v24, 42, v24, v24}
        _ = v38
    }
    v39 := GenS5{F0: 65535}
    if (-498 == -32768) {
        var v40 int32 = -1
        fmt.Println(v1, v3)
        v41 := []int{32768, (v24 & v24)}
        _ = v40
        _ = v41
    } else {
        if true {
            v42 := []int{}
            _ = v42
        } else {
            v43 := GenS2{F0: v2, F1: "hello", F2: v1}
            var v44 int = (v24 & v24)
            _ = v43
            _ = v44
        }
    }
    var v45 float32 = -2.25
    v46 := []int{int(v3)}
    var v47 uint16 = 260
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v46:", len(v46), v46)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v39
    _ = v4
    _ = v45
    _ = v46
    _ = v47
}
