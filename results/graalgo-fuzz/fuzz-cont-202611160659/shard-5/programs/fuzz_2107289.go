package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 bool }

func main() {
    var v1 uint = 609
    var v2 uint16 = 32767
    var v3 uint16 = (v2 / (((^(v2))) | 1))
    var v4 int16 = -1
    var v5 uint8 = 127
    var v6 int32 = -2147483648
    fmt.Println(v1, v3)
    fmt.Println(v2)
    v7 := GenS1{F0: (v1 + v1), F1: v6}
    v1 -= (v1 * (^(v1)))
    switch 32767 {
    case 2147483647:
        for fi8 := 0; fi8 < 1; fi8++ {
            if (!false) {
                for fi9 := 0; fi9 < 9; fi9++ {
                    fmt.Println(v1, fi8)
                    if true {
                        fmt.Println(v3, v3, v3)
                    }
                    _ = fi9
                }
                v10 := []int{(^(fi8)), fi8, fi8, (fi8 / (((fi8 + 42)) | 1))}
                v11 := GenS2{F0: 1}
                _ = v10
                _ = v11
            } else {
                fmt.Println(fi8, v6)
                var v12 int32 = (v6 | v6)
                _ = v12
            }
            v13 := []int{(fi8 * (fi8 ^ (fi8 * fi8)))}
            if true {
                v4 += 629
            }
            _ = fi8
            _ = v13
        }
    case 1:
        if ((!false) || false) {
            var v14 uint16 = 20
            fmt.Println(v1, v3)
            var v15 int64 = int64(v5)
            _ = v14
            _ = v15
        }
    case -1:
        if (1 != 1) {
            var v16 string = "hello"
            v4 -= (-(v4))
            _ = v16
        }
    default:
        var v17 int32 = (-(v6))
        _ = v17
    }
    v6 -= v6
    if (!((!(-533 == int(v6))) || true)) {
        v4 = (v4 | v4)
        v18 := []int{127, 179, -845, 42}
        _ = v18
    } else {
        fmt.Println(v6)
    }
    for fi19 := 0; fi19 < 16; fi19++ {
        v20 := []int{32768, fi19, (fi19 * fi19)}
        var v21 int64 = -9223372036854775808
        switch fi19 {
        case -914:
            for fi22 := 0; fi22 < 7; fi22++ {
                v23 := GenS1{F0: 255, F1: int32(v5)}
                var v24 uint64 = 4294967295
                _ = fi22
                _ = v23
                _ = v24
            }
        default:
            v25 := GenS2{F0: 0}
            _ = v25
        }
        _ = fi19
        _ = v20
        _ = v21
    }
    v26 := []int{}
    v26 = append(v26, 261)
    var v27 float64 = ((((1.5 / -2.25) - (((0.1 / 1.0) / (((float64(0.1) * 3.14159) / (0.0 + 0.1)) + (0.1 - ((-2.25 + float64(float32(1e10))) / (1e10 / 3.14159))))) / -2.25)) * 1e10) * 0.1)
    var v28 uint = v1
    var v29 uint = (v28 / ((42) | 1))
    fmt.Println(v6, v2, v6)
    if true {
        v30 := GenS3{F0: (40 < 127)}
        for fi31 := 0; fi31 < 18; fi31++ {
            fmt.Println(v29, v1, v27)
            if ((v3 % ((v2) | 1)) <= 65535) {
                var v32 uint8 = (^(v5))
                _ = v32
            }
            for fi33 := 0; fi33 < 2; fi33++ {
                var v34 uint16 = v2
                var v35 int64 = -145
                _ = fi33
                _ = v34
                _ = v35
            }
            _ = fi31
        }
        switch (("a" + (((("a" + ("a" + (("foo bar" + ("foo bar" + ("hello" + (("\n\t" + "a") + ("hello" + "\n\t"))))) + "foo bar"))) + "") + "a") + ("a" + "a"))) + ("" + "\n\t")) {
        case "hello":
            v36 := GenS2{F0: 1}
            _ = v36
        case "":
            v26 = append(v26, 2147483647)
        case "foo bar":
            var v37 string = (("hello" + ("a" + ("foo bar" + "\n\t"))) + ("a" + ""))
            _ = v37
        default:
            var v38 int64 = 0
            _ = v38
        }
        _ = v30
    } else {
        v29 = (^(v1))
    }
    var v39 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
