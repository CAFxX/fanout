package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16; F1 string; F2 uint8 }
type GenS4 struct { F0 uint64 }
func genFn0(p0 int64) int64 {
    if (67 <= 0) {
        var v7 int64 = (p0 & 42)
        if (!false) {
            for fi8 := 0; fi8 < 10; fi8++ {
                v9 := []int{(^(fi8)), (fi8 ^ -687), 128}
                p0 += p0
                fmt.Println(fi8, p0)
                _ = fi8
                _ = v9
            }
            var v10 int16 = 0
            _ = v10
        } else {
            var v11 uint64 = 40
            switch 9223372036854775807 {
            case 128:
                v12 := GenS1{F0: (^(v7)), F1: 154}
                _ = v12
            case -130:
                if true {
                    fmt.Println(p0)
                    for fi13 := 0; fi13 < 13; fi13++ {
                        v14 := GenS1{F0: (v7 >> ((0) & 63)), F1: 825}
                        v11 -= (v11 + v11)
                        v15 := GenS2{F0: false}
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                    v16 := []int{1}
                    _ = v16
                } else {
                    fmt.Println(v7)
                }
            case 32768:
                v17 := GenS2{F0: true}
                _ = v17
            default:
                switch -2147483648 {
                case 1:
                    fmt.Println(v11)
                default:
                    for fi18 := 0; fi18 < 9; fi18++ {
                        v7 -= 1
                        switch ("hello" + (("foo bar" + ((("foo bar" + ("a" + "")) + ("foo bar" + "foo bar")) + "")) + "a")) {
                        case "a":
                            fmt.Println(fi18, v11, v7)
                        case "":
                            v19 := GenS1{F0: p0, F1: -1}
                            _ = v19
                        case "hello":
                            fmt.Println(fi18)
                        }
                        _ = fi18
                    }
                }
            }
            fmt.Println(p0, v7, v7)
            _ = v11
        }
        var v20 string = "a"
        _ = v20
        _ = v7
    } else {
        if (("\n\t" + "a") == "") {
            p0 -= p0
            v21 := []int{}
            _ = v21
        }
        var v22 uint8 = 255
        v23 := GenS1{F0: p0, F1: 42}
        _ = v22
        _ = v23
    }
    p0 -= p0
    v24 := GenS1{F0: (p0 | (^(p0))), F1: 42}
    var v25 float64 = 1e10
    var v26 uint = 748
    _ = p0
    _ = v24
    _ = v25
    _ = v26
    return p0
}

func main() {
    var v1 int = 32768
    var v2 uint = 40
    v3 := make(map[int]bool)
    var v4 float32 = ((-2.25 / (((0.0 - -2.25) / -2.25) / (1e10 + 0.1))) / (1.5 + 0.1))
    v5 := make(map[int]string)
    var v6 int64 = -837
    v1 -= (-(v1))
    if (-454 != -63) {
        var v27 uint = (^(v2))
        _ = v27
    } else {
        v4 += 3.14159
        fmt.Println(v2, v6)
    }
    if (769 >= 1) {
        switch "" {
        case "hello":
            var v28 uint64 = 2
            _ = v28
        case "a":
            var v29 uint8 = 0
            _ = v29
        case "\n\t":
            fmt.Println(v1)
        default:
            fmt.Println(v4, v1, v2)
        }
        v2 += v2
    } else {
        if false {
            _ = len(v5)
        }
    }
    v6 -= genFn0(176)
    delete(v5, (v1 << (((v2 * v2)) & 63)))
    var v30 uint = (v2 + (^(v2)))
    var v31 float32 = (3.14159 / v4)
    switch v1 {
    case -9223372036854775808:
        var v32 bool = true
        _ = v32
    case 469:
        fmt.Println(v4, v2, v1)
    }
    v33 := GenS3{F0: 1, F1: "", F2: 127}
    v4 -= v31
    for fi34 := 0; fi34 < 18; fi34++ {
        for fi35 := 0; fi35 < 9; fi35++ {
            var v36 uint = (^(v30))
            var v37 string = (("hello" + "\n\t") + "a")
            v3[v1] = ((v30 ^ v36) < (v36 % ((v36) | 1)))
            v38 := []int{fi35, (fi34 << ((v36) & 63)), (fi34 ^ -2147483648)}
            _ = len(v5)
            _ = fi35
            _ = v36
            _ = v37
            _ = v38
        }
        var v39 int8 = 0
        _ = fi34
        _ = v39
    }
    if false {
        if true {
            switch (-(v1)) {
            case 9223372036854775807:
                if (((1e10 * 1e10) * 1e10) > ((float64(v31) / (0.1 * 0.1)) - 0.1)) {
                    for fi40 := 0; fi40 < 2; fi40++ {
                        v31 -= (float32(1e10) - v31)
                        _ = fi40
                    }
                    var v41 float32 = ((v31 - v31) * float32(1e10))
                    _ = v41
                }
            case 4294967295:
                v42 := GenS2{F0: true}
                _ = v42
            }
        } else {
            v43 := GenS4{F0: 1}
            v30 = 4294967295
            _ = v43
        }
        v6 = v6
    } else {
        fmt.Println(v30)
        var v44 uint = (v2 | v30)
        var v45 uint16 = 255
        _ = v44
        _ = v45
    }
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
