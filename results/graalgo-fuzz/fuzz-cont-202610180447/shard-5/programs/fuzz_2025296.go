package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16; F1 uint8; F2 int64 }
type GenS4 struct { F0 uint64; F1 int8 }
func genFn0() int16 {
    var v5 int32 = -2147483648
    fmt.Println(v5)
    switch 128 {
    case 128:
        if true {
            v6 := GenS1{F0: 2147483647, F1: 42, F2: 4294967295}
            v5 = (v5 - -2147483648)
            _ = v6
        }
    }
    v5 += (v5 << ((18446744073709551615) & 63))
    v5 = -2147483648
    _ = v5
    return -1
}
func genFn1() uint8 {
    if (502 != 0) {
        var v7 uint32 = 4294967295
        if ((float64((-2.25 + 0.1)) - ((-2.25 / 1.5) / ((1.5 + -2.25) / 3.14159))) == (1e10 - 1e10)) {
            if false {
                fmt.Println(v7)
                var v8 int64 = -1
                if (("foo bar" + ("" + ("foo bar" + ""))) > (("a" + "foo bar") + "hello")) {
                    v9 := []int{}
                    for fi10 := 0; fi10 < 20; fi10++ {
                        var v11 uint32 = 4294967295
                        for fi12 := 0; fi12 < 18; fi12++ {
                            var v13 int = fi10
                            v7 = 1
                            v7 -= (^(v11))
                            v8 = 0
                            _ = fi12
                            _ = v13
                        }
                        switch (("" + "hello") + "") {
                        case "foo bar":
                            v14 := GenS2{F0: "hello"}
                            _ = v14
                        case "a":
                            var v15 uint64 = 1
                            _ = v15
                        case "\n\t":
                            v9 = append(v9, (fi10 / ((int(v11)) | 1)))
                        default:
                            fmt.Println(v8, v7)
                        }
                        _ = fi10
                        _ = v11
                    }
                    _ = v9
                }
                _ = v8
            } else {
                v16 := []int{-1, -1, -9223372036854775808, -115}
                v7 += 519
                v16 = append(v16, -888)
                _ = v16
            }
            v7 = (v7 ^ (v7 >> ((4294967295) & 63)))
            v17 := GenS3{F0: -32768, F1: 151, F2: 9223372036854775807}
            _ = v17
        }
        v7 = v7
        _ = v7
    }
    var v18 bool = false
    fmt.Println(v18)
    _ = v18
    return 1
}

func main() {
    var v1 uint8 = 1
    v2 := make(map[string]bool)
    var v3 int16 = 1
    var v4 uint = 689
    var v19 string = "a"
    fmt.Println(v1)
    switch (v19 + v19) {
    case "hello":
        fmt.Println(v19, v4)
    case "foo bar":
        v4 = (v4 | v4)
    case "a":
        v4 -= (v4 ^ (v4 % ((v4) | 1)))
    default:
        v20 := []int{71}
        _ = v20
    }
    fmt.Println(v19, v1, v19)
    v21 := GenS1{F0: -47, F1: 4294967295, F2: v4}
    fmt.Println(v1)
    var v22 int16 = 42
    v2[""] = false
    if (((1e10 / ((3.14159 * 0.0) - 3.14159)) != float32(0.0)) && (!((v1 << ((920) & 63)) == v1))) {
        if false {
            v1 = genFn1()
        } else {
            fmt.Println(v19)
        }
    } else {
        v23 := GenS2{F0: (v19 + v19)}
        v19 = v19
        _ = v23
    }
    for fi24 := 0; fi24 < 2; fi24++ {
        switch (-(fi24)) {
        case 45:
            var v25 int = fi24
            _ = v25
        case 128:
            var v26 int = 2147483647
            _ = v26
        case -1:
            v1 -= 128
        default:
            v27 := GenS4{F0: 130, F1: 42}
            _ = v27
        }
        _ = fi24
    }
    if (1 != 65535) {
        var v28 float32 = 1e10
        fmt.Println(v19)
        _ = v28
    } else {
        var v29 int64 = 9223372036854775807
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
