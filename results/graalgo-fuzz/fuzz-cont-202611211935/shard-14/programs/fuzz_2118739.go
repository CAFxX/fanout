package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8; F1 int; F2 uint }
type GenS4 struct { F0 int64; F1 uint }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint64; F1 uint16 }
type GenS7 struct { F0 uint32; F1 uint16 }
func genFn0(p0 uint32) int32 {
    var v7 int32 = 42
    var v8 bool = true
    if (((p0 / ((1) | 1)) < (p0 >> ((1) & 63))) && v8) {
        if ((v7 & (v7 % ((v7) | 1))) <= v7) {
            var v9 uint16 = 70
            v8 = (32767 == 42)
            _ = v9
        } else {
            p0 += 65535
            for fi10 := 0; fi10 < 2; fi10++ {
                switch fi10 {
                case 1:
                    v11 := GenS1{F0: 42}
                    _ = v11
                case 32768:
                    if false {
                        var v12 int = (fi10 | fi10)
                        var v13 float64 = (float64(1e10) * 1e10)
                        v8 = (!false)
                        _ = v12
                        _ = v13
                    } else {
                        switch "\n\t" {
                        case "foo bar":
                            if false {
                                var v14 uint32 = (p0 >> ((917) & 63))
                                if false {
                                    var v15 uint32 = (v14 ^ (v14 + p0))
                                    _ = v15
                                }
                                _ = v14
                            }
                        }
                        var v16 uint64 = 42
                        var v17 string = "a"
                        _ = v16
                        _ = v17
                    }
                }
                for fi18 := 0; fi18 < 14; fi18++ {
                    v19 := GenS2{F0: 280}
                    var v20 int64 = -896
                    v7 += v7
                    switch "foo bar" {
                    case "a":
                        var v21 int8 = -65
                        _ = v21
                    case "":
                        if false {
                            var v22 int = fi18
                            p0 -= (p0 % (((p0 | (p0 ^ p0))) | 1))
                            var v23 int8 = 100
                            _ = v22
                            _ = v23
                        } else {
                            var v24 uint8 = uint8(p0)
                            v24 = v24
                            v25 := GenS1{F0: 738}
                            _ = v24
                            _ = v25
                        }
                        v8 = (1e10 < -2.25)
                        v26 := []int{fi10, -2147483648, (fi18 | fi10), 0}
                        _ = v26
                    case "hello":
                        fmt.Println(v8, fi18)
                    }
                    var v27 int64 = v20
                    fmt.Println(v20)
                    _ = fi18
                    _ = v19
                    _ = v20
                    _ = v27
                }
                p0 -= (p0 - p0)
                _ = fi10
            }
        }
    } else {
        v28 := GenS1{F0: 109}
        v29 := GenS3{F0: 127, F1: 2, F2: 4294967295}
        _ = v28
        _ = v29
    }
    v30 := GenS4{F0: 42, F1: 293}
    v31 := GenS1{F0: 655}
    _ = p0
    _ = v30
    _ = v31
    _ = v7
    _ = v8
    return (v7 / (((v7 & v7)) | 1))
}

func main() {
    var v1 int32 = 0
    var v2 int8 = 100
    v3 := make(map[string]string)
    var v4 int64 = 672
    v5 := []string{"foo bar", "hello", "", "a"}
    var v6 uint64 = 0
    if (((float32((3.14159 * (1e10 + 1.5))) != (float32(3.14159) + -2.25)) || (false || false)) && true) {
        v1 = v1
    }
    fmt.Println(v6, v1)
    var v32 uint = 255
    v33 := GenS2{F0: 4294967295}
    if false {
        v6 = (v6 << ((v32) & 63))
    } else {
        if len(v5) > 0 {
            v5[0] = "\n\t"
        }
        if false {
            for fi34 := 0; fi34 < 3; fi34++ {
                v32 += (v32 << ((v32) & 63))
                if true {
                    for fi35 := 0; fi35 < 3; fi35++ {
                        if true {
                            _ = len(v3)
                            v4 = v4
                        } else {
                            switch "hello" {
                            case "\n\t":
                                var v36 int64 = 9223372036854775807
                                _ = v36
                            case "foo bar":
                                if false {
                                    v37 := GenS3{F0: 255, F1: fi34, F2: (^(v32))}
                                    var v38 uint = v32
                                    _ = v37
                                    _ = v38
                                }
                            default:
                                v6 += v6
                            }
                            if false {
                                var v39 int16 = 32767
                                _ = v39
                            } else {
                                switch "a" {
                                case "a":
                                    var v40 uint16 = 673
                                    _ = v40
                                case "\n\t":
                                    v1 = (v1 / ((v1) | 1))
                                case "foo bar":
                                    v41 := GenS4{F0: 0, F1: (v32 - (v32 / ((v32) | 1)))}
                                    _ = v41
                                default:
                                    switch ("a" + "\n\t") {
                                    case "\n\t":
                                        v42 := v3[("hello" + "hello")]
                                        _ = v42
                                    }
                                }
                            }
                        }
                        fmt.Println(v6)
                        _ = fi35
                    }
                    fmt.Println(fi34, v32)
                } else {
                    delete(v3, ((("\n\t" + ("" + ("\n\t" + ((("a" + "") + (("a" + "a") + "a")) + "\n\t")))) + "\n\t") + "a"))
                    v43 := GenS2{F0: 2}
                    _ = v43
                }
                v1 += v1
                _ = fi34
            }
            for fi44 := 0; fi44 < 9; fi44++ {
                var v45 uint8 = 30
                if (fi44 < -9223372036854775808) {
                    for fi46 := 0; fi46 < 15; fi46++ {
                        v1 = v1
                        var v47 int8 = v2
                        _ = fi46
                        _ = v47
                    }
                    var v48 float32 = float32(1e10)
                    _ = v48
                }
                _ = fi44
                _ = v45
            }
            v4 += (v4 - v4)
        } else {
            var v49 int16 = -217
            var v50 bool = (!true)
            _ = v49
            _ = v50
        }
        fmt.Println(v4, v4)
    }
    if (("\n\t" + "hello") >= "hello") {
        if (!true) {
            v51 := v3["\n\t"]
            v32 = v32
            v4 = (v4 >> (((v32 * (v32 ^ 1))) & 63))
            _ = v51
        }
    }
    v52 := GenS5{F0: 1}
    var v53 int64 = (v4 * (v4 + v4))
    var v54 uint8 = 1
    for fi55 := 0; fi55 < 16; fi55++ {
        _ = len(v3)
        if len(v5) > 0 {
            v5[0] = "foo bar"
        }
        _ = fi55
    }
    _ = len(v3)
    v56 := GenS6{F0: (v6 - (v6 & v6)), F1: 2}
    if false {
        fmt.Println(v1)
        v57 := GenS7{F0: 65535, F1: 1}
        _ = v57
    }
    delete(v3, (("\n\t" + ("hello" + "hello")) + "\n\t"))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v54
    _ = v56
    _ = v6
}
