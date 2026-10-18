package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 int32; F2 int8 }
type GenS3 struct { F0 string; F1 uint64 }
type GenS4 struct { F0 uint8 }
func genFn0(p0 int32, p1 uint8, p2 int32) bool {
    if true {
        fmt.Println(p2, p0, p2)
        for fi7 := 0; fi7 < 1; fi7++ {
            for fi8 := 0; fi8 < 7; fi8++ {
                var v9 int32 = p0
                v10 := []int{fi7, int(p1)}
                _ = fi8
                _ = v10
                _ = v9
            }
            _ = fi7
        }
        var v11 float32 = 1.5
        _ = v11
    }
    fmt.Println(p1, p2, p1)
    for fi12 := 0; fi12 < 12; fi12++ {
        p1 += (p1 + (p1 >> ((4294967295) & 63)))
        var v13 int32 = (p0 >> ((65535) & 63))
        _ = fi12
        _ = v13
    }
    var v14 int64 = 9223372036854775807
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return true
}
func genFn1(p0 int, p1 uint, p2 uint32) int64 {
    var v15 uint8 = 2
    var v16 bool = (!true)
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    return -931
}

func main() {
    v1 := []uint32{42}
    v2 := make(map[int]string)
    var v3 int8 = 0
    var v4 uint64 = 0
    var v5 float64 = 3.14159
    v6 := make(map[string]bool)
    var v17 int16 = -842
    if (3.14159 < 0.1) {
        for fi18 := 0; fi18 < 17; fi18++ {
            v19 := GenS1{F0: -1}
            v5 = v5
            _ = fi18
            _ = v19
        }
        fmt.Println(v4, v4, v5)
        if false {
            fmt.Println(v17)
            v3 = (-(v3))
            switch ("\n\t" + "") {
            case "hello":
                v3 = (v3 | (v3 >> ((65535) & 63)))
            }
        }
    } else {
        for fi20 := 0; fi20 < 4; fi20++ {
            fmt.Println(v17)
            switch "" {
            case "\n\t":
                var v21 string = "hello"
                _ = v21
            case "hello":
                if false {
                    if true {
                        for fi22 := 0; fi22 < 20; fi22++ {
                            if len(v1) > 0 {
                                v1[0] = 485
                            }
                            v5 -= 1e10
                            v23 := GenS2{F0: 0, F1: -2147483648, F2: (v3 + 127)}
                            _ = fi22
                            _ = v23
                        }
                        var v24 uint8 = 0
                        v3 = v3
                        _ = v24
                    }
                }
            default:
                switch "" {
                case "":
                    if (!false) {
                        fmt.Println(fi20, v5, v17)
                    }
                case "\n\t":
                    var v25 bool = true
                    _ = v25
                case "foo bar":
                    for fi26 := 0; fi26 < 6; fi26++ {
                        v27 := GenS2{F0: 118, F1: 42, F2: v3}
                        delete(v6, ("hello" + ((("\n\t" + "foo bar") + "hello") + "hello")))
                        if ((v4 / ((v4) | 1)) != 0) {
                            v28 := GenS1{F0: 1}
                            _ = v28
                        }
                        _ = fi26
                        _ = v27
                    }
                }
            }
            v17 -= -592
            _ = fi20
        }
        switch 9223372036854775807 {
        case 1:
            _ = len(v1)
        case 2147483647:
            var v29 int = -722
            _ = v29
        case -680:
            fmt.Println(v17)
        }
    }
    var v30 uint16 = 2
    for fi31 := 0; fi31 < 10; fi31++ {
        v3 = (v3 + v3)
        if false {
            var v32 int8 = v3
            v30 = (v30 | (v30 * (v30 % ((v30) | 1))))
            var v33 uint = 42
            _ = v32
            _ = v33
        } else {
            var v34 float64 = v5
            v35 := GenS3{F0: "", F1: 237}
            if true {
                v5 = float64(float32(1e10))
            } else {
                switch "foo bar" {
                case "\n\t":
                    switch fi31 {
                    case 4294967295:
                        var v36 uint64 = 498
                        _ = v36
                    case 269:
                        v1 = append(v1, 2147483648)
                    }
                default:
                    var v37 int = -2147483648
                    _ = v37
                }
                var v38 int64 = 944
                _ = v38
            }
            _ = v34
            _ = v35
        }
        fmt.Println(fi31, v3, fi31)
        _ = fi31
    }
    v39 := GenS1{F0: 1}
    v3 = v3
    var v40 bool = (v4 == v4)
    _ = len(v2)
    fmt.Println(v30, v17, v4)
    v30 = 65535
    for fi41 := 0; fi41 < 1; fi41++ {
        v42 := GenS4{F0: 127}
        _ = fi41
        _ = v42
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v30
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
