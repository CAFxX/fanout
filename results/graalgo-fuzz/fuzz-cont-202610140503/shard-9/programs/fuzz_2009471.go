package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int16; F1 string; F2 int32 }
func genFn0(p0 string, p1 uint64, p2 uint) int32 {
    v6 := []int{128, 926, int(p2), -1}
    fmt.Println(p2, p1, p1)
    fmt.Println(p2)
    v7 := GenS1{F0: -236, F1: 42}
    v8 := GenS1{F0: 1000, F1: -2147483648}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 42
}
func genFn1() uint32 {
    var v9 uint64 = 472
    if (!(32767 <= 1000)) {
        var v10 uint64 = v9
        _ = v10
    } else {
        v11 := []int{127, 32768}
        _ = v11
    }
    v12 := []int{-2147483648}
    var v13 int = 127
    _ = v12
    _ = v13
    _ = v9
    return 2147483648
}

func main() {
    var v1 uint64 = 2
    var v2 uint16 = 766
    var v3 float32 = 1e10
    var v4 int8 = 42
    var v5 bool = true
    v5 = v5
    v3 -= -2.25
    v14 := GenS2{F0: 4294967295}
    v1 += 0
    switch "foo bar" {
    case "foo bar":
        v15 := []int{4294967295, 2147483647, 42}
        _ = v15
    case "\n\t":
        if (-32768 > -363) {
            var v16 int64 = 0
            switch 128 {
            case 762:
                switch "foo bar" {
                case "a":
                    fmt.Println(v2)
                case "\n\t":
                    v17 := []int{}
                    _ = v17
                case "":
                    v2 = (v2 | v2)
                }
            case 32767:
                v18 := GenS3{F0: 1}
                _ = v18
            case 127:
                var v19 int64 = v16
                _ = v19
            default:
                for fi20 := 0; fi20 < 8; fi20++ {
                    for fi21 := 0; fi21 < 17; fi21++ {
                        var v22 float32 = (v3 / ((float32(0.0) / -2.25) * v3))
                        fmt.Println(v5, v1, v3)
                        _ = fi21
                        _ = v22
                    }
                    fmt.Println(v3)
                    if (0 == 0) {
                        if (v2 < (v2 / (((v2 + v2)) | 1))) {
                            if v5 {
                                v23 := GenS4{F0: ("a" + "")}
                                v2 = v2
                                _ = v23
                            }
                            var v24 float64 = 0.1
                            var v25 int32 = genFn0(("a" + "\n\t"), (v1 | 0), 0)
                            _ = v24
                            _ = v25
                        } else {
                            switch "a" {
                            case "":
                                if ("hello" >= "foo bar") {
                                    var v26 uint16 = v2
                                    switch int(v2) {
                                    case -146:
                                        var v27 int16 = 42
                                        _ = v27
                                    }
                                    _ = v26
                                }
                            case "foo bar":
                                v28 := []int{fi20, -1, fi20}
                                _ = v28
                            case "hello":
                                v29 := GenS2{F0: (fi20 + fi20)}
                                _ = v29
                            default:
                                v2 += (^(v2))
                            }
                            v16 = 9223372036854775807
                        }
                    } else {
                        if (65535 != 2) {
                            v4 -= 4
                            v30 := GenS4{F0: "a"}
                            _ = v30
                        }
                    }
                    _ = fi20
                }
            }
            _ = v16
        } else {
            fmt.Println(v3)
            if false {
                var v31 int32 = -2147483648
                _ = v31
            }
            v32 := GenS5{F0: 32767, F1: ("foo bar" + ("" + ("foo bar" + "\n\t"))), F2: -115}
            _ = v32
        }
    }
    v5 = (v3 <= (float32((0.1 * 3.14159)) - (v3 * (v3 * (v3 * (float32(1.5) / 1.0))))))
    for fi33 := 0; fi33 < 17; fi33++ {
        v2 -= v2
        var v34 uint16 = (v2 | 2)
        _ = fi33
        _ = v34
    }
    var v35 int8 = v4
    for fi36 := 0; fi36 < 19; fi36++ {
        var v37 int16 = 1
        _ = fi36
        _ = v37
    }
    v5 = v5
    switch "hello" {
    case "foo bar":
        var v38 float32 = v3
        _ = v38
    }
    if v5 {
        v39 := GenS5{F0: 1000, F1: "hello", F2: 42}
        v3 -= (3.14159 / 1.0)
        v40 := GenS4{F0: ("\n\t" + "")}
        _ = v39
        _ = v40
    }
    if (!(!v5)) {
        var v41 int32 = 811
        var v42 uint32 = genFn1()
        v41 += int32(v2)
        _ = v41
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
