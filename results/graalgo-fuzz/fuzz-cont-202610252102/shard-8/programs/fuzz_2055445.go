package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 int64; F1 uint16 }
func genFn0(p0 int16) int64 {
    var v4 uint16 = 42
    v4 = 42
    v4 -= (v4 * (v4 ^ 255))
    for fi5 := 0; fi5 < 2; fi5++ {
        var v6 uint8 = 93
        v7 := []int{}
        _ = fi5
        _ = v6
        _ = v7
    }
    _ = p0
    _ = v4
    return 42
}

func main() {
    var v1 float64 = ((1.5 / 1e10) - float64(3.14159))
    var v2 bool = (42 <= 42)
    v3 := make(map[uint32]uint32)
    var v8 uint32 = 65535
    for fi9 := 0; fi9 < 3; fi9++ {
        v8 += (v8 | (v8 / ((v8) | 1)))
        for fi10 := 0; fi10 < 2; fi10++ {
            var v11 int16 = 1000
            _ = fi10
            _ = v11
        }
        fmt.Println(fi9, fi9, v2)
        _ = fi9
    }
    fmt.Println(v8, v1)
    _ = len(v3)
    for fi12 := 0; fi12 < 2; fi12++ {
        delete(v3, v8)
        var v13 uint64 = 0
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        delete(v3, 324)
        for fi15 := 0; fi15 < 7; fi15++ {
            v16 := GenS1{F0: (!v2)}
            switch ("a" + (("hello" + "\n\t") + ("" + "a"))) {
            case "hello":
                switch (-(fi15)) {
                case 9223372036854775807:
                    v8 -= (^(v8))
                case 2147483647:
                    var v17 uint16 = 42
                    _ = v17
                }
            case "\n\t":
                var v18 int16 = 0
                _ = v18
            case "":
                var v19 int64 = 178
                _ = v19
            default:
                v2 = (v2 && true)
            }
            switch "a" {
            case "\n\t":
                v20 := []int{}
                _ = v20
            case "foo bar":
                var v21 uint8 = 0
                _ = v21
            case "":
                var v22 string = ("\n\t" + ("hello" + ("hello" + "")))
                _ = v22
            }
            _ = fi15
            _ = v16
        }
        for fi23 := 0; fi23 < 8; fi23++ {
            v1 = ((v1 - v1) - v1)
            v24 := GenS2{F0: 18446744073709551615, F1: "hello"}
            if (4294967295 == 0) {
                v25 := GenS2{F0: 4294967295, F1: ("" + "")}
                v26 := GenS3{F0: 305, F1: 65535}
                _ = v25
                _ = v26
            } else {
                v27 := GenS1{F0: v2}
                v28 := GenS1{F0: (0 == 911)}
                var v29 int = (fi23 & -1)
                v30 := []int{(fi14 % ((v29) | 1)), 9223372036854775807, (v29 << ((4294967295) & 63)), 42}
                v31 := GenS3{F0: 1, F1: 65535}
                v1 = 3.14159
                _ = v27
                _ = v28
                _ = v29
                _ = v30
                _ = v31
            }
            _ = fi23
            _ = v24
        }
        _ = fi14
    }
    if (-1 == -542) {
        v32 := []int{32767, 4294967295, 4294967295, 4294967295}
        _ = len(v3)
        var v33 int64 = -411
        _ = v32
        _ = v33
    }
    v2 = (!("a" <= ""))
    for fi34 := 0; fi34 < 18; fi34++ {
        v1 -= v1
        _ = fi34
    }
    if false {
        v1 = float64(((((1e10 - 0.0) - 3.14159) - (float32(((v1 * v1) + v1)) * ((0.0 - -2.25) + (0.0 * (float32((-2.25 - v1)) * ((1.5 - (1.5 + 0.1)) + float32(0.0))))))) * -2.25))
    }
    switch ((("a" + "foo bar") + "hello") + "") {
    case "\n\t":
        v1 = v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v8
}
