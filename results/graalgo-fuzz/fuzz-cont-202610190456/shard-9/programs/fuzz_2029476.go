package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 int32 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 string) string {
    var v7 int16 = -32768
    if true {
        p0 = "hello"
    } else {
        if false {
            fmt.Println(v7, v7)
        }
    }
    var v8 string = p0
    for fi9 := 0; fi9 < 14; fi9++ {
        v10 := GenS1{F0: 718, F1: 79, F2: -2147483648}
        v11 := []int{(fi9 - fi9), -9223372036854775808, (^(fi9))}
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = p0
    _ = v7
    _ = v8
    return ""
}

func main() {
    var v1 bool = (!true)
    var v2 uint8 = 1
    v3 := make(map[uint32]string)
    var v4 uint = 65535
    var v5 uint32 = 2147483648
    v6 := make(map[string]uint32)
    if false {
        v12 := []int{}
        _ = len(v12)
        var v13 int64 = -9223372036854775808
        _ = v12
        _ = v13
    } else {
        if (-920 >= 2147483647) {
            if (11 != 37) {
                v14 := GenS1{F0: 0, F1: 100, F2: -1}
                _ = v14
            }
        } else {
            if v1 {
                for fi15 := 0; fi15 < 10; fi15++ {
                    if (v1 || v1) {
                        v5 = (^(v5))
                        v2 = (v2 * (v2 | (v2 >> (((v4 << (((v4 - (v4 & 659))) & 63))) & 63))))
                    } else {
                        fmt.Println(v5)
                        v1 = false
                    }
                    v6["a"] = uint32(v2)
                    v16 := []int{}
                    _ = fi15
                    _ = v16
                }
                var v17 int8 = 113
                v4 -= v4
                _ = v17
            }
            v6[("foo bar" + "\n\t")] = (v5 | v5)
            var v18 int = 1
            _ = v18
        }
        v1 = ((v5 < (^(v5))) && (!false))
    }
    v5 += 65535
    for fi19 := 0; fi19 < 13; fi19++ {
        var v20 float32 = (-2.25 / float32(1.5))
        _ = fi19
        _ = v20
    }
    var v21 int16 = 32767
    v22 := v3[(^(v5))]
    v2 += (^(v2))
    v23 := GenS1{F0: 42, F1: 127, F2: 0}
    var v24 int16 = (v21 | -1)
    var v25 float64 = (1.5 - 1e10)
    v26 := v6[v22]
    switch v22 {
    case "\n\t":
        if v1 {
            if v1 {
                var v27 int32 = -2147483648
                v26 = 2147483647
                _ = v27
            }
            fmt.Println(v25)
            fmt.Println(v25, v21, v24)
        }
    case "a":
        switch -555 {
        case 128:
            if (v1 || ((465 < (v21 ^ v21)) || ((v22 + "hello") > v22))) {
                var v28 int32 = -2147483648
                var v29 uint8 = 0
                _ = v28
                _ = v29
            }
        default:
            switch 128 {
            case 4294967295:
                if ((v1 || (true || (v1 && ((float64(1.5) - (v25 / 3.14159)) < v25)))) && v1) {
                    v4 = 65535
                }
            case 32767:
                for fi30 := 0; fi30 < 4; fi30++ {
                    var v31 string = ("" + v22)
                    if false {
                        var v32 bool = true
                        v33 := GenS2{F0: uint8(v26)}
                        delete(v6, v31)
                        _ = v32
                        _ = v33
                    }
                    v34 := GenS1{F0: 42, F1: -128, F2: 42}
                    _ = fi30
                    _ = v31
                    _ = v34
                }
            }
        }
    case "hello":
        for fi35 := 0; fi35 < 19; fi35++ {
            v36 := GenS2{F0: (^(v2))}
            var v37 float64 = 0.1
            _ = fi35
            _ = v36
            _ = v37
        }
    default:
        var v38 uint16 = 42
        _ = v38
    }
    var v39 float32 = 3.14159
    v21 -= 42
    v40 := []int{}
    v41 := GenS2{F0: v2}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
