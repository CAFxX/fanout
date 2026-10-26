package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool; F1 int; F2 int32 }
type GenS4 struct { F0 int32 }

func main() {
    v1 := make(map[int]int)
    var v2 int32 = 2147483647
    v3 := make(map[int]string)
    var v4 int64 = 508
    for fi5 := 0; fi5 < 9; fi5++ {
        for fi6 := 0; fi6 < 16; fi6++ {
            fmt.Println(v2, fi6, fi6)
            var v7 bool = true
            _ = fi6
            _ = v7
        }
        v8 := GenS1{F0: 127, F1: -104}
        v9 := GenS2{F0: 100}
        _ = fi5
        _ = v8
        _ = v9
    }
    switch "" {
    case "hello":
        v2 += 0
    default:
        var v10 uint64 = 2
        _ = v10
    }
    if false {
        switch "\n\t" {
        case "foo bar":
            for fi11 := 0; fi11 < 13; fi11++ {
                for fi12 := 0; fi12 < 6; fi12++ {
                    v2 = v2
                    v4 = -9223372036854775808
                    _ = fi12
                }
                var v13 uint16 = 42
                _ = fi11
                _ = v13
            }
        case "hello":
            v4 = v4
        case "\n\t":
            v14 := GenS1{F0: 26, F1: 42}
            _ = v14
        default:
            v15 := []int{0, -944}
            _ = v15
        }
        for fi16 := 0; fi16 < 17; fi16++ {
            _ = len(v1)
            _ = fi16
        }
        var v17 uint64 = 2
        _ = v17
    }
    var v18 int32 = v2
    for fi19 := 0; fi19 < 20; fi19++ {
        for fi20 := 0; fi20 < 16; fi20++ {
            if (true || (v2 == v18)) {
                v21 := GenS1{F0: 0, F1: 100}
                var v22 bool = true
                _ = v21
                _ = v22
            }
            _ = fi20
        }
        var v23 float64 = 1e10
        _ = fi19
        _ = v23
    }
    if false {
        if false {
            if (0 < 1) {
                var v24 float32 = 1e10
                _ = v24
            } else {
                if (32767 < 42) {
                    if (-130 != -2147483648) {
                        v4 -= 1
                        for fi25 := 0; fi25 < 1; fi25++ {
                            if false {
                                var v26 bool = (((fi25 % ((fi25) | 1)) <= fi25) || false)
                                _ = v26
                            }
                            v2 += (v2 & v2)
                            _ = fi25
                        }
                    } else {
                        var v27 uint8 = 2
                        _ = v27
                    }
                    var v28 uint16 = 255
                    var v29 bool = (210 == 2)
                    _ = v28
                    _ = v29
                } else {
                    if true {
                        delete(v3, 32767)
                        v30 := []int{9223372036854775807}
                        v18 -= v18
                        _ = v30
                    }
                    var v31 int64 = v4
                    var v32 uint32 = 65535
                    _ = v31
                    _ = v32
                }
                switch ("hello" + ("a" + "hello")) {
                case "a":
                    v33 := []int{127, -574, 9223372036854775807, 729}
                    _ = v33
                case "foo bar":
                    v34 := GenS3{F0: (0 == -490), F1: 2147483647, F2: v2}
                    _ = v34
                default:
                    v18 += v18
                }
                v4 -= v4
            }
            v35 := GenS1{F0: 127, F1: 127}
            _ = v35
        }
        v2 = v2
    }
    var v36 int8 = 127
    v37 := v3[623]
    v38 := GenS4{F0: (v18 * (^(v18)))}
    v37 = ""
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
}
