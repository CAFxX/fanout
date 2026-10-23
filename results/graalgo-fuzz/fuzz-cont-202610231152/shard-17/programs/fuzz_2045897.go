package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8; F1 int64; F2 int8 }
type GenS4 struct { F0 int64; F1 uint16; F2 string }
func genFn0(p0 uint, p1 bool) uint {
    if (-128 <= 0) {
        var v6 uint16 = 0
        v7 := GenS1{F0: 1}
        _ = v6
        _ = v7
    }
    var v8 float32 = ((1.5 + (0.1 - 1.5)) - 3.14159)
    var v9 float64 = float64(((((1e10 * ((v8 + float32(-2.25)) * ((v8 / 3.14159) / v8))) / v8) * 1.5) + (((v8 - v8) / 3.14159) * ((((v8 * (v8 - float32((1.5 - -2.25)))) - float32(0.0)) * -2.25) + v8))))
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 793
}
func genFn1(p0 int32) int {
    if true {
        v10 := []int{993, 128, -723, -165}
        _ = v10
    }
    switch "hello" {
    case "":
        if (true || false) {
            var v11 uint16 = 2
            fmt.Println(v11, p0)
            _ = v11
        } else {
            switch (("" + "\n\t") + ("hello" + "foo bar")) {
            case "":
                for fi12 := 0; fi12 < 4; fi12++ {
                    v13 := []int{fi12, -9223372036854775808}
                    fmt.Println(fi12, fi12)
                    var v14 uint64 = 4294967295
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            }
            for fi15 := 0; fi15 < 5; fi15++ {
                v16 := GenS2{F0: 2}
                switch fi15 {
                case 32768:
                    var v17 uint = genFn0(278, false)
                    _ = v17
                case 9223372036854775807:
                    p0 -= p0
                default:
                    v18 := []int{42, fi15, 668}
                    _ = v18
                }
                _ = fi15
                _ = v16
            }
        }
    }
    fmt.Println(p0)
    _ = p0
    return 2147483647
}

func main() {
    var v1 int = -621
    var v2 bool = (127 > (v1 ^ (v1 + (v1 + (v1 / ((1) | 1))))))
    v3 := make(map[int]bool)
    var v4 uint = 1
    var v5 uint = (v4 / ((v4) | 1))
    var v19 int32 = 42
    v5 += v4
    if (0.0 >= 1e10) {
        var v20 uint16 = 781
        for fi21 := 0; fi21 < 8; fi21++ {
            v4 += v5
            if v2 {
                for fi22 := 0; fi22 < 17; fi22++ {
                    var v23 float64 = (3.14159 / 1e10)
                    v24 := GenS2{F0: 0}
                    var v25 int = genFn1(v19)
                    switch ("\n\t" + ("\n\t" + "a")) {
                    case "\n\t":
                        var v26 string = "foo bar"
                        _ = v26
                    }
                    if (false || (18446744073709551615 < 734)) {
                        v4 = v5
                        var v27 float32 = (float32(v23) * -2.25)
                        var v28 uint = 556
                        v5 -= v28
                        v29 := GenS3{F0: 128, F1: 0, F2: -126}
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                    _ = fi22
                    _ = v23
                    _ = v24
                    _ = v25
                }
                v30 := []int{(fi21 % (((fi21 << ((121) & 63))) | 1)), 4294967295}
                _ = v30
            } else {
                if (1 == -1) {
                    var v31 int = -2147483648
                    v4 = 4294967295
                    _ = v31
                }
                v4 -= (^(v5))
            }
            _ = fi21
        }
        _ = v20
    } else {
        fmt.Println(v2, v1)
        v5 = v4
        var v32 int32 = (v19 << (((v5 >> ((v4) & 63))) & 63))
        _ = v32
    }
    if false {
        v4 = v4
        v4 = v4
    } else {
        v33 := GenS4{F0: -9223372036854775808, F1: 32767, F2: "\n\t"}
        var v34 uint64 = 18446744073709551615
        _ = v33
        _ = v34
    }
    v35 := GenS4{F0: 42, F1: 0, F2: "\n\t"}
    var v36 float32 = 1.5
    v1 -= v1
    var v37 int = (v1 << (((v4 / ((genFn0(genFn0(genFn0(genFn0(genFn0(419, v2), v2), v2), v2), v2)) | 1))) & 63))
    switch "\n\t" {
    case "":
        v4 += v5
    }
    switch "a" {
    case "":
        switch "hello" {
        case "foo bar":
            fmt.Println(v4)
        }
    case "a":
        v4 -= (v5 + (v5 * (v5 % (((v5 / ((v5) | 1))) | 1))))
    case "hello":
        var v38 uint16 = 815
        _ = v38
    default:
        var v39 int32 = v19
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
