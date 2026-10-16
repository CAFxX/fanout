package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 uint32; F2 uint8 }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 int32; F1 int; F2 int16 }

func main() {
    v1 := make(map[string]bool)
    var v2 int32 = -776
    v3 := make(map[uint32]string)
    v4 := make(map[int]string)
    switch "foo bar" {
    case "\n\t":
        v5 := v3[2147483648]
        _ = v5
    case "foo bar":
        var v6 uint16 = 638
        _ = v6
    default:
        v2 += (v2 % ((v2) | 1))
    }
    for fi7 := 0; fi7 < 17; fi7++ {
        v2 = v2
        v3[914] = "hello"
        v8 := GenS1{F0: 1000}
        _ = fi7
        _ = v8
    }
    v2 -= -810
    v2 -= v2
    v2 -= (v2 + (v2 << ((42) & 63)))
    delete(v4, -630)
    fmt.Println(v2)
    var v9 float32 = 3.14159
    var v10 uint8 = 42
    var v11 int = 128
    for fi12 := 0; fi12 < 19; fi12++ {
        for fi13 := 0; fi13 < 18; fi13++ {
            var v14 uint = 0
            v11 += fi12
            _ = fi13
            _ = v14
        }
        for fi15 := 0; fi15 < 8; fi15++ {
            delete(v3, 216)
            v10 = (v10 >> ((263) & 63))
            var v16 uint16 = 32767
            _ = fi15
            _ = v16
        }
        if true {
            if (((-32768 == 42) && (false && (-1 < v2))) || false) {
                fmt.Println(fi12, fi12, fi12)
                if (!false) {
                    var v17 int = 32767
                    v18 := v1["a"]
                    v1["foo bar"] = true
                    _ = v17
                    _ = v18
                }
                _ = len(v1)
            }
            if ((v2 != (v2 >> ((2) & 63))) || ((v10 & v10) == 128)) {
                switch (fi12 + (fi12 / ((128) | 1))) {
                case 0:
                    v10 = v10
                case -41:
                    v2 = (v2 << ((1) & 63))
                case -400:
                    delete(v4, v11)
                }
                switch "a" {
                case "a":
                    var v19 int = (v11 % ((fi12) | 1))
                    _ = v19
                case "foo bar":
                    _ = len(v4)
                case "\n\t":
                    var v20 int16 = -599
                    _ = v20
                }
            }
        }
        _ = fi12
    }
    fmt.Println(v9)
    var v21 string = (("hello" + "\n\t") + "foo bar")
    switch v11 {
    case 127:
        fmt.Println(v11, v21)
    case 747:
        v22 := GenS1{F0: -1}
        _ = v22
    case -9223372036854775808:
        if false {
            for fi23 := 0; fi23 < 5; fi23++ {
                v24 := GenS1{F0: -32768}
                _ = fi23
                _ = v24
            }
            var v25 int16 = -256
            _ = v25
        } else {
            var v26 uint64 = 42
            v21 = "hello"
            _ = v26
        }
    default:
        for fi27 := 0; fi27 < 9; fi27++ {
            var v28 uint32 = 0
            v29 := GenS2{F0: 0, F1: v28, F2: 42}
            var v30 uint = 364
            _ = fi27
            _ = v28
            _ = v29
            _ = v30
        }
    }
    _ = len(v1)
    if (!true) {
        switch (v11 + (v11 | (^(v11)))) {
        case 42:
            fmt.Println(v2, v11, v11)
        default:
            fmt.Println(v11, v10, v9)
        }
        for fi31 := 0; fi31 < 10; fi31++ {
            if (!false) {
                v32 := []int{4294967295, 128, -2147483648}
                var v33 int16 = -784
                _ = v32
                _ = v33
            } else {
                var v34 int8 = 115
                for fi35 := 0; fi35 < 8; fi35++ {
                    v11 = (fi35 ^ (fi31 & fi31))
                    switch (fi31 >> ((330) & 63)) {
                    case 9223372036854775807:
                        if true {
                            v21 = "hello"
                            var v36 bool = (!true)
                            _ = v36
                        } else {
                            delete(v4, fi35)
                            v37 := []int{}
                            var v38 float64 = (-2.25 * 0.0)
                            _ = v37
                            _ = v38
                        }
                    case -1:
                        fmt.Println(v34, v10, v11)
                    case 128:
                        v9 = 3.14159
                    default:
                        switch (v11 % ((2147483647) | 1)) {
                        case -412:
                            v11 = fi35
                        default:
                            switch int(v2) {
                            case 127:
                                v39 := []int{}
                                _ = v39
                            case -229:
                                fmt.Println(v11, v34)
                            default:
                                switch "" {
                                case "a":
                                    var v40 float64 = 0.1
                                    _ = v40
                                case "":
                                    v41 := GenS1{F0: -102}
                                    _ = v41
                                case "\n\t":
                                    v21 = v21
                                default:
                                    var v42 bool = (true || true)
                                    _ = v42
                                }
                            }
                        }
                    }
                    v43 := GenS3{F0: false, F1: int32(v10)}
                    _ = fi35
                    _ = v43
                }
                var v44 int16 = 252
                _ = v34
                _ = v44
            }
            if false {
                var v45 float32 = ((((v9 * v9) + 0.0) - 3.14159) * 0.0)
                var v46 uint16 = 547
                _ = v45
                _ = v46
            }
            v47 := GenS4{F0: (v2 * (v2 + (v2 / ((v2) | 1)))), F1: (v11 >> ((42) & 63)), F2: 1}
            _ = fi31
            _ = v47
        }
        var v48 float32 = float32(0.1)
        _ = v48
    } else {
        var v49 int64 = -1
        _ = v49
    }
    var v50 int64 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v50:", v50)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v50
    _ = v9
}
