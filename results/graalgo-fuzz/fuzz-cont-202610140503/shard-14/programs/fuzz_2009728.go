package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8; F1 uint32; F2 uint32 }
type GenS4 struct { F0 uint64; F1 uint32; F2 int64 }

func main() {
    var v1 uint = 255
    var v2 uint8 = 127
    v3 := make(map[uint32]string)
    var v4 bool = (470 != 2)
    var v5 int32 = 1
    v6 := make(map[int]string)
    v7 := GenS1{F0: "a", F1: v2, F2: ((("a" + "a") + "a") + ("a" + (("foo bar" + "foo bar") + "hello")))}
    if v4 {
        if v4 {
            switch ("\n\t" + "a") {
            case "":
                var v8 uint8 = v2
                _ = v8
            case "a":
                var v9 float64 = 1e10
                _ = v9
            default:
                for fi10 := 0; fi10 < 10; fi10++ {
                    if ((!(255 <= 42)) && (!(2 == 42))) {
                        switch "a" {
                        case "hello":
                            if (!(v4 || ((v5 - (-(v5))) != v5))) {
                                v3[2147483647] = ""
                                fmt.Println(v4, v2, v4)
                                var v11 uint = v1
                                _ = v11
                            } else {
                                if (true && ((0.0 - 1e10) != (0.1 + float32(0.1)))) {
                                    if true {
                                        var v12 uint16 = 255
                                        _ = v12
                                    } else {
                                        switch fi10 {
                                        case 42:
                                            var v13 uint16 = 42
                                            _ = v13
                                        }
                                        v5 = v5
                                    }
                                    var v14 int16 = -272
                                    _ = v14
                                }
                                if (((1.5 + 1e10) - float64(float32(0.1))) >= (((float64(0.0) - 1e10) * 1e10) + (-2.25 * float64(1.5)))) {
                                    v15 := GenS1{F0: "foo bar", F1: v2, F2: ((("hello" + "foo bar") + "hello") + "hello")}
                                    var v16 int8 = 42
                                    switch ("foo bar" + "foo bar") {
                                    case "foo bar":
                                        v17 := GenS1{F0: "hello", F1: v2, F2: "a"}
                                        _ = v17
                                    default:
                                        v18 := []int{(^(fi10)), (fi10 | (fi10 & fi10))}
                                        _ = v18
                                    }
                                    _ = v15
                                    _ = v16
                                }
                            }
                        }
                        fmt.Println(v5, fi10, v1)
                    }
                    if false {
                        switch (fi10 - 128) {
                        case -9223372036854775808:
                            _ = len(v3)
                        case 9223372036854775807:
                            v19 := GenS2{F0: 40}
                            _ = v19
                        case 42:
                            v20 := []int{fi10, 991, fi10}
                            _ = v20
                        }
                        v21 := GenS1{F0: "a", F1: (v2 * (v2 | v2)), F2: ""}
                        _ = v21
                    } else {
                        fmt.Println(v2, v1)
                        v22 := []int{(fi10 / ((fi10) | 1)), 42, fi10, (-(fi10))}
                        switch (fi10 / ((1) | 1)) {
                        case 32768:
                            v23 := GenS3{F0: v2, F1: 250, F2: 0}
                            _ = v23
                        default:
                            var v24 float32 = (1e10 / 0.1)
                            _ = v24
                        }
                        _ = v22
                    }
                    for fi25 := 0; fi25 < 13; fi25++ {
                        if v4 {
                            switch fi25 {
                            case 856:
                                v26 := []int{(fi10 / (((fi25 + int(v2))) | 1))}
                                _ = v26
                            case -1:
                                var v27 uint = (v1 % ((v1) | 1))
                                _ = v27
                            case 32768:
                                v1 -= 42
                            default:
                                if true {
                                    v28 := GenS2{F0: 55}
                                    _ = len(v6)
                                    _ = v28
                                } else {
                                    delete(v3, 0)
                                    var v29 uint64 = 42
                                    v30 := GenS3{F0: 2, F1: uint32(v1), F2: 1}
                                    _ = v29
                                    _ = v30
                                }
                            }
                        } else {
                            v2 = (v2 - 255)
                            v4 = v4
                            v31 := GenS3{F0: v2, F1: 65535, F2: 271}
                            var v32 uint32 = 268
                            _ = v31
                            _ = v32
                        }
                        v33 := []int{}
                        _ = fi25
                        _ = v33
                    }
                    _ = fi10
                }
            }
            if v4 {
                if (!(false || v4)) {
                    v34 := GenS2{F0: -85}
                    _ = v34
                }
                v2 += 1
            }
        } else {
            var v35 uint = v1
            _ = v35
        }
        for fi36 := 0; fi36 < 20; fi36++ {
            v5 -= v5
            var v37 int16 = -1
            var v38 uint16 = 0
            _ = fi36
            _ = v37
            _ = v38
        }
    }
    v39 := []int{-697, 127}
    var v40 int8 = 127
    delete(v3, 2147483647)
    var v41 uint32 = 4294967295
    v2 = (v2 * 21)
    v42 := GenS2{F0: -128}
    fmt.Println(v41)
    if len(v39) >= 2 {
        v39 = v39[1:len(v39)]
    }
    var v43 uint16 = 992
    var v44 float32 = (-2.25 / 1.5)
    for fi45 := 0; fi45 < 15; fi45++ {
        var v46 int = fi45
        fmt.Println(v2, v2)
        _ = fi45
        _ = v46
    }
    switch ("hello" + ("a" + "foo bar")) {
    case "\n\t":
        var v47 string = ("foo bar" + "a")
        _ = v47
    case "":
        _ = len(v6)
    case "hello":
        for fi48 := 0; fi48 < 13; fi48++ {
            v49 := GenS4{F0: 4294967295, F1: (v41 << ((v1) & 63)), F2: 9223372036854775807}
            v4 = (false || (879 > 1))
            _ = fi48
            _ = v49
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v6
    _ = v7
}
