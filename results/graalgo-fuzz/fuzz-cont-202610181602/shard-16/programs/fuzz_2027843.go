package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int }
func genFn0() uint8 {
    var v7 int = -910
    var v8 uint64 = 4294967295
    _ = v7
    _ = v8
    return 1
}

func main() {
    v1 := make(map[string]int)
    v2 := []int{-401, -2147483648}
    v3 := []uint32{295, 339}
    var v4 int8 = -53
    var v5 uint = 0
    var v6 bool = (false || (false || true))
    switch "\n\t" {
    case "hello":
        var v9 string = (("\n\t" + "") + "\n\t")
        _ = v9
    case "\n\t":
        var v10 int32 = -457
        _ = v10
    }
    switch ("foo bar" + "\n\t") {
    case "foo bar":
        switch -577 {
        case -9223372036854775808:
            for fi11 := 0; fi11 < 19; fi11++ {
                v4 = v4
                v12 := GenS1{F0: (fi11 * (fi11 & fi11)), F1: 9223372036854775807}
                v13 := GenS2{F0: v5, F1: -1}
                _ = fi11
                _ = v12
                _ = v13
            }
        default:
            v5 += 2
        }
    case "hello":
        v1["a"] = 128
    }
    v5 += v5
    v6 = ((^(v5)) <= (v5 ^ v5))
    v5 -= (^(v5))
    v6 = v6
    fmt.Println(v5, v5)
    var v14 uint8 = 82
    if v6 {
        if (("" + "a") == "\n\t") {
            var v15 bool = (-1 > 42)
            _ = v15
        } else {
            for fi16 := 0; fi16 < 14; fi16++ {
                var v17 uint16 = 2
                v17 -= (v17 ^ (v17 ^ uint16(v5)))
                _ = len(v2)
                _ = fi16
                _ = v17
            }
        }
        for fi18 := 0; fi18 < 13; fi18++ {
            v6 = (1 == (v4 >> ((42) & 63)))
            _ = fi18
        }
        if (int64(v14) != -1) {
            _ = len(v3)
        } else {
            var v19 int64 = -1
            v3 = append(v3, 813)
            var v20 int16 = 32767
            _ = v19
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 16; fi21++ {
        v5 += (v5 - 255)
        v22 := GenS1{F0: (fi21 ^ (fi21 & fi21)), F1: 42}
        _ = fi21
        _ = v22
    }
    v3 = append(v3, 347)
    if ((!v6) || (-789 != -9223372036854775808)) {
        v5 = v5
        if v6 {
            var v23 uint16 = 65535
            v14 = v14
            v6 = (2147483647 < 1)
            _ = v23
        } else {
            fmt.Println(v4)
            var v24 uint32 = 42
            _ = v24
        }
        v5 -= 315
    } else {
        if (!v6) {
            if v6 {
                v5 += 42
                v6 = v6
            }
        }
        v25 := GenS3{F0: -286}
        _ = v25
    }
    var v26 uint64 = 2
    switch (((("a" + ("hello" + "hello")) + "foo bar") + "a") + "\n\t") {
    case "hello":
        for fi27 := 0; fi27 < 19; fi27++ {
            v4 -= (v4 & 127)
            if ("" != "") {
                v28 := v1["hello"]
                var v29 uint64 = v26
                _ = v28
                _ = v29
            }
            if (((3.14159 + -2.25) == 3.14159) && (v6 && true)) {
                var v30 float64 = 3.14159
                v31 := GenS4{F0: (fi27 << ((v5) & 63))}
                _ = v30
                _ = v31
            }
            _ = fi27
        }
    default:
        if len(v2) > 0 {
            v2[0] = -2147483648
        }
    }
    if ((float32((0.1 - float64((0.1 * 1.5)))) * ((3.14159 + 0.1) - 0.1)) >= -2.25) {
        for fi32 := 0; fi32 < 7; fi32++ {
            if (true || false) {
                var v33 uint8 = 2
                for fi34 := 0; fi34 < 9; fi34++ {
                    v5 = 0
                    if (!false) {
                        v26 = 18446744073709551615
                        if ((v5 * v5) > v5) {
                            var v35 uint8 = (v14 / ((genFn0()) | 1))
                            _ = v35
                        }
                    }
                    _ = len(v2)
                    v33 += v33
                    _ = fi34
                }
                _ = v33
            }
            v36 := GenS2{F0: 0, F1: 1}
            _ = fi32
            _ = v36
        }
    } else {
        for fi37 := 0; fi37 < 5; fi37++ {
            var v38 uint64 = (^(v26))
            _ = fi37
            _ = v38
        }
        _ = len(v2)
    }
    delete(v1, "a")
    var v39 uint64 = 2
    switch -1 {
    case -2147483648:
        v4 = (v4 * (v4 - (v4 + v4)))
    case 2147483647:
        var v40 uint64 = (^(v39))
        _ = v40
    case 343:
        switch "" {
        case "":
            v41 := GenS3{F0: 0}
            _ = v41
        case "\n\t":
            for fi42 := 0; fi42 < 9; fi42++ {
                var v43 float32 = float32(1.5)
                _ = fi42
                _ = v43
            }
        case "foo bar":
            v3 = append(v3, 2147483648)
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v26
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
