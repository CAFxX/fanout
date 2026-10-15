package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 string) bool {
    var v5 uint64 = 0
    v6 := []int{}
    var v7 uint64 = v5
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return false
}

func main() {
    var v1 uint16 = 1
    v2 := []int{-2147483648, 128, 4294967295}
    v3 := make(map[uint32]bool)
    v4 := []string{"foo bar", "\n\t"}
    for fi8 := 0; fi8 < 6; fi8++ {
        fmt.Println(fi8)
        v1 = 32767
        if true {
            var v9 uint64 = 2
            v1 -= (^(v1))
            _ = v9
        }
        _ = fi8
    }
    var v10 uint32 = 974
    for fi11 := 0; fi11 < 7; fi11++ {
        for fi12 := 0; fi12 < 6; fi12++ {
            var v13 float32 = 0.0
            _ = fi12
            _ = v13
        }
        var v14 uint8 = 255
        _ = fi11
        _ = v14
    }
    var v15 int16 = 574
    for fi16 := 0; fi16 < 13; fi16++ {
        v1 -= (^(v1))
        var v17 bool = true
        var v18 bool = v17
        _ = fi16
        _ = v17
        _ = v18
    }
    var v19 int16 = v15
    v19 = (v19 + (v19 << ((907) & 63)))
    switch ("" + (("foo bar" + ("a" + "foo bar")) + "")) {
    case "\n\t":
        v4 = append(v4, "hello")
    case "a":
        v1 += (v1 % (((v1 >> ((255) & 63))) | 1))
    case "":
        var v20 int8 = -76
        _ = v20
    default:
        var v21 int = 128
        _ = v21
    }
    v1 += v1
    fmt.Println(v19, v15, v15)
    delete(v3, v10)
    var v22 int = -9223372036854775808
    v22 = v22
    v1 = v1
    if false {
        switch "hello" {
        case "":
            v19 -= 1000
        case "a":
            switch v22 {
            case 127:
                for fi23 := 0; fi23 < 14; fi23++ {
                    if (true || (false || (v1 != (v1 + (v1 - v1))))) {
                        v22 = int(v1)
                        var v24 uint16 = 65535
                        _ = v24
                    } else {
                        v4 = append(v4, ("a" + ((("a" + "hello") + "a") + ("a" + ("hello" + ("a" + "foo bar"))))))
                    }
                    v25 := GenS1{F0: (v15 + v15)}
                    if true {
                        if (3.14159 == 0.0) {
                            v19 = (v15 / ((int16(v1)) | 1))
                            if (65535 <= (v10 << ((2) & 63))) {
                                if (185 >= 2) {
                                    v22 -= v22
                                    var v26 float64 = 0.0
                                    _ = v26
                                } else {
                                    v1 += (v1 % (((^(v1))) | 1))
                                    var v27 bool = (2 > 65535)
                                    _ = v27
                                }
                            }
                            if len(v2) > 0 {
                                v2[0] = fi23
                            }
                        }
                        switch fi23 {
                        case 825:
                            v15 = v15
                        case -9223372036854775808:
                            switch "foo bar" {
                            case "\n\t":
                                _ = len(v4)
                            }
                        case 555:
                            var v28 int = int(v15)
                            _ = v28
                        }
                        fmt.Println(v10)
                    }
                    _ = fi23
                    _ = v25
                }
            default:
                v3[uint32(v1)] = true
            }
        }
        if true {
            v29 := GenS1{F0: v15}
            var v30 uint64 = 109
            _ = v29
            _ = v30
        }
        var v31 int16 = v15
        _ = v31
    }
    for fi32 := 0; fi32 < 3; fi32++ {
        var v33 float32 = float32(0.1)
        fmt.Println(v33, v33)
        var v34 string = "a"
        _ = fi32
        _ = v33
        _ = v34
    }
    v35 := GenS1{F0: v19}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v15
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v35
    _ = v4
}
