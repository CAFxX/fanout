package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 string; F1 int64 }
func genFn0(p0 uint64, p1 uint64) bool {
    var v6 uint16 = 0
    v7 := []int{}
    var v8 string = "a"
    switch v8 {
    case "hello":
        fmt.Println(v6, p1, v6)
    case "":
        switch "foo bar" {
        case "hello":
            var v9 uint = 2
            _ = v9
        case "\n\t":
            p0 += (p1 + p1)
        default:
            var v10 int16 = 32767
            _ = v10
        }
    default:
        if len(v7) > 0 {
            v7[0] = 1
        }
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return ((1.5 < -2.25) && (-281 != 1))
}
func genFn1(p0 int16) uint32 {
    if true {
        var v11 int32 = 2147483647
        _ = v11
    }
    v12 := GenS1{F0: 0}
    var v13 float64 = 3.14159
    _ = p0
    _ = v12
    _ = v13
    return 557
}

func main() {
    var v1 float64 = (1.5 * 1.5)
    v2 := []int{0}
    v3 := []string{"a", ("foo bar" + ((("\n\t" + "a") + ((("\n\t" + (("\n\t" + ("\n\t" + ("a" + "foo bar"))) + "a")) + "\n\t") + "")) + "a"))}
    var v4 int8 = -1
    v5 := make(map[int]string)
    switch "a" {
    case "":
        v1 += (float64(-2.25) / 3.14159)
    case "\n\t":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "a":
        v2 = append(v2, 128)
    }
    v4 = v4
    switch -1 {
    case -81:
        fmt.Println(v1)
    case 9223372036854775807:
        v4 += v4
    default:
        switch "\n\t" {
        case "foo bar":
            for fi14 := 0; fi14 < 5; fi14++ {
                var v15 int16 = -32768
                v3 = append(v3, "foo bar")
                _ = len(v2)
                _ = fi14
                _ = v15
            }
        case "":
            for fi16 := 0; fi16 < 1; fi16++ {
                var v17 float32 = 0.0
                v17 += v17
                switch (fi16 * 127) {
                case 0:
                    var v18 int8 = v4
                    _ = v18
                case 437:
                    fmt.Println(fi16)
                case -1:
                    fmt.Println(v1, v17, fi16)
                }
                _ = fi16
                _ = v17
            }
        case "\n\t":
            switch 689 {
            case -848:
                v1 -= 0.1
            case 713:
                for fi19 := 0; fi19 < 17; fi19++ {
                    fmt.Println(v4, v4, v4)
                    if false {
                        var v20 int8 = (v4 * v4)
                        switch fi19 {
                        case 892:
                            v20 = v20
                        }
                        var v21 uint = 42
                        _ = v20
                        _ = v21
                    } else {
                        delete(v5, (fi19 << ((839) & 63)))
                        switch (fi19 >> ((4294967295) & 63)) {
                        case 4294967295:
                            v1 = 1.5
                        default:
                            switch ("\n\t" + "hello") {
                            case "hello":
                                var v22 int64 = 397
                                _ = v22
                            case "foo bar":
                                var v23 float32 = 1.5
                                _ = v23
                            }
                        }
                    }
                    if (42 < 0) {
                        v3 = append(v3, ("foo bar" + ("" + "")))
                        for fi24 := 0; fi24 < 11; fi24++ {
                            var v25 int8 = (v4 - -11)
                            v4 = v25
                            v26 := GenS1{F0: 539}
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                        v1 = (v1 - (float64(-2.25) / (((-2.25 * v1) + float64(-2.25)) * (((1e10 / v1) - 0.0) / v1))))
                    }
                    _ = fi19
                }
            default:
                v1 = (v1 + (v1 / v1))
            }
        default:
            v1 -= float64(1.5)
        }
    }
    for fi27 := 0; fi27 < 20; fi27++ {
        if (1 <= 2) {
            v1 = v1
        } else {
            switch "\n\t" {
            case "hello":
                v1 = 1.5
            }
            var v28 float64 = v1
            _ = v28
        }
        for fi29 := 0; fi29 < 10; fi29++ {
            if len(v2) > 0 {
                v2[0] = (-(fi27))
            }
            v1 = 1.5
            var v30 int8 = (v4 - 1)
            _ = fi29
            _ = v30
        }
        var v31 uint8 = 62
        _ = fi27
        _ = v31
    }
    switch -203 {
    case -263:
        v32 := GenS2{F0: 1, F1: 2}
        _ = v32
    case 0:
        v1 -= v1
    }
    for fi33 := 0; fi33 < 2; fi33++ {
        v34 := GenS2{F0: 1, F1: 1}
        var v35 float32 = 3.14159
        var v36 uint = 255
        _ = fi33
        _ = v34
        _ = v35
        _ = v36
    }
    fmt.Println(v1, v1)
    v37 := GenS3{F0: ("\n\t" + ""), F1: 657}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
}
