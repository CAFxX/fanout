package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 int64; F1 bool; F2 uint8 }
type GenS4 struct { F0 int8; F1 uint64 }
func genFn0(p0 uint, p1 uint32, p2 int) uint8 {
    v5 := []int{p2, (p2 & p2)}
    for fi6 := 0; fi6 < 8; fi6++ {
        switch (p2 - (p2 / ((-212) | 1))) {
        case -9223372036854775808:
            v7 := GenS1{F0: "a"}
            _ = v7
        }
        var v8 int16 = 1
        var v9 uint32 = 2147483647
        _ = fi6
        _ = v8
        _ = v9
    }
    var v10 uint16 = 2
    for fi11 := 0; fi11 < 16; fi11++ {
        fmt.Println(p1)
        fmt.Println(fi11, p2, p1)
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v5
    return 128
}
func genFn1(p0 int32, p1 int32, p2 uint16) bool {
    var v12 bool = false
    var v13 int64 = 910
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return (!(v12 || (v12 && v12)))
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int16 = -32768
    var v3 int = -2147483648
    var v4 int8 = -1
    if (!((1 < 255) || (42 != 42))) {
        v14 := GenS2{F0: ("" + "a"), F1: 255}
        var v15 uint32 = 252
        fmt.Println(v2)
        _ = v14
        _ = v15
    } else {
        if false {
            v4 = -128
        }
        for fi16 := 0; fi16 < 4; fi16++ {
            fmt.Println(v2, v4, v3)
            v17 := GenS1{F0: (("a" + "a") + ("a" + "a"))}
            var v18 uint32 = 2147483648
            _ = fi16
            _ = v17
            _ = v18
        }
        v3 -= v3
    }
    var v19 bool = (!true)
    fmt.Println(v3, v3)
    var v20 int32 = -2147483648
    var v21 float64 = 3.14159
    _ = len(v1)
    v22 := GenS1{F0: "a"}
    if v19 {
        var v23 int64 = -9223372036854775808
        _ = v23
    } else {
        v19 = (255 >= 255)
        if (1 <= 2) {
            switch ("\n\t" + "foo bar") {
            case "a":
                var v24 float64 = ((-2.25 * v21) * v21)
                _ = v24
            default:
                if v19 {
                    switch "a" {
                    case "foo bar":
                        var v25 int64 = -983
                        _ = v25
                    case "":
                        var v26 uint16 = 332
                        _ = v26
                    }
                    v27 := []int{(v3 | 32768), (v3 - -2147483648)}
                    v28 := GenS2{F0: "a", F1: 0}
                    _ = v27
                    _ = v28
                } else {
                    var v29 uint16 = 696
                    _ = v29
                }
            }
            v3 = (v3 * v3)
            v30 := GenS3{F0: -9223372036854775808, F1: ((3.14159 >= (v21 + v21)) || true), F2: 42}
            _ = v30
        }
    }
    if ((v3 % ((v3) | 1)) >= (v3 + 0)) {
        var v31 uint32 = 2147483647
        v32 := GenS4{F0: int8(v20), F1: 0}
        if ((v31 ^ v31) >= (v31 + 1)) {
            var v33 int = v3
            _ = v33
        }
        _ = v31
        _ = v32
    } else {
        v34 := v1[42]
        if (42 != genFn0(1, 2147483647, v3)) {
            var v35 uint32 = 2147483647
            var v36 int = v3
            _ = v35
            _ = v36
        }
        v19 = (0 < 2147483647)
        _ = v34
    }
    var v37 int32 = (v20 & v20)
    v19 = ((42 == 0) || (true || false))
    if (18446744073709551615 != 0) {
        var v38 string = ("" + "\n\t")
        _ = v38
    }
    for fi39 := 0; fi39 < 13; fi39++ {
        v1[2] = false
        _ = fi39
    }
    v19 = (!(418 != 688))
    v3 = v3
    v4 = (v4 * (v4 << ((65535) & 63)))
    v40 := GenS1{F0: ""}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v37
    _ = v4
    _ = v40
}
