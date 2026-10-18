package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string; F1 int32; F2 bool }
type GenS3 struct { F0 uint; F1 int64 }
func genFn0(p0 int8, p1 int16) int {
    v7 := []int{4294967295, -22, 4294967295}
    for fi8 := 0; fi8 < 3; fi8++ {
        if (p0 <= (p0 | (p0 | p0))) {
            v9 := GenS1{F0: 42}
            fmt.Println(p0)
            _ = v9
        }
        switch ("hello" + "") {
        case "hello":
            _ = len(v7)
        case "foo bar":
            _ = len(v7)
        }
        _ = fi8
    }
    var v10 int32 = 740
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    return 9223372036854775807
}
func genFn1(p0 uint32, p1 uint32) uint32 {
    for fi11 := 0; fi11 < 9; fi11++ {
        var v12 uint32 = 866
        v13 := []int{(fi11 / ((genFn0(0, 13)) | 1)), (-(fi11)), fi11, (-(fi11))}
        v12 += (p0 - p1)
        _ = fi11
        _ = v12
        _ = v13
    }
    fmt.Println(p1, p1)
    for fi14 := 0; fi14 < 8; fi14++ {
        for fi15 := 0; fi15 < 1; fi15++ {
            v16 := GenS1{F0: -2147483648}
            v17 := []int{(^(fi15)), (^(fi15))}
            _ = fi15
            _ = v16
            _ = v17
        }
        if true {
            if (("foo bar" + "foo bar") >= "a") {
                v18 := []int{-9223372036854775808}
                _ = v18
            }
            if false {
                v19 := []int{-2147483648, (-(fi14)), (-(fi14)), (fi14 - genFn0(0, 1000))}
                p1 -= (p1 >> ((565) & 63))
                _ = len(v19)
                _ = v19
            } else {
                v20 := []int{fi14, (fi14 << ((0) & 63)), genFn0(100, -773)}
                v21 := GenS1{F0: -2147483648}
                for fi22 := 0; fi22 < 19; fi22++ {
                    if false {
                        v23 := GenS2{F0: ("hello" + ""), F1: 1, F2: ("" == "\n\t")}
                        p1 -= p1
                        _ = v23
                    }
                    if true {
                        p0 += p1
                        p0 = p0
                        p0 -= (p1 << ((4294967295) & 63))
                    } else {
                        if (!(!((0 == 2147483647) && (((0.0 * 1.5) - 0.0) != (0.1 / 1.5))))) {
                            p0 = p0
                        } else {
                            p0 = (p0 << ((255) & 63))
                        }
                        switch fi22 {
                        case 9223372036854775807:
                            fmt.Println(fi22, p0)
                        case 2147483647:
                            _ = len(v20)
                        case -1:
                            v24 := GenS1{F0: 42}
                            _ = v24
                        }
                        p1 = 630
                        if (!false) {
                            if (2 != 128) {
                                fmt.Println(fi14)
                            } else {
                                _ = len(v20)
                            }
                            fmt.Println(p1)
                            p1 += 2147483647
                            if (true || true) {
                                p0 = p1
                            } else {
                                _ = len(v20)
                            }
                        }
                    }
                    var v25 int32 = 571
                    _ = fi22
                    _ = v25
                }
                _ = v20
                _ = v21
            }
            p1 -= 4294967295
        } else {
            v26 := []int{(fi14 >> ((964) & 63)), (fi14 << ((537) & 63))}
            _ = v26
        }
        _ = fi14
    }
    _ = p0
    _ = p1
    return (p0 >> ((2) & 63))
}

func main() {
    var v1 int16 = 859
    var v2 uint64 = 4294967295
    var v3 uint8 = 128
    var v4 float64 = 1e10
    var v5 uint32 = 1
    v6 := make(map[int]uint32)
    var v27 float32 = 0.1
    var v28 uint8 = (v3 * 42)
    if false {
        if false {
            fmt.Println(v5, v28)
            fmt.Println(v27, v4, v1)
            v29 := GenS3{F0: 265, F1: -9223372036854775808}
            _ = v29
        }
    } else {
        if true {
            v5 = (v5 / (((^(v5))) | 1))
            v30 := []int{}
            if len(v30) >= 2 {
                v30 = v30[1:len(v30)]
            }
            _ = v30
        } else {
            v27 = ((v27 + ((0.0 * 0.1) - 3.14159)) / (-2.25 / 1e10))
        }
        switch "" {
        case "hello":
            var v31 uint8 = v3
            _ = v31
        case "a":
            fmt.Println(v3)
        case "":
            if false {
                fmt.Println(v5, v2, v5)
            }
        }
    }
    var v32 float64 = ((v4 / 1e10) * 1e10)
    v6[127] = (v5 / ((2147483647) | 1))
    v33 := GenS1{F0: 826}
    v2 = (v2 / (((^(v2))) | 1))
    v5 = v5
    v27 = v27
    var v34 string = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
