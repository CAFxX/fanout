package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 string; F1 int; F2 int16 }
type GenS3 struct { F0 uint16; F1 uint8; F2 uint8 }
func genFn0(p0 int64, p1 bool) uint32 {
    p1 = (131 >= 1)
    if true {
        v5 := GenS1{F0: 2, F1: 0}
        v6 := []int{}
        _ = v5
        _ = v6
    }
    v7 := []int{-205, -349}
    var v8 int = 349
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 2147483647
}
func genFn1(p0 bool, p1 int, p2 int64) uint8 {
    if (("hello" + ("foo bar" + "foo bar")) != "\n\t") {
        for fi9 := 0; fi9 < 10; fi9++ {
            if (265 > 65535) {
                for fi10 := 0; fi10 < 3; fi10++ {
                    fmt.Println(fi10, fi9, p1)
                    fmt.Println(p2, p1)
                    p1 -= p1
                    _ = fi10
                }
            } else {
                switch "hello" {
                case "hello":
                    fmt.Println(p2, p0)
                default:
                    for fi11 := 0; fi11 < 16; fi11++ {
                        var v12 bool = (("\n\t" + ("" + "")) >= ("\n\t" + "\n\t"))
                        var v13 int = (fi9 * (p1 ^ fi11))
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                }
            }
            v14 := []int{(fi9 / ((-9223372036854775808) | 1)), (fi9 % (((p1 + 42)) | 1)), fi9, 128}
            var v15 int16 = 1000
            _ = fi9
            _ = v14
            _ = v15
        }
        for fi16 := 0; fi16 < 6; fi16++ {
            for fi17 := 0; fi17 < 14; fi17++ {
                if p0 {
                    if (p1 == -1) {
                        var v18 int64 = p2
                        if false {
                            v19 := []int{(p1 & fi16), (p1 & (p1 + (fi17 & fi17))), p1, -1}
                            _ = v19
                        } else {
                            if (0.1 > (((1.5 / 3.14159) / 3.14159) / 1.5)) {
                                var v20 uint64 = 322
                                _ = v20
                            }
                            p0 = (4294967295 > 42)
                            var v21 string = (("foo bar" + "\n\t") + "")
                            _ = v21
                        }
                        _ = v18
                    }
                    p0 = (genFn0((p2 >> ((2) & 63)), (true || p0)) < 2147483647)
                    var v22 float64 = float64(((-2.25 + (float32(1.5) * (((1.5 + (0.0 * float32(-2.25))) + float32(1.5)) - (3.14159 - 1.5)))) * ((0.0 - 1e10) / float32(-2.25))))
                    _ = v22
                }
                var v23 bool = (!("\n\t" < ""))
                _ = fi17
                _ = v23
            }
            _ = fi16
        }
    } else {
        v24 := GenS2{F0: "", F1: (^(p1)), F2: int16(p2)}
        _ = v24
    }
    v25 := GenS1{F0: 18446744073709551615, F1: 42}
    v26 := []int{9223372036854775807}
    var v27 uint8 = 128
    var v28 uint = 905
    _ = p0
    _ = p1
    _ = p2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    return v27
}

func main() {
    v1 := []uint32{}
    v2 := []uint32{65535, 512, 2147483648}
    var v3 bool = (4294967295 != 18446744073709551615)
    v4 := make(map[int]int)
    _ = len(v4)
    for fi29 := 0; fi29 < 6; fi29++ {
        for fi30 := 0; fi30 < 16; fi30++ {
            var v31 int64 = -915
            var v32 uint64 = 592
            fmt.Println(v31, v3, v32)
            _ = fi30
            _ = v31
            _ = v32
        }
        if (0 > 143) {
            for fi33 := 0; fi33 < 7; fi33++ {
                if true {
                    var v34 uint32 = 2147483647
                    v35 := GenS3{F0: 42, F1: 0, F2: 0}
                    _ = v34
                    _ = v35
                } else {
                    fmt.Println(fi33, fi33, fi33)
                    var v36 int32 = 249
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v36
                }
                _ = fi33
            }
        }
        _ = fi29
    }
    var v37 uint = 42
    for fi38 := 0; fi38 < 11; fi38++ {
        v39 := GenS2{F0: "a", F1: fi38, F2: 975}
        if (v3 && v3) {
            var v40 int32 = -2147483648
            var v41 uint16 = 1
            _ = v40
            _ = v41
        } else {
            for fi42 := 0; fi42 < 14; fi42++ {
                switch (fi38 | 1) {
                case 42:
                    if ("" >= (("hello" + "\n\t") + (("\n\t" + "") + "a"))) {
                        _ = len(v4)
                        v37 = v37
                    } else {
                        fmt.Println(v3, fi38)
                    }
                case 9223372036854775807:
                    if false {
                        v43 := v4[(fi42 ^ -141)]
                        fmt.Println(v37, fi38)
                        _ = v43
                    }
                case 4294967295:
                    if (true || v3) {
                        v44 := v4[(fi42 % ((fi38) | 1))]
                        _ = v44
                    }
                }
                _ = fi42
            }
            var v45 uint32 = 1
            _ = v45
        }
        _ = fi38
        _ = v39
    }
    v3 = v3
    fmt.Println(v37, v37)
    v46 := GenS3{F0: 114, F1: 2, F2: 250}
    var v47 uint64 = 18446744073709551615
    var v48 int16 = 32767
    switch int(v47) {
    case 0:
        if v3 {
            switch 42 {
            case 572:
                if v3 {
                    v3 = (uint8(v37) < 255)
                } else {
                    switch -2147483648 {
                    case 1:
                        if (179 < 199) {
                            fmt.Println(v3)
                            fmt.Println(v37)
                        } else {
                            v2 = append(v2, uint32(v47))
                            if true {
                                v49 := GenS3{F0: 255, F1: 1, F2: 128}
                                _ = v49
                            }
                            fmt.Println(v37)
                        }
                    case 32768:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                }
            case 9223372036854775807:
                if (true && v3) {
                    var v50 int64 = 42
                    _ = v50
                }
            case 343:
                var v51 uint32 = 7
                _ = v51
            default:
                var v52 int32 = -2147483648
                _ = v52
            }
            v53 := GenS2{F0: "foo bar", F1: 567, F2: (v48 & v48)}
            if (("foo bar" + ("foo bar" + ("foo bar" + ""))) >= "a") {
                if v3 {
                    v3 = false
                } else {
                    for fi54 := 0; fi54 < 5; fi54++ {
                        for fi55 := 0; fi55 < 15; fi55++ {
                            var v56 float64 = (3.14159 / 1e10)
                            v57 := v4[fi55]
                            fmt.Println(v57)
                            _ = fi55
                            _ = v56
                            _ = v57
                        }
                        var v58 uint = v37
                        _ = fi54
                        _ = v58
                    }
                }
            }
            _ = v53
        } else {
            fmt.Println(v37)
        }
    case 9223372036854775807:
        fmt.Println(v48, v3)
    default:
        if false {
            var v59 uint16 = 32767
            _ = v59
        }
    }
    for fi60 := 0; fi60 < 6; fi60++ {
        fmt.Println(v48, v37, v3)
        _ = len(v1)
        var v61 int8 = 1
        _ = fi60
        _ = v61
    }
    _ = len(v4)
    fmt.Println(v3, v47)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v46
    _ = v47
    _ = v48
}
