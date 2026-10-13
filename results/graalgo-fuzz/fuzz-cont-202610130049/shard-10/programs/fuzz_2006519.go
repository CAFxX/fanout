package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64; F1 uint16; F2 int32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int32; F1 int16; F2 int64 }

func main() {
    var v1 int64 = 1
    var v2 float32 = float32(-2.25)
    var v3 uint = 815
    v4 := make(map[uint32]string)
    var v5 bool = (("\n\t" >= ("a" + "\n\t")) && false)
    var v6 int8 = -102
    if v5 {
        for fi7 := 0; fi7 < 10; fi7++ {
            var v8 uint = 2
            for fi9 := 0; fi9 < 8; fi9++ {
                switch ("a" + "hello") {
                case "\n\t":
                    switch fi7 {
                    case 0:
                        v10 := GenS1{F0: v3}
                        _ = v10
                    case 638:
                        v11 := []int{fi9, fi7, (fi7 % ((fi9) | 1)), (fi7 * fi9)}
                        _ = v11
                    case 1:
                        v12 := []int{}
                        _ = v12
                    }
                case "a":
                    fmt.Println(v1, fi7, v5)
                case "":
                    if (("a" > "\n\t") && ((v8 | (^(v3))) > (v8 ^ (v8 & 18446744073709551615)))) {
                        v13 := GenS2{F0: (32767 != -1)}
                        v14 := []int{}
                        v3 = v8
                        var v15 bool = v5
                        if (true || false) {
                            v16 := GenS3{F0: 4294967295, F1: 32767, F2: -241}
                            _ = v16
                        } else {
                            var v17 uint64 = 2
                            if false {
                                fmt.Println(v17)
                                _ = len(v14)
                            } else {
                                v14 = append(v14, fi9)
                            }
                            v5 = ((^(v17)) == 18446744073709551615)
                            if (0 == -1) {
                                v4[496] = ("\n\t" + "")
                            } else {
                                v18 := GenS1{F0: v8}
                                if v15 {
                                    var v19 int = (fi7 ^ fi9)
                                    fmt.Println(v3, v8)
                                    _ = v19
                                } else {
                                    var v20 uint64 = 4294967295
                                    _ = v20
                                }
                                _ = v18
                            }
                            var v21 float32 = 1.5
                            var v22 uint32 = 2147483647
                            _ = v17
                            _ = v21
                            _ = v22
                        }
                        v23 := GenS3{F0: 4294967295, F1: 42, F2: 335}
                        _ = v13
                        _ = v14
                        _ = v15
                        _ = v23
                    } else {
                        switch (fi9 * (fi9 * (fi7 - (fi7 ^ fi9)))) {
                        case -1:
                            var v24 uint8 = 50
                            _ = v24
                        case -50:
                            v25 := GenS4{F0: 4294967295}
                            _ = v25
                        }
                    }
                }
                v26 := GenS4{F0: 2147483647}
                _ = fi9
                _ = v26
            }
            _ = fi7
            _ = v8
        }
        v1 += (v1 & (^(v1)))
        fmt.Println(v6)
    }
    for fi27 := 0; fi27 < 10; fi27++ {
        delete(v4, 4294967295)
        var v28 uint32 = 42
        _ = fi27
        _ = v28
    }
    switch 32767 {
    case 9223372036854775807:
        v6 += (v6 >> (((v3 << (((v3 + 65535)) & 63))) & 63))
    }
    for fi29 := 0; fi29 < 13; fi29++ {
        var v30 bool = v5
        v6 += v6
        _ = fi29
        _ = v30
    }
    v31 := GenS5{F0: 0, F1: 0, F2: v1}
    var v32 uint = (v3 << ((v3) & 63))
    for fi33 := 0; fi33 < 13; fi33++ {
        for fi34 := 0; fi34 < 18; fi34++ {
            v35 := GenS2{F0: v5}
            var v36 int16 = 0
            v6 = (^(v6))
            v37 := GenS2{F0: (!v5)}
            _ = fi34
            _ = v35
            _ = v36
            _ = v37
        }
        var v38 float32 = v2
        _ = fi33
        _ = v38
    }
    v4[1] = ("a" + ("foo bar" + ("a" + "\n\t")))
    var v39 int64 = (v1 >> ((0) & 63))
    if (!(32767 != 32767)) {
        if v5 {
            if ((0.1 + (v2 * float32(1e10))) > 0.1) {
                v1 += (v1 ^ v1)
                fmt.Println(v6, v5, v39)
                fmt.Println(v2, v39)
            } else {
                v39 += -1
                fmt.Println(v2, v6)
            }
        }
    } else {
        for fi40 := 0; fi40 < 14; fi40++ {
            v1 += (v39 | v1)
            for fi41 := 0; fi41 < 16; fi41++ {
                var v42 int16 = 1000
                _ = fi41
                _ = v42
            }
            _ = fi40
        }
    }
    v43 := []int{2147483647, -745, -616}
    v1 = 9223372036854775807
    fmt.Println(v6)
    v44 := GenS3{F0: 1, F1: 1, F2: -450}
    if (v6 > (^(v6))) {
        v45 := GenS4{F0: 2147483648}
        switch "" {
        case "hello":
            var v46 string = ""
            _ = v46
        case "a":
            fmt.Println(v32, v39)
        case "foo bar":
            v47 := GenS3{F0: 819, F1: 2, F2: -912}
            _ = v47
        default:
            v48 := GenS4{F0: 0}
            _ = v48
        }
        v49 := GenS5{F0: 750, F1: 32767, F2: v1}
        _ = v45
        _ = v49
    }
    var v50 int16 = 566
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v39
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v50
    _ = v6
}
