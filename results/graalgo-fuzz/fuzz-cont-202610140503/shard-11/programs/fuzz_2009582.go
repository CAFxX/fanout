package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 uint }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 uint16; F1 int32; F2 uint32 }
type GenS4 struct { F0 int; F1 uint16; F2 int32 }
func genFn0(p0 int32, p1 uint32) uint32 {
    p1 -= 2
    for fi4 := 0; fi4 < 2; fi4++ {
        fmt.Println(p0, p0, p0)
        _ = fi4
    }
    _ = p0
    _ = p1
    return p1
}

func main() {
    var v1 int64 = 693
    var v2 int = 127
    v3 := make(map[string]int)
    for fi5 := 0; fi5 < 8; fi5++ {
        v6 := GenS1{F0: 42, F1: ((("hello" + "") + "a") + ""), F2: 1}
        _ = fi5
        _ = v6
    }
    switch v2 {
    case 162:
        if false {
            for fi7 := 0; fi7 < 1; fi7++ {
                v1 += v1
                if ((42 >= 884) || false) {
                    switch ("a" + ("" + "")) {
                    case "":
                        if true {
                            var v8 string = "hello"
                            if false {
                                delete(v3, v8)
                                v1 = 0
                            }
                            _ = v8
                        } else {
                            v2 += -638
                            v2 -= (^(fi7))
                            if true {
                                var v9 int16 = 1000
                                v9 = -117
                                _ = v9
                            } else {
                                v10 := []int{-9223372036854775808, fi7}
                                _ = v10
                            }
                        }
                    case "hello":
                        switch ("a" + "") {
                        case "\n\t":
                            fmt.Println(v1)
                        default:
                            for fi11 := 0; fi11 < 8; fi11++ {
                                v1 = 1
                                _ = fi11
                            }
                        }
                    case "foo bar":
                        for fi12 := 0; fi12 < 7; fi12++ {
                            if (!(("foo bar" == ("hello" + "\n\t")) && (!(1 == 255)))) {
                                v13 := GenS1{F0: 0, F1: "a", F2: 0}
                                fmt.Println(v1)
                                _ = v13
                            } else {
                                if true {
                                    var v14 int = 395
                                    var v15 float64 = -2.25
                                    delete(v3, "\n\t")
                                    _ = v14
                                    _ = v15
                                } else {
                                    var v16 int64 = v1
                                    _ = v16
                                }
                                v2 -= (fi7 / (((fi12 % ((fi7) | 1))) | 1))
                                fmt.Println(v1, fi7)
                            }
                            var v17 float64 = 0.0
                            v2 += (fi12 & (fi7 | fi7))
                            _ = fi12
                            _ = v17
                        }
                    }
                    v18 := GenS2{F0: 42, F1: "foo bar"}
                    _ = v18
                }
                v19 := GenS3{F0: 0, F1: -1, F2: 65535}
                _ = fi7
                _ = v19
            }
            var v20 int32 = 640
            _ = v20
        } else {
            var v21 uint64 = 42
            for fi22 := 0; fi22 < 13; fi22++ {
                v23 := GenS4{F0: (fi22 >> ((0) & 63)), F1: 591, F2: -687}
                _ = fi22
                _ = v23
            }
            _ = v21
        }
    }
    var v24 uint = 0
    v25 := []int{(^(v2)), 2147483647, (v2 ^ (v2 + (v2 / ((845) | 1)))), (v2 + 127)}
    var v26 float64 = (0.0 - 1.5)
    v1 = v1
    for fi27 := 0; fi27 < 12; fi27++ {
        var v28 uint64 = 42
        v25 = append(v25, (-(fi27)))
        _ = fi27
        _ = v28
    }
    v2 = (-(v2))
    var v29 int = (v2 << ((65535) & 63))
    v24 = v24
    for fi30 := 0; fi30 < 19; fi30++ {
        v2 -= 784
        for fi31 := 0; fi31 < 12; fi31++ {
            var v32 uint64 = 1
            var v33 uint64 = (v32 % ((0) | 1))
            v33 = v32
            _ = fi31
            _ = v32
            _ = v33
        }
        _ = fi30
    }
    var v34 uint32 = 42
    v29 += v2
    var v35 uint8 = 128
    var v36 int64 = 823
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v34
    _ = v35
    _ = v36
}
