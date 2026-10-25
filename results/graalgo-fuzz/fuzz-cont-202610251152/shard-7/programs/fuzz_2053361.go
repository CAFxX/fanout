package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32; F1 uint8; F2 int16 }
type GenS4 struct { F0 uint; F1 uint16 }
func genFn0(p0 int16) uint {
    var v4 bool = (false || true)
    v5 := []int{357}
    var v6 uint = 1
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return v6
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := make(map[string]bool)
    var v3 uint = 2
    v1 += v1
    if true {
        if false {
            v1 += 1
            v7 := GenS1{F0: 2147483647, F1: -330}
            _ = v7
        }
        v3 = v3
    } else {
        v2[("hello" + "hello")] = (42 > 65535)
    }
    var v8 int64 = 9223372036854775807
    switch -736 {
    case 42:
        var v9 int = 32767
        _ = v9
    case 0:
        fmt.Println(v1)
    case 128:
        fmt.Println(v3)
    default:
        v3 = (v3 % ((genFn0(1)) | 1))
    }
    v10 := GenS2{F0: -9223372036854775808}
    if ((!true) || ("a" >= ("\n\t" + "hello"))) {
        var v11 int = -656
        v12 := []int{v11, -837, -9223372036854775808, -49}
        if (!true) {
            fmt.Println(v1, v3, v3)
            v1 += v1
        }
        _ = v11
        _ = v12
    } else {
        v13 := GenS1{F0: 2147483647, F1: 87}
        _ = v13
    }
    switch 669 {
    case 1:
        if (65535 == 2) {
            for fi14 := 0; fi14 < 1; fi14++ {
                if true {
                    if (true && false) {
                        switch ("foo bar" + "hello") {
                        case "\n\t":
                            for fi15 := 0; fi15 < 5; fi15++ {
                                switch (fi14 ^ -1) {
                                case 4294967295:
                                    if false {
                                        if (true && true) {
                                            _ = len(v2)
                                            v16 := []int{fi14, fi14}
                                            _ = v16
                                        } else {
                                            var v17 int32 = 2147483647
                                            _ = v17
                                        }
                                    }
                                case -342:
                                    delete(v2, ("" + ""))
                                case 2147483647:
                                    v18 := GenS1{F0: 42, F1: 1000}
                                    _ = v18
                                }
                                _ = fi15
                            }
                        case "":
                            var v19 float64 = 1.5
                            _ = v19
                        case "a":
                            var v20 int64 = int64(v3)
                            _ = v20
                        default:
                            var v21 int32 = 1
                            _ = v21
                        }
                        switch ("" + "hello") {
                        case "":
                            var v22 string = "a"
                            _ = v22
                        }
                    }
                } else {
                    var v23 float64 = (-2.25 + ((1e10 * 3.14159) - 1.5))
                    v24 := []int{(fi14 & fi14), (fi14 + fi14), 4294967295, (^(fi14))}
                    _ = v23
                    _ = v24
                }
                _ = fi14
            }
            for fi25 := 0; fi25 < 7; fi25++ {
                var v26 int8 = 6
                v1 -= (v1 * (^(v8)))
                var v27 int8 = (v26 / ((-118) | 1))
                _ = fi25
                _ = v26
                _ = v27
            }
        }
    case 32768:
        v28 := GenS1{F0: 2, F1: 32767}
        _ = v28
    case 835:
        var v29 int16 = 32767
        _ = v29
    default:
        v30 := GenS3{F0: -2147483648, F1: 160, F2: 0}
        _ = v30
    }
    var v31 string = ("a" + "")
    if true {
        v32 := []int{115, 4294967295, -1, int(v1)}
        _ = v32
    }
    for fi33 := 0; fi33 < 13; fi33++ {
        var v34 int16 = 42
        for fi35 := 0; fi35 < 6; fi35++ {
            var v36 string = v31
            if true {
                v37 := v2[v31]
                _ = v37
            }
            if false {
                v1 -= (-(v1))
                v38 := []int{fi33, 9223372036854775807, 683}
                v2[((v31 + "foo bar") + "foo bar")] = false
                v1 = (v8 - v8)
                _ = v38
            } else {
                fmt.Println(v31, v1, v3)
                v39 := GenS1{F0: 565, F1: (v34 | 972)}
                var v40 float32 = (1e10 + 1e10)
                _ = v39
                _ = v40
            }
            _ = fi35
            _ = v36
        }
        v41 := []int{fi33, 1, (fi33 | (fi33 * (fi33 % ((114) | 1)))), (fi33 / ((fi33) | 1))}
        _ = fi33
        _ = v34
        _ = v41
    }
    var v42 uint = (^(v3))
    v43 := GenS2{F0: 4294967295}
    v44 := []int{9223372036854775807, 128, 2147483647, -450}
    var v45 uint64 = 18446744073709551615
    fmt.Println(v31, v31)
    fmt.Println(v42, v3)
    v46 := GenS4{F0: (^(v42)), F1: uint16(v3)}
    v31 = ""
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v45:", v45)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v31
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v8
}
