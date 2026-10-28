package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 int32; F1 int8; F2 int64 }
type GenS3 struct { F0 uint64 }

func main() {
    v1 := make(map[string]bool)
    var v2 bool = (true && true)
    v3 := make(map[string]string)
    v4 := make(map[int]uint32)
    var v5 uint = 4294967295
    v6 := make(map[int]int)
    if false {
        v7 := v3[("\n\t" + "hello")]
        var v8 float32 = float32(float64(1.5))
        var v9 int = -1
        _ = v7
        _ = v8
        _ = v9
    }
    delete(v6, 1)
    v5 = 11
    for fi10 := 0; fi10 < 5; fi10++ {
        var v11 uint8 = 83
        var v12 int8 = 100
        for fi13 := 0; fi13 < 10; fi13++ {
            fmt.Println(fi10, v11, fi10)
            _ = fi13
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    switch "foo bar" {
    case "a":
        v5 += v5
    default:
        if (!(!v2)) {
            fmt.Println(v2, v2)
            if (((("foo bar" + "") + ("\n\t" + (("hello" + "hello") + ("a" + ("a" + "a"))))) + "foo bar") <= ("" + "foo bar")) {
                v3["a"] = ("a" + ("a" + "\n\t"))
                if (144 > 127) {
                    v14 := GenS1{F0: 128, F1: (v2 && (81 == 42))}
                    v2 = (67 < 65535)
                    _ = v14
                }
            } else {
                var v15 int16 = 0
                switch "foo bar" {
                case "":
                    v15 += v15
                default:
                    v15 += v15
                }
                var v16 uint16 = 1
                _ = v15
                _ = v16
            }
            var v17 int64 = 347
            _ = v17
        }
    }
    v18 := GenS1{F0: 42, F1: (127 == 57)}
    switch "foo bar" {
    case "":
        var v19 string = "hello"
        _ = v19
    }
    switch 977 {
    case -648:
        var v20 int32 = -1
        _ = v20
    case 127:
        v21 := []int{-9223372036854775808}
        _ = v21
    case 128:
        v2 = true
    default:
        v2 = false
    }
    var v22 float32 = (1.5 / float32(0.1))
    v2 = v2
    var v23 int = 1
    v24 := GenS1{F0: 246, F1: (-9223372036854775808 == -1)}
    v25 := []int{656, (v23 + 127)}
    switch (v23 | (^(v23))) {
    case 127:
        var v26 int64 = -304
        _ = v26
    case -370:
        for fi27 := 0; fi27 < 11; fi27++ {
            for fi28 := 0; fi28 < 7; fi28++ {
                if v2 {
                    switch (-(fi28)) {
                    case -2147483648:
                        if (0 == 1) {
                            v2 = (uint8(v5) < 213)
                        }
                    case 42:
                        var v29 bool = true
                        _ = v29
                    }
                    if ((2147483647 == -2147483648) || v2) {
                        _ = len(v25)
                        if (((v2 || (v2 || (uint64(v5) <= 42))) || (v2 && (!true))) || (16 < 1)) {
                            fmt.Println(v5, fi27, v22)
                            var v30 int64 = 9223372036854775807
                            if ((v2 && v2) && (42 < 278)) {
                                v30 -= 339
                                v5 = v5
                                v30 = v30
                            }
                            _ = v30
                        } else {
                            var v31 float64 = ((-2.25 - 3.14159) / (3.14159 * 0.1))
                            v32 := GenS1{F0: 128, F1: (uint64(v5) == 384)}
                            _ = v31
                            _ = v32
                        }
                        delete(v6, 4294967295)
                    } else {
                        fmt.Println(v23)
                    }
                    switch (fi27 & v23) {
                    case 9223372036854775807:
                        v5 -= (v5 % ((v5) | 1))
                    case 128:
                        v23 -= fi27
                    default:
                        fmt.Println(v2)
                        var v33 string = ""
                        var v34 uint32 = 65535
                        _ = v33
                        _ = v34
                    }
                } else {
                    var v35 float32 = 3.14159
                    v22 += (v22 / (3.14159 - v22))
                    var v36 bool = true
                    _ = v35
                    _ = v36
                }
                v37 := GenS1{F0: 2, F1: (true && v2)}
                if (v5 <= v5) {
                    v2 = true
                    var v38 float32 = v22
                    var v39 bool = v2
                    _ = v38
                    _ = v39
                }
                _ = fi28
                _ = v37
            }
            v40 := GenS1{F0: 127, F1: (127 > -31)}
            fmt.Println(v2, v22, v2)
            _ = fi27
            _ = v40
        }
    case 4294967295:
        if (-440 < 1) {
            var v41 string = "hello"
            _ = v41
        }
    }
    if v2 {
        fmt.Println(v5)
        v42 := GenS1{F0: 255, F1: ((18446744073709551615 == 4294967295) || true)}
        var v43 int8 = 97
        _ = v42
        _ = v43
    } else {
        if v2 {
            if (!(v2 || v2)) {
                v23 -= (-(v23))
                var v44 int64 = -1
                var v45 uint16 = 620
                _ = v44
                _ = v45
            } else {
                var v46 uint = v5
                if ((204 == 135) && v2) {
                    v22 -= float32(float64((float32((0.1 * (0.0 - (1e10 * float64(0.0))))) / v22)))
                    v47 := GenS2{F0: 0, F1: -1, F2: 9223372036854775807}
                    _ = v47
                }
                _ = v46
            }
            var v48 int16 = 32767
            var v49 string = "a"
            _ = v48
            _ = v49
        } else {
            for fi50 := 0; fi50 < 13; fi50++ {
                var v51 int64 = 94
                fmt.Println(v2)
                _ = fi50
                _ = v51
            }
        }
        if (false || true) {
            for fi52 := 0; fi52 < 3; fi52++ {
                switch "\n\t" {
                case "hello":
                    var v53 string = "a"
                    _ = v53
                case "\n\t":
                    for fi54 := 0; fi54 < 4; fi54++ {
                        var v55 int64 = 42
                        _ = fi54
                        _ = v55
                    }
                }
                v23 -= int(v5)
                _ = fi52
            }
        } else {
            var v56 int8 = 69
            _ = v56
        }
        var v57 uint16 = 798
        _ = v57
    }
    v2 = v2
    v58 := GenS3{F0: 0}
    if v2 {
        if len(v25) >= 2 {
            v25 = v25[1:len(v25)]
        }
        if v2 {
            var v59 uint64 = 0
            var v60 int = (v23 + v23)
            _ = v59
            _ = v60
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v6
}
