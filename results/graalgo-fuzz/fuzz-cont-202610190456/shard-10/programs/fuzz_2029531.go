package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint16; F1 string }
type GenS5 struct { F0 string; F1 int16; F2 int32 }
type GenS6 struct { F0 int32; F1 string; F2 bool }
type GenS7 struct { F0 int8; F1 string; F2 int16 }
func genFn0(p0 uint8) int {
    p0 -= p0
    p0 = p0
    if true {
        var v7 uint32 = 923
        _ = v7
    }
    var v8 uint64 = 4294967295
    _ = p0
    _ = v8
    return 4294967295
}
func genFn1(p0 int64, p1 uint64) uint16 {
    var v9 int16 = 1
    v10 := []int{genFn0(1), -9223372036854775808}
    v11 := GenS1{F0: 0, F1: 42, F2: 2}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v9
    return 0
}

func main() {
    var v1 uint64 = 730
    var v2 int32 = 0
    var v3 uint64 = v1
    var v4 uint16 = 1
    var v5 uint8 = 2
    var v6 uint = 476
    if false {
        fmt.Println(v5)
        var v12 uint64 = (^(v3))
        var v13 int8 = -71
        _ = v12
        _ = v13
    }
    var v14 int16 = 1
    fmt.Println(v6, v4)
    if false {
        for fi15 := 0; fi15 < 13; fi15++ {
            v16 := []int{genFn0((v5 << ((4) & 63))), fi15, fi15}
            v2 = v2
            fmt.Println(fi15, v5, v5)
            _ = fi15
            _ = v16
        }
        if true {
            v17 := GenS2{F0: 2}
            v6 -= 0
            _ = v17
        } else {
            var v18 uint8 = (v5 * (v5 >> ((2) & 63)))
            _ = v18
        }
    } else {
        v19 := GenS3{F0: -1}
        if false {
            v6 += (v6 << ((689) & 63))
            var v20 int8 = 1
            var v21 int32 = v2
            _ = v20
            _ = v21
        } else {
            fmt.Println(v6, v6, v5)
            v22 := []int{664, 406, 2147483647}
            for fi23 := 0; fi23 < 4; fi23++ {
                v5 -= v5
                v6 += v6
                v4 += (v4 ^ v4)
                _ = fi23
            }
            _ = v22
        }
        _ = v19
    }
    var v24 uint32 = 54
    v2 = int32(v4)
    var v25 int16 = int16(v4)
    for fi26 := 0; fi26 < 6; fi26++ {
        v27 := GenS4{F0: (v4 - (^(v4))), F1: (("a" + ("a" + ("hello" + "foo bar"))) + ("" + ("hello" + ("" + ((("\n\t" + "\n\t") + ("a" + ("" + (("" + (("\n\t" + "hello") + "a")) + (("hello" + ("foo bar" + "\n\t")) + "foo bar"))))) + "foo bar")))))}
        var v28 int8 = 1
        for fi29 := 0; fi29 < 18; fi29++ {
            v4 = (^(v4))
            v1 = v3
            v4 -= 255
            if (127 < v5) {
                if false {
                    switch "a" {
                    case "":
                        v24 -= 0
                    default:
                        var v30 uint32 = (^(v24))
                        _ = v30
                    }
                    if true {
                        v31 := GenS5{F0: "", F1: (v25 - v14), F2: (v2 << (((v6 >> (((^(v6))) & 63))) & 63))}
                        var v32 uint64 = 18446744073709551615
                        _ = v31
                        _ = v32
                    }
                    v33 := GenS6{F0: (v2 + (^(v2))), F1: ("hello" + "foo bar"), F2: ((false || true) && true)}
                    _ = v33
                } else {
                    v34 := []int{1}
                    fmt.Println(v5, fi26)
                    fmt.Println(v2)
                    _ = v34
                }
                v35 := []int{(fi29 + (fi26 - genFn0((v5 - 1))))}
                v4 = 1
                var v36 int32 = (^(v2))
                var v37 float32 = float32((float64((1e10 - 1.5)) - ((float64((3.14159 - (1e10 - (1.5 + -2.25)))) * (1e10 - float64(0.0))) * 0.1)))
                _ = v35
                _ = v36
                _ = v37
            }
            if false {
                switch "" {
                case "foo bar":
                    switch "foo bar" {
                    case "":
                        v38 := GenS6{F0: v2, F1: "a", F2: false}
                        _ = v38
                    case "foo bar":
                        if ((^(v2)) >= v2) {
                            var v39 int64 = 1
                            _ = v39
                        }
                    }
                case "\n\t":
                    v5 = (v5 / ((v5) | 1))
                default:
                    v40 := GenS1{F0: 1, F1: v24, F2: genFn1(0, 18446744073709551615)}
                    _ = v40
                }
                fmt.Println(v1)
                var v41 uint32 = 2147483647
                _ = v41
            } else {
                v42 := []int{(fi29 & 9223372036854775807), fi26}
                if len(v42) > 0 {
                    v42[0] = -1
                }
                v28 = -109
                _ = v42
            }
            _ = fi29
        }
        _ = fi26
        _ = v27
        _ = v28
    }
    for fi43 := 0; fi43 < 8; fi43++ {
        var v44 int8 = -128
        var v45 int = (fi43 - int(v5))
        var v46 int = 4294967295
        _ = fi43
        _ = v44
        _ = v45
        _ = v46
    }
    v47 := GenS7{F0: 127, F1: "foo bar", F2: v25}
    v5 = v5
    var v48 int16 = v14
    v49 := GenS7{F0: 98, F1: "", F2: v14}
    for fi50 := 0; fi50 < 20; fi50++ {
        for fi51 := 0; fi51 < 7; fi51++ {
            var v52 float32 = 1.5
            _ = fi51
            _ = v52
        }
        if ((uint64(v6) >= uint64(v5)) || (fi50 < fi50)) {
            var v53 uint = uint(v24)
            _ = v53
        } else {
            for fi54 := 0; fi54 < 15; fi54++ {
                var v55 uint32 = (v24 + (^(v24)))
                _ = fi54
                _ = v55
            }
            var v56 float32 = 3.14159
            _ = v56
        }
        _ = fi50
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v6
}
