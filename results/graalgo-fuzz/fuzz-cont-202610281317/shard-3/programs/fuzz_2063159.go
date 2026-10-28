package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 int32 }
type GenS2 struct { F0 uint64; F1 string; F2 string }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 bool; F1 int16 }
func genFn0(p0 uint64, p1 uint64) int {
    v4 := GenS1{F0: 616, F1: 42, F2: 42}
    if true {
        var v5 uint64 = 241
        _ = v5
    } else {
        p0 = 42
    }
    for fi6 := 0; fi6 < 11; fi6++ {
        if (true || (259 != 9223372036854775807)) {
            if (!true) {
                if true {
                    for fi7 := 0; fi7 < 2; fi7++ {
                        fmt.Println(fi6, p0, fi7)
                        _ = fi7
                    }
                } else {
                    if false {
                        var v8 bool = false
                        var v9 uint32 = 65535
                        _ = v8
                        _ = v9
                    }
                    switch 9223372036854775807 {
                    case 81:
                        v10 := GenS1{F0: -9223372036854775808, F1: 1, F2: 42}
                        _ = v10
                    case 32768:
                        v11 := GenS1{F0: 42, F1: 777, F2: 42}
                        _ = v11
                    case -279:
                        p0 += p1
                    default:
                        p0 = p1
                    }
                    var v12 uint16 = 2
                    _ = v12
                }
                fmt.Println(fi6, p1)
                v13 := []int{fi6}
                _ = v13
            } else {
                p1 = (p0 | p1)
            }
            p1 = (p0 % (((p0 >> ((1) & 63))) | 1))
            if false {
                v14 := []int{9223372036854775807}
                var v15 int32 = 0
                fmt.Println(v15)
                _ = v14
                _ = v15
            }
        } else {
            var v16 int = (^(fi6))
            fmt.Println(fi6, v16, fi6)
            if false {
                p1 += (^(p1))
                var v17 string = (("a" + (((("hello" + "hello") + "foo bar") + ("a" + ("hello" + "a"))) + "\n\t")) + ("foo bar" + ""))
                _ = v17
            }
            _ = v16
        }
        for fi18 := 0; fi18 < 19; fi18++ {
            if false {
                fmt.Println(p0, fi6)
                var v19 string = "a"
                _ = v19
            }
            _ = fi18
        }
        _ = fi6
    }
    v20 := []int{9223372036854775807, 0}
    _ = p0
    _ = p1
    _ = v20
    _ = v4
    return 4294967295
}

func main() {
    var v1 int32 = 0
    v2 := []uint32{384, 0}
    v3 := []int{2147483647, 32767}
    v21 := GenS2{F0: 1, F1: "foo bar", F2: "\n\t"}
    var v22 int32 = (v1 - 1)
    var v23 string = "hello"
    var v24 int = -450
    var v25 uint64 = 1
    if (!false) {
        for fi26 := 0; fi26 < 16; fi26++ {
            var v27 uint32 = 822
            if false {
                if false {
                    for fi28 := 0; fi28 < 20; fi28++ {
                        var v29 int64 = 1
                        _ = fi28
                        _ = v29
                    }
                    switch fi26 {
                    case 4294967295:
                        if (false || false) {
                            var v30 float64 = (3.14159 + 1e10)
                            v31 := GenS2{F0: v25, F1: "", F2: (v23 + v23)}
                            _ = v30
                            _ = v31
                        }
                    default:
                        v1 -= v22
                    }
                    var v32 uint16 = 1
                    _ = v32
                }
                if false {
                    if true {
                        var v33 int64 = 0
                        v27 = 338
                        _ = v33
                    }
                    if (!((v22 >> ((uint(v27)) & 63)) > v1)) {
                        for fi34 := 0; fi34 < 6; fi34++ {
                            var v35 uint8 = 128
                            _ = fi34
                            _ = v35
                        }
                        var v36 uint16 = 293
                        _ = v36
                    } else {
                        _ = len(v3)
                        v37 := GenS3{F0: 541, F1: 42}
                        v38 := GenS2{F0: v25, F1: v23, F2: v23}
                        _ = v37
                        _ = v38
                    }
                }
            } else {
                var v39 int32 = -973
                _ = v39
            }
            v2 = append(v2, (v27 % (((^(v27))) | 1)))
            _ = fi26
            _ = v27
        }
        for fi40 := 0; fi40 < 2; fi40++ {
            for fi41 := 0; fi41 < 3; fi41++ {
                v1 = (v1 | 206)
                _ = fi41
            }
            _ = fi40
        }
        v22 -= v22
    } else {
        v22 += 606
        if len(v2) > 0 {
            v2[0] = 42
        }
    }
    var v42 uint8 = 142
    fmt.Println(v25)
    v43 := GenS4{F0: false, F1: int16(v42)}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v42
    _ = v43
}
