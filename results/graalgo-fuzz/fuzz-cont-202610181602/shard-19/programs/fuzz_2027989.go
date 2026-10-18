package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
func genFn0(p0 bool) int32 {
    for fi7 := 0; fi7 < 6; fi7++ {
        p0 = false
        _ = fi7
    }
    for fi8 := 0; fi8 < 11; fi8++ {
        v9 := []int{(^(fi8)), (fi8 << ((2) & 63)), 4294967295}
        p0 = (fi8 >= (fi8 / ((42) | 1)))
        _ = fi8
        _ = v9
    }
    p0 = (42 != 569)
    var v10 int32 = 1
    _ = p0
    _ = v10
    return v10
}

func main() {
    var v1 uint16 = 32767
    v2 := make(map[string]string)
    v3 := make(map[string]string)
    var v4 float64 = 0.1
    v5 := []uint32{42}
    var v6 int = 127
    if (!(128 > v6)) {
        if ((1e10 - -2.25) > 3.14159) {
            if (!true) {
                var v11 uint8 = 35
                v1 = 1
                _ = v11
            } else {
                delete(v2, "\n\t")
                v6 -= v6
                var v12 string = (("" + "") + "foo bar")
                _ = v12
            }
        }
        v13 := GenS1{F0: "foo bar", F1: 564}
        _ = v13
    }
    fmt.Println(v4, v1)
    v14 := GenS1{F0: "\n\t", F1: 2}
    if true {
        v6 -= (v6 ^ v6)
        delete(v2, ((("a" + ("hello" + "")) + ("" + "\n\t")) + "\n\t"))
        if true {
            v4 = (float64(3.14159) - 1e10)
            v4 = v4
        } else {
            v15 := GenS1{F0: "\n\t", F1: 2147483647}
            fmt.Println(v6, v6)
            _ = v15
        }
    }
    switch v6 {
    case -9223372036854775808:
        switch v6 {
        case -9223372036854775808:
            _ = len(v3)
        case 127:
            v1 = 0
        default:
            switch (("" + "") + ("foo bar" + "a")) {
            case "\n\t":
                var v16 int32 = -415
                _ = v16
            }
        }
    }
    if ((0.1 / ((3.14159 - (0.0 + float32(1e10))) * (-2.25 / 1e10))) < ((1e10 * ((3.14159 / 0.1) * ((((-2.25 * -2.25) - 1e10) + ((3.14159 / ((float32((((0.0 - (v4 + ((v4 + ((v4 * ((float64(0.0) * 1.5) - float64((float32(-2.25) / 3.14159)))) / v4)) / v4))) - float64((0.1 / -2.25))) / v4)) * 3.14159) + ((3.14159 + 0.0) + 0.1))) - ((((0.1 * 0.0) / ((3.14159 - 1e10) / 1e10)) - ((3.14159 / 3.14159) * float32(v4))) / (3.14159 * float32(v4))))) * (float32((v4 / v4)) / 0.1)))) - (-2.25 * -2.25))) {
        for fi17 := 0; fi17 < 2; fi17++ {
            v4 += 1.5
            for fi18 := 0; fi18 < 9; fi18++ {
                var v19 string = ("foo bar" + "")
                _ = fi18
                _ = v19
            }
            _ = fi17
        }
        v6 += 0
        delete(v3, "")
    }
    if ((!true) && true) {
        fmt.Println(v6)
    }
    var v20 int16 = 32767
    fmt.Println(v1)
    var v21 int32 = genFn0(true)
    for fi22 := 0; fi22 < 19; fi22++ {
        var v23 uint8 = 1
        fmt.Println(v4, v23, v4)
        _ = fi22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
