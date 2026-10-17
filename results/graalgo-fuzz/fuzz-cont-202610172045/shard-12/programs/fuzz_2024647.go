package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 int; F1 uint; F2 int8 }
func genFn0(p0 uint16, p1 string, p2 uint64) int32 {
    var v7 uint = uint(p2)
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 274
}
func genFn1(p0 uint16, p1 int8, p2 int64) int16 {
    if true {
        v8 := []int{}
        switch ("foo bar" + "hello") {
        case "a":
            var v9 bool = true
            _ = v9
        }
        _ = v8
    }
    fmt.Println(p2)
    if false {
        var v10 int16 = -1
        var v11 int = 663
        _ = v10
        _ = v11
    }
    p0 += p0
    _ = p0
    _ = p1
    _ = p2
    return 32767
}

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[uint32]uint32)
    var v3 uint64 = 699
    var v4 int8 = 42
    v5 := make(map[uint32]string)
    v6 := make(map[int]bool)
    var v12 uint64 = (v3 * v3)
    v3 -= (v12 / ((2) | 1))
    var v13 uint32 = 2
    if false {
        var v14 bool = false
        v15 := GenS1{F0: 127, F1: 2}
        _ = v14
        _ = v15
    }
    if false {
        if false {
            var v16 uint16 = 0
            var v17 float64 = 1.5
            v3 -= (v12 - 497)
            _ = v16
            _ = v17
        }
    } else {
        switch 186 {
        case 32768:
            v12 = v12
        case 128:
            v1 = (v1 * (v1 << ((4294967295) & 63)))
        case 1:
            if (!true) {
                v4 += v4
                v1 -= (^(v1))
                switch "\n\t" {
                case "hello":
                    fmt.Println(v3, v4, v3)
                case "":
                    for fi18 := 0; fi18 < 3; fi18++ {
                        v19 := []int{fi18}
                        v20 := GenS2{F0: (fi18 * (fi18 & (fi18 / (((fi18 / ((fi18) | 1))) | 1)))), F1: (v1 / ((v1) | 1)), F2: v4}
                        var v21 uint16 = 1
                        _ = fi18
                        _ = v19
                        _ = v20
                        _ = v21
                    }
                }
            }
        default:
            v5[(v13 ^ 71)] = "foo bar"
        }
        if (("" >= "\n\t") || ((1 < 1) || false)) {
            if true {
                fmt.Println(v3, v12, v1)
                v22 := GenS2{F0: -250, F1: 42, F2: -37}
                for fi23 := 0; fi23 < 20; fi23++ {
                    var v24 uint32 = 65535
                    _ = fi23
                    _ = v24
                }
                _ = v22
            }
            switch (("a" + ("a" + "hello")) + "\n\t") {
            case "\n\t":
                v25 := []int{-1}
                _ = v25
            case "a":
                v26 := GenS1{F0: 37, F1: 497}
                _ = v26
            case "":
                v1 -= v1
            }
        }
    }
    if (false || (v12 < v12)) {
        var v27 float64 = float64(1.5)
        var v28 int = -9223372036854775808
        switch (^(v28)) {
        case -969:
            if (!true) {
                for fi29 := 0; fi29 < 18; fi29++ {
                    switch (v28 << ((141) & 63)) {
                    case -2147483648:
                        for fi30 := 0; fi30 < 12; fi30++ {
                            v13 -= 2147483647
                            var v31 bool = (true || ((v12 & (v12 | 2)) == (^(v12))))
                            _ = fi30
                            _ = v31
                        }
                    }
                    switch ((("foo bar" + "") + "\n\t") + (((((("" + "foo bar") + "") + "foo bar") + "foo bar") + "a") + ("a" + ""))) {
                    case "":
                        fmt.Println(v13, v27, v1)
                    case "a":
                        for fi32 := 0; fi32 < 18; fi32++ {
                            var v33 uint64 = v12
                            var v34 int32 = genFn0(uint16(v13), "foo bar", v12)
                            var v35 float64 = 1.5
                            _ = fi32
                            _ = v33
                            _ = v34
                            _ = v35
                        }
                    default:
                        v36 := GenS2{F0: v28, F1: v1, F2: (^(v4))}
                        _ = v36
                    }
                    _ = fi29
                }
            } else {
                for fi37 := 0; fi37 < 5; fi37++ {
                    v38 := GenS1{F0: 70, F1: 255}
                    fmt.Println(v3, v4, v3)
                    _ = fi37
                    _ = v38
                }
            }
        default:
            var v39 int8 = v4
            _ = v39
        }
        _ = v27
        _ = v28
    } else {
        var v40 float32 = -2.25
        v12 += (v12 ^ (v12 * 0))
        _ = v40
    }
    switch 151 {
    case 127:
        v41 := GenS1{F0: (^(v4)), F1: 42}
        _ = v41
    case -9223372036854775808:
        var v42 int = -239
        _ = v42
    case 9223372036854775807:
        v43 := GenS2{F0: 2147483647, F1: v1, F2: 100}
        _ = v43
    }
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
