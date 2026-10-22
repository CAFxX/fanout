package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 int8 }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint; F1 uint64 }
func genFn0(p0 int32) uint {
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 int64 = 1000
        if (!(-140 <= 42)) {
            v8 = v8
            var v9 uint8 = 42
            v10 := GenS1{F0: (fi7 << ((18446744073709551615) & 63)), F1: 2147483648, F2: 1}
            _ = v10
            _ = v9
        }
        fmt.Println(fi7, fi7, fi7)
        _ = fi7
        _ = v8
    }
    for fi11 := 0; fi11 < 17; fi11++ {
        var v12 string = ("" + (((("hello" + "a") + "hello") + "hello") + "foo bar"))
        v12 = "hello"
        _ = fi11
        _ = v12
    }
    v13 := GenS1{F0: 4294967295, F1: 0, F2: 83}
    _ = p0
    _ = v13
    return 673
}

func main() {
    v1 := make(map[string]string)
    var v2 int64 = -319
    var v3 int32 = 1
    v4 := make(map[string]string)
    v5 := []int{-2147483648}
    var v6 uint16 = 1
    fmt.Println(v2, v3)
    switch "hello" {
    case "\n\t":
        v6 -= v6
    case "a":
        switch 255 {
        case 2147483647:
            v2 += (^(v2))
        case 329:
            var v14 string = (("a" + "foo bar") + (("foo bar" + ("foo bar" + "\n\t")) + "\n\t"))
            _ = v14
        case 105:
            var v15 int = 9223372036854775807
            _ = v15
        }
    case "":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    fmt.Println(v3, v3)
    v16 := GenS2{F0: v2, F1: 2147483647}
    var v17 uint32 = 667
    for fi18 := 0; fi18 < 12; fi18++ {
        var v19 uint32 = uint32(v6)
        v19 = (v17 ^ (v17 << ((4294967295) & 63)))
        _ = fi18
        _ = v19
    }
    var v20 int32 = v3
    _ = len(v5)
    v21 := GenS1{F0: 0, F1: v17, F2: 1}
    switch "" {
    case "":
        if len(v5) > 0 {
            v5[0] = -694
        }
    case "foo bar":
        var v22 uint8 = 212
        _ = v22
    case "a":
        fmt.Println(v3, v6)
    }
    v23 := GenS3{F0: "", F1: 82}
    v24 := GenS4{F0: 85}
    switch "" {
    case "":
        var v25 int32 = (-(v20))
        _ = v25
    case "\n\t":
        var v26 string = (("hello" + "hello") + "hello")
        _ = v26
    }
    var v27 uint64 = uint64(v17)
    v28 := GenS1{F0: 1, F1: (v17 << ((255) & 63)), F2: -1}
    if false {
        switch 32767 {
        case 366:
            var v29 string = ("a" + ("" + "\n\t"))
            _ = v29
        case 4294967295:
            for fi30 := 0; fi30 < 19; fi30++ {
                v2 -= v2
                if (-128 <= 69) {
                    fmt.Println(v2, v6, fi30)
                    if len(v5) > 0 {
                        v5[0] = (fi30 >> ((2) & 63))
                    }
                }
                _ = fi30
            }
        default:
            if (!("\n\t" != "hello")) {
                v4["\n\t"] = "\n\t"
            }
        }
        if (!(false || true)) {
            var v31 int = 32768
            v32 := GenS4{F0: 83}
            v2 -= -9223372036854775808
            _ = v31
            _ = v32
        } else {
            fmt.Println(v17)
        }
    } else {
        if true {
            v17 += 2
        }
    }
    switch "hello" {
    case "a":
        var v33 string = "foo bar"
        _ = v33
    case "hello":
        switch 0 {
        case 32768:
            if false {
                if (42 != 122) {
                    switch 9223372036854775807 {
                    case 42:
                        if false {
                            v17 += v17
                        } else {
                            var v34 string = "a"
                            var v35 float32 = (0.1 / 0.1)
                            _ = v34
                            _ = v35
                        }
                    case 0:
                        var v36 bool = false
                        _ = v36
                    default:
                        if true {
                            switch "\n\t" {
                            case "":
                                var v37 int32 = (-(v3))
                                _ = v37
                            case "a":
                                var v38 bool = ((^(v20)) != 584)
                                _ = v38
                            default:
                                v39 := GenS5{F0: 356, F1: (^(v27))}
                                _ = v39
                            }
                            fmt.Println(v6, v17, v3)
                            v17 -= 577
                        }
                    }
                } else {
                    v6 = (v6 & (^(v6)))
                    fmt.Println(v2, v27, v20)
                    v17 = 2
                }
            }
        case 42:
            v40 := GenS2{F0: (v2 | (v2 >> ((1) & 63))), F1: 32767}
            _ = v40
        case -191:
            delete(v1, "a")
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v24
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
