package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint8 }
func genFn0(p0 bool) int {
    if (!(320 > -1)) {
        v4 := []int{-9223372036854775808, 128, 955}
        v5 := GenS1{F0: 0}
        p0 = true
        _ = v4
        _ = v5
    } else {
        if false {
            p0 = p0
            var v6 float64 = 3.14159
            var v7 int = -9223372036854775808
            _ = v6
            _ = v7
        } else {
            fmt.Println(p0)
            fmt.Println(p0)
            v8 := GenS1{F0: 65535}
            _ = v8
        }
        v9 := GenS2{F0: 798, F1: 4294967295}
        v10 := GenS3{F0: "foo bar"}
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 12; fi11++ {
        v12 := GenS4{F0: ((150 < 42) && (42 <= 128))}
        fmt.Println(fi11)
        v13 := []int{}
        _ = fi11
        _ = v12
        _ = v13
    }
    _ = p0
    return 0
}

func main() {
    v1 := []int{772, 128}
    var v2 string = "foo bar"
    v3 := make(map[string]int)
    v2 = v2
    var v14 uint = 42
    if false {
        v14 = v14
        var v15 uint32 = 269
        _ = v15
    }
    if true {
        v2 = (v2 + ("foo bar" + "foo bar"))
        if false {
            for fi16 := 0; fi16 < 8; fi16++ {
                if true {
                    if (!true) {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v17 := GenS4{F0: false}
                        _ = v17
                    } else {
                        if (!(0.1 <= 3.14159)) {
                            v18 := GenS2{F0: 0, F1: 1}
                            var v19 int = (fi16 << ((30) & 63))
                            _ = v18
                            _ = v19
                        } else {
                            v2 = v2
                            switch (fi16 * (fi16 + fi16)) {
                            case 128:
                                var v20 float32 = ((1.5 - (0.1 * (1.5 - -2.25))) / -2.25)
                                _ = v20
                            case 965:
                                v14 = (v14 * (v14 * v14))
                            case -9223372036854775808:
                                v21 := GenS2{F0: -1, F1: (v14 << (((v14 >> ((v14) & 63))) & 63))}
                                _ = v21
                            }
                            var v22 int = int(v14)
                            _ = v22
                        }
                    }
                } else {
                    v23 := v3[v2]
                    _ = v23
                }
                var v24 int32 = 0
                v24 = (-(v24))
                _ = fi16
                _ = v24
            }
        }
        var v25 int64 = 50
        _ = v25
    }
    v2 = v2
    for fi26 := 0; fi26 < 11; fi26++ {
        v27 := GenS5{F0: 128}
        v14 = (v14 & v14)
        _ = fi26
        _ = v27
    }
    if ((0 > 981) || false) {
        v28 := v3[v2]
        v29 := GenS5{F0: 145}
        _ = v28
        _ = v29
    } else {
        fmt.Println(v2)
    }
    if (-959 >= -1) {
        if false {
            v14 += 4294967295
            switch "\n\t" {
            case "hello":
                if false {
                    fmt.Println(v14, v14)
                    if (!(v2 == "")) {
                        var v30 uint16 = 0
                        fmt.Println(v14, v2, v30)
                        _ = v30
                    } else {
                        v31 := GenS1{F0: 2147483648}
                        _ = v31
                    }
                    _ = len(v3)
                } else {
                    fmt.Println(v2, v2)
                }
            default:
                v32 := GenS5{F0: 0}
                _ = v32
            }
            v14 = (v14 | (^(v14)))
        } else {
            var v33 int = 32768
            _ = v33
        }
        v2 = v2
    }
    for fi34 := 0; fi34 < 16; fi34++ {
        delete(v3, "\n\t")
        var v35 uint32 = 2147483647
        v1 = append(v1, (-(fi34)))
        _ = fi34
        _ = v35
    }
    v14 = 4294967295
    v14 -= v14
    if (1 >= 109) {
        v14 += v14
        v2 = ("hello" + "hello")
    }
    if len(v1) > 0 {
        v1[0] = 42
    }
    _ = len(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v2
    _ = v3
}
