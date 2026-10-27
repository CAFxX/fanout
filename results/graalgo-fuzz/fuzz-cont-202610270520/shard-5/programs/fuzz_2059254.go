package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 int; F1 int64 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float32 = -2.25
    var v3 int32 = 0
    v4 := []int{1, 368, 0}
    v5 := []uint32{617, 794}
    var v6 int8 = 1
    for fi7 := 0; fi7 < 9; fi7++ {
        v3 -= v3
        v6 = 123
        switch fi7 {
        case 9223372036854775807:
            for fi8 := 0; fi8 < 18; fi8++ {
                v3 = 413
                v9 := GenS1{F0: false, F1: 90}
                fmt.Println(fi8, fi8, v6)
                _ = fi8
                _ = v9
            }
        case 0:
            fmt.Println(fi7)
        case -128:
            var v10 uint16 = 525
            _ = v10
        }
        _ = fi7
    }
    if false {
        var v11 uint32 = 1
        var v12 uint32 = (v11 % ((v11) | 1))
        var v13 uint64 = 4294967295
        _ = v11
        _ = v12
        _ = v13
    } else {
        v3 = v3
        v4 = append(v4, -388)
        v14 := GenS2{F0: -9223372036854775808, F1: -1}
        _ = v14
    }
    var v15 uint8 = 120
    if (!((2147483648 != 777) && false)) {
        var v16 int16 = 1000
        _ = v16
    } else {
        for fi17 := 0; fi17 < 3; fi17++ {
            v15 = v15
            var v18 uint32 = 959
            for fi19 := 0; fi19 < 10; fi19++ {
                v1[65535] = 4294967295
                if (693 < 18446744073709551615) {
                    v2 = v2
                    fmt.Println(v15)
                    var v20 int8 = -1
                    _ = v20
                }
                _ = fi19
            }
            _ = fi17
            _ = v18
        }
        fmt.Println(v6, v2)
        if false {
            switch ("\n\t" + "a") {
            case "foo bar":
                fmt.Println(v15)
            }
            if true {
                v3 = v3
                var v21 uint8 = (v15 % (((v15 & 98)) | 1))
                fmt.Println(v6, v21)
                _ = v21
            } else {
                v3 = (v3 % ((v3) | 1))
            }
        } else {
            if (!false) {
                v6 -= v6
            }
            v22 := GenS2{F0: 0, F1: 823}
            _ = v22
        }
    }
    if ((!(true && (false || false))) && false) {
        v23 := GenS1{F0: false, F1: v15}
        _ = v23
    }
    var v24 float64 = 0.0
    var v25 string = "foo bar"
    for fi26 := 0; fi26 < 16; fi26++ {
        v27 := GenS1{F0: (fi26 >= (fi26 | (fi26 & fi26))), F1: v15}
        v6 = v6
        var v28 bool = true
        _ = fi26
        _ = v27
        _ = v28
    }
    v29 := GenS1{F0: (((-2.25 / 1.5) + (3.14159 + float64(v2))) >= v24), F1: 255}
    var v30 int32 = (v3 & v3)
    fmt.Println(v24)
    v2 -= ((0.1 * v2) + v2)
    fmt.Println(v25, v3)
    if true {
        var v31 string = "a"
        _ = v31
    } else {
        fmt.Println(v2, v30)
    }
    v30 -= (v30 + v30)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
