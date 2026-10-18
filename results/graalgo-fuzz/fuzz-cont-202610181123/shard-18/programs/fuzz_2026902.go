package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 int32; F2 int }
type GenS3 struct { F0 int16; F1 int; F2 int8 }
type GenS4 struct { F0 bool; F1 int8 }
func genFn0() uint16 {
    v5 := []int{2147483647, 32767, 32768, -1}
    switch "foo bar" {
    case "\n\t":
        if ((-2.25 * float64(0.1)) != ((1e10 + 1.5) / 0.1)) {
            if len(v5) > 0 {
                v5[0] = 397
            }
            for fi6 := 0; fi6 < 15; fi6++ {
                var v7 string = "foo bar"
                if false {
                    v7 = ""
                    v7 = (((((v7 + v7) + v7) + v7) + "hello") + v7)
                    var v8 uint16 = 808
                    _ = v8
                }
                _ = fi6
                _ = v7
            }
            if false {
                var v9 int = 0
                var v10 uint32 = 42
                v9 += (^(v9))
                _ = v10
                _ = v9
            }
        } else {
            v11 := GenS1{F0: 2}
            _ = v11
        }
    case "":
        switch 2147483647 {
        case 1:
            for fi12 := 0; fi12 < 15; fi12++ {
                fmt.Println(fi12)
                var v13 float64 = float64(0.0)
                var v14 uint64 = 2
                _ = fi12
                _ = v13
                _ = v14
            }
        }
    case "a":
        if len(v5) > 0 {
            v5[0] = 2147483647
        }
    default:
        var v15 int = 816
        _ = v15
    }
    switch -341 {
    case 1:
        var v16 float64 = -2.25
        _ = v16
    default:
        for fi17 := 0; fi17 < 14; fi17++ {
            for fi18 := 0; fi18 < 11; fi18++ {
                var v19 int64 = 42
                switch fi18 {
                case 0:
                    if false {
                        var v20 int = 792
                        _ = v20
                    }
                default:
                    switch -9223372036854775808 {
                    case -294:
                        if true {
                            if false {
                                var v21 bool = (!(("foo bar" + ("" + ("foo bar" + "\n\t"))) > "hello"))
                                _ = v21
                            } else {
                                v22 := GenS1{F0: 0}
                                _ = v22
                            }
                            fmt.Println(fi17, v19)
                            v23 := GenS2{F0: 1, F1: 1, F2: (fi17 % ((fi18) | 1))}
                            _ = v23
                        } else {
                            v19 += (v19 * v19)
                            var v24 int = fi17
                            if false {
                                v25 := GenS2{F0: 100, F1: 1, F2: fi18}
                                _ = v25
                            }
                            v24 += (^(fi17))
                            _ = v24
                        }
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    var v26 uint32 = 2
                    _ = v26
                }
                v19 += v19
                var v27 int64 = 226
                _ = fi18
                _ = v19
                _ = v27
            }
            var v28 uint16 = 42
            _ = fi17
            _ = v28
        }
    }
    _ = v5
    return 1
}
func genFn1(p0 int) bool {
    if (65535 <= 1) {
        var v29 string = "a"
        p0 = p0
        v29 = (("hello" + "\n\t") + v29)
        _ = v29
    } else {
        var v30 int16 = -1
        var v31 int64 = 1
        v32 := GenS3{F0: (v30 | (v30 | (v30 >> ((426) & 63)))), F1: p0, F2: 127}
        _ = v30
        _ = v31
        _ = v32
    }
    p0 = (p0 - p0)
    v33 := []int{(p0 + (p0 % (((^(p0))) | 1))), (p0 / ((p0) | 1)), (-(p0))}
    v34 := GenS1{F0: 995}
    p0 += (p0 | 127)
    _ = p0
    _ = v33
    _ = v34
    return (42 > -2147483648)
}

func main() {
    var v1 uint16 = 874
    v2 := []int{}
    v3 := make(map[string]string)
    var v4 float32 = 1e10
    v4 = v4
    v1 -= (^(v1))
    var v35 int32 = -2147483648
    var v36 uint64 = 149
    if false {
        v37 := GenS2{F0: 105, F1: v35, F2: 0}
        _ = v37
    }
    switch 1 {
    case 32768:
        v38 := v3["foo bar"]
        _ = v38
    }
    v39 := GenS1{F0: v1}
    var v40 bool = (4294967295 == uint32(v36))
    if len(v2) > 0 {
        v2[0] = 32767
    }
    v41 := GenS4{F0: false, F1: -111}
    v40 = ((-2.25 + (v4 * v4)) <= v4)
    fmt.Println(v35, v4, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v39
    _ = v4
    _ = v40
    _ = v41
}
