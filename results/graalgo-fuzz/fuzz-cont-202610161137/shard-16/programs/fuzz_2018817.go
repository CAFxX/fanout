package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 int16 }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint64; F1 uint32 }

func main() {
    v1 := make(map[string]bool)
    v2 := []uint32{2147483648, 65535, 42, 4294967295}
    v3 := []uint32{292, 65535, 65535}
    v4 := make(map[uint32]bool)
    for fi5 := 0; fi5 < 14; fi5++ {
        var v6 int64 = 9223372036854775807
        var v7 int32 = -2147483648
        var v8 uint = 18446744073709551615
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    var v9 uint8 = 255
    var v10 uint32 = 2
    switch (("a" + "foo bar") + "a") {
    case "\n\t":
        for fi11 := 0; fi11 < 19; fi11++ {
            var v12 float32 = (1e10 * float32(1e10))
            _ = fi11
            _ = v12
        }
    }
    v13 := GenS1{F0: 65535, F1: 42, F2: -1}
    var v14 int = 0
    if ((true || false) || false) {
        for fi15 := 0; fi15 < 9; fi15++ {
            switch fi15 {
            case 9223372036854775807:
                v16 := v4[v10]
                _ = v16
            case -2147483648:
                for fi17 := 0; fi17 < 20; fi17++ {
                    if (!(((v14 + (^(fi15))) == v14) || true)) {
                        if true {
                            if true {
                                v4[(^(v10))] = ((!true) || false)
                            } else {
                                v14 = fi17
                            }
                            var v18 int32 = 445
                            _ = v18
                        } else {
                            if (1 == 671) {
                                fmt.Println(fi17, v9, v14)
                            } else {
                                fmt.Println(fi17, v9)
                                v19 := GenS2{F0: 0, F1: 2}
                                var v20 float64 = 0.1
                                v21 := GenS2{F0: 827, F1: 367}
                                var v22 int16 = -1
                                v3 = append(v3, 1)
                                _ = v19
                                _ = v20
                                _ = v21
                                _ = v22
                            }
                            var v23 uint16 = 23
                            var v24 uint8 = v9
                            _ = v23
                            _ = v24
                        }
                        var v25 uint64 = 879
                        if (uint64(v9) < (v25 << ((355) & 63))) {
                            var v26 float32 = 3.14159
                            switch v14 {
                            case 42:
                                var v27 float64 = 3.14159
                                _ = v27
                            }
                            v4[(v10 % (((v10 - 2)) | 1))] = (-9223372036854775808 >= 0)
                            v10 += v10
                            if (1.5 > ((((v26 - (v26 - v26)) / 1.0) * v26) - 0.1)) {
                                v26 = v26
                            }
                            _ = v26
                        }
                        _ = v25
                    } else {
                        v9 -= v9
                        v9 += (v9 ^ (v9 & 234))
                        if true {
                            var v28 uint32 = (v10 + 4294967295)
                            _ = len(v4)
                            _ = v28
                        }
                        if true {
                            fmt.Println(v14, v10)
                        }
                    }
                    _ = fi17
                }
            }
            if true {
                var v29 uint32 = (^(v10))
                var v30 int32 = -745
                v31 := GenS1{F0: 0, F1: 100, F2: 42}
                _ = v29
                _ = v30
                _ = v31
            } else {
                fmt.Println(v14)
                for fi32 := 0; fi32 < 10; fi32++ {
                    v14 = (^(fi15))
                    delete(v1, ("a" + "foo bar"))
                    _ = fi32
                }
            }
            _ = fi15
        }
    } else {
        v9 += 2
    }
    v33 := GenS2{F0: 42, F1: 622}
    v3 = append(v3, v10)
    v34 := v4[v10]
    if false {
        v14 = 376
    } else {
        for fi35 := 0; fi35 < 9; fi35++ {
            var v36 uint64 = 449
            _ = fi35
            _ = v36
        }
        var v37 string = ("hello" + "foo bar")
        v34 = (!v34)
        _ = v37
    }
    v34 = (154 > 119)
    fmt.Println(v14, v10, v14)
    if false {
        for fi38 := 0; fi38 < 3; fi38++ {
            v39 := GenS3{F0: v14}
            var v40 uint8 = (v9 * 127)
            _ = fi38
            _ = v39
            _ = v40
        }
        v41 := v4[(^(v10))]
        _ = v41
    }
    for fi42 := 0; fi42 < 16; fi42++ {
        switch fi42 {
        case 460:
            switch "a" {
            case "a":
                switch (("foo bar" + ((((("\n\t" + "a") + "") + "a") + "\n\t") + "\n\t")) + "\n\t") {
                case "foo bar":
                    var v43 uint8 = v9
                    _ = v43
                case "a":
                    switch (fi42 + fi42) {
                    case -2147483648:
                        delete(v1, "hello")
                    case 1:
                        for fi44 := 0; fi44 < 3; fi44++ {
                            delete(v4, 498)
                            if v34 {
                                var v45 string = "foo bar"
                                _ = v45
                            }
                            _ = fi44
                        }
                    default:
                        v9 -= 127
                    }
                }
            case "":
                var v46 uint32 = v10
                _ = v46
            case "\n\t":
                var v47 float64 = 0.1
                _ = v47
            default:
                if (1 > 942) {
                    fmt.Println(fi42, fi42)
                }
            }
        case 9223372036854775807:
            for fi48 := 0; fi48 < 2; fi48++ {
                v1[(("foo bar" + "") + "a")] = (1 == 127)
                if ((!true) && v34) {
                    v9 += 127
                    v34 = (v34 && ((v9 / (((v9 % ((v9) | 1))) | 1)) < (v9 << ((1) & 63))))
                }
                v34 = (false && (!false))
                _ = fi48
            }
        }
        _ = fi42
    }
    for fi49 := 0; fi49 < 7; fi49++ {
        var v50 int8 = 1
        v51 := GenS4{F0: 2, F1: v10}
        _ = fi49
        _ = v50
        _ = v51
    }
    v9 += v9
    if v34 {
        if len(v3) > 0 {
            v3[0] = v10
        }
        v34 = (-127 <= 100)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v9
}
