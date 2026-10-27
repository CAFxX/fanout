package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16 }
func genFn0() uint16 {
    var v5 uint64 = 0
    v6 := GenS1{F0: (v5 | (v5 + v5))}
    fmt.Println(v5)
    for fi7 := 0; fi7 < 5; fi7++ {
        for fi8 := 0; fi8 < 3; fi8++ {
            fmt.Println(v5, fi7, fi7)
            _ = fi8
        }
        for fi9 := 0; fi9 < 3; fi9++ {
            v5 += v5
            switch fi7 {
            case 4294967295:
                v5 -= v5
            case 9223372036854775807:
                switch fi9 {
                case -9223372036854775808:
                    var v10 float64 = 0.1
                    _ = v10
                case 128:
                    var v11 uint16 = 32767
                    var v12 int16 = -32768
                    v13 := []int{(fi9 - (fi9 + fi7))}
                    _ = v11
                    _ = v12
                    _ = v13
                default:
                    switch ("\n\t" + "foo bar") {
                    case "\n\t":
                        fmt.Println(v5)
                        var v14 float64 = 3.14159
                        _ = v14
                    case "":
                        var v15 int16 = -32768
                        _ = v15
                    }
                }
            }
            var v16 float64 = float64(-2.25)
            _ = fi9
            _ = v16
        }
        _ = fi7
    }
    _ = v5
    _ = v6
    return 2
}

func main() {
    var v1 uint32 = 2147483648
    v2 := make(map[int]bool)
    var v3 int32 = 531
    v4 := []string{("\n\t" + "foo bar"), ("a" + "a"), "", "hello"}
    if (-1 >= 789) {
        var v17 int8 = -1
        fmt.Println(v1, v17)
        _ = v17
    } else {
        v3 = (v3 | v3)
        _ = len(v2)
    }
    v1 += (^(v1))
    v1 = v1
    var v18 float32 = 0.0
    for fi19 := 0; fi19 < 16; fi19++ {
        fmt.Println(v3)
        _ = fi19
    }
    v18 = ((((((v18 + (1.5 - float32(((1.5 + -2.25) / 1.5)))) * v18) / v18) + v18) / (float32((0.0 - 3.14159)) + (1e10 * v18))) + v18)
    switch "\n\t" {
    case "hello":
        _ = len(v4)
    default:
        fmt.Println(v1)
    }
    var v20 uint64 = 2
    v1 = 691
    if (!true) {
        v2[-252] = false
    }
    var v21 int = 2147483647
    var v22 bool = ((v1 & v1) != 42)
    for fi23 := 0; fi23 < 14; fi23++ {
        for fi24 := 0; fi24 < 18; fi24++ {
            if v22 {
                fmt.Println(fi23)
                if v22 {
                    fmt.Println(v3)
                    v25 := GenS2{F0: 42}
                    _ = v25
                } else {
                    v18 = v18
                    var v26 uint64 = (^(v20))
                    v3 = v3
                    var v27 int32 = (v3 + 2147483647)
                    _ = v26
                    _ = v27
                }
            }
            v28 := GenS1{F0: v20}
            _ = fi24
            _ = v28
        }
        var v29 float64 = 1.5
        _ = fi23
        _ = v29
    }
    v30 := GenS3{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v30
    _ = v4
}
