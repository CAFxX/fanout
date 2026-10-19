package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 uint32 }
type GenS2 struct { F0 uint64; F1 uint; F2 int }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 int; F1 int16 }
type GenS5 struct { F0 int8 }
func genFn0(p0 string, p1 int, p2 bool) int16 {
    switch ((p0 + p0) + p0) {
    case "foo bar":
        switch p1 {
        case -2147483648:
            var v6 float32 = 1e10
            _ = v6
        }
    case "":
        fmt.Println(p1)
    case "a":
        fmt.Println(p2, p2)
    }
    switch p0 {
    case "\n\t":
        for fi7 := 0; fi7 < 14; fi7++ {
            fmt.Println(p0)
            v8 := []int{32768}
            _ = fi7
            _ = v8
        }
    case "":
        if p2 {
            var v9 bool = false
            for fi10 := 0; fi10 < 20; fi10++ {
                fmt.Println(fi10, p2, v9)
                _ = fi10
            }
            var v11 uint8 = 147
            _ = v11
            _ = v9
        } else {
            for fi12 := 0; fi12 < 19; fi12++ {
                switch p0 {
                case "foo bar":
                    fmt.Println(p0)
                }
                _ = fi12
            }
            if (!(3.14159 > float32((1.5 * 0.0)))) {
                p2 = (42 == 100)
                var v13 uint32 = 1
                var v14 uint64 = 1
                _ = v13
                _ = v14
            }
        }
    }
    p1 += p1
    v15 := []int{(^(p1))}
    for fi16 := 0; fi16 < 17; fi16++ {
        p1 = fi16
        _ = fi16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return -32768
}

func main() {
    var v1 uint8 = 147
    v2 := []string{"foo bar", ""}
    var v3 uint64 = 1
    var v4 uint64 = (v3 | (v3 + 426))
    var v5 float32 = 0.1
    v17 := GenS1{F0: 0, F1: 128, F2: 4294967295}
    for fi18 := 0; fi18 < 2; fi18++ {
        switch "foo bar" {
        case "hello":
            var v19 int64 = 0
            _ = v19
        case "foo bar":
            v5 = 1.5
        case "a":
            for fi20 := 0; fi20 < 14; fi20++ {
                var v21 uint16 = 2
                _ = fi20
                _ = v21
            }
        default:
            v3 -= (^(v4))
        }
        fmt.Println(v4)
        _ = fi18
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v22 int64 = -237
    if (-32768 >= 32767) {
        v22 = v22
        switch ("\n\t" + "foo bar") {
        case "foo bar":
            for fi23 := 0; fi23 < 10; fi23++ {
                for fi24 := 0; fi24 < 9; fi24++ {
                    v2 = append(v2, "foo bar")
                    v3 = (v4 & (v3 | 18446744073709551615))
                    v22 += v22
                    _ = fi24
                }
                v25 := GenS1{F0: 18446744073709551615, F1: (fi23 / (((fi23 >> ((516) & 63))) | 1)), F2: 513}
                var v26 uint = 0
                _ = fi23
                _ = v25
                _ = v26
            }
        }
        v1 -= 255
    }
    if false {
        var v27 uint = 225
        switch 32767 {
        case -801:
            for fi28 := 0; fi28 < 8; fi28++ {
                v29 := GenS2{F0: v4, F1: v27, F2: fi28}
                switch (fi28 - fi28) {
                case -9223372036854775808:
                    var v30 int16 = -32768
                    _ = v30
                case 128:
                    v4 += v3
                case 672:
                    v31 := GenS1{F0: v27, F1: -2147483648, F2: 2}
                    _ = v31
                }
                if (326 == -2147483648) {
                    var v32 int16 = genFn0("foo bar", 9223372036854775807, false)
                    v5 = v5
                    v2 = append(v2, "hello")
                    _ = v32
                } else {
                    var v33 float32 = v5
                    _ = v33
                }
                _ = fi28
                _ = v29
            }
        case 11:
            var v34 int8 = -128
            _ = v34
        default:
            var v35 int64 = v22
            _ = v35
        }
        v2 = append(v2, "")
        _ = v27
    } else {
        if len(v2) > 0 {
            v2[0] = "hello"
        }
    }
    switch "" {
    case "foo bar":
        v36 := GenS3{F0: 77, F1: -850}
        _ = v36
    case "a":
        var v37 uint16 = 198
        _ = v37
    case "hello":
        switch 32768 {
        case 32768:
            for fi38 := 0; fi38 < 1; fi38++ {
                var v39 string = "a"
                _ = fi38
                _ = v39
            }
        case 128:
            if ((v3 % (((v4 - v3)) | 1)) > 1) {
                v40 := GenS4{F0: 82, F1: int16(v1)}
                v41 := GenS5{F0: 0}
                _ = v40
                _ = v41
            } else {
                v22 = 1
                var v42 float64 = (float64(0.1) - 0.0)
                v4 = 4294967295
                _ = v42
            }
        case 4294967295:
            var v43 string = ("\n\t" + "\n\t")
            _ = v43
        }
    }
    for fi44 := 0; fi44 < 10; fi44++ {
        v5 = v5
        _ = fi44
    }
    v5 = ((v5 / float32(-2.25)) + (float32(0.1) / v5))
    var v45 uint = 0
    switch "foo bar" {
    case "a":
        v3 += (v3 ^ (v3 / ((v3) | 1)))
    case "foo bar":
        if false {
            v5 -= ((v5 - ((((v5 - float32(0.0)) + (1e10 / -2.25)) * v5) * (v5 - 0.0))) * 3.14159)
            v5 = ((0.0 * -2.25) + ((v5 + v5) + (1e10 / v5)))
        }
    }
    switch ("hello" + ("a" + "")) {
    case "":
        v4 = (v3 << ((42) & 63))
    }
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    var v46 uint32 = 4294967295
    _ = len(v2)
    switch 0 {
    case 239:
        v47 := GenS3{F0: (v45 + v45), F1: -32768}
        _ = v47
    case -9223372036854775808:
        if ((v22 & -9223372036854775808) == v22) {
            var v48 float32 = 1e10
            _ = v48
        }
    }
    switch -7 {
    case 232:
        for fi49 := 0; fi49 < 15; fi49++ {
            var v50 int = -9223372036854775808
            if (0 != v22) {
                if false {
                    var v51 bool = false
                    v46 = (v46 * (v46 + v46))
                    var v52 float32 = v5
                    _ = v51
                    _ = v52
                } else {
                    var v53 int = (v50 - fi49)
                    var v54 uint8 = v1
                    for fi55 := 0; fi55 < 14; fi55++ {
                        var v56 bool = ((^(v46)) < (v46 * v46))
                        _ = fi55
                        _ = v56
                    }
                    _ = v53
                    _ = v54
                }
                for fi57 := 0; fi57 < 15; fi57++ {
                    var v58 int16 = 1000
                    v2 = append(v2, "a")
                    var v59 uint32 = v46
                    switch ("" + (("\n\t" + "\n\t") + ((("foo bar" + "\n\t") + "\n\t") + ("hello" + "hello")))) {
                    case "\n\t":
                        var v60 uint = (v45 * (v45 / (((^(v45))) | 1)))
                        _ = v60
                    case "a":
                        v58 = (v58 % (((v58 >> (((v45 + (v45 / ((v45) | 1)))) & 63))) | 1))
                        v22 = (-(v22))
                        v61 := GenS5{F0: 100}
                        _ = v61
                    default:
                        var v62 bool = false
                        v5 -= (0.0 + float32(0.0))
                        var v63 uint16 = 0
                        _ = v62
                        _ = v63
                    }
                    v59 = 1
                    _ = fi57
                    _ = v58
                    _ = v59
                }
                fmt.Println(v45, v3)
            }
            _ = fi49
            _ = v50
        }
    case 42:
        fmt.Println(v46, v1)
    default:
        var v64 int = -2147483648
        _ = v64
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
}
