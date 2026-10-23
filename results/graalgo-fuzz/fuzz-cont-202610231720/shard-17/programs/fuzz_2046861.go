package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64; F1 string; F2 string }
type GenS4 struct { F0 int32; F1 int64 }
func genFn0(p0 bool, p1 uint8, p2 bool) bool {
    var v6 float64 = float64((1.5 / -2.25))
    var v7 int32 = 219
    v8 := GenS1{F0: 124}
    var v9 uint16 = 913
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return (0 < -1)
}

func main() {
    var v1 int64 = 42
    v2 := []int{128, 1, -1, -2147483648}
    var v3 uint = 617
    var v4 uint16 = 0
    v5 := make(map[string]int)
    if true {
        var v10 uint16 = (v4 / (((v4 ^ v4)) | 1))
        _ = len(v5)
        for fi11 := 0; fi11 < 20; fi11++ {
            v4 -= (v4 ^ v4)
            if false {
                v1 = -9223372036854775808
            } else {
                if true {
                    v5[("foo bar" + "foo bar")] = fi11
                    var v12 float32 = ((((-2.25 - 1.5) - 1.5) / -2.25) * 3.14159)
                    _ = v12
                }
            }
            _ = fi11
        }
        _ = v10
    } else {
        var v13 uint32 = 428
        var v14 uint32 = 4294967295
        _ = v13
        _ = v14
    }
    v15 := GenS2{F0: 38}
    v4 += v4
    if true {
        v16 := GenS3{F0: v1, F1: (("hello" + "hello") + ("" + "hello")), F2: ""}
        _ = v16
    }
    var v17 bool = false
    var v18 int64 = (v1 - (v1 - 0))
    switch "hello" {
    case "hello":
        v19 := GenS2{F0: 42}
        _ = v19
    case "":
        v1 += (v18 << ((v3) & 63))
    case "\n\t":
        if v17 {
            if len(v2) > 0 {
                v2[0] = 1
            }
        } else {
            for fi20 := 0; fi20 < 14; fi20++ {
                v17 = false
                _ = fi20
            }
            v21 := GenS2{F0: 100}
            _ = v21
        }
    }
    v22 := v5[(("" + ("a" + "a")) + "\n\t")]
    v17 = v17
    switch "foo bar" {
    case "hello":
        for fi23 := 0; fi23 < 17; fi23++ {
            if len(v2) > 0 {
                v2[0] = (v22 | v22)
            }
            v1 -= v1
            _ = fi23
        }
    case "\n\t":
        for fi24 := 0; fi24 < 5; fi24++ {
            v25 := GenS2{F0: -128}
            var v26 int = v22
            _ = fi24
            _ = v25
            _ = v26
        }
    case "a":
        if len(v2) > 0 {
            v2[0] = v22
        }
    default:
        for fi27 := 0; fi27 < 8; fi27++ {
            v17 = false
            for fi28 := 0; fi28 < 2; fi28++ {
                v5[("foo bar" + "foo bar")] = v22
                delete(v5, "foo bar")
                v29 := GenS1{F0: 2147483648}
                _ = fi28
                _ = v29
            }
            _ = fi27
        }
    }
    switch (v22 + v22) {
    case 128:
        switch "a" {
        case "":
            var v30 float64 = (float64(3.14159) - 1e10)
            _ = v30
        case "foo bar":
            switch v22 {
            case 32768:
                var v31 int32 = 2147483647
                _ = v31
            case -9223372036854775808:
                fmt.Println(v4)
            case -107:
                v2 = append(v2, (v22 | (v22 + v22)))
            }
        }
    }
    v22 += v22
    if (!false) {
        v17 = v17
    } else {
        _ = len(v2)
        switch "a" {
        case "a":
            v32 := GenS4{F0: -391, F1: (-(v1))}
            _ = v32
        case "\n\t":
            switch -9223372036854775808 {
            case 4294967295:
                for fi33 := 0; fi33 < 12; fi33++ {
                    for fi34 := 0; fi34 < 16; fi34++ {
                        v5["a"] = fi34
                        _ = fi34
                    }
                    _ = fi33
                }
            default:
                var v35 uint8 = 127
                _ = v35
            }
        case "":
            v36 := GenS3{F0: v18, F1: "", F2: "foo bar"}
            _ = v36
        default:
            for fi37 := 0; fi37 < 2; fi37++ {
                var v38 uint64 = 2
                _ = fi37
                _ = v38
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
