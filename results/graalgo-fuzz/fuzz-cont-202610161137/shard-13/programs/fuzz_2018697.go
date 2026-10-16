package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 int16; F1 bool }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint16; F1 uint32; F2 int32 }
func genFn0(p0 int32) int64 {
    var v7 int8 = 42
    if true {
        var v8 int8 = (v7 * v7)
        v9 := []int{-2147483648}
        _ = v8
        _ = v9
    } else {
        var v10 uint = 0
        _ = v10
    }
    if (true || (-9223372036854775808 >= 0)) {
        v11 := GenS1{F0: 9223372036854775807}
        if (((v7 - -104) == (v7 - (-(v7)))) && (false && (!true))) {
            var v12 string = ((("" + "") + "\n\t") + ("" + "\n\t"))
            _ = v12
        } else {
            for fi13 := 0; fi13 < 4; fi13++ {
                v7 -= 127
                var v14 int64 = -9223372036854775808
                _ = fi13
                _ = v14
            }
            v7 += v7
        }
        var v15 int32 = p0
        _ = v11
        _ = v15
    }
    _ = p0
    _ = v7
    return -1
}
func genFn1() uint {
    switch 42 {
    case -1:
        var v16 int32 = 42
        _ = v16
    case 127:
        v17 := []int{}
        _ = v17
    case -34:
        var v18 int32 = 2147483647
        _ = v18
    }
    for fi19 := 0; fi19 < 2; fi19++ {
        for fi20 := 0; fi20 < 17; fi20++ {
            if (true || true) {
                fmt.Println(fi19, fi20)
                var v21 int64 = 9223372036854775807
                _ = v21
            } else {
                switch (fi19 | fi19) {
                case 4294967295:
                    fmt.Println(fi19)
                default:
                    var v22 int32 = 43
                    _ = v22
                }
                var v23 int8 = 1
                _ = v23
            }
            _ = fi20
        }
        var v24 string = ("\n\t" + "hello")
        _ = fi19
        _ = v24
    }
    var v25 int32 = 0
    _ = v25
    return 376
}

func main() {
    v1 := []uint32{2147483647, 2147483647, 65535, 157}
    v2 := make(map[uint32]bool)
    var v3 int8 = -128
    v4 := make(map[uint32]bool)
    var v5 uint = 18446744073709551615
    v6 := make(map[uint32]string)
    var v26 int64 = 0
    fmt.Println(v3)
    switch "hello" {
    case "hello":
        fmt.Println(v5, v26, v3)
    case "a":
        v27 := v4[856]
        _ = v27
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v28 := GenS1{F0: genFn0(-1)}
    switch -9223372036854775808 {
    case 127:
        switch -2147483648 {
        case -2147483648:
            var v29 uint8 = 128
            _ = v29
        default:
            _ = len(v1)
        }
    default:
        for fi30 := 0; fi30 < 17; fi30++ {
            if (-32768 > -370) {
                if (!true) {
                    if (18446744073709551615 < 2) {
                        var v31 string = ("\n\t" + "")
                        v32 := GenS2{F0: false, F1: (v3 ^ v3)}
                        _ = v31
                        _ = v32
                    }
                    fmt.Println(v5, fi30, v3)
                } else {
                    switch (fi30 << (((v5 >> (((v5 ^ v5)) & 63))) & 63)) {
                    case 326:
                        fmt.Println(v5)
                    case 4294967295:
                        var v33 float32 = (float32(1e10) + (0.0 + (1e10 / -2.25)))
                        _ = v33
                    case 1:
                        for fi34 := 0; fi34 < 17; fi34++ {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            v35 := GenS3{F0: 1000, F1: (!((v5 | genFn1()) != v5))}
                            if false {
                                if ((!(3.14159 != (1.5 + 1e10))) && (false && (true || (227 >= 255)))) {
                                    if (!true) {
                                        v3 += 42
                                    } else {
                                        fmt.Println(v3, fi30)
                                        var v36 int32 = 0
                                        _ = v36
                                    }
                                    var v37 int32 = 483
                                    _ = v37
                                } else {
                                    v2[4294967295] = true
                                    fmt.Println(v3, fi30, v5)
                                    var v38 string = ("" + (("" + "foo bar") + ("\n\t" + "hello")))
                                    _ = v38
                                }
                            } else {
                                var v39 uint32 = 2147483648
                                v26 -= (v26 ^ v26)
                                _ = v39
                            }
                            v40 := GenS4{F0: 0}
                            switch int(v5) {
                            case 32768:
                                var v41 uint16 = 255
                                _ = v41
                            case 32767:
                                if (v5 > 727) {
                                    var v42 int8 = 127
                                    _ = v42
                                }
                                var v43 int8 = -14
                                _ = v43
                            }
                            _ = fi34
                            _ = v35
                            _ = v40
                        }
                    default:
                        var v44 int16 = -1
                        _ = v44
                    }
                }
                fmt.Println(v26, v5)
            }
            _ = fi30
        }
    }
    fmt.Println(v5, v3)
    var v45 uint8 = 23
    v46 := GenS5{F0: 65535, F1: 244, F2: 42}
    v3 -= (v3 ^ (v3 ^ v3))
    fmt.Println(v26)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v26
    _ = v28
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
