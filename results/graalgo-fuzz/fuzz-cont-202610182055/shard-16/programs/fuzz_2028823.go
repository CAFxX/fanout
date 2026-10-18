package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 uint16 }
type GenS2 struct { F0 uint; F1 int8; F2 bool }
type GenS3 struct { F0 int; F1 bool }
func genFn0(p0 bool, p1 uint16, p2 uint64) int32 {
    var v5 uint32 = 1
    var v6 float64 = 0.0
    p2 = (p2 * 1)
    if (-1 >= 2147483647) {
        p2 -= (p2 / ((42) | 1))
    } else {
        var v7 int = -9223372036854775808
        v8 := GenS1{F0: p2, F1: 261, F2: uint16(p2)}
        _ = v7
        _ = v8
    }
    if (!p0) {
        fmt.Println(v5)
        v9 := []int{-555, 32767}
        var v10 float32 = -2.25
        _ = v10
        _ = v9
    } else {
        v11 := GenS1{F0: (p2 % (((p2 % ((uint64(p1)) | 1))) | 1)), F1: -846, F2: 32767}
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return -1
}
func genFn1(p0 uint16) int {
    var v12 float64 = 1e10
    v12 += 0.0
    fmt.Println(p0)
    v13 := []int{}
    _ = p0
    _ = v12
    _ = v13
    return 32767
}

func main() {
    v1 := make(map[string]string)
    v2 := []string{}
    var v3 bool = (42 != 0)
    var v4 uint = 65535
    v14 := GenS1{F0: 0, F1: 9223372036854775807, F2: 42}
    var v15 uint64 = 463
    v1[(("" + ("foo bar" + ("" + ""))) + (("foo bar" + ("foo bar" + "")) + ("hello" + ("a" + "foo bar"))))] = (("\n\t" + "foo bar") + "foo bar")
    for fi16 := 0; fi16 < 3; fi16++ {
        var v17 uint64 = v15
        _ = fi16
        _ = v17
    }
    if v3 {
        var v18 int8 = -70
        v19 := GenS1{F0: (v15 >> (((v4 * v4)) & 63)), F1: 42, F2: 42}
        _ = v18
        _ = v19
    }
    delete(v1, "a")
    switch -1 {
    case -399:
        v15 += v15
    case -2147483648:
        switch 4294967295 {
        case 1:
            for fi20 := 0; fi20 < 10; fi20++ {
                switch fi20 {
                case 2147483647:
                    v21 := GenS1{F0: (v15 - v15), F1: 721, F2: 752}
                    _ = v21
                default:
                    v22 := GenS2{F0: (v4 >> (((v4 * v4)) & 63)), F1: 100, F2: (v4 == 65535)}
                    _ = v22
                }
                v3 = true
                _ = fi20
            }
        case 128:
            if true {
                for fi23 := 0; fi23 < 13; fi23++ {
                    fmt.Println(v15, fi23)
                    _ = fi23
                }
                for fi24 := 0; fi24 < 3; fi24++ {
                    var v25 uint32 = 4294967295
                    var v26 int8 = 100
                    for fi27 := 0; fi27 < 5; fi27++ {
                        var v28 uint16 = uint16(v4)
                        v15 = v15
                        _ = fi27
                        _ = v28
                    }
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                v4 -= (v4 ^ 474)
            } else {
                v15 += v15
                v2 = append(v2, "\n\t")
            }
        default:
            var v29 int8 = -1
            _ = v29
        }
    }
    v15 += v15
    v15 = 4294967295
    for fi30 := 0; fi30 < 18; fi30++ {
        v3 = v3
        var v31 int = (fi30 * 4294967295)
        var v32 float32 = (3.14159 + (1.5 * float32(float64(1e10))))
        _ = fi30
        _ = v31
        _ = v32
    }
    if (-1 < 1) {
        v4 = v4
        var v33 float32 = (1.5 + 1.5)
        _ = v33
    } else {
        if (("\n\t" + ("" + "foo bar")) >= "hello") {
            if (2 <= 0) {
                if len(v2) > 0 {
                    v2[0] = ((("hello" + ((("a" + ("" + "a")) + (("a" + "a") + "a")) + (("" + "foo bar") + ""))) + "hello") + ("foo bar" + "hello"))
                }
                _ = len(v2)
            }
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v34 := GenS3{F0: 0, F1: (!((42 <= -1) && ((1e10 == (1.5 / 1.0)) && false)))}
            _ = v34
        }
        v4 = 2
        var v35 float64 = ((3.14159 * 0.0) + -2.25)
        _ = v35
    }
    for fi36 := 0; fi36 < 2; fi36++ {
        _ = len(v1)
        _ = fi36
    }
    v37 := GenS3{F0: genFn1(42), F1: (211 == 65535)}
    for fi38 := 0; fi38 < 3; fi38++ {
        v39 := v1[((("\n\t" + "a") + ("a" + "a")) + ("" + ("hello" + (("hello" + ("" + ("\n\t" + (("foo bar" + "hello") + "hello")))) + "\n\t"))))]
        _ = fi38
        _ = v39
    }
    v3 = (false || ("hello" != ""))
    var v40 int = 1
    for fi41 := 0; fi41 < 8; fi41++ {
        var v42 uint = v4
        _ = fi41
        _ = v42
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v40
}
