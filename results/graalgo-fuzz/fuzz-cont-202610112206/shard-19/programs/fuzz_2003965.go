package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 int8) uint8 {
    switch ((("\n\t" + "hello") + ("hello" + ("hello" + "a"))) + "a") {
    case "foo bar":
        if ("a" <= ("" + (((("" + ("\n\t" + ("\n\t" + ("a" + "\n\t")))) + "hello") + ("" + (("\n\t" + ((("foo bar" + ("a" + "\n\t")) + "\n\t") + "hello")) + "hello"))) + "\n\t"))) {
            var v5 int16 = 275
            v6 := GenS1{F0: 0}
            p0 = (p0 * 127)
            _ = v5
            _ = v6
        }
    }
    for fi7 := 0; fi7 < 16; fi7++ {
        for fi8 := 0; fi8 < 17; fi8++ {
            var v9 int32 = -421
            p0 = p0
            v10 := GenS1{F0: 1}
            _ = fi8
            _ = v10
            _ = v9
        }
        var v11 int32 = 838
        v12 := []int{(fi7 ^ (fi7 << ((255) & 63))), fi7, fi7}
        _ = fi7
        _ = v11
        _ = v12
    }
    if true {
        for fi13 := 0; fi13 < 6; fi13++ {
            var v14 int32 = 42
            v14 += (v14 << ((255) & 63))
            _ = fi13
            _ = v14
        }
    }
    p0 = (p0 - (^(p0)))
    if (-951 >= -1) {
        fmt.Println(p0)
        v15 := []int{669}
        _ = v15
    }
    _ = p0
    return 42
}
func genFn1(p0 uint16, p1 uint16, p2 int64) int {
    v16 := []int{-395, 70}
    p0 -= p0
    var v17 uint32 = 1
    if (!true) {
        p1 = (p0 ^ p0)
    } else {
        _ = len(v16)
    }
    p1 += (p1 << ((891) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    return -719
}

func main() {
    v1 := []string{"\n\t", "foo bar", "\n\t", "foo bar"}
    v2 := make(map[int]int)
    v3 := make(map[uint32]bool)
    v4 := []int{2147483647, -563, -98}
    _ = len(v2)
    if false {
        var v18 uint8 = 183
        delete(v3, 2)
        _ = v18
    }
    var v19 int32 = -1
    switch 9223372036854775807 {
    case -9223372036854775808:
        if len(v1) > 0 {
            v1[0] = "foo bar"
        }
    case 128:
        switch -1 {
        case 128:
            if (false || (65535 > 1)) {
                if (!(3.14159 > (1e10 + (0.0 * (((-2.25 + (float32((-2.25 + 3.14159)) + 1.5)) + ((1e10 - ((1e10 + (float32(-2.25) * 3.14159)) * 0.0)) - 1.5)) + -2.25))))) {
                    var v20 bool = (!(862 < 1))
                    if len(v4) > 0 {
                        v4[0] = 89
                    }
                    _ = v20
                }
            } else {
                v19 = v19
                var v21 int8 = 35
                _ = v21
            }
        }
    case -167:
        if ((-2147483648 < v19) && true) {
            var v22 uint64 = 354
            for fi23 := 0; fi23 < 19; fi23++ {
                v22 += v22
                _ = fi23
            }
            var v24 uint64 = v22
            _ = v22
            _ = v24
        } else {
            if (40 < 42) {
                switch 9223372036854775807 {
                case 0:
                    v19 = (^(v19))
                case 128:
                    v19 += (v19 & v19)
                }
                if (972 != -66) {
                    if (-1 <= 148) {
                        delete(v3, 42)
                    }
                }
                v25 := GenS1{F0: 18446744073709551615}
                _ = v25
            }
        }
    }
    var v26 uint32 = 65535
    _ = len(v4)
    _ = len(v2)
    var v27 uint8 = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
