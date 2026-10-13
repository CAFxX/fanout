package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int32 }
type GenS2 struct { F0 int64; F1 int; F2 uint16 }
type GenS3 struct { F0 int; F1 uint16; F2 int64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint8 }
func genFn0(p0 uint32) uint64 {
    if false {
        v7 := []int{167, -920, 1}
        _ = v7
    }
    v8 := GenS1{F0: 1, F1: "foo bar", F2: 42}
    p0 -= p0
    var v9 uint64 = 2
    _ = p0
    _ = v8
    _ = v9
    return 718
}

func main() {
    var v1 int64 = 1
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]string)
    v4 := make(map[string]bool)
    v5 := make(map[int]int)
    v6 := make(map[string]int)
    v10 := GenS1{F0: 0, F1: "foo bar", F2: -2147483648}
    var v11 int = 4294967295
    v1 = v1
    var v12 int16 = 0
    for fi13 := 0; fi13 < 6; fi13++ {
        for fi14 := 0; fi14 < 13; fi14++ {
            var v15 uint16 = 701
            if (42 < 65535) {
                var v16 int8 = 42
                if (351 >= 345) {
                    if (false || ((2 > 18446744073709551615) || true)) {
                        if true {
                            if (255 <= 1) {
                                v15 = 1
                            }
                        } else {
                            if false {
                                switch ("foo bar" + "a") {
                                case "hello":
                                    switch "" {
                                    case "hello":
                                        v17 := []int{}
                                        _ = v17
                                    case "a":
                                        var v18 int64 = v1
                                        _ = v18
                                    case "\n\t":
                                        v1 = v1
                                    }
                                case "":
                                    var v19 bool = false
                                    _ = v19
                                case "a":
                                    v20 := []int{int(v12), v11}
                                    v12 = 1
                                    fmt.Println(v16)
                                    _ = v20
                                default:
                                    var v21 int16 = (v12 - (v12 ^ v12))
                                    _ = v21
                                }
                                v22 := GenS1{F0: v12, F1: "\n\t", F2: -1}
                                fmt.Println(v16, v15, v16)
                                _ = v22
                            } else {
                                var v23 uint = 18446744073709551615
                                _ = v23
                            }
                        }
                        if ((v16 - 119) < v16) {
                            var v24 bool = false
                            switch "" {
                            case "hello":
                                _ = len(v2)
                            case "a":
                                v25 := v6[""]
                                _ = v25
                            }
                            v1 += 9223372036854775807
                            _ = v24
                        } else {
                            _ = len(v3)
                            v26 := GenS1{F0: (v12 ^ (-(v12))), F1: "\n\t", F2: 551}
                            _ = v26
                        }
                        v27 := GenS2{F0: v1, F1: -1, F2: v15}
                        v1 = v1
                        if true {
                            v6[""] = v11
                            if true {
                                v1 -= v1
                                fmt.Println(v15, v1, v11)
                            } else {
                                var v28 int16 = (v12 >> ((0) & 63))
                                if (v15 >= v15) {
                                    fmt.Println(v12)
                                    v28 = (^(v28))
                                    if true {
                                        v1 = (v1 - v1)
                                        v29 := []int{(^(fi13)), int(v28), (fi14 - -691)}
                                        _ = v29
                                    }
                                }
                                var v30 int16 = v12
                                _ = v28
                                _ = v30
                            }
                            var v31 uint32 = 1
                            fmt.Println(fi13, v11, v15)
                            _ = v31
                        } else {
                            var v32 uint = 42
                            _ = v32
                        }
                        _ = v27
                    } else {
                        v33 := GenS3{F0: (fi13 % (((fi14 >> ((4294967295) & 63))) | 1)), F1: 2, F2: v1}
                        _ = v33
                    }
                    switch (^(fi13)) {
                    case 4294967295:
                        v16 += 1
                    }
                    v34 := GenS2{F0: (v1 << ((215) & 63)), F1: -378, F2: v15}
                    v11 = (fi14 + v11)
                    _ = v34
                } else {
                    if true {
                        var v35 int64 = (v1 % ((v1) | 1))
                        _ = v35
                    }
                    _ = len(v6)
                    var v36 string = "\n\t"
                    _ = v36
                }
                var v37 bool = (("foo bar" + "\n\t") == ((("hello" + (((("hello" + "") + (("hello" + "\n\t") + (("a" + (("" + "") + ("a" + "\n\t"))) + "a"))) + "a") + "hello")) + "foo bar") + (((("a" + ((("hello" + ("a" + (("hello" + ("foo bar" + "foo bar")) + ((("foo bar" + (("hello" + ("foo bar" + "\n\t")) + "\n\t")) + (("\n\t" + "\n\t") + "\n\t")) + ("hello" + "\n\t"))))) + "hello") + "a")) + "hello") + "a") + ("a" + "\n\t"))))
                _ = v16
                _ = v37
            } else {
                v1 += 9223372036854775807
                var v38 int = (fi13 * (fi13 / ((32768) | 1)))
                var v39 uint64 = 4294967295
                v1 += v1
                _ = v38
                _ = v39
            }
            switch (fi13 + (v11 % ((fi14) | 1))) {
            case 9223372036854775807:
                v40 := []int{}
                _ = v40
            case 979:
                fmt.Println(fi14, v1)
            case -515:
                fmt.Println(v12, v11)
            }
            v15 = v15
            v41 := []int{}
            v42 := GenS2{F0: (-(v1)), F1: 42, F2: 2}
            v11 += 2147483647
            _ = fi14
            _ = v15
            _ = v41
            _ = v42
        }
        for fi43 := 0; fi43 < 9; fi43++ {
            switch ("a" + "a") {
            case "hello":
                var v44 int8 = 0
                _ = v44
            default:
                v45 := GenS4{F0: 117}
                _ = v45
            }
            v12 = v12
            v46 := GenS2{F0: v1, F1: (-(fi13)), F2: 30}
            _ = fi43
            _ = v46
        }
        _ = fi13
    }
    switch "" {
    case "a":
        var v47 bool = true
        _ = v47
    default:
        var v48 int8 = int8(v1)
        _ = v48
    }
    for fi49 := 0; fi49 < 4; fi49++ {
        v11 = v11
        _ = fi49
    }
    switch ("" + "\n\t") {
    case "":
        var v50 int = v11
        _ = v50
    case "hello":
        if true {
            for fi51 := 0; fi51 < 5; fi51++ {
                var v52 float64 = 0.1
                _ = fi51
                _ = v52
            }
            v4[""] = true
        }
    default:
        v53 := GenS2{F0: (v1 ^ v1), F1: (^(v11)), F2: 42}
        _ = v53
    }
    fmt.Println(v1, v12, v12)
    var v54 uint32 = 2147483647
    v55 := []int{v11}
    switch ("" + ("" + ("\n\t" + "a"))) {
    case "\n\t":
        v56 := GenS5{F0: 188}
        _ = v56
    case "":
        if (42 <= 0) {
            var v57 int32 = 2147483647
            _ = v57
        } else {
            fmt.Println(v54, v1, v11)
            for fi58 := 0; fi58 < 5; fi58++ {
                switch int(v12) {
                case 824:
                    switch fi58 {
                    case 450:
                        v54 = v54
                    case 1:
                        var v59 uint16 = 255
                        _ = v59
                    default:
                        var v60 uint64 = 888
                        _ = v60
                    }
                case -2147483648:
                    if len(v55) > 0 {
                        v55[0] = (v11 / (((v11 & fi58)) | 1))
                    }
                }
                var v61 string = "foo bar"
                _ = fi58
                _ = v61
            }
        }
    case "a":
        v54 = (v54 | v54)
    default:
        v11 -= 671
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v54:", v54)
    fmt.Println("slicelen v55:", len(v55), v55)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v55
    _ = v6
}
