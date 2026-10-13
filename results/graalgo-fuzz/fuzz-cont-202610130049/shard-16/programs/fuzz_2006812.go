package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 string }
type GenS2 struct { F0 int32; F1 bool; F2 int32 }

func main() {
    var v1 int8 = -128
    var v2 int32 = -2147483648
    var v3 string = ((("foo bar" + ("foo bar" + ("a" + ("" + "foo bar")))) + ("" + "hello")) + "a")
    var v4 int64 = 42
    v5 := make(map[uint32]bool)
    switch -232 {
    case 483:
        for fi6 := 0; fi6 < 6; fi6++ {
            v7 := []int{-2147483648, (fi6 % ((fi6) | 1)), fi6, 897}
            v2 -= v2
            v8 := GenS1{F0: 4294967295, F1: (^(v4)), F2: ""}
            _ = fi6
            _ = v7
            _ = v8
        }
    case 1:
        if (false || (0 <= 0)) {
            v2 += 847
        }
    case 4294967295:
        v2 = v2
    }
    v9 := []int{696}
    fmt.Println(v2, v1, v1)
    for fi10 := 0; fi10 < 3; fi10++ {
        v2 -= v2
        _ = fi10
    }
    for fi11 := 0; fi11 < 6; fi11++ {
        var v12 int32 = -2147483648
        for fi13 := 0; fi13 < 4; fi13++ {
            if true {
                v14 := GenS2{F0: 1, F1: (386 != 32767), F2: v12}
                var v15 uint8 = 93
                _ = v14
                _ = v15
            } else {
                delete(v5, 2)
                var v16 float64 = -2.25
                _ = v16
            }
            _ = fi13
        }
        var v17 int8 = -1
        _ = fi11
        _ = v12
        _ = v17
    }
    v4 = (v4 | v4)
    v18 := GenS2{F0: (v2 % (((v2 >> ((145) & 63))) | 1)), F1: (255 <= 828), F2: v2}
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    if (0 <= 65535) {
        var v19 uint64 = 18446744073709551615
        for fi20 := 0; fi20 < 1; fi20++ {
            for fi21 := 0; fi21 < 6; fi21++ {
                var v22 uint32 = 4294967295
                if true {
                    var v23 uint32 = 65535
                    v24 := GenS2{F0: v2, F1: false, F2: v2}
                    if true {
                        fmt.Println(v3, v1, v23)
                    } else {
                        var v25 uint32 = (v22 - (v22 & v22))
                        delete(v5, (v23 & v23))
                        _ = v25
                    }
                    _ = v23
                    _ = v24
                }
                fmt.Println(fi21, v2)
                _ = fi21
                _ = v22
            }
            _ = fi20
        }
        _ = v19
    } else {
        switch v3 {
        case "hello":
            fmt.Println(v3)
        }
    }
    switch v3 {
    case "":
        v2 -= -423
    case "\n\t":
        var v26 float32 = ((1.5 - 0.0) + -2.25)
        _ = v26
    case "a":
        v3 = v3
    }
    fmt.Println(v3, v3, v4)
    var v27 int16 = 1000
    var v28 uint32 = 1
    var v29 float64 = 1.5
    _ = len(v9)
    var v30 uint64 = 841
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v18
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v9
}
