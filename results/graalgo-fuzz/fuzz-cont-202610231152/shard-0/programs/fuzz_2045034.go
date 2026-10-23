package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int; F1 int64 }
func genFn0(p0 uint32, p1 int, p2 bool) uint16 {
    p2 = false
    fmt.Println(p0)
    if (4294967295 != 18446744073709551615) {
        if (788 >= 787) {
            for fi4 := 0; fi4 < 6; fi4++ {
                v5 := []int{p1, 32768}
                _ = len(v5)
                _ = fi4
                _ = v5
            }
        }
    }
    v6 := []int{(p1 / ((p1) | 1)), p1, (p1 ^ p1)}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 2
}

func main() {
    var v1 uint = 2
    var v2 uint64 = 531
    v3 := make(map[int]string)
    v7 := []int{128, -2147483648, 0}
    v1 -= 255
    v1 = (v1 & (v1 / ((v1) | 1)))
    switch ("foo bar" + "hello") {
    case "a":
        if (255 < 209) {
            fmt.Println(v2, v2)
        } else {
            v2 = (v2 % (((v2 * 1)) | 1))
            if true {
                v2 = v2
                for fi8 := 0; fi8 < 12; fi8++ {
                    _ = len(v3)
                    v9 := GenS1{F0: (^(v1)), F1: "foo bar", F2: (^(fi8))}
                    _ = fi8
                    _ = v9
                }
            } else {
                if true {
                    v10 := GenS2{F0: (42 == v1)}
                    _ = v10
                } else {
                    if false {
                        var v11 int = -1
                        var v12 float32 = (((((1e10 - 1e10) * float32((((1.5 + float64((((1.5 + 0.0) / 1.5) - (0.0 - 3.14159)))) + 0.1) * 1e10))) - 0.1) * -2.25) / ((1e10 + -2.25) + float32(1e10)))
                        switch (("hello" + "") + "\n\t") {
                        case "a":
                            fmt.Println(v12, v12)
                        case "hello":
                            v13 := v3[v11]
                            _ = v13
                        case "foo bar":
                            fmt.Println(v2, v12, v2)
                        }
                        _ = v11
                        _ = v12
                    } else {
                        v14 := GenS3{F0: ""}
                        if (("hello" + "hello") != "foo bar") {
                            v1 = v1
                            v7 = append(v7, 0)
                        }
                        _ = v14
                    }
                }
            }
            delete(v3, 128)
        }
    case "":
        if false {
            v15 := v3[925]
            for fi16 := 0; fi16 < 11; fi16++ {
                if true {
                    v17 := GenS3{F0: (("\n\t" + "hello") + "\n\t")}
                    _ = v17
                } else {
                    v15 = ("hello" + v15)
                    if true {
                        v18 := GenS3{F0: (v15 + "foo bar")}
                        v7 = append(v7, (fi16 + fi16))
                        _ = v18
                    }
                }
                v19 := GenS3{F0: "\n\t"}
                _ = fi16
                _ = v19
            }
            fmt.Println(v2, v15, v2)
            _ = v15
        } else {
            if true {
                v20 := GenS3{F0: "foo bar"}
                var v21 uint8 = 127
                _ = v20
                _ = v21
            } else {
                if true {
                    v22 := v3[42]
                    switch v22 {
                    case "hello":
                        if (-1 != 0) {
                            v2 = (v2 % ((4294967295) | 1))
                            v23 := GenS4{F0: -2147483648, F1: 1}
                            _ = v23
                        }
                    }
                    _ = v22
                } else {
                    var v24 int16 = 0
                    _ = v24
                }
                var v25 int32 = 568
                _ = v25
            }
            v2 = 1
        }
    default:
        v1 += 634
    }
    if false {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    }
    for fi26 := 0; fi26 < 12; fi26++ {
        var v27 uint16 = 32767
        delete(v3, (fi26 % ((fi26) | 1)))
        var v28 int8 = 42
        _ = fi26
        _ = v27
        _ = v28
    }
    if true {
        var v29 int32 = 386
        v2 = (v2 >> ((v1) & 63))
        _ = v29
    }
    v1 += (v1 - (v1 << ((v1) & 63)))
    var v30 uint32 = 2147483648
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v7
}
