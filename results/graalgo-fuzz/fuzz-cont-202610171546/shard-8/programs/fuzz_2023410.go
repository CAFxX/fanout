package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 string }
type GenS2 struct { F0 uint16; F1 int8; F2 int }
type GenS3 struct { F0 uint32; F1 uint16 }
func genFn0(p0 uint32) int16 {
    var v6 int64 = 0
    for fi7 := 0; fi7 < 16; fi7++ {
        v6 = 9223372036854775807
        _ = fi7
    }
    _ = p0
    _ = v6
    return 1
}
func genFn1(p0 uint8) uint {
    v8 := GenS1{F0: 0, F1: ((577 <= 1) && (255 > 1)), F2: ""}
    for fi9 := 0; fi9 < 6; fi9++ {
        v10 := []int{fi9, fi9, fi9, fi9}
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = v8
    return 4294967295
}

func main() {
    v1 := []int{128, 630}
    var v2 int8 = 0
    var v3 string = ("\n\t" + ("" + "hello"))
    v4 := make(map[uint32]bool)
    var v5 uint = 2
    var v11 bool = ((1 == 393) && (0 != 4294967295))
    for fi12 := 0; fi12 < 5; fi12++ {
        v13 := v4[uint32(v5)]
        _ = fi12
        _ = v13
    }
    v14 := GenS1{F0: (^(v5)), F1: ((831 >= 595) && (65535 <= 42)), F2: v3}
    fmt.Println(v5)
    var v15 string = "foo bar"
    for fi16 := 0; fi16 < 10; fi16++ {
        var v17 int = (fi16 + fi16)
        _ = fi16
        _ = v17
    }
    var v18 int8 = v2
    if v11 {
        var v19 int64 = 1
        _ = len(v4)
        _ = v19
    }
    switch (((v15 + ((v15 + (v15 + v3)) + v15)) + v3) + v15) {
    case "\n\t":
        for fi20 := 0; fi20 < 14; fi20++ {
            v21 := GenS1{F0: (^(v5)), F1: v11, F2: v15}
            if len(v1) > 0 {
                v1[0] = fi20
            }
            fmt.Println(fi20, fi20, v3)
            _ = fi20
            _ = v21
        }
    default:
        if ((v3 + v3) <= "") {
            fmt.Println(v5)
            v22 := GenS1{F0: 0, F1: (v11 && v11), F2: v3}
            _ = v22
        } else {
            for fi23 := 0; fi23 < 13; fi23++ {
                if (2147483647 <= 50) {
                    delete(v4, 1)
                    v24 := GenS1{F0: v5, F1: false, F2: "a"}
                    _ = v24
                } else {
                    var v25 int = fi23
                    v26 := GenS2{F0: 32, F1: v18, F2: fi23}
                    for fi27 := 0; fi27 < 14; fi27++ {
                        var v28 int32 = 0
                        _ = fi27
                        _ = v28
                    }
                    _ = v25
                    _ = v26
                }
                var v29 int = -9223372036854775808
                _ = fi23
                _ = v29
            }
            for fi30 := 0; fi30 < 13; fi30++ {
                if v11 {
                    if (true && v11) {
                        var v31 string = v15
                        v11 = v11
                        _ = v31
                    } else {
                        fmt.Println(v18, v11, fi30)
                    }
                    v32 := GenS3{F0: 71, F1: 42}
                    _ = v32
                }
                v33 := GenS1{F0: v5, F1: v11, F2: v3}
                fmt.Println(fi30)
                _ = fi30
                _ = v33
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
