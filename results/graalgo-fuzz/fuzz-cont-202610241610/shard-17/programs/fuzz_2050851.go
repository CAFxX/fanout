package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 int32 }
type GenS2 struct { F0 int8; F1 bool; F2 uint64 }
type GenS3 struct { F0 bool; F1 int16 }
func genFn0() uint16 {
    var v6 uint32 = 2147483647
    v6 = (^(v6))
    var v7 int8 = 66
    switch "\n\t" {
    case "hello":
        fmt.Println(v7, v7)
    default:
        if false {
            var v8 int8 = (^(v7))
            var v9 uint16 = 65535
            _ = v8
            _ = v9
        }
    }
    var v10 float32 = (((0.0 * 0.1) - 3.14159) * -2.25)
    _ = v10
    _ = v6
    _ = v7
    return 65535
}
func genFn1() string {
    var v11 float64 = 0.0
    var v12 float32 = ((1e10 / float32(v11)) * ((0.0 * (3.14159 * -2.25)) * (((1.5 + 1e10) - 1e10) * (((0.0 + float32((1.5 + v11))) / (1.5 / 1.5)) * float32(0.0)))))
    _ = v11
    _ = v12
    return ""
}

func main() {
    var v1 int32 = -2147483648
    var v2 uint8 = 179
    var v3 bool = true
    v4 := []int{141, 413}
    v5 := []string{"\n\t", ("foo bar" + ("\n\t" + ("\n\t" + ""))), (("foo bar" + "a") + "foo bar"), "foo bar"}
    for fi13 := 0; fi13 < 13; fi13++ {
        var v14 int16 = 1000
        _ = fi13
        _ = v14
    }
    v3 = false
    switch "a" {
    case "\n\t":
        switch 127 {
        case -2147483648:
            v15 := GenS1{F0: "\n\t", F1: 2147483647, F2: (-(v1))}
            _ = v15
        case 2147483647:
            v2 -= 1
        case -9223372036854775808:
            if (637 < 682) {
                for fi16 := 0; fi16 < 17; fi16++ {
                    var v17 float32 = (0.1 + 0.1)
                    if (int8(v1) < int8(v1)) {
                        v1 = (v1 >> ((2) & 63))
                    }
                    _ = fi16
                    _ = v17
                }
                for fi18 := 0; fi18 < 7; fi18++ {
                    for fi19 := 0; fi19 < 1; fi19++ {
                        var v20 uint32 = 0
                        _ = fi19
                        _ = v20
                    }
                    v21 := GenS1{F0: "foo bar", F1: 1, F2: 1}
                    _ = fi18
                    _ = v21
                }
            } else {
                var v22 int8 = 100
                v23 := GenS1{F0: "foo bar", F1: 42, F2: (v1 ^ (v1 << ((42) & 63)))}
                var v24 float64 = (-2.25 + float64(float32(1e10)))
                _ = v22
                _ = v23
                _ = v24
            }
        }
    case "foo bar":
        v25 := GenS2{F0: 127, F1: v3, F2: 42}
        _ = v25
    }
    fmt.Println(v3, v1)
    if (1 >= 0) {
        switch "\n\t" {
        case "":
            v26 := GenS3{F0: false, F1: 42}
            _ = v26
        case "a":
            switch ("hello" + ("\n\t" + ("" + ""))) {
            case "\n\t":
                v27 := GenS3{F0: true, F1: -1}
                _ = v27
            case "":
                var v28 float32 = 1e10
                _ = v28
            case "foo bar":
                if false {
                    switch -2147483648 {
                    case 32767:
                        for fi29 := 0; fi29 < 13; fi29++ {
                            fmt.Println(v1)
                            var v30 float64 = 0.0
                            v30 = v30
                            _ = fi29
                            _ = v30
                        }
                    default:
                        var v31 int = -312
                        _ = v31
                    }
                } else {
                    var v32 int32 = -2147483648
                    v2 += v2
                    var v33 uint64 = 1
                    _ = v32
                    _ = v33
                }
            }
        default:
            var v34 int = 127
            _ = v34
        }
        v1 += v1
    }
    var v35 uint8 = (v2 % (((v2 * v2)) | 1))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v1 = 296
    for fi36 := 0; fi36 < 2; fi36++ {
        v37 := GenS3{F0: (2147483648 > 65535), F1: 759}
        v3 = (204 == 42)
        _ = fi36
        _ = v37
    }
    v1 -= v1
    v35 += (v35 % ((v35) | 1))
    var v38 int8 = 1
    var v39 int = -419
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
