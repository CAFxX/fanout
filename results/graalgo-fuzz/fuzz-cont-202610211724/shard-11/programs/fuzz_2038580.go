package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string; F1 string }

func main() {
    v1 := []int{}
    var v2 int16 = -687
    v3 := make(map[string]string)
    var v4 bool = (65535 >= 255)
    var v5 float32 = (1e10 - 1.5)
    if (4294967295 < 32767) {
        for fi6 := 0; fi6 < 20; fi6++ {
            if (100 >= 127) {
                fmt.Println(v5, v2, v2)
                if v4 {
                    v2 += v2
                    v2 = v2
                }
                var v7 int = (fi6 & (fi6 | fi6))
                _ = v7
            } else {
                v8 := GenS1{F0: 18446744073709551615, F1: (4294967295 == 2147483647)}
                for fi9 := 0; fi9 < 1; fi9++ {
                    var v10 int16 = -243
                    fmt.Println(fi9, fi6)
                    v11 := GenS1{F0: 4294967295, F1: false}
                    v12 := GenS1{F0: 18446744073709551615, F1: v4}
                    if (255 != 1) {
                        _ = len(v1)
                        v13 := GenS2{F0: ("\n\t" + "foo bar")}
                        _ = v13
                    } else {
                        v14 := GenS3{F0: "\n\t", F1: "\n\t"}
                        v5 -= ((v5 - v5) + (float32((3.14159 + (0.0 - 0.1))) - ((float32(1.5) * 1e10) - v5)))
                        if (v4 || (fi6 >= 0)) {
                            switch "foo bar" {
                            case "foo bar":
                                var v15 float64 = 0.0
                                _ = v15
                            case "\n\t":
                                var v16 bool = v4
                                _ = v16
                            case "hello":
                                fmt.Println(fi6, v5, fi9)
                            }
                            var v17 bool = v4
                            _ = v17
                        }
                        _ = v14
                    }
                    _ = fi9
                    _ = v10
                    _ = v11
                    _ = v12
                }
                _ = v8
            }
            var v18 uint64 = 613
            _ = fi6
            _ = v18
        }
    } else {
        fmt.Println(v5, v2)
        fmt.Println(v4)
    }
    for fi19 := 0; fi19 < 9; fi19++ {
        v20 := GenS1{F0: 4294967295, F1: ("\n\t" == "a")}
        for fi21 := 0; fi21 < 18; fi21++ {
            if (255 == 127) {
                v22 := v3["a"]
                var v23 uint32 = 0
                var v24 float64 = (-2.25 * 1e10)
                var v25 string = ("a" + v22)
                _ = v22
                _ = v23
                _ = v24
                _ = v25
            } else {
                var v26 int16 = (v2 / ((924) | 1))
                var v27 int16 = -1
                _ = v26
                _ = v27
            }
            v3[""] = "foo bar"
            var v28 int16 = 32767
            _ = fi21
            _ = v28
        }
        _ = fi19
        _ = v20
    }
    v3["\n\t"] = ("" + ("hello" + "\n\t"))
    if (v4 || v4) {
        v29 := GenS3{F0: "", F1: "\n\t"}
        for fi30 := 0; fi30 < 14; fi30++ {
            v4 = (-9223372036854775808 == -9223372036854775808)
            _ = fi30
        }
        _ = v29
    }
    switch "" {
    case "\n\t":
        if (false || v4) {
            v4 = ((v5 == float32(-2.25)) || (853 > 0))
        } else {
            for fi31 := 0; fi31 < 19; fi31++ {
                for fi32 := 0; fi32 < 19; fi32++ {
                    v2 = (v2 ^ 1000)
                    _ = fi32
                }
                _ = fi31
            }
            v4 = (v4 || true)
            var v33 int8 = 100
            _ = v33
        }
    case "":
        v4 = true
    }
    if (!true) {
        v3["foo bar"] = "a"
    }
    for fi34 := 0; fi34 < 4; fi34++ {
        v4 = v4
        _ = fi34
    }
    v2 = (v2 + v2)
    _ = len(v1)
    v35 := GenS3{F0: "\n\t", F1: "hello"}
    if (!v4) {
        var v36 uint64 = 953
        v36 = v36
        _ = v36
    }
    for fi37 := 0; fi37 < 14; fi37++ {
        var v38 int8 = -128
        _ = fi37
        _ = v38
    }
    v39 := GenS2{F0: ("\n\t" + "")}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v39
    _ = v4
    _ = v5
}
