package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 int32; F1 int8 }
type GenS5 struct { F0 uint64; F1 int32; F2 int32 }

func main() {
    v1 := make(map[string]string)
    v2 := []int{-9223372036854775808, 32768}
    var v3 bool = (!(!(-526 > 1000)))
    delete(v1, "foo bar")
    for fi4 := 0; fi4 < 15; fi4++ {
        v5 := GenS1{F0: (909 == 2147483647), F1: 42, F2: -32768}
        if (112 <= -1) {
            v3 = true
            v6 := GenS2{F0: "a"}
            switch (-(fi4)) {
            case -366:
                fmt.Println(fi4)
            }
            _ = v6
        }
        switch (-(fi4)) {
        case 2147483647:
            v7 := GenS2{F0: ("hello" + ("\n\t" + ("foo bar" + "hello")))}
            _ = v7
        case 128:
            v3 = true
        }
        _ = fi4
        _ = v5
    }
    v3 = v3
    _ = len(v2)
    v8 := GenS3{F0: 0, F1: 32768}
    v9 := GenS1{F0: (v3 && false), F1: 1, F2: 0}
    for fi10 := 0; fi10 < 2; fi10++ {
        var v11 uint64 = 850
        _ = fi10
        _ = v11
    }
    fmt.Println(v3)
    if ((0 != 255) || (-210 > 42)) {
        v12 := GenS4{F0: -2147483648, F1: 0}
        for fi13 := 0; fi13 < 16; fi13++ {
            if (0 != 1000) {
                var v14 uint32 = 345
                _ = v14
            } else {
                if (true || false) {
                    var v15 uint64 = 42
                    _ = v15
                }
                fmt.Println(v3)
            }
            for fi16 := 0; fi16 < 20; fi16++ {
                var v17 float64 = 3.14159
                _ = fi16
                _ = v17
            }
            switch (fi13 - fi13) {
            case 799:
                if v3 {
                    var v18 uint32 = 127
                    _ = v18
                } else {
                    switch fi13 {
                    case 32767:
                        var v19 uint16 = 255
                        _ = v19
                    case 170:
                        v20 := GenS5{F0: 754, F1: 2147483647, F2: 2147483647}
                        _ = v20
                    default:
                        v3 = (!(2 < 902))
                    }
                }
            case 127:
                v3 = ((3.14159 / 1.0) >= (((1.5 * 0.0) / 1.0) + 1e10))
            default:
                switch (fi13 % (((fi13 << ((1) & 63))) | 1)) {
                case 682:
                    fmt.Println(v3, fi13)
                }
            }
            _ = fi13
        }
        _ = v12
    }
    for fi21 := 0; fi21 < 14; fi21++ {
        v22 := GenS4{F0: 60, F1: 127}
        _ = fi21
        _ = v22
    }
    if (!(42 < -1)) {
        switch 683 {
        case 1:
            v23 := GenS3{F0: 0, F1: -419}
            _ = v23
        case 32767:
            _ = len(v2)
        }
        v3 = (!true)
    } else {
        var v24 int16 = -793
        _ = v24
    }
    var v25 uint8 = 185
    var v26 int64 = 549
    for fi27 := 0; fi27 < 13; fi27++ {
        v28 := v1["a"]
        var v29 uint32 = 2147483647
        _ = fi27
        _ = v28
        _ = v29
    }
    var v30 bool = (true && ("foo bar" != ("a" + "a")))
    for fi31 := 0; fi31 < 11; fi31++ {
        v3 = (v3 || (-1 != -32768))
        _ = fi31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v8
    _ = v9
}
