package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 string; F1 uint16 }

func main() {
    var v1 uint16 = 0
    var v2 int8 = 127
    var v3 float64 = (((1e10 + (0.1 * 0.0)) / ((1.5 + 0.0) + 3.14159)) - 1.5)
    var v4 float64 = v3
    var v5 int = 32767
    if true {
        for fi6 := 0; fi6 < 10; fi6++ {
            if false {
                v2 = (-(v2))
                switch fi6 {
                case -9223372036854775808:
                    if false {
                        fmt.Println(v3, v5)
                        var v7 int = fi6
                        v8 := []int{}
                        _ = v7
                        _ = v8
                    }
                case 1:
                    var v9 uint8 = 209
                    _ = v9
                }
            } else {
                var v10 uint16 = v1
                v11 := []int{(fi6 << ((0) & 63)), v5}
                var v12 float32 = ((-2.25 / (3.14159 - 0.1)) / (0.1 + -2.25))
                _ = v10
                _ = v11
                _ = v12
            }
            var v13 int32 = -812
            v14 := GenS1{F0: "hello", F1: 55}
            _ = fi6
            _ = v13
            _ = v14
        }
        for fi15 := 0; fi15 < 1; fi15++ {
            fmt.Println(v4)
            _ = fi15
        }
    }
    switch v5 {
    case -2147483648:
        v2 = (-(v2))
    default:
        switch ("foo bar" + ("" + ((("foo bar" + "\n\t") + "") + ("hello" + "hello")))) {
        case "foo bar":
            v5 = v5
        case "\n\t":
            v16 := GenS1{F0: "a", F1: 42}
            _ = v16
        }
    }
    for fi17 := 0; fi17 < 7; fi17++ {
        v18 := GenS2{F0: 180, F1: 70}
        var v19 uint32 = 2147483648
        _ = fi17
        _ = v18
        _ = v19
    }
    if false {
        if true {
            v20 := []int{v5, 32767, (v5 << ((18446744073709551615) & 63)), v5}
            v4 -= v3
            _ = v20
        } else {
            switch v5 {
            case -2147483648:
                v21 := []int{}
                _ = v21
            case 9223372036854775807:
                fmt.Println(v5, v4)
            case 3:
                for fi22 := 0; fi22 < 10; fi22++ {
                    v23 := []int{}
                    for fi24 := 0; fi24 < 19; fi24++ {
                        v1 += 1
                        v3 -= 0.1
                        _ = fi24
                    }
                    _ = fi22
                    _ = v23
                }
            }
            for fi25 := 0; fi25 < 4; fi25++ {
                var v26 uint = 42
                v27 := []int{-2147483648}
                var v28 int32 = 0
                _ = fi25
                _ = v26
                _ = v27
                _ = v28
            }
        }
    } else {
        v29 := GenS3{F0: ("\n\t" + "hello"), F1: v1}
        _ = v29
    }
    if ((v2 / ((v2) | 1)) != 25) {
        if (true && (!true)) {
            v30 := []int{(v5 | v5), -595, (v5 ^ (v5 + -9223372036854775808))}
            var v31 int = 42
            _ = v30
            _ = v31
        } else {
            v5 -= (v5 ^ v5)
            fmt.Println(v2)
            fmt.Println(v3, v4, v1)
        }
        for fi32 := 0; fi32 < 14; fi32++ {
            var v33 float64 = -2.25
            _ = fi32
            _ = v33
        }
    } else {
        fmt.Println(v3, v4, v5)
        var v34 uint = 1
        _ = v34
    }
    var v35 int64 = -1
    v36 := []int{v5, -225, (v5 - v5)}
    var v37 uint = 18446744073709551615
    v38 := GenS1{F0: "foo bar", F1: 100}
    if len(v36) >= 2 {
        v36 = v36[1:len(v36)]
    }
    if (1 == -2147483648) {
        v36 = append(v36, 42)
    } else {
        v39 := GenS3{F0: "foo bar", F1: v1}
        if false {
            if len(v36) >= 2 {
                v36 = v36[1:len(v36)]
            }
            _ = len(v36)
            v2 = v2
        }
        var v40 uint = v37
        _ = v39
        _ = v40
    }
    v41 := GenS3{F0: "foo bar", F1: v1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v41
    _ = v5
}
