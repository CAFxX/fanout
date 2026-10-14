package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 int64 }
type GenS2 struct { F0 uint32; F1 int8; F2 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint32; F1 int; F2 int }
type GenS6 struct { F0 string; F1 int8; F2 bool }
type GenS7 struct { F0 int8 }
func genFn0(p0 uint64, p1 int, p2 uint16) int16 {
    p2 = p2
    v6 := []int{4294967295, (p1 >> ((255) & 63)), (p1 | (p1 << ((4294967295) & 63))), p1}
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 string = ""
        p0 = p0
        _ = fi7
        _ = v8
    }
    var v9 uint8 = 60
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v9
    return -32768
}
func genFn1(p0 uint8, p1 uint64) string {
    var v10 string = "foo bar"
    v11 := GenS1{F0: 4294967295, F1: (-2147483648 < -310), F2: -1}
    var v12 uint32 = 482
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    return ("hello" + v10)
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]bool)
    var v3 uint32 = 2147483647
    v4 := []string{}
    var v5 float64 = 3.14159
    v13 := GenS2{F0: (v3 * (^(v3))), F1: 127, F2: 0}
    var v14 int16 = 0
    for fi15 := 0; fi15 < 17; fi15++ {
        for fi16 := 0; fi16 < 1; fi16++ {
            fmt.Println(v3, v14, v5)
            _ = fi16
        }
        _ = len(v4)
        for fi17 := 0; fi17 < 4; fi17++ {
            var v18 int64 = -9223372036854775808
            if len(v4) > 0 {
                v4[0] = "a"
            }
            var v19 int8 = -1
            if true {
                fmt.Println(v19, fi15)
                if (2147483647 <= -1) {
                    switch "foo bar" {
                    case "a":
                        v20 := GenS3{F0: v14}
                        _ = v20
                    case "":
                        var v21 uint16 = 243
                        _ = v21
                    }
                    var v22 float64 = -2.25
                    v23 := GenS4{F0: 475}
                    _ = v22
                    _ = v23
                }
                v14 -= v14
                v24 := GenS5{F0: (v3 % ((v3) | 1)), F1: (fi15 + fi17), F2: 32767}
                var v25 uint16 = 985
                _ = v24
                _ = v25
            }
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = fi15
    }
    fmt.Println(v14, v14)
    v26 := GenS6{F0: "\n\t", F1: 102, F2: (true && (!(1 == 65535)))}
    var v27 int8 = -1
    var v28 uint16 = 767
    if (!(1 >= 665)) {
        for fi29 := 0; fi29 < 12; fi29++ {
            var v30 int8 = v27
            var v31 int = 42
            v3 = 2147483647
            _ = fi29
            _ = v30
            _ = v31
        }
        delete(v1, (((("foo bar" + "foo bar") + "hello") + "hello") + "hello"))
    } else {
        v32 := GenS7{F0: 42}
        v33 := GenS7{F0: (v27 ^ 88)}
        _ = v32
        _ = v33
    }
    v3 = v3
    for fi34 := 0; fi34 < 16; fi34++ {
        fmt.Println(fi34, v5, v5)
        var v35 int = 127
        var v36 int64 = 0
        _ = fi34
        _ = v35
        _ = v36
    }
    for fi37 := 0; fi37 < 18; fi37++ {
        for fi38 := 0; fi38 < 16; fi38++ {
            fmt.Println(v5, v3, v27)
            var v39 uint16 = v28
            if (3.14159 != v5) {
                if (2 >= 4294967295) {
                    v40 := GenS1{F0: fi38, F1: true, F2: int64(fi38)}
                    v41 := GenS2{F0: (v3 % (((^(v3))) | 1)), F1: (-(v27)), F2: v14}
                    v42 := GenS6{F0: "hello", F1: v27, F2: (v39 > 389)}
                    _ = v40
                    _ = v41
                    _ = v42
                } else {
                    if (255 != 1) {
                        v1["a"] = v3
                        v5 += -2.25
                        var v43 int8 = -128
                        _ = v43
                    }
                    if false {
                        var v44 float32 = 0.1
                        var v45 uint32 = v3
                        v3 -= 2
                        _ = v44
                        _ = v45
                    } else {
                        fmt.Println(fi38, v39, v28)
                    }
                    var v46 int64 = 0
                    _ = v46
                }
                if true {
                    var v47 int16 = genFn0(721, 4294967295, (v39 % (((v28 % ((v28) | 1))) | 1)))
                    v4 = append(v4, "")
                    _ = v47
                }
                switch (fi38 << ((65535) & 63)) {
                case 4294967295:
                    v3 += 476
                    fmt.Println(v28, fi38, v14)
                    v28 += (v39 % (((v39 % ((v39) | 1))) | 1))
                case 183:
                    var v48 int16 = 160
                    _ = v48
                }
            }
            _ = fi38
            _ = v39
        }
        fmt.Println(v5, v14)
        _ = fi37
    }
    switch -1 {
    case -1:
        v5 += 0.0
    case 128:
        v49 := v1["foo bar"]
        _ = v49
    default:
        v3 = 332
    }
    for fi50 := 0; fi50 < 2; fi50++ {
        v14 = -1
        fmt.Println(v28, v27)
        _ = fi50
    }
    v51 := GenS3{F0: v14}
    for fi52 := 0; fi52 < 14; fi52++ {
        _ = len(v4)
        _ = fi52
    }
    var v53 int64 = -9223372036854775808
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v53
}
