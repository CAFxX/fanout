package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 int64; F1 uint8; F2 bool }
type GenS4 struct { F0 uint16; F1 uint32; F2 string }
type GenS5 struct { F0 int64; F1 int64 }
type GenS6 struct { F0 int; F1 int64 }
func genFn0(p0 int64, p1 uint64) uint8 {
    if (-1 == -1) {
        for fi5 := 0; fi5 < 15; fi5++ {
            if true {
                var v6 uint32 = 65535
                _ = v6
            }
            if ((-79 < -1) || ("" >= ("a" + (("" + ("" + ("a" + ("\n\t" + ((("\n\t" + (("foo bar" + "foo bar") + "")) + "hello") + "\n\t"))))) + "\n\t")))) {
                switch (("foo bar" + "a") + "\n\t") {
                case "foo bar":
                    v7 := []int{-1}
                    _ = v7
                case "hello":
                    var v8 float64 = 1e10
                    _ = v8
                }
            } else {
                p0 = -585
                v9 := []int{(-(fi5)), fi5}
                p0 = int64(p1)
                _ = v9
            }
            v10 := []int{fi5, (fi5 / (((fi5 << ((721) & 63))) | 1)), 0}
            _ = fi5
            _ = v10
        }
        switch ("hello" + "foo bar") {
        case "":
            if (!(!(!false))) {
                for fi11 := 0; fi11 < 8; fi11++ {
                    var v12 uint32 = 835
                    _ = fi11
                    _ = v12
                }
                var v13 int16 = 1
                _ = v13
            }
        case "foo bar":
            if true {
                var v14 int = -699
                if false {
                    v15 := []int{(^(v14)), (v14 - 127), v14}
                    p1 = 18446744073709551615
                    p0 += (p0 / ((p0) | 1))
                    _ = v15
                } else {
                    p0 -= (-(p0))
                    v14 += (v14 >> ((1) & 63))
                }
                v16 := GenS1{F0: 1, F1: (^(p0))}
                _ = v14
                _ = v16
            } else {
                fmt.Println(p0, p1)
            }
        case "\n\t":
            p1 += p1
        default:
            var v17 int64 = 10
            _ = v17
        }
    } else {
        for fi18 := 0; fi18 < 8; fi18++ {
            var v19 float32 = 1.5
            v19 = v19
            var v20 string = (("" + "foo bar") + "")
            _ = fi18
            _ = v19
            _ = v20
        }
        fmt.Println(p1, p1)
    }
    p0 += p0
    for fi21 := 0; fi21 < 4; fi21++ {
        if true {
            for fi22 := 0; fi22 < 20; fi22++ {
                switch -169 {
                case 1:
                    v23 := []int{}
                    _ = v23
                default:
                    p1 += 4294967295
                }
                if true {
                    v24 := GenS2{F0: 69, F1: -1}
                    v25 := GenS3{F0: (p0 % ((p0) | 1)), F1: 0, F2: (!(!false))}
                    if (false || (!true)) {
                        v26 := GenS4{F0: 32767, F1: 65535, F2: ((("" + "a") + "hello") + "a")}
                        _ = v26
                    } else {
                        v27 := GenS2{F0: 2, F1: 539}
                        fmt.Println(p0, fi22)
                        switch ("" + ("hello" + "a")) {
                        case "a":
                            v28 := GenS1{F0: 0, F1: (p0 + (-(p0)))}
                            _ = v28
                        case "\n\t":
                            var v29 uint32 = uint32(p1)
                            _ = v29
                        case "foo bar":
                            fmt.Println(p0, fi21, fi22)
                        default:
                            p1 = 0
                        }
                        _ = v27
                    }
                    _ = v24
                    _ = v25
                } else {
                    switch ("\n\t" + "a") {
                    case "foo bar":
                        switch (fi21 + fi21) {
                        case 1:
                            fmt.Println(fi22)
                            v30 := []int{(fi21 * 0), 42, -530}
                            _ = v30
                        default:
                            var v31 uint64 = (p1 >> ((383) & 63))
                            _ = v31
                        }
                    }
                    p0 += (p0 / ((p0) | 1))
                }
                v32 := GenS5{F0: p0, F1: (^(p0))}
                _ = fi22
                _ = v32
            }
        } else {
            if false {
                v33 := GenS4{F0: 447, F1: 2147483648, F2: ("" + "foo bar")}
                fmt.Println(p0, fi21)
                v34 := []int{-1, 9223372036854775807}
                _ = v33
                _ = v34
            }
        }
        _ = fi21
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    return 128
}

func main() {
    var v1 float32 = 1.5
    var v2 uint32 = 828
    v3 := []string{(("" + ("foo bar" + "foo bar")) + "a")}
    v4 := make(map[uint32]int)
    v4[v2] = 1
    var v35 float32 = v1
    v1 = ((float32(1e10) + ((v1 / float32(((float64(v35) * (0.0 * 0.0)) / -2.25))) * v1)) / v1)
    v35 += v35
    var v36 int32 = 2147483647
    v4[(v2 * 4294967295)] = 42
    v37 := GenS6{F0: -2147483648, F1: -9223372036854775808}
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
}
