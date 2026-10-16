package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 int16; F1 int }
type GenS4 struct { F0 int64; F1 uint8; F2 uint8 }
func genFn0(p0 int16, p1 uint16, p2 uint32) int {
    switch ("a" + ("a" + ("" + "\n\t"))) {
    case "foo bar":
        if true {
            var v5 int16 = int16(p2)
            _ = v5
        }
    default:
        var v6 int64 = -9223372036854775808
        _ = v6
    }
    var v7 uint64 = 4294967295
    if false {
        if (0 > -9223372036854775808) {
            var v8 uint64 = v7
            v9 := GenS1{F0: -994}
            _ = v8
            _ = v9
        } else {
            v10 := GenS1{F0: 0}
            v11 := []int{9223372036854775807, -105}
            _ = v10
            _ = v11
        }
    } else {
        v12 := []int{-452}
        _ = v12
    }
    var v13 string = ("a" + "foo bar")
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v7
    return 0
}

func main() {
    v1 := make(map[string]int)
    var v2 int = -2147483648
    var v3 string = "foo bar"
    v4 := []string{v3}
    v2 = 42
    _ = len(v1)
    _ = len(v1)
    for fi14 := 0; fi14 < 16; fi14++ {
        var v15 int = (v2 % ((genFn0(0, 1, 446)) | 1))
        var v16 uint32 = 1
        _ = fi14
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 18; fi17++ {
        if len(v4) > 0 {
            v4[0] = v3
        }
        if len(v4) > 0 {
            v4[0] = v3
        }
        var v18 uint64 = 42
        _ = fi17
        _ = v18
    }
    v19 := GenS2{F0: 0, F1: 0}
    v3 = (("hello" + "hello") + v3)
    switch v3 {
    case "foo bar":
        fmt.Println(v2, v3)
    case "\n\t":
        if true {
            if (true && false) {
                var v20 uint16 = 400
                for fi21 := 0; fi21 < 4; fi21++ {
                    for fi22 := 0; fi22 < 16; fi22++ {
                        switch ("hello" + "hello") {
                        case "\n\t":
                            v23 := GenS2{F0: 0, F1: -2147483648}
                            _ = v23
                        case "foo bar":
                            v2 = fi21
                        case "hello":
                            v24 := GenS1{F0: -2147483648}
                            _ = v24
                        }
                        _ = fi22
                    }
                    _ = fi21
                }
                _ = v20
            }
            if (v3 < v3) {
                if true {
                    v3 = v3
                    switch v3 {
                    case "":
                        var v25 int16 = 0
                        _ = v25
                    case "hello":
                        v2 = -625
                    }
                    var v26 uint = 18446744073709551615
                    _ = v26
                } else {
                    var v27 int16 = 1000
                    if ("foo bar" < ("" + (("\n\t" + v3) + "\n\t"))) {
                        if false {
                            var v28 int = (v2 + (v2 | (v2 + (v2 ^ 310))))
                            v29 := GenS1{F0: 1}
                            v28 = genFn0(1, 2, 2147483647)
                            _ = v28
                            _ = v29
                        }
                        v1[(v3 + ((v3 + v3) + v3))] = -1
                    }
                    v30 := GenS3{F0: 1000, F1: (-(v2))}
                    _ = v27
                    _ = v30
                }
            } else {
                var v31 int8 = 42
                switch v3 {
                case "hello":
                    var v32 uint16 = 255
                    _ = v32
                case "\n\t":
                    _ = len(v1)
                default:
                    v33 := GenS3{F0: 221, F1: 128}
                    _ = v33
                }
                _ = v31
            }
        } else {
            for fi34 := 0; fi34 < 14; fi34++ {
                v2 = (fi34 - fi34)
                _ = fi34
            }
        }
    case "a":
        v3 = (v3 + v3)
    }
    for fi35 := 0; fi35 < 9; fi35++ {
        v2 += (v2 & 0)
        _ = fi35
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v2 -= v2
    v3 = v3
    v36 := GenS4{F0: 42, F1: 8, F2: 0}
    if (true && (v2 != (v2 / (((^(v2))) | 1)))) {
        v2 -= (v2 >> ((65535) & 63))
        v1["\n\t"] = (v2 >> ((549) & 63))
    }
    for fi37 := 0; fi37 < 9; fi37++ {
        v4 = append(v4, (v3 + v3))
        _ = fi37
    }
    switch "foo bar" {
    case "\n\t":
        for fi38 := 0; fi38 < 13; fi38++ {
            v3 = v3
            _ = fi38
        }
    case "":
        var v39 uint32 = 539
        _ = v39
    case "hello":
        var v40 uint = 271
        _ = v40
    default:
        for fi41 := 0; fi41 < 13; fi41++ {
            var v42 uint64 = 1
            v3 = v3
            v43 := GenS2{F0: 1, F1: 955}
            _ = fi41
            _ = v42
            _ = v43
        }
    }
    var v44 uint64 = 18446744073709551615
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v44
}
