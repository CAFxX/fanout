package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 uint16 }
type GenS2 struct { F0 uint8; F1 uint64; F2 int64 }
type GenS3 struct { F0 int32; F1 int16 }
type GenS4 struct { F0 string; F1 int; F2 int8 }
func genFn0(p0 uint, p1 int32, p2 int16) int8 {
    v4 := []int{0, 127, -1}
    if len(v4) > 0 {
        v4[0] = -380
    }
    p0 = (p0 - 65535)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -60
}
func genFn1(p0 string) uint16 {
    v5 := []int{-1, -785}
    if len(v5) > 0 {
        v5[0] = -9223372036854775808
    }
    var v6 int32 = 0
    fmt.Println(p0, v6)
    var v7 uint = 289
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 157
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]bool)
    v3 := []int{1}
    v8 := GenS1{F0: 1, F1: -128, F2: genFn1(("a" + (("a" + "\n\t") + "foo bar")))}
    var v9 bool = true
    switch "" {
    case "foo bar":
        for fi10 := 0; fi10 < 8; fi10++ {
            v9 = (0 < 338)
            for fi11 := 0; fi11 < 17; fi11++ {
                fmt.Println(fi11, fi10)
                switch ("" + "") {
                case "\n\t":
                    if (!true) {
                        fmt.Println(fi11, v9)
                        var v12 int = (^(fi10))
                        _ = v12
                    } else {
                        fmt.Println(fi10, fi11)
                    }
                }
                _ = fi11
            }
            _ = fi10
        }
    case "a":
        if v9 {
            if (!((1e10 / float64(-2.25)) >= 0.1)) {
                v9 = (0 < -1)
                var v13 float32 = 1e10
                _ = v13
            } else {
                if v9 {
                    var v14 uint8 = 1
                    var v15 string = ("\n\t" + "\n\t")
                    _ = v14
                    _ = v15
                } else {
                    if len(v3) > 0 {
                        v3[0] = -511
                    }
                    var v16 uint32 = 2147483647
                    _ = v16
                }
            }
        } else {
            v9 = v9
        }
    case "\n\t":
        for fi17 := 0; fi17 < 6; fi17++ {
            var v18 float64 = float64((-2.25 * 1e10))
            v19 := GenS1{F0: 2147483647, F1: -128, F2: genFn1(("hello" + (("foo bar" + "a") + "hello")))}
            var v20 bool = v9
            _ = fi17
            _ = v18
            _ = v19
            _ = v20
        }
    default:
        v9 = (!true)
    }
    v9 = true
    var v21 bool = false
    var v22 float32 = 1e10
    v21 = true
    if (674 <= 42) {
        for fi23 := 0; fi23 < 4; fi23++ {
            v9 = v21
            if (v21 && (127 == 206)) {
                var v24 float32 = (((-2.25 / v22) * 0.1) * v22)
                _ = v24
            }
            v9 = ("" >= "foo bar")
            _ = fi23
        }
        var v25 uint16 = 65535
        _ = v25
    }
    var v26 uint64 = 845
    var v27 bool = false
    for fi28 := 0; fi28 < 17; fi28++ {
        if (1 == 161) {
            v9 = (65535 > 107)
            if v9 {
                v9 = (float32((-2.25 - 3.14159)) >= -2.25)
                v29 := GenS2{F0: 78, F1: v26, F2: 9223372036854775807}
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = v29
            } else {
                var v30 int64 = 1
                switch fi28 {
                case 429:
                    var v31 bool = (2 >= 0)
                    _ = v31
                case 32767:
                    v26 = (v26 - 1)
                case 970:
                    switch (fi28 - fi28) {
                    case 42:
                        if v21 {
                            fmt.Println(v26, fi28, v27)
                            delete(v2, 2147483647)
                            v32 := GenS3{F0: 0, F1: -925}
                            _ = v32
                        } else {
                            v21 = true
                        }
                    case 4294967295:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                }
                switch fi28 {
                case 362:
                    var v33 bool = v21
                    _ = v33
                default:
                    v9 = (v22 >= (((0.0 + (v22 + v22)) - v22) + float32(3.14159)))
                }
                _ = v30
            }
        }
        _ = fi28
    }
    v21 = v21
    v27 = true
    var v34 uint32 = 0
    var v35 float64 = 1e10
    if v21 {
        v36 := GenS2{F0: 255, F1: (v26 * 1), F2: 0}
        v21 = ("foo bar" != "a")
        _ = v36
    } else {
        var v37 uint8 = 42
        delete(v2, v34)
        v38 := GenS4{F0: ((("\n\t" + "\n\t") + (("hello" + "hello") + ((("" + ((("a" + "\n\t") + "foo bar") + (("" + "foo bar") + "a"))) + ("" + "foo bar")) + ("" + "")))) + (("hello" + ("foo bar" + "")) + "foo bar")), F1: -990, F2: 100}
        _ = v37
        _ = v38
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v26
    _ = v27
    _ = v3
    _ = v34
    _ = v35
    _ = v8
    _ = v9
}
