package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 uint32 }
type GenS4 struct { F0 int16; F1 uint }
func genFn0(p0 int8) int {
    v4 := GenS1{F0: (p0 >> ((663) & 63)), F1: 2}
    var v5 int16 = 1000
    for fi6 := 0; fi6 < 8; fi6++ {
        v7 := []int{fi6}
        fmt.Println(v5)
        p0 -= (p0 - (p0 * p0))
        _ = fi6
        _ = v7
    }
    if (!((219 > 186) || true)) {
        var v8 string = "foo bar"
        _ = v8
    } else {
        if (194 < 128) {
            var v9 int8 = (p0 >> ((18446744073709551615) & 63))
            if false {
                for fi10 := 0; fi10 < 5; fi10++ {
                    if (324 >= 0) {
                        v11 := GenS1{F0: p0, F1: 1}
                        for fi12 := 0; fi12 < 15; fi12++ {
                            v9 -= v9
                            if ((!true) && ((true && (!(4294967295 <= 42))) || (true && (0 != 4294967295)))) {
                                v9 = (-(v9))
                                v13 := []int{(fi10 ^ 2147483647), (-(fi10)), (fi10 & 32768)}
                                _ = v13
                            }
                            var v14 int32 = -1
                            _ = fi12
                            _ = v14
                        }
                        var v15 uint32 = 414
                        _ = v11
                        _ = v15
                    }
                    _ = fi10
                }
            }
            v16 := []int{-2147483648, 42, -152, 9223372036854775807}
            _ = v16
            _ = v9
        }
        fmt.Println(p0, p0)
    }
    _ = p0
    _ = v4
    _ = v5
    return -9223372036854775808
}
func genFn1(p0 int) bool {
    var v17 uint8 = 2
    v18 := GenS2{F0: -1}
    v19 := GenS3{F0: "", F1: 2}
    switch (^(p0)) {
    case 154:
        if true {
            var v20 float32 = -2.25
            v21 := []int{(p0 << ((42) & 63)), (p0 ^ p0), (p0 ^ (p0 % (((p0 * (p0 / ((p0) | 1)))) | 1)))}
            var v22 uint8 = 112
            _ = v20
            _ = v21
            _ = v22
        } else {
            for fi23 := 0; fi23 < 3; fi23++ {
                p0 -= (fi23 * (-(fi23)))
                _ = fi23
            }
        }
    }
    _ = p0
    _ = v17
    _ = v18
    _ = v19
    return false
}

func main() {
    v1 := make(map[string]string)
    var v2 float32 = 0.1
    v3 := []uint32{2147483647, 65535}
    var v24 bool = true
    var v25 uint64 = 181
    var v26 uint64 = 0
    v27 := GenS3{F0: "\n\t", F1: 673}
    v24 = (0.0 != 0.1)
    v28 := GenS4{F0: 1000, F1: 729}
    var v29 int = genFn0(-13)
    switch genFn0(-128) {
    case 32768:
        switch ("\n\t" + "hello") {
        case "foo bar":
            v25 -= v25
        case "":
            v25 = 1
        case "\n\t":
            var v30 int32 = -2147483648
            _ = v30
        }
    case 42:
        v31 := GenS3{F0: "hello", F1: 825}
        _ = v31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
}
