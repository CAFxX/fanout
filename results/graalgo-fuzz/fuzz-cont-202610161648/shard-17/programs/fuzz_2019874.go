package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 uint; F2 uint16 }
type GenS3 struct { F0 int64; F1 uint32; F2 int32 }
type GenS4 struct { F0 uint16; F1 int32 }
func genFn0(p0 int64, p1 uint, p2 uint8) int {
    v5 := GenS1{F0: (p1 - uint(p2))}
    v6 := GenS1{F0: p1}
    v7 := []int{}
    p1 += (p1 | (^(p1)))
    var v8 int32 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 42
}
func genFn1() uint8 {
    v9 := GenS1{F0: 42}
    v10 := []int{128, 4294967295, -1}
    var v11 int16 = 0
    switch "hello" {
    case "a":
        switch genFn0(9223372036854775807, 65535, 128) {
        case 2147483647:
            if true {
                v12 := GenS1{F0: 1}
                _ = v12
            }
        case -514:
            v11 = v11
        }
    default:
        for fi13 := 0; fi13 < 9; fi13++ {
            fmt.Println(v11, fi13)
            _ = fi13
        }
    }
    _ = v10
    _ = v11
    _ = v9
    return 0
}

func main() {
    var v1 int8 = 100
    v2 := []int{}
    v3 := make(map[string]bool)
    v4 := []string{(((("" + "\n\t") + "") + "hello") + "\n\t"), ("foo bar" + ("" + "foo bar")), ("\n\t" + "a")}
    if (!true) {
        v14 := GenS1{F0: 916}
        v15 := GenS2{F0: 582, F1: 0, F2: 42}
        v1 += (v1 | (v1 + v1))
        _ = v14
        _ = v15
    } else {
        v16 := GenS3{F0: 884, F1: 831, F2: -1}
        var v17 uint64 = 158
        v1 -= v1
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 18; fi18++ {
        v19 := v3["foo bar"]
        var v20 int = fi18
        switch "hello" {
        case "":
            if ((-(v1)) <= v1) {
                var v21 uint64 = 42
                fmt.Println(v19)
                var v22 int32 = 1
                _ = v21
                _ = v22
            } else {
                if len(v4) > 0 {
                    v4[0] = ("a" + "\n\t")
                }
                v19 = (4294967295 > 2)
                if (384 == 32767) {
                    var v23 uint16 = 977
                    _ = v23
                } else {
                    var v24 int16 = 265
                    if false {
                        if (!(v1 >= -1)) {
                            var v25 uint32 = 2147483648
                            v26 := GenS3{F0: 0, F1: 104, F2: 592}
                            for fi27 := 0; fi27 < 18; fi27++ {
                                var v28 string = (("hello" + "a") + ("a" + "foo bar"))
                                _ = fi27
                                _ = v28
                            }
                            _ = v25
                            _ = v26
                        }
                        fmt.Println(v1, fi18)
                    }
                    _ = len(v2)
                    _ = v24
                }
            }
        case "hello":
            if len(v2) > 0 {
                v2[0] = fi18
            }
        }
        _ = fi18
        _ = v19
        _ = v20
    }
    if true {
        var v29 uint8 = 42
        _ = v29
    }
    for fi30 := 0; fi30 < 2; fi30++ {
        for fi31 := 0; fi31 < 4; fi31++ {
            v3[("\n\t" + ("a" + ("\n\t" + ("foo bar" + "hello"))))] = false
            v32 := GenS2{F0: 0, F1: 65535, F2: 510}
            _ = fi31
            _ = v32
        }
        var v33 uint8 = 42
        _ = fi30
        _ = v33
    }
    switch "hello" {
    case "":
        var v34 int8 = v1
        _ = v34
    case "\n\t":
        _ = len(v2)
    case "hello":
        v35 := GenS3{F0: 9223372036854775807, F1: 167, F2: -1}
        _ = v35
    default:
        if (-2147483648 >= 29) {
            v1 = 42
        } else {
            switch 1 {
            case 127:
                if true {
                    var v36 int64 = 247
                    _ = v36
                } else {
                    fmt.Println(v1)
                    _ = len(v4)
                    var v37 bool = true
                    _ = v37
                }
            default:
                v1 = v1
            }
            v1 += (v1 ^ (v1 / (((v1 / (((v1 * v1)) | 1))) | 1)))
        }
    }
    switch -1 {
    case 631:
        _ = len(v4)
    case 4294967295:
        v1 += 42
    case 32767:
        v1 = (-(v1))
    default:
        if false {
            if true {
                v1 += v1
                delete(v3, "a")
            } else {
                var v38 float32 = 3.14159
                fmt.Println(v1)
                if false {
                    var v39 int64 = -1
                    if ((!true) && ("hello" == "hello")) {
                        var v40 string = "hello"
                        var v41 uint16 = 954
                        v39 = v39
                        _ = v40
                        _ = v41
                    }
                    _ = v39
                } else {
                    var v42 int64 = 9223372036854775807
                    var v43 int64 = 42
                    var v44 bool = (float64(((v38 / 3.14159) - v38)) != -2.25)
                    _ = v42
                    _ = v43
                    _ = v44
                }
                _ = v38
            }
            var v45 int = 35
            for fi46 := 0; fi46 < 2; fi46++ {
                if false {
                    var v47 int32 = 42
                    var v48 uint8 = 255
                    var v49 string = "hello"
                    _ = v47
                    _ = v48
                    _ = v49
                } else {
                    v1 = (v1 << ((65535) & 63))
                    v50 := GenS4{F0: 32767, F1: -2147483648}
                    _ = v50
                }
                var v51 int8 = (v1 ^ (-(v1)))
                fmt.Println(fi46, v1, v1)
                _ = fi46
                _ = v51
            }
            _ = v45
        } else {
            if len(v4) > 0 {
                v4[0] = ("foo bar" + ("foo bar" + "hello"))
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            fmt.Println(v1)
        }
    }
    for fi52 := 0; fi52 < 6; fi52++ {
        var v53 uint32 = 2147483648
        for fi54 := 0; fi54 < 18; fi54++ {
            var v55 int64 = 0
            _ = fi54
            _ = v55
        }
        var v56 uint8 = 128
        _ = fi52
        _ = v53
        _ = v56
    }
    var v57 float32 = (3.14159 * 1e10)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v57
}
