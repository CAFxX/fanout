package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 int64; F2 uint8 }
type GenS3 struct { F0 int32; F1 uint8 }
func genFn0(p0 int, p1 int64, p2 uint32) int {
    fmt.Println(p1)
    switch 32767 {
    case -1:
        p1 = (p1 ^ p1)
    }
    var v6 int8 = -128
    v7 := []int{32768}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return p0
}
func genFn1(p0 int8) uint16 {
    var v8 uint8 = 128
    v9 := GenS1{F0: 4294967295}
    v8 = 2
    _ = p0
    _ = v8
    _ = v9
    return 1
}

func main() {
    v1 := []string{"foo bar", ("" + (((("hello" + "a") + ("a" + "a")) + "hello") + "\n\t"))}
    v2 := []string{}
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]bool)
    var v5 uint32 = 2147483647
    var v10 int32 = -1
    switch "a" {
    case "hello":
        var v11 string = ("" + "hello")
        _ = v11
    case "\n\t":
        _ = len(v1)
    case "foo bar":
        if true {
            var v12 uint8 = 2
            var v13 uint32 = (^(v5))
            v14 := GenS2{F0: 42, F1: -1, F2: 1}
            _ = v12
            _ = v13
            _ = v14
        } else {
            fmt.Println(v5, v10)
        }
    default:
        for fi15 := 0; fi15 < 2; fi15++ {
            if (-32768 >= 32767) {
                if false {
                    v10 += (v10 - int32(fi15))
                    fmt.Println(fi15, v5)
                } else {
                    v10 += v10
                }
                var v16 float32 = 1.5
                for fi17 := 0; fi17 < 1; fi17++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v18 uint = 255
                    _ = fi17
                    _ = v18
                }
                _ = v16
            }
            _ = fi15
        }
    }
    switch "\n\t" {
    case "a":
        _ = len(v1)
    case "":
        if false {
            for fi19 := 0; fi19 < 15; fi19++ {
                v5 = (v5 / (((v5 | (v5 * v5))) | 1))
                var v20 uint8 = 1
                var v21 uint8 = (v20 % ((0) | 1))
                _ = fi19
                _ = v20
                _ = v21
            }
        } else {
            v5 = (^(v5))
            if false {
                if (false || true) {
                    if (v10 < 0) {
                        var v22 uint = 0
                        _ = v22
                    }
                    var v23 int8 = 100
                    var v24 uint8 = 188
                    _ = v23
                    _ = v24
                } else {
                    fmt.Println(v10, v10)
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v10 -= (v10 >> ((4294967295) & 63))
                }
                var v25 uint64 = 18446744073709551615
                _ = v25
            } else {
                var v26 int16 = 42
                var v27 uint16 = 255
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = v26
                _ = v27
            }
        }
    case "\n\t":
        for fi28 := 0; fi28 < 7; fi28++ {
            v5 = v5
            _ = fi28
        }
    default:
        var v29 float32 = (3.14159 * 3.14159)
        _ = v29
    }
    var v30 uint = 0
    if (!((true && (32767 != 2147483647)) || (1 > (v5 | v5)))) {
        for fi31 := 0; fi31 < 3; fi31++ {
            if ((-598 <= v10) || false) {
                v32 := GenS1{F0: genFn0(-1, -9223372036854775808, v5)}
                _ = v32
            } else {
                var v33 int64 = -709
                if ((v33 / ((-934) | 1)) >= v33) {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    if ((fi31 != (fi31 + (-(fi31)))) && false) {
                        v2 = append(v2, "foo bar")
                    }
                    for fi34 := 0; fi34 < 20; fi34++ {
                        if false {
                            v30 += (v30 / (((v30 << ((v30) & 63))) | 1))
                        }
                        if false {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            if false {
                                if (true && true) {
                                    v33 -= (v33 & (v33 << (((v30 % (((v30 ^ 958)) | 1))) & 63)))
                                    fmt.Println(v10, fi34)
                                    v35 := GenS1{F0: (^(fi31))}
                                    _ = v35
                                }
                                var v36 uint8 = 127
                                _ = v36
                            }
                        }
                        _ = fi34
                    }
                }
                for fi37 := 0; fi37 < 6; fi37++ {
                    v5 += 42
                    if (0.0 >= 0.1) {
                        var v38 float32 = (0.1 - (0.0 - 1e10))
                        _ = v38
                    }
                    v5 += (^(v5))
                    var v39 string = ""
                    _ = fi37
                    _ = v39
                }
                _ = v33
            }
            v40 := GenS3{F0: v10, F1: 36}
            var v41 float32 = 0.1
            _ = fi31
            _ = v40
            _ = v41
        }
        v42 := GenS1{F0: genFn0(127, -523, v5)}
        v43 := v4[v5]
        _ = v42
        _ = v43
    } else {
        v30 = (v30 ^ (v30 * (v30 & (v30 ^ (^(v30))))))
        v30 -= 0
        _ = len(v1)
    }
    if false {
        v44 := GenS2{F0: int8(v5), F1: 9223372036854775807, F2: 42}
        v2 = append(v2, "\n\t")
        _ = v44
    }
    for fi45 := 0; fi45 < 19; fi45++ {
        var v46 int32 = v10
        var v47 float32 = (1e10 + 0.1)
        _ = fi45
        _ = v46
        _ = v47
    }
    if (0.1 == 1.5) {
        var v48 uint8 = 42
        if len(v1) > 0 {
            v1[0] = "a"
        }
        _ = v48
    } else {
        for fi49 := 0; fi49 < 6; fi49++ {
            v50 := GenS3{F0: 800, F1: 2}
            _ = fi49
            _ = v50
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
