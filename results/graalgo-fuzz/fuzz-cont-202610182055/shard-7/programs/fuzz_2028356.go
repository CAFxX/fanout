package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 uint64 }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint64; F1 int; F2 uint64 }
type GenS7 struct { F0 int32; F1 uint8 }
func genFn0(p0 bool, p1 int32) int8 {
    p1 = 42
    v5 := GenS1{F0: 0, F1: 0, F2: 454}
    for fi6 := 0; fi6 < 9; fi6++ {
        var v7 uint = 0
        var v8 bool = true
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 uint8 = 1
    v9 = (v9 | v9)
    _ = p0
    _ = p1
    _ = v5
    _ = v9
    return 0
}
func genFn1(p0 int16, p1 int32) uint16 {
    v10 := []int{2147483647}
    var v11 float32 = -2.25
    v11 -= v11
    v10 = append(v10, 0)
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 32767
}

func main() {
    var v1 string = "hello"
    var v2 string = "\n\t"
    v3 := make(map[string]bool)
    var v4 string = v2
    v12 := GenS1{F0: -38, F1: 65535, F2: 2}
    var v13 uint32 = 845
    switch v1 {
    case "foo bar":
        switch "\n\t" {
        case "":
            v4 = (v1 + v4)
        case "\n\t":
            v13 -= (v13 % ((v13) | 1))
        case "a":
            v2 = v4
        }
    }
    var v14 int = 0
    var v15 string = ("foo bar" + "foo bar")
    if (true || (false && false)) {
        v15 = v4
        v16 := GenS2{F0: v14, F1: ((!(3.14159 > -2.25)) || ((3.14159 + 0.1) < (-2.25 * ((float32(1.5) - -2.25) / float32((1e10 + 1e10))))))}
        var v17 bool = true
        _ = v16
        _ = v17
    } else {
        v4 = v2
    }
    v18 := []int{(v14 & 2147483647), v14}
    var v19 bool = (313 < 32767)
    for fi20 := 0; fi20 < 2; fi20++ {
        if v19 {
            v13 = v13
        }
        _ = fi20
    }
    fmt.Println(v1, v1)
    fmt.Println(v15, v15)
    switch v4 {
    case "foo bar":
        fmt.Println(v15)
    case "a":
        var v21 string = (v2 + ((v1 + (v15 + (v4 + ""))) + v2))
        _ = v21
    }
    v22 := GenS1{F0: -1, F1: 1, F2: 2}
    for fi23 := 0; fi23 < 14; fi23++ {
        v24 := GenS2{F0: -2147483648, F1: false}
        if true {
            switch v14 {
            case 32767:
                switch 4294967295 {
                case -661:
                    v25 := GenS3{F0: (^(v13))}
                    _ = v25
                case -980:
                    var v26 bool = v19
                    _ = v26
                case 128:
                    var v27 int32 = 0
                    _ = v27
                default:
                    for fi28 := 0; fi28 < 19; fi28++ {
                        if (((((v2 + v1) + "a") + ("hello" + v2)) + "a") < v4) {
                            if len(v18) >= 2 {
                                v18 = v18[1:len(v18)]
                            }
                        } else {
                            var v29 uint64 = 836
                            var v30 bool = false
                            if false {
                                v31 := GenS4{F0: 2}
                                _ = v31
                            }
                            _ = v29
                            _ = v30
                        }
                        v4 = v15
                        switch v1 {
                        case "":
                            var v32 int = (^(fi28))
                            _ = v32
                        case "foo bar":
                            if v19 {
                                delete(v3, v4)
                                var v33 uint64 = 11
                                switch "a" {
                                case "":
                                    v13 = v13
                                default:
                                    v34 := GenS5{F0: 867}
                                    _ = v34
                                }
                                _ = v33
                            } else {
                                var v35 string = v2
                                var v36 uint64 = 241
                                if (0 > -1) {
                                    v37 := GenS6{F0: v36, F1: (-(fi28)), F2: 0}
                                    _ = v37
                                }
                                v38 := GenS3{F0: v13}
                                _ = v35
                                _ = v36
                                _ = v38
                            }
                        case "hello":
                            var v39 int16 = 0
                            _ = v39
                        }
                        _ = fi28
                    }
                }
            case -124:
                for fi40 := 0; fi40 < 16; fi40++ {
                    var v41 int16 = 169
                    _ = fi40
                    _ = v41
                }
            case 990:
                v13 -= 0
            }
        }
        var v42 uint64 = 18446744073709551615
        _ = fi23
        _ = v24
        _ = v42
    }
    v43 := GenS7{F0: 42, F1: 1}
    var v44 uint8 = 63
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v43
    _ = v44
}
