package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 int16 }
type GenS2 struct { F0 uint32; F1 bool; F2 bool }
type GenS3 struct { F0 int32; F1 int32 }
type GenS4 struct { F0 int64; F1 uint64; F2 uint16 }
func genFn0(p0 uint, p1 uint16, p2 int8) int8 {
    var v4 uint64 = 4294967295
    v5 := []int{32767}
    p0 = (p0 | p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 1
}
func genFn1(p0 uint8, p1 int64, p2 uint64) int64 {
    var v6 int32 = -2147483648
    if (1 >= (v6 * -1)) {
        v7 := []int{-9223372036854775808, 32768, 1, -1}
        p2 -= 18446744073709551615
        _ = v7
    }
    var v8 int16 = 32767
    v9 := []int{}
    p2 = (p2 % (((^(p2))) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v8
    _ = v9
    return p1
}

func main() {
    v1 := []string{"foo bar", (("a" + "hello") + "foo bar"), "a", "hello"}
    var v2 uint32 = 2147483648
    v3 := []uint32{(v2 | v2), (v2 % ((v2) | 1)), v2, 42}
    v2 += 4294967295
    var v10 string = ""
    var v11 uint64 = 18446744073709551615
    switch (v10 + v10) {
    case "hello":
        switch -249 {
        case 961:
            v2 -= v2
        case -199:
            if true {
                var v12 int8 = -6
                if true {
                    v2 = v2
                }
                if len(v3) > 0 {
                    v3[0] = v2
                }
                _ = v12
            } else {
                v2 = v2
                var v13 float32 = (-2.25 * 0.1)
                v10 = v10
                _ = v13
            }
        }
    case "foo bar":
        v10 = v10
    default:
        fmt.Println(v10, v2)
    }
    v14 := GenS1{F0: -219, F1: (!false), F2: -718}
    v15 := GenS2{F0: v2, F1: (32767 != 32767), F2: (true || true)}
    var v16 int64 = 42
    _ = len(v3)
    for fi17 := 0; fi17 < 16; fi17++ {
        var v18 bool = false
        var v19 int32 = 1
        _ = fi17
        _ = v18
        _ = v19
    }
    fmt.Println(v2, v2)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v20 := GenS2{F0: 42, F1: (-1 != 42), F2: true}
    fmt.Println(v2, v10)
    _ = len(v3)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch "foo bar" {
    case "foo bar":
        if false {
            for fi21 := 0; fi21 < 20; fi21++ {
                var v22 uint32 = (^(v2))
                for fi23 := 0; fi23 < 7; fi23++ {
                    var v24 int32 = int32(fi23)
                    var v25 int8 = 65
                    v26 := GenS1{F0: 470, F1: false, F2: 1}
                    _ = fi23
                    _ = v24
                    _ = v25
                    _ = v26
                }
                switch "\n\t" {
                case "a":
                    v27 := GenS3{F0: -546, F1: -222}
                    _ = v27
                default:
                    switch "\n\t" {
                    case "hello":
                        var v28 uint64 = v11
                        _ = v28
                    default:
                        var v29 string = (("a" + v10) + v10)
                        _ = v29
                    }
                }
                _ = fi21
                _ = v22
            }
            v3 = append(v3, 0)
            v30 := GenS3{F0: -2147483648, F1: 2147483647}
            _ = v30
        } else {
            fmt.Println(v2, v16, v16)
        }
    case "":
        v16 = v16
    case "a":
        if (3.14159 < 3.14159) {
            if false {
                var v31 uint8 = 127
                v32 := GenS1{F0: 127, F1: (-1 < 9223372036854775807), F2: 27}
                _ = v31
                _ = v32
            } else {
                if false {
                    var v33 int8 = -14
                    if false {
                        var v34 bool = (v16 < v16)
                        _ = v34
                    }
                    v3 = append(v3, (v2 & (^(v2))))
                    _ = v33
                } else {
                    v35 := GenS4{F0: (-(v16)), F1: (v11 - (v11 & v11)), F2: 255}
                    fmt.Println(v10, v16, v10)
                    var v36 string = "hello"
                    _ = v35
                    _ = v36
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                fmt.Println(v2, v16, v10)
            }
            var v37 int64 = genFn1(127, (v16 % ((v16) | 1)), v11)
            _ = v37
        } else {
            v38 := GenS4{F0: v16, F1: (v11 | (v11 & (v11 % ((42) | 1)))), F2: 2}
            v39 := GenS4{F0: 15, F1: 18446744073709551615, F2: 42}
            var v40 float64 = 1.5
            _ = v38
            _ = v39
            _ = v40
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v3
}
