package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint; F1 int }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int8 }
func genFn0(p0 int) string {
    switch p0 {
    case 1:
        switch "" {
        case "\n\t":
            var v6 int64 = -322
            _ = v6
        case "foo bar":
            if (!(!(!false))) {
                for fi7 := 0; fi7 < 10; fi7++ {
                    p0 += (^(p0))
                    var v8 bool = (0 == 42)
                    _ = fi7
                    _ = v8
                }
            }
        }
    case 0:
        var v9 int16 = 402
        _ = v9
    case -906:
        for fi10 := 0; fi10 < 11; fi10++ {
            p0 = 9223372036854775807
            if true {
                for fi11 := 0; fi11 < 16; fi11++ {
                    v12 := []int{(fi11 >> ((401) & 63)), (fi11 - p0), (fi10 / (((fi10 | 42)) | 1))}
                    switch -9223372036854775808 {
                    case 9223372036854775807:
                        p0 = 4294967295
                    case -770:
                        var v13 bool = false
                        var v14 int8 = 42
                        _ = len(v12)
                        _ = v13
                        _ = v14
                    case 1:
                        v15 := GenS1{F0: 1, F1: 0, F2: "hello"}
                        _ = v15
                    default:
                        p0 = (^(fi11))
                    }
                    p0 = (fi11 / (((fi10 + fi10)) | 1))
                    _ = fi11
                    _ = v12
                }
                p0 -= 765
            } else {
                v16 := []int{(p0 % (((fi10 * p0)) | 1)), fi10, 832}
                if (1 > 1) {
                    p0 -= p0
                    var v17 int8 = 127
                    _ = v17
                } else {
                    v18 := GenS2{F0: 451}
                    if len(v16) > 0 {
                        v16[0] = 32767
                    }
                    p0 -= fi10
                    _ = v18
                }
                if false {
                    p0 -= (-(p0))
                } else {
                    for fi19 := 0; fi19 < 15; fi19++ {
                        var v20 bool = (-509 != p0)
                        fmt.Println(fi19)
                        _ = fi19
                        _ = v20
                    }
                    fmt.Println(p0, p0)
                }
                _ = v16
            }
            _ = fi10
        }
    }
    var v21 string = "\n\t"
    var v22 string = v21
    var v23 float64 = -2.25
    _ = p0
    _ = v21
    _ = v22
    _ = v23
    return (((v22 + "hello") + "a") + "foo bar")
}

func main() {
    var v1 uint32 = 744
    var v2 int = 458
    var v3 int32 = int32(v2)
    v4 := make(map[uint32]int)
    v5 := []uint32{208, v1, 2, v1}
    if ((!(((v3 + v3) != (^(v3))) || false)) && false) {
        if false {
            for fi24 := 0; fi24 < 15; fi24++ {
                var v25 int = (-(fi24))
                v25 = fi24
                _ = fi24
                _ = v25
            }
        }
    }
    v3 = -866
    if false {
        var v26 uint64 = 1
        var v27 int = v2
        _ = v26
        _ = v27
    }
    v28 := GenS1{F0: 1, F1: 1, F2: ""}
    var v29 int8 = 1
    if true {
        v2 = (v2 | v2)
    } else {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        var v30 string = "\n\t"
        v31 := GenS3{F0: 4294967295, F1: (-(v2))}
        _ = v30
        _ = v31
    }
    if (float64(1e10) == -2.25) {
        delete(v4, (^(v1)))
        v32 := GenS4{F0: -9223372036854775808}
        v3 = 0
        _ = v32
    }
    v33 := GenS2{F0: 497}
    fmt.Println(v2)
    switch 4294967295 {
    case -1:
        for fi34 := 0; fi34 < 16; fi34++ {
            v35 := GenS5{F0: (v29 << ((42) & 63))}
            _ = fi34
            _ = v35
        }
    case -42:
        var v36 uint16 = 195
        _ = v36
    }
    if len(v5) > 0 {
        v5[0] = (v1 * v1)
    }
    var v37 float32 = 3.14159
    v38 := GenS2{F0: 1}
    _ = len(v4)
    if false {
        v1 = 2
        v39 := GenS5{F0: 1}
        var v40 float32 = 0.0
        _ = v39
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v33
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
