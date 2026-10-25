package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint }

func main() {
    var v1 bool = true
    var v2 int64 = 42
    v3 := make(map[string]int)
    v4 := make(map[uint32]bool)
    var v5 uint8 = 1
    for fi6 := 0; fi6 < 1; fi6++ {
        switch -1 {
        case 127:
            for fi7 := 0; fi7 < 17; fi7++ {
                var v8 uint64 = 0
                _ = fi7
                _ = v8
            }
        case 512:
            if ((true && (!(!v1))) || (!(!(float32(-2.25) >= 1.5)))) {
                _ = len(v3)
            } else {
                var v9 int32 = 42
                v10 := GenS1{F0: -1, F1: (true && (-437 < 1))}
                _ = v10
                _ = v9
            }
        default:
            var v11 int64 = (v2 - -1)
            _ = v11
        }
        if (int64(fi6) != (v2 + -9223372036854775808)) {
            v12 := GenS1{F0: 1000, F1: (v1 || v1)}
            v1 = v1
            for fi13 := 0; fi13 < 10; fi13++ {
                if (2147483648 < 2) {
                    var v14 int = 42
                    v5 -= (v5 + v5)
                    fmt.Println(fi6, v2)
                    _ = v14
                } else {
                    if ((float32((1e10 * 1e10)) / 1e10) > 1.5) {
                        _ = len(v4)
                    }
                    var v15 int8 = -60
                    var v16 uint32 = 865
                    _ = v15
                    _ = v16
                }
                _ = fi13
            }
            _ = v12
        } else {
            v2 = v2
            fmt.Println(v1, fi6, v5)
        }
        _ = fi6
    }
    fmt.Println(v2, v5)
    _ = len(v4)
    for fi17 := 0; fi17 < 17; fi17++ {
        v4[853] = (!("foo bar" != "a"))
        if true {
            v18 := GenS2{F0: (v5 - (v5 ^ v5))}
            switch "a" {
            case "foo bar":
                v19 := GenS1{F0: 1, F1: v1}
                _ = v19
            default:
                v20 := []int{fi17, (fi17 & fi17), fi17}
                _ = v20
            }
            _ = v18
        } else {
            var v21 float32 = 1e10
            _ = v21
        }
        fmt.Println(v2)
        _ = fi17
    }
    v22 := []int{}
    fmt.Println(v5, v5, v2)
    v5 -= v5
    var v23 int = 128
    var v24 int = (v23 - (v23 | v23))
    v22 = append(v22, v24)
    if v1 {
        if true {
            fmt.Println(v2)
        }
        switch "hello" {
        case "foo bar":
            fmt.Println(v24)
        default:
            v25 := GenS3{F0: 18446744073709551615}
            _ = v25
        }
        _ = len(v22)
    } else {
        v26 := GenS1{F0: -32768, F1: false}
        v22 = append(v22, v23)
        _ = v26
    }
    if ((v5 + (^(v5))) != 255) {
        var v27 float64 = float64((float32(0.0) * 3.14159))
        switch 2147483647 {
        case 9223372036854775807:
            v23 += 4294967295
        }
        _ = v27
    }
    v1 = v1
    if true {
        switch v23 {
        case -149:
            v28 := GenS1{F0: 32767, F1: (2147483647 >= 1)}
            _ = v28
        case 30:
            fmt.Println(v24, v24)
        case 32767:
            switch "foo bar" {
            case "foo bar":
                var v29 uint16 = 65535
                _ = v29
            default:
                if (false && v1) {
                    for fi30 := 0; fi30 < 9; fi30++ {
                        for fi31 := 0; fi31 < 10; fi31++ {
                            v32 := GenS3{F0: 657}
                            _ = fi31
                            _ = v32
                        }
                        _ = fi30
                    }
                    var v33 int = v23
                    _ = v33
                }
            }
        }
        v3[(("a" + ("\n\t" + ("" + "\n\t"))) + "a")] = 2147483647
        v4[2] = (32767 < 645)
    } else {
        v34 := GenS2{F0: v5}
        _ = v34
    }
    var v35 bool = (!(42 <= 0))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
