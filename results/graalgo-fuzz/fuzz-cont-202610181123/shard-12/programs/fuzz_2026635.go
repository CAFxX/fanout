package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 bool }
type GenS2 struct { F0 uint8; F1 int16; F2 uint8 }
type GenS3 struct { F0 uint8; F1 int64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 string }
func genFn0() bool {
    v6 := []int{-1, 115, 4294967295, -9223372036854775808}
    for fi7 := 0; fi7 < 17; fi7++ {
        for fi8 := 0; fi8 < 5; fi8++ {
            v9 := GenS1{F0: 1, F1: 360, F2: false}
            if len(v6) > 0 {
                v6[0] = fi7
            }
            if (!true) {
                fmt.Println(fi8, fi7)
            }
            v10 := GenS1{F0: 128, F1: 579, F2: (42 != 42)}
            switch "" {
            case "\n\t":
                fmt.Println(fi7, fi8)
            case "a":
                fmt.Println(fi7, fi7)
            case "foo bar":
                v6 = append(v6, (fi8 & (fi7 / ((fi7) | 1))))
            default:
                var v11 int64 = -1
                _ = v11
            }
            var v12 float64 = -2.25
            fmt.Println(fi8, fi8)
            _ = fi8
            _ = v10
            _ = v12
            _ = v9
        }
        if true {
            var v13 uint = 783
            if true {
                v13 = v13
                switch fi7 {
                case 128:
                    v13 = (v13 & (v13 & (^(v13))))
                }
                fmt.Println(v13, v13)
            } else {
                var v14 uint = 4294967295
                var v15 float32 = float32(3.14159)
                _ = v14
                _ = v15
            }
            for fi16 := 0; fi16 < 3; fi16++ {
                if (0 <= 517) {
                    v13 -= 727
                    v13 += v13
                }
                _ = fi16
            }
            _ = v13
        }
        v17 := GenS1{F0: 255, F1: 4294967295, F2: (false && false)}
        _ = fi7
        _ = v17
    }
    var v18 int32 = 562
    if false {
        var v19 uint = 0
        fmt.Println(v18, v18)
        fmt.Println(v18, v19)
        _ = v19
    }
    _ = v18
    _ = v6
    return false
}

func main() {
    v1 := []uint32{848}
    v2 := []int{0, -617, 42, 2147483647}
    var v3 float64 = 1.5
    v4 := []int{-810, 859}
    v5 := make(map[uint32]int)
    v3 += (v3 + v3)
    var v20 string = "foo bar"
    if ((-1 < 30) && (((v20 + "a") == ((("" + ("hello" + "\n\t")) + v20) + "")) || (v3 >= (0.1 * 3.14159)))) {
        if (!true) {
            v1 = append(v1, 1)
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        for fi21 := 0; fi21 < 20; fi21++ {
            switch v20 {
            case "foo bar":
                var v22 uint8 = 60
                _ = v22
            default:
                v20 = (v20 + "\n\t")
            }
            for fi23 := 0; fi23 < 1; fi23++ {
                if (v3 != v3) {
                    fmt.Println(v20, fi23)
                }
                fmt.Println(fi23)
                _ = fi23
            }
            _ = fi21
        }
    } else {
        if (618 <= 65535) {
            fmt.Println(v20)
        } else {
            if false {
                for fi24 := 0; fi24 < 4; fi24++ {
                    var v25 bool = (0 > 63)
                    _ = fi24
                    _ = v25
                }
            } else {
                for fi26 := 0; fi26 < 10; fi26++ {
                    var v27 int32 = 1
                    _ = fi26
                    _ = v27
                }
                if (((41 == 0) || (18446744073709551615 < 4294967295)) && false) {
                    v5[94] = 32768
                }
                switch ("a" + v20) {
                case "\n\t":
                    var v28 int64 = -9223372036854775808
                    _ = v28
                case "hello":
                    v29 := GenS2{F0: 2, F1: 42, F2: 127}
                    _ = v29
                case "":
                    v20 = v20
                default:
                    var v30 bool = (65535 >= 2)
                    _ = v30
                }
            }
        }
        v20 = (v20 + v20)
        v20 = "hello"
    }
    v31 := GenS1{F0: 219, F1: 1, F2: true}
    v4 = append(v4, -705)
    v32 := GenS2{F0: 127, F1: 916, F2: 141}
    v33 := GenS3{F0: 255, F1: 9223372036854775807}
    var v34 uint = 4294967295
    switch -9223372036854775808 {
    case -1:
        v35 := GenS4{F0: -479}
        _ = v35
    case 42:
        if false {
            if (true && (((v3 + v3) - v3) <= float64(3.14159))) {
                var v36 int8 = -128
                var v37 bool = (!false)
                _ = len(v5)
                _ = v36
                _ = v37
            } else {
                var v38 int64 = -975
                for fi39 := 0; fi39 < 10; fi39++ {
                    var v40 uint16 = 65535
                    _ = fi39
                    _ = v40
                }
                var v41 float64 = v3
                _ = v38
                _ = v41
            }
            v42 := GenS2{F0: 255, F1: 42, F2: 109}
            fmt.Println(v20, v3, v3)
            _ = v42
        } else {
            var v43 int16 = -814
            v44 := v5[42]
            _ = v43
            _ = v44
        }
    }
    if (!true) {
        var v45 int16 = 0
        _ = v45
    } else {
        var v46 uint32 = 2147483647
        _ = v46
    }
    switch (((v20 + "") + "a") + "") {
    case "a":
        v3 = (0.1 + 1.5)
    }
    fmt.Println(v3, v34)
    v3 = (v3 / (3.14159 + v3))
    if (((1.5 >= 1.5) || false) && false) {
        if (536 != 127) {
            v20 = "a"
            fmt.Println(v3, v34, v34)
            if true {
                v5[2] = 2147483647
            } else {
                if true {
                    v20 = v20
                    v34 -= v34
                } else {
                    if len(v1) > 0 {
                        v1[0] = 42
                    }
                }
                var v47 int16 = 0
                delete(v5, 0)
                _ = v47
            }
        } else {
            v48 := GenS3{F0: 136, F1: 483}
            v49 := GenS5{F0: "\n\t"}
            v50 := GenS1{F0: 16, F1: 1, F2: false}
            _ = v48
            _ = v49
            _ = v50
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
