package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int32 }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 string; F1 uint64 }
type GenS4 struct { F0 int8; F1 string }
func genFn0(p0 int32) int {
    for fi5 := 0; fi5 < 15; fi5++ {
        for fi6 := 0; fi6 < 19; fi6++ {
            v7 := []int{}
            if len(v7) > 0 {
                v7[0] = fi5
            }
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            _ = fi6
            _ = v7
        }
        var v8 uint8 = 1
        _ = fi5
        _ = v8
    }
    switch ((("a" + "") + "hello") + "\n\t") {
    case "hello":
        if true {
            var v9 uint16 = 458
            _ = v9
        }
    case "\n\t":
        p0 += (p0 | 235)
    case "a":
        var v10 int = 32767
        _ = v10
    }
    switch ("foo bar" + (("a" + "foo bar") + "hello")) {
    case "\n\t":
        for fi11 := 0; fi11 < 6; fi11++ {
            var v12 int = (fi11 | (fi11 | 4294967295))
            for fi13 := 0; fi13 < 2; fi13++ {
                if ((!true) && true) {
                    var v14 string = (("a" + "\n\t") + "foo bar")
                    v15 := []int{522, (v12 * int(p0)), fi13}
                    v16 := GenS1{F0: -128, F1: 65535, F2: (-(p0))}
                    v14 = (v14 + (v14 + (v14 + ("\n\t" + v14))))
                    _ = v14
                    _ = v15
                    _ = v16
                }
                if false {
                    v12 = fi13
                    if true {
                        var v17 int32 = p0
                        var v18 uint16 = 880
                        _ = v17
                        _ = v18
                    }
                }
                if true {
                    p0 += 102
                    v19 := []int{fi11, v12}
                    _ = v19
                } else {
                    p0 = p0
                }
                _ = fi13
            }
            switch "\n\t" {
            case "\n\t":
                fmt.Println(v12, fi11)
            case "hello":
                var v20 float64 = 0.0
                _ = v20
            }
            _ = fi11
            _ = v12
        }
    case "a":
        v21 := GenS1{F0: -128, F1: 1, F2: p0}
        _ = v21
    case "":
        switch 4294967295 {
        case 127:
            for fi22 := 0; fi22 < 17; fi22++ {
                for fi23 := 0; fi23 < 18; fi23++ {
                    p0 = p0
                    p0 = -2147483648
                    _ = fi23
                }
                for fi24 := 0; fi24 < 20; fi24++ {
                    fmt.Println(fi22, fi24)
                    p0 -= p0
                    var v25 int32 = (^(p0))
                    _ = fi24
                    _ = v25
                }
                _ = fi22
            }
        case -2147483648:
            v26 := GenS1{F0: -128, F1: 65535, F2: (p0 | (p0 - -2147483648))}
            _ = v26
        case -9223372036854775808:
            var v27 bool = false
            _ = v27
        default:
            fmt.Println(p0)
        }
    }
    v28 := []int{2147483647}
    if (42 <= -9223372036854775808) {
        p0 = (-(p0))
    } else {
        var v29 uint = 276
        p0 = 0
        _ = v29
    }
    _ = p0
    _ = v28
    return 4294967295
}
func genFn1(p0 string) uint8 {
    v30 := GenS2{F0: ("hello" + p0), F1: -2147483648}
    var v31 uint8 = 42
    v31 -= v31
    _ = p0
    _ = v30
    _ = v31
    return 2
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 190
    v3 := []uint32{2}
    var v4 int32 = 0
    if false {
        for fi32 := 0; fi32 < 14; fi32++ {
            fmt.Println(v2, v4, v4)
            fmt.Println(v4, v4)
            v33 := GenS3{F0: "a", F1: 4294967295}
            _ = fi32
            _ = v33
        }
    }
    var v34 bool = (!(true || (v2 <= v2)))
    var v35 float64 = (1e10 - ((1.5 - 0.0) + (0.1 * 1e10)))
    v36 := GenS2{F0: "\n\t", F1: v4}
    v37 := GenS4{F0: 1, F1: "foo bar"}
    var v38 float64 = ((0.1 + (v35 / (v35 / (v35 + ((1e10 * v35) + v35))))) / v35)
    for fi39 := 0; fi39 < 17; fi39++ {
        switch (fi39 | fi39) {
        case 0:
            v35 -= 3.14159
        case 872:
            switch "" {
            case "foo bar":
                fmt.Println(fi39)
            }
        default:
            v40 := GenS3{F0: ("foo bar" + "\n\t"), F1: uint64(v2)}
            _ = v40
        }
        _ = fi39
    }
    fmt.Println(v38)
    v35 = v38
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
}
