package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64; F1 uint8; F2 bool }
type GenS4 struct { F0 int; F1 uint32 }
func genFn0() int32 {
    var v6 int16 = 1000
    if false {
        for fi7 := 0; fi7 < 1; fi7++ {
            fmt.Println(v6, v6)
            var v8 string = "a"
            v9 := GenS1{F0: 42}
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    v6 -= v6
    _ = v6
    return 0
}
func genFn1(p0 string, p1 int64) string {
    for fi10 := 0; fi10 < 19; fi10++ {
        v11 := GenS2{F0: (p1 ^ 1)}
        if false {
            for fi12 := 0; fi12 < 4; fi12++ {
                p0 = "foo bar"
                p1 = (-(p1))
                v13 := GenS3{F0: 42, F1: 2, F2: (1e10 > -2.25)}
                _ = fi12
                _ = v13
            }
            p1 = p1
            fmt.Println(fi10)
        } else {
            fmt.Println(p0, p1)
            switch ((p0 + "a") + "a") {
            case "a":
                var v14 uint8 = 42
                _ = v14
            case "":
                v15 := []int{}
                _ = v15
            case "hello":
                switch fi10 {
                case 616:
                    var v16 float64 = 1.5
                    _ = v16
                case -660:
                    v17 := GenS1{F0: 32767}
                    _ = v17
                case 628:
                    p0 = p0
                default:
                    var v18 int8 = 100
                    _ = v18
                }
            default:
                for fi19 := 0; fi19 < 14; fi19++ {
                    v20 := []int{-2147483648}
                    switch "foo bar" {
                    case "hello":
                        if (true || false) {
                            if (true || (!false)) {
                                v21 := GenS3{F0: 1, F1: 255, F2: (true && (true || true))}
                                p1 = 0
                                _ = v21
                            }
                            var v22 int8 = 127
                            v23 := GenS3{F0: 4294967295, F1: 66, F2: true}
                            _ = v22
                            _ = v23
                        }
                    }
                    _ = fi19
                    _ = v20
                }
            }
            for fi24 := 0; fi24 < 16; fi24++ {
                p0 = p0
                if true {
                    switch p0 {
                    case "\n\t":
                        p1 -= (p1 | 410)
                    case "a":
                        if ((^(p1)) <= 9223372036854775807) {
                            v25 := GenS2{F0: p1}
                            var v26 float64 = (3.14159 / 1.5)
                            _ = v25
                            _ = v26
                        } else {
                            v27 := GenS4{F0: 32767, F1: 869}
                            p0 = p0
                            _ = v27
                        }
                    case "foo bar":
                        fmt.Println(fi10, p0)
                    default:
                        var v28 int8 = -78
                        _ = v28
                    }
                } else {
                    v29 := GenS1{F0: 1000}
                    _ = v29
                }
                _ = fi24
            }
        }
        var v30 int32 = -186
        _ = fi10
        _ = v11
        _ = v30
    }
    p0 = p0
    _ = p0
    _ = p1
    return (p0 + p0)
}

func main() {
    var v1 int = -9223372036854775808
    var v2 bool = true
    var v3 int16 = 1
    var v4 float64 = 1.5
    var v5 string = "\n\t"
    for fi31 := 0; fi31 < 15; fi31++ {
        switch (v1 ^ (v1 & v1)) {
        case -9223372036854775808:
            for fi32 := 0; fi32 < 1; fi32++ {
                var v33 uint64 = 42
                _ = fi32
                _ = v33
            }
        case 128:
            v3 = v3
        default:
            v3 -= (^(v3))
        }
        _ = fi31
    }
    for fi34 := 0; fi34 < 7; fi34++ {
        switch (v1 / (((fi34 ^ fi34)) | 1)) {
        case -2147483648:
            switch "\n\t" {
            case "\n\t":
                var v35 uint8 = 128
                _ = v35
            }
        case 127:
            v36 := GenS2{F0: 9223372036854775807}
            _ = v36
        }
        v37 := []int{(fi34 & (v1 & 0)), (v1 - fi34)}
        if len(v37) > 0 {
            v37[0] = fi34
        }
        _ = fi34
        _ = v37
    }
    v38 := []int{v1}
    var v39 int16 = (v3 ^ v3)
    fmt.Println(v5)
    v38 = append(v38, v1)
    fmt.Println(v4)
    v40 := GenS2{F0: 1}
    var v41 int8 = -1
    v3 -= (v39 * (v3 | (-(v3))))
    switch "\n\t" {
    case "\n\t":
        var v42 string = v5
        _ = v42
    case "":
        v3 += (v3 | (v3 % (((^(v39))) | 1)))
    case "a":
        v4 = v4
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
