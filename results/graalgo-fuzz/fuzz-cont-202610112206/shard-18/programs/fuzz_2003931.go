package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 string }
type GenS2 struct { F0 int32; F1 int16; F2 int32 }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 int32; F1 int; F2 string }
func genFn0() uint {
    var v4 uint = 2
    for fi5 := 0; fi5 < 5; fi5++ {
        for fi6 := 0; fi6 < 17; fi6++ {
            if true {
                v4 = v4
            }
            switch "a" {
            case "foo bar":
                v7 := GenS1{F0: 32767, F1: "a", F2: ""}
                _ = v7
            default:
                v8 := []int{(-(fi5)), fi5, 0, 32767}
                _ = v8
            }
            v4 += (v4 / ((v4) | 1))
            _ = fi6
        }
        if (false && false) {
            var v9 uint64 = 18446744073709551615
            var v10 uint64 = 42
            _ = v10
            _ = v9
        }
        fmt.Println(fi5, v4)
        _ = fi5
    }
    if (((true && (1 <= 0)) && true) || false) {
        v11 := []int{4294967295}
        _ = v11
    } else {
        fmt.Println(v4)
    }
    _ = v4
    return (^(v4))
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]int)
    var v3 uint64 = 585
    var v12 int = 1
    delete(v2, "")
    var v13 int = v12
    switch ("foo bar" + ("" + "foo bar")) {
    case "a":
        switch 1 {
        case -256:
            var v14 int8 = 100
            _ = v14
        }
    case "foo bar":
        var v15 uint32 = 42
        _ = v15
    case "":
        v16 := GenS2{F0: 1, F1: 42, F2: 1}
        _ = v16
    }
    switch (v12 | (v13 / ((32768) | 1))) {
    case -5:
        fmt.Println(v13, v12)
    case 127:
        var v17 int32 = 171
        _ = v17
    case 4294967295:
        v18 := []int{v13, 32767, v12}
        _ = v18
    }
    switch "a" {
    case "\n\t":
        var v19 float64 = (0.0 - 0.0)
        _ = v19
    case "":
        v3 += 437
    default:
        switch 1 {
        case 0:
            v13 = v13
        case 32768:
            v12 -= v13
        default:
            var v20 float32 = 0.0
            _ = v20
        }
    }
    v21 := GenS3{F0: 9223372036854775807, F1: (v3 << ((162) & 63))}
    switch ("foo bar" + "foo bar") {
    case "\n\t":
        fmt.Println(v12, v12)
    case "foo bar":
        v1["a"] = 175
    case "hello":
        if (!((127 < -7) && (42 > -194))) {
            v3 -= (v3 ^ (v3 ^ 2))
            var v22 string = ""
            _ = v22
        } else {
            var v23 bool = true
            if (!(742 > 2)) {
                v24 := GenS4{F0: 42, F1: v13, F2: ""}
                if false {
                    delete(v1, (("\n\t" + "foo bar") + ""))
                    fmt.Println(v23)
                    v25 := GenS4{F0: -153, F1: v12, F2: ""}
                    _ = v25
                }
                v26 := GenS1{F0: -778, F1: ("hello" + "hello"), F2: "a"}
                _ = v24
                _ = v26
            } else {
                _ = len(v2)
            }
            _ = v23
        }
    }
    v27 := GenS4{F0: 899, F1: v13, F2: ("hello" + "a")}
    v13 -= v13
    switch "foo bar" {
    case "":
        v13 = (^(v13))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v27
    _ = v3
}
