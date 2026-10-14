package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int64 }
func genFn0() string {
    var v5 int = -9223372036854775808
    v6 := GenS1{F0: -523, F1: ("foo bar" + "\n\t"), F2: 0}
    var v7 uint64 = 4294967295
    v8 := GenS1{F0: 257, F1: (("\n\t" + "hello") + ("a" + (("a" + "foo bar") + ("hello" + ("" + "a"))))), F2: (v7 | (^(v7)))}
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return "a"
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint = 903
    v3 := make(map[string]string)
    v4 := []string{("" + ""), "", "", ("" + (("foo bar" + "") + ("\n\t" + "foo bar")))}
    var v9 int8 = -1
    if (0.0 < 1e10) {
        var v10 float64 = 1e10
        if len(v4) > 0 {
            v4[0] = "\n\t"
        }
        _ = v10
    } else {
        v11 := GenS1{F0: 1, F1: "", F2: 42}
        v9 = (v9 ^ v9)
        for fi12 := 0; fi12 < 12; fi12++ {
            v2 = (v2 + 930)
            fmt.Println(v2)
            for fi13 := 0; fi13 < 18; fi13++ {
                fmt.Println(fi13, v2)
                if ((v9 * 42) != v9) {
                    var v14 uint64 = 2
                    if false {
                        v15 := GenS1{F0: -326, F1: "", F2: (v14 & v14)}
                        if (fi12 <= 42) {
                            v3[(("a" + ("foo bar" + "")) + "")] = "a"
                            v16 := GenS2{F0: (128 <= 1)}
                            v9 = v9
                            _ = v16
                        }
                        v9 -= (v9 / ((v9) | 1))
                        _ = v15
                    } else {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        if false {
                            v17 := v3["\n\t"]
                            _ = v17
                        } else {
                            v18 := GenS3{F0: 1}
                            _ = v18
                        }
                        fmt.Println(fi12)
                    }
                    _ = v14
                }
                _ = fi13
            }
            _ = fi12
        }
        _ = v11
    }
    switch "" {
    case "a":
        var v19 int = 9223372036854775807
        _ = v19
    case "hello":
        if true {
            var v20 uint32 = 1
            _ = v20
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    }
    v2 += v2
    v21 := GenS4{F0: "foo bar"}
    for fi22 := 0; fi22 < 19; fi22++ {
        for fi23 := 0; fi23 < 14; fi23++ {
            v9 -= (v9 / ((v9) | 1))
            _ = fi23
        }
        var v24 uint8 = 127
        _ = fi22
        _ = v24
    }
    if false {
        var v25 bool = true
        _ = v25
    }
    for fi26 := 0; fi26 < 18; fi26++ {
        v2 = v2
        v27 := GenS2{F0: false}
        if (!(false || true)) {
            for fi28 := 0; fi28 < 7; fi28++ {
                v2 += 4294967295
                fmt.Println(v2, v9, fi26)
                var v29 uint8 = 1
                _ = fi28
                _ = v29
            }
        } else {
            if len(v4) > 0 {
                v4[0] = ("" + ((((((("a" + ("\n\t" + "foo bar")) + "\n\t") + "") + "\n\t") + "hello") + "\n\t") + (("\n\t" + "a") + "foo bar")))
            }
            switch (-(fi26)) {
            case -1:
                v30 := GenS1{F0: 1, F1: ("foo bar" + ("\n\t" + "hello")), F2: 18446744073709551615}
                _ = v30
            case 9223372036854775807:
                for fi31 := 0; fi31 < 18; fi31++ {
                    if (-2.25 < 0.1) {
                        fmt.Println(fi26)
                        if (42 < 42) {
                            v32 := GenS4{F0: "foo bar"}
                            _ = v32
                        } else {
                            v33 := GenS4{F0: (("\n\t" + ("" + ("a" + "hello"))) + ((("" + "") + ("foo bar" + "a")) + "hello"))}
                            v2 = (^(v2))
                            _ = v33
                        }
                        if ((!(2147483647 < 0)) || (!false)) {
                            var v34 float64 = 1.5
                            if (fi31 <= (fi26 * 1)) {
                                v9 -= v9
                            }
                            _ = v34
                        }
                    } else {
                        v35 := GenS5{F0: 21}
                        v36 := GenS5{F0: 9223372036854775807}
                        if len(v4) > 0 {
                            v4[0] = "hello"
                        }
                        _ = v35
                        _ = v36
                    }
                    _ = fi31
                }
            case 4294967295:
                if (!true) {
                    v4 = append(v4, "hello")
                }
            }
        }
        _ = fi26
        _ = v27
    }
    var v37 uint8 = 127
    if (2147483647 < -716) {
        v38 := GenS2{F0: (false && ((3.14159 < -2.25) || true))}
        fmt.Println(v37, v37)
        v39 := GenS3{F0: 395}
        _ = v38
        _ = v39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v37
    _ = v4
    _ = v9
}
