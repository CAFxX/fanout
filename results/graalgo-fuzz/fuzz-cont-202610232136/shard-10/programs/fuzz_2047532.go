package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 uint16; F1 int16; F2 string }
type GenS3 struct { F0 int32; F1 uint }
type GenS4 struct { F0 bool; F1 int8; F2 uint64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint8 }
func genFn0(p0 bool) int32 {
    p0 = false
    p0 = p0
    switch "hello" {
    case "foo bar":
        var v6 int = -9223372036854775808
        _ = v6
    case "\n\t":
        var v7 uint8 = 128
        _ = v7
    case "a":
        var v8 int = 32767
        _ = v8
    default:
        for fi9 := 0; fi9 < 17; fi9++ {
            switch (fi9 & (fi9 * 0)) {
            case -1:
                for fi10 := 0; fi10 < 7; fi10++ {
                    v11 := GenS1{F0: 255, F1: "a"}
                    v12 := []int{fi9, fi10, fi9}
                    fmt.Println(fi9)
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            }
            _ = fi9
        }
    }
    _ = p0
    return -1
}

func main() {
    var v1 int8 = 112
    v2 := make(map[string]string)
    var v3 uint32 = 42
    v4 := []string{"foo bar"}
    var v5 float64 = 0.0
    if ("a" != "") {
        switch -413 {
        case 2147483647:
            for fi13 := 0; fi13 < 7; fi13++ {
                v1 = (v1 * v1)
                if (true || (false && (94 > 32767))) {
                    var v14 uint32 = (^(v3))
                    _ = v14
                } else {
                    switch fi13 {
                    case 128:
                        switch "foo bar" {
                        case "a":
                            v15 := GenS2{F0: 42, F1: 1, F2: ("\n\t" + "a")}
                            _ = v15
                        case "":
                            v3 += v3
                        case "\n\t":
                            var v16 uint8 = 127
                            _ = v16
                        }
                    case 394:
                        for fi17 := 0; fi17 < 16; fi17++ {
                            var v18 uint8 = 92
                            var v19 int32 = 320
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                    }
                    delete(v2, "\n\t")
                }
                _ = fi13
            }
        case 32767:
            if false {
                switch -9223372036854775808 {
                case 1:
                    v5 = (float64(((((1.5 * 1e10) - 3.14159) * ((1.5 * -2.25) * -2.25)) / float32((v5 * ((float64(1e10) / v5) / v5))))) + float64((-2.25 / 1.0)))
                default:
                    fmt.Println(v1, v1)
                }
            } else {
                var v20 string = "foo bar"
                var v21 uint8 = 0
                var v22 int8 = v1
                _ = v20
                _ = v21
                _ = v22
            }
        case 1:
            delete(v2, "a")
        default:
            fmt.Println(v1)
        }
    }
    v23 := GenS1{F0: 1, F1: "\n\t"}
    if (true || false) {
        for fi24 := 0; fi24 < 18; fi24++ {
            var v25 float64 = 1.5
            var v26 int = 961
            _ = fi24
            _ = v25
            _ = v26
        }
    } else {
        v27 := GenS1{F0: 42, F1: ""}
        var v28 uint = 18446744073709551615
        v5 -= (v5 + float64(1e10))
        _ = v27
        _ = v28
    }
    var v29 uint = 187
    v30 := GenS3{F0: 0, F1: (^(v29))}
    for fi31 := 0; fi31 < 5; fi31++ {
        v32 := GenS4{F0: (false && true), F1: (v1 * (^(v1))), F2: uint64(v29)}
        var v33 uint16 = 255
        var v34 uint32 = v3
        _ = fi31
        _ = v32
        _ = v33
        _ = v34
    }
    _ = len(v4)
    v29 -= (v29 << ((v29) & 63))
    v1 = (-(v1))
    switch "" {
    case "hello":
        v35 := GenS5{F0: 255}
        _ = v35
    }
    for fi36 := 0; fi36 < 16; fi36++ {
        fmt.Println(v29, v3, v3)
        switch fi36 {
        case 93:
            v37 := GenS6{F0: -32768}
            _ = v37
        case -2147483648:
            v5 = 1e10
        }
        _ = fi36
    }
    v1 = (v1 & v1)
    v5 = v5
    v38 := GenS7{F0: 127}
    fmt.Println(v29, v3)
    v29 = uint(v3)
    fmt.Println(v29)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v29
    _ = v3
    _ = v30
    _ = v38
    _ = v4
    _ = v5
}
