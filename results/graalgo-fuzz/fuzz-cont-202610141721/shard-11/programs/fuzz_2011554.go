package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 bool; F1 int16; F2 string }
type GenS3 struct { F0 int64; F1 uint32; F2 int32 }
type GenS4 struct { F0 int32; F1 string }

func main() {
    var v1 int64 = 1
    var v2 float32 = 3.14159
    v3 := make(map[string]uint32)
    var v4 int = 2147483647
    v5 := make(map[int]string)
    v6 := make(map[int]int)
    for fi7 := 0; fi7 < 20; fi7++ {
        for fi8 := 0; fi8 < 15; fi8++ {
            fmt.Println(fi8, v4, fi7)
            v3["a"] = 0
            v1 -= -1
            _ = fi8
        }
        for fi9 := 0; fi9 < 11; fi9++ {
            v4 -= fi9
            delete(v6, fi7)
            _ = fi9
        }
        var v10 bool = (true || false)
        _ = fi7
        _ = v10
    }
    v11 := []int{(v4 >> ((255) & 63)), (v4 * 4294967295), 707}
    v12 := v3["\n\t"]
    if true {
        var v13 int16 = 1
        var v14 uint32 = (v12 ^ (^(v12)))
        _ = v13
        _ = v14
    }
    var v15 int16 = 42
    var v16 uint32 = v12
    for fi17 := 0; fi17 < 4; fi17++ {
        var v18 uint8 = 214
        if (!(false || true)) {
            switch ("hello" + "a") {
            case "a":
                if (true || true) {
                    var v19 uint8 = (v18 * v18)
                    v3[("foo bar" + "")] = v16
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                    _ = v19
                } else {
                    switch (v4 * (v4 % ((1) | 1))) {
                    case 128:
                        fmt.Println(v18)
                    case -342:
                        for fi20 := 0; fi20 < 3; fi20++ {
                            v1 = int64(v4)
                            var v21 int32 = 0
                            _ = fi20
                            _ = v21
                        }
                    case 2147483647:
                        if (v18 > v18) {
                            v2 = (v2 * ((((v2 - (0.1 - (v2 - 0.1))) * v2) + (3.14159 + v2)) + 1e10))
                            if (1 <= 2) {
                                var v22 float64 = 3.14159
                                _ = v22
                            } else {
                                switch (-(fi17)) {
                                case -330:
                                    v15 -= -32768
                                case 32768:
                                    v11 = append(v11, 42)
                                case -1:
                                    var v23 int64 = (v1 + v1)
                                    _ = v23
                                }
                                v2 += 1e10
                            }
                        } else {
                            v24 := v6[(v4 + (v4 - v4))]
                            v25 := v6[fi17]
                            v26 := GenS1{F0: v1, F1: 725}
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                    }
                }
            case "foo bar":
                v4 -= fi17
            case "\n\t":
                var v27 float32 = 0.0
                _ = v27
            }
            if true {
                v1 -= -1
                v15 = v15
            }
        } else {
            switch fi17 {
            case 127:
                v28 := GenS1{F0: int64(v4), F1: 4294967295}
                _ = v28
            default:
                for fi29 := 0; fi29 < 1; fi29++ {
                    v4 -= (fi29 >> ((42) & 63))
                    _ = len(v11)
                    var v30 bool = false
                    _ = fi29
                    _ = v30
                }
            }
            v11 = append(v11, v4)
        }
        _ = fi17
        _ = v18
    }
    switch -9223372036854775808 {
    case 1:
        var v31 int = v4
        _ = v31
    case 42:
        v1 = (v1 >> ((4294967295) & 63))
    default:
        if (true && ((("hello" <= "hello") || true) && true)) {
            v12 = 42
        } else {
            fmt.Println(v2)
            var v32 uint32 = v12
            if false {
                v16 = (v16 + (v32 - (v16 + (v12 & (v12 % ((v12) | 1))))))
                v33 := GenS2{F0: true, F1: v15, F2: ("\n\t" + ((("" + "") + "") + ""))}
                _ = v33
            }
            _ = v32
        }
    }
    for fi34 := 0; fi34 < 6; fi34++ {
        switch v4 {
        case 0:
            fmt.Println(v15, v4)
        case -831:
            var v35 uint32 = v16
            _ = v35
        }
        _ = fi34
    }
    var v36 int16 = (v15 + 1000)
    if (false || (!(42 != 65535))) {
        fmt.Println(v16, v36)
        v37 := GenS1{F0: 1, F1: 494}
        _ = v37
    } else {
        v38 := GenS3{F0: 189, F1: v12, F2: 2147483647}
        var v39 int16 = v36
        v40 := GenS4{F0: 0, F1: "\n\t"}
        _ = v38
        _ = v39
        _ = v40
    }
    if ((255 >= 401) && false) {
        var v41 uint8 = 23
        _ = v41
    } else {
        var v42 uint8 = 68
        var v43 int = -106
        _ = v42
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
