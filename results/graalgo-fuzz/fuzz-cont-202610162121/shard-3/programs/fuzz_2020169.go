package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 string; F1 string; F2 int64 }
type GenS3 struct { F0 string }
func genFn0() int {
    var v6 int16 = -1
    v7 := GenS1{F0: 2147483647, F1: (-2147483648 >= -965)}
    v6 -= (v6 & v6)
    v8 := []int{32767, 2147483647, -9223372036854775808, 128}
    v6 = v6
    _ = v6
    _ = v7
    _ = v8
    return 802
}

func main() {
    var v1 uint = 728
    v2 := make(map[uint32]bool)
    var v3 int64 = -1
    v4 := make(map[string]string)
    v5 := make(map[string]uint32)
    var v9 int = -2147483648
    switch "hello" {
    case "\n\t":
        if false {
            for fi10 := 0; fi10 < 3; fi10++ {
                if (false || ((((863 >= 42) || false) && false) && false)) {
                    var v11 int16 = 42
                    v3 -= 42
                    _ = v11
                } else {
                    switch "foo bar" {
                    case "hello":
                        for fi12 := 0; fi12 < 17; fi12++ {
                            v13 := []int{}
                            switch ("" + "") {
                            case "foo bar":
                                v14 := v5["a"]
                                _ = v14
                            case "hello":
                                v15 := GenS1{F0: int32(v3), F1: (35 != 260)}
                                _ = v15
                            }
                            var v16 int = (-(v9))
                            _ = fi12
                            _ = v13
                            _ = v16
                        }
                    case "\n\t":
                        v17 := GenS2{F0: "\n\t", F1: "foo bar", F2: v3}
                        _ = v17
                    default:
                        switch (fi10 % ((genFn0()) | 1)) {
                        case -842:
                            var v18 uint16 = 2
                            _ = v18
                        default:
                            v1 = 42
                        }
                    }
                }
                switch (v9 + (v9 | genFn0())) {
                case -9223372036854775808:
                    v1 = v1
                case 128:
                    fmt.Println(v9)
                case 743:
                    v19 := GenS1{F0: 165, F1: ((!((1e10 - 0.1) <= 1.5)) && false)}
                    _ = v19
                }
                _ = fi10
            }
            v20 := []int{(^(v9)), 32768, 530}
            _ = v20
        } else {
            v21 := GenS2{F0: "\n\t", F1: ("foo bar" + ("a" + "hello")), F2: (v3 << ((42) & 63))}
            _ = v21
        }
    default:
        var v22 uint8 = 2
        _ = v22
    }
    v9 = (v9 ^ (v9 % ((v9) | 1)))
    v1 = (v1 | v1)
    var v23 int64 = v3
    v9 = v9
    _ = len(v4)
    var v24 uint = v1
    v1 = 65535
    v25 := []int{}
    v24 += 18446744073709551615
    v25 = append(v25, (v9 + -9223372036854775808))
    for fi26 := 0; fi26 < 15; fi26++ {
        v24 = v24
        var v27 int = (v9 >> (((v24 << ((1) & 63))) & 63))
        _ = fi26
        _ = v27
    }
    switch (v9 - 524) {
    case 42:
        v24 -= 85
    }
    switch (v9 - v9) {
    case 17:
        v23 += (v23 & v3)
    case -643:
        switch 224 {
        case 9223372036854775807:
            var v28 int16 = 0
            _ = v28
        default:
            for fi29 := 0; fi29 < 5; fi29++ {
                var v30 int = (v9 - fi29)
                _ = fi29
                _ = v30
            }
        }
    }
    if true {
        fmt.Println(v23)
        for fi31 := 0; fi31 < 13; fi31++ {
            fmt.Println(v24, fi31)
            _ = fi31
        }
        switch v9 {
        case -9223372036854775808:
            var v32 bool = false
            _ = v32
        case 0:
            var v33 uint8 = 194
            _ = v33
        case 32767:
            if true {
                var v34 uint32 = 4294967295
                _ = v34
            } else {
                _ = len(v25)
                fmt.Println(v3, v24, v1)
                v35 := GenS3{F0: ""}
                _ = v35
            }
        }
    }
    v36 := GenS2{F0: "", F1: (("" + "hello") + (("\n\t" + "foo bar") + "\n\t")), F2: v23}
    if len(v25) >= 2 {
        v25 = v25[1:len(v25)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v9
}
