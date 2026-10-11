package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 uint64; F2 int }
type GenS4 struct { F0 int; F1 int16 }
type GenS5 struct { F0 int16; F1 int32; F2 bool }
type GenS6 struct { F0 uint64; F1 uint64; F2 uint16 }
func genFn0(p0 uint8, p1 string, p2 int8) string {
    for fi7 := 0; fi7 < 13; fi7++ {
        switch fi7 {
        case -301:
            p1 = "foo bar"
        default:
            var v8 uint8 = 18
            _ = v8
        }
        v9 := []int{(fi7 / ((32767) | 1)), (fi7 ^ (fi7 + 127)), (^(fi7)), 748}
        var v10 int = (fi7 ^ fi7)
        _ = fi7
        _ = v10
        _ = v9
    }
    p1 = p1
    p1 = p1
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    return "\n\t"
}
func genFn1(p0 uint8, p1 int) uint32 {
    if false {
        p0 += p0
        p1 = (-(p1))
        p1 -= -152
    } else {
        if false {
            if (true && true) {
                p1 -= 4294967295
            } else {
                var v11 uint = 65535
                fmt.Println(v11, v11)
                var v12 uint64 = 42
                _ = v11
                _ = v12
            }
        } else {
            var v13 uint8 = 2
            for fi14 := 0; fi14 < 18; fi14++ {
                p1 += (p1 | p1)
                _ = fi14
            }
            if false {
                var v15 int16 = -32768
                _ = v15
            } else {
                if (0 > 2) {
                    v16 := GenS1{F0: 32767, F1: 568}
                    switch "foo bar" {
                    case "hello":
                        v17 := []int{p1, p1}
                        _ = v17
                    case "foo bar":
                        v13 = 127
                    case "a":
                        v13 = p0
                    }
                    v18 := []int{p1, p1}
                    _ = v16
                    _ = v18
                } else {
                    v19 := []int{(p1 * (p1 & p1)), 127, p1}
                    _ = v19
                }
                v20 := GenS2{F0: 0}
                _ = v20
            }
            _ = v13
        }
        fmt.Println(p1, p0)
        for fi21 := 0; fi21 < 9; fi21++ {
            p1 = fi21
            _ = fi21
        }
    }
    switch p1 {
    case 934:
        if (false || ((!false) || (true && true))) {
            v22 := []int{p1, 32767}
            fmt.Println(p1)
            p0 += (p0 << ((358) & 63))
            _ = v22
        } else {
            for fi23 := 0; fi23 < 7; fi23++ {
                v24 := GenS3{F0: 2147483647, F1: 18446744073709551615, F2: (-(fi23))}
                fmt.Println(p0, p1)
                _ = fi23
                _ = v24
            }
        }
    }
    if (true && true) {
        var v25 uint64 = 4294967295
        for fi26 := 0; fi26 < 3; fi26++ {
            fmt.Println(v25)
            _ = fi26
        }
        _ = v25
    }
    v27 := []int{(^(p1)), (p1 % (((p1 << ((18446744073709551615) & 63))) | 1)), (p1 >> ((0) & 63)), (p1 | p1)}
    _ = p0
    _ = p1
    _ = v27
    return 0
}

func main() {
    var v1 int64 = 42
    v2 := []string{"\n\t", ("hello" + "a"), "foo bar", "hello"}
    v3 := make(map[string]string)
    var v4 uint64 = 42
    v5 := make(map[int]string)
    var v6 string = "hello"
    if (false || false) {
        var v28 string = v6
        var v29 bool = false
        fmt.Println(v1, v1)
        _ = v28
        _ = v29
    } else {
        switch 967 {
        case -1:
            v5[-794] = (v6 + v6)
        case 330:
            delete(v3, "hello")
        default:
            v5[9223372036854775807] = v6
        }
        fmt.Println(v1, v1, v6)
    }
    for fi30 := 0; fi30 < 18; fi30++ {
        fmt.Println(v4)
        v31 := GenS4{F0: fi30, F1: -32768}
        _ = fi30
        _ = v31
    }
    var v32 uint16 = 255
    for fi33 := 0; fi33 < 8; fi33++ {
        v34 := GenS3{F0: -2147483648, F1: 0, F2: fi33}
        _ = fi33
        _ = v34
    }
    v1 = v1
    fmt.Println(v32)
    fmt.Println(v4, v1)
    if (42 <= 0) {
        if len(v2) > 0 {
            v2[0] = v6
        }
        if ((42 <= -931) && true) {
            var v35 int = 4294967295
            v35 -= v35
            var v36 float64 = ((float64((0.1 * (1.5 / 1.0))) + 1.5) * float64(1e10))
            _ = v35
            _ = v36
        } else {
            if (2147483648 < 2147483647) {
                var v37 string = v6
                _ = v37
            } else {
                var v38 int64 = (^(v1))
                var v39 string = ((("hello" + v6) + v6) + v6)
                _ = v38
                _ = v39
            }
        }
    } else {
        v1 = 468
        fmt.Println(v32, v4)
    }
    var v40 uint16 = (v32 >> ((2) & 63))
    v41 := GenS5{F0: 172, F1: -1, F2: false}
    var v42 uint64 = (v4 / ((v4) | 1))
    v43 := v3[(v6 + ((((((("a" + v6) + ("\n\t" + (("\n\t" + ("hello" + (v6 + v6))) + "a"))) + ("foo bar" + v6)) + (((v6 + (v6 + v6)) + (((v6 + ("hello" + "foo bar")) + ("foo bar" + (v6 + "foo bar"))) + v6)) + v6)) + v6) + v6) + ("hello" + "")))]
    var v44 int32 = 2147483647
    if len(v2) > 0 {
        v2[0] = v43
    }
    if false {
        fmt.Println(v42)
        v45 := GenS4{F0: 168, F1: -32768}
        var v46 int8 = 42
        _ = v45
        _ = v46
    }
    switch -507 {
    case 0:
        v47 := GenS6{F0: (v4 << ((4294967295) & 63)), F1: (v4 * (^(v42))), F2: v40}
        _ = v47
    case 2147483647:
        switch 32767 {
        case 407:
            if true {
                v5[128] = ("foo bar" + v6)
            }
        }
    case 9223372036854775807:
        v5[-183] = v6
    }
    v48 := GenS4{F0: 128, F1: -871}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v48
    _ = v5
    _ = v6
}
