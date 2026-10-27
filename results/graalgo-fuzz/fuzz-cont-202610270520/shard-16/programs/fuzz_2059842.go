package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 int; F1 int16; F2 string }
func genFn0(p0 int16) string {
    var v7 float32 = -2.25
    fmt.Println(p0)
    _ = p0
    _ = v7
    return "a"
}

func main() {
    var v1 float32 = -2.25
    var v2 uint32 = 0
    var v3 string = (("a" + "foo bar") + (("hello" + ("" + "foo bar")) + "hello"))
    var v4 string = v3
    v5 := make(map[string]bool)
    v6 := make(map[uint32]string)
    var v8 uint64 = 2
    for fi9 := 0; fi9 < 8; fi9++ {
        var v10 int8 = 42
        _ = fi9
        _ = v10
    }
    switch v4 {
    case "foo bar":
        for fi11 := 0; fi11 < 11; fi11++ {
            var v12 bool = (!(4294967295 >= 42))
            var v13 uint8 = 94
            _ = fi11
            _ = v12
            _ = v13
        }
    default:
        var v14 uint16 = 605
        _ = v14
    }
    v15 := GenS1{F0: v8, F1: 2147483647}
    v3 = "\n\t"
    v16 := GenS1{F0: (v8 >> ((255) & 63)), F1: 1}
    for fi17 := 0; fi17 < 18; fi17++ {
        if (0 < -521) {
            fmt.Println(v3, fi17)
            switch (v4 + (v3 + (v3 + "a"))) {
            case "hello":
                fmt.Println(v3, v4)
            }
            v18 := GenS2{F0: 0, F1: 2}
            _ = v18
        }
        _ = fi17
    }
    if false {
        v19 := []int{32767, 4294967295, -9223372036854775808, -1}
        var v20 int32 = -1
        _ = v19
        _ = v20
    }
    delete(v6, v2)
    v21 := GenS3{F0: 777, F1: 1000, F2: (v4 + "foo bar")}
    if false {
        for fi22 := 0; fi22 < 3; fi22++ {
            if ((-1 != 42) || false) {
                v23 := GenS2{F0: 971, F1: 109}
                var v24 uint32 = v2
                _ = v23
                _ = v24
            } else {
                var v25 uint32 = 325
                v2 = v2
                _ = v25
            }
            _ = fi22
        }
    }
    v4 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
