package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16; F1 int8 }
func genFn0(p0 int8, p1 bool) bool {
    p1 = ((float32((0.0 * -2.25)) / 1.5) >= (-2.25 + ((1.5 * (float32(-2.25) - 3.14159)) + (-2.25 + (1.5 / 1.0)))))
    var v7 uint32 = 140
    _ = p0
    _ = p1
    _ = v7
    return p1
}
func genFn1(p0 uint32) int8 {
    fmt.Println(p0)
    var v8 float32 = float32((1.5 / (1.5 / float64(0.1))))
    var v9 uint32 = p0
    for fi10 := 0; fi10 < 8; fi10++ {
        if (925 != 1) {
            var v11 uint = 349
            _ = v11
        }
        v9 -= 42
        _ = fi10
    }
    p0 -= (^(p0))
    _ = p0
    _ = v8
    _ = v9
    return -121
}

func main() {
    v1 := []int{0, 284, -53}
    v2 := []string{(("foo bar" + ("foo bar" + "")) + (("a" + ("\n\t" + "\n\t")) + "foo bar")), "foo bar", ("hello" + ((("hello" + "hello") + "a") + "hello")), "a"}
    v3 := []uint32{}
    v4 := []int{127, -745, 32767}
    var v5 uint16 = 255
    v6 := []int{127}
    var v12 uint64 = 923
    v12 += v12
    v5 = (v5 / ((410) | 1))
    var v13 uint32 = 361
    if (-62 != 127) {
        var v14 int8 = int8(v12)
        _ = v14
    }
    var v15 uint32 = v13
    for fi16 := 0; fi16 < 6; fi16++ {
        switch "a" {
        case "\n\t":
            switch ("a" + ("hello" + "")) {
            case "hello":
                fmt.Println(v12, v5, v13)
            case "":
                switch 2147483647 {
                case 796:
                    var v17 string = ("a" + ("\n\t" + "foo bar"))
                    _ = v17
                case 127:
                    switch "a" {
                    case "hello":
                        v18 := GenS1{F0: -793, F1: 99, F2: 0}
                        _ = v18
                    case "foo bar":
                        v5 -= v5
                    case "":
                        v19 := GenS2{F0: true}
                        _ = v19
                    default:
                        for fi20 := 0; fi20 < 17; fi20++ {
                            v15 += 2
                            v15 += (v13 ^ uint32(v5))
                            _ = fi20
                        }
                    }
                default:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
            case "a":
                var v21 int64 = -463
                _ = v21
            }
        }
        switch ("a" + "\n\t") {
        case "\n\t":
            v13 += (^(v15))
        }
        _ = fi16
    }
    v22 := GenS2{F0: true}
    switch -2147483648 {
    case 32767:
        for fi23 := 0; fi23 < 8; fi23++ {
            var v24 uint32 = (v15 % ((65535) | 1))
            if false {
                for fi25 := 0; fi25 < 14; fi25++ {
                    v12 = (^(v12))
                    _ = fi25
                }
            }
            for fi26 := 0; fi26 < 18; fi26++ {
                if (255 == 255) {
                    if true {
                        var v27 float32 = 0.1
                        v5 -= (v5 | (v5 ^ 1))
                        v28 := GenS3{F0: -32768, F1: -128}
                        _ = v27
                        _ = v28
                    }
                    v12 = 2
                }
                v1 = append(v1, fi23)
                _ = fi26
            }
            _ = fi23
            _ = v24
        }
    default:
        var v29 uint32 = v13
        _ = v29
    }
    if len(v2) > 0 {
        v2[0] = ("\n\t" + "foo bar")
    }
    v12 = v12
    var v30 float64 = 1.5
    fmt.Println(v30, v30)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v22
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
