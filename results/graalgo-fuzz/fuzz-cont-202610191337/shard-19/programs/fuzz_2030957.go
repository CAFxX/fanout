package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint8; F1 int8 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]int)
    v4 := make(map[int]int)
    var v5 uint8 = 92
    var v6 int = 42
    var v7 int8 = 42
    v7 = v7
    for fi8 := 0; fi8 < 2; fi8++ {
        v6 = v6
        var v9 bool = true
        _ = fi8
        _ = v9
    }
    if ((float64(float32((0.0 * (0.1 - 1.5)))) / (-2.25 + 0.1)) >= 3.14159) {
        v10 := GenS1{F0: "\n\t"}
        _ = v10
    } else {
        var v11 uint32 = 2147483648
        _ = v11
    }
    for fi12 := 0; fi12 < 17; fi12++ {
        v13 := GenS1{F0: "a"}
        var v14 uint32 = 686
        _ = fi12
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 7; fi15++ {
        v5 = 0
        switch fi15 {
        case 2147483647:
            var v16 string = ("hello" + "\n\t")
            _ = v16
        default:
            v5 += (^(v5))
        }
        for fi17 := 0; fi17 < 8; fi17++ {
            v7 = 0
            var v18 uint32 = 4294967295
            v19 := []int{(fi15 + (^(v6)))}
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = fi15
    }
    if (v6 <= 127) {
        var v20 string = (("a" + "foo bar") + "hello")
        _ = v20
    } else {
        v21 := GenS1{F0: ("a" + "hello")}
        v5 += (v5 % (((v5 / ((242) | 1))) | 1))
        for fi22 := 0; fi22 < 12; fi22++ {
            if (42 <= 42) {
                var v23 uint64 = 0
                _ = v23
            } else {
                for fi24 := 0; fi24 < 9; fi24++ {
                    switch ((("\n\t" + "a") + (("foo bar" + ("\n\t" + ("" + "\n\t"))) + "a")) + "") {
                    case "":
                        v7 -= v7
                    case "foo bar":
                        var v25 uint32 = 359
                        _ = v25
                    default:
                        var v26 int8 = (v7 << ((65535) & 63))
                        _ = v26
                    }
                    _ = fi24
                }
            }
            _ = fi22
        }
        _ = v21
    }
    v27 := GenS2{F0: false, F1: 186}
    v6 = v6
    fmt.Println(v5, v6)
    var v28 int64 = -363
    for fi29 := 0; fi29 < 2; fi29++ {
        if (v5 != v5) {
            v30 := []int{(fi29 % ((fi29) | 1)), fi29, v6}
            for fi31 := 0; fi31 < 12; fi31++ {
                var v32 int64 = v28
                var v33 bool = true
                v6 -= (fi29 & v6)
                if len(v30) > 0 {
                    v30[0] = fi31
                }
                _ = fi31
                _ = v32
                _ = v33
            }
            v7 = v7
            _ = v30
        }
        _ = fi29
    }
    switch ("" + ("\n\t" + "foo bar")) {
    case "hello":
        for fi34 := 0; fi34 < 3; fi34++ {
            var v35 int = (fi34 % (((-(v6))) | 1))
            for fi36 := 0; fi36 < 15; fi36++ {
                var v37 int8 = 42
                v5 = v5
                v35 -= 2147483647
                _ = fi36
                _ = v37
            }
            if true {
                for fi38 := 0; fi38 < 6; fi38++ {
                    v39 := GenS3{F0: -1}
                    var v40 float64 = ((0.0 + (1.5 - (1e10 - 0.1))) - 1e10)
                    var v41 float32 = -2.25
                    _ = fi38
                    _ = v39
                    _ = v40
                    _ = v41
                }
                if (!(509 >= 18446744073709551615)) {
                    v42 := []int{127}
                    _ = v42
                } else {
                    fmt.Println(v5, v6, v5)
                    v43 := GenS4{F0: (v5 - v5), F1: v7}
                    v44 := []int{(^(v6)), v35, -979, v6}
                    _ = v43
                    _ = v44
                }
            }
            _ = fi34
            _ = v35
        }
    case "":
        v7 = (v7 | v7)
    default:
        v45 := []int{386, v6, (v6 ^ 417)}
        _ = v45
    }
    v28 = 263
    var v46 int16 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v6
    _ = v7
}
