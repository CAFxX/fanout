package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint64; F1 int16; F2 uint8 }
func genFn0() uint8 {
    switch -823 {
    case -2147483648:
        var v6 int32 = 2147483647
        _ = v6
    case -9223372036854775808:
        v7 := []int{-200, 1, 484}
        _ = v7
    case 345:
        v8 := []int{-753, 42, 127, 128}
        _ = v8
    default:
        for fi9 := 0; fi9 < 6; fi9++ {
            var v10 uint32 = 0
            if false {
                var v11 string = "hello"
                _ = v11
            } else {
                if (0 > 18446744073709551615) {
                    v12 := []int{0}
                    _ = v12
                } else {
                    var v13 int64 = 1
                    v13 = (v13 | 38)
                    var v14 uint32 = 65535
                    _ = v13
                    _ = v14
                }
            }
            _ = fi9
            _ = v10
        }
    }
    v15 := []int{32767, 128, 42, -397}
    v16 := GenS1{F0: 0}
    _ = v15
    _ = v16
    return 202
}

func main() {
    v1 := []string{("\n\t" + ""), ((("foo bar" + "a") + "foo bar") + ""), ("a" + "hello")}
    v2 := make(map[int]string)
    var v3 uint = 0
    var v4 uint16 = 1
    var v5 int8 = 100
    if (v5 < (v5 >> ((v3) & 63))) {
        v17 := GenS1{F0: 42}
        v18 := GenS1{F0: 32768}
        _ = v17
        _ = v18
    } else {
        if false {
            v19 := GenS1{F0: -737}
            switch ("hello" + "foo bar") {
            case "\n\t":
                v20 := GenS2{F0: -518}
                _ = v20
            default:
                if ((!false) && (v5 == 42)) {
                    v1 = append(v1, ("foo bar" + "hello"))
                    var v21 uint16 = v4
                    _ = v21
                }
            }
            for fi22 := 0; fi22 < 16; fi22++ {
                if ("foo bar" <= "hello") {
                    var v23 uint16 = (v4 * v4)
                    for fi24 := 0; fi24 < 12; fi24++ {
                        var v25 int8 = v5
                        _ = fi24
                        _ = v25
                    }
                    _ = v23
                }
                for fi26 := 0; fi26 < 4; fi26++ {
                    if true {
                        v27 := GenS2{F0: -1}
                        if false {
                            var v28 uint64 = 2
                            _ = v28
                        }
                        if ((v5 + -128) == 127) {
                            fmt.Println(fi22, fi26, v5)
                        } else {
                            v3 = (v3 >> (((v3 + v3)) & 63))
                        }
                        _ = v27
                    }
                    _ = len(v1)
                    _ = fi26
                }
                var v29 uint64 = 42
                _ = fi22
                _ = v29
            }
            _ = v19
        }
    }
    var v30 float64 = (float64(1.5) * 1.5)
    switch ("" + ("a" + "foo bar")) {
    case "":
        v2[609] = ("hello" + "foo bar")
    case "foo bar":
        switch -439 {
        case -261:
            var v31 float32 = 0.1
            _ = v31
        }
    case "\n\t":
        var v32 uint16 = v4
        _ = v32
    default:
        if false {
            v30 = ((float64(3.14159) * 0.1) / v30)
            switch ((("hello" + ("foo bar" + "a")) + "\n\t") + "\n\t") {
            case "hello":
                if ((^(v3)) <= v3) {
                    v33 := GenS1{F0: 4294967295}
                    _ = v33
                } else {
                    switch -2147483648 {
                    case 0:
                        v5 -= (-(v5))
                    case -2147483648:
                        v30 = 1e10
                    case -9223372036854775808:
                        v34 := GenS3{F0: 18446744073709551615, F1: 428, F2: 156}
                        _ = v34
                    default:
                        var v35 uint16 = 42
                        _ = v35
                    }
                    v4 = v4
                }
            case "a":
                v1 = append(v1, "foo bar")
            case "":
                v4 -= v4
            }
        }
    }
    v4 -= 64
    var v36 int64 = 9223372036854775807
    var v37 int16 = -1
    v38 := GenS1{F0: -148}
    v3 = v3
    fmt.Println(v36, v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
