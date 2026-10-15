package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32; F1 bool; F2 uint }
type GenS4 struct { F0 bool }
func genFn0() bool {
    for fi7 := 0; fi7 < 7; fi7++ {
        fmt.Println(fi7)
        v8 := []int{fi7}
        _ = fi7
        _ = v8
    }
    var v9 int = 127
    var v10 uint16 = 117
    if false {
        for fi11 := 0; fi11 < 6; fi11++ {
            switch (^(v9)) {
            case -164:
                v9 = (v9 % ((-808) | 1))
            case 2147483647:
                var v12 uint = 2
                _ = v12
            }
            var v13 uint16 = (v10 << ((255) & 63))
            _ = fi11
            _ = v13
        }
    } else {
        v14 := GenS1{F0: (v10 ^ v10), F1: -128}
        v10 = 65535
        for fi15 := 0; fi15 < 14; fi15++ {
            fmt.Println(fi15)
            _ = fi15
        }
        _ = v14
    }
    _ = v10
    _ = v9
    return (!true)
}

func main() {
    v1 := make(map[string]string)
    v2 := []uint32{4294967295, 0}
    var v3 float32 = 0.1
    var v4 int32 = -2147483648
    v5 := make(map[int]uint32)
    var v6 uint32 = 1
    fmt.Println(v6)
    delete(v1, "foo bar")
    if false {
        switch "\n\t" {
        case "a":
            v3 -= 1.5
        case "hello":
            v4 -= 42
        case "foo bar":
            if true {
                if true {
                    v16 := GenS1{F0: 805, F1: -128}
                    _ = v16
                } else {
                    fmt.Println(v6, v4)
                    v17 := GenS1{F0: 0, F1: -31}
                    var v18 bool = false
                    _ = v17
                    _ = v18
                }
            }
        }
        var v19 int8 = 91
        v20 := GenS2{F0: ""}
        _ = v19
        _ = v20
    }
    v21 := GenS3{F0: v4, F1: ((0.0 * 1.5) != -2.25), F2: 42}
    v2 = append(v2, v6)
    v6 += (^(v6))
    var v22 int8 = -1
    fmt.Println(v3, v4)
    for fi23 := 0; fi23 < 17; fi23++ {
        v24 := GenS4{F0: false}
        _ = fi23
        _ = v24
    }
    fmt.Println(v6, v4, v22)
    _ = len(v5)
    var v25 float64 = 3.14159
    if ((!false) || (true && ((!((v22 - v22) == v22)) && ("" != "a")))) {
        for fi26 := 0; fi26 < 6; fi26++ {
            if true {
                v27 := GenS4{F0: (18446744073709551615 != 35)}
                v2 = append(v2, v6)
                v28 := GenS1{F0: 505, F1: v22}
                _ = v27
                _ = v28
            }
            fmt.Println(v3)
            _ = fi26
        }
        var v29 int16 = -718
        v3 -= v3
        _ = v29
    }
    v25 = v25
    if (!false) {
        switch "hello" {
        case "":
            v25 = 0.0
        case "a":
            if ((false || true) && (721 == 462)) {
                var v30 string = ("\n\t" + "foo bar")
                _ = v30
            } else {
                delete(v5, -2147483648)
                _ = len(v5)
                var v31 int32 = v4
                _ = v31
            }
        default:
            var v32 int8 = (v22 + v22)
            _ = v32
        }
        var v33 uint16 = 2
        fmt.Println(v4, v25)
        _ = v33
    } else {
        v4 = 0
        v22 -= int8(v6)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
