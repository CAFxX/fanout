package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 string }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 uint; F1 uint32 }

func main() {
    var v1 uint8 = 255
    var v2 int32 = 42
    v3 := make(map[string]bool)
    var v4 int64 = -9223372036854775808
    if true {
        if true {
            switch "hello" {
            case "":
                v1 = v1
            case "\n\t":
                _ = len(v3)
            default:
                var v5 uint64 = 18446744073709551615
                _ = v5
            }
            fmt.Println(v2, v1)
            _ = len(v3)
        }
        if true {
            v4 = (^(v4))
        }
        for fi6 := 0; fi6 < 8; fi6++ {
            v4 += v4
            _ = fi6
        }
    } else {
        switch -1 {
        case 315:
            v7 := []int{32767, -2147483648}
            _ = v7
        }
        if true {
            v4 = (v4 - v4)
            var v8 uint32 = 4294967295
            _ = v8
        } else {
            var v9 bool = false
            if v9 {
                v10 := GenS1{F0: 32768, F1: 0, F2: ("" + (("a" + "foo bar") + "a"))}
                _ = v10
            }
            _ = v9
        }
        fmt.Println(v4)
    }
    switch 42 {
    case 9223372036854775807:
        var v11 uint32 = 65535
        _ = v11
    default:
        if (false && (!((v4 + 0) >= v4))) {
            var v12 string = "foo bar"
            _ = v12
        } else {
            v13 := []int{4294967295}
            if true {
                if true {
                    v14 := GenS1{F0: 32767, F1: 552, F2: "\n\t"}
                    if false {
                        _ = len(v13)
                        v15 := GenS2{F0: 308, F1: (v2 / ((v2) | 1))}
                        _ = v15
                    }
                    _ = v14
                }
            } else {
                v13 = append(v13, 1)
            }
            var v16 int32 = v2
            _ = v13
            _ = v16
        }
    }
    _ = len(v3)
    var v17 int32 = v2
    v2 += -982
    if (true && false) {
        var v18 int16 = -32768
        for fi19 := 0; fi19 < 3; fi19++ {
            var v20 uint32 = 2147483648
            var v21 uint16 = 32767
            v22 := []int{fi19, 1, fi19}
            _ = fi19
            _ = v20
            _ = v21
            _ = v22
        }
        var v23 int16 = v18
        _ = v18
        _ = v23
    }
    if (32768 <= -9223372036854775808) {
        v1 = (v1 | (v1 % ((0) | 1)))
        var v24 uint16 = 32767
        for fi25 := 0; fi25 < 2; fi25++ {
            fmt.Println(v1, fi25)
            var v26 int8 = 100
            _ = fi25
            _ = v26
        }
        _ = v24
    } else {
        fmt.Println(v1, v17, v2)
        var v27 uint16 = 42
        _ = len(v3)
        _ = v27
    }
    v28 := GenS1{F0: -469, F1: 1, F2: "a"}
    for fi29 := 0; fi29 < 1; fi29++ {
        v30 := GenS3{F0: (fi29 * fi29), F1: (("\n\t" + "foo bar") + "hello")}
        _ = fi29
        _ = v30
    }
    if (((3.14159 + 3.14159) * 3.14159) >= 0.1) {
        v3["foo bar"] = ((true && false) && (642 >= 65535))
        v17 = (v17 + -2147483648)
    }
    v31 := GenS3{F0: int(v1), F1: ("\n\t" + ("" + ((((("foo bar" + "a") + ("" + "\n\t")) + ("foo bar" + "a")) + "\n\t") + ("\n\t" + "\n\t"))))}
    switch 0 {
    case 0:
        for fi32 := 0; fi32 < 9; fi32++ {
            _ = len(v3)
            _ = fi32
        }
    case -9223372036854775808:
        v17 += (-(v2))
    case -2147483648:
        v33 := GenS4{F0: 4294967295, F1: 524}
        _ = v33
    }
    fmt.Println(v4, v17, v17)
    var v34 int16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v28
    _ = v3
    _ = v31
    _ = v34
    _ = v4
}
