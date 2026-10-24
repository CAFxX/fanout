package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 int64; F2 int64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint8; F1 uint }
type GenS5 struct { F0 int; F1 int8 }
func genFn0(p0 string) string {
    var v4 uint32 = 42
    var v5 bool = true
    if (118 >= 255) {
        if v5 {
            if (548 != 0) {
                v6 := []int{-1}
                _ = len(v6)
                _ = v6
            }
        }
        var v7 int16 = 42
        _ = v7
    } else {
        var v8 uint64 = 18446744073709551615
        var v9 int8 = 127
        v10 := GenS1{F0: ("foo bar" + p0)}
        _ = v10
        _ = v8
        _ = v9
    }
    _ = p0
    _ = v4
    _ = v5
    return ("" + ((p0 + "hello") + "foo bar"))
}
func genFn1() int64 {
    var v11 float64 = -2.25
    v11 -= v11
    v11 += v11
    v12 := GenS1{F0: "foo bar"}
    _ = v11
    _ = v12
    return 150
}

func main() {
    v1 := []int{-9223372036854775808}
    v2 := []string{"\n\t", "foo bar", "foo bar"}
    v3 := []int{2147483647, 1}
    if ((42 == 1) && true) {
        var v13 int16 = -694
        _ = v13
    }
    for fi14 := 0; fi14 < 12; fi14++ {
        switch 701 {
        case -9223372036854775808:
            var v15 bool = true
            _ = v15
        case 0:
            fmt.Println(fi14)
        case 2147483647:
            var v16 float32 = (1.5 + float32((0.0 - -2.25)))
            _ = v16
        }
        v17 := GenS1{F0: "\n\t"}
        v18 := GenS2{F0: true, F1: 855, F2: 1}
        _ = fi14
        _ = v17
        _ = v18
    }
    if ((!(!true)) || (344 <= 2)) {
        for fi19 := 0; fi19 < 14; fi19++ {
            var v20 int16 = -32768
            _ = fi19
            _ = v20
        }
        if (!(-1 < 2147483647)) {
            var v21 string = "hello"
            var v22 int64 = -1
            _ = v21
            _ = v22
        }
    }
    if ((-1 == 42) && true) {
        if len(v1) > 0 {
            v1[0] = 42
        }
    }
    if false {
        var v23 int = 872
        _ = v23
    }
    v24 := GenS3{F0: 0}
    var v25 int16 = 32767
    for fi26 := 0; fi26 < 10; fi26++ {
        if len(v3) > 0 {
            v3[0] = (fi26 / (((fi26 >> ((970) & 63))) | 1))
        }
        _ = fi26
    }
    v25 = (v25 | (v25 << ((368) & 63)))
    switch 2147483647 {
    case -9223372036854775808:
        v25 = (v25 / ((-1) | 1))
    case 1:
        v27 := GenS1{F0: "foo bar"}
        _ = v27
    case 4294967295:
        for fi28 := 0; fi28 < 6; fi28++ {
            for fi29 := 0; fi29 < 20; fi29++ {
                switch -453 {
                case 4:
                    v30 := GenS4{F0: 145, F1: 4294967295}
                    _ = v30
                }
                v31 := GenS1{F0: "a"}
                v32 := GenS5{F0: (fi28 >> ((2) & 63)), F1: 114}
                var v33 int16 = (v25 % (((v25 / ((v25) | 1))) | 1))
                var v34 float32 = 0.1
                _ = fi29
                _ = v31
                _ = v32
                _ = v33
                _ = v34
            }
            _ = len(v1)
            _ = fi28
        }
    }
    _ = len(v2)
    v25 = (-(v25))
    var v35 float32 = ((0.0 * ((0.0 * 0.1) + 0.1)) * ((((0.0 - -2.25) - (-2.25 - (float32(float64(0.0)) / 3.14159))) - 1e10) + 0.0))
    if len(v1) > 0 {
        v1[0] = -898
    }
    var v36 int64 = 0
    for fi37 := 0; fi37 < 4; fi37++ {
        fmt.Println(v36, fi37)
        _ = fi37
    }
    var v38 int64 = (v36 * v36)
    for fi39 := 0; fi39 < 7; fi39++ {
        switch (("\n\t" + "a") + "\n\t") {
        case "hello":
            v35 = (((v35 - 0.0) - v35) * (v35 + 0.1))
        case "a":
            v1 = append(v1, fi39)
        default:
            if (2 >= 42) {
                switch fi39 {
                case -2147483648:
                    if len(v2) > 0 {
                        v2[0] = "hello"
                    }
                case 127:
                    fmt.Println(v38, v35)
                case 2147483647:
                    v1 = append(v1, fi39)
                }
            }
        }
        var v40 int = -9223372036854775808
        if false {
            v40 += (fi39 + (v40 >> ((255) & 63)))
        } else {
            for fi41 := 0; fi41 < 9; fi41++ {
                v35 = v35
                fmt.Println(v40, v40)
                v38 = v38
                _ = fi41
            }
            v40 += 9223372036854775807
            for fi42 := 0; fi42 < 5; fi42++ {
                fmt.Println(v38)
                _ = fi42
            }
        }
        _ = fi39
        _ = v40
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v35
    _ = v36
    _ = v38
}
