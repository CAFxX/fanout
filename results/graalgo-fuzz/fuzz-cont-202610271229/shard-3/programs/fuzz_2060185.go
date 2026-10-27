package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 int; F1 uint; F2 bool }
func genFn0(p0 uint8, p1 int64, p2 int32) int {
    v5 := GenS1{F0: 314, F1: 1}
    v6 := []int{}
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 int16 = 0
        var v9 int64 = p1
        p1 = 42
        _ = fi7
        _ = v8
        _ = v9
    }
    if len(v6) > 0 {
        v6[0] = 9223372036854775807
    }
    var v10 uint64 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v5
    _ = v6
    return 2147483647
}
func genFn1(p0 uint16, p1 uint8) string {
    v11 := []int{}
    switch "hello" {
    case "\n\t":
        if (42 < -31) {
            fmt.Println(p0, p0)
        } else {
            _ = len(v11)
            v12 := GenS1{F0: 32767, F1: p0}
            switch 4294967295 {
            case 1:
                var v13 float32 = float32(0.0)
                _ = v13
            case -2147483648:
                var v14 uint64 = 705
                _ = v14
            }
            _ = v12
        }
    case "foo bar":
        for fi15 := 0; fi15 < 19; fi15++ {
            if len(v11) > 0 {
                v11[0] = 210
            }
            _ = fi15
        }
    case "hello":
        for fi16 := 0; fi16 < 14; fi16++ {
            p1 = p1
            v11 = append(v11, fi16)
            _ = fi16
        }
    default:
        if (false && (!(1 >= 0))) {
            p0 = (p0 + p0)
        } else {
            p0 = (p0 + (p0 * (p0 / (((^(p0))) | 1))))
            var v17 uint = 1
            var v18 int64 = 813
            _ = v17
            _ = v18
        }
    }
    if true {
        if len(v11) > 0 {
            v11[0] = 781
        }
        _ = len(v11)
    }
    for fi19 := 0; fi19 < 11; fi19++ {
        p1 = (p1 / ((p1) | 1))
        var v20 uint = 133
        v20 = 2
        _ = fi19
        _ = v20
    }
    v21 := GenS1{F0: genFn0(p1, 9223372036854775807, 42), F1: p0}
    _ = p0
    _ = p1
    _ = v11
    _ = v21
    return "a"
}

func main() {
    var v1 int16 = 32767
    v2 := make(map[string]bool)
    var v3 int = 42
    v4 := make(map[int]string)
    switch (^(v3)) {
    case 330:
        delete(v4, (v3 >> ((0) & 63)))
    case 2147483647:
        for fi22 := 0; fi22 < 20; fi22++ {
            v23 := GenS1{F0: v3, F1: 65535}
            var v24 uint8 = 128
            _ = fi22
            _ = v23
            _ = v24
        }
    case 4294967295:
        v25 := GenS1{F0: (v3 | v3), F1: 1}
        _ = v25
    default:
        v26 := v4[(v3 * (v3 << ((42) & 63)))]
        _ = v26
    }
    fmt.Println(v3, v1)
    var v27 int = (v3 | -233)
    if false {
        switch genFn0(127, 42, -2147483648) {
        case -27:
            var v28 string = ""
            _ = v28
        }
        for fi29 := 0; fi29 < 12; fi29++ {
            v30 := []int{}
            for fi31 := 0; fi31 < 16; fi31++ {
                switch ("" + ("" + "hello")) {
                case "foo bar":
                    fmt.Println(fi29)
                case "a":
                    if ((1 > 42) || true) {
                        v3 += 0
                        var v32 uint16 = 330
                        fmt.Println(fi29, fi31, v32)
                        _ = v32
                    }
                }
                if len(v30) >= 2 {
                    v30 = v30[1:len(v30)]
                }
                _ = fi31
            }
            _ = fi29
            _ = v30
        }
        fmt.Println(v1, v3)
    } else {
        for fi33 := 0; fi33 < 14; fi33++ {
            var v34 int8 = 0
            _ = fi33
            _ = v34
        }
        v3 += genFn0(128, -1, 2147483647)
        var v35 int = v3
        _ = v35
    }
    v36 := GenS2{F0: 127, F1: 0, F2: (1e10 == ((float32(((0.1 / 1.5) * float64(float32((1e10 / (1e10 - 1.5)))))) / ((float32((0.1 + (0.1 / -2.25))) + 0.1) - 3.14159)) - ((1e10 + 1e10) - -2.25)))}
    fmt.Println(v3, v27, v27)
    v1 -= v1
    if false {
        if (710 == 0) {
            var v37 int8 = 127
            _ = v37
        }
    }
    delete(v2, "a")
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v36
    _ = v4
}
