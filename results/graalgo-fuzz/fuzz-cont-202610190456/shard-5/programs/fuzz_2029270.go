package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 int8; F2 uint16 }

func main() {
    var v1 int64 = 1
    var v2 uint = 0
    v3 := make(map[uint32]uint32)
    var v4 string = ("" + "")
    for fi5 := 0; fi5 < 12; fi5++ {
        for fi6 := 0; fi6 < 17; fi6++ {
            fmt.Println(v2)
            var v7 int8 = 42
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    for fi8 := 0; fi8 < 10; fi8++ {
        var v9 int8 = 100
        var v10 float32 = 0.0
        var v11 int16 = int16(v1)
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    fmt.Println(v1)
    v12 := []int{-429}
    v13 := GenS1{F0: "\n\t"}
    switch v4 {
    case "foo bar":
        var v14 bool = (-2147483648 != 0)
        _ = v14
    }
    switch (v4 + v4) {
    case "foo bar":
        if (false || false) {
            if false {
                var v15 float64 = 3.14159
                if true {
                    if (!false) {
                        var v16 bool = false
                        v3[4294967295] = 42
                        var v17 uint32 = 4294967295
                        _ = v16
                        _ = v17
                    }
                    for fi18 := 0; fi18 < 16; fi18++ {
                        var v19 string = v4
                        v12 = append(v12, (-(fi18)))
                        _ = fi18
                        _ = v19
                    }
                }
                _ = v15
            }
        } else {
            switch v4 {
            case "\n\t":
                switch -1 {
                case -1:
                    fmt.Println(v2)
                case 2147483647:
                    var v20 string = v4
                    _ = v20
                case 9223372036854775807:
                    var v21 uint16 = 42
                    _ = v21
                }
            case "hello":
                if (false || true) {
                    var v22 int = 4294967295
                    var v23 bool = (!((v1 << (((^(v2))) & 63)) >= 42))
                    _ = v22
                    _ = v23
                } else {
                    fmt.Println(v4)
                }
            default:
                if ("\n\t" > "foo bar") {
                    var v24 int32 = -2147483648
                    v25 := GenS2{F0: -99, F1: 66, F2: 415}
                    var v26 uint64 = 0
                    _ = v24
                    _ = v25
                    _ = v26
                }
            }
            if false {
                var v27 int8 = 127
                _ = v27
            }
            v1 -= v1
        }
    }
    switch (v4 + ("hello" + ("hello" + (("a" + "") + "\n\t")))) {
    case "hello":
        v2 = (v2 * 303)
    case "foo bar":
        v1 = (v1 * 826)
    case "":
        var v28 float64 = (3.14159 * 1.5)
        _ = v28
    default:
        for fi29 := 0; fi29 < 15; fi29++ {
            var v30 uint32 = 2147483648
            v31 := GenS1{F0: "a"}
            v2 = v2
            _ = fi29
            _ = v30
            _ = v31
        }
    }
    if false {
        var v32 uint8 = 255
        v33 := GenS2{F0: 1, F1: 1, F2: 65535}
        if (!true) {
            fmt.Println(v32, v32)
            v34 := v3[65535]
            _ = v34
        } else {
            var v35 string = "foo bar"
            v32 -= v32
            v2 += v2
            _ = v35
        }
        _ = v32
        _ = v33
    } else {
        fmt.Println(v1)
        if (18446744073709551615 < 255) {
            var v36 int64 = v1
            _ = v36
        } else {
            v37 := GenS2{F0: -9223372036854775808, F1: 97, F2: 255}
            fmt.Println(v2, v4)
            _ = v37
        }
        fmt.Println(v4, v2, v1)
    }
    _ = len(v12)
    v38 := GenS2{F0: 9223372036854775807, F1: 103, F2: 42}
    var v39 int16 = 0
    var v40 int32 = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
}
