package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 int8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 bool; F1 string; F2 uint }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int; F1 string }

func main() {
    var v1 uint16 = 0
    v2 := make(map[string]bool)
    v3 := make(map[int]int)
    v4 := []string{"hello", (("hello" + ("" + "\n\t")) + ("foo bar" + "foo bar"))}
    var v5 uint64 = 18446744073709551615
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 int64 = 72
        var v8 uint64 = 2
        _ = fi6
        _ = v7
        _ = v8
    }
    v1 += v1
    if (true || false) {
        if (!(42 < 2147483647)) {
            v1 += 32767
            var v9 string = ""
            if false {
                v1 += 1
                if true {
                    for fi10 := 0; fi10 < 7; fi10++ {
                        if false {
                            var v11 float64 = -2.25
                            if (!(-869 >= 454)) {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            } else {
                                var v12 string = v9
                                _ = v12
                            }
                            _ = v11
                        }
                        _ = fi10
                    }
                } else {
                    v5 = (v5 * 910)
                    v13 := GenS1{F0: -2147483648, F1: 865, F2: 42}
                    _ = v13
                }
            }
            _ = v9
        }
        if true {
            v1 = v1
        } else {
            v14 := GenS2{F0: 926}
            if len(v4) > 0 {
                v4[0] = ("foo bar" + "hello")
            }
            _ = v14
        }
        v15 := GenS3{F0: false, F1: "foo bar", F2: 2}
        _ = v15
    } else {
        var v16 uint8 = 58
        _ = v16
    }
    v17 := GenS1{F0: 2147483647, F1: -1, F2: 127}
    v5 = v5
    switch "a" {
    case "\n\t":
        for fi18 := 0; fi18 < 12; fi18++ {
            var v19 int64 = 0
            _ = fi18
            _ = v19
        }
    case "hello":
        var v20 uint64 = 178
        _ = v20
    default:
        if true {
            var v21 int8 = 127
            _ = v21
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v22 bool = true
            _ = v22
        }
    }
    if len(v4) > 0 {
        v4[0] = ("hello" + ("foo bar" + ""))
    }
    v1 -= 638
    v1 = 65535
    for fi23 := 0; fi23 < 14; fi23++ {
        delete(v2, "foo bar")
        if false {
            for fi24 := 0; fi24 < 5; fi24++ {
                v25 := GenS4{F0: (("\n\t" + "hello") + "\n\t")}
                var v26 uint64 = 18446744073709551615
                _ = fi24
                _ = v25
                _ = v26
            }
            v1 += (v1 / ((v1) | 1))
        } else {
            v3[fi23] = -57
        }
        _ = fi23
    }
    _ = len(v2)
    v27 := GenS5{F0: 0, F1: (("foo bar" + ("a" + ("" + (("\n\t" + (("" + ((("" + ("a" + (("" + "") + ("a" + "hello")))) + "a") + ("foo bar" + "hello"))) + ("a" + ""))) + "")))) + "a")}
    v28 := GenS5{F0: 1, F1: "\n\t"}
    var v29 uint32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
