package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 int64; F2 bool }
type GenS3 struct { F0 int16; F1 uint32 }
func genFn0(p0 uint32, p1 int16, p2 uint64) uint16 {
    if (true || true) {
        v6 := GenS1{F0: -896}
        v7 := GenS1{F0: -9223372036854775808}
        var v8 int = 32767
        _ = v6
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 694
}
func genFn1(p0 uint32, p1 int16) int32 {
    v10 := []int{1, -516}
    var v11 int = -2147483648
    v12 := GenS2{F0: "", F1: -9223372036854775808, F2: true}
    fmt.Println(p1, p1, v11)
    var v13 uint16 = 255
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    return 0
}

func main() {
    var v1 uint16 = 65535
    v2 := make(map[string]uint32)
    v3 := []string{("foo bar" + "hello"), "", ("a" + "hello")}
    v4 := make(map[string]string)
    var v5 int64 = 0
    v1 -= v1
    var v14 uint32 = 550
    var v15 uint8 = 10
    var v16 string = ("" + ("hello" + ""))
    v16 = v16
    if (false || (false || true)) {
        v14 -= v14
        var v17 float64 = float64(-2.25)
        var v18 uint = 2
        _ = v17
        _ = v18
    }
    var v19 int8 = 100
    var v20 string = "hello"
    var v21 int8 = v19
    var v22 float64 = 0.0
    v19 = -1
    v19 = v19
    for fi23 := 0; fi23 < 18; fi23++ {
        if true {
            for fi24 := 0; fi24 < 5; fi24++ {
                if true {
                    switch fi23 {
                    case 4294967295:
                        var v25 int8 = (v19 * -1)
                        _ = v25
                    case 128:
                        v15 = (v15 & v15)
                    default:
                        var v26 uint8 = (^(v15))
                        _ = v26
                    }
                    var v27 int = 32768
                    if (v15 < (v15 << ((42) & 63))) {
                        v16 = v20
                        v28 := GenS2{F0: v20, F1: (v5 / ((v5) | 1)), F2: false}
                        var v29 float32 = 3.14159
                        _ = v28
                        _ = v29
                    }
                    _ = v27
                } else {
                    v16 = v20
                    var v30 uint64 = 1
                    if true {
                        if true {
                            _ = len(v4)
                        } else {
                            v31 := GenS3{F0: int16(v14), F1: (v14 | 657)}
                            v16 = "hello"
                            if false {
                                fmt.Println(fi24, v14, fi23)
                            }
                            _ = v31
                        }
                        var v32 uint64 = 50
                        var v33 bool = (!((v14 + (^(v14))) < v14))
                        _ = v32
                        _ = v33
                    } else {
                        if false {
                            var v34 uint8 = v15
                            fmt.Println(v1, v1, v30)
                            _ = v34
                        }
                    }
                    _ = v30
                }
                _ = fi24
            }
            v20 = v16
        }
        v20 = v20
        switch fi23 {
        case 127:
            v35 := GenS2{F0: ("a" + "\n\t"), F1: -812, F2: false}
            _ = v35
        case -1:
            var v36 int32 = -262
            _ = v36
        default:
            if len(v3) > 0 {
                v3[0] = v16
            }
        }
        _ = fi23
    }
    switch "\n\t" {
    case "a":
        var v37 float64 = float64((0.1 + 1e10))
        _ = v37
    default:
        for fi38 := 0; fi38 < 12; fi38++ {
            switch 128 {
            case 128:
                var v39 float64 = (v22 * (((-2.25 + v22) / 0.1) * v22))
                _ = v39
            case 1:
                fmt.Println(v14, v16, v14)
            case -1:
                var v40 int64 = -1
                _ = v40
            default:
                v14 += v14
            }
            _ = fi38
        }
    }
    var v41 int16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v41
    _ = v5
}
