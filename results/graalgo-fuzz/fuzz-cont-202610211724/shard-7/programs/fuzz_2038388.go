package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 int8; F1 int16 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int16; F1 uint64; F2 uint32 }

func main() {
    v1 := make(map[int]bool)
    v2 := []string{("" + "a")}
    var v3 uint8 = 2
    v4 := make(map[string]string)
    v5 := make(map[string]int)
    v6 := []string{}
    fmt.Println(v3)
    var v7 int64 = 42
    for fi8 := 0; fi8 < 10; fi8++ {
        switch ("\n\t" + "foo bar") {
        case "":
            fmt.Println(v3, v7, v7)
        default:
            var v9 uint = 815
            _ = v9
        }
        v7 = v7
        var v10 int16 = -406
        _ = fi8
        _ = v10
    }
    v11 := GenS1{F0: 2, F1: 755}
    var v12 int32 = -2147483648
    var v13 uint = 4294967295
    fmt.Println(v12, v7, v7)
    v14 := GenS1{F0: 2, F1: 2}
    v15 := GenS1{F0: 42, F1: 687}
    var v16 uint16 = 255
    var v17 int64 = v7
    if false {
        for fi18 := 0; fi18 < 12; fi18++ {
            fmt.Println(fi18, fi18)
            v7 = 355
            for fi19 := 0; fi19 < 17; fi19++ {
                v20 := GenS1{F0: 1, F1: v16}
                _ = fi19
                _ = v20
            }
            _ = fi18
        }
        v21 := GenS2{F0: v12, F1: v7}
        v6 = append(v6, "hello")
        _ = v21
    } else {
        v7 += (^(v7))
        if (!false) {
            v1[-258] = true
            for fi22 := 0; fi22 < 12; fi22++ {
                if true {
                    fmt.Println(v16, fi22)
                    v13 -= (v13 - (v13 + 0))
                    if ((v12 % (((v12 >> (((v13 >> ((2) & 63))) & 63))) | 1)) != v12) {
                        switch "foo bar" {
                        case "":
                            if true {
                                v13 += 18446744073709551615
                            } else {
                                var v23 bool = (((v13 << ((720) & 63)) != v13) && true)
                                for fi24 := 0; fi24 < 19; fi24++ {
                                    fmt.Println(v23)
                                    fmt.Println(fi24)
                                    v3 += (v3 / (((v3 + 242)) | 1))
                                    if v23 {
                                        delete(v4, "hello")
                                        fmt.Println(v7)
                                        var v25 int32 = 848
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        fmt.Println(v17)
                                        _ = v25
                                    } else {
                                        v26 := GenS3{F0: (fi22 - fi22), F1: (42 >= -128)}
                                        v12 += v12
                                        _ = v26
                                    }
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    fmt.Println(fi22, v12)
                                    v27 := GenS1{F0: 2, F1: (v16 ^ (v16 / ((v16) | 1)))}
                                    v23 = (!false)
                                    v2 = append(v2, (("a" + "foo bar") + ((("foo bar" + "foo bar") + "") + "a")))
                                    v12 += int32(fi22)
                                    switch ("foo bar" + "foo bar") {
                                    case "foo bar":
                                        delete(v5, "")
                                    }
                                    _ = fi24
                                    _ = v27
                                }
                                _ = v23
                            }
                        }
                        switch "foo bar" {
                        case "foo bar":
                            if true {
                                var v28 string = "foo bar"
                                _ = v28
                            }
                        case "\n\t":
                            if (false && false) {
                                v29 := GenS4{F0: 100, F1: 688}
                                v17 -= (v7 % ((1) | 1))
                                v30 := GenS4{F0: 127, F1: 1000}
                                _ = v29
                                _ = v30
                            } else {
                                v31 := GenS5{F0: 0}
                                v32 := v1[(fi22 % ((42) | 1))]
                                _ = v31
                                _ = v32
                            }
                        case "hello":
                            switch 571 {
                            case -2147483648:
                                delete(v4, ((("foo bar" + (("" + "a") + ("" + ("a" + "a")))) + "\n\t") + "foo bar"))
                            case 127:
                                v33 := GenS6{F0: 691, F1: 2, F2: 2147483647}
                                _ = v33
                            }
                        default:
                            v1[fi22] = true
                        }
                    }
                } else {
                    var v34 string = "hello"
                    switch "\n\t" {
                    case "":
                        _ = len(v5)
                    default:
                        v17 = v17
                    }
                    if false {
                        fmt.Println(v16, v12)
                        fmt.Println(v34)
                    }
                    _ = v34
                }
                v12 = v12
                v16 = 0
                _ = fi22
            }
            for fi35 := 0; fi35 < 9; fi35++ {
                var v36 int32 = v12
                _ = fi35
                _ = v36
            }
        } else {
            v12 -= v12
            var v37 int = 821
            _ = v37
        }
        var v38 uint16 = v16
        _ = v38
    }
    var v39 int32 = int32(v17)
    delete(v4, ("" + "a"))
    v12 += (v39 ^ 2147483647)
    fmt.Println(v17, v3)
    v40 := GenS4{F0: 0, F1: 1}
    fmt.Println(v3, v7, v12)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v7
}
