package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 uint8 }
type GenS2 struct { F0 string; F1 int32; F2 bool }
func genFn0() bool {
    v7 := []int{32767, 128, 128, 32768}
    var v8 int64 = 1
    for fi9 := 0; fi9 < 13; fi9++ {
        v10 := GenS1{F0: 2147483647, F1: "", F2: 127}
        switch "a" {
        case "hello":
            switch (^(fi9)) {
            case 128:
                if (9223372036854775807 <= (v8 << ((42) & 63))) {
                    fmt.Println(v8)
                    for fi11 := 0; fi11 < 20; fi11++ {
                        switch ("a" + "hello") {
                        case "":
                            if len(v7) > 0 {
                                v7[0] = fi9
                            }
                        case "\n\t":
                            v8 = (^(v8))
                        }
                        v7 = append(v7, (fi11 - (fi11 + (fi9 >> ((1) & 63)))))
                        v12 := GenS2{F0: (("hello" + ("" + "\n\t")) + ""), F1: -265, F2: (42 >= 2147483647)}
                        _ = fi11
                        _ = v12
                    }
                } else {
                    v8 = (v8 % (((v8 & (v8 % ((v8) | 1)))) | 1))
                    v8 = v8
                }
            case 127:
                var v13 int64 = -9223372036854775808
                _ = v13
            }
        case "\n\t":
            if false {
                v8 -= v8
                _ = len(v7)
            }
        default:
            fmt.Println(v8, v8)
        }
        switch 128 {
        case 128:
            var v14 int8 = 100
            _ = v14
        case 1:
            if true {
                for fi15 := 0; fi15 < 3; fi15++ {
                    fmt.Println(v8, fi15)
                    fmt.Println(fi15)
                    if len(v7) > 0 {
                        v7[0] = fi9
                    }
                    _ = fi15
                }
            }
        }
        _ = fi9
        _ = v10
    }
    var v16 uint16 = 1
    if (1.5 < ((1e10 - 0.1) + 1e10)) {
        v17 := GenS2{F0: ("\n\t" + ("foo bar" + ((("\n\t" + "") + "foo bar") + "hello"))), F1: -113, F2: false}
        var v18 float32 = (0.1 - 1.5)
        _ = v17
        _ = v18
    }
    _ = v16
    _ = v7
    _ = v8
    return true
}

func main() {
    var v1 uint32 = 42
    var v2 int64 = -1
    var v3 uint16 = 255
    var v4 float32 = (-2.25 + (1.5 + (1.5 * -2.25)))
    v5 := []int{}
    v6 := make(map[string]int)
    switch -1 {
    case -1:
        v19 := GenS1{F0: 42, F1: (("" + "foo bar") + "hello"), F2: 0}
        _ = v19
    case 9223372036854775807:
        if (1 == 1) {
            fmt.Println(v3)
            if len(v5) > 0 {
                v5[0] = -9223372036854775808
            }
            var v20 int32 = int32(v3)
            _ = v20
        }
    default:
        var v21 string = "a"
        _ = v21
    }
    for fi22 := 0; fi22 < 2; fi22++ {
        if len(v5) > 0 {
            v5[0] = fi22
        }
        _ = fi22
    }
    v2 = -1
    for fi23 := 0; fi23 < 19; fi23++ {
        switch ("hello" + "foo bar") {
        case "\n\t":
            v1 -= (v1 + (^(v1)))
        }
        _ = fi23
    }
    _ = len(v5)
    if true {
        var v24 int64 = (v2 << ((4294967295) & 63))
        _ = v24
    } else {
        v3 += 0
    }
    _ = len(v6)
    fmt.Println(v3, v3)
    if true {
        if (v3 > v3) {
            switch ((((("a" + "") + "") + (("hello" + ("hello" + "foo bar")) + ("\n\t" + ""))) + "") + "hello") {
            case "hello":
                v6[""] = 57
            }
        }
        var v25 uint = 0
        var v26 uint64 = 172
        _ = v25
        _ = v26
    }
    fmt.Println(v2)
    if true {
        switch "\n\t" {
        case "a":
            var v27 int16 = 32767
            _ = v27
        }
    }
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
