package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 uint16; F1 int; F2 int64 }
type GenS5 struct { F0 int32 }
func genFn0(p0 string, p1 uint, p2 uint64) uint64 {
    v4 := []int{0}
    if true {
        var v5 bool = true
        if (-2.25 <= ((((((3.14159 / -2.25) / 0.1) / (3.14159 - -2.25)) * 3.14159) * float64(float32(((0.1 + 3.14159) * 0.0)))) / 1e10)) {
            fmt.Println(v5)
            v6 := GenS1{F0: -9223372036854775808, F1: 2147483647}
            if (!v5) {
                v7 := GenS1{F0: 127, F1: 0}
                _ = v7
            } else {
                fmt.Println(v5, p1, p1)
                if false {
                    v8 := GenS2{F0: p1, F1: 42}
                    _ = v8
                }
            }
            _ = v6
        }
        for fi9 := 0; fi9 < 19; fi9++ {
            p1 = 65535
            _ = fi9
        }
        _ = v5
    }
    switch ((p0 + p0) + (p0 + p0)) {
    case "foo bar":
        p2 = 1
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return (p2 | (p2 ^ (^(p2))))
}

func main() {
    v1 := make(map[int]int)
    var v2 uint64 = 926
    v3 := []int{128, 42, 32767, 4294967295}
    var v10 bool = false
    if (39 > 65535) {
        v2 -= 0
    }
    v2 = 2
    v11 := GenS3{F0: v2, F1: "foo bar"}
    if (0 > 127) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v10 = (1 < 0)
        v2 -= 747
    } else {
        for fi12 := 0; fi12 < 18; fi12++ {
            var v13 float32 = (float32(3.14159) + (0.0 / 3.14159))
            v14 := GenS4{F0: 55, F1: -4, F2: 0}
            _ = fi12
            _ = v13
            _ = v14
        }
        var v15 string = ("a" + "\n\t")
        v16 := GenS3{F0: v2, F1: ""}
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 3; fi17++ {
        if true {
            for fi18 := 0; fi18 < 10; fi18++ {
                var v19 float64 = 0.1
                v10 = false
                fmt.Println(v2)
                v20 := GenS5{F0: 1}
                if v10 {
                    fmt.Println(v2, fi18, fi18)
                    v10 = (65535 != 65535)
                } else {
                    if (-2147483648 == 42) {
                        v10 = false
                    }
                    switch fi18 {
                    case 1:
                        v19 = 1e10
                    case -208:
                        var v21 uint16 = 666
                        _ = v21
                    }
                    fmt.Println(v2)
                }
                _ = fi18
                _ = v19
                _ = v20
            }
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        _ = fi17
    }
    for fi22 := 0; fi22 < 4; fi22++ {
        fmt.Println(fi22, v10)
        for fi23 := 0; fi23 < 15; fi23++ {
            v2 -= v2
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi23
        }
        _ = fi22
    }
    var v24 float32 = 3.14159
    _ = len(v3)
    if false {
        v25 := GenS2{F0: 65535, F1: 9223372036854775807}
        for fi26 := 0; fi26 < 11; fi26++ {
            v27 := GenS4{F0: 242, F1: (fi26 >> ((58) & 63)), F2: 1}
            _ = fi26
            _ = v27
        }
        _ = v25
    } else {
        for fi28 := 0; fi28 < 18; fi28++ {
            fmt.Println(v2)
            _ = fi28
        }
    }
    fmt.Println(v2)
    fmt.Println(v24, v24, v24)
    switch "a" {
    case "\n\t":
        var v29 uint = 528
        _ = v29
    }
    var v30 float32 = (float32((0.0 - 0.0)) / -2.25)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v24
    _ = v3
    _ = v30
}
