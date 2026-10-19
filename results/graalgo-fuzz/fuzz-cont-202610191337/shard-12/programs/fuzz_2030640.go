package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 int32 }
type GenS2 struct { F0 uint32; F1 uint64; F2 uint64 }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 uint; F1 int64; F2 int16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint64; F1 string }
func genFn0() uint32 {
    v4 := GenS1{F0: -1, F1: -2147483648, F2: 42}
    var v5 float32 = (1.5 + float32((0.1 * ((0.1 * 1.5) * ((((((3.14159 - 1.5) * 1e10) * 1.5) - 1e10) + (float64(-2.25) - 0.0)) / float64(0.1))))))
    v5 = v5
    if ("a" < "") {
        switch "foo bar" {
        case "foo bar":
            v6 := []int{0, 31, 9223372036854775807, 0}
            _ = v6
        case "hello":
            if true {
                fmt.Println(v5)
                v5 += v5
            } else {
                v7 := GenS2{F0: 2, F1: 1, F2: 1}
                _ = v7
            }
        }
        v8 := []int{-769}
        v5 = ((float32((1e10 / 3.14159)) * 1e10) - v5)
        _ = v8
    } else {
        if (!false) {
            v9 := GenS3{F0: -9223372036854775808, F1: 0}
            for fi10 := 0; fi10 < 8; fi10++ {
                var v11 int64 = 0
                _ = fi10
                _ = v11
            }
            v12 := []int{-137, 196, 1}
            _ = v12
            _ = v9
        }
        switch -9223372036854775808 {
        case -9223372036854775808:
            v13 := []int{981, 128, -2147483648}
            _ = v13
        case -836:
            v14 := []int{-948}
            _ = v14
        }
    }
    _ = v4
    _ = v5
    return 2147483648
}
func genFn1() int8 {
    if (874 == 65535) {
        var v15 int32 = 985
        _ = v15
    } else {
        var v16 uint32 = 2147483648
        if true {
            switch "a" {
            case "\n\t":
                for fi17 := 0; fi17 < 2; fi17++ {
                    var v18 int64 = 0
                    _ = fi17
                    _ = v18
                }
            }
            v19 := []int{-487}
            _ = v19
        } else {
            switch "\n\t" {
            case "foo bar":
                fmt.Println(v16)
            case "a":
                var v20 bool = (-2147483648 <= 146)
                _ = v20
            case "hello":
                for fi21 := 0; fi21 < 7; fi21++ {
                    if (("a" == "\n\t") && (fi21 == fi21)) {
                        v22 := GenS4{F0: 4294967295, F1: int64(v16), F2: -8}
                        _ = v22
                    }
                    if (-2147483648 <= 0) {
                        for fi23 := 0; fi23 < 19; fi23++ {
                            fmt.Println(fi21)
                            _ = fi23
                        }
                        v16 += genFn0()
                        v16 += genFn0()
                    } else {
                        v16 = v16
                        v16 = (v16 * v16)
                    }
                    switch (fi21 / ((fi21) | 1)) {
                    case -9223372036854775808:
                        v24 := []int{(fi21 | (fi21 >> ((1) & 63)))}
                        _ = v24
                    case 9223372036854775807:
                        v16 = (v16 + 2)
                    case 32767:
                        v25 := []int{fi21, fi21, fi21}
                        _ = v25
                    }
                    _ = fi21
                }
            default:
                v26 := GenS5{F0: (1 >= 1)}
                _ = v26
            }
            v16 = v16
        }
        _ = v16
    }
    var v27 bool = (932 <= 42)
    v28 := GenS1{F0: 679, F1: -2147483648, F2: -97}
    v29 := GenS6{F0: 1, F1: ("" + "a")}
    if (127 < -128) {
        for fi30 := 0; fi30 < 20; fi30++ {
            var v31 string = ("\n\t" + ("foo bar" + "hello"))
            switch ((v31 + v31) + v31) {
            case "":
                v31 = "a"
            }
            var v32 uint16 = 255
            _ = fi30
            _ = v31
            _ = v32
        }
    }
    _ = v27
    _ = v28
    _ = v29
    return -128
}

func main() {
    var v1 string = "\n\t"
    var v2 uint32 = 2147483647
    v3 := make(map[string]int)
    var v33 uint64 = 499
    fmt.Println(v1, v2, v1)
    var v34 float32 = 3.14159
    v1 = v1
    v35 := []int{9223372036854775807, 1}
    var v36 uint16 = 1
    var v37 float64 = 1e10
    v37 -= (float64(1.5) + (((((v37 + 3.14159) * v37) - v37) / float64(float32(v37))) / (((v37 / ((float64(-2.25) + (v37 + float64(v34))) / v37)) * (v37 * 1.5)) - 1.5)))
    v36 += v36
    var v38 int16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
}
