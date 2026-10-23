package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 int; F1 int64; F2 int }
type GenS3 struct { F0 int }

func main() {
    var v1 int8 = -128
    var v2 int = -2147483648
    var v3 int8 = -67
    v4 := make(map[int]int)
    var v5 uint64 = 768
    if false {
        if true {
            if false {
                var v6 bool = false
                _ = v6
            } else {
                switch "a" {
                case "":
                    var v7 int32 = 1
                    _ = v7
                }
                v2 += (v2 % (((v2 >> ((255) & 63))) | 1))
            }
            var v8 uint = 1
            if true {
                v9 := []int{42, v2, -657}
                v10 := v4[-1]
                if ((false || (v8 != v8)) && (v8 > 0)) {
                    for fi11 := 0; fi11 < 16; fi11++ {
                        if false {
                            v1 = (^(v1))
                        } else {
                            v3 += (v1 ^ v3)
                            v3 += (v1 & v3)
                        }
                        var v12 uint = v8
                        _ = fi11
                        _ = v12
                    }
                }
                _ = v10
                _ = v9
            } else {
                switch (v2 % ((453) | 1)) {
                case 4294967295:
                    var v13 int = (v2 << ((v8) & 63))
                    _ = v13
                case 1:
                    var v14 int8 = (v3 ^ -1)
                    _ = v14
                case 144:
                    v15 := GenS1{F0: 2147483647, F1: ("hello" + (("a" + "\n\t") + ""))}
                    _ = v15
                }
                var v16 int64 = 42
                v4[int(v5)] = v2
                _ = v16
            }
            _ = v8
        } else {
            v17 := v4[9223372036854775807]
            if (2147483647 >= 4294967295) {
                var v18 int8 = 42
                v17 += v2
                switch "\n\t" {
                case "":
                    for fi19 := 0; fi19 < 20; fi19++ {
                        var v20 float32 = 1e10
                        _ = fi19
                        _ = v20
                    }
                default:
                    if ((0.1 + 1e10) < (1e10 - -2.25)) {
                        v2 = v17
                        for fi21 := 0; fi21 < 10; fi21++ {
                            if false {
                                fmt.Println(fi21, v1, v1)
                                if (false || (-763 <= -2147483648)) {
                                    v22 := GenS2{F0: fi21, F1: 1, F2: (v2 >> ((65535) & 63))}
                                    v2 += (fi21 | (^(fi21)))
                                    var v23 int = (-(v17))
                                    _ = v22
                                    _ = v23
                                } else {
                                    v5 -= (v5 / ((2) | 1))
                                    v24 := []int{}
                                    delete(v4, (v2 & v2))
                                    _ = v24
                                }
                            } else {
                                v25 := GenS1{F0: 358, F1: "hello"}
                                delete(v4, 32767)
                                _ = v25
                            }
                            _ = fi21
                        }
                        for fi26 := 0; fi26 < 9; fi26++ {
                            var v27 float32 = ((0.1 * (3.14159 * 3.14159)) + float32((3.14159 - 0.1)))
                            _ = len(v4)
                            _ = fi26
                            _ = v27
                        }
                    } else {
                        var v28 uint = 1
                        _ = v28
                    }
                }
                _ = v18
            } else {
                if false {
                    v17 += (v2 + 0)
                    var v29 int32 = -89
                    _ = v29
                }
                for fi30 := 0; fi30 < 20; fi30++ {
                    v31 := GenS3{F0: (v2 << ((42) & 63))}
                    fmt.Println(v1, v2, v17)
                    _ = fi30
                    _ = v31
                }
                var v32 uint8 = 127
                _ = v32
            }
            v1 = (v1 ^ v1)
            _ = v17
        }
        var v33 int32 = 1
        for fi34 := 0; fi34 < 16; fi34++ {
            switch (^(v2)) {
            case 127:
                switch "" {
                case "a":
                    for fi35 := 0; fi35 < 18; fi35++ {
                        v5 += v5
                        _ = fi35
                    }
                case "":
                    var v36 int16 = 0
                    _ = v36
                default:
                    var v37 int16 = 0
                    _ = v37
                }
            case 261:
                v4[2147483647] = v2
            case 210:
                var v38 bool = (!false)
                _ = v38
            default:
                v39 := []int{}
                _ = v39
            }
            var v40 float64 = 1e10
            var v41 int = (^(v2))
            _ = fi34
            _ = v40
            _ = v41
        }
        _ = v33
    }
    v2 += (-(v2))
    v42 := []int{}
    v43 := GenS1{F0: 0, F1: ("" + ((("a" + "") + ("" + (("hello" + ((("\n\t" + "hello") + (("" + (("foo bar" + ("\n\t" + "\n\t")) + "hello")) + ("a" + "a"))) + (("hello" + ("foo bar" + (("" + ("hello" + "foo bar")) + "a"))) + ""))) + ""))) + "foo bar"))}
    fmt.Println(v1)
    var v44 int = -567
    var v45 int16 = int16(v5)
    var v46 uint32 = 2
    v3 = 100
    fmt.Println(v1, v45, v1)
    v5 += v5
    if true {
        v2 = 414
        v3 -= 106
        fmt.Println(v3, v3, v3)
    }
    var v47 uint64 = (^(v5))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v5
}
