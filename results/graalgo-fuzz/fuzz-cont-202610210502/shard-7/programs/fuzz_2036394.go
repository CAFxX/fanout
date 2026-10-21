package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 uint }
type GenS2 struct { F0 string; F1 string; F2 uint32 }
type GenS3 struct { F0 bool }
func genFn0(p0 uint32, p1 int16) int32 {
    switch "hello" {
    case "":
        for fi5 := 0; fi5 < 12; fi5++ {
            fmt.Println(p1, p1, fi5)
            _ = fi5
        }
    case "hello":
        v6 := []int{128, 918, -40, 9223372036854775807}
        _ = v6
    case "\n\t":
        v7 := []int{127}
        _ = v7
    default:
        var v8 string = "hello"
        _ = v8
    }
    var v9 uint8 = 127
    if (("" + "a") == "foo bar") {
        switch (("\n\t" + "a") + "foo bar") {
        case "a":
            for fi10 := 0; fi10 < 17; fi10++ {
                v11 := []int{fi10, fi10, (fi10 << ((341) & 63))}
                if false {
                    v9 += (v9 + 234)
                }
                switch "" {
                case "foo bar":
                    fmt.Println(p0, fi10)
                case "hello":
                    v12 := GenS1{F0: 42, F1: (-(p1)), F2: 0}
                    _ = v12
                }
                _ = fi10
                _ = v11
            }
        case "foo bar":
            p1 -= (^(p1))
        case "\n\t":
            var v13 int8 = 100
            _ = v13
        default:
            v14 := []int{int(p0)}
            _ = v14
        }
    }
    switch -1 {
    case 1:
        if false {
            var v15 bool = false
            if ((-2.25 / 1.5) == 0.0) {
                p0 = p0
                var v16 int = 2147483647
                fmt.Println(v16, p1, p1)
                _ = v16
            } else {
                v17 := []int{}
                _ = len(v17)
                _ = len(v17)
                _ = v17
            }
            _ = v15
        }
    default:
        v18 := GenS2{F0: "\n\t", F1: "a", F2: p0}
        _ = v18
    }
    _ = p0
    _ = p1
    _ = v9
    return -2147483648
}
func genFn1(p0 uint32) int64 {
    v19 := []int{-72, -9223372036854775808, -9223372036854775808, 4294967295}
    p0 += (^(p0))
    var v20 bool = (42 == 686)
    switch -141 {
    case 1:
        fmt.Println(p0, p0)
    case -458:
        v20 = true
    case 47:
        if len(v19) > 0 {
            v19[0] = -2147483648
        }
    default:
        fmt.Println(v20, v20)
    }
    if len(v19) > 0 {
        v19[0] = 2147483647
    }
    _ = p0
    _ = v19
    _ = v20
    return 1
}

func main() {
    var v1 float64 = 0.0
    v2 := make(map[int]int)
    var v3 int32 = -102
    var v4 uint8 = 1
    if (!true) {
        v21 := []int{-146, 4294967295, 127, 32768}
        v22 := GenS1{F0: 820, F1: 1, F2: 2}
        v3 += (^(v3))
        _ = v21
        _ = v22
    } else {
        var v23 int32 = (-(v3))
        for fi24 := 0; fi24 < 2; fi24++ {
            fmt.Println(fi24, v23, v23)
            fmt.Println(v1, fi24, v23)
            _ = fi24
        }
        _ = v23
    }
    v25 := []int{127, -9223372036854775808}
    v4 = v4
    switch -1 {
    case -1:
        v26 := GenS2{F0: "", F1: (("\n\t" + "\n\t") + ("foo bar" + ("hello" + "hello"))), F2: 1}
        _ = v26
    case 4294967295:
        v27 := GenS3{F0: (!(true || ("a" <= "\n\t")))}
        _ = v27
    }
    v1 += v1
    v2[127] = -780
    v1 = v1
    fmt.Println(v1, v3)
    v3 = (v3 % ((v3) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
}
