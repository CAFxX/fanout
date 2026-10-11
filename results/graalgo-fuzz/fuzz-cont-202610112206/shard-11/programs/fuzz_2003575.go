package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 string }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 string; F1 int8; F2 uint32 }
func genFn0(p0 uint16) bool {
    for fi6 := 0; fi6 < 14; fi6++ {
        p0 -= 337
        v7 := GenS1{F0: 127, F1: "hello", F2: ""}
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 11; fi8++ {
        v9 := []int{2147483647}
        _ = fi8
        _ = v9
    }
    _ = p0
    return true
}
func genFn1(p0 uint8) string {
    fmt.Println(p0)
    var v10 uint64 = 846
    switch "\n\t" {
    case "a":
        if ("hello" == "a") {
            switch 128 {
            case 4294967295:
                var v11 float64 = (0.0 - 0.0)
                _ = v11
            case 9223372036854775807:
                v12 := []int{969, 32767}
                _ = v12
            }
            for fi13 := 0; fi13 < 15; fi13++ {
                v14 := []int{}
                if (!false) {
                    switch "\n\t" {
                    case "\n\t":
                        if (float64((0.1 / -2.25)) < (1e10 / (3.14159 + (((0.1 / 1e10) - float64((0.1 * -2.25))) + 0.0)))) {
                            fmt.Println(p0, p0)
                            var v15 uint = 18446744073709551615
                            v10 = (^(v10))
                            _ = v15
                        }
                    case "hello":
                        for fi16 := 0; fi16 < 10; fi16++ {
                            if len(v14) > 0 {
                                v14[0] = (fi13 / ((989) | 1))
                            }
                            var v17 string = ("foo bar" + "a")
                            if false {
                                var v18 uint32 = 4294967295
                                _ = v18
                            } else {
                                v14 = append(v14, (fi16 / ((fi13) | 1)))
                                p0 = (^(p0))
                            }
                            fmt.Println(fi13, v17, fi13)
                            var v19 string = "a"
                            _ = fi16
                            _ = v17
                            _ = v19
                        }
                    case "":
                        v20 := GenS1{F0: (p0 << ((2) & 63)), F1: "", F2: "foo bar"}
                        _ = v20
                    default:
                        if len(v14) > 0 {
                            v14[0] = -9223372036854775808
                        }
                    }
                    p0 = 0
                }
                v10 = (v10 / ((42) | 1))
                _ = fi13
                _ = v14
            }
            var v21 int32 = 228
            _ = v21
        } else {
            v10 = v10
        }
    case "":
        var v22 float32 = (0.0 / 3.14159)
        _ = v22
    }
    _ = p0
    _ = v10
    return "hello"
}

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[string]string)
    var v3 uint = 18446744073709551615
    v4 := []int{}
    var v5 int16 = -611
    v2[("\n\t" + ("a" + ("hello" + "hello")))] = ""
    var v23 uint8 = 0
    if true {
        if true {
            v23 = v23
            for fi24 := 0; fi24 < 20; fi24++ {
                v5 += (v5 % ((v5) | 1))
                _ = fi24
            }
        }
    }
    v25 := v2[("foo bar" + ("foo bar" + "a"))]
    var v26 int64 = 773
    v25 = v25
    for fi27 := 0; fi27 < 8; fi27++ {
        v28 := GenS1{F0: (v23 % (((v23 % ((uint8(v3)) | 1))) | 1)), F1: (("a" + "foo bar") + (v25 + "")), F2: v25}
        v29 := GenS1{F0: v23, F1: v25, F2: "a"}
        if len(v4) > 0 {
            v4[0] = (fi27 - fi27)
        }
        _ = fi27
        _ = v28
        _ = v29
    }
    if false {
        fmt.Println(v25)
    } else {
        fmt.Println(v1)
    }
    v1 = (v1 / ((2147483647) | 1))
    fmt.Println(v1, v23)
    v26 = (v26 << (((v3 | v3)) & 63))
    v30 := GenS2{F0: 255, F1: -1}
    v1 = v1
    if false {
        switch (v25 + (v25 + "foo bar")) {
        case "":
            var v31 uint16 = 64
            _ = v31
        }
        var v32 uint = v3
        _ = v32
    }
    v1 = (v1 % (((v1 * v1)) | 1))
    for fi33 := 0; fi33 < 13; fi33++ {
        v34 := GenS1{F0: (v23 / ((128) | 1)), F1: v25, F2: v25}
        switch ("\n\t" + v25) {
        case "foo bar":
            v5 -= (v5 % (((v5 | 0)) | 1))
        case "\n\t":
            if false {
                v23 -= 255
            } else {
                for fi35 := 0; fi35 < 16; fi35++ {
                    var v36 int = (fi35 - -141)
                    _ = fi35
                    _ = v36
                }
                v37 := v2[v25]
                fmt.Println(v25, v1, v37)
                _ = v37
            }
        default:
            for fi38 := 0; fi38 < 3; fi38++ {
                v39 := GenS3{F0: v25, F1: 127, F2: 4294967295}
                v3 -= 18446744073709551615
                _ = fi38
                _ = v39
            }
        }
        _ = fi33
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
