package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 int64; F2 uint32 }
type GenS4 struct { F0 uint32 }
func genFn0(p0 uint32) int32 {
    p0 = p0
    switch (("" + "a") + "hello") {
    case "foo bar":
        for fi6 := 0; fi6 < 17; fi6++ {
            v7 := GenS1{F0: 255, F1: 624, F2: "\n\t"}
            _ = fi6
            _ = v7
        }
    case "hello":
        p0 = p0
    case "":
        fmt.Println(p0)
    default:
        switch 1 {
        case 1:
            v8 := []int{}
            _ = v8
        case -428:
            for fi9 := 0; fi9 < 10; fi9++ {
                fmt.Println(fi9, p0)
                _ = fi9
            }
        default:
            p0 = (^(p0))
        }
    }
    fmt.Println(p0)
    v10 := GenS1{F0: 20, F1: 1, F2: "hello"}
    p0 += 42
    _ = p0
    _ = v10
    return 0
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]uint32)
    var v3 int = -9223372036854775808
    var v4 uint16 = 924
    var v5 int = (v3 - (v3 >> ((18446744073709551615) & 63)))
    var v11 string = ("foo bar" + "\n\t")
    fmt.Println(v5)
    var v12 bool = true
    for fi13 := 0; fi13 < 10; fi13++ {
        if (228 != 0) {
            if v12 {
                v11 = v11
            } else {
                v3 = -1
                var v14 int8 = -26
                if v12 {
                    for fi15 := 0; fi15 < 3; fi15++ {
                        v16 := GenS1{F0: 128, F1: 2, F2: "foo bar"}
                        v17 := GenS1{F0: 42, F1: 1, F2: (v11 + v11)}
                        var v18 int64 = 9223372036854775807
                        var v19 bool = (v4 >= v4)
                        _ = fi15
                        _ = v16
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                }
                _ = v14
            }
            fmt.Println(v4)
            var v20 uint = uint(v4)
            _ = v20
        }
        fmt.Println(v5)
        _ = fi13
    }
    v21 := GenS2{F0: 4294967295}
    var v22 int16 = 1
    switch (v3 % ((9223372036854775807) | 1)) {
    case 4294967295:
        fmt.Println(v3, v3, v22)
    case 839:
        v4 = 113
    case -921:
        var v23 bool = (980 == 0)
        _ = v23
    default:
        switch "" {
        case "hello":
            v4 = (^(v4))
        default:
            var v24 int64 = -148
            _ = v24
        }
    }
    switch v11 {
    case "foo bar":
        if false {
            v25 := GenS3{F0: -709, F1: 214, F2: uint32(v4)}
            fmt.Println(v5)
            _ = v25
        } else {
            v26 := v1[42]
            _ = v26
        }
    case "":
        v27 := GenS3{F0: 468, F1: -9223372036854775808, F2: 2}
        _ = v27
    }
    var v28 int32 = 0
    v12 = (v12 || (v12 && false))
    v29 := GenS1{F0: 18, F1: 0, F2: "foo bar"}
    switch (((v11 + ("a" + "foo bar")) + (v11 + ("\n\t" + "foo bar"))) + v11) {
    case "hello":
        var v30 float64 = float64(1e10)
        _ = v30
    default:
        for fi31 := 0; fi31 < 7; fi31++ {
            v32 := GenS4{F0: 1}
            _ = fi31
            _ = v32
        }
    }
    v33 := GenS1{F0: 0, F1: 2, F2: (v11 + v11)}
    fmt.Println(v28, v22, v12)
    var v34 bool = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v21
    _ = v22
    _ = v28
    _ = v29
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
