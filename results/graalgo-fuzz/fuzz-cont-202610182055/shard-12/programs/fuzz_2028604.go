package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 string; F1 int32; F2 int64 }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 uint16; F1 uint32 }
type GenS5 struct { F0 uint8; F1 int8 }
type GenS6 struct { F0 int64; F1 uint; F2 int8 }
func genFn0(p0 uint16, p1 int16, p2 uint) uint8 {
    var v7 uint16 = 785
    v8 := GenS1{F0: 465, F1: 2}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 15
}
func genFn1(p0 int) uint64 {
    switch "" {
    case "hello":
        v9 := GenS1{F0: 2, F1: 42}
        _ = v9
    case "\n\t":
        switch "a" {
        case "\n\t":
            p0 -= (p0 | 128)
        case "a":
            v10 := []int{}
            _ = v10
        }
    }
    p0 += 762
    switch p0 {
    case -65:
        if false {
            for fi11 := 0; fi11 < 10; fi11++ {
                var v12 float32 = 3.14159
                v13 := GenS2{F0: "\n\t", F1: 0, F2: 1}
                var v14 int32 = 0
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
            }
            fmt.Println(p0)
        } else {
            v15 := GenS3{F0: p0, F1: "a"}
            _ = v15
        }
    case 0:
        var v16 int = p0
        _ = v16
    }
    _ = p0
    return 818
}

func main() {
    var v1 uint64 = 1
    var v2 string = ((("hello" + "") + (("\n\t" + "\n\t") + ("foo bar" + "foo bar"))) + ("\n\t" + "hello"))
    v3 := []int{}
    v4 := make(map[uint32]bool)
    var v5 int32 = 0
    var v6 int16 = 32767
    for fi17 := 0; fi17 < 6; fi17++ {
        v2 = (v2 + (((("hello" + (v2 + ("\n\t" + v2))) + "hello") + v2) + (("hello" + (("hello" + "a") + "foo bar")) + "hello")))
        _ = fi17
    }
    for fi18 := 0; fi18 < 13; fi18++ {
        v19 := GenS4{F0: 65535, F1: 460}
        switch "hello" {
        case "a":
            var v20 int16 = v6
            _ = v20
        default:
            switch fi18 {
            case 1:
                var v21 int32 = (^(v5))
                _ = v21
            default:
                if (3.14159 >= 0.0) {
                    for fi22 := 0; fi22 < 2; fi22++ {
                        v23 := GenS4{F0: 295, F1: 410}
                        var v24 int = 2147483647
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                } else {
                    switch 4294967295 {
                    case -1:
                        if (v5 >= v5) {
                            fmt.Println(v1, fi18)
                            if (!(((float64((float32(((0.0 * 1.5) / 3.14159)) / 1.5)) * ((-2.25 + -2.25) - -2.25)) != 0.0) && (42 < 877))) {
                                delete(v4, 2)
                            }
                        }
                    case 2147483647:
                        if true {
                            if true {
                                fmt.Println(v2, v1)
                                switch "a" {
                                case "hello":
                                    var v25 int8 = 127
                                    _ = v25
                                }
                            } else {
                                v26 := GenS2{F0: "\n\t", F1: v5, F2: -1}
                                v27 := GenS5{F0: 204, F1: 1}
                                _ = v26
                                _ = v27
                            }
                        } else {
                            var v28 uint64 = (v1 * v1)
                            v29 := GenS3{F0: fi18, F1: v2}
                            _ = v28
                            _ = v29
                        }
                    }
                }
            }
        }
        v30 := GenS6{F0: 1, F1: 2, F2: -70}
        _ = fi18
        _ = v19
        _ = v30
    }
    for fi31 := 0; fi31 < 3; fi31++ {
        v5 = (v5 * (v5 ^ v5))
        if false {
            v32 := GenS2{F0: v2, F1: v5, F2: -1}
            if (v6 > v6) {
                var v33 uint64 = v1
                v5 = (v5 + (v5 | (-(v5))))
                _ = v33
            } else {
                switch (fi31 >> ((352) & 63)) {
                case -438:
                    var v34 uint64 = (v1 * v1)
                    _ = v34
                }
                var v35 int = (-(fi31))
                _ = v35
            }
            switch "hello" {
            case "":
                var v36 uint = 2
                _ = v36
            case "\n\t":
                var v37 float32 = ((0.1 - 0.0) + (0.1 * 1e10))
                _ = v37
            case "foo bar":
                for fi38 := 0; fi38 < 2; fi38++ {
                    if (genFn0(882, (v6 >> ((840) & 63)), 4294967295) == genFn0(uint16(v1), -204, 1)) {
                        var v39 uint8 = 0
                        var v40 bool = ((-(fi38)) >= (^(fi31)))
                        _ = v39
                        _ = v40
                    }
                    _ = fi38
                }
            }
            _ = v32
        } else {
            _ = len(v4)
            fmt.Println(v6, v1, v6)
            var v41 string = (((v2 + "\n\t") + "hello") + (v2 + ("\n\t" + (v2 + ""))))
            _ = v41
        }
        _ = fi31
    }
    var v42 uint8 = 0
    v5 -= v5
    var v43 float64 = (1.5 + (3.14159 + 3.14159))
    v6 += (v6 << ((18446744073709551615) & 63))
    v5 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
