package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int }
func genFn0(p0 uint32) int16 {
    if true {
        p0 -= p0
        var v5 uint64 = 42
        _ = v5
    } else {
        switch ("" + ("\n\t" + "a")) {
        case "hello":
            p0 -= p0
        case "\n\t":
            p0 = p0
        default:
            var v6 string = ((("" + "foo bar") + ("hello" + "")) + ((("a" + ("\n\t" + ("a" + "foo bar"))) + "foo bar") + "a"))
            _ = v6
        }
        p0 -= p0
        fmt.Println(p0)
    }
    p0 = (p0 / (((p0 * (p0 ^ p0))) | 1))
    _ = p0
    return 0
}

func main() {
    var v1 uint8 = 255
    var v2 uint64 = 602
    var v3 float64 = (0.0 - 1e10)
    var v4 uint8 = v1
    if (555 <= 174) {
        v7 := []int{}
        _ = v7
    }
    fmt.Println(v4)
    fmt.Println(v3)
    v8 := GenS1{F0: "", F1: -18, F2: "\n\t"}
    for fi9 := 0; fi9 < 5; fi9++ {
        var v10 float32 = -2.25
        var v11 float32 = (-2.25 - v10)
        v11 = v11
        _ = fi9
        _ = v10
        _ = v11
    }
    switch "foo bar" {
    case "hello":
        v3 = (((v3 - 1e10) - v3) * float64((1e10 / float32(1e10))))
    default:
        v12 := GenS1{F0: "a", F1: 1, F2: (("a" + "hello") + "\n\t")}
        _ = v12
    }
    v13 := []int{}
    var v14 float64 = v3
    v3 -= ((float64(1.5) * ((float64(0.1) + 0.1) / (((1.5 * (v14 - 0.0)) + (3.14159 - 3.14159)) * v14))) + v3)
    if ("foo bar" <= "a") {
        if (4294967295 <= 1) {
            v13 = append(v13, 32768)
        }
    } else {
        v4 += v4
        if false {
            var v15 int = 1000
            v3 = ((((v14 - v3) / (v3 - -2.25)) - 1.5) - ((3.14159 / ((v3 + -2.25) * 1e10)) / float64(-2.25)))
            v16 := GenS2{F0: 493}
            _ = v15
            _ = v16
        }
        v17 := GenS2{F0: 1}
        _ = v17
    }
    v18 := GenS1{F0: "", F1: -1, F2: ("a" + "\n\t")}
    v19 := GenS3{F0: 42}
    if len(v13) > 0 {
        v13[0] = 42
    }
    for fi20 := 0; fi20 < 15; fi20++ {
        var v21 int = fi20
        switch "" {
        case "hello":
            var v22 uint = 18446744073709551615
            _ = v22
        }
        v23 := GenS3{F0: v21}
        _ = fi20
        _ = v21
        _ = v23
    }
    if false {
        switch -9223372036854775808 {
        case 127:
            for fi24 := 0; fi24 < 20; fi24++ {
                switch fi24 {
                case 2147483647:
                    var v25 uint64 = v2
                    _ = v25
                case 4294967295:
                    for fi26 := 0; fi26 < 13; fi26++ {
                        var v27 string = "a"
                        var v28 uint16 = 393
                        _ = fi26
                        _ = v27
                        _ = v28
                    }
                }
                var v29 uint32 = uint32(v2)
                _ = fi24
                _ = v29
            }
        case 2147483647:
            var v30 int64 = 42
            _ = v30
        }
    }
    for fi31 := 0; fi31 < 5; fi31++ {
        for fi32 := 0; fi32 < 2; fi32++ {
            var v33 uint16 = 2
            var v34 int8 = 4
            v14 = 0.1
            _ = fi32
            _ = v33
            _ = v34
        }
        var v35 int16 = 42
        _ = fi31
        _ = v35
    }
    v14 = v14
    var v36 int32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v8
}
