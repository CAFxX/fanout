package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 uint; F1 int8; F2 int32 }
type GenS3 struct { F0 uint }
func genFn0() bool {
    if ("\n\t" > "foo bar") {
        var v6 int32 = 42
        _ = v6
    } else {
        if (true && false) {
            v7 := []int{4294967295, 2147483647, -2147483648, 32768}
            _ = v7
        } else {
            var v8 int16 = 1000
            switch ((((("foo bar" + (("foo bar" + "a") + "\n\t")) + "foo bar") + ("\n\t" + "a")) + "foo bar") + "foo bar") {
            case "foo bar":
                v8 = (-(v8))
            case "hello":
                var v9 float64 = 3.14159
                _ = v9
            case "\n\t":
                var v10 int16 = v8
                _ = v10
            default:
                for fi11 := 0; fi11 < 7; fi11++ {
                    v12 := []int{(fi11 % (((^(fi11))) | 1)), 1, (^(fi11)), (fi11 & (fi11 ^ fi11))}
                    _ = fi11
                    _ = v12
                }
            }
            v8 = -893
            _ = v8
        }
        for fi13 := 0; fi13 < 10; fi13++ {
            var v14 bool = (11 <= 32767)
            _ = fi13
            _ = v14
        }
    }
    var v15 int = 957
    v15 = v15
    var v16 uint32 = 94
    _ = v15
    _ = v16
    return (false || false)
}
func genFn1(p0 int32) uint32 {
    if true {
        for fi17 := 0; fi17 < 12; fi17++ {
            var v18 int = (-(fi17))
            var v19 int8 = -1
            _ = fi17
            _ = v18
            _ = v19
        }
        var v20 float32 = ((-2.25 / (0.1 / (-2.25 / 1e10))) + 0.1)
        _ = v20
    } else {
        var v21 float64 = -2.25
        p0 = (^(p0))
        _ = v21
    }
    var v22 int32 = (p0 % ((p0) | 1))
    _ = p0
    _ = v22
    return 65535
}

func main() {
    var v1 bool = true
    v2 := []string{"\n\t", (((("" + "") + ("" + ("a" + "a"))) + "hello") + "hello")}
    var v3 uint8 = 127
    v4 := []uint32{2, 42, uint32(v3)}
    v5 := make(map[uint32]bool)
    v23 := GenS1{F0: -1, F1: 42}
    v24 := GenS1{F0: -1, F1: 2}
    for fi25 := 0; fi25 < 7; fi25++ {
        var v26 uint8 = v3
        _ = fi25
        _ = v26
    }
    if false {
        for fi27 := 0; fi27 < 10; fi27++ {
            var v28 int = fi27
            fmt.Println(v3)
            var v29 int = (-(v28))
            _ = fi27
            _ = v28
            _ = v29
        }
        v1 = (579 >= 4294967295)
        v30 := GenS1{F0: 9223372036854775807, F1: 849}
        _ = v30
    }
    var v31 bool = (v1 || (true || false))
    if (2147483648 <= 780) {
        var v32 uint64 = uint64(v3)
        if len(v2) > 0 {
            v2[0] = ("a" + "a")
        }
        var v33 float32 = 1e10
        _ = v32
        _ = v33
    }
    v34 := GenS1{F0: 109, F1: 42}
    var v35 int64 = 265
    v36 := v5[715]
    v37 := GenS2{F0: 1, F1: 1, F2: -499}
    fmt.Println(v3)
    for fi38 := 0; fi38 < 5; fi38++ {
        var v39 uint64 = 2
        for fi40 := 0; fi40 < 13; fi40++ {
            v5[42] = v36
            if (385 > -410) {
                var v41 uint8 = v3
                _ = v41
            }
            _ = fi40
        }
        _ = fi38
        _ = v39
    }
    _ = len(v2)
    fmt.Println(v35)
    for fi42 := 0; fi42 < 13; fi42++ {
        v43 := GenS1{F0: (v35 / (((v35 << ((42) & 63))) | 1)), F1: 0}
        _ = fi42
        _ = v43
    }
    switch "\n\t" {
    case "":
        v35 = v35
    default:
        for fi44 := 0; fi44 < 20; fi44++ {
            v45 := GenS1{F0: v35, F1: 42}
            v3 += v3
            var v46 uint64 = 728
            _ = fi44
            _ = v45
            _ = v46
        }
    }
    v47 := GenS3{F0: 534}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v31
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v47
    _ = v5
}
