package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 int16 }
func genFn0() int16 {
    var v4 int16 = 42
    v4 = (v4 + (v4 >> ((65535) & 63)))
    if false {
        if (1.5 < 0.0) {
            for fi5 := 0; fi5 < 11; fi5++ {
                fmt.Println(fi5, v4)
                _ = fi5
            }
            if false {
                if (42 >= 42) {
                    v6 := []int{}
                    for fi7 := 0; fi7 < 4; fi7++ {
                        v4 -= (v4 | (v4 << ((422) & 63)))
                        for fi8 := 0; fi8 < 6; fi8++ {
                            v4 = (v4 ^ (v4 ^ v4))
                            _ = fi8
                        }
                        fmt.Println(v4, fi7)
                        _ = fi7
                    }
                    _ = v6
                } else {
                    if false {
                        v9 := GenS1{F0: v4}
                        _ = v9
                    } else {
                        switch "\n\t" {
                        case "hello":
                            v10 := GenS2{F0: 255, F1: 321}
                            _ = v10
                        case "\n\t":
                            v11 := []int{1, -599}
                            _ = v11
                        default:
                            fmt.Println(v4)
                        }
                    }
                }
                for fi12 := 0; fi12 < 14; fi12++ {
                    for fi13 := 0; fi13 < 5; fi13++ {
                        v14 := []int{(fi12 + 42)}
                        var v15 float32 = (3.14159 + 0.1)
                        var v16 uint64 = 430
                        _ = fi13
                        _ = v14
                        _ = v15
                        _ = v16
                    }
                    fmt.Println(v4, v4)
                    switch ("a" + "a") {
                    case "":
                        v17 := []int{fi12, (-(fi12))}
                        _ = v17
                    case "foo bar":
                        var v18 uint64 = 4294967295
                        _ = v18
                    }
                    _ = fi12
                }
                if (!false) {
                    fmt.Println(v4)
                } else {
                    switch "\n\t" {
                    case "a":
                        v4 -= v4
                    case "":
                        v4 -= 828
                    }
                    fmt.Println(v4)
                }
            } else {
                for fi19 := 0; fi19 < 19; fi19++ {
                    var v20 uint32 = 1
                    v4 = (v4 / ((-699) | 1))
                    _ = fi19
                    _ = v20
                }
            }
            v21 := GenS2{F0: 132, F1: 2147483647}
            _ = v21
        }
        if false {
            for fi22 := 0; fi22 < 15; fi22++ {
                v23 := GenS2{F0: 1, F1: 679}
                _ = fi22
                _ = v23
            }
            v24 := GenS3{F0: 42}
            v25 := []int{}
            _ = v24
            _ = v25
        } else {
            v4 = 528
            if false {
                var v26 uint8 = 136
                v26 = v26
                var v27 string = "\n\t"
                _ = v26
                _ = v27
            } else {
                v4 = (v4 % ((575) | 1))
                v4 = (v4 * -1)
                v4 += -1
            }
            v4 = v4
        }
        v28 := GenS1{F0: (^(v4))}
        _ = v28
    } else {
        for fi29 := 0; fi29 < 10; fi29++ {
            var v30 int64 = 1
            _ = fi29
            _ = v30
        }
        v4 = v4
    }
    _ = v4
    return v4
}
func genFn1(p0 uint16) uint64 {
    v31 := []int{-711, 1, 9223372036854775807}
    var v32 uint16 = p0
    _ = len(v31)
    _ = p0
    _ = v31
    _ = v32
    return 1
}

func main() {
    var v1 float32 = 0.1
    var v2 int16 = -1
    var v3 float32 = v1
    v33 := GenS1{F0: v2}
    switch "hello" {
    case "foo bar":
        fmt.Println(v1, v3)
    default:
        v34 := []int{-1, -488}
        _ = v34
    }
    v35 := GenS3{F0: (v2 | 0)}
    v36 := []int{-1, 1}
    v37 := GenS2{F0: 1, F1: 723}
    var v38 uint = 2
    fmt.Println(v38)
    v1 -= 3.14159
    var v39 bool = false
    _ = len(v36)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
}
