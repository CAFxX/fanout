package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 string }
func genFn0(p0 uint8, p1 uint8, p2 uint8) bool {
    if true {
        var v5 float32 = float32((3.14159 + 0.1))
        _ = v5
    }
    switch "hello" {
    case "hello":
        for fi6 := 0; fi6 < 17; fi6++ {
            var v7 int8 = 1
            v7 += 96
            v7 = v7
            _ = fi6
            _ = v7
        }
    case "foo bar":
        var v8 string = "foo bar"
        _ = v8
    case "\n\t":
        switch 956 {
        case -9223372036854775808:
            if false {
                v9 := GenS1{F0: 32767, F1: "a", F2: "foo bar"}
                _ = v9
            } else {
                if true {
                    if false {
                        v10 := GenS1{F0: 2147483647, F1: "foo bar", F2: "\n\t"}
                        _ = v10
                    }
                    switch ("hello" + ("hello" + "hello")) {
                    case "a":
                        p2 += (p0 / (((p0 ^ p0)) | 1))
                    }
                }
            }
        case 32767:
            v11 := GenS1{F0: 32768, F1: "hello", F2: "hello"}
            _ = v11
        case -555:
            fmt.Println(p2)
        }
    }
    for fi12 := 0; fi12 < 18; fi12++ {
        var v13 uint32 = 168
        switch (fi12 % ((-2147483648) | 1)) {
        case 2147483647:
            v14 := GenS1{F0: (fi12 | (fi12 - fi12)), F1: (("\n\t" + ("" + "a")) + "\n\t"), F2: ("foo bar" + ("" + "a"))}
            _ = v14
        case 0:
            var v15 string = (("a" + ("" + "a")) + "hello")
            _ = v15
        case 32768:
            if (1000 >= 0) {
                v16 := []int{(fi12 * 0), (^(fi12)), 42, int(v13)}
                if false {
                    var v17 int8 = 42
                    if (1 >= -1) {
                        if ((0.1 + 1.5) > (-2.25 + 1e10)) {
                            switch "\n\t" {
                            case "hello":
                                p1 -= (p1 / ((1) | 1))
                            case "":
                                v17 -= -127
                            }
                            var v18 uint8 = p2
                            _ = v18
                        }
                    }
                    _ = v17
                } else {
                    if len(v16) > 0 {
                        v16[0] = (fi12 ^ fi12)
                    }
                }
                _ = v16
            }
        }
        p0 = (p2 & (p0 & uint8(v13)))
        _ = fi12
        _ = v13
    }
    var v19 uint32 = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    return ((p1 - p1) == (p2 & p2))
}

func main() {
    var v1 int8 = -123
    v2 := []int{669, -528}
    var v3 int8 = 1
    v4 := make(map[string]int)
    var v20 int8 = (v1 ^ v1)
    _ = len(v4)
    for fi21 := 0; fi21 < 15; fi21++ {
        for fi22 := 0; fi22 < 7; fi22++ {
            if ((!true) || false) {
                var v23 float32 = float32((-2.25 / 1.0))
                _ = v23
            } else {
                v24 := GenS1{F0: (^(fi21)), F1: "a", F2: "foo bar"}
                _ = v24
            }
            _ = fi22
        }
        _ = fi21
    }
    var v25 int64 = 9223372036854775807
    var v26 int = -182
    v3 = -128
    switch v26 {
    case 128:
        var v27 int8 = (v1 ^ (v20 + (v3 & v1)))
        _ = v27
    case 4294967295:
        v28 := GenS1{F0: 4294967295, F1: "hello", F2: "foo bar"}
        _ = v28
    default:
        v29 := GenS1{F0: v26, F1: "a", F2: "\n\t"}
        _ = v29
    }
    switch (v26 * (v26 & v26)) {
    case 42:
        fmt.Println(v26, v20)
    case 491:
        if (0.0 < (3.14159 - 0.0)) {
            v25 = (v25 + (v25 - (v25 >> ((255) & 63))))
            if (true && ("a" == "hello")) {
                v25 = v25
                v3 = v3
            } else {
                for fi30 := 0; fi30 < 5; fi30++ {
                    v1 += int8(v25)
                    _ = fi30
                }
                switch (("hello" + "foo bar") + "\n\t") {
                case "a":
                    v31 := GenS1{F0: (v26 & v26), F1: "a", F2: "a"}
                    _ = v31
                }
                v4[""] = v26
            }
        }
    case -620:
        var v32 bool = true
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v20
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
