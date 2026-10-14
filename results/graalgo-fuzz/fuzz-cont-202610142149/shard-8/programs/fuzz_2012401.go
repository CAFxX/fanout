package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32; F1 string }

func main() {
    var v1 float32 = (0.0 + 1.5)
    v2 := make(map[string]uint32)
    var v3 int8 = 56
    var v4 float64 = (-2.25 + 0.0)
    v5 := make(map[string]string)
    fmt.Println(v3, v3)
    v3 = v3
    v6 := []int{584}
    switch -158 {
    case 4294967295:
        var v7 int32 = 1
        _ = v7
    case 1:
        v8 := GenS1{F0: true, F1: (!(-379 != 870))}
        _ = v8
    }
    if (4294967295 < 530) {
        var v9 string = "a"
        delete(v2, v9)
        switch (((v9 + ("" + "\n\t")) + v9) + v9) {
        case "\n\t":
            switch -1 {
            case 32767:
                v10 := GenS2{F0: false}
                _ = v10
            case 665:
                v4 -= v4
            case 372:
                for fi11 := 0; fi11 < 10; fi11++ {
                    v4 -= 1e10
                    var v12 int = 32767
                    _ = fi11
                    _ = v12
                }
            }
        case "foo bar":
            fmt.Println(v4, v1)
        default:
            for fi13 := 0; fi13 < 13; fi13++ {
                var v14 bool = (false && true)
                var v15 uint8 = 53
                _ = fi13
                _ = v14
                _ = v15
            }
        }
        _ = v9
    } else {
        v3 = -128
        _ = len(v5)
    }
    for fi16 := 0; fi16 < 9; fi16++ {
        var v17 int64 = 1
        if false {
            var v18 uint = 1
            var v19 int64 = (-(v17))
            if false {
                v6 = append(v6, (^(fi16)))
                for fi20 := 0; fi20 < 11; fi20++ {
                    v17 += (v17 & v19)
                    v21 := GenS2{F0: false}
                    switch ("\n\t" + ("foo bar" + ("a" + "foo bar"))) {
                    case "foo bar":
                        if (false || true) {
                            fmt.Println(v19)
                            v6 = append(v6, fi16)
                            var v22 int16 = 0
                            _ = v22
                        }
                    case "\n\t":
                        v17 += v17
                    }
                    switch "" {
                    case "a":
                        fmt.Println(v19, v17)
                    case "foo bar":
                        v3 = (v3 + v3)
                    case "hello":
                        if true {
                            var v23 int = fi20
                            var v24 uint8 = 128
                            var v25 int8 = v3
                            switch fi20 {
                            case 9223372036854775807:
                                var v26 float64 = float64(0.0)
                                _ = v26
                            case -1:
                                v27 := GenS3{F0: "foo bar"}
                                _ = v27
                            }
                            var v28 uint32 = 65535
                            v3 = (v3 / ((-68) | 1))
                            _ = v23
                            _ = v24
                            _ = v25
                            _ = v28
                        }
                    }
                    _ = fi20
                    _ = v21
                }
            }
            _ = v18
            _ = v19
        } else {
            var v29 uint8 = 226
            _ = v29
        }
        _ = fi16
        _ = v17
    }
    switch -1 {
    case 32767:
        _ = len(v6)
    default:
        if false {
            if (969 == 1) {
                if false {
                    var v30 uint64 = 1
                    if len(v6) > 0 {
                        v6[0] = 0
                    }
                    var v31 int64 = 9223372036854775807
                    _ = v30
                    _ = v31
                }
                var v32 int = -309
                _ = v32
            }
            for fi33 := 0; fi33 < 8; fi33++ {
                v1 = 1e10
                _ = fi33
            }
            if ((42 <= 18446744073709551615) || (false && (0 > 0))) {
                for fi34 := 0; fi34 < 16; fi34++ {
                    var v35 bool = false
                    _ = fi34
                    _ = v35
                }
            }
        }
    }
    if (!(v3 > v3)) {
        delete(v2, ("foo bar" + ""))
        var v36 uint = 359
        _ = v36
    }
    v4 += v4
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v37 uint8 = 43
    fmt.Println(v3, v37, v37)
    for fi38 := 0; fi38 < 4; fi38++ {
        v39 := GenS4{F0: 1, F1: "hello"}
        v4 -= (float64((v1 / float32(v4))) * 1e10)
        _ = fi38
        _ = v39
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
