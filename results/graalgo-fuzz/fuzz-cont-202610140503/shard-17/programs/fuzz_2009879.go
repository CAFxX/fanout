package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 uint; F2 uint16 }
type GenS3 struct { F0 uint32; F1 string }
type GenS4 struct { F0 uint32; F1 int32; F2 int32 }
type GenS5 struct { F0 uint64 }
func genFn0(p0 string, p1 uint16, p2 int32) int64 {
    p2 += p2
    var v7 uint32 = 65535
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}

func main() {
    v1 := []string{("" + "foo bar"), "hello", "a"}
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]bool)
    var v4 float32 = (((3.14159 / 1.0) / 1.5) - 3.14159)
    var v5 uint8 = 255
    v6 := make(map[int]uint32)
    _ = len(v2)
    v4 -= (v4 + v4)
    if true {
        v1 = append(v1, ("foo bar" + ("\n\t" + "foo bar")))
        v4 += v4
        if (!(v5 == v5)) {
            v4 += (v4 - ((float32(1e10) + (3.14159 - v4)) * 1e10))
        } else {
            var v8 int8 = 86
            var v9 int64 = 0
            v9 = (^(v9))
            _ = v8
            _ = v9
        }
    }
    v10 := GenS1{F0: 857}
    v4 -= v4
    switch 300 {
    case 799:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case 0:
        switch ("hello" + "foo bar") {
        case "\n\t":
            var v11 uint8 = 46
            _ = v11
        default:
            v12 := GenS2{F0: 190, F1: 65535, F2: 42}
            _ = v12
        }
    case 128:
        if false {
            for fi13 := 0; fi13 < 6; fi13++ {
                for fi14 := 0; fi14 < 18; fi14++ {
                    switch (("" + "\n\t") + "\n\t") {
                    case "\n\t":
                        if (true || true) {
                            v5 -= (^(v5))
                        } else {
                            switch "a" {
                            case "hello":
                                fmt.Println(v4, fi14)
                            case "a":
                                fmt.Println(fi14)
                            default:
                                v15 := GenS3{F0: 0, F1: "foo bar"}
                                _ = v15
                            }
                        }
                    case "foo bar":
                        v5 = (v5 / ((42) | 1))
                        v16 := GenS3{F0: 42, F1: "a"}
                        v4 = (v4 - ((v4 - float32(float64((0.1 - v4)))) * v4))
                        _ = v16
                    case "hello":
                        if false {
                            v4 = (0.0 - ((v4 * v4) - v4))
                            v5 += (v5 & (v5 >> ((48) & 63)))
                        } else {
                            v2[65535] = (fi14 + 0)
                        }
                    }
                    v5 = (v5 << ((2) & 63))
                    v4 += ((3.14159 * v4) - 1e10)
                    _ = fi14
                }
                switch "foo bar" {
                case "a":
                    for fi17 := 0; fi17 < 17; fi17++ {
                        var v18 bool = false
                        v19 := GenS4{F0: 2, F1: 0, F2: -716}
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                }
                _ = fi13
            }
        }
    }
    for fi20 := 0; fi20 < 13; fi20++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if false {
            var v21 int32 = -1
            switch "\n\t" {
            case "\n\t":
                for fi22 := 0; fi22 < 17; fi22++ {
                    v4 += v4
                    var v23 uint32 = 2147483648
                    var v24 string = "\n\t"
                    _ = fi22
                    _ = v23
                    _ = v24
                }
            case "hello":
                var v25 uint16 = 2
                _ = v25
            }
            v26 := GenS5{F0: 0}
            _ = v21
            _ = v26
        }
        _ = fi20
    }
    for fi27 := 0; fi27 < 6; fi27++ {
        switch int(v5) {
        case -9223372036854775808:
            var v28 int64 = 814
            _ = v28
        case -2147483648:
            v5 = 61
        case -121:
            v5 -= (v5 & (v5 + 108))
        default:
            if (false || true) {
                delete(v6, -821)
                v4 -= ((v4 + ((0.0 / (v4 * (v4 - float32(0.0)))) * 0.0)) + v4)
            } else {
                if ((fi27 < 631) && true) {
                    v4 = (v4 - 0.1)
                } else {
                    _ = len(v1)
                    var v29 int64 = -665
                    for fi30 := 0; fi30 < 16; fi30++ {
                        var v31 int16 = 0
                        var v32 int8 = -128
                        _ = fi30
                        _ = v31
                        _ = v32
                    }
                    _ = v29
                }
                v33 := v2[2]
                _ = v33
            }
        }
        _ = fi27
    }
    fmt.Println(v5)
    var v34 float32 = v4
    var v35 bool = true
    fmt.Println(v4)
    var v36 string = "foo bar"
    var v37 string = ("" + v36)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
