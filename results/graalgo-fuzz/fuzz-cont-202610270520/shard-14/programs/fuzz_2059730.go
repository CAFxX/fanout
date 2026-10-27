package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64; F1 int32; F2 bool }
type GenS4 struct { F0 int8; F1 uint; F2 int64 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 string; F1 uint }
type GenS7 struct { F0 uint64; F1 int32 }

func main() {
    v1 := make(map[int]string)
    var v2 uint64 = 805
    var v3 uint32 = 0
    var v4 string = "foo bar"
    if (2 == 0) {
        v5 := GenS1{F0: 127, F1: false, F2: 65535}
        if false {
            v6 := GenS1{F0: 1, F1: true, F2: 8}
            if (true && false) {
                v7 := GenS1{F0: -65, F1: false, F2: 1}
                fmt.Println(v2, v2)
                v8 := []int{-844, 204, 32768}
                _ = v7
                _ = v8
            } else {
                if (!true) {
                    v9 := GenS2{F0: (906 >= 637)}
                    var v10 string = ((("a" + ((((v4 + "") + "\n\t") + ((v4 + "") + ("\n\t" + "hello"))) + v4)) + ("\n\t" + v4)) + "\n\t")
                    switch -440 {
                    case -961:
                        var v11 uint32 = 269
                        _ = v11
                    case -2147483648:
                        var v12 int32 = 0
                        _ = v12
                    }
                    _ = v10
                    _ = v9
                }
                v2 = 42
                v13 := []int{-2147483648, 32767, 9223372036854775807, -2147483648}
                _ = v13
            }
            _ = v6
        }
        _ = v5
    }
    var v14 uint32 = 2147483648
    v14 -= (v3 + v3)
    if true {
        var v15 bool = true
        if false {
            v16 := []int{354}
            v15 = false
            v17 := GenS2{F0: v15}
            _ = v16
            _ = v17
        } else {
            if ((162 > 128) || (346 > 185)) {
                v4 = v4
            } else {
                if true {
                    var v18 int = -2147483648
                    v2 = 576
                    _ = v18
                } else {
                    v4 = ((("a" + (v4 + v4)) + "foo bar") + (v4 + "a"))
                    fmt.Println(v14, v2)
                    v19 := GenS3{F0: (v2 >> ((187) & 63)), F1: 42, F2: false}
                    _ = v19
                }
                var v20 int32 = 815
                v3 -= v3
                _ = v20
            }
            v21 := []int{-809, 219, 32768, -221}
            v22 := GenS2{F0: (251 >= 128)}
            _ = v21
            _ = v22
        }
        for fi23 := 0; fi23 < 19; fi23++ {
            switch 1 {
            case -594:
                v3 = (^(v3))
            case 120:
                switch (v4 + (v4 + v4)) {
                case "a":
                    for fi24 := 0; fi24 < 3; fi24++ {
                        v25 := GenS3{F0: (v2 ^ 18446744073709551615), F1: int32(fi24), F2: (float64(3.14159) > 0.1)}
                        if true {
                            var v26 uint32 = 460
                            _ = v26
                        } else {
                            v27 := GenS4{F0: -128, F1: 42, F2: 9223372036854775807}
                            _ = v27
                        }
                        v3 = (v3 ^ (v3 - (v14 ^ (v14 * v14))))
                        v2 = 18446744073709551615
                        _ = fi24
                        _ = v25
                    }
                case "foo bar":
                    v28 := GenS4{F0: -128, F1: 0, F2: 1}
                    _ = v28
                }
            }
            var v29 float64 = 3.14159
            _ = fi23
            _ = v29
        }
        _ = v15
    }
    v30 := GenS5{F0: v2}
    v31 := []int{4294967295, -1}
    var v32 float32 = (3.14159 - -2.25)
    var v33 string = ("foo bar" + ((v4 + v4) + v4))
    for fi34 := 0; fi34 < 5; fi34++ {
        for fi35 := 0; fi35 < 11; fi35++ {
            v33 = "\n\t"
            var v36 uint64 = 4294967295
            v37 := GenS6{F0: ("hello" + v33), F1: uint(v3)}
            _ = fi35
            _ = v36
            _ = v37
        }
        _ = fi34
    }
    var v38 int16 = 1000
    for fi39 := 0; fi39 < 6; fi39++ {
        v33 = v33
        var v40 int32 = 0
        _ = fi39
        _ = v40
    }
    if (224 < 1) {
        if (-1 <= 1) {
            var v41 uint64 = v2
            _ = v41
        } else {
            if len(v31) > 0 {
                v31[0] = 450
            }
        }
        v42 := GenS7{F0: v2, F1: 2147483647}
        _ = v42
    } else {
        fmt.Println(v32)
        switch "" {
        case "":
            for fi43 := 0; fi43 < 18; fi43++ {
                if (!(-2147483648 <= 2147483647)) {
                    for fi44 := 0; fi44 < 11; fi44++ {
                        v2 += v2
                        fmt.Println(v38)
                        fmt.Println(v4, v33, fi43)
                        _ = fi44
                    }
                    v38 = v38
                    var v45 uint16 = 65535
                    _ = v45
                }
                fmt.Println(fi43, v38)
                for fi46 := 0; fi46 < 20; fi46++ {
                    fmt.Println(fi43, v2)
                    fmt.Println(v14, v4)
                    _ = fi46
                }
                _ = fi43
            }
        case "foo bar":
            for fi47 := 0; fi47 < 1; fi47++ {
                var v48 int = 128
                _ = fi47
                _ = v48
            }
        case "\n\t":
            v2 = (v2 & 437)
        default:
            fmt.Println(v14, v38)
        }
        var v49 int = 127
        _ = v49
    }
    if (((v2 - 0) != (^(v2))) && false) {
        delete(v1, 42)
    }
    v38 += v38
    _ = len(v1)
    var v50 int32 = -2147483648
    v38 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v38
    _ = v4
    _ = v50
}
