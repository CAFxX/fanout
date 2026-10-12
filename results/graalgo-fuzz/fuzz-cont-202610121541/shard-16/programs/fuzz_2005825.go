package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32; F1 uint }
type GenS4 struct { F0 int32; F1 uint16 }
type GenS5 struct { F0 uint16 }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]string)
    var v3 int32 = 2147483647
    fmt.Println(v3)
    fmt.Println(v3)
    v4 := []int{985}
    if true {
        v5 := GenS1{F0: 65535, F1: "", F2: 656}
        _ = v5
    } else {
        v6 := GenS1{F0: 2, F1: "hello", F2: 42}
        v7 := GenS2{F0: "a"}
        _ = v6
        _ = v7
    }
    v3 -= v3
    if false {
        if (!(-2.25 != 0.0)) {
            for fi8 := 0; fi8 < 6; fi8++ {
                var v9 int = (^(fi8))
                v3 = (v3 / (((-(v3))) | 1))
                _ = fi8
                _ = v9
            }
            var v10 int8 = 0
            var v11 int64 = 912
            _ = v10
            _ = v11
        } else {
            switch ("a" + (("" + ((("hello" + "a") + "\n\t") + "")) + ("\n\t" + (("hello" + ((("hello" + ((((("a" + "") + ("foo bar" + "hello")) + ("hello" + "hello")) + "\n\t") + "a")) + "hello") + "")) + "\n\t")))) {
            case "":
                v3 -= v3
            }
        }
    } else {
        switch "foo bar" {
        case "a":
            v12 := v1["a"]
            _ = v12
        }
    }
    delete(v1, "foo bar")
    fmt.Println(v3)
    delete(v2, 569)
    v3 = -647
    switch 2147483647 {
    case 32767:
        v2[42] = "\n\t"
    case -1:
        v3 = 0
    default:
        var v13 int32 = (v3 | (v3 & v3))
        _ = v13
    }
    v14 := GenS3{F0: v3, F1: 2}
    var v15 uint32 = 2147483648
    var v16 uint32 = (^(v15))
    var v17 float32 = (float32((1e10 - -2.25)) * 1e10)
    if true {
        v17 -= 0.0
        if true {
            var v18 int64 = -1
            v3 = -2147483648
            var v19 float64 = (1e10 * 3.14159)
            _ = v18
            _ = v19
        }
    } else {
        if (((float32((((-2.25 - float64(1.5)) - -2.25) / float64((v17 + v17)))) - v17) * (float32(1e10) * v17)) <= v17) {
            if ((v17 * 0.0) != v17) {
                v20 := GenS1{F0: 839, F1: ("foo bar" + "\n\t"), F2: (v16 + (v16 & (v16 - v15)))}
                v21 := GenS4{F0: v3, F1: 65535}
                for fi22 := 0; fi22 < 1; fi22++ {
                    var v23 int32 = v3
                    switch ("\n\t" + ("a" + ("a" + "a"))) {
                    case "a":
                        if false {
                            if false {
                                v23 = -996
                                for fi24 := 0; fi24 < 9; fi24++ {
                                    v25 := GenS3{F0: v23, F1: 169}
                                    var v26 uint64 = 1
                                    _ = fi24
                                    _ = v25
                                    _ = v26
                                }
                                if (true && true) {
                                    delete(v1, "a")
                                } else {
                                    if ((-(v23)) == (-(v3))) {
                                        v16 -= v16
                                        var v27 bool = false
                                        _ = v27
                                    } else {
                                        v28 := GenS3{F0: int32(v15), F1: 255}
                                        _ = v28
                                    }
                                }
                            }
                        }
                    case "foo bar":
                        if (-1 > 1) {
                            if (-2.25 < 1e10) {
                                fmt.Println(v16)
                            }
                            fmt.Println(v16, v15)
                        }
                    }
                    switch fi22 {
                    case -2147483648:
                        for fi29 := 0; fi29 < 7; fi29++ {
                            var v30 int8 = -30
                            v3 -= v23
                            _ = fi29
                            _ = v30
                        }
                    }
                    _ = fi22
                    _ = v23
                }
                _ = v20
                _ = v21
            } else {
                fmt.Println(v16, v17)
                for fi31 := 0; fi31 < 11; fi31++ {
                    var v32 float64 = -2.25
                    v33 := GenS3{F0: (v3 << ((67) & 63)), F1: 470}
                    _ = fi31
                    _ = v32
                    _ = v33
                }
                for fi34 := 0; fi34 < 2; fi34++ {
                    var v35 string = "hello"
                    _ = fi34
                    _ = v35
                }
            }
            var v36 uint32 = v15
            _ = v36
        }
        if (false || ((true || true) && true)) {
            switch 152 {
            case 32767:
                if (127 > -26) {
                    var v37 int64 = -977
                    _ = v37
                }
            case 395:
                v38 := GenS5{F0: 42}
                _ = v38
            case -1:
                switch (("a" + "hello") + "") {
                case "":
                    if true {
                        if true {
                            v16 -= v16
                            fmt.Println(v3)
                        }
                        if (1 == 4294967295) {
                            v16 -= (^(v15))
                        }
                    } else {
                        var v39 uint64 = 290
                        var v40 int16 = 32767
                        switch (("a" + "\n\t") + "\n\t") {
                        case "":
                            v39 = v39
                        case "foo bar":
                            switch "hello" {
                            case "a":
                                var v41 bool = (100 <= int8(v39))
                                _ = v41
                            }
                        case "hello":
                            _ = len(v4)
                        }
                        _ = v39
                        _ = v40
                    }
                default:
                    v42 := GenS5{F0: 255}
                    _ = v42
                }
            }
        } else {
            var v43 string = "a"
            var v44 int = 1
            var v45 uint32 = (^(v16))
            _ = v43
            _ = v44
            _ = v45
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
