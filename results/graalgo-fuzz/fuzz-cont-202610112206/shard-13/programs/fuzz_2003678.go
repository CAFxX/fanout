package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 uint32 }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 string }
func genFn0(p0 int, p1 uint64, p2 int64) int32 {
    v4 := []int{184, (-(p0)), p0}
    v5 := GenS1{F0: 127, F1: -128, F2: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 398
}
func genFn1(p0 bool, p1 uint) int {
    v6 := GenS1{F0: 85, F1: 0, F2: 404}
    v7 := GenS1{F0: 87, F1: -128, F2: 0}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 714
}

func main() {
    var v1 bool = true
    var v2 string = "a"
    var v3 uint32 = 4294967295
    for fi8 := 0; fi8 < 2; fi8++ {
        fmt.Println(v2, fi8)
        var v9 uint8 = uint8(v3)
        var v10 uint = 255
        _ = fi8
        _ = v10
        _ = v9
    }
    if false {
        var v11 int16 = int16(v3)
        v1 = (!v1)
        switch ("a" + "hello") {
        case "foo bar":
            v2 = v2
        case "":
            v12 := GenS2{F0: 42, F1: 255}
            _ = v12
        default:
            if (!(2 == 4294967295)) {
                v11 -= v11
            } else {
                fmt.Println(v3)
                if (1 >= 1) {
                    v1 = false
                }
                for fi13 := 0; fi13 < 16; fi13++ {
                    if v1 {
                        fmt.Println(v1)
                    } else {
                        if (!false) {
                            v2 = "\n\t"
                            for fi14 := 0; fi14 < 1; fi14++ {
                                fmt.Println(fi13, fi13)
                                var v15 uint = 65535
                                if true {
                                    var v16 uint16 = 0
                                    var v17 float64 = (float64(0.1) + (0.0 + 0.1))
                                    _ = v16
                                    _ = v17
                                } else {
                                    v18 := []int{2147483647, 32767}
                                    if len(v18) > 0 {
                                        v18[0] = (fi13 | (fi13 + 2147483647))
                                    }
                                    _ = v18
                                }
                                _ = fi14
                                _ = v15
                            }
                            for fi19 := 0; fi19 < 17; fi19++ {
                                v20 := []int{(fi13 % ((genFn1(false, 65535)) | 1)), -2147483648}
                                _ = fi19
                                _ = v20
                            }
                        }
                    }
                    v2 = v2
                    v21 := GenS3{F0: 2, F1: 0}
                    _ = fi13
                    _ = v21
                }
            }
        }
        _ = v11
    } else {
        v22 := GenS2{F0: 666, F1: 32767}
        var v23 uint64 = 42
        for fi24 := 0; fi24 < 14; fi24++ {
            var v25 bool = v1
            switch v2 {
            case "\n\t":
                v26 := GenS3{F0: 925, F1: -953}
                _ = v26
            }
            var v27 int16 = 42
            _ = fi24
            _ = v25
            _ = v27
        }
        _ = v22
        _ = v23
    }
    v28 := GenS2{F0: 2, F1: 270}
    v29 := GenS2{F0: 2, F1: 0}
    switch 32767 {
    case 9223372036854775807:
        v1 = (v3 <= 4294967295)
    default:
        var v30 bool = ((207 <= 1) && v1)
        _ = v30
    }
    var v31 uint16 = 1
    if ((-2147483648 < 128) || false) {
        v2 = v2
        v31 = v31
    } else {
        v1 = (!v1)
        switch -653 {
        case 713:
            v2 = v2
        default:
            switch v2 {
            case "hello":
                v1 = (-2.25 <= (-2.25 / (((1e10 + 3.14159) - 0.1) * (-2.25 + 0.0))))
            default:
                fmt.Println(v2, v1, v2)
            }
        }
        v32 := []int{genFn1(v1, 4294967295), 2147483647, -9223372036854775808, -456}
        _ = v32
    }
    switch -976 {
    case 127:
        switch 32767 {
        case 42:
            v3 += (v3 ^ v3)
        case 417:
            if true {
                for fi33 := 0; fi33 < 9; fi33++ {
                    v34 := []int{fi33, (-(fi33)), (^(fi33))}
                    _ = len(v34)
                    _ = fi33
                    _ = v34
                }
                v2 = ("hello" + v2)
                v35 := []int{genFn1((9223372036854775807 >= 1), 42), 627, -2147483648}
                _ = v35
            } else {
                v36 := GenS4{F0: (^(v3))}
                v37 := GenS5{F0: "foo bar"}
                _ = v36
                _ = v37
            }
        }
    case 9223372036854775807:
        v38 := []int{128, 42, genFn1(true, 0), 128}
        _ = v38
    case 1:
        v39 := []int{1, -597, 4294967295, 2147483647}
        _ = v39
    default:
        if false {
            fmt.Println(v3, v1)
            var v40 uint64 = 2
            _ = v40
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v31
}
