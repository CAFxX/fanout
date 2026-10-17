package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint8 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 string; F1 int8; F2 int }
type GenS5 struct { F0 bool; F1 bool }
type GenS6 struct { F0 uint64; F1 uint32 }
func genFn0(p0 bool, p1 int32) uint32 {
    p0 = (2 >= 2)
    var v7 int32 = 1
    switch "foo bar" {
    case "\n\t":
        p0 = ("foo bar" > "\n\t")
    case "foo bar":
        p0 = p0
    case "hello":
        var v8 int16 = 132
        _ = v8
    default:
        if p0 {
            var v9 bool = p0
            _ = v9
        } else {
            p0 = p0
        }
    }
    p1 += p1
    _ = p0
    _ = p1
    _ = v7
    return 4294967295
}

func main() {
    var v1 int = 42
    var v2 int8 = 127
    var v3 int8 = v2
    v4 := make(map[string]string)
    var v5 uint8 = 1
    var v6 uint32 = 828
    v5 = 42
    if true {
        var v10 int32 = -1
        _ = v10
    } else {
        var v11 float64 = float64(((0.1 / (0.1 + 1.5)) * 0.1))
        var v12 int8 = v3
        delete(v4, "hello")
        _ = v11
        _ = v12
    }
    v5 += (v5 % ((52) | 1))
    v13 := v4["a"]
    var v14 int32 = -859
    var v15 int = (v1 & (v1 << ((65535) & 63)))
    v16 := GenS1{F0: (-(v14)), F1: v1, F2: v5}
    v6 = (v6 >> ((65535) & 63))
    v17 := GenS1{F0: v14, F1: -9223372036854775808, F2: v5}
    for fi18 := 0; fi18 < 1; fi18++ {
        if (128 == v5) {
            for fi19 := 0; fi19 < 3; fi19++ {
                var v20 uint64 = 4294967295
                if (false && (1 == 1)) {
                    var v21 float32 = -2.25
                    _ = v21
                } else {
                    if true {
                        v6 += v6
                        var v22 float32 = float32(((-2.25 + (0.0 / 1.0)) + 0.0))
                        switch v13 {
                        case "":
                            delete(v4, (v13 + "\n\t"))
                        case "foo bar":
                            var v23 float64 = ((-2.25 * -2.25) + 0.0)
                            _ = v23
                        case "\n\t":
                            if true {
                                v6 -= v6
                                if false {
                                    v24 := GenS1{F0: v14, F1: fi19, F2: v5}
                                    var v25 bool = true
                                    _ = v24
                                    _ = v25
                                } else {
                                    switch v15 {
                                    case 127:
                                        _ = len(v4)
                                    }
                                }
                                v26 := GenS2{F0: 97, F1: -33}
                                switch (fi19 >> ((65535) & 63)) {
                                case -9223372036854775808:
                                    var v27 float32 = ((v22 - v22) - ((-2.25 + v22) / v22))
                                    _ = v27
                                case 1:
                                    delete(v4, ((v13 + "a") + (v13 + v13)))
                                case 32768:
                                    if false {
                                        v13 = v13
                                        v28 := []int{(fi19 >> ((1) & 63))}
                                        if len(v28) >= 2 {
                                            v28 = v28[1:len(v28)]
                                        }
                                        _ = v28
                                    } else {
                                        switch "hello" {
                                        case "\n\t":
                                            fmt.Println(v1, fi19, fi18)
                                        default:
                                            var v29 uint8 = (v5 << ((65535) & 63))
                                            _ = v29
                                        }
                                        v30 := GenS3{F0: (^(fi18)), F1: (v5 / ((42) | 1))}
                                        switch -839 {
                                        case -1:
                                            fmt.Println(v20, v13)
                                        default:
                                            var v31 uint8 = 92
                                            _ = v31
                                        }
                                        switch (fi19 & v15) {
                                        case 32767:
                                            if false {
                                                var v32 float64 = float64(3.14159)
                                                _ = v32
                                            }
                                        default:
                                            _ = len(v4)
                                        }
                                        _ = v30
                                    }
                                }
                                v33 := GenS1{F0: v14, F1: fi18, F2: v5}
                                var v34 uint8 = 42
                                _ = v26
                                _ = v33
                                _ = v34
                            }
                        }
                        if ("" != "\n\t") {
                            var v35 uint = 866
                            v36 := GenS4{F0: (((("\n\t" + v13) + "") + v13) + (v13 + ("hello" + (v13 + v13)))), F1: (v2 % ((int8(v5)) | 1)), F2: -9223372036854775808}
                            _ = v35
                            _ = v36
                        } else {
                            if true {
                                v20 = uint64(v6)
                            } else {
                                var v37 uint = 169
                                fmt.Println(v15, v20)
                                _ = v37
                            }
                            _ = len(v4)
                            fmt.Println(v6)
                        }
                        var v38 float64 = float64(1.5)
                        _ = v22
                        _ = v38
                    } else {
                        var v39 uint16 = 42
                        _ = v39
                    }
                    v40 := []int{v15}
                    v20 = (v20 >> ((18446744073709551615) & 63))
                    _ = v40
                }
                _ = fi19
                _ = v20
            }
            v41 := GenS5{F0: false, F1: (1 > v2)}
            _ = v41
        } else {
            fmt.Println(v1)
            var v42 bool = true
            v1 += (v1 ^ (v1 | 42))
            _ = v42
        }
        v14 -= v14
        _ = fi18
    }
    for fi43 := 0; fi43 < 4; fi43++ {
        v44 := GenS4{F0: v13, F1: (v3 % ((v2) | 1)), F2: fi43}
        _ = fi43
        _ = v44
    }
    _ = len(v4)
    v45 := []int{}
    v46 := GenS6{F0: 519, F1: (v6 * v6)}
    fmt.Println(v1, v14)
    var v47 uint32 = genFn0(true, (v14 + v14))
    if len(v45) >= 2 {
        v45 = v45[1:len(v45)]
    }
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}
