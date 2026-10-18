package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 int32; F1 int }
type GenS5 struct { F0 int8; F1 int }
func genFn0(p0 bool, p1 string) uint32 {
    if (!p0) {
        v7 := []int{-9223372036854775808, 32767}
        _ = v7
    } else {
        var v8 uint = 42
        fmt.Println(v8)
        v9 := GenS1{F0: -422, F1: -424}
        _ = v8
        _ = v9
    }
    p0 = (0 == 4294967295)
    v10 := GenS1{F0: 0, F1: 1}
    if (0 >= 128) {
        if p0 {
            var v11 string = (p1 + "a")
            var v12 int8 = 127
            _ = v11
            _ = v12
        } else {
            if (1e10 >= (1.5 + 1.5)) {
                fmt.Println(p0, p1)
                var v13 uint32 = 2
                v13 = (v13 << ((18446744073709551615) & 63))
                _ = v13
            } else {
                fmt.Println(p1, p0)
                p1 = p1
            }
            var v14 bool = (p0 || (p1 == ""))
            var v15 bool = (217 > 741)
            _ = v14
            _ = v15
        }
    }
    _ = p0
    _ = p1
    _ = v10
    return 106
}

func main() {
    var v1 int16 = 32767
    v2 := make(map[uint32]bool)
    var v3 bool = (9223372036854775807 >= 619)
    var v4 uint8 = 128
    v5 := make(map[string]string)
    var v6 uint8 = 0
    for fi16 := 0; fi16 < 2; fi16++ {
        v17 := GenS2{F0: 42}
        _ = fi16
        _ = v17
    }
    v1 -= (v1 >> ((217) & 63))
    fmt.Println(v4, v4, v3)
    var v18 int8 = 75
    var v19 float32 = (1e10 / 1.5)
    var v20 uint16 = 65535
    v21 := GenS3{F0: v3, F1: -963}
    for fi22 := 0; fi22 < 4; fi22++ {
        for fi23 := 0; fi23 < 17; fi23++ {
            v18 -= int8(fi23)
            v19 = v19
            _ = fi23
        }
        for fi24 := 0; fi24 < 2; fi24++ {
            switch (fi22 | fi24) {
            case 541:
                v25 := v2[952]
                _ = v25
            case -439:
                if true {
                    v20 = 23
                    if (-1 == fi22) {
                        if (!((((0.0 / 1.0) - v19) - v19) >= v19)) {
                            v26 := []int{127, -1, fi24}
                            _ = v26
                        } else {
                            if (v3 && v3) {
                                var v27 bool = ((v18 + v18) == (v18 | v18))
                                switch ("\n\t" + "a") {
                                case "foo bar":
                                    var v28 int16 = (v1 * v1)
                                    if v3 {
                                        v29 := GenS2{F0: 1000}
                                        var v30 uint = 255
                                        _ = v29
                                        _ = v30
                                    }
                                    _ = v28
                                default:
                                    var v31 int8 = v18
                                    _ = v31
                                }
                                _ = v27
                            } else {
                                fmt.Println(v4)
                                fmt.Println(v18, v18)
                            }
                            v6 = 0
                            var v32 int32 = 0
                            v3 = ((v19 < v19) || v3)
                            _ = v32
                        }
                    } else {
                        v33 := []int{}
                        v6 -= 177
                        var v34 int16 = (v1 >> ((42) & 63))
                        _ = v33
                        _ = v34
                    }
                } else {
                    switch "a" {
                    case "\n\t":
                        if (4294967295 == 953) {
                            fmt.Println(v1, v1)
                            v35 := []int{fi24, fi22}
                            _ = v35
                        } else {
                            var v36 int64 = -1
                            v6 += (v6 / ((128) | 1))
                            switch fi24 {
                            case -762:
                                _ = len(v2)
                            case 460:
                                v19 = 0.0
                            case 42:
                                v37 := []int{(-(fi24))}
                                v5["foo bar"] = "hello"
                                var v38 string = "a"
                                _ = v37
                                _ = v38
                            }
                            if v3 {
                                v4 += v6
                                fmt.Println(fi24)
                            }
                            v39 := GenS3{F0: v3, F1: 42}
                            _ = v36
                            _ = v39
                        }
                    case "hello":
                        switch (fi22 & 1) {
                        case 1:
                            v40 := GenS4{F0: -653, F1: fi24}
                            _ = v40
                        default:
                            fmt.Println(v19, v4, v19)
                        }
                    }
                    v41 := GenS5{F0: (v18 % (((v18 - -73)) | 1)), F1: (^(fi22))}
                    _ = len(v2)
                    _ = v41
                }
            case 32768:
                v6 = (v6 % (((v4 % (((v4 << ((384) & 63))) | 1))) | 1))
            default:
                v42 := v5["foo bar"]
                _ = v42
            }
            _ = fi24
        }
        _ = fi22
    }
    v43 := []int{919, 2147483647}
    v44 := GenS4{F0: 42, F1: 32767}
    switch "hello" {
    case "foo bar":
        for fi45 := 0; fi45 < 11; fi45++ {
            var v46 string = ("\n\t" + "\n\t")
            _ = fi45
            _ = v46
        }
    case "hello":
        if len(v43) >= 2 {
            v43 = v43[1:len(v43)]
        }
    default:
        v20 = 2
    }
    var v47 float64 = 1.5
    var v48 uint32 = 495
    if len(v43) >= 2 {
        v43 = v43[1:len(v43)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
