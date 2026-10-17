package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int16; F1 int64; F2 int64 }
func genFn0() uint64 {
    if true {
        for fi5 := 0; fi5 < 6; fi5++ {
            switch (fi5 & 128) {
            case -733:
                for fi6 := 0; fi6 < 18; fi6++ {
                    v7 := GenS1{F0: true, F1: 1, F2: 1000}
                    v8 := GenS1{F0: (!(true && true)), F1: 255, F2: -32768}
                    _ = fi6
                    _ = v7
                    _ = v8
                }
            case 68:
                for fi9 := 0; fi9 < 6; fi9++ {
                    var v10 string = "\n\t"
                    _ = fi9
                    _ = v10
                }
            }
            v11 := []int{fi5}
            _ = fi5
            _ = v11
        }
        var v12 uint16 = 2
        fmt.Println(v12)
        _ = v12
    }
    v13 := []int{}
    _ = v13
    return 4294967295
}
func genFn1() int8 {
    if ("a" == "a") {
        switch 128 {
        case -354:
            var v14 uint8 = 255
            _ = v14
        case 103:
            var v15 int8 = 127
            _ = v15
        case 380:
            v16 := []int{-2147483648, 32768, 826}
            _ = v16
        }
    } else {
        var v17 uint = 18446744073709551615
        v18 := GenS1{F0: (!true), F1: 255, F2: -956}
        v19 := GenS2{F0: (127 == -1)}
        _ = v17
        _ = v18
        _ = v19
    }
    var v20 int64 = 42
    v20 = v20
    _ = v20
    return -1
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int8 = 114
    var v3 int = 32768
    var v4 float32 = 0.0
    v21 := GenS1{F0: (-2.25 != ((0.1 - -2.25) - (0.1 - float64(v4)))), F1: 447, F2: 0}
    if (true && true) {
        fmt.Println(v4, v2, v2)
        _ = len(v1)
        v22 := []int{v3}
        _ = v22
    }
    var v23 int64 = 223
    fmt.Println(v4)
    v2 -= (v2 / (((v2 >> ((255) & 63))) | 1))
    v4 = v4
    var v24 uint8 = 125
    v25 := v1[0]
    if ("" <= "\n\t") {
        if (4294967295 <= (^(v3))) {
            switch -274 {
            case 841:
                fmt.Println(v3)
            }
            delete(v1, 0)
        } else {
            v23 = (v23 << ((65535) & 63))
            var v26 uint32 = 177
            _ = v26
        }
        var v27 int = 32767
        _ = v27
    }
    v28 := []int{(^(v25)), (v3 ^ v3), (v25 / ((-2147483648) | 1))}
    for fi29 := 0; fi29 < 1; fi29++ {
        for fi30 := 0; fi30 < 1; fi30++ {
            fmt.Println(fi29)
            fmt.Println(v25, fi30)
            v31 := GenS3{F0: (v24 | v24)}
            _ = fi30
            _ = v31
        }
        for fi32 := 0; fi32 < 14; fi32++ {
            v25 += -234
            v3 += 42
            fmt.Println(v23)
            if (-2.25 < 1.5) {
                if true {
                    v28 = append(v28, 0)
                    v4 = ((v4 / v4) + (v4 - v4))
                }
                v4 -= (3.14159 * (((-2.25 * 1.5) - v4) + v4))
            } else {
                fmt.Println(v23)
                switch 32767 {
                case 42:
                    fmt.Println(fi29)
                case -2147483648:
                    if true {
                        fmt.Println(v25, v4, fi29)
                    } else {
                        v3 = fi29
                    }
                default:
                    var v33 int = v25
                    _ = v33
                }
            }
            _ = fi32
        }
        _ = fi29
    }
    if (false || true) {
        v34 := GenS1{F0: false, F1: 1, F2: -32768}
        fmt.Println(v3)
        v2 = v2
        _ = v34
    } else {
        v1[413] = v25
        if ((false || false) || (!true)) {
            v35 := GenS4{F0: 32767, F1: 42, F2: v23}
            if true {
                if (uint32(v24) > 0) {
                    var v36 float32 = v4
                    var v37 int = v25
                    _ = v36
                    _ = v37
                }
                for fi38 := 0; fi38 < 4; fi38++ {
                    v24 -= (v24 - (v24 - v24))
                    fmt.Println(v2, v2, v2)
                    v4 = v4
                    _ = fi38
                }
                var v39 uint32 = 1
                _ = v39
            }
            _ = v35
        }
    }
    for fi40 := 0; fi40 < 12; fi40++ {
        v41 := GenS1{F0: true, F1: 4294967295, F2: -1}
        fmt.Println(v3, v2, v4)
        _ = fi40
        _ = v41
    }
    var v42 uint16 = 42
    switch (v3 % ((32767) | 1)) {
    case 515:
        if ("" > "hello") {
            fmt.Println(v4, v2, v25)
        } else {
            fmt.Println(v42)
            v3 = v3
            v23 -= v23
        }
    case 9223372036854775807:
        v2 -= 0
    case 453:
        var v43 uint8 = v24
        _ = v43
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v21
    _ = v23
    _ = v24
    _ = v25
    _ = v28
    _ = v3
    _ = v4
    _ = v42
}
