package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 int; F2 uint }
type GenS3 struct { F0 int8; F1 uint64 }
type GenS4 struct { F0 int }
func genFn0(p0 uint8, p1 int64, p2 bool) int32 {
    v5 := []int{}
    p0 += p0
    if p2 {
        var v6 bool = (127 < 0)
        if (true && v6) {
            v7 := GenS1{F0: 1}
            var v8 uint = 0
            fmt.Println(p1, p1, p0)
            _ = v7
            _ = v8
        } else {
            for fi9 := 0; fi9 < 8; fi9++ {
                var v10 float64 = -2.25
                v10 -= v10
                v10 = (1e10 - (v10 * (v10 / (v10 + -2.25))))
                _ = fi9
                _ = v10
            }
            if false {
                if len(v5) > 0 {
                    v5[0] = -214
                }
            }
            fmt.Println(p1, p1, p2)
        }
        if ((-2147483648 < -2147483648) && (42 >= 42)) {
            v6 = (65535 > 2)
            v11 := GenS2{F0: 0, F1: 874, F2: 2}
            var v12 int16 = 970
            _ = v11
            _ = v12
        }
        _ = v6
    } else {
        v13 := GenS2{F0: 918, F1: 531, F2: 255}
        fmt.Println(p2, p1)
        _ = v13
    }
    for fi14 := 0; fi14 < 17; fi14++ {
        var v15 int8 = -1
        fmt.Println(fi14, p0)
        _ = fi14
        _ = v15
    }
    switch ("hello" + "foo bar") {
    case "foo bar":
        var v16 uint32 = 2147483647
        _ = v16
    case "hello":
        fmt.Println(p0, p0)
    case "a":
        switch -496 {
        case 32767:
            switch 32767 {
            case 483:
                if (!p2) {
                    var v17 bool = true
                    v18 := GenS1{F0: p1}
                    _ = v17
                    _ = v18
                } else {
                    p0 = p0
                }
            case -2147483648:
                fmt.Println(p1, p2, p1)
            case 0:
                v19 := GenS1{F0: (^(p1))}
                _ = v19
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 1
}
func genFn1(p0 uint32, p1 uint8) int32 {
    v20 := []int{42, -925}
    fmt.Println(p0, p0)
    p0 -= p0
    _ = p0
    _ = p1
    _ = v20
    return -700
}

func main() {
    v1 := []string{"\n\t", "a", "hello"}
    v2 := []string{"hello", "a"}
    v3 := make(map[int]int)
    var v4 int64 = -1
    v21 := GenS2{F0: 574, F1: 0, F2: 255}
    v3[-875] = -588
    switch "a" {
    case "hello":
        var v22 uint64 = 255
        _ = v22
    case "foo bar":
        _ = len(v3)
    }
    if false {
        if (0 != 0) {
            var v23 float64 = (1.5 / 1.5)
            var v24 int16 = -908
            _ = v23
            _ = v24
        }
    } else {
        var v25 string = ("a" + (("foo bar" + ("foo bar" + "hello")) + "a"))
        switch ((v25 + "foo bar") + ("hello" + (("" + v25) + v25))) {
        case "\n\t":
            for fi26 := 0; fi26 < 13; fi26++ {
                if (((0.0 - (1.5 * -2.25)) / 1.0) >= 3.14159) {
                    var v27 bool = (-32768 == 0)
                    var v28 int32 = -266
                    _ = v27
                    _ = v28
                } else {
                    switch (("foo bar" + (v25 + "")) + v25) {
                    case "foo bar":
                        v29 := v3[127]
                        _ = v29
                    }
                    for fi30 := 0; fi30 < 6; fi30++ {
                        switch v25 {
                        case "":
                            fmt.Println(v4, fi26, fi26)
                        case "a":
                            v31 := GenS3{F0: -1, F1: 4294967295}
                            _ = v31
                        }
                        fmt.Println(fi30, fi30)
                        v32 := GenS2{F0: 42, F1: (-(fi30)), F2: 2}
                        _ = fi30
                        _ = v32
                    }
                    if len(v1) > 0 {
                        v1[0] = (v25 + "")
                    }
                }
                var v33 bool = true
                v34 := GenS2{F0: 1, F1: fi26, F2: 4294967295}
                _ = fi26
                _ = v33
                _ = v34
            }
        }
        if true {
            var v35 int32 = 2147483647
            var v36 uint32 = 0
            _ = v35
            _ = v36
        }
        _ = v25
    }
    switch 2147483647 {
    case 0:
        v4 = v4
    case 42:
        for fi37 := 0; fi37 < 17; fi37++ {
            for fi38 := 0; fi38 < 19; fi38++ {
                fmt.Println(v4, fi37)
                var v39 uint8 = 2
                _ = fi38
                _ = v39
            }
            _ = fi37
        }
    case 32768:
        var v40 uint64 = 2
        _ = v40
    default:
        var v41 string = (("" + "") + "a")
        _ = v41
    }
    v42 := GenS4{F0: 4294967295}
    var v43 int8 = 81
    if false {
        fmt.Println(v43, v4)
    }
    var v44 int8 = 42
    v44 = (v43 ^ (v43 >> ((127) & 63)))
    v43 = (v43 & -3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
}
