package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 bool; F1 uint64; F2 bool }
func genFn0(p0 uint32) uint32 {
    var v4 uint64 = 487
    var v5 float64 = ((1e10 * 0.0) * 3.14159)
    v6 := GenS1{F0: false, F1: v4}
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return (p0 / (((p0 & (p0 % ((p0) | 1)))) | 1))
}
func genFn1(p0 bool, p1 bool) bool {
    var v7 uint8 = 128
    switch -766 {
    case -368:
        fmt.Println(p1, p0, p0)
    case 0:
        for fi8 := 0; fi8 < 14; fi8++ {
            switch (-(fi8)) {
            case -9223372036854775808:
                for fi9 := 0; fi9 < 13; fi9++ {
                    var v10 uint32 = 1
                    fmt.Println(fi9)
                    _ = fi9
                    _ = v10
                }
            case 2147483647:
                v11 := []int{fi8, (-(fi8)), fi8}
                _ = v11
            }
            var v12 int = fi8
            p1 = p1
            _ = fi8
            _ = v12
        }
    }
    var v13 uint8 = 0
    if p0 {
        if (42 <= 4294967295) {
            switch 112 {
            case 2147483647:
                switch (((((("" + "") + ("\n\t" + "\n\t")) + "") + "foo bar") + "hello") + "\n\t") {
                case "a":
                    if true {
                        v7 = v13
                        switch ((("\n\t" + "foo bar") + "") + "a") {
                        case "foo bar":
                            fmt.Println(v13)
                        case "hello":
                            var v14 int16 = 0
                            _ = v14
                        }
                    } else {
                        switch -2147483648 {
                        case 2147483647:
                            var v15 int16 = 42
                            _ = v15
                        case 128:
                            v16 := GenS2{F0: "foo bar"}
                            _ = v16
                        }
                    }
                case "foo bar":
                    p0 = (1e10 != 0.0)
                case "":
                    if false {
                        fmt.Println(p0, p0)
                    } else {
                        var v17 uint32 = 2147483647
                        var v18 uint64 = 1
                        _ = v17
                        _ = v18
                    }
                }
            case 15:
                var v19 bool = (-862 > 384)
                _ = v19
            case 0:
                v20 := GenS3{F0: 429}
                _ = v20
            default:
                var v21 float64 = 0.1
                _ = v21
            }
        }
        for fi22 := 0; fi22 < 6; fi22++ {
            v23 := GenS1{F0: (-1 == -1), F1: 42}
            _ = fi22
            _ = v23
        }
        v24 := GenS2{F0: "foo bar"}
        _ = v24
    }
    if p0 {
        var v25 uint = 591
        if p0 {
            switch 9223372036854775807 {
            case 1:
                fmt.Println(p1, v25)
            case 364:
                var v26 uint64 = 18446744073709551615
                _ = v26
            default:
                v25 -= (^(v25))
            }
            v27 := []int{1, 9223372036854775807}
            _ = v27
        } else {
            p1 = p1
        }
        _ = v25
    } else {
        v28 := []int{500, -1, 1}
        _ = v28
    }
    _ = p0
    _ = p1
    _ = v13
    _ = v7
    return (2 >= 2)
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := make(map[int]string)
    v3 := make(map[string]uint32)
    v1 = 131
    for fi29 := 0; fi29 < 16; fi29++ {
        fmt.Println(fi29)
        _ = fi29
    }
    v1 += 967
    for fi30 := 0; fi30 < 17; fi30++ {
        if (0 >= 1) {
            v1 = (v1 * v1)
            for fi31 := 0; fi31 < 15; fi31++ {
                switch 127 {
                case 380:
                    if (!((fi31 ^ fi30) > (fi31 / ((fi31) | 1)))) {
                        v1 = v1
                    } else {
                        if false {
                            fmt.Println(fi31)
                        }
                        var v32 float64 = 0.1
                        v33 := GenS3{F0: 1}
                        _ = v32
                        _ = v33
                    }
                case 4294967295:
                    var v34 int32 = 930
                    _ = v34
                case 114:
                    if (false || ((-366 > -1) && false)) {
                        var v35 int = 32768
                        if (("hello" + "foo bar") > "a") {
                            v36 := GenS2{F0: (("foo bar" + "hello") + "")}
                            fmt.Println(fi30, fi30, v1)
                            delete(v2, fi31)
                            _ = v36
                        }
                        v35 -= 4294967295
                        _ = v35
                    } else {
                        if (42 >= 533) {
                            v1 += v1
                        }
                    }
                }
                v37 := GenS3{F0: 0}
                v38 := GenS3{F0: -824}
                _ = fi31
                _ = v37
                _ = v38
            }
        } else {
            var v39 float32 = float32((1e10 + 1.5))
            v40 := GenS4{F0: (0 <= 1000), F1: (v1 << ((804) & 63)), F2: (true || true)}
            v1 = v1
            _ = v39
            _ = v40
        }
        var v41 uint32 = 65535
        _ = fi30
        _ = v41
    }
    v42 := GenS2{F0: "foo bar"}
    var v43 int16 = -377
    fmt.Println(v1, v43)
    for fi44 := 0; fi44 < 6; fi44++ {
        v43 = (v43 - v43)
        var v45 int16 = 32767
        _ = fi44
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v42
    _ = v43
}
