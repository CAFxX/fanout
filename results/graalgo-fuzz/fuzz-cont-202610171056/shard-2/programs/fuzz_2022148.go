package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int64 }
type GenS2 struct { F0 bool; F1 int; F2 uint8 }
type GenS3 struct { F0 int64; F1 uint }
func genFn0(p0 uint64, p1 bool) int64 {
    p0 = (p0 * p0)
    fmt.Println(p0)
    _ = p0
    _ = p1
    return 232
}
func genFn1() uint16 {
    var v5 float32 = float32(0.1)
    var v6 uint16 = 65535
    switch (("" + "a") + ("foo bar" + "")) {
    case "\n\t":
        switch 32768 {
        case 1:
            v7 := GenS1{F0: 242, F1: 0, F2: -1}
            _ = v7
        }
    case "a":
        for fi8 := 0; fi8 < 7; fi8++ {
            v9 := GenS1{F0: 66, F1: 4294967295, F2: -448}
            var v10 float32 = (v5 + v5)
            if true {
                switch -9223372036854775808 {
                case -383:
                    fmt.Println(v10, v10)
                case -1:
                    var v11 float32 = (0.0 / v10)
                    _ = v11
                case 865:
                    v5 -= v10
                default:
                    v12 := []int{}
                    _ = v12
                }
                v6 = 32767
                var v13 uint32 = 4294967295
                _ = v13
            }
            _ = fi8
            _ = v10
            _ = v9
        }
    case "":
        fmt.Println(v5)
    default:
        fmt.Println(v5, v6)
    }
    _ = v5
    _ = v6
    return (v6 >> ((42) & 63))
}

func main() {
    var v1 int64 = -1
    var v2 int8 = 42
    var v3 uint64 = 18446744073709551615
    var v4 int32 = -343
    v14 := GenS1{F0: 1, F1: 0, F2: (v1 >> ((255) & 63))}
    if (("a" + "hello") != "hello") {
        v15 := []int{128, 444}
        var v16 uint16 = genFn1()
        _ = v15
        _ = v16
    } else {
        for fi17 := 0; fi17 < 3; fi17++ {
            var v18 int = (fi17 * fi17)
            v18 -= -32
            _ = fi17
            _ = v18
        }
    }
    fmt.Println(v3, v3)
    v19 := []int{1}
    var v20 bool = false
    var v21 uint32 = 273
    v22 := GenS1{F0: 386, F1: 255, F2: (v1 >> ((982) & 63))}
    var v23 uint = 1
    v20 = false
    for fi24 := 0; fi24 < 12; fi24++ {
        fmt.Println(v1)
        _ = fi24
    }
    v4 = 1
    if ((^(v2)) == (v2 - (v2 % (((v2 % ((-128) | 1))) | 1)))) {
        var v25 int8 = (v2 & 0)
        fmt.Println(v1, v23, v1)
        _ = v25
    } else {
        v26 := GenS1{F0: (v23 * v23), F1: (^(v23)), F2: (^(v1))}
        _ = v26
    }
    var v27 float32 = 3.14159
    for fi28 := 0; fi28 < 10; fi28++ {
        v21 += 42
        fmt.Println(v23, v2)
        if ((!v20) && v20) {
            v21 -= (v21 + (v21 + 755))
            v27 += v27
            var v29 uint64 = v3
            _ = v29
        } else {
            v19 = append(v19, 32768)
        }
        _ = fi28
    }
    if (int64(v23) < v1) {
        for fi30 := 0; fi30 < 5; fi30++ {
            if v20 {
                fmt.Println(v20, v4)
                for fi31 := 0; fi31 < 9; fi31++ {
                    if v20 {
                        v32 := GenS2{F0: ((("" == ("" + "")) && (v1 <= genFn0(v3, ((v21 % ((v21) | 1)) >= (v21 + v21))))) && v20), F1: -255, F2: 255}
                        v20 = (!false)
                        _ = v32
                    }
                    _ = fi31
                }
            }
            for fi33 := 0; fi33 < 10; fi33++ {
                var v34 float32 = v27
                switch "a" {
                case "foo bar":
                    var v35 bool = v20
                    _ = v35
                default:
                    fmt.Println(fi30)
                }
                v36 := GenS3{F0: v1, F1: uint(v3)}
                _ = fi33
                _ = v34
                _ = v36
            }
            _ = len(v19)
            _ = fi30
        }
        if len(v19) >= 2 {
            v19 = v19[1:len(v19)]
        }
        v1 = 0
    } else {
        var v37 uint32 = 86
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v27
    _ = v3
    _ = v4
}
