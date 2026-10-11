package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int8; F1 uint }
func genFn0(p0 string, p1 uint8) bool {
    v6 := []int{}
    fmt.Println(p1, p1)
    for fi7 := 0; fi7 < 3; fi7++ {
        if ((-70 < 42) && (((float64((3.14159 * 0.0)) - (3.14159 + 1e10)) + ((float64((-2.25 * 3.14159)) + (0.1 - (0.1 - 1.5))) / (1.5 + 0.1))) <= (1e10 * float64(0.1)))) {
            p0 = p0
            p1 -= (p1 / ((p1) | 1))
            for fi8 := 0; fi8 < 6; fi8++ {
                var v9 uint16 = uint16(p1)
                var v10 uint32 = 65535
                _ = fi8
                _ = v10
                _ = v9
            }
        }
        v11 := GenS1{F0: 255}
        p0 = (p0 + ("a" + "hello"))
        _ = fi7
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v6
    return false
}
func genFn1(p0 uint16) bool {
    v12 := []int{-981, 32768, 128}
    fmt.Println(p0)
    fmt.Println(p0)
    _ = p0
    _ = v12
    return false
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]int)
    v3 := []int{607}
    v4 := make(map[uint32]int)
    var v5 float64 = 0.1
    switch "\n\t" {
    case "":
        for fi13 := 0; fi13 < 18; fi13++ {
            _ = len(v4)
            v5 = float64(-2.25)
            _ = fi13
        }
    case "hello":
        var v14 int16 = 45
        _ = v14
    }
    var v15 float32 = 3.14159
    v16 := GenS1{F0: 255}
    v17 := GenS1{F0: 1}
    v18 := GenS1{F0: 32767}
    v15 -= v15
    var v19 uint32 = 4294967295
    var v20 int8 = -50
    switch ("hello" + ("a" + "foo bar")) {
    case "\n\t":
        var v21 int8 = v20
        _ = v21
    case "":
        var v22 int16 = 0
        _ = v22
    default:
        if false {
            for fi23 := 0; fi23 < 14; fi23++ {
                v24 := GenS2{F0: 99, F1: fi23}
                v25 := GenS3{F0: -32768}
                _ = fi23
                _ = v24
                _ = v25
            }
            if (true && (true || true)) {
                var v26 uint16 = 32767
                var v27 uint64 = 1
                for fi28 := 0; fi28 < 11; fi28++ {
                    var v29 bool = true
                    _ = len(v2)
                    fmt.Println(v19)
                    _ = fi28
                    _ = v29
                }
                _ = v26
                _ = v27
            } else {
                var v30 string = "\n\t"
                for fi31 := 0; fi31 < 9; fi31++ {
                    if true {
                        fmt.Println(fi31)
                        v32 := GenS4{F0: -43, F1: 65535}
                        for fi33 := 0; fi33 < 3; fi33++ {
                            fmt.Println(v20, v19)
                            _ = fi33
                        }
                        _ = v32
                    } else {
                        if true {
                            v3 = append(v3, (fi31 << ((0) & 63)))
                            v20 -= (^(v20))
                            if ((v15 + 0.1) >= v15) {
                                var v34 int64 = -9223372036854775808
                                v35 := GenS2{F0: 0, F1: fi31}
                                _ = v34
                                _ = v35
                            } else {
                                v5 = (0.1 / 3.14159)
                            }
                        } else {
                            for fi36 := 0; fi36 < 10; fi36++ {
                                v20 += (^(v20))
                                v5 -= float64(v15)
                                fmt.Println(v30, v15)
                                _ = fi36
                            }
                            v30 = v30
                            if (0 >= 81) {
                                if false {
                                    var v37 int8 = -128
                                    v15 = 0.0
                                    _ = len(v3)
                                    _ = v37
                                }
                                var v38 int64 = 9223372036854775807
                                _ = v38
                            } else {
                                v15 = (v15 - (1e10 * 0.0))
                                var v39 bool = (true || (!true))
                                if (42 != -2147483648) {
                                    var v40 int64 = 9223372036854775807
                                    fmt.Println(v5)
                                    _ = v40
                                }
                                _ = v39
                            }
                        }
                    }
                    var v41 uint16 = 65535
                    _ = fi31
                    _ = v41
                }
                _ = v30
            }
            v42 := v2["\n\t"]
            _ = v42
        } else {
            var v43 uint = 594
            fmt.Println(v5, v43)
            var v44 uint64 = 2
            _ = v43
            _ = v44
        }
    }
    for fi45 := 0; fi45 < 8; fi45++ {
        v46 := GenS1{F0: 0}
        _ = fi45
        _ = v46
    }
    var v47 uint32 = v19
    v5 += v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v47
    _ = v5
}
