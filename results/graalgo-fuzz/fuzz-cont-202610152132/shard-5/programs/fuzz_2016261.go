package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 int }
func genFn0() int32 {
    for fi5 := 0; fi5 < 8; fi5++ {
        fmt.Println(fi5)
        var v6 int32 = -1
        _ = fi5
        _ = v6
    }
    var v7 string = "\n\t"
    v7 = v7
    var v8 bool = true
    _ = v7
    _ = v8
    return 1
}
func genFn1(p0 int64, p1 uint) string {
    switch "foo bar" {
    case "":
        p1 -= p1
    case "hello":
        v9 := []int{2147483647, 541}
        _ = v9
    case "a":
        var v10 uint32 = 1
        _ = v10
    default:
        v11 := GenS1{F0: 0, F1: -74}
        _ = v11
    }
    var v12 int = 128
    if true {
        var v13 string = "hello"
        var v14 bool = (-2147483648 == -2147483648)
        v15 := GenS2{F0: (-(v12)), F1: 1}
        _ = v13
        _ = v14
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v12
    return "\n\t"
}

func main() {
    var v1 int64 = 1
    v2 := make(map[string]int)
    v3 := make(map[string]bool)
    var v4 uint32 = 2147483647
    fmt.Println(v1, v4)
    _ = len(v2)
    delete(v2, ("" + "\n\t"))
    if (42 > 32767) {
        var v16 uint8 = 1
        v16 = 236
        v1 += v1
        _ = v16
    } else {
        if false {
            if (v1 <= (v1 - 302)) {
                v4 += (v4 | 0)
                switch (("foo bar" + "\n\t") + "") {
                case "":
                    v4 = 42
                default:
                    var v17 uint32 = v4
                    _ = v17
                }
                switch 45 {
                case -391:
                    var v18 uint16 = 32767
                    _ = v18
                case 2147483647:
                    for fi19 := 0; fi19 < 7; fi19++ {
                        v1 = -407
                        _ = fi19
                    }
                default:
                    fmt.Println(v1, v4)
                }
            } else {
                var v20 float32 = 1.5
                delete(v3, "\n\t")
                for fi21 := 0; fi21 < 18; fi21++ {
                    v20 += 1e10
                    var v22 string = "\n\t"
                    _ = fi21
                    _ = v22
                }
                _ = v20
            }
        } else {
            switch 32767 {
            case -2147483648:
                v4 = v4
            case 601:
                for fi23 := 0; fi23 < 12; fi23++ {
                    for fi24 := 0; fi24 < 6; fi24++ {
                        switch "foo bar" {
                        case "hello":
                            if false {
                                var v25 int8 = 1
                                v26 := GenS1{F0: (v4 % (((v4 % ((1) | 1))) | 1)), F1: (v1 - (v1 >> ((18446744073709551615) & 63)))}
                                v27 := v2["\n\t"]
                                _ = v25
                                _ = v26
                                _ = v27
                            }
                        case "foo bar":
                            v28 := []int{}
                            _ = v28
                        }
                        v1 = v1
                        _ = fi24
                    }
                    switch (^(fi23)) {
                    case -2147483648:
                        if false {
                            for fi29 := 0; fi29 < 19; fi29++ {
                                var v30 int32 = 42
                                _ = fi29
                                _ = v30
                            }
                        } else {
                            var v31 uint = 652
                            fmt.Println(v31, v4)
                            for fi32 := 0; fi32 < 17; fi32++ {
                                v4 = (v4 % (((v4 % (((v4 * (v4 >> (((v31 ^ 431)) & 63)))) | 1))) | 1))
                                if false {
                                    if (42 == 1) {
                                        v33 := v2[""]
                                        _ = v33
                                    } else {
                                        v34 := []int{(fi23 ^ (fi32 | int(v4)))}
                                        var v35 string = "foo bar"
                                        var v36 uint32 = 42
                                        _ = v34
                                        _ = v35
                                        _ = v36
                                    }
                                } else {
                                    switch fi23 {
                                    case -1:
                                        var v37 uint32 = 144
                                        _ = v37
                                    case -885:
                                        if false {
                                            switch ("" + "a") {
                                            case "\n\t":
                                                v4 -= (v4 ^ (v4 * 2))
                                            default:
                                                v2["a"] = 772
                                            }
                                        }
                                    case 32767:
                                        delete(v3, "hello")
                                    default:
                                        fmt.Println(fi32)
                                    }
                                }
                                v31 = v31
                                _ = fi32
                            }
                            _ = v31
                        }
                    default:
                        v1 += v1
                    }
                    _ = fi23
                }
            default:
                v38 := GenS1{F0: (v4 ^ 0), F1: v1}
                _ = v38
            }
        }
        v39 := GenS1{F0: (v4 % (((^(v4))) | 1)), F1: v1}
        _ = v39
    }
    v40 := []int{}
    v4 = v4
    var v41 uint64 = 18446744073709551615
    var v42 uint8 = 255
    v43 := GenS2{F0: 128, F1: 42}
    v44 := GenS3{F0: 32768}
    var v45 int64 = v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
}
