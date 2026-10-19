package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 int; F1 int; F2 uint32 }

func main() {
    var v1 uint = 193
    v2 := []string{}
    var v3 int16 = 42
    v1 = 0
    v3 += v3
    for fi4 := 0; fi4 < 10; fi4++ {
        switch fi4 {
        case -196:
            var v5 uint64 = 757
            _ = v5
        }
        switch fi4 {
        case 128:
            v3 = v3
        case 0:
            v6 := GenS1{F0: -1, F1: true}
            _ = v6
        default:
            var v7 string = ("foo bar" + "hello")
            _ = v7
        }
        for fi8 := 0; fi8 < 13; fi8++ {
            if (false && (!((false && true) && false))) {
                v9 := GenS1{F0: 1, F1: (42 != -101)}
                _ = v9
            }
            _ = fi8
        }
        _ = fi4
    }
    if ("\n\t" > "") {
        v1 = (v1 % ((v1) | 1))
    }
    v1 = (v1 | (v1 >> ((0) & 63)))
    fmt.Println(v1, v1)
    fmt.Println(v1, v3)
    if false {
        fmt.Println(v3)
    }
    if true {
        switch -1 {
        case -349:
            var v10 int64 = -1
            _ = v10
        }
        v11 := GenS1{F0: -1, F1: false}
        v1 = (v1 * (v1 + 485))
        _ = v11
    } else {
        v1 = v1
    }
    var v12 uint = v1
    _ = len(v2)
    switch 32767 {
    case 9223372036854775807:
        v3 = (v3 & (v3 >> (((v12 - 4294967295)) & 63)))
    case 4294967295:
        if ("a" >= "hello") {
            for fi13 := 0; fi13 < 20; fi13++ {
                if false {
                    v1 = 4294967295
                    var v14 float32 = 1.5
                    var v15 float64 = 1.5
                    _ = v14
                    _ = v15
                } else {
                    if len(v2) > 0 {
                        v2[0] = ("a" + (((("\n\t" + "a") + "a") + "a") + (("a" + ("" + "")) + "")))
                    }
                    switch (-(fi13)) {
                    case 2147483647:
                        var v16 float32 = 0.0
                        _ = v16
                    case 127:
                        var v17 string = ("foo bar" + ("hello" + "a"))
                        _ = v17
                    case -864:
                        v18 := GenS1{F0: 1, F1: true}
                        _ = v18
                    }
                    v19 := GenS1{F0: -9223372036854775808, F1: (2 > 765)}
                    _ = v19
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v20 float64 = float64(0.0)
                _ = fi13
                _ = v20
            }
            switch -2147483648 {
            case -28:
                v1 = (v12 ^ 0)
            default:
                if (true || true) {
                    var v21 int32 = -248
                    switch "" {
                    case "":
                        _ = len(v2)
                    case "foo bar":
                        v3 += 649
                    }
                    v1 = (v1 % ((441) | 1))
                    _ = v21
                } else {
                    v12 -= (v1 - v1)
                    v12 = (v12 * (v12 << (((v1 & (v12 >> ((v12) & 63)))) & 63)))
                    var v22 float64 = 0.1
                    _ = v22
                }
            }
            if true {
                v1 = v12
            }
        }
    default:
        v1 -= v1
    }
    if ((!(42 == 92)) && (3.14159 >= 0.1)) {
        if (false || (256 <= v12)) {
            var v23 uint32 = 42
            _ = v23
        }
        v12 = (^(v12))
    } else {
        var v24 int32 = 1
        _ = v24
    }
    var v25 uint64 = 1
    switch "" {
    case "a":
        v26 := GenS2{F0: -1, F1: 42, F2: 65535}
        _ = v26
    }
    var v27 float64 = 0.1
    v27 -= ((float64(0.0) + v27) - (v27 - float64(float32(v27))))
    v28 := GenS1{F0: 31, F1: (-9223372036854775808 >= 9223372036854775807)}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v2
    _ = v25
    _ = v27
    _ = v28
    _ = v3
}
