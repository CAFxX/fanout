package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 int64 }
type GenS2 struct { F0 uint16; F1 uint16; F2 uint32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int8; F1 uint32; F2 uint64 }
type GenS5 struct { F0 uint16; F1 int16 }
type GenS6 struct { F0 uint; F1 int; F2 uint32 }
type GenS7 struct { F0 int32 }
func genFn0(p0 uint, p1 uint8, p2 uint64) uint16 {
    p2 = 1
    if false {
        if true {
            var v5 int64 = -1
            v6 := GenS1{F0: 42, F1: v5, F2: -1}
            v7 := GenS2{F0: 255, F1: 520, F2: 0}
            _ = v5
            _ = v6
            _ = v7
        }
        fmt.Println(p1, p0, p0)
    }
    v8 := GenS1{F0: -1, F1: -9223372036854775808, F2: 42}
    if true {
        v9 := GenS3{F0: -246}
        fmt.Println(p1, p0, p2)
        _ = v9
    } else {
        fmt.Println(p0, p2, p1)
    }
    p1 += p1
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 278
}
func genFn1(p0 uint8) int {
    v10 := GenS2{F0: 32767, F1: 65535, F2: 736}
    var v11 string = "\n\t"
    if (!false) {
        p0 = (p0 + 0)
        var v12 uint64 = 0
        _ = v12
    }
    switch 42 {
    case 32767:
        v13 := GenS3{F0: 1}
        _ = v13
    case 17:
        if (!true) {
            v11 = v11
            if (-126 != 0) {
                var v14 bool = true
                switch 1 {
                case 475:
                    switch ((("a" + v11) + v11) + (v11 + v11)) {
                    case "foo bar":
                        for fi15 := 0; fi15 < 6; fi15++ {
                            var v16 int64 = 775
                            _ = fi15
                            _ = v16
                        }
                    }
                case 1:
                    v17 := GenS4{F0: -76, F1: 4294967295, F2: 60}
                    _ = v17
                }
                var v18 int32 = 1
                _ = v14
                _ = v18
            }
            if false {
                v19 := GenS2{F0: 397, F1: 255, F2: 2147483647}
                _ = v19
            }
        } else {
            var v20 uint8 = p0
            v21 := []int{32767, 128}
            v22 := GenS5{F0: 32767, F1: 1000}
            _ = v20
            _ = v21
            _ = v22
        }
    case 569:
        var v23 float32 = float32(1.5)
        _ = v23
    default:
        if (736 != 160) {
            var v24 uint64 = 197
            v25 := GenS6{F0: 18446744073709551615, F1: 9223372036854775807, F2: 338}
            var v26 int16 = -1
            _ = v24
            _ = v25
            _ = v26
        }
    }
    _ = p0
    _ = v10
    _ = v11
    return 405
}

func main() {
    v1 := make(map[string]string)
    var v2 bool = false
    var v3 uint16 = 42
    v4 := []uint32{uint32(v3), 2, 772}
    v2 = true
    if v2 {
        v27 := GenS1{F0: 1, F1: 570, F2: 9223372036854775807}
        var v28 int64 = -9223372036854775808
        _ = v27
        _ = v28
    }
    v2 = v2
    if ((!(128 == 42)) || (!true)) {
        switch 32768 {
        case -731:
            fmt.Println(v2, v2)
        case -1:
            v2 = true
        case -2147483648:
            v29 := GenS2{F0: 2, F1: genFn0(65535, 2, 42), F2: 0}
            _ = v29
        }
        for fi30 := 0; fi30 < 13; fi30++ {
            var v31 string = ("" + "\n\t")
            var v32 float64 = (0.0 - 3.14159)
            v33 := v1[(("foo bar" + "foo bar") + ("\n\t" + (v31 + (v31 + (v31 + ((v31 + v31) + ("" + (((v31 + (v31 + "\n\t")) + (v31 + (v31 + v31))) + "a"))))))))]
            _ = fi30
            _ = v31
            _ = v32
            _ = v33
        }
        fmt.Println(v2, v3)
    } else {
        v2 = ((!(false && ((!(237 > 2)) && v2))) || (!(int32(v3) >= 1)))
        switch "a" {
        case "\n\t":
            if len(v4) > 0 {
                v4[0] = 4294967295
            }
        case "foo bar":
            for fi34 := 0; fi34 < 15; fi34++ {
                var v35 float64 = (0.1 + float64(0.1))
                fmt.Println(v2, v2)
                var v36 int8 = 42
                _ = fi34
                _ = v35
                _ = v36
            }
        case "":
            var v37 float64 = (0.1 + 1.5)
            _ = v37
        default:
            v3 = v3
        }
    }
    fmt.Println(v2, v3)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch 127 {
    case 0:
        v3 = v3
    case -2147483648:
        v3 -= (v3 * genFn0(340, 178, 0))
    }
    v38 := GenS5{F0: 1, F1: 32767}
    var v39 bool = v2
    v40 := GenS7{F0: 171}
    switch (("\n\t" + "hello") + ("a" + "hello")) {
    case "hello":
        if true {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        } else {
            switch "hello" {
            case "\n\t":
                var v41 int8 = 0
                _ = v41
            case "hello":
                switch "\n\t" {
                case "hello":
                    var v42 int = 2147483647
                    _ = v42
                }
            case "a":
                v1["\n\t"] = "a"
            default:
                v2 = true
            }
        }
    default:
        var v43 uint = 0
        _ = v43
    }
    v2 = false
    if true {
        var v44 int8 = -128
        if (0 > -1) {
            var v45 uint64 = 51
            _ = v45
        } else {
            for fi46 := 0; fi46 < 8; fi46++ {
                for fi47 := 0; fi47 < 16; fi47++ {
                    v44 = v44
                    v2 = ((float64(((0.0 / (-2.25 + 1e10)) - ((float32(1.5) * 0.0) + (-2.25 * (0.0 / 3.14159))))) * 0.0) == 3.14159)
                    _ = fi47
                }
                var v48 string = "foo bar"
                _ = fi46
                _ = v48
            }
        }
        v44 += v44
        _ = v44
    } else {
        var v49 uint16 = (v3 - genFn0(65535, 0, 0))
        _ = v49
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
}
