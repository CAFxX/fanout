package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 string; F1 int32; F2 int16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint; F1 int32 }
func genFn0(p0 int8, p1 uint8) int32 {
    p1 = (p1 + p1)
    p0 += 127
    if (((("\n\t" + "hello") + "\n\t") + "hello") < "") {
        p0 -= 17
        if false {
            fmt.Println(p0, p0)
        }
        v6 := GenS1{F0: 42, F1: 65535}
        _ = v6
    } else {
        for fi7 := 0; fi7 < 16; fi7++ {
            for fi8 := 0; fi8 < 10; fi8++ {
                v9 := []int{}
                var v10 uint = 65535
                _ = fi8
                _ = v10
                _ = v9
            }
            _ = fi7
        }
        v11 := []int{}
        v12 := GenS2{F0: ("\n\t" + (((("\n\t" + "") + "foo bar") + "\n\t") + "foo bar")), F1: 0, F2: 32767}
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    return -2147483648
}
func genFn1(p0 int32, p1 bool) uint32 {
    for fi13 := 0; fi13 < 13; fi13++ {
        for fi14 := 0; fi14 < 5; fi14++ {
            fmt.Println(fi13, p1, fi13)
            p0 = p0
            _ = fi14
        }
        _ = fi13
    }
    var v15 int8 = 42
    v15 = (v15 / ((v15) | 1))
    v16 := GenS3{F0: ("\n\t" + "a")}
    _ = p0
    _ = p1
    _ = v15
    _ = v16
    return 0
}

func main() {
    var v1 int16 = -175
    var v2 float32 = (-2.25 / (0.1 - float32(1e10)))
    var v3 string = ("hello" + "")
    var v4 int8 = -68
    var v5 int32 = -2147483648
    switch "a" {
    case "":
        v5 = v5
    case "hello":
        fmt.Println(v1, v2, v4)
    default:
        var v17 int = 32768
        _ = v17
    }
    switch 32768 {
    case -821:
        v5 = (v5 - 822)
    case 127:
        if false {
            v18 := []int{42}
            fmt.Println(v2, v4)
            v18 = append(v18, int(v5))
            _ = v18
        } else {
            v3 = (((("foo bar" + v3) + (v3 + v3)) + "\n\t") + v3)
            fmt.Println(v5, v5, v2)
        }
    case 32767:
        v19 := []int{2147483647, 32768}
        _ = v19
    default:
        fmt.Println(v2)
    }
    fmt.Println(v2, v4, v5)
    var v20 uint8 = 255
    fmt.Println(v5, v2)
    for fi21 := 0; fi21 < 7; fi21++ {
        v5 += v5
        _ = fi21
    }
    v2 -= v2
    v2 += ((v2 - v2) - 3.14159)
    if true {
        v1 += 1
        for fi22 := 0; fi22 < 5; fi22++ {
            for fi23 := 0; fi23 < 19; fi23++ {
                fmt.Println(fi22)
                _ = fi23
            }
            switch "a" {
            case "foo bar":
                if true {
                    for fi24 := 0; fi24 < 16; fi24++ {
                        var v25 float32 = 1e10
                        _ = fi24
                        _ = v25
                    }
                    v1 = (v1 * v1)
                    v26 := []int{(fi22 * 127), fi22}
                    _ = v26
                } else {
                    v2 -= (v2 * (((-2.25 - 0.0) * v2) - float32(1.5)))
                    fmt.Println(v3)
                }
            case "":
                if (!false) {
                    v27 := []int{2147483647, 4294967295, fi22}
                    var v28 int8 = 127
                    v29 := GenS4{F0: 788, F1: (v5 & genFn0(127, v20))}
                    _ = v27
                    _ = v28
                    _ = v29
                } else {
                    if ((v20 / (((v20 / ((v20) | 1))) | 1)) > v20) {
                        v30 := GenS4{F0: 2, F1: (v5 / ((v5) | 1))}
                        _ = v30
                    } else {
                        if true {
                            v5 = (v5 >> ((1) & 63))
                        }
                    }
                }
            default:
                v31 := []int{fi22, fi22, 2147483647}
                _ = v31
            }
            v32 := GenS2{F0: "hello", F1: 1, F2: 1000}
            _ = fi22
            _ = v32
        }
        for fi33 := 0; fi33 < 16; fi33++ {
            v34 := GenS3{F0: "\n\t"}
            _ = fi33
            _ = v34
        }
    } else {
        v2 += v2
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
