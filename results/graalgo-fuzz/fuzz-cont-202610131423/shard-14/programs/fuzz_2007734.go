package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 int16 }
type GenS2 struct { F0 int16 }
func genFn0(p0 int64) uint32 {
    for fi4 := 0; fi4 < 11; fi4++ {
        switch ("foo bar" + "hello") {
        case "\n\t":
            if (2147483648 == 5) {
                var v5 uint64 = 42
                _ = v5
            } else {
                if (0.1 == (0.1 * (((1.5 + 1.5) / float64((3.14159 / (1e10 + (1.5 * (((1e10 + (0.0 * float32((-2.25 - float64(float32(((-2.25 + 0.1) / -2.25))))))) + float32((0.1 - 1e10))) * ((0.0 - 1e10) + 0.0))))))) + 3.14159))) {
                    var v6 float32 = -2.25
                    _ = v6
                } else {
                    if false {
                        v7 := []int{fi4}
                        var v8 string = "foo bar"
                        switch v8 {
                        case "\n\t":
                            for fi9 := 0; fi9 < 18; fi9++ {
                                v10 := GenS1{F0: fi9, F1: v8, F2: -609}
                                fmt.Println(p0)
                                _ = len(v7)
                                _ = fi9
                                _ = v10
                            }
                        case "foo bar":
                            switch v8 {
                            case "hello":
                                v8 = "a"
                            case "":
                                p0 -= p0
                            case "\n\t":
                                if (65535 <= 2) {
                                    fmt.Println(fi4)
                                }
                            }
                        case "hello":
                            for fi11 := 0; fi11 < 11; fi11++ {
                                p0 -= (p0 << ((18446744073709551615) & 63))
                                _ = fi11
                            }
                        }
                        _ = v7
                        _ = v8
                    } else {
                        for fi12 := 0; fi12 < 4; fi12++ {
                            v13 := GenS1{F0: (fi4 >> ((377) & 63)), F1: ("" + "\n\t"), F2: 49}
                            var v14 uint = 605
                            v14 = v14
                            v15 := []int{fi12, fi12, fi12}
                            _ = fi12
                            _ = v13
                            _ = v14
                            _ = v15
                        }
                    }
                }
            }
        case "a":
            p0 += (p0 ^ (-(p0)))
        case "foo bar":
            fmt.Println(p0)
        default:
            v16 := []int{(fi4 / ((fi4) | 1)), 2147483647, (^(fi4))}
            _ = v16
        }
        fmt.Println(fi4)
        v17 := GenS1{F0: (fi4 - fi4), F1: "a", F2: 32767}
        _ = fi4
        _ = v17
    }
    v18 := GenS1{F0: 32768, F1: "", F2: 32767}
    v19 := GenS2{F0: 1}
    p0 -= p0
    _ = p0
    _ = v18
    _ = v19
    return 2147483647
}

func main() {
    var v1 uint16 = 32767
    var v2 uint8 = 127
    var v3 int8 = 75
    v1 = (v1 % (((v1 ^ 292)) | 1))
    var v20 float32 = 0.0
    var v21 int32 = 2147483647
    fmt.Println(v21, v1, v20)
    var v22 float32 = -2.25
    v23 := []int{778, 32768, 9223372036854775807, 1}
    v21 += 0
    if len(v23) > 0 {
        v23[0] = 1
    }
    if false {
        var v24 float64 = (1.5 / float64((v20 + v22)))
        fmt.Println(v22, v24, v20)
        v1 = v1
        _ = v24
    } else {
        v25 := GenS2{F0: 1000}
        fmt.Println(v21, v1)
        for fi26 := 0; fi26 < 6; fi26++ {
            switch ((("" + "a") + "\n\t") + (("a" + "\n\t") + ("a" + ((((("a" + "hello") + "hello") + "") + ((("\n\t" + (("a" + "a") + "a")) + "hello") + "a")) + "foo bar")))) {
            case "a":
                var v27 int64 = -9223372036854775808
                _ = v27
            case "hello":
                _ = len(v23)
            default:
                if len(v23) >= 2 {
                    v23 = v23[1:len(v23)]
                }
            }
            _ = fi26
        }
        _ = v25
    }
    v21 = (v21 << ((65535) & 63))
    switch -101 {
    case -9223372036854775808:
        fmt.Println(v3)
    default:
        fmt.Println(v1, v3)
    }
    var v28 int16 = 1
    var v29 uint16 = (v1 - (v1 * (v1 & 1)))
    v23 = append(v23, -337)
    var v30 int64 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v30
}
