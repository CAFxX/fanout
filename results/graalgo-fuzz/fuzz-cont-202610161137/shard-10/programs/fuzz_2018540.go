package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 int16; F2 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint16; F1 uint16 }
type GenS6 struct { F0 uint8; F1 string }
type GenS7 struct { F0 bool; F1 uint }
func genFn0(p0 uint64) uint {
    var v4 uint64 = (p0 - 0)
    v5 := []int{127, -9223372036854775808, 128}
    if (431 < 814) {
        var v6 bool = false
        _ = v6
    }
    _ = len(v5)
    _ = p0
    _ = v4
    _ = v5
    return 730
}
func genFn1(p0 uint, p1 bool, p2 uint8) uint32 {
    var v7 uint32 = 679
    v8 := GenS1{F0: ("" + "hello")}
    v9 := GenS1{F0: "hello"}
    var v10 uint16 = 779
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return v7
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]uint32)
    v3 := make(map[string]bool)
    for fi11 := 0; fi11 < 5; fi11++ {
        fmt.Println(fi11)
        v12 := GenS1{F0: ""}
        _ = fi11
        _ = v12
    }
    var v13 int16 = 32767
    if false {
        switch 128 {
        case 790:
            fmt.Println(v13)
        default:
            var v14 int16 = (^(v13))
            _ = v14
        }
    } else {
        _ = len(v1)
    }
    if false {
        v15 := GenS1{F0: "\n\t"}
        _ = v15
    } else {
        switch 127 {
        case -9223372036854775808:
            var v16 uint = 1
            _ = v16
        default:
            v17 := GenS2{F0: 127, F1: v13, F2: -91}
            _ = v17
        }
        v18 := GenS3{F0: 255}
        _ = v18
    }
    v19 := GenS4{F0: 0}
    if true {
        for fi20 := 0; fi20 < 15; fi20++ {
            v21 := GenS4{F0: 91}
            for fi22 := 0; fi22 < 19; fi22++ {
                var v23 int64 = 352
                if true {
                    v23 += v23
                    v23 = (v23 % (((v23 + (^(v23)))) | 1))
                    _ = len(v2)
                    v24 := GenS5{F0: 65535, F1: 65535}
                    _ = v24
                }
                _ = fi22
                _ = v23
            }
            _ = fi20
            _ = v21
        }
    }
    v13 += v13
    v13 += (v13 & -827)
    for fi25 := 0; fi25 < 12; fi25++ {
        switch (fi25 >> ((2) & 63)) {
        case 0:
            for fi26 := 0; fi26 < 19; fi26++ {
                v13 -= 0
                if (0 != 368) {
                    var v27 string = "hello"
                    _ = v27
                }
                var v28 int8 = 47
                _ = fi26
                _ = v28
            }
        case 605:
            if true {
                var v29 int = (^(fi25))
                if false {
                    if false {
                        fmt.Println(v29, fi25, v13)
                        var v30 int16 = v13
                        v31 := GenS6{F0: 2, F1: "\n\t"}
                        _ = v30
                        _ = v31
                    }
                } else {
                    var v32 int8 = 9
                    _ = v32
                }
                _ = v29
            }
        default:
            v2[32767] = 2147483647
        }
        _ = fi25
    }
    v33 := []int{42, -9223372036854775808, 4294967295, -1}
    v34 := GenS2{F0: 100, F1: 1, F2: 1}
    for fi35 := 0; fi35 < 16; fi35++ {
        if (!true) {
            switch "foo bar" {
            case "a":
                switch ("\n\t" + "foo bar") {
                case "hello":
                    var v36 string = "\n\t"
                    _ = v36
                case "foo bar":
                    v1[fi35] = "hello"
                }
            case "foo bar":
                v13 -= v13
            case "hello":
                for fi37 := 0; fi37 < 17; fi37++ {
                    var v38 float32 = float32(float64((1e10 / 1e10)))
                    v13 += (v13 * (v13 % ((v13) | 1)))
                    v39 := GenS7{F0: (!(!false)), F1: 1}
                    _ = fi37
                    _ = v38
                    _ = v39
                }
            default:
                v13 = v13
            }
            v40 := GenS3{F0: 65535}
            var v41 float64 = 0.0
            _ = v40
            _ = v41
        } else {
            var v42 string = ("\n\t" + ("\n\t" + (("hello" + "a") + ("" + "foo bar"))))
            v42 = v42
            for fi43 := 0; fi43 < 5; fi43++ {
                var v44 float32 = 1.5
                _ = fi43
                _ = v44
            }
            _ = v42
        }
        v33 = append(v33, fi35)
        _ = fi35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v33:", len(v33), v33)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v33
    _ = v34
}
