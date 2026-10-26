package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 uint64 }
type GenS2 struct { F0 uint; F1 uint8; F2 int }
type GenS3 struct { F0 int8; F1 int16; F2 string }
func genFn0(p0 int32) int32 {
    fmt.Println(p0)
    var v4 int8 = 42
    p0 -= (p0 * 2147483647)
    if (false || (1000 <= -1)) {
        for fi5 := 0; fi5 < 20; fi5++ {
            var v6 uint = 255
            if false {
                v4 = v4
                fmt.Println(p0, p0)
            } else {
                v6 += 42
                var v7 int16 = 663
                _ = v7
            }
            _ = fi5
            _ = v6
        }
        v8 := GenS1{F0: 1, F1: 375, F2: 42}
        _ = v8
    } else {
        for fi9 := 0; fi9 < 5; fi9++ {
            v4 = v4
            v10 := []int{-2147483648}
            _ = fi9
            _ = v10
        }
        if false {
            v11 := GenS2{F0: 0, F1: 1, F2: -828}
            _ = v11
        }
    }
    _ = p0
    _ = v4
    return p0
}

func main() {
    v1 := make(map[string]int)
    var v2 uint = 255
    var v3 string = ("foo bar" + "foo bar")
    v3 = v3
    switch ("\n\t" + "hello") {
    case "foo bar":
        for fi12 := 0; fi12 < 9; fi12++ {
            v2 = v2
            _ = fi12
        }
    default:
        for fi13 := 0; fi13 < 6; fi13++ {
            var v14 uint64 = 2
            if false {
                v2 -= v2
            } else {
                _ = len(v1)
                v14 += v14
            }
            _ = fi13
            _ = v14
        }
    }
    if (false || true) {
        v15 := v1[v3]
        _ = v15
    } else {
        v3 = (v3 + ((("a" + "foo bar") + "\n\t") + (v3 + (v3 + (v3 + v3)))))
    }
    var v16 int32 = 1
    for fi17 := 0; fi17 < 2; fi17++ {
        var v18 string = ((v3 + v3) + v3)
        _ = fi17
        _ = v18
    }
    v19 := []int{}
    var v20 uint16 = 65535
    var v21 uint32 = 42
    v16 = (v16 >> ((uint(v21)) & 63))
    if (-937 <= 32767) {
        for fi22 := 0; fi22 < 17; fi22++ {
            v23 := v1[v3]
            v1[v3] = v23
            _ = fi22
            _ = v23
        }
        var v24 uint16 = (v20 % ((v20) | 1))
        _ = v24
    }
    v25 := GenS2{F0: (v2 / ((v2) | 1)), F1: 205, F2: 829}
    _ = len(v19)
    v26 := GenS2{F0: (v2 ^ 2), F1: 0, F2: 2147483647}
    for fi27 := 0; fi27 < 18; fi27++ {
        switch v3 {
        case "":
            v20 -= (v20 ^ (^(v20)))
        }
        var v28 int64 = 9223372036854775807
        _ = fi27
        _ = v28
    }
    for fi29 := 0; fi29 < 16; fi29++ {
        for fi30 := 0; fi30 < 3; fi30++ {
            v1[v3] = (fi30 << ((v2) & 63))
            if true {
                var v31 int8 = 100
                var v32 int8 = -1
                var v33 float32 = 1e10
                _ = v31
                _ = v32
                _ = v33
            } else {
                v2 = v2
                var v34 int64 = -863
                if (!(false && true)) {
                    var v35 int16 = 0
                    switch ("hello" + v3) {
                    case "\n\t":
                        if true {
                            if ((^(v35)) >= (^(v35))) {
                                if ("hello" < v3) {
                                    var v36 float32 = 1e10
                                    var v37 uint = 42
                                    v38 := GenS3{F0: -128, F1: (v35 | -32768), F2: ""}
                                    v35 += v35
                                    _ = v36
                                    _ = v37
                                    _ = v38
                                }
                                var v39 uint64 = 289
                                fmt.Println(v35, v35)
                                _ = v39
                            } else {
                                fmt.Println(v3, v2)
                            }
                            var v40 uint8 = 1
                            fmt.Println(v34)
                            if true {
                                var v41 uint32 = 894
                                _ = v41
                            }
                            switch "foo bar" {
                            case "foo bar":
                                v1[v3] = fi29
                            default:
                                fmt.Println(v40, v2)
                            }
                            v1[(((v3 + "a") + ("" + "foo bar")) + (v3 + (v3 + v3)))] = -955
                            fmt.Println(fi30, v21, v16)
                            _ = len(v19)
                            switch v3 {
                            case "hello":
                                v19 = append(v19, (fi30 << ((v2) & 63)))
                            default:
                                var v42 uint16 = (v20 * 0)
                                _ = v42
                            }
                            _ = v40
                        } else {
                            v34 -= (v34 + v34)
                        }
                    case "a":
                        fmt.Println(v34)
                    default:
                        v2 = v2
                    }
                    _ = v35
                }
                _ = v34
            }
            v2 = (v2 >> ((v2) & 63))
            _ = fi30
        }
        var v43 float64 = (3.14159 - -2.25)
        v19 = append(v19, (fi29 + (fi29 & fi29)))
        _ = fi29
        _ = v43
    }
    for fi44 := 0; fi44 < 8; fi44++ {
        for fi45 := 0; fi45 < 3; fi45++ {
            v46 := v1["\n\t"]
            var v47 uint32 = v21
            v47 -= 65535
            _ = fi45
            _ = v46
            _ = v47
        }
        _ = fi44
    }
    v16 -= int32(v20)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v26
    _ = v3
}
