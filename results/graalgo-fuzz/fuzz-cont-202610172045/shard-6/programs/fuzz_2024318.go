package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
func genFn0() bool {
    var v7 int8 = 1
    v8 := GenS1{F0: 0, F1: 171}
    fmt.Println(v7)
    var v9 int8 = v7
    var v10 int = 2147483647
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return true
}
func genFn1(p0 int32, p1 int32) int64 {
    v11 := GenS1{F0: 4294967295, F1: 0}
    var v12 int64 = 702
    for fi13 := 0; fi13 < 11; fi13++ {
        p1 -= -2147483648
        p1 = (^(p1))
        var v14 bool = true
        _ = fi13
        _ = v14
    }
    var v15 uint = 65535
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v15
    return v12
}

func main() {
    v1 := []uint32{1}
    v2 := make(map[int]bool)
    v3 := make(map[int]bool)
    v4 := make(map[string]uint32)
    v5 := make(map[uint32]string)
    v6 := []string{"foo bar", ("hello" + (("a" + "") + ("hello" + "hello"))), ((("\n\t" + ("hello" + "")) + "") + "\n\t")}
    var v16 uint64 = 471
    for fi17 := 0; fi17 < 2; fi17++ {
        fmt.Println(v16, fi17)
        _ = fi17
    }
    var v18 float64 = 3.14159
    v18 = (1.5 + v18)
    for fi19 := 0; fi19 < 9; fi19++ {
        var v20 int64 = 948
        if (false && false) {
            fmt.Println(v20, v18)
        } else {
            var v21 int8 = -128
            _ = len(v6)
            v22 := GenS1{F0: 65535, F1: 42}
            _ = v21
            _ = v22
        }
        var v23 float64 = 0.1
        _ = fi19
        _ = v20
        _ = v23
    }
    _ = len(v6)
    v24 := GenS1{F0: 4294967295, F1: 128}
    v16 -= (^(v16))
    v16 = (^(v16))
    fmt.Println(v18, v18)
    switch "" {
    case "hello":
        switch ((("hello" + "hello") + "\n\t") + "hello") {
        case "":
            var v25 int = 42
            _ = v25
        case "a":
            v26 := v3[2147483647]
            _ = v26
        case "\n\t":
            if (-1 > 32767) {
                fmt.Println(v18)
                var v27 uint = 14
                _ = len(v2)
                _ = v27
            } else {
                v18 += (((0.1 * v18) - v18) + (v18 + v18))
                switch int(v16) {
                case -2147483648:
                    for fi28 := 0; fi28 < 18; fi28++ {
                        var v29 bool = (3.14159 <= 1.5)
                        _ = fi28
                        _ = v29
                    }
                case 9223372036854775807:
                    switch (("foo bar" + "\n\t") + "\n\t") {
                    case "a":
                        var v30 float64 = float64(((0.1 - 0.1) * 0.0))
                        _ = v30
                    default:
                        for fi31 := 0; fi31 < 3; fi31++ {
                            _ = len(v2)
                            v18 += 1.5
                            _ = fi31
                        }
                    }
                default:
                    for fi32 := 0; fi32 < 10; fi32++ {
                        for fi33 := 0; fi33 < 15; fi33++ {
                            v34 := GenS1{F0: 704, F1: 2}
                            v18 -= float64(1.5)
                            _ = fi33
                            _ = v34
                        }
                        var v35 uint8 = 255
                        var v36 int32 = -1
                        _ = fi32
                        _ = v35
                        _ = v36
                    }
                }
                v6 = append(v6, "\n\t")
            }
        }
    case "foo bar":
        var v37 int32 = -1
        _ = v37
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v18
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
