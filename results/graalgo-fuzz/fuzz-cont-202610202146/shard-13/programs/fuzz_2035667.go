package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 int32; F2 uint64 }
func genFn0() uint16 {
    var v6 uint16 = 0
    var v7 int8 = 127
    _ = v6
    _ = v7
    return v6
}

func main() {
    v1 := []uint32{0, 4294967295, 372, 0}
    var v2 int64 = 0
    v3 := make(map[int]string)
    var v4 uint32 = 2147483648
    var v5 string = "a"
    var v8 uint32 = v4
    var v9 uint8 = 245
    v10 := GenS1{F0: -3}
    if (!true) {
        for fi11 := 0; fi11 < 6; fi11++ {
            var v12 bool = true
            v13 := GenS1{F0: 127}
            v3[(^(fi11))] = ("\n\t" + v5)
            _ = fi11
            _ = v12
            _ = v13
        }
    } else {
        v5 = ("foo bar" + "\n\t")
    }
    fmt.Println(v9, v8)
    fmt.Println(v5, v9)
    var v14 int8 = 8
    var v15 int = 128
    if (1 > 18446744073709551615) {
        for fi16 := 0; fi16 < 2; fi16++ {
            if (!true) {
                var v17 bool = (true && true)
                v8 += v8
                for fi18 := 0; fi18 < 3; fi18++ {
                    fmt.Println(v17, v17)
                    var v19 uint8 = 57
                    _ = fi18
                    _ = v19
                }
                _ = v17
            } else {
                switch "\n\t" {
                case "a":
                    for fi20 := 0; fi20 < 7; fi20++ {
                        v5 = "foo bar"
                        v14 = (v14 << ((42) & 63))
                        fmt.Println(v5)
                        _ = fi20
                    }
                case "":
                    v8 -= v4
                case "foo bar":
                    v21 := GenS2{F0: v14, F1: -22, F2: 4294967295}
                    _ = v21
                }
            }
            v9 = 187
            v22 := GenS1{F0: (v14 * (^(v14)))}
            _ = fi16
            _ = v22
        }
    }
    _ = len(v3)
    var v23 uint16 = 32767
    v5 = v5
    var v24 int64 = -378
    v1 = append(v1, v8)
    v2 = (^(v2))
    if true {
        v25 := v3[int(v14)]
        if ((v8 - (v8 / ((v4) | 1))) <= 65535) {
            var v26 int32 = int32(v24)
            var v27 int32 = v26
            _ = v26
            _ = v27
        } else {
            var v28 uint8 = (v9 - v9)
            v29 := GenS2{F0: (v14 + v14), F1: -2147483648, F2: 4294967295}
            _ = v28
            _ = v29
        }
        _ = v25
    } else {
        v1 = append(v1, v4)
        for fi30 := 0; fi30 < 16; fi30++ {
            for fi31 := 0; fi31 < 20; fi31++ {
                var v32 uint8 = v9
                _ = fi31
                _ = v32
            }
            _ = fi30
        }
        var v33 uint8 = 1
        _ = v33
    }
    for fi34 := 0; fi34 < 5; fi34++ {
        v35 := v3[fi34]
        _ = fi34
        _ = v35
    }
    switch 119 {
    case 1:
        var v36 int8 = v14
        _ = v36
    case -1:
        var v37 int64 = v24
        _ = v37
    default:
        var v38 uint16 = (v23 % (((^(v23))) | 1))
        _ = v38
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
