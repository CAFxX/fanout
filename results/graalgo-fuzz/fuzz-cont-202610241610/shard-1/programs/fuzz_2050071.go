package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 int8; F2 uint }
type GenS3 struct { F0 int16; F1 uint8; F2 bool }

func main() {
    v1 := make(map[uint32]string)
    var v2 int = 9223372036854775807
    v3 := make(map[int]int)
    for fi4 := 0; fi4 < 10; fi4++ {
        v2 = v2
        if (46 >= 1) {
            v5 := GenS1{F0: 128}
            _ = v5
        } else {
            v6 := GenS1{F0: 176}
            _ = v6
        }
        for fi7 := 0; fi7 < 4; fi7++ {
            _ = len(v1)
            var v8 uint16 = 391
            if (452 <= 49) {
                v2 = -2147483648
                if true {
                    switch -954 {
                    case 32768:
                        if true {
                            v8 += (v8 - 0)
                            _ = len(v3)
                        }
                    case 32767:
                        v9 := []int{fi7}
                        _ = v9
                    default:
                        switch "foo bar" {
                        case "foo bar":
                            if (127 != 100) {
                                switch ("hello" + "hello") {
                                case "":
                                    var v10 int32 = -2147483648
                                    _ = v10
                                case "a":
                                    v11 := v1[724]
                                    _ = v11
                                default:
                                    v12 := GenS2{F0: 127, F1: -1, F2: 2}
                                    _ = v12
                                }
                            }
                        case "hello":
                            v2 -= 0
                            var v13 float32 = ((3.14159 * float32(((-2.25 - 0.0) / ((float64((3.14159 + 1.5)) / 1.5) + 0.0)))) + -2.25)
                            v8 = (v8 / ((v8) | 1))
                            var v14 int = 127
                            _ = v13
                            _ = v14
                        case "":
                            v3[(v2 | 128)] = (fi7 << ((65535) & 63))
                        default:
                            var v15 string = ("hello" + (("\n\t" + "foo bar") + (("" + "\n\t") + ("" + "foo bar"))))
                            _ = v15
                        }
                    }
                } else {
                    v2 = 32767
                    if false {
                        switch "foo bar" {
                        case "foo bar":
                            v2 = 9223372036854775807
                        case "hello":
                            var v16 uint64 = 4294967295
                            _ = v16
                        case "\n\t":
                            fmt.Println(fi4, fi4)
                        }
                        delete(v3, (fi7 ^ -1))
                    }
                    if true {
                        var v17 uint = 0
                        _ = v17
                    } else {
                        if true {
                            var v18 string = ""
                            fmt.Println(v18, v18)
                            v19 := v3[(fi4 ^ (-(v2)))]
                            v18 = v18
                            _ = v18
                            _ = v19
                        }
                    }
                }
            }
            var v20 int = 42
            _ = fi7
            _ = v20
            _ = v8
        }
        _ = fi4
    }
    switch ("\n\t" + ("a" + ("foo bar" + ""))) {
    case "\n\t":
        v2 = (v2 ^ (v2 << ((4294967295) & 63)))
    case "a":
        for fi21 := 0; fi21 < 6; fi21++ {
            fmt.Println(v2, fi21)
            delete(v1, 47)
            _ = fi21
        }
    default:
        v22 := []int{v2}
        _ = v22
    }
    for fi23 := 0; fi23 < 17; fi23++ {
        v24 := GenS1{F0: 255}
        _ = fi23
        _ = v24
    }
    v2 -= v2
    fmt.Println(v2)
    v2 -= v2
    v25 := GenS2{F0: 127, F1: 42, F2: 65535}
    v26 := []int{v2, v2, (v2 - v2), (v2 ^ 493)}
    v3[v2] = 464
    var v27 float64 = (1.5 * ((float64((3.14159 - 0.0)) / 1.0) - 3.14159))
    var v28 uint32 = 4294967295
    v29 := GenS3{F0: 19, F1: 0, F2: true}
    var v30 uint = 18446744073709551615
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
}
