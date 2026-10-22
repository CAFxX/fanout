package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8; F1 bool }
func genFn0() uint {
    for fi5 := 0; fi5 < 19; fi5++ {
        v6 := GenS1{F0: 2, F1: "foo bar"}
        _ = fi5
        _ = v6
    }
    var v7 int32 = 42
    var v8 bool = true
    var v9 uint64 = 202
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 uint16 = 32767
    v2 := []uint32{2147483648, 42, 42, 1}
    var v3 uint64 = 494
    v4 := make(map[int]bool)
    for fi10 := 0; fi10 < 7; fi10++ {
        var v11 int32 = int32(fi10)
        if true {
            fmt.Println(v3, v11)
            _ = len(v4)
            delete(v4, fi10)
        } else {
            var v12 int16 = -227
            for fi13 := 0; fi13 < 11; fi13++ {
                if ((!(81 > 42)) || (false && false)) {
                    v11 = v11
                    fmt.Println(v3, fi13)
                    v14 := GenS2{F0: -1}
                    _ = v14
                } else {
                    v1 -= 0
                    var v15 float32 = 0.1
                    v11 += -1
                    _ = v15
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi13
            }
            _ = v12
        }
        _ = fi10
        _ = v11
    }
    for fi16 := 0; fi16 < 5; fi16++ {
        var v17 float64 = (1e10 * (1.5 + 1.5))
        for fi18 := 0; fi18 < 9; fi18++ {
            v19 := GenS3{F0: 115}
            _ = fi18
            _ = v19
        }
        switch 127 {
        case 32767:
            if (true && (607 != 0)) {
                var v20 float32 = 1.5
                _ = v20
            }
        }
        _ = fi16
        _ = v17
    }
    v3 = v3
    var v21 uint8 = 0
    for fi22 := 0; fi22 < 16; fi22++ {
        v23 := GenS1{F0: 655, F1: ("foo bar" + ("a" + ("foo bar" + "hello")))}
        _ = fi22
        _ = v23
    }
    v1 = (v1 << ((1) & 63))
    var v24 int = 32767
    v25 := GenS4{F0: 86, F1: false}
    if (0.0 >= 1e10) {
        var v26 int32 = -2147483648
        if (744 > v3) {
            delete(v4, v24)
            var v27 uint16 = v1
            _ = v27
        }
        _ = v26
    }
    var v28 uint64 = (v3 / (((v3 / ((uint64(v1)) | 1))) | 1))
    switch "\n\t" {
    case "\n\t":
        v3 = 64
    default:
        v29 := v4[v24]
        _ = v29
    }
    switch ("a" + (("" + "\n\t") + ((("foo bar" + "a") + ("\n\t" + "")) + ((("a" + "hello") + ("hello" + ("" + "a"))) + "")))) {
    case "hello":
        _ = len(v2)
    default:
        v21 = v21
    }
    if ((v21 >> ((255) & 63)) != v21) {
        for fi30 := 0; fi30 < 20; fi30++ {
            if (!false) {
                var v31 bool = (-1 < 2147483647)
                _ = v31
            } else {
                v2 = append(v2, 1)
                if ((^(v1)) > (v1 - (v1 ^ 42))) {
                    if false {
                        var v32 float32 = (float32((float64(((3.14159 * 1.5) - (0.0 / 0.1))) * 1.5)) - (((float32((-2.25 - 0.0)) / 3.14159) - -2.25) + (((0.1 * -2.25) / (3.14159 + (1.5 * 0.1))) / 1.0)))
                        v33 := GenS4{F0: -1, F1: false}
                        _ = v32
                        _ = v33
                    } else {
                        var v34 uint16 = (v1 % ((v1) | 1))
                        _ = v34
                    }
                    fmt.Println(v28, v3)
                } else {
                    for fi35 := 0; fi35 < 7; fi35++ {
                        var v36 float32 = 0.1
                        v24 += 32767
                        _ = fi35
                        _ = v36
                    }
                    v28 -= (^(v28))
                }
                var v37 float32 = 0.1
                _ = v37
            }
            var v38 string = "foo bar"
            var v39 int64 = -26
            _ = fi30
            _ = v38
            _ = v39
        }
        fmt.Println(v21)
        fmt.Println(v1, v3, v28)
    } else {
        var v40 uint32 = 159
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v28
    _ = v3
    _ = v4
}
