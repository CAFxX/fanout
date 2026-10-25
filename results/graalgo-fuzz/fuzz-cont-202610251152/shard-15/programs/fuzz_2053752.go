package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint32 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 uint }
func genFn0(p0 uint32, p1 int8) uint {
    var v7 uint16 = 476
    v7 += 330
    var v8 uint32 = (p0 + (p0 & 42))
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 2
}
func genFn1() uint32 {
    var v9 bool = (203 > 0)
    if v9 {
        v9 = (3.14159 >= 1e10)
    } else {
        v10 := []int{-146, 127, 32767, 32768}
        fmt.Println(v9)
        _ = v10
    }
    _ = v9
    return 42
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int16 = 32767
    v3 := []string{("" + (("hello" + "hello") + "")), "a"}
    var v4 float64 = 0.0
    v5 := []string{"", "foo bar", "hello", "\n\t"}
    var v6 bool = (false || (!(true && ("" == "a"))))
    var v11 uint64 = 95
    v6 = (32767 != 408)
    v12 := v1[(("" + "hello") + "")]
    var v13 bool = (2 > 2)
    v12 = 2147483648
    switch 1 {
    case -817:
        if (!("a" <= ("hello" + ("" + "a")))) {
            fmt.Println(v6, v13)
            v14 := GenS1{F0: v2, F1: 103, F2: v12}
            v15 := v1["foo bar"]
            _ = v14
            _ = v15
        }
    default:
        var v16 int8 = -128
        _ = v16
    }
    switch 9223372036854775807 {
    case 341:
        if true {
            v17 := v1[(("\n\t" + "hello") + "foo bar")]
            _ = v17
        } else {
            v18 := GenS2{F0: 139, F1: genFn0(v12, 1)}
            for fi19 := 0; fi19 < 3; fi19++ {
                v20 := GenS3{F0: 422}
                switch -2147483648 {
                case -892:
                    for fi21 := 0; fi21 < 5; fi21++ {
                        if len(v5) > 0 {
                            v5[0] = ""
                        }
                        switch "foo bar" {
                        case "":
                            v2 = v2
                        case "hello":
                            var v22 int64 = 9223372036854775807
                            _ = v22
                        case "\n\t":
                            v1[""] = uint32(v11)
                        }
                        switch "hello" {
                        case "\n\t":
                            v4 += v4
                        }
                        _ = fi21
                    }
                case 9223372036854775807:
                    var v23 uint64 = (^(v11))
                    _ = v23
                case 0:
                    if v13 {
                        var v24 int64 = 0
                        _ = v24
                    }
                }
                _ = fi19
                _ = v20
            }
            var v25 float64 = 3.14159
            _ = v18
            _ = v25
        }
    }
    for fi26 := 0; fi26 < 6; fi26++ {
        var v27 uint8 = 127
        _ = fi26
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
