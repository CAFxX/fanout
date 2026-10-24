package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 uint64 }
type GenS2 struct { F0 uint16; F1 uint32; F2 int }
type GenS3 struct { F0 int16; F1 uint; F2 uint8 }
func genFn0(p0 int8) int16 {
    v5 := []int{9223372036854775807, 42}
    fmt.Println(p0)
    var v6 int = -9223372036854775808
    if (153 < 18446744073709551615) {
        v7 := GenS1{F0: 65535, F1: -59, F2: 1}
        _ = v7
    }
    _ = p0
    _ = v5
    _ = v6
    return 32767
}
func genFn1(p0 uint32, p1 bool, p2 uint64) bool {
    var v8 float32 = (float32((0.1 / 3.14159)) - 0.1)
    p1 = ((p1 && p1) && p1)
    var v9 int32 = 0
    var v10 bool = (!false)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return ((!v10) || (42 == (p2 * p2)))
}

func main() {
    v1 := []string{"foo bar", ("\n\t" + "a"), ("\n\t" + "")}
    v2 := []string{"", "a", "a"}
    var v3 int16 = -512
    v4 := make(map[int]string)
    if true {
        if false {
            v3 -= (^(v3))
        } else {
            v3 = v3
            v11 := GenS1{F0: 610, F1: int32(v3), F2: 1}
            _ = v11
        }
        _ = len(v2)
    }
    v3 += 0
    if true {
        v3 += v3
        var v12 uint = 261
        if false {
            var v13 float32 = 3.14159
            var v14 int8 = 43
            _ = v13
            _ = v14
        } else {
            var v15 uint = v12
            v3 = (-(v3))
            switch "" {
            case "foo bar":
                switch int(v3) {
                case -40:
                    var v16 uint32 = 4294967295
                    _ = v16
                case 127:
                    fmt.Println(v12, v12)
                default:
                    if true {
                        var v17 uint32 = 2147483648
                        v15 = (^(v15))
                        switch 32768 {
                        case 2147483647:
                            v18 := v4[1]
                            _ = v18
                        case 984:
                            var v19 string = ""
                            _ = v19
                        case -2147483648:
                            var v20 uint = 2
                            _ = v20
                        default:
                            for fi21 := 0; fi21 < 14; fi21++ {
                                v17 = 1
                                var v22 bool = (false && (!(true && (1 != 2))))
                                _ = fi21
                                _ = v22
                            }
                        }
                        _ = v17
                    } else {
                        v23 := GenS1{F0: 524, F1: 2147483647, F2: 2}
                        fmt.Println(v3, v12, v15)
                        var v24 uint16 = 901
                        _ = v23
                        _ = v24
                    }
                }
            case "hello":
                fmt.Println(v15, v3)
            case "\n\t":
                _ = len(v1)
            }
            _ = v15
        }
        _ = v12
    } else {
        var v25 int16 = v3
        for fi26 := 0; fi26 < 13; fi26++ {
            if false {
                for fi27 := 0; fi27 < 8; fi27++ {
                    v2 = append(v2, "hello")
                    if len(v2) > 0 {
                        v2[0] = "a"
                    }
                    var v28 int32 = -341
                    _ = fi27
                    _ = v28
                }
            } else {
                v29 := v4[fi26]
                var v30 float64 = float64(1.5)
                _ = v29
                _ = v30
            }
            fmt.Println(fi26, v25, v25)
            v31 := GenS2{F0: 563, F1: 65535, F2: fi26}
            _ = fi26
            _ = v31
        }
        _ = v25
    }
    v32 := GenS1{F0: 0, F1: 2147483647, F2: 1}
    var v33 bool = (!false)
    v34 := GenS3{F0: -403, F1: 0, F2: 127}
    v33 = false
    v35 := GenS3{F0: -1, F1: 255, F2: 128}
    for fi36 := 0; fi36 < 13; fi36++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi36
    }
    var v37 uint32 = 0
    v3 = (v3 | v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v37
    _ = v4
}
