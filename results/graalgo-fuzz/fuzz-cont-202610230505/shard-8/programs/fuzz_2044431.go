package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16; F1 uint32; F2 int }
func genFn0(p0 uint64) bool {
    var v7 bool = (!false)
    v8 := []int{1, -375, -322}
    var v9 int16 = 1
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return (!((0 >= 711) || (0 > 2)))
}
func genFn1(p0 int32, p1 uint8, p2 uint8) uint {
    v10 := GenS1{F0: (true && false), F1: p1, F2: 2}
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 97
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []string{}
    v3 := []string{"hello"}
    var v4 int32 = 42
    var v5 uint = 18446744073709551615
    var v6 float64 = 1.5
    v6 = v6
    _ = len(v1)
    v6 += v6
    if (false && true) {
        v5 -= (^(v5))
        v5 += genFn1((v4 ^ v4), 74, 0)
    }
    v11 := GenS1{F0: (0 != 255), F1: 42, F2: uint32(v5)}
    if ("" == "") {
        v2 = append(v2, ((("a" + ("a" + ("a" + "\n\t"))) + ("foo bar" + "\n\t")) + "foo bar"))
        v5 = 56
    }
    for fi12 := 0; fi12 < 20; fi12++ {
        fmt.Println(fi12, v5)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        delete(v1, fi12)
        _ = fi12
    }
    fmt.Println(v4, v6)
    if (true || false) {
        for fi13 := 0; fi13 < 16; fi13++ {
            if len(v2) > 0 {
                v2[0] = ""
            }
            if ((9223372036854775807 <= 9223372036854775807) && (0.0 == ((3.14159 / (((1.5 * 1.5) + 1.5) + 1e10)) * ((3.14159 + (1e10 + (1e10 / -2.25))) / 0.1)))) {
                if (-2147483648 >= v4) {
                    var v14 int8 = -128
                    v4 -= 637
                    switch ("foo bar" + ("hello" + (("\n\t" + "\n\t") + "a"))) {
                    case "a":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case "hello":
                        var v15 uint32 = 234
                        _ = v15
                    }
                    _ = v14
                } else {
                    var v16 int32 = (v4 ^ (-(v4)))
                    fmt.Println(v5, fi13)
                    _ = v16
                }
                var v17 string = "\n\t"
                _ = v17
            } else {
                v18 := GenS1{F0: true, F1: 255, F2: 4294967295}
                _ = v18
            }
            if (false || (-1 == 822)) {
                v6 -= v6
                v4 -= v4
            } else {
                v19 := GenS1{F0: false, F1: 1, F2: 42}
                if true {
                    v5 = (^(v5))
                    v20 := GenS1{F0: (float32(v6) <= float32(v6)), F1: 186, F2: 336}
                    _ = v20
                }
                _ = v19
            }
            _ = fi13
        }
        switch "hello" {
        case "\n\t":
            switch 4294967295 {
            case -9223372036854775808:
                var v21 bool = false
                _ = v21
            case 9223372036854775807:
                v22 := GenS2{F0: 732}
                _ = v22
            }
        case "a":
            for fi23 := 0; fi23 < 10; fi23++ {
                for fi24 := 0; fi24 < 20; fi24++ {
                    fmt.Println(v6)
                    var v25 uint64 = 18446744073709551615
                    fmt.Println(v5)
                    _ = fi24
                    _ = v25
                }
                _ = fi23
            }
        default:
            var v26 int32 = v4
            _ = v26
        }
        var v27 uint = genFn1(-1, 255, 2)
        _ = v27
    }
    if true {
        for fi28 := 0; fi28 < 13; fi28++ {
            var v29 uint16 = 42
            switch fi28 {
            case -9223372036854775808:
                v5 += (v5 * genFn1((v4 * v4), 42, 255))
            case 725:
                if true {
                    for fi30 := 0; fi30 < 9; fi30++ {
                        var v31 float64 = (float64(((1.5 + 1.5) / 1.0)) - 0.1)
                        v32 := GenS2{F0: (fi28 / ((32768) | 1))}
                        v5 -= v5
                        delete(v1, 32767)
                        if false {
                            v33 := GenS3{F0: v29, F1: 443, F2: fi28}
                            _ = v33
                        } else {
                            v31 = (((1.5 + (v6 + v31)) + v6) - 3.14159)
                            var v34 float32 = 1.5
                            _ = v34
                        }
                        _ = fi30
                        _ = v31
                        _ = v32
                    }
                }
            }
            _ = fi28
            _ = v29
        }
        v6 -= 0.1
    } else {
        var v35 uint8 = 63
        _ = v35
    }
    v2 = append(v2, "\n\t")
    _ = len(v1)
    for fi36 := 0; fi36 < 10; fi36++ {
        delete(v1, -591)
        v6 += v6
        _ = fi36
    }
    delete(v1, -2147483648)
    if len(v3) > 0 {
        v3[0] = "hello"
    }
    v5 -= v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
