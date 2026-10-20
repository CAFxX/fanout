package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int16; F1 bool }
type GenS5 struct { F0 int64; F1 uint64; F2 uint16 }
func genFn0() int32 {
    switch 9223372036854775807 {
    case -198:
        if (1 < 127) {
            var v7 uint32 = 42
            _ = v7
        } else {
            if true {
                if (42 > 128) {
                    var v8 uint64 = 536
                    for fi9 := 0; fi9 < 17; fi9++ {
                        switch fi9 {
                        case 294:
                            v10 := GenS1{F0: -488, F1: -29, F2: "hello"}
                            _ = v10
                        default:
                            v11 := []int{}
                            _ = v11
                        }
                        var v12 float32 = 0.1
                        var v13 uint32 = 0
                        _ = fi9
                        _ = v12
                        _ = v13
                    }
                    v8 = v8
                    _ = v8
                }
            }
            var v14 float32 = ((((0.1 * (3.14159 - 0.1)) / 1.0) * 0.0) / 3.14159)
            _ = v14
        }
    case 127:
        switch "hello" {
        case "":
            if (!(0.0 != 0.0)) {
                v15 := []int{821}
                v16 := GenS1{F0: 42, F1: 2147483647, F2: ""}
                _ = v15
                _ = v16
            }
        default:
            var v17 int = 127
            _ = v17
        }
    }
    var v18 float64 = 0.0
    fmt.Println(v18)
    _ = v18
    return 1
}
func genFn1(p0 uint16, p1 uint32, p2 int) int64 {
    var v19 float64 = float64(3.14159)
    if ("a" <= "a") {
        if false {
            v20 := GenS2{F0: 0}
            for fi21 := 0; fi21 < 8; fi21++ {
                v22 := []int{(p2 | (-(p2))), p2, (fi21 | fi21)}
                v19 = v19
                v23 := GenS3{F0: (p0 * uint16(p1))}
                _ = fi21
                _ = v22
                _ = v23
            }
            _ = v20
        }
    } else {
        var v24 int16 = 42
        v25 := []int{}
        var v26 uint64 = 448
        _ = v24
        _ = v25
        _ = v26
    }
    fmt.Println(p1, p2, p2)
    for fi27 := 0; fi27 < 12; fi27++ {
        p2 -= 32768
        _ = fi27
    }
    var v28 int64 = int64(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v28
    return v28
}

func main() {
    var v1 int64 = -1
    var v2 int16 = -32768
    v3 := make(map[int]bool)
    var v4 int16 = v2
    var v5 uint32 = 2
    v6 := make(map[string]bool)
    v29 := GenS4{F0: v4, F1: true}
    if false {
        var v30 uint = 4294967295
        _ = v30
    } else {
        fmt.Println(v2, v4)
        var v31 bool = true
        var v32 uint = 2
        _ = v31
        _ = v32
    }
    for fi33 := 0; fi33 < 18; fi33++ {
        for fi34 := 0; fi34 < 18; fi34++ {
            var v35 uint64 = 18446744073709551615
            v5 -= (v5 >> ((255) & 63))
            _ = fi34
            _ = v35
        }
        if false {
            v5 = v5
            fmt.Println(v2)
            switch (fi33 ^ fi33) {
            case 32767:
                for fi36 := 0; fi36 < 17; fi36++ {
                    fmt.Println(v2, v4)
                    _ = fi36
                }
            default:
                v4 += 1
            }
        }
        _ = fi33
    }
    var v37 int32 = 1
    if true {
        v2 = v4
    }
    v38 := GenS3{F0: 0}
    v39 := GenS5{F0: v1, F1: 0, F2: 255}
    v2 += (v2 << ((276) & 63))
    fmt.Println(v1)
    delete(v6, "\n\t")
    var v40 int = 2147483647
    v2 = v4
    v41 := GenS5{F0: -502, F1: 266, F2: 2}
    for fi42 := 0; fi42 < 13; fi42++ {
        var v43 string = "hello"
        var v44 int64 = -9223372036854775808
        for fi45 := 0; fi45 < 7; fi45++ {
            v5 -= 42
            if (v5 != 2147483648) {
                var v46 string = ("hello" + v43)
                _ = v46
            }
            _ = fi45
        }
        _ = fi42
        _ = v43
        _ = v44
    }
    switch "" {
    case "":
        for fi47 := 0; fi47 < 7; fi47++ {
            fmt.Println(v40, v37)
            fmt.Println(v2)
            v48 := []int{4294967295}
            _ = fi47
            _ = v48
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
