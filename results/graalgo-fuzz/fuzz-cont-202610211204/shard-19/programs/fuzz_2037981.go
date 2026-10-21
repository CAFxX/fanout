package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 uint32 }
type GenS2 struct { F0 uint64; F1 int8 }
func genFn0(p0 int8) uint8 {
    var v6 uint16 = 255
    v7 := []int{32767, 1, 128, -279}
    for fi8 := 0; fi8 < 5; fi8++ {
        if (((!false) && true) || true) {
            switch fi8 {
            case 9223372036854775807:
                if (!(!true)) {
                    v9 := GenS1{F0: 65535, F1: false, F2: 42}
                    if (127 >= p0) {
                        v10 := GenS1{F0: 2, F1: (65535 != v6), F2: 0}
                        v7 = append(v7, fi8)
                        _ = v10
                    }
                    _ = v9
                }
            }
            var v11 int64 = 42
            var v12 uint16 = v6
            _ = v11
            _ = v12
        } else {
            _ = len(v7)
            fmt.Println(v6, fi8)
            var v13 int8 = p0
            _ = v13
        }
        p0 = p0
        p0 = 42
        _ = fi8
    }
    v14 := GenS2{F0: 18446744073709551615, F1: -128}
    _ = p0
    _ = v14
    _ = v6
    _ = v7
    return 255
}
func genFn1(p0 uint64) uint8 {
    v15 := []int{-971}
    v16 := GenS2{F0: (p0 % (((p0 | p0)) | 1)), F1: -24}
    var v17 int8 = 42
    _ = p0
    _ = v15
    _ = v16
    _ = v17
    return 42
}

func main() {
    var v1 uint = 89
    var v2 uint16 = 5
    var v3 float32 = 3.14159
    v4 := make(map[uint32]uint32)
    var v5 uint16 = (v2 * v2)
    switch "hello" {
    case "hello":
        switch ("\n\t" + ("\n\t" + ("foo bar" + "foo bar"))) {
        case "hello":
            v18 := GenS1{F0: v1, F1: ((65535 <= 42) && (2 >= v1)), F2: 268}
            _ = v18
        }
    default:
        v19 := GenS1{F0: 0, F1: false, F2: 65535}
        _ = v19
    }
    for fi20 := 0; fi20 < 15; fi20++ {
        var v21 int = fi20
        _ = fi20
        _ = v21
    }
    v22 := []int{127}
    var v23 string = (("hello" + "hello") + "foo bar")
    if (0 <= 1) {
        fmt.Println(v2)
    }
    var v24 int32 = 2147483647
    if true {
        var v25 int16 = 600
        _ = v25
    } else {
        for fi26 := 0; fi26 < 17; fi26++ {
            for fi27 := 0; fi27 < 19; fi27++ {
                fmt.Println(v1, v24, fi26)
                var v28 uint16 = v5
                v3 = -2.25
                _ = fi27
                _ = v28
            }
            if (v2 != (^(v2))) {
                v4[462] = 2
                v1 -= (v1 ^ 255)
            }
            if false {
                var v29 bool = (!(int16(v1) <= 840))
                _ = v29
            }
            _ = fi26
        }
        _ = len(v4)
    }
    var v30 int16 = 32767
    v31 := GenS1{F0: (v1 | 1), F1: false, F2: 42}
    v4[0] = 1
    var v32 uint = 501
    v33 := v4[522]
    if true {
        if (!(true && ((v23 == ((v23 + v23) + "")) && false))) {
            var v34 int32 = v24
            v35 := v4[(v33 << ((255) & 63))]
            var v36 uint8 = 255
            _ = v34
            _ = v35
            _ = v36
        } else {
            var v37 uint = v32
            _ = v37
        }
        for fi38 := 0; fi38 < 12; fi38++ {
            var v39 uint32 = 980
            _ = fi38
            _ = v39
        }
        if (!(false && false)) {
            var v40 bool = ((^(v33)) > v33)
            var v41 bool = true
            _ = v40
            _ = v41
        }
    }
    if false {
        v33 -= v33
    }
    fmt.Println(v3, v24)
    v1 += 0
    _ = len(v22)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
