package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 int8; F1 int64 }
type GenS4 struct { F0 uint16; F1 uint16; F2 int32 }
type GenS5 struct { F0 string; F1 uint32; F2 uint }
type GenS6 struct { F0 uint16 }
func genFn0() int16 {
    v5 := []int{310, -242, 1}
    var v6 uint64 = 866
    v5 = append(v5, 4294967295)
    var v7 int = -1
    _ = v5
    _ = v6
    _ = v7
    return 42
}

func main() {
    v1 := []int{9223372036854775807, -2147483648, -1, 0}
    v2 := []string{"a", ""}
    v3 := make(map[string]int)
    var v4 int32 = 215
    var v8 uint8 = 2
    var v9 string = "foo bar"
    v8 += (v8 | v8)
    v9 = (((v9 + ((v9 + (v9 + "a")) + (v9 + v9))) + (v9 + v9)) + "a")
    for fi10 := 0; fi10 < 19; fi10++ {
        fmt.Println(v8, v9)
        _ = fi10
    }
    switch 2147483647 {
    case -563:
        switch "hello" {
        case "":
            fmt.Println(v9, v4, v9)
        default:
            if true {
                if (true || ((false || ("hello" == (("hello" + (v9 + v9)) + v9))) || true)) {
                    _ = len(v2)
                    var v11 float64 = 0.0
                    _ = v11
                } else {
                    if (0.0 > (1e10 * 1.5)) {
                        v2 = append(v2, v9)
                    }
                    if (-1 != 227) {
                        v12 := GenS1{F0: 484, F1: (v4 & (v4 | (v4 + v4)))}
                        _ = v12
                    }
                    v8 -= (^(v8))
                }
                for fi13 := 0; fi13 < 10; fi13++ {
                    if true {
                        var v14 uint8 = (v8 / ((v8) | 1))
                        _ = v14
                    } else {
                        v15 := GenS2{F0: (fi13 % ((fi13) | 1)), F1: 127}
                        if true {
                            v2 = append(v2, v9)
                            if true {
                                v16 := GenS2{F0: (fi13 ^ fi13), F1: fi13}
                                _ = v16
                            }
                            var v17 int = (fi13 | 9223372036854775807)
                            _ = v17
                        } else {
                            var v18 int64 = 1
                            var v19 int16 = 53
                            _ = v18
                            _ = v19
                        }
                        _ = v15
                    }
                    _ = fi13
                }
            }
        }
    case 472:
        var v20 uint = 65535
        _ = v20
    default:
        if (genFn0() <= 0) {
            v21 := GenS3{F0: 0, F1: 42}
            for fi22 := 0; fi22 < 3; fi22++ {
                if true {
                    _ = len(v2)
                } else {
                    for fi23 := 0; fi23 < 17; fi23++ {
                        if true {
                            v4 += 0
                            var v24 bool = true
                            var v25 int64 = 42
                            if (true && false) {
                                v9 = v9
                                if true {
                                    switch fi22 {
                                    case -845:
                                        if len(v2) > 0 {
                                            v2[0] = (v9 + v9)
                                        }
                                        var v26 uint64 = 207
                                        _ = v26
                                    }
                                    switch (v9 + (v9 + v9)) {
                                    case "":
                                        v25 = v25
                                    case "foo bar":
                                        v27 := v3[(v9 + "foo bar")]
                                        _ = v27
                                    }
                                } else {
                                    v25 -= (v25 % ((0) | 1))
                                    var v28 int64 = -9223372036854775808
                                    _ = v28
                                }
                                v8 -= 255
                            } else {
                                var v29 int = 1
                                var v30 uint8 = (v8 >> ((983) & 63))
                                v31 := v3[v9]
                                _ = v29
                                _ = v30
                                _ = v31
                            }
                            _ = v24
                            _ = v25
                        } else {
                            v32 := GenS4{F0: 255, F1: 2, F2: (v4 ^ (-(v4)))}
                            v8 += v8
                            _ = v32
                        }
                        v9 = v9
                        switch (fi22 | fi23) {
                        case 332:
                            var v33 int = fi23
                            _ = v33
                        case -990:
                            switch fi22 {
                            case 32767:
                                delete(v3, ((v9 + "hello") + ((v9 + v9) + (v9 + (v9 + (("hello" + ((("\n\t" + ((("" + "hello") + v9) + ((("hello" + v9) + "a") + "hello"))) + ("" + (((("hello" + v9) + v9) + (v9 + ((v9 + v9) + v9))) + ("" + (v9 + ("a" + v9)))))) + v9)) + v9))))))
                            case 127:
                                v34 := GenS5{F0: "a", F1: 62, F2: 458}
                                _ = v34
                            }
                        }
                        _ = fi23
                    }
                    if false {
                        if true {
                            switch "foo bar" {
                            case "":
                                if false {
                                    var v35 int16 = 32767
                                    for fi36 := 0; fi36 < 13; fi36++ {
                                        fmt.Println(v9)
                                        v3[("\n\t" + (v9 + (("a" + ((v9 + ("a" + ((((v9 + ("" + v9)) + (v9 + v9)) + ("\n\t" + ("" + "foo bar"))) + (v9 + v9)))) + "\n\t")) + v9)))] = fi22
                                        v37 := GenS2{F0: (fi22 >> ((1) & 63)), F1: fi22}
                                        _ = fi36
                                        _ = v37
                                    }
                                    var v38 int32 = 2147483647
                                    _ = v35
                                    _ = v38
                                }
                            case "hello":
                                v9 = "hello"
                            }
                            fmt.Println(v8)
                            _ = len(v1)
                        }
                    }
                }
                var v39 int16 = int16(v8)
                _ = fi22
                _ = v39
            }
            v9 = v9
            _ = v21
        } else {
            var v40 uint32 = 42
            v9 = v9
            v3[v9] = 4294967295
            _ = v40
        }
    }
    for fi41 := 0; fi41 < 16; fi41++ {
        if true {
            v42 := GenS6{F0: 521}
            _ = v42
        }
        _ = fi41
    }
    if (32767 != 42) {
        switch ((v9 + ("a" + v9)) + ((v9 + v9) + ("\n\t" + v9))) {
        case "foo bar":
            var v43 uint16 = 983
            _ = v43
        case "\n\t":
            fmt.Println(v9)
        }
        var v44 int64 = -1
        var v45 int32 = v4
        _ = v44
        _ = v45
    } else {
        var v46 int32 = 2147483647
        for fi47 := 0; fi47 < 3; fi47++ {
            v4 += (^(v46))
            v4 += 42
            var v48 uint32 = 2
            _ = fi47
            _ = v48
        }
        _ = v46
    }
    v1 = append(v1, 0)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
