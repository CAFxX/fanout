package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 int16 }
type GenS2 struct { F0 int; F1 string }
func genFn0(p0 int32, p1 bool) uint {
    if (((float64(-2.25) + float64((0.0 + 1.5))) - -2.25) > ((float64(1.5) + (0.0 * (((1e10 - float64((float32(float64(0.0)) * 1.5))) / float64(-2.25)) + 1.5))) * 3.14159)) {
        p1 = ("a" < ((("hello" + ("hello" + ("a" + ""))) + "foo bar") + "foo bar"))
    } else {
        p0 -= p0
    }
    switch 79 {
    case 128:
        fmt.Println(p0, p1)
    }
    var v5 int64 = 42
    for fi6 := 0; fi6 < 2; fi6++ {
        v7 := []int{fi6, (fi6 * fi6), (fi6 * 740), (^(fi6))}
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v5
    return 4294967295
}
func genFn1() uint {
    var v8 uint16 = 315
    var v9 string = "\n\t"
    _ = v8
    _ = v9
    return 855
}

func main() {
    var v1 string = ("" + ("a" + (("a" + "") + "\n\t")))
    var v2 float32 = -2.25
    v3 := []uint32{}
    var v4 int = 2147483647
    for fi10 := 0; fi10 < 2; fi10++ {
        var v11 int = (fi10 & v4)
        var v12 float64 = (float64(float32(3.14159)) - (3.14159 / -2.25))
        _ = fi10
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 12; fi13++ {
        if true {
            if (v1 >= "") {
                var v14 bool = true
                switch ("a" + (v1 + v1)) {
                case "a":
                    for fi15 := 0; fi15 < 16; fi15++ {
                        switch 42 {
                        case 909:
                            var v16 uint32 = 0
                            _ = v16
                        default:
                            if (65535 > 42) {
                                switch (v1 + ((v1 + ((v1 + (v1 + "a")) + (((v1 + "") + "hello") + v1))) + v1)) {
                                case "hello":
                                    var v17 float32 = 0.0
                                    _ = v17
                                }
                                v4 -= fi15
                                v2 = v2
                            }
                        }
                        v1 = (v1 + ("a" + ((v1 + (v1 + "foo bar")) + v1)))
                        v18 := GenS1{F0: fi15, F1: (v1 + (("" + v1) + "\n\t")), F2: -1}
                        var v19 int64 = 963
                        _ = fi15
                        _ = v18
                        _ = v19
                    }
                case "":
                    var v20 uint = 42
                    _ = v20
                case "\n\t":
                    v4 = (fi13 >> ((2) & 63))
                default:
                    var v21 float32 = v2
                    _ = v21
                }
                v1 = ""
                _ = v14
            }
        }
        fmt.Println(v2, fi13)
        v22 := GenS1{F0: (-(fi13)), F1: v1, F2: 273}
        _ = fi13
        _ = v22
    }
    switch v1 {
    case "\n\t":
        for fi23 := 0; fi23 < 8; fi23++ {
            if false {
                v24 := GenS1{F0: 9223372036854775807, F1: ("" + ("" + (v1 + ((v1 + v1) + v1)))), F2: 0}
                fmt.Println(v1, v4, v1)
                v3 = append(v3, 1)
                _ = v24
            }
            _ = fi23
        }
    case "a":
        v25 := GenS1{F0: (v4 * (v4 % (((v4 % (((v4 | 32768)) | 1))) | 1))), F1: "hello", F2: 0}
        _ = v25
    }
    v1 = v1
    v1 = v1
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v3 = append(v3, 1)
    v1 = "foo bar"
    switch "hello" {
    case "":
        v26 := GenS2{F0: v4, F1: "\n\t"}
        _ = v26
    case "hello":
        v27 := GenS2{F0: 32768, F1: v1}
        _ = v27
    default:
        var v28 uint16 = 32767
        _ = v28
    }
    if (!(-1 != -40)) {
        switch (^(v4)) {
        case 0:
            v1 = v1
        case 42:
            var v29 int8 = -1
            _ = v29
        case 127:
            v4 += (v4 & (-(v4)))
        }
        v1 = "hello"
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
