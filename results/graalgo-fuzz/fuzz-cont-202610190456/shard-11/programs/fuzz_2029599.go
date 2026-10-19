package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64; F1 uint32; F2 int16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint; F1 int16 }
type GenS7 struct { F0 int8; F1 uint }
type GenS8 struct { F0 uint64; F1 int64; F2 int16 }
func genFn0(p0 uint) int16 {
    v5 := GenS1{F0: 234, F1: -79}
    if (-42 > -31) {
        v6 := []int{307, 127, 1}
        fmt.Println(p0)
        _ = v6
    }
    for fi7 := 0; fi7 < 16; fi7++ {
        if false {
            var v8 int16 = -32768
            switch "foo bar" {
            case "":
                v9 := GenS2{F0: 34}
                _ = v9
            case "hello":
                var v10 int16 = v8
                _ = v10
            case "\n\t":
                switch (fi7 / (((fi7 % ((fi7) | 1))) | 1)) {
                case -40:
                    var v11 uint = (^(p0))
                    _ = v11
                default:
                    switch ("hello" + "\n\t") {
                    case "foo bar":
                        p0 = (p0 << (((^(p0))) & 63))
                    default:
                        if (1 > 0) {
                            p0 = p0
                            p0 += p0
                            for fi12 := 0; fi12 < 14; fi12++ {
                                if (!(int(p0) > (fi12 * fi12))) {
                                    v8 = (v8 >> ((42) & 63))
                                    v8 += -554
                                    var v13 uint16 = 74
                                    _ = v13
                                } else {
                                    var v14 int8 = -128
                                    switch "foo bar" {
                                    case "a":
                                        if false {
                                            switch -211 {
                                            case -2147483648:
                                                fmt.Println(v14, fi12, fi7)
                                            default:
                                                if true {
                                                    p0 -= p0
                                                    var v15 uint32 = 279
                                                    var v16 string = "\n\t"
                                                    _ = v15
                                                    _ = v16
                                                }
                                                fmt.Println(p0)
                                                p0 = p0
                                            }
                                            v17 := GenS3{F0: int64(v8), F1: 2147483648, F2: -1}
                                            v18 := GenS4{F0: v14}
                                            _ = v17
                                            _ = v18
                                        }
                                    default:
                                        v19 := GenS5{F0: 0}
                                        _ = v19
                                    }
                                    var v20 float32 = 1e10
                                    _ = v14
                                    _ = v20
                                }
                                fmt.Println(p0)
                                _ = fi12
                            }
                        }
                    }
                }
            }
            v21 := []int{}
            _ = v21
            _ = v8
        } else {
            fmt.Println(fi7, p0)
        }
        switch "hello" {
        case "":
            fmt.Println(fi7)
        default:
            for fi22 := 0; fi22 < 13; fi22++ {
                v23 := []int{(-(fi7)), (-(fi22)), -9223372036854775808, 721}
                _ = fi22
                _ = v23
            }
        }
        _ = fi7
    }
    _ = p0
    _ = v5
    return 32767
}
func genFn1(p0 string, p1 int8) int32 {
    v24 := GenS6{F0: 18446744073709551615, F1: genFn0(65535)}
    switch ("" + ("" + "foo bar")) {
    case "hello":
        var v25 bool = true
        _ = v25
    case "":
        p0 = p0
    }
    v26 := GenS7{F0: (p1 ^ -1), F1: 702}
    p0 = p0
    _ = p0
    _ = p1
    _ = v24
    _ = v26
    return 42
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]int)
    v3 := []string{"a"}
    var v4 int32 = 280
    v27 := GenS8{F0: 1, F1: -323, F2: 42}
    v28 := v1[166]
    if false {
        v29 := GenS5{F0: 501}
        _ = len(v1)
        _ = v29
    }
    fmt.Println(v28, v28)
    delete(v2, 2)
    for fi30 := 0; fi30 < 8; fi30++ {
        var v31 bool = true
        v32 := GenS7{F0: 1, F1: 18446744073709551615}
        var v33 uint16 = 42
        _ = fi30
        _ = v31
        _ = v32
        _ = v33
    }
    if true {
        if true {
            if (-997 >= 1000) {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        }
    }
    v4 -= v4
    var v34 int8 = -1
    v35 := GenS3{F0: 42, F1: 685, F2: 875}
    for fi36 := 0; fi36 < 7; fi36++ {
        var v37 int = 685
        _ = fi36
        _ = v37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v34
    _ = v35
    _ = v4
}
