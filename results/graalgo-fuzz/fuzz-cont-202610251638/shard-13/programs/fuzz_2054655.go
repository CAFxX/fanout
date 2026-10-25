package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 int32; F1 uint32; F2 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint8; F1 uint8 }
func genFn0(p0 int8) int {
    v5 := []int{4294967295}
    v5 = append(v5, -377)
    v6 := GenS1{F0: 42, F1: 195}
    p0 = (p0 * (p0 % ((-43) | 1)))
    p0 = -1
    _ = p0
    _ = v5
    _ = v6
    return 834
}

func main() {
    v1 := []string{}
    v2 := make(map[string]string)
    var v3 uint8 = 2
    var v4 float64 = float64((1.5 * (0.1 / -2.25)))
    var v7 float64 = float64(3.14159)
    for fi8 := 0; fi8 < 12; fi8++ {
        switch fi8 {
        case 9223372036854775807:
            switch "hello" {
            case "hello":
                var v9 int64 = 0
                _ = v9
            default:
                fmt.Println(v7, v7)
            }
        default:
            var v10 int = genFn0(-59)
            _ = v10
        }
        _ = fi8
    }
    v11 := GenS2{F0: -574, F1: 689, F2: "foo bar"}
    v12 := v2["foo bar"]
    var v13 int32 = -2147483648
    switch 933 {
    case -339:
        v14 := GenS1{F0: 0, F1: 2}
        _ = v14
    }
    v15 := GenS2{F0: (^(v13)), F1: 2147483648, F2: ((v12 + (v12 + "a")) + v12)}
    if true {
        v3 = v3
        fmt.Println(v4, v7)
        for fi16 := 0; fi16 < 16; fi16++ {
            fmt.Println(v4, v7)
            v17 := v2[v12]
            _ = fi16
            _ = v17
        }
    }
    _ = len(v1)
    for fi18 := 0; fi18 < 20; fi18++ {
        fmt.Println(v4, fi18)
        v7 += v7
        _ = fi18
    }
    for fi19 := 0; fi19 < 3; fi19++ {
        switch v12 {
        case "a":
            var v20 uint16 = 2
            _ = v20
        case "foo bar":
            _ = len(v1)
        case "\n\t":
            v12 = (v12 + "foo bar")
        }
        v21 := GenS1{F0: -479, F1: 255}
        for fi22 := 0; fi22 < 1; fi22++ {
            var v23 float64 = v4
            switch v12 {
            case "foo bar":
                switch "a" {
                case "hello":
                    switch 128 {
                    case 32767:
                        v24 := GenS2{F0: v13, F1: 2147483648, F2: ((("" + ("" + "")) + "\n\t") + v12)}
                        _ = v24
                    default:
                        switch (fi19 * (-(fi22))) {
                        case 0:
                            v12 = "a"
                        }
                    }
                    var v25 bool = true
                    _ = v25
                case "a":
                    v13 = v13
                default:
                    var v26 string = v12
                    _ = v26
                }
            default:
                fmt.Println(v4, fi22, v23)
            }
            v12 = (v12 + "a")
            _ = fi22
            _ = v23
        }
        _ = fi19
        _ = v21
    }
    v27 := GenS3{F0: (1e10 >= 3.14159)}
    for fi28 := 0; fi28 < 19; fi28++ {
        for fi29 := 0; fi29 < 11; fi29++ {
            v30 := v2["a"]
            v13 = v13
            v4 -= v7
            _ = fi29
            _ = v30
        }
        v31 := GenS3{F0: false}
        switch 9223372036854775807 {
        case -2147483648:
            if len(v1) > 0 {
                v1[0] = ("\n\t" + v12)
            }
        case 1:
            v7 -= (v7 + float64((-2.25 - 1e10)))
        default:
            for fi32 := 0; fi32 < 12; fi32++ {
                v1 = append(v1, (v12 + v12))
                _ = fi32
            }
        }
        _ = fi28
        _ = v31
    }
    switch 4294967295 {
    case 802:
        var v33 float64 = float64(0.0)
        _ = v33
    case -852:
        var v34 uint = 0
        _ = v34
    }
    if false {
        v1 = append(v1, ((v12 + v12) + ("hello" + v12)))
        v35 := GenS4{F0: (v3 * v3), F1: v3}
        var v36 float64 = v7
        _ = v35
        _ = v36
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v7
}
