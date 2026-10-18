package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 string }
type GenS2 struct { F0 uint64; F1 int8; F2 int8 }
type GenS3 struct { F0 uint16 }

func main() {
    var v1 bool = (1e10 <= 1.5)
    v2 := make(map[string]bool)
    var v3 int = 32768
    var v4 uint16 = 1
    switch 128 {
    case -1:
        if (0.0 < (0.0 - (0.1 / 1e10))) {
            v1 = (2147483648 != 31)
        } else {
            v5 := GenS1{F0: -1, F1: v3, F2: "a"}
            v6 := []int{(v3 / (((v3 | v3)) | 1)), 350, v3, (v3 - 1)}
            for fi7 := 0; fi7 < 14; fi7++ {
                v4 = v4
                v8 := GenS1{F0: 1, F1: fi7, F2: "a"}
                _ = fi7
                _ = v8
            }
            _ = v5
            _ = v6
        }
    case 42:
        for fi9 := 0; fi9 < 9; fi9++ {
            var v10 bool = v1
            _ = fi9
            _ = v10
        }
    case -336:
        v11 := []int{32768, (v3 ^ (v3 << ((18446744073709551615) & 63))), (v3 - (^(v3))), (v3 ^ v3)}
        _ = v11
    default:
        switch (v3 >> ((620) & 63)) {
        case 32768:
            v3 = (v3 - v3)
        case 9223372036854775807:
            v3 += 32768
        case -394:
            v12 := []int{}
            _ = v12
        }
    }
    fmt.Println(v4, v1)
    var v13 int32 = 2147483647
    var v14 bool = ((v1 && v1) || (v4 == v4))
    if (!(!true)) {
        delete(v2, ("" + ("a" + (("foo bar" + "a") + ""))))
        v2[("hello" + ("foo bar" + (("a" + "\n\t") + "a")))] = true
        if (v14 && true) {
            var v15 uint32 = 2
            _ = v15
        } else {
            var v16 int32 = v13
            if (!false) {
                v13 = (^(v16))
                if (-1 != -9223372036854775808) {
                    fmt.Println(v4)
                    var v17 int64 = 42
                    fmt.Println(v16, v14, v1)
                    _ = v17
                }
            }
            if (32767 > 1000) {
                v1 = v1
                if false {
                    var v18 int64 = -905
                    _ = v18
                }
                var v19 string = ("a" + (("hello" + ("" + "a")) + ("a" + ("\n\t" + (("a" + "hello") + ("hello" + ""))))))
                _ = v19
            }
            _ = v16
        }
    } else {
        v20 := []int{(^(v3)), (v3 >> ((18446744073709551615) & 63))}
        var v21 uint64 = 934
        var v22 int32 = (v13 << ((529) & 63))
        _ = v20
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 5; fi23++ {
        v4 += (v4 / (((^(v4))) | 1))
        var v24 uint8 = 128
        var v25 int8 = 42
        _ = fi23
        _ = v24
        _ = v25
    }
    for fi26 := 0; fi26 < 16; fi26++ {
        v13 = (v13 % ((2147483647) | 1))
        _ = fi26
    }
    var v27 float64 = 0.1
    if (v1 && false) {
        v2[""] = v14
    } else {
        switch (v3 >> ((2) & 63)) {
        case -611:
            v2["hello"] = true
        }
    }
    var v28 uint64 = 580
    var v29 uint8 = 128
    switch (v3 / (((-(v3))) | 1)) {
    case 32767:
        for fi30 := 0; fi30 < 3; fi30++ {
            fmt.Println(v29, v14, v4)
            v31 := []int{(v3 >> ((255) & 63)), fi30, (v3 | 32768)}
            _ = fi30
            _ = v31
        }
    case 0:
        if ("a" > ("hello" + "hello")) {
            if (float64(3.14159) >= (v27 - 1e10)) {
                v32 := GenS2{F0: v28, F1: 127, F2: 0}
                v33 := GenS3{F0: v4}
                v14 = (!(v14 || (!v14)))
                _ = v32
                _ = v33
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
