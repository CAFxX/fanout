package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 int64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int16; F1 uint8; F2 uint32 }
func genFn0(p0 int16, p1 int, p2 string) bool {
    p2 = ((("hello" + (p2 + p2)) + p2) + p2)
    if true {
        if (65535 < 853) {
            v6 := GenS1{F0: 42, F1: 4294967295, F2: 42}
            var v7 int32 = 0
            if ((1 < 2147483648) || true) {
                var v8 int8 = 0
                _ = v8
            } else {
                v9 := GenS1{F0: 1, F1: 1, F2: 1}
                _ = v9
            }
            _ = v6
            _ = v7
        } else {
            var v10 uint32 = 103
            for fi11 := 0; fi11 < 13; fi11++ {
                p2 = p2
                p1 += (p1 << ((1) & 63))
                _ = fi11
            }
            fmt.Println(p1, p2)
            _ = v10
        }
    }
    switch (("a" + "a") + "") {
    case "foo bar":
        p1 = (p1 / (((-(p1))) | 1))
    case "a":
        var v12 bool = false
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    return (0 < 65535)
}

func main() {
    var v1 int32 = -140
    var v2 uint64 = 4294967295
    var v3 uint16 = 806
    v4 := make(map[int]string)
    v5 := make(map[string]bool)
    for fi13 := 0; fi13 < 13; fi13++ {
        v14 := GenS2{F0: 899}
        if (-624 == 702) {
            switch (fi13 % ((1) | 1)) {
            case -397:
                v15 := GenS3{F0: 0}
                _ = v15
            case -1:
                v16 := []int{(-(fi13))}
                _ = v16
            case -751:
                if (75 > -840) {
                    v17 := GenS4{F0: 582, F1: 127, F2: 425}
                    if (!true) {
                        for fi18 := 0; fi18 < 9; fi18++ {
                            var v19 float32 = (1.5 / ((3.14159 * 1.5) + -2.25))
                            var v20 float32 = v19
                            if ("\n\t" <= ("a" + "")) {
                                fmt.Println(v3, v2, v1)
                                v19 += v19
                                var v21 bool = true
                                v1 = (^(v1))
                                switch "\n\t" {
                                case "":
                                    switch "\n\t" {
                                    case "foo bar":
                                        var v22 uint = 787
                                        _ = v22
                                    }
                                case "hello":
                                    v2 -= v2
                                case "\n\t":
                                    if false {
                                        var v23 int = 128
                                        fmt.Println(v23, v1, v1)
                                        _ = v23
                                    }
                                default:
                                    var v24 int = (fi18 << ((65535) & 63))
                                    _ = v24
                                }
                                _ = v21
                            } else {
                                var v25 uint16 = (v3 >> ((uint(v3)) & 63))
                                v25 = v25
                                _ = v25
                            }
                            var v26 int8 = -122
                            if (!(v19 >= float32(0.1))) {
                                var v27 uint32 = 4294967295
                                fmt.Println(v26, fi18)
                                var v28 uint64 = v2
                                _ = v27
                                _ = v28
                            }
                            _ = fi18
                            _ = v19
                            _ = v20
                            _ = v26
                        }
                        var v29 uint64 = v2
                        _ = v29
                    }
                    var v30 float64 = 1e10
                    _ = v17
                    _ = v30
                }
            }
        } else {
            v3 += 420
            if (0.0 >= -2.25) {
                var v31 uint32 = 0
                _ = v31
            } else {
                fmt.Println(v1, v1, v2)
                var v32 uint16 = (v3 / (((^(v3))) | 1))
                var v33 uint32 = 1
                _ = v32
                _ = v33
            }
        }
        _ = fi13
        _ = v14
    }
    var v34 int8 = 127
    v5["foo bar"] = (((-1 != 829) && false) && (true && (0.1 > (0.1 - 3.14159))))
    if (v1 < (-(v1))) {
        if (!((v3 / ((v3) | 1)) != v3)) {
            v1 += int32(v2)
            if true {
                fmt.Println(v3, v34)
                delete(v5, "foo bar")
            }
            var v35 uint = 0
            _ = v35
        } else {
            var v36 uint16 = 32767
            var v37 int64 = 0
            v38 := []int{}
            _ = v36
            _ = v37
            _ = v38
        }
    } else {
        switch -9223372036854775808 {
        case -89:
            fmt.Println(v2, v1)
        }
        var v39 int8 = 127
        _ = v39
    }
    v2 = v2
    v3 -= (v3 << ((1) & 63))
    var v40 uint64 = v2
    var v41 int32 = v1
    v41 += 574
    v42 := GenS1{F0: 42, F1: v40, F2: -609}
    var v43 uint16 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v5
}
