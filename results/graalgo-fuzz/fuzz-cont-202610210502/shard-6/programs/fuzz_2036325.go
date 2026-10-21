package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int }
func genFn0() uint16 {
    if true {
        var v5 string = ("" + "")
        v6 := []int{-487, 1}
        _ = v5
        _ = v6
    } else {
        v7 := GenS1{F0: "a"}
        _ = v7
    }
    if (-2.25 > 3.14159) {
        for fi8 := 0; fi8 < 17; fi8++ {
            switch (-(fi8)) {
            case 830:
                fmt.Println(fi8)
            case 982:
                var v9 int = (fi8 & fi8)
                _ = v9
            case 1:
                var v10 int8 = 0
                _ = v10
            default:
                fmt.Println(fi8)
            }
            v11 := []int{(fi8 * 9223372036854775807), (^(fi8))}
            _ = fi8
            _ = v11
        }
        for fi12 := 0; fi12 < 9; fi12++ {
            v13 := GenS2{F0: 127}
            var v14 int8 = -44
            for fi15 := 0; fi15 < 4; fi15++ {
                v16 := GenS1{F0: "a"}
                v17 := []int{(fi15 & fi15), (fi12 | 9223372036854775807)}
                _ = fi15
                _ = v16
                _ = v17
            }
            _ = fi12
            _ = v13
            _ = v14
        }
        for fi18 := 0; fi18 < 18; fi18++ {
            v19 := GenS1{F0: ""}
            if false {
                var v20 float32 = 0.1
                for fi21 := 0; fi21 < 6; fi21++ {
                    v22 := GenS2{F0: 46}
                    _ = fi21
                    _ = v22
                }
                _ = v20
            }
            _ = fi18
            _ = v19
        }
    }
    return 915
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint16 = 711
    v3 := []string{("" + "a"), "a", "hello"}
    v4 := make(map[string]bool)
    v2 = v2
    switch "\n\t" {
    case "":
        var v23 uint64 = 18446744073709551615
        _ = v23
    case "\n\t":
        var v24 float64 = (float64((0.1 * float32(-2.25))) * float64(-2.25))
        _ = v24
    case "hello":
        var v25 float64 = ((3.14159 / 0.1) / 1e10)
        _ = v25
    }
    var v26 uint32 = 2147483648
    v2 -= v2
    var v27 uint64 = 391
    v2 = genFn0()
    v28 := GenS1{F0: "a"}
    v27 = 241
    if false {
        v2 = (v2 | 42)
    }
    for fi29 := 0; fi29 < 2; fi29++ {
        switch "foo bar" {
        case "\n\t":
            v30 := GenS1{F0: ("hello" + "")}
            _ = v30
        default:
            v31 := GenS2{F0: 1}
            _ = v31
        }
        fmt.Println(v26)
        _ = fi29
    }
    _ = len(v4)
    v32 := GenS3{F0: -9223372036854775808}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v32
    _ = v4
}
