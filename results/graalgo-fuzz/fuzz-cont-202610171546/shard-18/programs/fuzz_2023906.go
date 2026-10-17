package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 uint; F1 int64; F2 string }
func genFn0(p0 uint8, p1 int, p2 int8) int {
    for fi6 := 0; fi6 < 11; fi6++ {
        fmt.Println(p2, p1)
        switch fi6 {
        case 658:
            var v7 int = 948
            _ = v7
        case -1:
            if false {
                p0 -= p0
                var v8 float64 = ((-2.25 * 1.5) / 3.14159)
                v9 := GenS1{F0: 152, F1: 421}
                _ = v8
                _ = v9
            }
        case -467:
            v10 := GenS1{F0: 2, F1: 42}
            _ = v10
        }
        for fi11 := 0; fi11 < 16; fi11++ {
            fmt.Println(p2)
            _ = fi11
        }
        _ = fi6
    }
    var v12 uint8 = 198
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return p1
}
func genFn1(p0 uint8, p1 uint8, p2 int32) uint {
    switch -835 {
    case 32767:
        if false {
            switch 390 {
            case 32768:
                v13 := []int{4294967295, 1}
                _ = v13
            default:
                var v14 uint = 154
                _ = v14
            }
            var v15 uint8 = p0
            _ = v15
        }
    }
    fmt.Println(p1, p0, p0)
    fmt.Println(p1)
    var v16 int16 = 32767
    v17 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    return 731
}

func main() {
    var v1 float32 = (float32(-2.25) / float32(0.1))
    var v2 int = -1
    var v3 uint = 4294967295
    var v4 int16 = 0
    var v5 float64 = (3.14159 * 1e10)
    switch "foo bar" {
    case "":
        v4 += v4
    }
    v18 := []int{9223372036854775807}
    var v19 uint8 = 2
    var v20 int8 = 8
    var v21 int32 = -2147483648
    v5 -= (v5 / 1e10)
    for fi22 := 0; fi22 < 10; fi22++ {
        switch ("hello" + "") {
        case "\n\t":
            v5 = v5
        default:
            var v23 int8 = v20
            _ = v23
        }
        _ = fi22
    }
    for fi24 := 0; fi24 < 13; fi24++ {
        var v25 string = "a"
        for fi26 := 0; fi26 < 16; fi26++ {
            switch 127 {
            case -2:
                if len(v18) > 0 {
                    v18[0] = fi26
                }
            case -545:
                v1 = v1
            case -640:
                if (v20 != v20) {
                    var v27 int64 = -9223372036854775808
                    var v28 bool = (0 < 255)
                    if v28 {
                        var v29 int8 = v20
                        var v30 string = v25
                        v3 -= v3
                        _ = v29
                        _ = v30
                    }
                    _ = v27
                    _ = v28
                } else {
                    v31 := GenS1{F0: 368, F1: (^(v3))}
                    var v32 uint64 = 1
                    _ = v31
                    _ = v32
                }
            }
            _ = fi26
        }
        _ = fi24
        _ = v25
    }
    v33 := GenS1{F0: 2, F1: (v3 - 1)}
    var v34 uint8 = 127
    for fi35 := 0; fi35 < 2; fi35++ {
        fmt.Println(v20, v5)
        _ = fi35
    }
    var v36 uint16 = 42
    for fi37 := 0; fi37 < 1; fi37++ {
        v36 += (^(v36))
        v19 = uint8(v3)
        switch fi37 {
        case 1:
            v38 := GenS1{F0: 1000, F1: genFn1((v34 | (v19 ^ (v34 * v34))), v34, v21)}
            _ = v38
        case -248:
            switch (("hello" + "\n\t") + "a") {
            case "\n\t":
                v36 -= (v36 % (((^(v36))) | 1))
            }
        case 2147483647:
            v39 := GenS2{F0: 4294967295, F1: "hello"}
            _ = v39
        default:
            if (!false) {
                v40 := GenS2{F0: 496, F1: (("foo bar" + ("a" + "a")) + "")}
                v21 += v21
                _ = v40
            }
        }
        _ = fi37
    }
    v5 = 0.0
    if (!true) {
        fmt.Println(v3)
        switch "\n\t" {
        case "foo bar":
            v18 = append(v18, genFn0(v34, 690, (v20 - 42)))
        case "\n\t":
            v41 := GenS2{F0: 42, F1: "foo bar"}
            _ = v41
        default:
            v42 := GenS3{F0: (^(v3)), F1: 182, F2: "hello"}
            _ = v42
        }
    } else {
        if false {
            if len(v18) >= 2 {
                v18 = v18[1:len(v18)]
            }
            fmt.Println(v34, v2, v19)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v33
    _ = v34
    _ = v36
    _ = v4
    _ = v5
}
