package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 string; F1 uint64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int16; F1 uint32; F2 string }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint64; F1 string }
func genFn0(p0 uint64, p1 uint8) int32 {
    p1 = p1
    var v5 uint8 = (p1 >> ((0) & 63))
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 uint = 463
        v7 = 0
        var v8 string = ("" + "foo bar")
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v5
    return 0
}
func genFn1(p0 int16, p1 int32, p2 bool) int64 {
    v9 := []int{386}
    v10 := GenS1{F0: ("" + ("" + (("a" + "foo bar") + ("\n\t" + "\n\t")))), F1: true}
    v9 = append(v9, 42)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 1
}

func main() {
    var v1 int64 = 42
    v2 := []int{}
    var v3 int16 = 42
    v4 := make(map[string]bool)
    v11 := GenS1{F0: "hello", F1: false}
    for fi12 := 0; fi12 < 17; fi12++ {
        if (-1 != 0) {
            var v13 int = (fi12 - fi12)
            _ = v13
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v14 float64 = (3.14159 + 0.0)
            _ = v14
        }
        _ = fi12
    }
    delete(v4, ("\n\t" + "\n\t"))
    for fi15 := 0; fi15 < 12; fi15++ {
        v1 = v1
        if (!true) {
            v1 = (^(v1))
            v16 := GenS2{F0: 0, F1: v1}
            if (!false) {
                var v17 uint16 = 11
                var v18 int64 = (v1 / (((v1 % ((-529) | 1))) | 1))
                if true {
                    for fi19 := 0; fi19 < 11; fi19++ {
                        v1 = v1
                        _ = fi19
                    }
                }
                _ = v17
                _ = v18
            }
            _ = v16
        } else {
            for fi20 := 0; fi20 < 12; fi20++ {
                v21 := GenS2{F0: 255, F1: -297}
                if false {
                    v22 := GenS3{F0: "a", F1: 346}
                    fmt.Println(v1, fi15)
                    _ = v22
                } else {
                    switch (fi15 / (((fi20 - fi15)) | 1)) {
                    case 723:
                        v2 = append(v2, 32768)
                    case 32768:
                        _ = len(v2)
                    }
                }
                if false {
                    var v23 uint64 = 4294967295
                    _ = v23
                } else {
                    v4["hello"] = false
                    v3 = v3
                    if true {
                        var v24 uint16 = 0
                        v25 := GenS1{F0: (("a" + "") + (("foo bar" + "hello") + "foo bar")), F1: (!(!((!true) || false)))}
                        switch -267 {
                        case 32767:
                            if false {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            } else {
                                var v26 float32 = 3.14159
                                if (!(!false)) {
                                    var v27 bool = (!true)
                                    v27 = v27
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    if v27 {
                                        var v28 bool = true
                                        var v29 int = 4294967295
                                        _ = v28
                                        _ = v29
                                    } else {
                                        var v30 float32 = 0.1
                                        v27 = (42 >= 0)
                                        v27 = (1 >= 0)
                                        _ = v30
                                    }
                                    _ = v27
                                }
                                var v31 string = "\n\t"
                                _ = v26
                                _ = v31
                            }
                        case 128:
                            if (false || false) {
                                v1 -= (-(v1))
                                var v32 float64 = 3.14159
                                _ = v32
                            } else {
                                if (false && ("\n\t" >= "foo bar")) {
                                    delete(v4, ("foo bar" + "\n\t"))
                                    v3 += 32767
                                    var v33 float64 = -2.25
                                    _ = v33
                                }
                                var v34 float32 = float32((1.5 + 0.0))
                                v35 := GenS3{F0: "a", F1: 42}
                                v36 := GenS2{F0: 0, F1: (v1 << ((uint(v24)) & 63))}
                                if false {
                                    v2 = append(v2, fi15)
                                    var v37 int64 = (v1 | 0)
                                    if false {
                                        v1 -= v1
                                        var v38 int32 = -2147483648
                                        _ = v38
                                    } else {
                                        v39 := GenS4{F0: 331}
                                        v3 -= (^(v3))
                                        _ = v39
                                    }
                                    _ = v37
                                }
                                v24 += v24
                                switch "a" {
                                case "a":
                                    switch fi20 {
                                    case -9223372036854775808:
                                        if true {
                                            var v40 int8 = 83
                                            var v41 float32 = ((v34 * v34) * float32((1.5 + (((1.5 + (float64(v34) * 0.1)) / -2.25) / 1.5))))
                                            delete(v4, "\n\t")
                                            _ = v40
                                            _ = v41
                                        }
                                    default:
                                        fmt.Println(fi20, v1)
                                    }
                                case "":
                                    if (!(0.1 == 0.1)) {
                                        switch (("" + "foo bar") + "") {
                                        case "a":
                                            v24 = (v24 | v24)
                                        }
                                        v42 := GenS5{F0: 42}
                                        _ = v42
                                    }
                                case "\n\t":
                                    v43 := GenS6{F0: "a"}
                                    _ = v43
                                }
                                _ = v34
                                _ = v35
                                _ = v36
                            }
                        case 0:
                            v1 = genFn1(v3, 1, false)
                        default:
                            if true {
                                v2 = append(v2, (fi20 % ((-1) | 1)))
                                v1 += (^(v1))
                            }
                            v2 = append(v2, (fi20 | -2147483648))
                        }
                        _ = v24
                        _ = v25
                    }
                }
                _ = fi20
                _ = v21
            }
            v44 := GenS4{F0: 255}
            _ = v44
        }
        v45 := GenS7{F0: (v3 - (^(v3))), F1: 4294967295, F2: (("foo bar" + "foo bar") + ("foo bar" + ((("\n\t" + "") + "\n\t") + "foo bar")))}
        _ = fi15
        _ = v45
    }
    v46 := GenS8{F0: -426}
    if (856 >= 42) {
        if (((v3 - (v3 | v3)) < 697) || true) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v47 uint64 = 828
            var v48 float64 = (1e10 - (3.14159 / 1e10))
            _ = v47
            _ = v48
        } else {
            v1 += v1
            v49 := GenS9{F0: 1, F1: ""}
            if (true || true) {
                for fi50 := 0; fi50 < 3; fi50++ {
                    for fi51 := 0; fi51 < 6; fi51++ {
                        var v52 string = "hello"
                        var v53 uint64 = 2
                        _ = fi51
                        _ = v52
                        _ = v53
                    }
                    _ = fi50
                }
                fmt.Println(v1)
            }
            _ = v49
        }
    } else {
        var v54 int64 = 0
        v3 += (v3 >> ((65535) & 63))
        var v55 int32 = genFn0(42, 2)
        _ = v54
        _ = v55
    }
    fmt.Println(v3)
    if len(v2) > 0 {
        v2[0] = -909
    }
    var v56 uint16 = 32767
    v56 -= v56
    for fi57 := 0; fi57 < 15; fi57++ {
        var v58 int8 = 127
        _ = fi57
        _ = v58
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v56:", v56)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v56
}
