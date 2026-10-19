package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 uint32 }
type GenS2 struct { F0 int8; F1 uint32 }
func genFn0(p0 uint8, p1 string, p2 int) int {
    var v6 int64 = -628
    fmt.Println(p2, p2)
    v7 := []int{(p2 ^ p2), (p2 << ((18446744073709551615) & 63)), 127}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return -365
}
func genFn1() int16 {
    if (!true) {
        for fi8 := 0; fi8 < 7; fi8++ {
            var v9 uint64 = 1
            _ = fi8
            _ = v9
        }
    }
    var v10 int = 128
    var v11 uint64 = 1
    _ = v10
    _ = v11
    return -32768
}

func main() {
    v1 := make(map[int]bool)
    v2 := []int{32767}
    var v3 uint64 = 2
    v4 := make(map[uint32]string)
    var v5 int = -2147483648
    v3 = v3
    fmt.Println(v5, v3)
    var v12 int16 = genFn1()
    switch ("" + (("foo bar" + "a") + ("\n\t" + ""))) {
    case "":
        v13 := GenS1{F0: 2, F1: genFn0(1, (("hello" + "") + "a"), v5), F2: 0}
        _ = v13
    }
    var v14 bool = true
    v15 := GenS1{F0: 2, F1: v5, F2: uint32(v3)}
    v16 := GenS1{F0: 2, F1: (v5 + (v5 % ((genFn0(127, "\n\t", 811)) | 1))), F2: 4294967295}
    v17 := GenS1{F0: 1, F1: v5, F2: 487}
    for fi18 := 0; fi18 < 1; fi18++ {
        if ((v3 << ((0) & 63)) <= v3) {
            switch ((("" + "hello") + ("" + "\n\t")) + "a") {
            case "a":
                var v19 int = genFn0(39, ("" + "\n\t"), (fi18 >> ((42) & 63)))
                _ = v19
            case "":
                switch (^(v5)) {
                case 234:
                    var v20 float64 = 0.1
                    _ = v20
                case 32767:
                    if (1 != 42) {
                        for fi21 := 0; fi21 < 15; fi21++ {
                            delete(v1, v5)
                            _ = len(v2)
                            v4[42] = "foo bar"
                            var v22 int64 = 9223372036854775807
                            fmt.Println(v14, v3)
                            _ = fi21
                            _ = v22
                        }
                    }
                }
            case "foo bar":
                v23 := GenS2{F0: -95, F1: 2147483648}
                _ = v23
            default:
                for fi24 := 0; fi24 < 12; fi24++ {
                    v5 -= genFn0(147, "foo bar", fi24)
                    _ = fi24
                }
            }
            var v25 int32 = 0
            _ = v25
        }
        _ = len(v2)
        _ = fi18
    }
    v26 := GenS2{F0: int8(v3), F1: 694}
    v3 += (^(v3))
    if ("" >= ("\n\t" + "hello")) {
        v5 -= (v5 + 4294967295)
        var v27 float64 = 0.0
        _ = v27
    }
    if (1e10 < 0.1) {
        v28 := GenS2{F0: 1, F1: 117}
        v29 := GenS2{F0: 0, F1: 2}
        _ = v28
        _ = v29
    }
    v12 = v12
    var v30 int16 = (^(v12))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v26
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
