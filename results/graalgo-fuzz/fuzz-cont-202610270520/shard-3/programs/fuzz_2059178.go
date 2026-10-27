package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 int32; F1 uint64 }
type GenS5 struct { F0 uint8 }
func genFn0(p0 int32, p1 int32) int32 {
    p1 -= (p0 + p1)
    fmt.Println(p0)
    v4 := GenS1{F0: "foo bar"}
    for fi5 := 0; fi5 < 3; fi5++ {
        for fi6 := 0; fi6 < 20; fi6++ {
            var v7 uint16 = 32767
            if ((!(true || false)) || (1 >= 9223372036854775807)) {
                p0 = (-(p1))
                var v8 int64 = -245
                switch (fi5 / ((fi6) | 1)) {
                case -671:
                    p1 += 324
                }
                _ = v8
            } else {
                switch ("" + "") {
                case "hello":
                    p0 = (^(p1))
                case "\n\t":
                    if (!(!(int64(v7) < 303))) {
                        switch "foo bar" {
                        case "foo bar":
                            v9 := []int{1, fi6, fi6, (^(fi6))}
                            _ = v9
                        case "a":
                            var v10 float32 = float32(-2.25)
                            _ = v10
                        }
                    }
                default:
                    v11 := GenS1{F0: ""}
                    _ = v11
                }
            }
            _ = fi6
            _ = v7
        }
        var v12 uint8 = 255
        _ = fi5
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v4
    return (p1 + (p1 / ((-39) | 1)))
}

func main() {
    var v1 bool = (("hello" <= "") && (!(!false)))
    v2 := make(map[string]string)
    v3 := make(map[string]string)
    fmt.Println(v1)
    v13 := GenS2{F0: 0}
    var v14 bool = false
    v15 := []int{4294967295, -1, 9223372036854775807}
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    var v16 string = "\n\t"
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    v17 := GenS3{F0: 42, F1: 992}
    if true {
        if (true || v14) {
            v1 = (!(!true))
            var v18 float64 = 3.14159
            switch 4294967295 {
            case 32768:
                if (2147483647 == 501) {
                    switch -2147483648 {
                    case -934:
                        v19 := v3[(v16 + (v16 + v16))]
                        _ = v19
                    }
                }
            case 0:
                if (408 != 65535) {
                    v20 := GenS4{F0: 677, F1: 2}
                    _ = v20
                } else {
                    fmt.Println(v14)
                }
            case -728:
                switch "" {
                case "hello":
                    v21 := GenS5{F0: 205}
                    _ = v21
                case "a":
                    var v22 uint16 = 255
                    _ = v22
                }
            default:
                v23 := GenS3{F0: 793, F1: 687}
                _ = v23
            }
            _ = v18
        } else {
            var v24 uint16 = 65535
            _ = v24
        }
    }
    v15 = append(v15, -2147483648)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
