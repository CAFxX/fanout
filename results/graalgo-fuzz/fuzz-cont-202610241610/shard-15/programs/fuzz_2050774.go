package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 uint; F1 int64; F2 int16 }
type GenS4 struct { F0 string; F1 bool; F2 uint8 }
type GenS5 struct { F0 int; F1 uint8; F2 uint64 }
type GenS6 struct { F0 bool }
func genFn0(p0 int32) string {
    switch 32767 {
    case 413:
        for fi6 := 0; fi6 < 18; fi6++ {
            for fi7 := 0; fi7 < 4; fi7++ {
                v8 := []int{fi6, fi7, 0}
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                _ = fi7
                _ = v8
            }
            v9 := []int{(fi6 - 4294967295), (^(fi6)), fi6}
            var v10 uint32 = 2147483648
            _ = fi6
            _ = v10
            _ = v9
        }
    case -9223372036854775808:
        if (-1 <= 1) {
            switch 9223372036854775807 {
            case -9223372036854775808:
                if true {
                    if (!false) {
                        fmt.Println(p0)
                        var v11 float64 = (-2.25 + 0.1)
                        v12 := GenS1{F0: 0, F1: p0}
                        _ = v11
                        _ = v12
                    } else {
                        if (435 != 4294967295) {
                            v13 := GenS2{F0: 0, F1: 1}
                            _ = v13
                        }
                    }
                    var v14 int32 = p0
                    _ = v14
                } else {
                    var v15 uint64 = 18446744073709551615
                    fmt.Println(p0, v15)
                    _ = v15
                }
            default:
                switch ("\n\t" + "") {
                case "":
                    var v16 uint8 = 186
                    _ = v16
                default:
                    var v17 int8 = 121
                    _ = v17
                }
            }
            if true {
                var v18 uint = 42
                _ = v18
            } else {
                p0 -= p0
            }
        }
    }
    switch 9223372036854775807 {
    case 4294967295:
        v19 := GenS2{F0: 127, F1: 65}
        _ = v19
    default:
        v20 := GenS1{F0: 4294967295, F1: p0}
        _ = v20
    }
    var v21 string = "foo bar"
    v22 := []int{-176, -1, 106, 32767}
    switch -2147483648 {
    case 1:
        switch "hello" {
        case "":
            var v23 bool = false
            _ = v23
        }
    case 4294967295:
        var v24 uint = 255
        _ = v24
    default:
        var v25 int8 = -1
        _ = v25
    }
    _ = p0
    _ = v21
    _ = v22
    return v21
}
func genFn1(p0 uint, p1 int16) int8 {
    var v26 uint8 = 96
    if false {
        p1 = p1
        for fi27 := 0; fi27 < 20; fi27++ {
            var v28 int = fi27
            v29 := []int{v28}
            _ = fi27
            _ = v28
            _ = v29
        }
        v30 := []int{-74, -1, 2147483647, 32767}
        _ = v30
    }
    _ = p0
    _ = p1
    _ = v26
    return 100
}

func main() {
    var v1 uint32 = 42
    var v2 uint16 = 333
    v3 := []uint32{v1}
    var v4 float64 = 3.14159
    v5 := make(map[string]bool)
    v1 = v1
    var v31 uint32 = (^(v1))
    var v32 float32 = 1.5
    v33 := v5["a"]
    fmt.Println(v1)
    v2 = (^(v2))
    fmt.Println(v4, v2)
    switch "foo bar" {
    case "hello":
        for fi34 := 0; fi34 < 3; fi34++ {
            var v35 bool = v33
            for fi36 := 0; fi36 < 9; fi36++ {
                if v33 {
                    fmt.Println(fi36, v4)
                    var v37 int64 = -9223372036854775808
                    if (v32 > ((v32 / 3.14159) - 1.5)) {
                        v4 -= v4
                    }
                    _ = v37
                } else {
                    v38 := GenS3{F0: 42, F1: 692, F2: 629}
                    v4 = (3.14159 + v4)
                    _ = v38
                }
                _ = fi36
            }
            v39 := GenS4{F0: "", F1: (((-128 != 100) || (0 <= 608)) && false), F2: 150}
            _ = fi34
            _ = v35
            _ = v39
        }
    }
    fmt.Println(v4)
    v5["a"] = (-90 >= 0)
    if (false && (42 > 4294967295)) {
        var v40 int64 = -1
        v33 = v33
        _ = v40
    } else {
        var v41 bool = false
        _ = v41
    }
    if (975 >= 4294967295) {
        var v42 uint16 = (v2 & v2)
        _ = len(v5)
        _ = v42
    } else {
        switch "\n\t" {
        case "a":
            _ = len(v3)
        case "hello":
            for fi43 := 0; fi43 < 16; fi43++ {
                if v33 {
                    for fi44 := 0; fi44 < 11; fi44++ {
                        v1 = 4294967295
                        _ = fi44
                    }
                    switch "hello" {
                    case "a":
                        v45 := GenS5{F0: (fi43 * 1), F1: 1, F2: 627}
                        _ = v45
                    default:
                        var v46 int64 = 951
                        _ = v46
                    }
                }
                v4 = (((v4 * 1.5) - v4) - float64((v32 * float32((v4 / v4)))))
                fmt.Println(v31)
                _ = fi43
            }
        case "foo bar":
            v47 := GenS3{F0: 800, F1: 896, F2: -32768}
            _ = v47
        default:
            if (!v33) {
                v48 := GenS6{F0: (524 < 4294967295)}
                var v49 float32 = ((v32 / (0.1 + v32)) - 0.1)
                for fi50 := 0; fi50 < 16; fi50++ {
                    if (true || false) {
                        switch "" {
                        case "hello":
                            v51 := GenS1{F0: v1, F1: -628}
                            _ = v51
                        }
                        if (!("hello" == ("a" + "hello"))) {
                            switch "foo bar" {
                            case "\n\t":
                                var v52 int = (fi50 >> ((195) & 63))
                                _ = v52
                            case "foo bar":
                                fmt.Println(v31)
                            case "":
                                for fi53 := 0; fi53 < 17; fi53++ {
                                    v49 = ((v32 / float32(v4)) + v49)
                                    v4 += v4
                                    switch fi50 {
                                    case -9223372036854775808:
                                        v49 -= ((float32(float64(v49)) - 0.1) - 1.5)
                                    }
                                    _ = fi53
                                }
                            }
                        } else {
                            v4 = (v4 * v4)
                            v4 -= (v4 / 1.5)
                        }
                    }
                    _ = fi50
                }
                _ = v48
                _ = v49
            } else {
                var v54 uint16 = v2
                _ = v54
            }
        }
    }
    v4 = v4
    v32 = float32(1.5)
    var v55 uint = 0
    if (128 != 4294967295) {
        fmt.Println(v2)
        v4 = v4
        for fi56 := 0; fi56 < 15; fi56++ {
            v57 := GenS6{F0: false}
            var v58 uint16 = (v2 * 1)
            var v59 int32 = 42
            _ = fi56
            _ = v57
            _ = v58
            _ = v59
        }
    } else {
        v55 += 65535
        v33 = v33
        var v60 float64 = (v4 * 1e10)
        _ = v60
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v55
}
