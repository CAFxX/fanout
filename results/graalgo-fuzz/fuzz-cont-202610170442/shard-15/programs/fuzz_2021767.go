package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 bool }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 uint32; F1 uint8; F2 bool }
type GenS4 struct { F0 int; F1 uint }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 string; F1 string; F2 int8 }
func genFn0() uint32 {
    var v6 uint32 = 2147483647
    if false {
        var v7 int8 = 100
        v6 = (v6 % (((v6 / (((v6 ^ 2)) | 1))) | 1))
        _ = v7
    }
    if (v6 <= (v6 + (v6 << ((255) & 63)))) {
        v6 -= 2147483647
    } else {
        if (!true) {
            if false {
                var v8 uint = 736
                _ = v8
            }
            for fi9 := 0; fi9 < 16; fi9++ {
                switch 32767 {
                case 178:
                    fmt.Println(v6, fi9)
                case 4294967295:
                    v10 := []int{(fi9 ^ fi9)}
                    _ = v10
                default:
                    switch "\n\t" {
                    case "\n\t":
                        var v11 float64 = float64(3.14159)
                        _ = v11
                    case "hello":
                        var v12 int16 = -1
                        _ = v12
                    case "foo bar":
                        var v13 float64 = (1e10 / 1e10)
                        _ = v13
                    default:
                        v6 += 2147483647
                    }
                }
                fmt.Println(v6)
                var v14 float32 = 1.5
                _ = fi9
                _ = v14
            }
            fmt.Println(v6)
        }
    }
    _ = v6
    return v6
}
func genFn1(p0 bool) uint32 {
    var v15 int16 = 1
    v16 := GenS1{F0: -2147483648, F1: 255, F2: ((-(v15)) > -32768)}
    for fi17 := 0; fi17 < 10; fi17++ {
        switch (fi17 - fi17) {
        case 127:
            v18 := GenS2{F0: ("foo bar" + "\n\t"), F1: -38}
            _ = v18
        case 32767:
            switch ("a" + "") {
            case "a":
                p0 = p0
            case "hello":
                fmt.Println(p0, p0, p0)
            case "":
                var v19 int8 = 100
                _ = v19
            default:
                var v20 int = fi17
                _ = v20
            }
        }
        v21 := GenS2{F0: ("a" + (("hello" + ("a" + "")) + "hello")), F1: 42}
        v22 := []int{fi17, (fi17 >> ((65535) & 63)), -2147483648}
        _ = fi17
        _ = v21
        _ = v22
    }
    p0 = true
    _ = p0
    _ = v15
    _ = v16
    return 242
}

func main() {
    v1 := []string{(("" + "a") + "a"), "hello"}
    v2 := make(map[string]uint32)
    var v3 uint64 = 0
    v4 := []uint32{1, 0}
    var v5 float64 = 1e10
    v23 := GenS3{F0: uint32(v3), F1: 2, F2: false}
    var v24 uint64 = 18446744073709551615
    var v25 float64 = (1.5 + float64(((0.0 * 1e10) + 1.5)))
    var v26 uint = 383
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if true {
        var v27 float64 = (float64(-2.25) - (v5 / (v5 * 1e10)))
        fmt.Println(v27, v25, v24)
        _ = v27
    }
    if false {
        v2["hello"] = 563
    } else {
        switch -317 {
        case 42:
            v28 := GenS2{F0: "\n\t", F1: 127}
            _ = v28
        case -873:
            fmt.Println(v5, v3, v5)
        default:
            v24 += v24
        }
        v29 := GenS3{F0: 693, F1: 45, F2: ((-2.25 > (1e10 / 0.1)) && (!(false && true)))}
        _ = v29
    }
    if false {
        switch 32767 {
        case -2147483648:
            v30 := GenS2{F0: "", F1: 127}
            _ = v30
        }
    }
    fmt.Println(v5, v24)
    if len(v1) > 0 {
        v1[0] = ""
    }
    v25 = v5
    var v31 int8 = 105
    for fi32 := 0; fi32 < 9; fi32++ {
        for fi33 := 0; fi33 < 2; fi33++ {
            v2["hello"] = 65535
            var v34 uint32 = 582
            var v35 int = fi32
            _ = fi33
            _ = v34
            _ = v35
        }
        _ = fi32
    }
    if (!true) {
        v36 := GenS4{F0: 840, F1: (v26 ^ 4294967295)}
        _ = v36
    } else {
        v37 := GenS5{F0: v31}
        _ = v37
    }
    switch ("foo bar" + "\n\t") {
    case "":
        var v38 int8 = 100
        _ = v38
    default:
        v26 += (v26 * 4294967295)
    }
    var v39 int32 = -1
    v40 := GenS6{F0: v24}
    v41 := GenS7{F0: "hello", F1: "\n\t", F2: v31}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v31
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
