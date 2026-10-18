package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 uint8; F2 int32 }
type GenS3 struct { F0 int64; F1 uint32 }
func genFn0() uint64 {
    v6 := []int{32767}
    v7 := GenS1{F0: -86}
    _ = v6
    _ = v7
    return 291
}
func genFn1(p0 uint8, p1 int64) bool {
    var v8 bool = (!true)
    if (!false) {
        for fi9 := 0; fi9 < 3; fi9++ {
            var v10 uint = 255
            switch (fi9 << (((v10 - v10)) & 63)) {
            case 32768:
                p1 = (p1 ^ 897)
            default:
                if v8 {
                    switch (-(fi9)) {
                    case 127:
                        for fi11 := 0; fi11 < 1; fi11++ {
                            v12 := GenS1{F0: fi9}
                            switch ((((("hello" + (("a" + ("\n\t" + "\n\t")) + ("foo bar" + ("a" + ("a" + "foo bar"))))) + "hello") + "\n\t") + "") + "a") {
                            case "":
                                switch "foo bar" {
                                case "hello":
                                    if ((!v8) || v8) {
                                        p1 += 0
                                    }
                                    v13 := GenS1{F0: fi9}
                                    _ = v13
                                case "\n\t":
                                    v14 := GenS1{F0: int(p1)}
                                    _ = v14
                                default:
                                    v15 := []int{}
                                    _ = v15
                                }
                            case "hello":
                                if (0 < 1) {
                                    p1 = 9223372036854775807
                                    var v16 uint64 = genFn0()
                                    v8 = (((v8 || (0 >= 32767)) && (1 >= 100)) || (296 >= 32767))
                                    _ = v16
                                }
                            default:
                                p0 += (p0 << (((v10 % ((v10) | 1))) & 63))
                            }
                            _ = fi11
                            _ = v12
                        }
                    }
                    var v17 uint = 884
                    switch "hello" {
                    case "hello":
                        fmt.Println(v10, v17, v17)
                    default:
                        for fi18 := 0; fi18 < 7; fi18++ {
                            fmt.Println(v10, p1, fi9)
                            if true {
                                fmt.Println(fi18, v17)
                                if (301 < 0) {
                                    var v19 int32 = 802
                                    switch ("hello" + "a") {
                                    case "hello":
                                        v19 -= (v19 ^ (v19 * (v19 + 1)))
                                    default:
                                        fmt.Println(v19, fi9, v8)
                                    }
                                    v17 = v17
                                    v20 := []int{}
                                    _ = v19
                                    _ = v20
                                }
                                var v21 uint8 = (p0 & 61)
                                _ = v21
                            }
                            var v22 int8 = -117
                            _ = fi18
                            _ = v22
                        }
                    }
                    _ = v17
                }
            }
            v23 := GenS2{F0: p0, F1: 1, F2: -590}
            _ = fi9
            _ = v10
            _ = v23
        }
    }
    var v24 uint = 1
    _ = p0
    _ = p1
    _ = v24
    _ = v8
    return (!v8)
}

func main() {
    var v1 uint16 = 139
    v2 := make(map[string]bool)
    v3 := make(map[string]int)
    var v4 int32 = 42
    var v5 uint = 2
    v4 = v4
    if (!true) {
        var v25 string = ""
        _ = v25
    }
    v5 -= uint(v1)
    if false {
        var v26 string = ("a" + ("foo bar" + "\n\t"))
        _ = v26
    } else {
        switch ("foo bar" + "a") {
        case "":
            for fi27 := 0; fi27 < 7; fi27++ {
                v5 += v5
                _ = fi27
            }
        default:
            if ((v4 == v4) && false) {
                v5 = v5
                if ((v4 % ((v4) | 1)) < v4) {
                    if ((0 <= 0) || (32767 < -32768)) {
                        fmt.Println(v5, v1, v1)
                        v28 := v2[(("hello" + "hello") + ("a" + ""))]
                        _ = v28
                    } else {
                        var v29 int64 = 241
                        v30 := GenS3{F0: int64(v4), F1: uint32(v1)}
                        v5 = (v5 ^ v5)
                        _ = v29
                        _ = v30
                    }
                    var v31 bool = false
                    _ = v31
                } else {
                    v32 := []int{42}
                    _ = v32
                }
            }
        }
        var v33 int64 = 0
        v4 -= v4
        _ = v33
    }
    for fi34 := 0; fi34 < 8; fi34++ {
        v35 := GenS1{F0: fi34}
        v1 = v1
        var v36 uint64 = uint64(v1)
        _ = fi34
        _ = v35
        _ = v36
    }
    switch 108 {
    case -1:
        var v37 uint = v5
        _ = v37
    case 127:
        var v38 float64 = ((1.5 * 1.5) - 0.1)
        _ = v38
    case -6:
        v3["hello"] = -9223372036854775808
    }
    delete(v3, "\n\t")
    v4 -= v4
    v4 += -2147483648
    v39 := GenS1{F0: -952}
    v1 = v1
    v40 := []int{2147483647}
    v4 -= (^(v4))
    var v41 int = -53
    var v42 uint32 = 2
    for fi43 := 0; fi43 < 8; fi43++ {
        if false {
            v44 := GenS1{F0: 128}
            _ = v44
        }
        _ = fi43
    }
    v5 += (^(v5))
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
}
