package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 uint; F2 int }
type GenS3 struct { F0 uint; F1 string; F2 uint64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 bool; F1 string }
type GenS6 struct { F0 bool; F1 int64; F2 uint32 }
type GenS7 struct { F0 uint16 }
func genFn0(p0 uint32, p1 uint16) uint64 {
    var v5 int8 = 42
    v6 := GenS1{F0: 1}
    switch 4294967295 {
    case 4294967295:
        switch "\n\t" {
        case "foo bar":
            v7 := []int{}
            _ = v7
        case "":
            var v8 int64 = -9223372036854775808
            _ = v8
        }
    default:
        v9 := GenS1{F0: -42}
        _ = v9
    }
    v10 := GenS1{F0: -518}
    for fi11 := 0; fi11 < 13; fi11++ {
        if ((p0 + p0) == (p0 ^ (p0 - 2147483647))) {
            for fi12 := 0; fi12 < 11; fi12++ {
                p1 += (p1 / (((p1 >> ((4294967295) & 63))) | 1))
                v13 := []int{}
                _ = fi12
                _ = v13
            }
            var v14 float32 = ((0.1 - (-2.25 + 0.1)) / 0.1)
            p1 = (^(p1))
            _ = v14
        }
        var v15 int16 = 837
        p0 = p0
        _ = fi11
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v5
    _ = v6
    return 1
}
func genFn1() int16 {
    v16 := GenS1{F0: -696}
    for fi17 := 0; fi17 < 18; fi17++ {
        v18 := []int{fi17, fi17, fi17, fi17}
        var v19 bool = false
        v19 = v19
        _ = fi17
        _ = v18
        _ = v19
    }
    _ = v16
    return 32767
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]bool)
    v3 := make(map[string]int)
    var v4 int8 = -1
    var v20 uint8 = 76
    switch ("" + "hello") {
    case "a":
        var v21 uint32 = 936
        _ = v21
    case "foo bar":
        if ((!((!(!false)) && false)) && (-2147483648 != 42)) {
            var v22 uint = 42
            _ = v22
        }
    }
    switch 127 {
    case 127:
        if (true && (2 <= 42)) {
            delete(v1, 2147483647)
            var v23 string = "hello"
            switch 484 {
            case -2147483648:
                v24 := GenS2{F0: (v4 + (-(v4))), F1: 4294967295, F2: -409}
                _ = v24
            case 9223372036854775807:
                if ((!true) || (!(!(65535 > 694)))) {
                    var v25 string = "hello"
                    if false {
                        fmt.Println(v20, v25)
                        fmt.Println(v20)
                        v4 = v4
                    } else {
                        var v26 int64 = 0
                        _ = v26
                    }
                    if false {
                        for fi27 := 0; fi27 < 10; fi27++ {
                            var v28 float32 = (((-2.25 + ((-2.25 - 3.14159) + -2.25)) / -2.25) - float32(0.0))
                            _ = fi27
                            _ = v28
                        }
                        v20 -= v20
                        if true {
                            for fi29 := 0; fi29 < 16; fi29++ {
                                v30 := v1[819]
                                v31 := GenS3{F0: 0, F1: "foo bar", F2: 42}
                                v25 = v23
                                _ = fi29
                                _ = v30
                                _ = v31
                            }
                            v32 := GenS4{F0: 32768}
                            var v33 int16 = 42
                            _ = v32
                            _ = v33
                        }
                    }
                    _ = v25
                } else {
                    fmt.Println(v4, v23)
                    v4 -= v4
                    fmt.Println(v4, v23)
                }
            }
            _ = v23
        } else {
            var v34 float32 = float32((3.14159 / -2.25))
            v35 := GenS5{F0: true, F1: ""}
            _ = v34
            _ = v35
        }
    default:
        var v36 uint32 = 42
        _ = v36
    }
    fmt.Println(v4, v4)
    var v37 uint64 = 0
    v38 := GenS6{F0: true, F1: 9223372036854775807, F2: 1}
    fmt.Println(v4)
    v39 := GenS7{F0: 682}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
}
