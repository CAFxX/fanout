package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 int8; F1 string; F2 uint }
type GenS3 struct { F0 uint; F1 uint16 }
type GenS4 struct { F0 int32; F1 uint32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int16; F1 string }
func genFn0(p0 string) int32 {
    if (float32((0.1 / 3.14159)) >= 0.1) {
        switch -353 {
        case 32767:
            v6 := []int{2147483647, 42}
            _ = v6
        case 1:
            p0 = p0
        case -756:
            v7 := GenS1{F0: 255, F1: 976}
            _ = v7
        default:
            fmt.Println(p0)
        }
        p0 = (p0 + ((("foo bar" + p0) + "hello") + "\n\t"))
    }
    var v8 int64 = 9223372036854775807
    fmt.Println(p0)
    v9 := GenS1{F0: 127, F1: 2147483647}
    _ = p0
    _ = v8
    _ = v9
    return -1
}

func main() {
    var v1 int32 = -1
    v2 := []int{}
    var v3 int32 = 1
    var v4 uint16 = 42
    var v5 uint = 65535
    v4 = 255
    var v10 uint32 = 4294967295
    fmt.Println(v5, v3)
    var v11 bool = ((!false) || true)
    if (v11 && ((true || v11) || v11)) {
        v11 = (42 >= 734)
    }
    switch "hello" {
    case "hello":
        var v12 uint32 = 2147483648
        _ = v12
    case "\n\t":
        switch 42 {
        case -432:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case -331:
            var v13 uint = v5
            _ = v13
        }
    case "a":
        for fi14 := 0; fi14 < 10; fi14++ {
            v10 -= v10
            for fi15 := 0; fi15 < 18; fi15++ {
                var v16 string = ("a" + "foo bar")
                switch (fi15 << (((^(v5))) & 63)) {
                case 128:
                    if (604 <= 2) {
                        var v17 uint16 = 42
                        fmt.Println(v1, v10, v11)
                        if ((v4 | v4) != (v17 | v17)) {
                            if (991 != (v5 - (v5 + v5))) {
                                v18 := GenS2{F0: -128, F1: "hello", F2: 936}
                                _ = v18
                            } else {
                                v2 = append(v2, fi15)
                                v19 := GenS3{F0: (v5 & (^(v5))), F1: v17}
                                _ = v19
                            }
                            v4 -= v4
                            var v20 int = fi15
                            fmt.Println(v1, v20)
                            _ = v20
                        }
                        _ = v17
                    } else {
                        fmt.Println(v5, fi14)
                    }
                case 32768:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                default:
                    v5 = (v5 / ((v5) | 1))
                }
                if len(v2) > 0 {
                    v2[0] = -104
                }
                _ = fi15
                _ = v16
            }
            v5 = v5
            _ = fi14
        }
    default:
        v11 = (!(1 <= 1000))
    }
    for fi21 := 0; fi21 < 3; fi21++ {
        v22 := GenS4{F0: v3, F1: v10}
        var v23 float32 = 0.0
        v2 = append(v2, (fi21 / ((fi21) | 1)))
        _ = fi21
        _ = v22
        _ = v23
    }
    if len(v2) > 0 {
        v2[0] = 18
    }
    var v24 float32 = 1e10
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if (true && (-233 > -9223372036854775808)) {
        if (42 < 2) {
            if (!v11) {
                if v11 {
                    v25 := GenS5{F0: v1}
                    var v26 int = 701
                    for fi27 := 0; fi27 < 1; fi27++ {
                        v2 = append(v2, (fi27 ^ (fi27 / (((^(fi27))) | 1))))
                        _ = fi27
                    }
                    _ = v25
                    _ = v26
                }
            }
            var v28 uint = (^(v5))
            _ = v28
        }
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v29 uint16 = (v4 | v4)
    v30 := GenS6{F0: 42, F1: ("" + "hello")}
    v5 = v5
    v3 = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v24
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
