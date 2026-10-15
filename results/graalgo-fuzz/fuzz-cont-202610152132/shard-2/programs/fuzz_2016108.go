package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 uint; F1 bool; F2 bool }
type GenS3 struct { F0 uint32; F1 bool; F2 uint8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int8; F1 uint }
func genFn0(p0 uint) uint32 {
    if true {
        fmt.Println(p0)
        var v5 float32 = float32(((3.14159 + 1.5) - 3.14159))
        switch 415 {
        case 32767:
            v6 := []int{42, -1, 32768, -241}
            _ = v6
        case 42:
            fmt.Println(v5)
        }
        _ = v5
    } else {
        var v7 uint32 = 65535
        v8 := []int{-1}
        _ = v7
        _ = v8
    }
    fmt.Println(p0)
    _ = p0
    return 2147483647
}
func genFn1(p0 int32, p1 uint, p2 int8) uint64 {
    var v9 string = "hello"
    v10 := []int{-2147483648, 0, 42}
    v9 = v9
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 292
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]uint32)
    v3 := []uint32{2147483648, 1, 88}
    var v4 int8 = 0
    v11 := v1[("" + (("hello" + ("" + "\n\t")) + "foo bar"))]
    switch 128 {
    case -365:
        v12 := v2["a"]
        _ = v12
    default:
        var v13 float64 = (1e10 - 3.14159)
        _ = v13
    }
    v3 = append(v3, genFn0(0))
    for fi14 := 0; fi14 < 14; fi14++ {
        fmt.Println(v4, v11, fi14)
        v15 := GenS1{F0: "", F1: 0}
        _ = fi14
        _ = v15
    }
    v3 = append(v3, 2147483648)
    var v16 string = v11
    v17 := GenS2{F0: 18446744073709551615, F1: true, F2: (!(1 < 18446744073709551615))}
    v4 = v4
    switch 1 {
    case 128:
        if (false && false) {
            var v18 string = "hello"
            if (!true) {
                if true {
                    for fi19 := 0; fi19 < 18; fi19++ {
                        if (true && (true || (!true))) {
                            if false {
                                delete(v1, (((("" + v18) + v18) + v16) + "a"))
                            }
                            v20 := GenS2{F0: 2, F1: (-558 > -236), F2: (1 != 312)}
                            for fi21 := 0; fi21 < 17; fi21++ {
                                v4 -= (-(v4))
                                _ = fi21
                            }
                            _ = v20
                        }
                        var v22 float64 = 0.0
                        v23 := GenS3{F0: 42, F1: true, F2: 1}
                        _ = fi19
                        _ = v22
                        _ = v23
                    }
                    var v24 int16 = 0
                    _ = v24
                } else {
                    if (!true) {
                        _ = len(v2)
                        var v25 int16 = -881
                        _ = v25
                    } else {
                        if false {
                            v11 = v11
                            v1[v16] = v18
                            if true {
                                var v26 int32 = 0
                                _ = v26
                            } else {
                                if false {
                                    v27 := GenS4{F0: v18}
                                    _ = v27
                                } else {
                                    var v28 bool = (false || true)
                                    _ = v28
                                }
                                if true {
                                    var v29 int16 = 32767
                                    for fi30 := 0; fi30 < 11; fi30++ {
                                        v16 = ((v16 + "") + (v16 + v18))
                                        v4 += 42
                                        var v31 int = fi30
                                        _ = fi30
                                        _ = v31
                                    }
                                    _ = v29
                                }
                            }
                        }
                    }
                    fmt.Println(v18, v16, v16)
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            } else {
                for fi32 := 0; fi32 < 2; fi32++ {
                    if true {
                        v11 = v16
                        fmt.Println(fi32, v16, fi32)
                        if ((true && (false && (false && (!(false && false))))) && false) {
                            for fi33 := 0; fi33 < 8; fi33++ {
                                v34 := GenS5{F0: (v4 % (((v4 - 127)) | 1)), F1: 42}
                                var v35 int32 = -2147483648
                                _ = fi33
                                _ = v34
                                _ = v35
                            }
                            var v36 uint8 = 128
                            fmt.Println(v36, v18, v16)
                            _ = v36
                        } else {
                            fmt.Println(fi32)
                        }
                    }
                    _ = fi32
                }
                fmt.Println(v4, v16, v4)
                v4 -= (v4 >> ((0) & 63))
            }
            v11 = v18
            _ = v18
        } else {
            v37 := GenS5{F0: (v4 * 0), F1: 255}
            var v38 int32 = -1
            _ = v37
            _ = v38
        }
    default:
        delete(v1, (v11 + ((v16 + v16) + v11)))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
