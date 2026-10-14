package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int32; F1 int; F2 uint64 }
type GenS5 struct { F0 bool; F1 int; F2 int8 }
type GenS6 struct { F0 uint; F1 int8; F2 int16 }
type GenS7 struct { F0 uint16; F1 int32 }
func genFn0() int32 {
    v6 := GenS1{F0: 42}
    v7 := GenS2{F0: 32767, F1: 32767}
    _ = v6
    _ = v7
    return -728
}
func genFn1(p0 int64) uint {
    p0 -= 0
    p0 += (p0 | (-(p0)))
    if true {
        p0 = (p0 | 9223372036854775807)
        fmt.Println(p0)
        switch "" {
        case "a":
            fmt.Println(p0)
        case "\n\t":
            var v8 int16 = 938
            _ = v8
        }
    } else {
        for fi9 := 0; fi9 < 20; fi9++ {
            p0 = (^(p0))
            v10 := []int{fi9}
            _ = fi9
            _ = v10
        }
        p0 -= (p0 % (((^(p0))) | 1))
    }
    _ = p0
    return 0
}

func main() {
    var v1 uint64 = 42
    v2 := []uint32{298, 4294967295, 65535, 254}
    var v3 uint16 = 42
    v4 := make(map[uint32]bool)
    var v5 float32 = 0.0
    if true {
        if true {
            v11 := GenS3{F0: -757}
            _ = v11
        } else {
            v2 = append(v2, 144)
        }
        v1 = v1
    }
    if false {
        for fi12 := 0; fi12 < 6; fi12++ {
            switch (fi12 >> ((42) & 63)) {
            case 110:
                for fi13 := 0; fi13 < 5; fi13++ {
                    v14 := GenS4{F0: 2147483647, F1: (fi13 / ((0) | 1)), F2: 18446744073709551615}
                    fmt.Println(v3, v3)
                    var v15 int32 = -392
                    delete(v4, 0)
                    if true {
                        var v16 uint16 = (v3 * v3)
                        var v17 uint8 = 0
                        fmt.Println(v15, v3, v5)
                        _ = v16
                        _ = v17
                    }
                    fmt.Println(v1, fi12)
                    v18 := v4[2]
                    _ = fi13
                    _ = v14
                    _ = v15
                    _ = v18
                }
            case 1:
                v1 += v1
            case 4294967295:
                v5 = v5
            }
            for fi19 := 0; fi19 < 5; fi19++ {
                var v20 int = (fi12 & 127)
                _ = fi19
                _ = v20
            }
            var v21 uint = 4294967295
            _ = fi12
            _ = v21
        }
        for fi22 := 0; fi22 < 12; fi22++ {
            v23 := GenS2{F0: 1000, F1: 2}
            v3 -= v3
            var v24 uint16 = 255
            _ = fi22
            _ = v23
            _ = v24
        }
    } else {
        v3 = v3
        v25 := GenS3{F0: 437}
        v5 -= 1e10
        _ = v25
    }
    fmt.Println(v5)
    if false {
        if true {
            v26 := GenS1{F0: v1}
            _ = v26
        } else {
            var v27 int = 9223372036854775807
            for fi28 := 0; fi28 < 10; fi28++ {
                v29 := v4[834]
                switch ("foo bar" + "\n\t") {
                case "foo bar":
                    if true {
                        for fi30 := 0; fi30 < 10; fi30++ {
                            v31 := GenS2{F0: 32767, F1: 42}
                            _ = fi30
                            _ = v31
                        }
                        v5 += v5
                        fmt.Println(v29, v29)
                    }
                case "hello":
                    if (v3 >= v3) {
                        var v32 int16 = 32767
                        if len(v2) > 0 {
                            v2[0] = 165
                        }
                        _ = v32
                    } else {
                        v33 := GenS5{F0: v29, F1: fi28, F2: 0}
                        _ = v33
                    }
                default:
                    var v34 float64 = ((1.5 + 1.5) / (1.5 + float64(0.1)))
                    _ = v34
                }
                for fi35 := 0; fi35 < 12; fi35++ {
                    var v36 uint64 = (v1 / ((v1) | 1))
                    _ = fi35
                    _ = v36
                }
                _ = fi28
                _ = v29
            }
            _ = v27
        }
        if len(v2) > 0 {
            v2[0] = 42
        }
        var v37 string = ""
        _ = v37
    }
    var v38 float64 = ((-2.25 - -2.25) / (0.1 * 3.14159))
    if ((1 <= -678) && (-9223372036854775808 != -1)) {
        v39 := GenS6{F0: 85, F1: 39, F2: 42}
        delete(v4, 0)
        _ = v39
    } else {
        switch "a" {
        case "hello":
            switch "" {
            case "":
                if false {
                    _ = len(v4)
                    var v40 float64 = 3.14159
                    fmt.Println(v1)
                    _ = v40
                }
            default:
                if (((!(9223372036854775807 > 373)) || false) || (false && ((-549 > -1) && (42 == 127)))) {
                    v3 = (^(v3))
                    delete(v4, 452)
                } else {
                    v3 -= v3
                    v41 := GenS2{F0: 267, F1: (v3 - 0)}
                    _ = v41
                }
            }
        case "":
            var v42 uint64 = 2
            _ = v42
        case "foo bar":
            fmt.Println(v3, v3, v3)
        }
    }
    var v43 uint16 = v3
    var v44 float32 = 1e10
    v45 := GenS7{F0: (v3 + v43), F1: -2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v5
}
