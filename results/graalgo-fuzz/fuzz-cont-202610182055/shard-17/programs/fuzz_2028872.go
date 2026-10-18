package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint; F1 int; F2 bool }
type GenS4 struct { F0 int16; F1 int; F2 int32 }
type GenS5 struct { F0 int64; F1 int32 }
func genFn0(p0 int32, p1 int8, p2 int16) string {
    p2 += p2
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    return "hello"
}
func genFn1() int8 {
    if true {
        var v4 uint8 = 127
        var v5 uint = 77
        _ = v4
        _ = v5
    }
    var v6 uint32 = 42
    v6 = v6
    _ = v6
    return -6
}

func main() {
    v1 := make(map[string]int)
    v2 := []int{128}
    var v3 uint = 341
    v3 -= (v3 * (v3 << ((4294967295) & 63)))
    v3 += (v3 * (v3 % (((v3 % ((v3) | 1))) | 1)))
    var v7 int64 = -9223372036854775808
    switch -2147483648 {
    case -9223372036854775808:
        v8 := GenS1{F0: -1, F1: 128}
        _ = v8
    case -588:
        var v9 int16 = 32767
        _ = v9
    default:
        for fi10 := 0; fi10 < 9; fi10++ {
            var v11 uint32 = 2147483648
            var v12 int = (^(fi10))
            var v13 uint16 = 0
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
    }
    var v14 float32 = 1e10
    v15 := GenS1{F0: 173, F1: 255}
    var v16 float64 = float64(((v14 - v14) / (v14 + (float32(0.0) * (v14 / float32(3.14159))))))
    if true {
        v16 += float64(1e10)
        v17 := GenS1{F0: 0, F1: 255}
        _ = v17
    } else {
        for fi18 := 0; fi18 < 14; fi18++ {
            if false {
                var v19 uint16 = 65535
                if true {
                    var v20 uint16 = v19
                    for fi21 := 0; fi21 < 19; fi21++ {
                        var v22 int = 128
                        if (!true) {
                            v3 -= v3
                        }
                        _ = fi21
                        _ = v22
                    }
                    if false {
                        for fi23 := 0; fi23 < 2; fi23++ {
                            var v24 uint8 = 116
                            if (("foo bar" + (("\n\t" + "\n\t") + "\n\t")) <= "foo bar") {
                                switch "" {
                                case "hello":
                                    var v25 int = fi23
                                    _ = v25
                                default:
                                    var v26 int = fi18
                                    v19 = (^(v20))
                                    v7 = v7
                                    _ = v26
                                }
                                fmt.Println(fi18, v3)
                            }
                            v27 := GenS1{F0: -1, F1: (^(v24))}
                            fmt.Println(fi18)
                            switch ("\n\t" + ((("" + "hello") + "\n\t") + "")) {
                            case "\n\t":
                                v28 := GenS2{F0: -2147483648}
                                _ = v28
                            case "foo bar":
                                v24 = (v24 % ((1) | 1))
                            case "a":
                                fmt.Println(v14)
                                if true {
                                    v7 -= v7
                                    v14 = v14
                                } else {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    v29 := GenS3{F0: (v3 * v3), F1: fi18, F2: ((42 < -121) || true)}
                                    _ = v29
                                }
                                var v30 uint16 = 65535
                                _ = v30
                            default:
                                fmt.Println(v19)
                                var v31 uint64 = uint64(v24)
                                _ = v31
                            }
                            var v32 bool = ((v16 / (float64((1e10 / (v14 / 0.1))) + (float64(v14) + v16))) == ((v16 / 1.0) + 0.1))
                            v33 := GenS4{F0: -840, F1: fi23, F2: int32(v7)}
                            var v34 uint = (v3 >> ((v3) & 63))
                            _ = fi23
                            _ = v24
                            _ = v27
                            _ = v32
                            _ = v33
                            _ = v34
                        }
                        v16 -= ((v16 - (v16 - (v16 - (v16 - v16)))) - (float64(((v14 * v14) + (v14 - 1e10))) / (float64(v14) - v16)))
                        _ = len(v2)
                    } else {
                        var v35 uint = (v3 >> ((v3) & 63))
                        _ = v35
                    }
                    _ = v20
                } else {
                    var v36 int64 = (v7 / (((v7 * 9223372036854775807)) | 1))
                    _ = v36
                }
                _ = v19
            }
            fmt.Println(v14, v16)
            for fi37 := 0; fi37 < 7; fi37++ {
                v38 := v1["foo bar"]
                v16 += v16
                var v39 bool = true
                _ = fi37
                _ = v38
                _ = v39
            }
            _ = fi18
        }
    }
    var v40 int16 = 32767
    var v41 int32 = 503
    v42 := GenS5{F0: 42, F1: v41}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v40
    _ = v41
    _ = v42
    _ = v7
}
