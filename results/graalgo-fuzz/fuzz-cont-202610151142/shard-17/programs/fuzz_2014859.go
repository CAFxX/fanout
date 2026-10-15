package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 uint8; F2 int8 }
type GenS3 struct { F0 bool; F1 int8; F2 uint8 }
type GenS4 struct { F0 uint16; F1 uint8; F2 uint64 }

func main() {
    var v1 int32 = -1
    var v2 int64 = -9223372036854775808
    var v3 float32 = (0.0 - float32(float64(3.14159)))
    v4 := make(map[int]bool)
    var v5 int8 = 0
    v6 := []string{"a", "a", ((("a" + "hello") + "a") + ((("a" + "a") + "foo bar") + "hello"))}
    switch (("\n\t" + ("\n\t" + "hello")) + ("\n\t" + "hello")) {
    case "":
        var v7 bool = false
        _ = v7
    case "a":
        var v8 int16 = -32768
        _ = v8
    default:
        v9 := GenS1{F0: "\n\t"}
        _ = v9
    }
    if (2 > 282) {
        fmt.Println(v3, v1, v3)
        var v10 int32 = -637
        _ = v10
    }
    if false {
        v5 -= (v5 & (^(v5)))
        if (602 < 65535) {
            var v11 uint64 = 1
            var v12 int8 = int8(v1)
            if false {
                if len(v6) > 0 {
                    v6[0] = "a"
                }
            }
            _ = v11
            _ = v12
        }
    }
    var v13 uint16 = 36
    var v14 int64 = -699
    v15 := v4[199]
    _ = len(v6)
    fmt.Println(v13)
    fmt.Println(v5, v15, v13)
    var v16 int64 = v2
    v15 = v15
    _ = len(v6)
    v14 += (-(v14))
    v6 = append(v6, "\n\t")
    v5 = (v5 | (v5 ^ (v5 + (v5 * (v5 ^ 1)))))
    var v17 int32 = (-(v1))
    if (!((false && false) || ((0 == 42) && (973 > 42)))) {
        if (4294967295 <= 65535) {
            if (!v15) {
                var v18 int16 = 685
                _ = v18
            } else {
                v3 = v3
                if v15 {
                    var v19 bool = ((^(v13)) > v13)
                    fmt.Println(v16, v5)
                    _ = v19
                }
            }
            if (v15 || true) {
                v15 = ((v5 * -1) <= (v5 % ((v5) | 1)))
                switch (("" + "hello") + ("a" + "\n\t")) {
                case "\n\t":
                    v4[-1] = (float64(v3) == float64((1.5 / (v3 / (-2.25 * v3)))))
                }
            }
            for fi20 := 0; fi20 < 15; fi20++ {
                v15 = (false || true)
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = fi20
            }
        }
        if v15 {
            v6 = append(v6, ("" + ""))
            for fi21 := 0; fi21 < 11; fi21++ {
                for fi22 := 0; fi22 < 10; fi22++ {
                    var v23 uint = 1
                    _ = fi22
                    _ = v23
                }
                for fi24 := 0; fi24 < 4; fi24++ {
                    v16 += -943
                    _ = fi24
                }
                _ = fi21
            }
            if v15 {
                v25 := GenS1{F0: "a"}
                v26 := GenS1{F0: ""}
                for fi27 := 0; fi27 < 13; fi27++ {
                    if (false && v15) {
                        v6 = append(v6, "a")
                        if v15 {
                            var v28 float32 = v3
                            _ = v28
                        }
                        fmt.Println(v5)
                    }
                    var v29 bool = false
                    for fi30 := 0; fi30 < 6; fi30++ {
                        v31 := GenS2{F0: 987, F1: 127, F2: 90}
                        v32 := GenS3{F0: (-91 >= 1), F1: (v5 * 0), F2: uint8(v13)}
                        if len(v6) > 0 {
                            v6[0] = "hello"
                        }
                        var v33 int = int(v5)
                        v34 := GenS4{F0: (^(v13)), F1: 239, F2: 2}
                        v2 = v2
                        fmt.Println(fi27, v17)
                        _ = fi30
                        _ = v31
                        _ = v32
                        _ = v33
                        _ = v34
                    }
                    _ = fi27
                    _ = v29
                }
                _ = v25
                _ = v26
            } else {
                for fi35 := 0; fi35 < 5; fi35++ {
                    var v36 int64 = v16
                    _ = fi35
                    _ = v36
                }
            }
        } else {
            var v37 int = 0
            var v38 uint64 = 4294967295
            var v39 uint64 = v38
            _ = v37
            _ = v38
            _ = v39
        }
        var v40 float64 = 1.5
        _ = v40
    } else {
        var v41 uint16 = v13
        fmt.Println(v16, v1, v5)
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
