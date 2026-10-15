package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32; F1 string; F2 string }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint8; F1 uint; F2 int }
func genFn0(p0 uint8, p1 int8) int {
    switch "\n\t" {
    case "\n\t":
        var v6 int64 = 429
        _ = v6
    case "foo bar":
        for fi7 := 0; fi7 < 1; fi7++ {
            v8 := GenS1{F0: 178, F1: (p0 >> ((65535) & 63))}
            switch fi7 {
            case 128:
                for fi9 := 0; fi9 < 4; fi9++ {
                    v10 := GenS1{F0: 4294967295, F1: p0}
                    _ = fi9
                    _ = v10
                }
            case -259:
                for fi11 := 0; fi11 < 18; fi11++ {
                    fmt.Println(p0, p0, fi11)
                    p0 = p0
                    _ = fi11
                }
            default:
                v12 := []int{fi7, fi7}
                _ = v12
            }
            if true {
                if true {
                    v13 := []int{(fi7 % (((fi7 << ((1) & 63))) | 1)), (fi7 - 377), (fi7 & int(p0))}
                    _ = v13
                } else {
                    v14 := GenS2{F0: p1}
                    v15 := GenS3{F0: "hello"}
                    _ = v14
                    _ = v15
                }
                var v16 int16 = -969
                _ = v16
            }
            _ = fi7
            _ = v8
        }
    case "a":
        p0 += (p0 & (p0 % ((128) | 1)))
    default:
        switch (("\n\t" + "") + "hello") {
        case "foo bar":
            var v17 int8 = -69
            _ = v17
        case "a":
            p0 = p0
        default:
            p0 = (^(p0))
        }
    }
    if true {
        if (1 != 1) {
            for fi18 := 0; fi18 < 14; fi18++ {
                v19 := GenS4{F0: -270, F1: "hello", F2: "\n\t"}
                _ = fi18
                _ = v19
            }
            v20 := []int{319, 127, -687, 645}
            _ = v20
        }
    } else {
        if true {
            v21 := []int{4294967295, 2147483647}
            _ = v21
        }
        p0 = 255
    }
    _ = p0
    _ = p1
    return -9223372036854775808
}

func main() {
    v1 := make(map[int]bool)
    v2 := []string{}
    v3 := make(map[int]string)
    var v4 float64 = (float64(1e10) / ((3.14159 / -2.25) / 3.14159))
    var v5 int64 = -9223372036854775808
    fmt.Println(v5)
    v22 := GenS1{F0: 4294967295, F1: 144}
    for fi23 := 0; fi23 < 2; fi23++ {
        v24 := GenS5{F0: 127}
        fmt.Println(v5, fi23, v5)
        switch (^(fi23)) {
        case 158:
            for fi25 := 0; fi25 < 7; fi25++ {
                v4 = (v4 + ((v4 / v4) + (v4 * (1e10 * ((v4 - 0.0) / ((0.1 * 3.14159) / float64((((1.5 * 1.5) + 1.5) - -2.25))))))))
                _ = fi25
            }
        }
        _ = fi23
        _ = v24
    }
    if (-143 < genFn0(127, 1)) {
        if true {
            fmt.Println(v4)
            if (true && (-161 < -1)) {
                delete(v1, -962)
                switch "foo bar" {
                case "foo bar":
                    var v26 string = "hello"
                    _ = v26
                case "":
                    v4 = (((((0.1 * v4) * (float64(3.14159) * v4)) / (0.1 - float64((1e10 / 1.0)))) / (float64((((float32(-2.25) / -2.25) * 3.14159) / (((3.14159 / float32(((-2.25 / v4) + (v4 / 1e10)))) - (0.1 / -2.25)) * 3.14159))) + (1.5 - (-2.25 + (v4 - float64((float32(1e10) + ((-2.25 - 3.14159) + float32(v4))))))))) + 1e10)
                }
                v1[-2147483648] = ((true || (128 > 9223372036854775807)) || false)
            }
            v5 = v5
        }
        var v27 float32 = ((1e10 * ((1.5 + 1.5) - 1.5)) / 1e10)
        _ = v27
    } else {
        v1[127] = true
        switch 1 {
        case 42:
            v4 = v4
        }
    }
    if false {
        v4 = v4
    }
    v28 := GenS3{F0: ""}
    switch "\n\t" {
    case "":
        v29 := GenS4{F0: 0, F1: (("a" + "foo bar") + "foo bar"), F2: "a"}
        _ = v29
    }
    v30 := GenS6{F0: 0, F1: 255, F2: -9223372036854775808}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v28
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
