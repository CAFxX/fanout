package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 string }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 uint64; F1 int }
type GenS5 struct { F0 int8; F1 int64; F2 uint64 }
type GenS6 struct { F0 int32; F1 int8 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint32; F1 string; F2 int }
func genFn0(p0 uint) uint {
    p0 += (^(p0))
    var v6 int32 = -836
    for fi7 := 0; fi7 < 16; fi7++ {
        v8 := []int{(fi7 >> ((p0) & 63))}
        var v9 string = ("foo bar" + (("\n\t" + ((("a" + "\n\t") + "hello") + "hello")) + ""))
        _ = fi7
        _ = v8
        _ = v9
    }
    p0 = (p0 >> ((p0) & 63))
    _ = p0
    _ = v6
    return 42
}
func genFn1(p0 int8, p1 uint32) int8 {
    var v10 string = "hello"
    if true {
        v11 := GenS1{F0: (p1 % (((p1 | (^(p1)))) | 1)), F1: (p1 / (((p1 / ((1) | 1))) | 1)), F2: (v10 + v10)}
        _ = v11
    } else {
        if false {
            if (42 >= -2147483648) {
                v12 := GenS1{F0: 2147483647, F1: (p1 >> ((135) & 63)), F2: v10}
                p0 = 77
                p0 += p0
                _ = v12
            }
            v10 = v10
        }
    }
    switch int(p1) {
    case -1:
        v13 := GenS2{F0: 2, F1: 42}
        _ = v13
    case -9223372036854775808:
        var v14 int8 = -1
        _ = v14
    default:
        for fi15 := 0; fi15 < 20; fi15++ {
            var v16 bool = true
            _ = fi15
            _ = v16
        }
    }
    p0 = p0
    v17 := []int{9223372036854775807}
    _ = p0
    _ = p1
    _ = v10
    _ = v17
    return (^(p0))
}

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[string]int)
    v3 := make(map[string]string)
    var v4 uint8 = 240
    var v5 uint64 = 774
    var v18 int64 = 0
    for fi19 := 0; fi19 < 12; fi19++ {
        var v20 string = (("hello" + "hello") + "")
        var v21 float64 = v1
        for fi22 := 0; fi22 < 6; fi22++ {
            v23 := []int{}
            var v24 int32 = 986
            _ = fi22
            _ = v23
            _ = v24
        }
        _ = fi19
        _ = v20
        _ = v21
    }
    if ((^(v4)) <= v4) {
        for fi25 := 0; fi25 < 20; fi25++ {
            if ((0 != 2) && ((2 > 12) || false)) {
                for fi26 := 0; fi26 < 12; fi26++ {
                    if (true || true) {
                        if ((-952 != -410) && true) {
                            v1 += (0.0 * float64((-2.25 * (3.14159 * 1e10))))
                            switch "foo bar" {
                            case "foo bar":
                                if true {
                                    var v27 uint64 = uint64(v4)
                                    v18 = v18
                                    v28 := GenS2{F0: uint(v27), F1: 197}
                                    _ = v27
                                    _ = v28
                                }
                            case "a":
                                v29 := GenS3{F0: 9223372036854775807, F1: 1}
                                _ = v29
                            }
                            var v30 uint16 = 255
                            v31 := GenS2{F0: genFn0(0), F1: 0}
                            _ = v30
                            _ = v31
                        }
                        v32 := []int{(fi25 & (-(fi26))), fi26}
                        if false {
                            var v33 uint = 1
                            _ = v33
                        } else {
                            v1 = ((3.14159 + 0.0) + v1)
                            var v34 float32 = 1e10
                            v35 := GenS4{F0: (v5 | v5), F1: fi25}
                            _ = v34
                            _ = v35
                        }
                        _ = v32
                    }
                    v36 := GenS3{F0: fi26, F1: 42}
                    v18 = v18
                    _ = fi26
                    _ = v36
                }
            }
            v3["foo bar"] = "foo bar"
            if true {
                var v37 int16 = 32767
                v18 = (v18 / ((v18) | 1))
                _ = v37
            }
            _ = fi25
        }
    }
    v5 += v5
    v1 -= v1
    switch "hello" {
    case "":
        var v38 int = -251
        _ = v38
    }
    v39 := GenS5{F0: -128, F1: v18, F2: v5}
    v40 := GenS1{F0: 42, F1: 327, F2: ("hello" + ("" + ("a" + "a")))}
    v41 := GenS6{F0: 42, F1: 1}
    var v42 string = "foo bar"
    for fi43 := 0; fi43 < 11; fi43++ {
        switch fi43 {
        case -847:
            delete(v2, v42)
        case 128:
            v44 := GenS7{F0: 1}
            _ = v44
        case 4294967295:
            v45 := GenS8{F0: 0, F1: v42, F2: fi43}
            _ = v45
        }
        var v46 string = v42
        var v47 string = v46
        _ = fi43
        _ = v46
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
}
