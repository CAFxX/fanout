package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 int8; F1 int64; F2 int8 }
type GenS3 struct { F0 bool; F1 uint32 }
type GenS4 struct { F0 string; F1 int32; F2 int16 }

func main() {
    v1 := make(map[string]int)
    v2 := []uint32{2147483648, 0, 836}
    v3 := []uint32{4294967295, 0}
    v4 := make(map[string]bool)
    var v5 int = 2147483647
    for fi6 := 0; fi6 < 3; fi6++ {
        switch ("a" + (((("foo bar" + ("a" + "a")) + (("\n\t" + (("hello" + "hello") + ("hello" + ""))) + "\n\t")) + "\n\t") + "")) {
        case "\n\t":
            v7 := GenS1{F0: -9223372036854775808, F1: 0}
            _ = v7
        }
        var v8 uint64 = 1
        _ = fi6
        _ = v8
    }
    if (0 <= 18446744073709551615) {
        var v9 uint16 = 621
        for fi10 := 0; fi10 < 4; fi10++ {
            for fi11 := 0; fi11 < 7; fi11++ {
                v9 -= v9
                _ = fi11
            }
            fmt.Println(fi10, v5, v9)
            _ = fi10
        }
        v9 += v9
        _ = v9
    } else {
        v5 += v5
    }
    fmt.Println(v5)
    fmt.Println(v5)
    fmt.Println(v5)
    v2 = append(v2, 4294967295)
    var v12 uint32 = 553
    if false {
        v3 = append(v3, v12)
        v13 := GenS2{F0: 0, F1: 1, F2: -90}
        var v14 bool = ((true || ("foo bar" >= ((("hello" + "foo bar") + ((("foo bar" + ("hello" + "\n\t")) + "\n\t") + ("hello" + ("hello" + "hello")))) + "foo bar"))) && true)
        _ = v13
        _ = v14
    }
    var v15 uint8 = 43
    for fi16 := 0; fi16 < 12; fi16++ {
        if false {
            v17 := v1[(((("" + (("foo bar" + ("" + ("hello" + ("hello" + ("hello" + "hello"))))) + "a")) + "foo bar") + "hello") + "")]
            fmt.Println(v15, v17, v5)
            if (100 > -128) {
                var v18 int = fi16
                fmt.Println(v18, fi16, fi16)
                v17 = -449
                _ = v18
            } else {
                for fi19 := 0; fi19 < 8; fi19++ {
                    var v20 int8 = -103
                    _ = len(v2)
                    _ = fi19
                    _ = v20
                }
                v21 := GenS2{F0: -128, F1: 42, F2: 127}
                if true {
                    for fi22 := 0; fi22 < 16; fi22++ {
                        switch "foo bar" {
                        case "foo bar":
                            v17 += v5
                        case "\n\t":
                            v23 := GenS3{F0: (-340 > -733), F1: v12}
                            _ = v23
                        case "hello":
                            var v24 bool = true
                            _ = v24
                        default:
                            var v25 bool = ("\n\t" < "")
                            _ = v25
                        }
                        v26 := v4["hello"]
                        switch "hello" {
                        case "foo bar":
                            var v27 uint64 = 0
                            _ = v27
                        default:
                            fmt.Println(v26, v12)
                        }
                        _ = fi22
                        _ = v26
                    }
                    v28 := GenS2{F0: 25, F1: -9223372036854775808, F2: 127}
                    v3 = append(v3, (v12 | v12))
                    _ = v28
                }
                _ = v21
            }
            _ = v17
        }
        v29 := GenS4{F0: "a", F1: 42, F2: 984}
        _ = fi16
        _ = v29
    }
    v5 = (v5 + (v5 & 1))
    v1["hello"] = (v5 ^ v5)
    for fi30 := 0; fi30 < 2; fi30++ {
        v5 = v5
        var v31 uint = 42
        v32 := GenS4{F0: "\n\t", F1: 42, F2: 1000}
        _ = fi30
        _ = v31
        _ = v32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
