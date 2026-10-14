package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 int8; F2 uint64 }
type GenS3 struct { F0 int16; F1 int; F2 uint64 }

func main() {
    var v1 bool = (!false)
    var v2 string = "\n\t"
    var v3 float32 = -2.25
    v4 := []string{"a", (v2 + "foo bar")}
    v5 := []int{}
    var v6 uint32 = 2147483648
    switch "a" {
    case "a":
        var v7 uint16 = 500
        _ = v7
    default:
        fmt.Println(v1)
    }
    v6 += (v6 | 941)
    switch ("" + v2) {
    case "hello":
        if false {
            v6 = 2
            if len(v4) > 0 {
                v4[0] = "a"
            }
        }
    case "a":
        var v8 uint32 = 42
        _ = v8
    case "\n\t":
        if ((v2 + v2) <= (v2 + (("a" + v2) + (v2 + v2)))) {
            v2 = v2
        }
    default:
        for fi9 := 0; fi9 < 10; fi9++ {
            v2 = v2
            if v1 {
                v10 := GenS1{F0: 0}
                v3 = v3
                _ = v10
            }
            for fi11 := 0; fi11 < 3; fi11++ {
                v12 := GenS2{F0: -1, F1: -1, F2: 2}
                if (true && ("foo bar" >= v2)) {
                    v13 := GenS2{F0: 127, F1: -1, F2: 1}
                    if v1 {
                        if v1 {
                            var v14 int8 = -24
                            fmt.Println(v1, fi11)
                            _ = v14
                        }
                        switch fi11 {
                        case -9223372036854775808:
                            v15 := GenS1{F0: 766}
                            _ = v15
                        case 2147483647:
                            var v16 bool = true
                            _ = v16
                        case 4294967295:
                            fmt.Println(v2, fi11, v3)
                        }
                    }
                    fmt.Println(v3)
                    _ = v13
                }
                var v17 int = fi9
                switch (("hello" + v2) + ((("a" + (v2 + ((v2 + "foo bar") + ("\n\t" + "")))) + v2) + v2)) {
                case "a":
                    if false {
                        if len(v5) > 0 {
                            v5[0] = 342
                        }
                        v18 := GenS1{F0: 1}
                        if v1 {
                            var v19 float64 = (0.0 + (1e10 + 1e10))
                            v4 = append(v4, ((v2 + ("" + v2)) + "a"))
                            fmt.Println(fi11, fi9)
                            _ = v19
                        } else {
                            v20 := GenS3{F0: 599, F1: v17, F2: 4294967295}
                            _ = v20
                        }
                        _ = v18
                    }
                }
                _ = fi11
                _ = v12
                _ = v17
            }
            _ = fi9
        }
    }
    if v1 {
        var v21 uint32 = (v6 - v6)
        var v22 int = 32767
        _ = v21
        _ = v22
    } else {
        v23 := GenS3{F0: 362, F1: 9223372036854775807, F2: 623}
        var v24 uint64 = 4294967295
        _ = v23
        _ = v24
    }
    switch 32767 {
    case 864:
        v4 = append(v4, (v2 + ("a" + ("\n\t" + v2))))
    case 0:
        var v25 float32 = ((0.0 * v3) + v3)
        _ = v25
    case 128:
        var v26 uint64 = 0
        _ = v26
    default:
        var v27 int32 = 42
        _ = v27
    }
    if (!true) {
        v3 += v3
    }
    v6 = 115
    fmt.Println(v3)
    switch v2 {
    case "":
        v5 = append(v5, 42)
    default:
        v6 = (v6 ^ (v6 + v6))
    }
    _ = len(v5)
    if v1 {
        fmt.Println(v3, v1)
        for fi28 := 0; fi28 < 18; fi28++ {
            v3 = v3
            _ = fi28
        }
        v29 := GenS3{F0: 199, F1: 32767, F2: 18446744073709551615}
        _ = v29
    }
    for fi30 := 0; fi30 < 14; fi30++ {
        for fi31 := 0; fi31 < 1; fi31++ {
            fmt.Println(fi31, fi30, v6)
            v32 := GenS3{F0: -32768, F1: (^(fi30)), F2: 1}
            _ = fi31
            _ = v32
        }
        _ = fi30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
