package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 int16; F1 int; F2 uint32 }
type GenS3 struct { F0 int64; F1 int; F2 bool }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int; F1 uint16 }
func genFn0() uint64 {
    var v6 int8 = -1
    var v7 string = (("" + (("a" + "\n\t") + (("hello" + "a") + ("\n\t" + "a")))) + "foo bar")
    _ = v6
    _ = v7
    return 4294967295
}
func genFn1(p0 int16, p1 bool, p2 int32) int8 {
    var v8 int8 = 112
    v9 := []int{4294967295}
    if p1 {
        v10 := GenS1{F0: false, F1: 1}
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 42
}

func main() {
    v1 := []string{}
    var v2 int = -9223372036854775808
    var v3 int16 = -32768
    v4 := make(map[int]int)
    var v5 int = 32767
    if true {
        for fi11 := 0; fi11 < 4; fi11++ {
            for fi12 := 0; fi12 < 1; fi12++ {
                fmt.Println(v5, v5)
                v13 := GenS1{F0: false, F1: genFn1(-32768, (true || (255 == 1)), 42)}
                _ = fi12
                _ = v13
            }
            _ = fi11
        }
        var v14 uint8 = 0
        _ = v14
    }
    v2 -= (v2 * (v5 + v5))
    v15 := GenS1{F0: ((512 >= 32767) || true), F1: -69}
    var v16 uint64 = 1
    v17 := GenS1{F0: (false && false), F1: -1}
    v18 := GenS2{F0: v3, F1: (v5 * (-(v5))), F2: 4294967295}
    v3 = v3
    if (1 == -137) {
        if (1 != 42) {
            fmt.Println(v2)
            switch ("a" + ((("a" + "a") + ("hello" + ("a" + "hello"))) + "foo bar")) {
            case "":
                _ = len(v1)
            }
        }
    } else {
        if ((v16 >> ((2) & 63)) != v16) {
            v19 := GenS1{F0: true, F1: genFn1((v3 ^ 1000), true, 1)}
            _ = v19
        } else {
            v20 := GenS2{F0: v3, F1: (v2 - v5), F2: 0}
            v21 := GenS2{F0: (v3 * v3), F1: -9223372036854775808, F2: 2}
            _ = v20
            _ = v21
        }
        v22 := v4[v5]
        _ = v22
    }
    var v23 uint64 = (v16 * v16)
    _ = len(v4)
    fmt.Println(v5)
    switch (("a" + "hello") + "") {
    case "\n\t":
        v24 := GenS3{F0: 9223372036854775807, F1: 557, F2: false}
        _ = v24
    case "a":
        v25 := GenS4{F0: 42}
        _ = v25
    case "":
        fmt.Println(v2, v23)
    default:
        v26 := GenS5{F0: v5, F1: 42}
        _ = v26
    }
    for fi27 := 0; fi27 < 17; fi27++ {
        var v28 uint8 = 128
        _ = fi27
        _ = v28
    }
    for fi29 := 0; fi29 < 4; fi29++ {
        v30 := GenS4{F0: 0}
        _ = fi29
        _ = v30
    }
    for fi31 := 0; fi31 < 16; fi31++ {
        var v32 uint8 = 117
        for fi33 := 0; fi33 < 8; fi33++ {
            if (!((v3 << ((42) & 63)) == v3)) {
                if (v16 <= (v23 >> ((397) & 63))) {
                    var v34 string = "\n\t"
                    fmt.Println(v32)
                    v5 += (-(fi31))
                    _ = v34
                } else {
                    var v35 string = "hello"
                    v35 = "a"
                    var v36 int32 = -149
                    _ = v35
                    _ = v36
                }
                v3 = 1
                if false {
                    if true {
                        var v37 uint = 2
                        v5 = 42
                        v4[236] = 827
                        _ = v37
                    }
                }
            } else {
                v23 = 42
            }
            v38 := GenS2{F0: (v3 + 0), F1: v5, F2: 2}
            v2 = fi33
            _ = fi33
            _ = v38
        }
        var v39 int16 = (-(v3))
        _ = fi31
        _ = v32
        _ = v39
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
