package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 uint8; F1 uint32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int32; F1 int8; F2 int }

func main() {
    var v1 bool = true
    var v2 bool = (32767 < 1)
    v3 := make(map[int]uint32)
    for fi4 := 0; fi4 < 16; fi4++ {
        var v5 uint8 = 255
        if (v1 && (!v1)) {
            for fi6 := 0; fi6 < 4; fi6++ {
                var v7 uint = 18446744073709551615
                var v8 bool = v2
                var v9 uint32 = 4294967295
                _ = fi6
                _ = v7
                _ = v8
                _ = v9
            }
            switch fi4 {
            case 128:
                if (v1 && ((1.5 + (0.0 - -2.25)) > float64((((((-2.25 - 3.14159) - (0.1 * 0.1)) - float32((0.1 + float64((-2.25 / 1.0))))) * -2.25) * float32(-2.25))))) {
                    v3[127] = 0
                    var v10 float32 = 3.14159
                    _ = v10
                } else {
                    for fi11 := 0; fi11 < 19; fi11++ {
                        fmt.Println(fi4)
                        v2 = v1
                        if false {
                            v1 = (v1 && v1)
                            v12 := v3[337]
                            v13 := v3[(fi11 << ((0) & 63))]
                            _ = v12
                            _ = v13
                        }
                        v14 := GenS1{F0: -128, F1: v1}
                        _ = fi11
                        _ = v14
                    }
                }
            case -2147483648:
                fmt.Println(v2, v1)
            case 0:
                v5 += (^(v5))
            }
        }
        v15 := GenS1{F0: 0, F1: v2}
        _ = fi4
        _ = v15
        _ = v5
    }
    var v16 int64 = 9223372036854775807
    v17 := GenS1{F0: 0, F1: ("hello" <= "hello")}
    fmt.Println(v2, v1)
    var v18 int8 = 1
    var v19 float32 = 0.0
    for fi20 := 0; fi20 < 1; fi20++ {
        v19 = (v19 - v19)
        _ = fi20
    }
    v1 = (!(9223372036854775807 != -353))
    var v21 uint64 = 2
    v22 := v3[127]
    for fi23 := 0; fi23 < 19; fi23++ {
        v22 += v22
        for fi24 := 0; fi24 < 20; fi24++ {
            var v25 float64 = 3.14159
            v21 -= (v21 - v21)
            if (65535 >= 255) {
                v26 := GenS1{F0: (v18 - v18), F1: (("" + "\n\t") != "hello")}
                switch fi23 {
                case 0:
                    v21 = v21
                }
                v27 := GenS2{F0: "a", F1: (v16 << ((0) & 63))}
                if v1 {
                    switch "foo bar" {
                    case "":
                        var v28 uint8 = 128
                        _ = v28
                    case "\n\t":
                        if true {
                            v29 := GenS1{F0: v18, F1: true}
                            if (331 <= 2) {
                                v1 = (!((!false) && (!v2)))
                                v1 = v1
                            }
                            var v30 int64 = (-(v16))
                            _ = v29
                            _ = v30
                        } else {
                            fmt.Println(fi24, fi24, v16)
                            v1 = true
                            var v31 uint64 = (v21 ^ 42)
                            _ = v31
                        }
                    }
                }
                _ = v26
                _ = v27
            } else {
                v32 := []int{(fi23 | int(v22)), (fi23 / (((^(fi23))) | 1)), -1, fi24}
                switch ("\n\t" + "hello") {
                case "":
                    var v33 uint8 = 42
                    _ = v33
                case "hello":
                    switch "a" {
                    case "a":
                        delete(v3, -9223372036854775808)
                    case "\n\t":
                        switch (fi23 | fi23) {
                        case 1:
                            switch ("" + "hello") {
                            case "":
                                if true {
                                    fmt.Println(v22, v22, v25)
                                    var v34 uint16 = 32767
                                    v35 := GenS3{F0: 127, F1: 2147483647}
                                    _ = v34
                                    _ = v35
                                }
                            }
                            v36 := v3[(fi23 & fi24)]
                            var v37 int = fi24
                            _ = v36
                            _ = v37
                        case 505:
                            v32 = append(v32, (fi24 * (fi23 - int(v18))))
                        case 4294967295:
                            v2 = (((-2.25 / ((v19 - v19) * 3.14159)) * v19) != -2.25)
                        }
                    }
                case "\n\t":
                    if len(v32) >= 2 {
                        v32 = v32[1:len(v32)]
                    }
                default:
                    v38 := GenS3{F0: 251, F1: v22}
                    _ = v38
                }
                switch 32768 {
                case -1:
                    switch -2147483648 {
                    case 4294967295:
                        if false {
                            v39 := GenS2{F0: "", F1: 1}
                            fmt.Println(v22)
                            v40 := GenS4{F0: 166}
                            _ = v39
                            _ = v40
                        }
                    case 418:
                        var v41 int16 = 1000
                        _ = v41
                    case -439:
                        v1 = (!(!false))
                    default:
                        v42 := GenS3{F0: 255, F1: 42}
                        _ = v42
                    }
                case 1:
                    v21 += (v21 << ((330) & 63))
                case 0:
                    if true {
                        v3[42] = (v22 + 1)
                        var v43 string = "\n\t"
                        _ = v43
                    }
                }
                _ = v32
            }
            _ = fi24
            _ = v25
        }
        v16 -= v16
        _ = fi23
    }
    v44 := GenS3{F0: 127, F1: (v22 + (v22 << ((4294967295) & 63)))}
    v45 := []int{-1, int(v18), 42, 0}
    var v46 uint8 = 127
    fmt.Println(v21)
    v47 := GenS5{F0: -2147483648, F1: (v18 * -128), F2: 270}
    var v48 int = -796
    if (v2 && true) {
        fmt.Println(v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("v46:", v46)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v48
}
