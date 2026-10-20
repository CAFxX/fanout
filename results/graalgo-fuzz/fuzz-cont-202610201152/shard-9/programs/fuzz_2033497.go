package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 uint16 }
func genFn0(p0 bool, p1 int8, p2 int16) string {
    for fi5 := 0; fi5 < 14; fi5++ {
        for fi6 := 0; fi6 < 10; fi6++ {
            var v7 uint32 = 65535
            var v8 uint16 = 1
            v9 := []int{fi6, fi5, (fi6 + (fi6 | fi6)), fi5}
            _ = fi6
            _ = v7
            _ = v8
            _ = v9
        }
        v10 := GenS1{F0: 0, F1: "a", F2: 42}
        p2 = 1
        _ = fi5
        _ = v10
    }
    if p0 {
        if false {
            v11 := []int{184, -2147483648, 653, 4294967295}
            _ = v11
        } else {
            v12 := GenS1{F0: 42, F1: "a", F2: 135}
            fmt.Println(p0, p0)
            _ = v12
        }
        v13 := GenS1{F0: 941, F1: ("foo bar" + "foo bar"), F2: 50}
        v14 := GenS1{F0: -9223372036854775808, F1: ("" + "foo bar"), F2: 2}
        _ = v13
        _ = v14
    } else {
        for fi15 := 0; fi15 < 11; fi15++ {
            p2 = (p2 + p2)
            fmt.Println(p1, p0)
            p1 = (p1 / (((p1 >> ((1) & 63))) | 1))
            _ = fi15
        }
    }
    for fi16 := 0; fi16 < 15; fi16++ {
        fmt.Println(p2, p2)
        _ = fi16
    }
    _ = p0
    _ = p1
    _ = p2
    return "foo bar"
}
func genFn1(p0 string, p1 int) bool {
    v17 := []int{}
    if (true || (!true)) {
        fmt.Println(p1)
    }
    if len(v17) >= 2 {
        v17 = v17[1:len(v17)]
    }
    _ = p0
    _ = p1
    _ = v17
    return (0.0 >= 0.0)
}

func main() {
    v1 := make(map[string]uint32)
    var v2 bool = (false && true)
    v3 := make(map[string]bool)
    v4 := make(map[string]string)
    v2 = (false || false)
    v2 = v2
    if (856 < 483) {
        v18 := GenS1{F0: 9223372036854775807, F1: ("foo bar" + ("\n\t" + "a")), F2: 42}
        if (v2 || (!v2)) {
            v2 = (42 != 128)
        }
        if (!(v2 && ((63 < -615) && (v2 || (false || (!v2)))))) {
            var v19 uint = 398
            _ = v19
        } else {
            for fi20 := 0; fi20 < 16; fi20++ {
                var v21 bool = v2
                v22 := GenS1{F0: 907, F1: "", F2: 6}
                var v23 uint32 = 2
                _ = fi20
                _ = v21
                _ = v22
                _ = v23
            }
            v2 = (2 >= 1)
        }
        _ = v18
    } else {
        var v24 uint64 = 0
        fmt.Println(v2)
        _ = v24
    }
    var v25 uint8 = 128
    var v26 int64 = 42
    v27 := v3["a"]
    v28 := []int{1, 4294967295, 32768}
    if ((((1.5 - (1.5 + 0.1)) - (0.0 - 1e10)) + 3.14159) > 0.0) {
        if (float64(1.5) != (3.14159 * 1.5)) {
            for fi29 := 0; fi29 < 16; fi29++ {
                if v27 {
                    var v30 bool = (-2.25 >= 3.14159)
                    for fi31 := 0; fi31 < 13; fi31++ {
                        var v32 uint = 18446744073709551615
                        var v33 bool = (100 <= 30)
                        v34 := GenS1{F0: (v26 ^ (v26 / (((v26 >> ((v32) & 63))) | 1))), F1: "", F2: 32767}
                        _ = fi31
                        _ = v32
                        _ = v33
                        _ = v34
                    }
                    _ = v30
                } else {
                    v25 = v25
                    if len(v28) >= 2 {
                        v28 = v28[1:len(v28)]
                    }
                }
                var v35 uint64 = 18446744073709551615
                var v36 string = (((("\n\t" + "a") + ("" + "")) + ("" + ("hello" + "hello"))) + ("foo bar" + "a"))
                _ = fi29
                _ = v35
                _ = v36
            }
            if ("a" != "foo bar") {
                var v37 uint32 = 879
                _ = v37
            }
            for fi38 := 0; fi38 < 7; fi38++ {
                var v39 float32 = 0.0
                _ = fi38
                _ = v39
            }
        }
        v28 = append(v28, 4294967295)
    } else {
        if true {
            fmt.Println(v25)
            v2 = (!v27)
        } else {
            var v40 uint = 65535
            for fi41 := 0; fi41 < 1; fi41++ {
                v40 += v40
                var v42 int32 = 42
                _ = fi41
                _ = v42
            }
            var v43 int = -653
            _ = v40
            _ = v43
        }
        v44 := GenS1{F0: v26, F1: "hello", F2: 316}
        v4["foo bar"] = "hello"
        _ = v44
    }
    fmt.Println(v2, v27, v27)
    if v27 {
        var v45 float32 = float32(0.0)
        _ = v45
    }
    var v46 string = "a"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v46
}
