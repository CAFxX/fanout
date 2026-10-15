package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32; F1 int16 }
type GenS4 struct { F0 uint; F1 int }
type GenS5 struct { F0 int; F1 uint }

func main() {
    v1 := make(map[int]uint32)
    var v2 float32 = (1e10 + 1e10)
    var v3 uint8 = 127
    v4 := make(map[uint32]bool)
    v5 := make(map[uint32]int)
    v3 = v3
    if true {
        v2 = float32(-2.25)
    }
    v2 = v2
    if false {
        for fi6 := 0; fi6 < 19; fi6++ {
            switch ((((("foo bar" + ("" + (("" + "a") + "\n\t"))) + "foo bar") + "foo bar") + "foo bar") + "\n\t") {
            case "hello":
                v7 := GenS1{F0: 42}
                _ = v7
            case "foo bar":
                v3 += v3
            case "a":
                var v8 int8 = -117
                _ = v8
            default:
                if (1 <= 198) {
                    for fi9 := 0; fi9 < 2; fi9++ {
                        v2 -= -2.25
                        if false {
                            v2 = 3.14159
                        }
                        v3 = (v3 * v3)
                        _ = fi9
                    }
                    v10 := GenS2{F0: 1000}
                    _ = v10
                }
            }
            for fi11 := 0; fi11 < 9; fi11++ {
                var v12 uint16 = 65535
                v1[(fi6 - (fi11 / ((fi11) | 1)))] = 65535
                _ = fi11
                _ = v12
            }
            _ = fi6
        }
        for fi13 := 0; fi13 < 5; fi13++ {
            var v14 float64 = ((3.14159 / 1.0) * ((0.1 + 1.5) / 1e10))
            v15 := []int{(fi13 / ((fi13) | 1)), (fi13 - fi13), fi13}
            _ = fi13
            _ = v14
            _ = v15
        }
        fmt.Println(v2)
    }
    v16 := GenS3{F0: 445, F1: 759}
    var v17 int64 = 1
    if true {
        switch "hello" {
        case "hello":
            if (157 >= 2147483648) {
                v18 := GenS2{F0: -32768}
                v2 -= 1e10
                if true {
                    switch "\n\t" {
                    case "":
                        var v19 float64 = ((0.0 / 1.0) * (0.1 / 0.1))
                        _ = v19
                    case "foo bar":
                        for fi20 := 0; fi20 < 20; fi20++ {
                            if (255 < 897) {
                                switch ("" + ("hello" + ("\n\t" + "hello"))) {
                                case "a":
                                    v21 := []int{(fi20 << ((65535) & 63)), (fi20 - fi20)}
                                    _ = v21
                                case "\n\t":
                                    v1[fi20] = 4294967295
                                case "hello":
                                    for fi22 := 0; fi22 < 17; fi22++ {
                                        if (127 <= -73) {
                                            v23 := GenS4{F0: 2, F1: 32767}
                                            var v24 bool = true
                                            _ = v23
                                            _ = v24
                                        }
                                        v17 -= (v17 ^ (^(v17)))
                                        _ = fi22
                                    }
                                }
                            } else {
                                fmt.Println(v2, v17)
                            }
                            _ = fi20
                        }
                    default:
                        v25 := GenS4{F0: 4294967295, F1: 128}
                        _ = v25
                    }
                    v26 := []int{-2147483648, 127, 1, 0}
                    delete(v4, 36)
                    _ = v26
                } else {
                    switch 926 {
                    case -393:
                        var v27 uint = 329
                        _ = v27
                    default:
                        v28 := GenS5{F0: 9223372036854775807, F1: 4294967295}
                        _ = v28
                    }
                    v17 -= 9223372036854775807
                }
                _ = v18
            }
        case "foo bar":
            switch 32768 {
            case 4294967295:
                v29 := []int{127, 127}
                _ = v29
            case -1:
                var v30 int = -677
                _ = v30
            }
        case "":
            v2 += v2
        }
    } else {
        v31 := []int{-1, -362, 32768}
        for fi32 := 0; fi32 < 17; fi32++ {
            var v33 int16 = 560
            var v34 int32 = -2147483648
            v35 := GenS1{F0: 18446744073709551615}
            _ = fi32
            _ = v33
            _ = v34
            _ = v35
        }
        switch "hello" {
        case "\n\t":
            switch "\n\t" {
            case "\n\t":
                var v36 int64 = (^(v17))
                _ = v36
            }
        case "":
            v37 := GenS2{F0: 1}
            _ = v37
        }
        _ = v31
    }
    var v38 int32 = 42
    if (!(!(false || (("a" >= ("\n\t" + "a")) || (true && (0 >= 269)))))) {
        fmt.Println(v3)
    }
    v38 = -2147483648
    delete(v1, 126)
    v3 -= (v3 << ((956) & 63))
    var v39 uint64 = 18446744073709551615
    if true {
        _ = len(v4)
        var v40 string = ("hello" + "hello")
        _ = len(v5)
        _ = v40
    } else {
        var v41 uint64 = (^(v39))
        var v42 uint16 = 32767
        _ = v41
        _ = v42
    }
    v17 += (^(v17))
    fmt.Println(v39, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
