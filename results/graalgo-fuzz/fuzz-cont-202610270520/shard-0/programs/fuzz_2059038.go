package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16; F1 uint8; F2 string }
type GenS4 struct { F0 int }
func genFn0(p0 uint32, p1 string, p2 int8) int8 {
    switch ((p1 + p1) + "hello") {
    case "foo bar":
        var v7 bool = true
        _ = v7
    case "hello":
        p2 = p2
    case "\n\t":
        if (1 == 172) {
            fmt.Println(p2)
            fmt.Println(p2, p0, p2)
        }
    default:
        if false {
            var v8 uint8 = 0
            _ = v8
        } else {
            p1 = ((p1 + p1) + ("hello" + p1))
            var v9 int16 = -1
            _ = v9
        }
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    return p2
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := make(map[string]bool)
    var v3 bool = true
    var v4 uint = 213
    v5 := []string{"a", ("foo bar" + (("" + "a") + "hello")), ("hello" + (("\n\t" + (("a" + ("" + "a")) + (("hello" + "foo bar") + ""))) + ("a" + ""))), ""}
    var v6 int64 = 1
    v1 = v1
    for fi10 := 0; fi10 < 3; fi10++ {
        v11 := GenS1{F0: v1, F1: uint8(v4), F2: int32(fi10)}
        fmt.Println(v4, fi10)
        _ = fi10
        _ = v11
    }
    v12 := GenS1{F0: v1, F1: 42, F2: 42}
    v4 -= v4
    var v13 int16 = 22
    v3 = false
    var v14 float32 = (3.14159 - 0.0)
    var v15 uint16 = uint16(v1)
    v16 := GenS2{F0: 0}
    switch 849 {
    case 42:
        for fi17 := 0; fi17 < 6; fi17++ {
            switch (fi17 - (fi17 << (((v4 | uint(v15))) & 63))) {
            case 2147483647:
                if v3 {
                    var v18 int = fi17
                    _ = v18
                }
            }
            for fi19 := 0; fi19 < 3; fi19++ {
                v3 = v3
                v13 -= v13
                v15 -= 0
                _ = fi19
            }
            _ = fi17
        }
    case -2147483648:
        v20 := GenS2{F0: -128}
        _ = v20
    }
    var v21 int = 128
    if ((!(!v3)) && v3) {
        for fi22 := 0; fi22 < 6; fi22++ {
            var v23 int64 = v6
            for fi24 := 0; fi24 < 2; fi24++ {
                switch -619 {
                case 32768:
                    v21 = (^(fi24))
                case -740:
                    if v3 {
                        if v3 {
                            if ((v15 & v15) < v15) {
                                fmt.Println(v21)
                                if v3 {
                                    fmt.Println(v6, v23, v13)
                                    _ = len(v5)
                                    v6 = (v6 | (^(v23)))
                                }
                            } else {
                                v25 := GenS1{F0: 42, F1: 1, F2: -513}
                                var v26 bool = v3
                                _ = v25
                                _ = v26
                            }
                            v23 = -9223372036854775808
                            v15 = v15
                            if v3 {
                                v21 += (v21 & -264)
                                var v27 bool = ((v1 / ((v1) | 1)) != 123)
                                _ = v27
                            } else {
                                var v28 float64 = 3.14159
                                v21 += (fi22 - v21)
                                var v29 float64 = v28
                                _ = v28
                                _ = v29
                            }
                        }
                        v30 := GenS1{F0: uint64(v15), F1: 142, F2: 2147483647}
                        _ = v30
                    }
                case 128:
                    if v3 {
                        v31 := GenS3{F0: (v15 | v15), F1: 1, F2: (("a" + (("\n\t" + "hello") + "")) + "\n\t")}
                        _ = v31
                    }
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                var v32 uint32 = 42
                _ = fi24
                _ = v32
            }
            _ = fi22
            _ = v23
        }
    }
    if (-2147483648 != 1) {
        v2["foo bar"] = v3
        var v33 bool = ((!false) && v3)
        _ = v33
    } else {
        var v34 uint32 = 1
        _ = v34
    }
    v5 = append(v5, ("foo bar" + (("a" + ("foo bar" + "hello")) + ("a" + (("hello" + "") + "a")))))
    _ = len(v5)
    for fi35 := 0; fi35 < 11; fi35++ {
        v36 := GenS4{F0: (fi35 ^ (-(fi35)))}
        _ = fi35
        _ = v36
    }
    if true {
        fmt.Println(v1)
        v37 := GenS1{F0: v1, F1: 255, F2: 42}
        fmt.Println(v3, v14, v14)
        _ = v37
    } else {
        var v38 int8 = 42
        var v39 uint8 = 128
        v40 := GenS2{F0: (v38 - (^(v38)))}
        _ = v38
        _ = v39
        _ = v40
    }
    switch -1 {
    case 127:
        v41 := GenS4{F0: (v21 + (v21 - v21))}
        _ = v41
    case 568:
        v5 = append(v5, "a")
    case -704:
        if len(v5) > 0 {
            v5[0] = ("\n\t" + "hello")
        }
    default:
        v6 -= (v6 ^ 99)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
