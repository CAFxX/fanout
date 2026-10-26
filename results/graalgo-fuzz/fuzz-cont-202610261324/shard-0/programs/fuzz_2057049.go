package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint8 }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 uint) uint32 {
    var v6 int32 = 0
    v7 := GenS1{F0: (p0 / ((42) | 1)), F1: -128, F2: 1}
    v8 := GenS1{F0: p0, F1: 56, F2: 173}
    switch "\n\t" {
    case "hello":
        for fi9 := 0; fi9 < 5; fi9++ {
            for fi10 := 0; fi10 < 4; fi10++ {
                v11 := []int{(^(fi10)), -1, 127}
                var v12 uint = p0
                v6 += 1
                if len(v11) > 0 {
                    v11[0] = (fi9 * fi9)
                }
                _ = len(v11)
                var v13 uint32 = 2147483648
                _ = len(v11)
                _ = fi10
                _ = v11
                _ = v12
                _ = v13
            }
            var v14 int32 = (^(v6))
            _ = fi9
            _ = v14
        }
    case "a":
        if false {
            if (2 <= 32767) {
                var v15 int64 = 0
                v6 = v6
                fmt.Println(v6)
                _ = v15
            } else {
                v6 = v6
                for fi16 := 0; fi16 < 8; fi16++ {
                    v17 := GenS1{F0: p0, F1: 42, F2: 42}
                    _ = fi16
                    _ = v17
                }
            }
            v18 := GenS1{F0: 255, F1: 0, F2: uint8(p0)}
            var v19 int8 = -1
            _ = v18
            _ = v19
        }
    default:
        if true {
            v6 += (v6 - (v6 | (-(v6))))
        } else {
            var v20 string = ("" + "hello")
            _ = v20
        }
    }
    if true {
        fmt.Println(p0, p0)
    } else {
        v21 := GenS1{F0: 255, F1: -128, F2: 178}
        var v22 int64 = 9223372036854775807
        fmt.Println(p0, p0, v6)
        _ = v21
        _ = v22
    }
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 float32 = float32((0.0 + (3.14159 + (-2.25 / float64((0.1 + -2.25))))))
    v2 := make(map[string]bool)
    var v3 uint32 = 42
    var v4 bool = false
    var v5 uint64 = 186
    var v23 int8 = 127
    var v24 int32 = 2147483647
    v25 := GenS1{F0: 65535, F1: (v23 - (v23 & v23)), F2: uint8(v3)}
    v26 := []int{0, 128, 647}
    v2["a"] = v4
    var v27 int32 = v24
    var v28 float64 = -2.25
    if v4 {
        switch 803 {
        case -9223372036854775808:
            v29 := v2["a"]
            _ = v29
        default:
            var v30 int64 = 0
            _ = v30
        }
    }
    var v31 int64 = 0
    switch 42 {
    case 2147483647:
        v28 += (((float64(1.5) - 0.0) / ((1e10 + v28) / 3.14159)) - 3.14159)
    case 128:
        v24 = v27
    }
    var v32 int32 = v27
    fmt.Println(v31, v28, v23)
    if (!(547 > 2)) {
        if (1 < 1) {
            if len(v26) >= 2 {
                v26 = v26[1:len(v26)]
            }
            v26 = append(v26, -2147483648)
            fmt.Println(v28, v28, v4)
        }
        for fi33 := 0; fi33 < 18; fi33++ {
            _ = len(v2)
            _ = fi33
        }
        _ = len(v2)
    } else {
        switch -858 {
        case -2147483648:
            v28 -= -2.25
        case 9223372036854775807:
            v23 += (v23 >> ((674) & 63))
        default:
            for fi34 := 0; fi34 < 15; fi34++ {
                v24 = (v24 * (^(v32)))
                v35 := GenS2{F0: (v5 & 4294967295), F1: 128}
                _ = fi34
                _ = v35
            }
        }
    }
    if (!false) {
        for fi36 := 0; fi36 < 2; fi36++ {
            if len(v26) >= 2 {
                v26 = v26[1:len(v26)]
            }
            v24 = 0
            v31 = (-(v31))
            _ = fi36
        }
        switch "" {
        case "hello":
            _ = len(v2)
        case "":
            v1 += v1
        }
    } else {
        v28 -= v28
        var v37 int16 = -215
        _ = v37
    }
    var v38 string = "foo bar"
    if v4 {
        for fi39 := 0; fi39 < 10; fi39++ {
            v24 = (v27 & v24)
            v3 += (v3 ^ (^(v3)))
            _ = fi39
        }
        v40 := GenS2{F0: 4294967295, F1: 42}
        _ = v40
    } else {
        fmt.Println(v31, v32, v1)
        if (v4 && ((v32 == -307) && v4)) {
            var v41 uint8 = 127
            if len(v26) > 0 {
                v26[0] = 2147483647
            }
            var v42 float32 = float32(v28)
            _ = v41
            _ = v42
        } else {
            for fi43 := 0; fi43 < 10; fi43++ {
                delete(v2, ("a" + ("hello" + v38)))
                v5 = (v5 ^ v5)
                _ = fi43
            }
            v44 := GenS3{F0: 0}
            _ = v44
        }
        for fi45 := 0; fi45 < 20; fi45++ {
            v5 = v5
            if len(v26) >= 2 {
                v26 = v26[1:len(v26)]
            }
            _ = fi45
        }
    }
    var v46 int8 = v23
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v38
    _ = v4
    _ = v46
    _ = v5
}
