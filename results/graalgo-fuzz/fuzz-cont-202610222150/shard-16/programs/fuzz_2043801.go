package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 int32 }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 uint; F1 int; F2 string }
type GenS4 struct { F0 uint32; F1 uint32 }
type GenS5 struct { F0 uint16; F1 uint }

func main() {
    v1 := make(map[int]string)
    var v2 uint64 = 4294967295
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]int)
    v5 := []int{-9223372036854775808, -997}
    v6 := make(map[string]bool)
    for fi7 := 0; fi7 < 13; fi7++ {
        v2 = v2
        _ = fi7
    }
    v8 := GenS1{F0: -437, F1: 2147483647, F2: 922}
    var v9 uint8 = uint8(v2)
    v9 += (v9 * (v9 % ((9) | 1)))
    var v10 float64 = 0.1
    if (729 >= 275) {
        if false {
            for fi11 := 0; fi11 < 12; fi11++ {
                switch ("hello" + ("a" + (("hello" + "") + "hello"))) {
                case "\n\t":
                    _ = len(v5)
                case "hello":
                    var v12 int8 = 114
                    _ = v12
                case "foo bar":
                    v10 -= ((v10 * (float64(float32(0.1)) + (-2.25 * (v10 + ((v10 + (v10 - v10)) / v10))))) * (v10 - (0.1 * v10)))
                }
                for fi13 := 0; fi13 < 9; fi13++ {
                    v14 := GenS2{F0: v9, F1: 0}
                    v15 := GenS2{F0: (v9 << ((2) & 63)), F1: 0}
                    if (!false) {
                        var v16 int8 = -1
                        v17 := v6[(("hello" + (("foo bar" + ("\n\t" + (("foo bar" + (("\n\t" + "a") + ((("hello" + ("\n\t" + ("a" + ("" + "a")))) + "foo bar") + "hello"))) + "a"))) + "foo bar")) + "\n\t")]
                        _ = v16
                        _ = v17
                    }
                    var v18 float32 = ((1.5 / 0.1) / float32(v10))
                    fmt.Println(fi11)
                    _ = fi13
                    _ = v14
                    _ = v15
                    _ = v18
                }
                _ = fi11
            }
            v19 := GenS3{F0: 42, F1: 1, F2: ""}
            _ = v19
        } else {
            v20 := GenS3{F0: 27, F1: int(v9), F2: "hello"}
            v21 := GenS1{F0: 205, F1: -695, F2: -1}
            _ = v20
            _ = v21
        }
        v22 := GenS4{F0: 514, F1: 2147483647}
        switch "\n\t" {
        case "foo bar":
            var v23 string = "foo bar"
            _ = v23
        default:
            switch 2147483647 {
            case 1:
                v9 = v9
            case 552:
                _ = len(v5)
            case 9223372036854775807:
                if false {
                    for fi24 := 0; fi24 < 6; fi24++ {
                        fmt.Println(fi24)
                        for fi25 := 0; fi25 < 8; fi25++ {
                            v2 = v2
                            _ = len(v5)
                            fmt.Println(v2, fi25)
                            _ = fi25
                        }
                        v2 -= (v2 ^ (^(v2)))
                        _ = fi24
                    }
                    v10 -= -2.25
                } else {
                    var v26 uint = 4294967295
                    _ = v26
                }
            }
        }
        _ = v22
    } else {
        for fi27 := 0; fi27 < 9; fi27++ {
            fmt.Println(v10, fi27, v2)
            v28 := GenS5{F0: 205, F1: 255}
            _ = fi27
            _ = v28
        }
        var v29 int = 32768
        if (!false) {
            fmt.Println(v10, v9)
        } else {
            for fi30 := 0; fi30 < 16; fi30++ {
                _ = len(v6)
                _ = fi30
            }
            v9 = (v9 | v9)
        }
        _ = v29
    }
    if len(v5) > 0 {
        v5[0] = -150
    }
    var v31 int = -2147483648
    v10 = 1e10
    fmt.Println(v2, v10, v9)
    if (true && (!true)) {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v32 := GenS1{F0: -1, F1: 0, F2: 0}
        _ = v32
    }
    switch "foo bar" {
    case "hello":
        for fi33 := 0; fi33 < 14; fi33++ {
            var v34 string = "a"
            if ((^(v9)) != v9) {
                v31 -= (fi33 | (v31 & 127))
                _ = len(v3)
            } else {
                v2 += 1
                v2 -= (v2 - (v2 >> ((uint(v9)) & 63)))
            }
            _ = fi33
            _ = v34
        }
    }
    v2 -= (v2 >> ((619) & 63))
    v35 := GenS5{F0: 382, F1: 255}
    v36 := GenS2{F0: (v9 * v9), F1: int64(v2)}
    for fi37 := 0; fi37 < 5; fi37++ {
        if true {
            v38 := GenS2{F0: v9, F1: 42}
            _ = v38
        } else {
            var v39 string = "foo bar"
            _ = v39
        }
        _ = fi37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v31
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
