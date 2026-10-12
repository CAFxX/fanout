package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 uint; F1 string; F2 int8 }
func genFn0() int32 {
    switch ((("\n\t" + "") + ("hello" + "hello")) + "hello") {
    case "\n\t":
        v7 := []int{-186, -2147483648}
        _ = v7
    default:
        var v8 uint = 65535
        _ = v8
    }
    var v9 bool = false
    _ = v9
    return 42
}

func main() {
    v1 := []string{}
    var v2 uint64 = 42
    v3 := []int{int(v2)}
    v4 := make(map[string]string)
    v5 := make(map[int]uint32)
    v6 := []string{("hello" + ("a" + (("a" + ("" + "foo bar")) + ""))), "\n\t", ("foo bar" + "a")}
    var v10 uint32 = uint32(v2)
    var v11 uint64 = 627
    v12 := GenS1{F0: -856}
    if len(v3) > 0 {
        v3[0] = 32768
    }
    for fi13 := 0; fi13 < 18; fi13++ {
        if ((("a" + "\n\t") != "foo bar") || false) {
            switch "hello" {
            case "a":
                v14 := GenS2{F0: 2, F1: 1}
                _ = v14
            case "hello":
                v11 = v11
            }
            fmt.Println(v2)
            v4["\n\t"] = ("foo bar" + "a")
        } else {
            v10 += v10
            v10 += v10
            var v15 uint16 = 1
            _ = v15
        }
        var v16 float32 = (((float32((1.5 / float64((0.0 + ((-2.25 * (((0.0 + (1.5 - 1.5)) - -2.25) + 0.1)) + 1e10))))) * 0.1) - 3.14159) + -2.25)
        if len(v6) > 0 {
            v6[0] = "a"
        }
        _ = fi13
        _ = v16
    }
    v11 = v11
    v2 -= (v11 / ((v2) | 1))
    var v17 float64 = -2.25
    v18 := GenS3{F0: 42, F1: "foo bar", F2: 127}
    fmt.Println(v2)
    fmt.Println(v10, v11)
    _ = len(v6)
    if (-9223372036854775808 != -904) {
        if (1 <= 42) {
            v19 := GenS3{F0: 30, F1: "foo bar", F2: -128}
            if true {
                fmt.Println(v2, v11)
                var v20 int32 = 2147483647
                _ = len(v4)
                _ = v20
            } else {
                fmt.Println(v17, v2)
                if false {
                    v17 = v17
                    v11 -= v11
                    v3 = append(v3, 538)
                }
                if (127 <= 12) {
                    var v21 uint32 = v10
                    var v22 float64 = 0.1
                    _ = v21
                    _ = v22
                }
            }
            switch -9223372036854775808 {
            case -2147483648:
                v23 := GenS1{F0: 0}
                _ = v23
            case 0:
                delete(v5, 1)
            case 1:
                var v24 int16 = 852
                _ = v24
            }
            _ = v19
        }
    }
    if false {
        v25 := GenS2{F0: 0, F1: 836}
        fmt.Println(v2, v17, v11)
        _ = v25
    } else {
        fmt.Println(v11, v17, v10)
        var v26 int32 = -1
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v26
    }
    fmt.Println(v17, v11, v17)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
