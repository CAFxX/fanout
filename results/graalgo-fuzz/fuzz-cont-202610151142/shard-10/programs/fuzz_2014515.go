package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 uint8; F2 int16 }
type GenS4 struct { F0 int; F1 int8; F2 uint32 }
func genFn0(p0 bool, p1 uint8) uint16 {
    v7 := GenS1{F0: 65535}
    if (4294967295 == 4294967295) {
        for fi8 := 0; fi8 < 3; fi8++ {
            var v9 uint8 = (p1 * (^(p1)))
            switch "\n\t" {
            case "":
                var v10 string = ("foo bar" + "foo bar")
                _ = v10
            case "hello":
                for fi11 := 0; fi11 < 19; fi11++ {
                    var v12 int16 = -1
                    v12 = (v12 - (v12 * (v12 - int16(fi8))))
                    v12 -= v12
                    _ = fi11
                    _ = v12
                }
            }
            v13 := []int{fi8, 12}
            _ = fi8
            _ = v13
            _ = v9
        }
        p1 = (^(p1))
        var v14 uint32 = 593
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v7
    return 290
}

func main() {
    v1 := make(map[string]bool)
    v2 := []uint32{2147483648, 431}
    v3 := make(map[string]bool)
    v4 := make(map[string]uint32)
    var v5 uint32 = 0
    v6 := make(map[string]string)
    switch "hello" {
    case "foo bar":
        v15 := GenS2{F0: 356}
        _ = v15
    default:
        for fi16 := 0; fi16 < 6; fi16++ {
            for fi17 := 0; fi17 < 14; fi17++ {
                v18 := GenS3{F0: 85, F1: 42, F2: 1000}
                _ = fi17
                _ = v18
            }
            v5 -= v5
            var v19 uint32 = (v5 * (v5 & v5))
            _ = fi16
            _ = v19
        }
    }
    v5 -= 2147483648
    switch (("" + "hello") + "hello") {
    case "":
        if (174 <= 250) {
            if false {
                var v20 int8 = -35
                _ = v20
            } else {
                fmt.Println(v5)
                v5 -= 0
            }
            for fi21 := 0; fi21 < 6; fi21++ {
                v5 = 42
                if true {
                    var v22 float32 = 0.1
                    _ = v22
                } else {
                    if true {
                        v5 = 253
                        for fi23 := 0; fi23 < 13; fi23++ {
                            switch ("\n\t" + "foo bar") {
                            case "foo bar":
                                var v24 float32 = 0.1
                                _ = v24
                            case "\n\t":
                                if false {
                                    if (798 > 42) {
                                        if true {
                                            v25 := v6[((((("hello" + "a") + (("\n\t" + "") + "")) + ("a" + "a")) + "\n\t") + "a")]
                                            v25 = "foo bar"
                                            _ = v25
                                        } else {
                                            var v26 uint64 = 42
                                            _ = v26
                                        }
                                        var v27 uint = 2
                                        if len(v2) > 0 {
                                            v2[0] = v5
                                        }
                                        switch ("\n\t" + "\n\t") {
                                        case "":
                                            var v28 float64 = ((0.0 / 0.1) * 1e10)
                                            _ = v28
                                        case "\n\t":
                                            v29 := GenS2{F0: v27}
                                            _ = v29
                                        default:
                                            var v30 int8 = -128
                                            _ = v30
                                        }
                                        _ = v27
                                    } else {
                                        v5 = 1
                                    }
                                }
                            default:
                                v5 += 4294967295
                            }
                            _ = fi23
                        }
                    } else {
                        var v31 float32 = float32(1e10)
                        _ = v31
                    }
                    fmt.Println(fi21, fi21)
                    v32 := GenS1{F0: 42}
                    _ = v32
                }
                var v33 int32 = 1
                _ = fi21
                _ = v33
            }
            v34 := GenS2{F0: 116}
            _ = v34
        }
    case "foo bar":
        fmt.Println(v5)
    case "a":
        fmt.Println(v5)
    }
    var v35 uint64 = 649
    v5 = (v5 ^ 2147483648)
    v36 := GenS3{F0: 2147483647, F1: 1, F2: 32767}
    v37 := GenS1{F0: (^(v5))}
    var v38 string = "hello"
    v38 = "foo bar"
    v39 := GenS2{F0: 2}
    v5 += 47
    var v40 uint = 42
    var v41 uint64 = (v35 >> (((v40 << ((4294967295) & 63))) & 63))
    if false {
        v35 = v35
        var v42 int16 = -444
        var v43 uint32 = v5
        _ = v42
        _ = v43
    } else {
        var v44 int32 = -538
        v45 := v6[((v38 + v38) + (v38 + "\n\t"))]
        _ = v44
        _ = v45
    }
    v46 := GenS4{F0: 0, F1: -128, F2: 4294967295}
    var v47 int = 2147483647
    fmt.Println(v40, v38, v40)
    for fi48 := 0; fi48 < 1; fi48++ {
        v40 = v40
        if (-1 <= 2147483647) {
            switch v38 {
            case "foo bar":
                v41 = (v35 ^ v35)
            }
        }
        v49 := v1[""]
        _ = fi48
        _ = v49
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}
