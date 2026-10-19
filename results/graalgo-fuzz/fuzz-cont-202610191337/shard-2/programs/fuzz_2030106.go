package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 bool }
type GenS2 struct { F0 uint8; F1 int32; F2 int8 }
type GenS3 struct { F0 int64; F1 string }
func genFn0() string {
    if false {
        var v5 string = "hello"
        _ = v5
    } else {
        if true {
            for fi6 := 0; fi6 < 6; fi6++ {
                var v7 uint8 = 37
                _ = fi6
                _ = v7
            }
        } else {
            var v8 int = 32768
            _ = v8
        }
        switch 32767 {
        case 128:
            v9 := GenS1{F0: 259, F1: 1, F2: (!(!true))}
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        var v11 float64 = ((0.1 * -2.25) - (((-2.25 * float64(-2.25)) + (1.5 + (0.0 + 3.14159))) * 0.1))
        v11 = v11
        _ = fi10
        _ = v11
    }
    v12 := []int{}
    if (1 <= 1) {
        if true {
            var v13 int = 986
            for fi14 := 0; fi14 < 20; fi14++ {
                fmt.Println(v13, v13)
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
                _ = fi14
            }
            v13 = v13
            _ = v13
        } else {
            var v15 float64 = 1e10
            _ = len(v12)
            _ = v15
        }
        var v16 bool = (0 < 18446744073709551615)
        _ = v16
    }
    _ = v12
    return ""
}
func genFn1(p0 int, p1 uint16, p2 bool) int16 {
    if ((-9223372036854775808 > -9223372036854775808) || (0 < 42)) {
        v17 := GenS2{F0: 187, F1: 680, F2: 127}
        _ = v17
    }
    var v18 uint8 = 1
    if false {
        switch (-(p0)) {
        case 903:
            v19 := GenS1{F0: 2, F1: 871, F2: p2}
            _ = v19
        }
    } else {
        var v20 uint8 = v18
        _ = v20
    }
    v18 -= (v18 >> ((1) & 63))
    v21 := []int{(p0 / (((p0 | 1)) | 1)), p0, p0}
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v21
    return 910
}

func main() {
    var v1 float64 = -2.25
    var v2 float64 = v1
    var v3 uint8 = 0
    var v4 uint = 255
    v22 := []int{32767, 4294967295, -9223372036854775808, 32767}
    var v23 uint8 = (v3 & v3)
    var v24 int = 128
    v25 := GenS3{F0: 64, F1: "a"}
    fmt.Println(v23, v1)
    for fi26 := 0; fi26 < 10; fi26++ {
        fmt.Println(v24, fi26)
        v23 = (v3 | v23)
        _ = fi26
    }
    var v27 bool = (((2147483647 >= 2147483647) || (!false)) && (-103 <= 127))
    fmt.Println(v1, v2, v3)
    var v28 uint32 = 874
    switch v24 {
    case 2147483647:
        var v29 uint = 873
        _ = v29
    case -1:
        if len(v22) > 0 {
            v22[0] = int(v28)
        }
    }
    var v30 bool = v27
    var v31 float32 = 3.14159
    var v32 int8 = 16
    v1 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
}
