package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool; F1 int }

func main() {
    v1 := make(map[uint32]int)
    var v2 float32 = float32(3.14159)
    v3 := make(map[int]string)
    v4 := make(map[string]bool)
    switch "\n\t" {
    case "":
        _ = len(v4)
    case "\n\t":
        if (false && true) {
            var v5 string = ("foo bar" + (("a" + ("" + "hello")) + ""))
            for fi6 := 0; fi6 < 17; fi6++ {
                v2 = ((v2 + (v2 * (float32(1.5) / 1e10))) + float32((((3.14159 + float64(float32(-2.25))) * 3.14159) - 1e10)))
                _ = fi6
            }
            _ = v5
        }
    }
    v2 -= v2
    for fi7 := 0; fi7 < 11; fi7++ {
        if false {
            var v8 uint64 = 42
            v9 := GenS1{F0: 127, F1: fi7}
            if ("\n\t" == "") {
                var v10 float32 = v2
                for fi11 := 0; fi11 < 11; fi11++ {
                    v2 = v2
                    switch (^(fi7)) {
                    case 4294967295:
                        if (!(!(false || true))) {
                            var v12 uint8 = 2
                            var v13 float32 = ((v2 - (v10 + v10)) * 0.0)
                            _ = v12
                            _ = v13
                        } else {
                            var v14 int16 = -1
                            var v15 int16 = -32768
                            _ = v14
                            _ = v15
                        }
                    case -509:
                        v16 := v1[2]
                        _ = v16
                    case 32767:
                        v17 := v4[("foo bar" + "\n\t")]
                        _ = v17
                    }
                    v18 := GenS1{F0: -9223372036854775808, F1: fi7}
                    switch "" {
                    case "hello":
                        v19 := []int{fi7, (-(fi7)), fi11, fi11}
                        _ = v19
                    case "\n\t":
                        if true {
                            v8 -= (^(v8))
                            var v20 uint = 959
                            _ = v20
                        } else {
                            switch fi11 {
                            case -2147483648:
                                if true {
                                    v21 := GenS1{F0: (^(fi11)), F1: 128}
                                    fmt.Println(v2)
                                    v8 = (^(v8))
                                    _ = v21
                                }
                            default:
                                var v22 uint16 = 213
                                _ = v22
                            }
                        }
                        v2 = ((v10 / float32(float64(v2))) + (1.5 + 1.5))
                        delete(v1, 65535)
                    case "a":
                        v8 = (v8 << ((936) & 63))
                    default:
                        if (16 >= 1) {
                            if (true && ((!((v8 % ((v8) | 1)) >= (v8 << ((65535) & 63)))) || false)) {
                                var v23 uint16 = 255
                                v10 = 0.1
                                var v24 uint = 42
                                var v25 int64 = -9223372036854775808
                                if true {
                                    v24 -= 750
                                } else {
                                    var v26 bool = ("a" <= "\n\t")
                                    switch "foo bar" {
                                    case "\n\t":
                                        v27 := GenS1{F0: (fi7 | fi7), F1: fi11}
                                        fmt.Println(fi11, v24)
                                        _ = v27
                                    case "foo bar":
                                        fmt.Println(fi7, v8)
                                    }
                                    fmt.Println(fi7, v26, v25)
                                    _ = v26
                                }
                                _ = v23
                                _ = v24
                                _ = v25
                            } else {
                                var v28 float32 = (v2 * 1.5)
                                fmt.Println(v2, fi7, v8)
                                v10 -= ((((v2 - ((((0.1 * float32(0.0)) * (float32((((((0.0 - 3.14159) - float64((0.0 / float32(float64(v28))))) * 0.0) + 0.0) - 1e10)) / 1.5)) * v10) + 3.14159)) / (0.0 * (v10 / (float32(3.14159) - 0.0)))) - 0.0) / (3.14159 * v10))
                                _ = v28
                            }
                        }
                    }
                    _ = fi11
                    _ = v18
                }
                _ = v10
            }
            _ = v8
            _ = v9
        }
        var v29 bool = false
        var v30 float32 = v2
        _ = fi7
        _ = v29
        _ = v30
    }
    var v31 float64 = 0.1
    for fi32 := 0; fi32 < 3; fi32++ {
        fmt.Println(v31)
        _ = fi32
    }
    var v33 int64 = -302
    v34 := []int{2147483647}
    var v35 float64 = 1.5
    var v36 uint32 = 225
    for fi37 := 0; fi37 < 2; fi37++ {
        for fi38 := 0; fi38 < 19; fi38++ {
            _ = len(v34)
            _ = fi38
        }
        if len(v34) > 0 {
            v34[0] = fi37
        }
        var v39 uint16 = 923
        _ = fi37
        _ = v39
    }
    var v40 uint64 = 73
    var v41 float32 = float32(float64(v2))
    v42 := GenS1{F0: 690, F1: -2147483648}
    if len(v34) >= 2 {
        v34 = v34[1:len(v34)]
    }
    for fi43 := 0; fi43 < 4; fi43++ {
        var v44 uint8 = 255
        v45 := GenS2{F0: ("foo bar" + "a")}
        v2 = -2.25
        _ = fi43
        _ = v44
        _ = v45
    }
    if true {
        var v46 uint32 = (v36 * v36)
        _ = len(v4)
        v47 := GenS3{F0: false, F1: 127}
        _ = v46
        _ = v47
    }
    var v48 int = 1
    v33 -= v33
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v48
}
