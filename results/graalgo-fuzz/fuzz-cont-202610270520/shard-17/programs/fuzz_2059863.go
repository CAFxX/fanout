package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32; F1 uint; F2 uint32 }
type GenS3 struct { F0 bool }
func genFn0(p0 uint, p1 uint64, p2 uint16) uint8 {
    var v7 float64 = 1.5
    v8 := []int{128, 42, -30}
    var v9 int8 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return uint8(p2)
}

func main() {
    var v1 float64 = (-2.25 / 1.5)
    v2 := []int{0, 90, 0, 128}
    v3 := make(map[int]int)
    var v4 int8 = 1
    var v5 uint8 = 172
    v6 := make(map[string]string)
    var v10 int64 = 9223372036854775807
    v1 = 0.1
    var v11 float32 = 1e10
    v4 = v4
    if true {
        var v12 float32 = 1e10
        v11 += (((-2.25 + ((0.1 - float32(v1)) + ((((v12 + float32(1.5)) - (v11 / float32(1.5))) + v11) * 0.0))) / v12) + ((v12 - v12) * (((float32((v1 / (((0.1 - v1) + v1) / ((1e10 - v1) - 0.0)))) * (v11 - (v12 + 1.5))) * float32(-2.25)) / 1.0)))
        _ = v12
    }
    if (!false) {
        var v13 string = "a"
        delete(v3, 9223372036854775807)
        switch 42 {
        case -914:
            var v14 int = -910
            _ = v14
        case 0:
            fmt.Println(v5)
        case 23:
            if (false && true) {
                for fi15 := 0; fi15 < 2; fi15++ {
                    if true {
                        if (!(1 != 65535)) {
                            v1 += float64(v11)
                            fmt.Println(v1)
                        }
                        for fi16 := 0; fi16 < 5; fi16++ {
                            var v17 int = (fi15 / (((^(fi16))) | 1))
                            _ = fi16
                            _ = v17
                        }
                    } else {
                        _ = len(v3)
                        fmt.Println(fi15, fi15)
                        v1 -= v1
                    }
                    _ = fi15
                }
            }
        }
        _ = v13
    } else {
        v11 += (0.1 * 0.1)
        switch ("hello" + ("\n\t" + "\n\t")) {
        case "foo bar":
            for fi18 := 0; fi18 < 15; fi18++ {
                var v19 float64 = (v1 + v1)
                if (("foo bar" + ("hello" + ("foo bar" + "\n\t"))) > "") {
                    for fi20 := 0; fi20 < 13; fi20++ {
                        if false {
                            var v21 int16 = 0
                            fmt.Println(v10, v19, fi18)
                            _ = v21
                        } else {
                            v22 := GenS1{F0: 42}
                            _ = v22
                        }
                        fmt.Println(v11)
                        _ = fi20
                    }
                } else {
                    var v23 int16 = 0
                    _ = v23
                }
                var v24 uint = 790
                _ = fi18
                _ = v19
                _ = v24
            }
        case "a":
            v25 := GenS1{F0: (v4 + v4)}
            _ = v25
        case "":
            fmt.Println(v4, v10, v4)
        default:
            v4 += v4
        }
    }
    v10 += v10
    var v26 int16 = 42
    if false {
        var v27 float64 = v1
        _ = v27
    } else {
        for fi28 := 0; fi28 < 17; fi28++ {
            if (-1 > 42) {
                var v29 int16 = v26
                _ = v29
            }
            var v30 int = (fi28 & (fi28 + (fi28 - 733)))
            switch "hello" {
            case "foo bar":
                v6["foo bar"] = ""
            case "":
                switch fi28 {
                case -1:
                    v31 := GenS2{F0: 65535, F1: 1, F2: 1}
                    _ = v31
                case 127:
                    if false {
                        var v32 int64 = (v10 % (((^(v10))) | 1))
                        var v33 uint64 = 549
                        v34 := GenS2{F0: 366, F1: 962, F2: 4294967295}
                        _ = v32
                        _ = v33
                        _ = v34
                    }
                default:
                    if false {
                        var v35 bool = true
                        _ = v35
                    } else {
                        for fi36 := 0; fi36 < 6; fi36++ {
                            if false {
                                v37 := GenS1{F0: int8(v30)}
                                v38 := GenS2{F0: 42, F1: 934, F2: 2147483648}
                                _ = v37
                                _ = v38
                            } else {
                                if true {
                                    v39 := GenS1{F0: v4}
                                    _ = v39
                                } else {
                                    _ = len(v3)
                                }
                                if false {
                                    var v40 uint16 = 42
                                    var v41 string = "\n\t"
                                    _ = v40
                                    _ = v41
                                }
                            }
                            _ = fi36
                        }
                    }
                }
            }
            _ = fi28
            _ = v30
        }
        fmt.Println(v26)
    }
    var v42 int32 = 808
    var v43 int16 = (v26 - -32768)
    v44 := GenS3{F0: (true && (int8(v43) < v4))}
    v43 += (v26 * (^(v43)))
    fmt.Println(v26)
    if (980 != 255) {
        v3[2147483647] = 42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
