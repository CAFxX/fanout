package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int }
func genFn0(p0 uint8) int8 {
    if true {
        var v6 float64 = 0.0
        for fi7 := 0; fi7 < 3; fi7++ {
            v8 := []int{fi7, (fi7 % (((fi7 >> ((66) & 63))) | 1)), fi7, (fi7 ^ (^(fi7)))}
            var v9 uint64 = 18446744073709551615
            var v10 float64 = v6
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
        var v11 int32 = 0
        _ = v11
        _ = v6
    } else {
        for fi12 := 0; fi12 < 11; fi12++ {
            var v13 int64 = 1
            if ((("\n\t" + "\n\t") + "a") <= "hello") {
                v14 := []int{}
                switch 828 {
                case 1:
                    if (-1 > 42) {
                        p0 = (p0 / ((p0) | 1))
                        p0 -= (^(p0))
                    } else {
                        v14 = append(v14, (fi12 % (((fi12 | -798)) | 1)))
                        v15 := GenS1{F0: 4294967295}
                        _ = v15
                    }
                default:
                    fmt.Println(fi12, p0)
                }
                for fi16 := 0; fi16 < 9; fi16++ {
                    p0 -= (p0 >> ((0) & 63))
                    var v17 float64 = -2.25
                    var v18 string = "a"
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                _ = v14
            }
            _ = fi12
            _ = v13
        }
        p0 += p0
        var v19 uint64 = 121
        _ = v19
    }
    for fi20 := 0; fi20 < 12; fi20++ {
        v21 := []int{fi20, fi20, (fi20 - fi20)}
        _ = fi20
        _ = v21
    }
    _ = p0
    return 100
}

func main() {
    var v1 bool = false
    v2 := make(map[string]int)
    var v3 uint16 = 32767
    v4 := make(map[int]int)
    var v5 float32 = -2.25
    var v22 uint = 2
    var v23 float64 = (3.14159 + 1e10)
    v3 = v3
    fmt.Println(v23, v23, v3)
    delete(v2, "a")
    if v1 {
        switch ("" + "\n\t") {
        case "a":
            for fi24 := 0; fi24 < 9; fi24++ {
                v25 := GenS2{F0: false, F1: true}
                if false {
                    var v26 float32 = v5
                    _ = v26
                } else {
                    var v27 int = fi24
                    if v1 {
                        if false {
                            if (false || (v1 && ((2147483648 >= 65535) && (-1 > 68)))) {
                                v3 -= (v3 / ((v3) | 1))
                            }
                        }
                        var v28 int16 = 1000
                        _ = v28
                    } else {
                        v5 = float32(-2.25)
                    }
                    _ = v27
                }
                _ = fi24
                _ = v25
            }
        case "hello":
            v29 := v4[2147483647]
            _ = v29
        case "":
            v30 := GenS3{F0: ("\n\t" + ("foo bar" + "a"))}
            _ = v30
        default:
            v4[-2147483648] = 2147483647
        }
        var v31 int16 = 1000
        v32 := GenS4{F0: -2147483648}
        _ = v31
        _ = v32
    }
    fmt.Println(v1, v3)
    v33 := GenS3{F0: "a"}
    if (!false) {
        var v34 int32 = 2147483647
        _ = v34
    } else {
        var v35 uint = v22
        v36 := GenS2{F0: ((v23 + -2.25) <= v23), F1: (!true)}
        _ = v35
        _ = v36
    }
    var v37 int8 = 100
    v4[-46] = 1
    var v38 uint = 394
    var v39 uint16 = v3
    for fi40 := 0; fi40 < 20; fi40++ {
        for fi41 := 0; fi41 < 6; fi41++ {
            if true {
                v42 := []int{fi40}
                _ = v42
            }
            _ = fi41
        }
        _ = fi40
    }
    fmt.Println(v1, v3, v37)
    switch "a" {
    case "hello":
        v5 -= 3.14159
    }
    v37 = (v37 % (((v37 / ((v37) | 1))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v33
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
