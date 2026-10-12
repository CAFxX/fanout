package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 uint32; F1 string; F2 int8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64 }

func main() {
    var v1 int8 = -128
    var v2 int32 = -2147483648
    v3 := make(map[uint32]int)
    v4 := GenS1{F0: 2147483647, F1: -2147483648}
    for fi5 := 0; fi5 < 4; fi5++ {
        v6 := []int{-1}
        _ = fi5
        _ = v6
    }
    _ = len(v3)
    v7 := GenS1{F0: 887, F1: -864}
    v2 -= v2
    if true {
        var v8 int16 = 0
        _ = v8
    } else {
        if (42 >= 1) {
            switch "foo bar" {
            case "hello":
                switch "" {
                case "":
                    for fi9 := 0; fi9 < 10; fi9++ {
                        if true {
                            for fi10 := 0; fi10 < 8; fi10++ {
                                v2 = v2
                                v11 := []int{fi10}
                                if len(v11) > 0 {
                                    v11[0] = (^(fi10))
                                }
                                _ = fi10
                                _ = v11
                            }
                            var v12 float32 = 3.14159
                            v13 := GenS1{F0: 2147483647, F1: fi9}
                            _ = v12
                            _ = v13
                        }
                        _ = fi9
                    }
                case "a":
                    v1 += v1
                case "hello":
                    v1 = v1
                default:
                    var v14 uint = 4294967295
                    _ = v14
                }
            case "":
                var v15 int16 = 658
                _ = v15
            default:
                var v16 uint8 = 1
                _ = v16
            }
            fmt.Println(v2, v2)
            var v17 string = "\n\t"
            _ = v17
        } else {
            v18 := []int{32768, 42, 42}
            _ = v18
        }
        for fi19 := 0; fi19 < 3; fi19++ {
            fmt.Println(v1)
            if false {
                v20 := GenS2{F0: 42, F1: ("" + "hello"), F2: v1}
                v1 = 42
                _ = v20
            } else {
                var v21 uint64 = 42
                for fi22 := 0; fi22 < 2; fi22++ {
                    if (624 != uint32(v21)) {
                        switch -1 {
                        case 42:
                            delete(v3, 2147483647)
                        }
                    }
                    switch ("foo bar" + ((((("hello" + ("foo bar" + (("" + "\n\t") + ""))) + "a") + "hello") + "a") + "\n\t")) {
                    case "foo bar":
                        var v23 uint16 = 0
                        _ = v23
                    }
                    v21 = v21
                    v2 = (v2 ^ v2)
                    if ((!false) || true) {
                        var v24 int8 = -62
                        v24 = 56
                        _ = v24
                    } else {
                        if true {
                            if false {
                                v25 := GenS3{F0: v1}
                                switch (fi19 / ((fi19) | 1)) {
                                case 1:
                                    var v26 uint16 = 849
                                    _ = v26
                                case 4294967295:
                                    v3[607] = fi19
                                }
                                _ = v25
                            } else {
                                var v27 int = fi19
                                _ = v27
                            }
                            var v28 int32 = (v2 / (((v2 & -967)) | 1))
                            _ = v28
                        } else {
                            if false {
                                v29 := GenS1{F0: 4294967295, F1: -2147483648}
                                var v30 int64 = 9223372036854775807
                                fmt.Println(fi22, v1, v30)
                                if true {
                                    var v31 uint32 = 307
                                    _ = v31
                                }
                                _ = v29
                                _ = v30
                            } else {
                                if true {
                                    var v32 float32 = (0.0 * 3.14159)
                                    _ = v32
                                } else {
                                    fmt.Println(fi19, fi19, fi19)
                                    if (("a" + (("foo bar" + "") + (("a" + ("" + "\n\t")) + "foo bar"))) > ("foo bar" + "hello")) {
                                        v21 = v21
                                        v1 = (-(v1))
                                        switch (fi22 >> ((675) & 63)) {
                                        case 9223372036854775807:
                                            var v33 uint32 = 4294967295
                                            _ = v33
                                        case 127:
                                            if true {
                                                var v34 int32 = v2
                                                v35 := []int{}
                                                var v36 uint64 = v21
                                                _ = v34
                                                _ = v35
                                                _ = v36
                                            }
                                        }
                                        v37 := GenS2{F0: 2, F1: (((("\n\t" + "a") + "a") + "hello") + (("a" + "foo bar") + "a")), F2: (v1 % ((0) | 1))}
                                        _ = v37
                                    } else {
                                        delete(v3, 41)
                                    }
                                    v1 = v1
                                }
                                var v38 float32 = 0.1
                                v39 := []int{fi19, (-(fi22)), -80}
                                _ = v38
                                _ = v39
                            }
                            if (true && (true && true)) {
                                v40 := GenS3{F0: v1}
                                v41 := GenS1{F0: 4294967295, F1: (-(fi22))}
                                var v42 uint8 = 110
                                var v43 float64 = 0.1
                                _ = v40
                                _ = v41
                                _ = v42
                                _ = v43
                            }
                        }
                    }
                    _ = fi22
                }
                v44 := GenS4{F0: (v21 ^ v21)}
                _ = v21
                _ = v44
            }
            _ = fi19
        }
        var v45 uint = 42
        _ = v45
    }
    switch "a" {
    case "\n\t":
        var v46 uint8 = 254
        _ = v46
    }
    var v47 int = -1
    v48 := GenS4{F0: 1}
    if (!(0 == 42)) {
        switch v47 {
        case 9223372036854775807:
            _ = len(v3)
        default:
            var v49 uint16 = 529
            _ = v49
        }
        if (!(v1 >= 42)) {
            var v50 uint32 = 2
            var v51 int64 = 145
            _ = v50
            _ = v51
        } else {
            var v52 uint8 = 1
            v3[65535] = v47
            fmt.Println(v52, v2)
            _ = v52
        }
    }
    v1 += (v1 | (v1 >> ((18446744073709551615) & 63)))
    v2 -= 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v7
}
