package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 string }
type GenS2 struct { F0 int8; F1 int8 }
func genFn0() string {
    var v6 uint64 = 566
    if true {
        v7 := GenS1{F0: (true || (((("foo bar" + "") + "hello") + "hello") < "")), F1: 149, F2: ((("foo bar" + "foo bar") + "foo bar") + "foo bar")}
        if true {
            fmt.Println(v6)
        }
        for fi8 := 0; fi8 < 14; fi8++ {
            for fi9 := 0; fi9 < 10; fi9++ {
                switch fi8 {
                case 9223372036854775807:
                    if (!(!(fi9 == -675))) {
                        v6 += (v6 + (v6 << ((4294967295) & 63)))
                        v6 -= v6
                    }
                case 2147483647:
                    var v10 bool = (42 >= 42)
                    _ = v10
                default:
                    var v11 uint = 128
                    _ = v11
                }
                _ = fi9
            }
            _ = fi8
        }
        _ = v7
    } else {
        v12 := []int{}
        if true {
            var v13 int8 = 32
            var v14 string = ("\n\t" + ((("\n\t" + ("foo bar" + "\n\t")) + ("" + ("\n\t" + ("hello" + "hello")))) + ("a" + ("\n\t" + ("" + ("" + ("" + "hello")))))))
            v15 := GenS1{F0: false, F1: 227, F2: "a"}
            _ = v13
            _ = v14
            _ = v15
        }
        _ = v12
    }
    _ = v6
    return ("" + "\n\t")
}
func genFn1(p0 uint64, p1 uint32, p2 bool) uint32 {
    var v16 bool = (2 < 0)
    switch -619 {
    case 127:
        var v17 int16 = 1
        _ = v17
    }
    v18 := []int{}
    fmt.Println(p2, v16)
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v18
    return 2147483648
}

func main() {
    var v1 int8 = 127
    var v2 uint32 = 4294967295
    var v3 int32 = 42
    var v4 float32 = 1.5
    v5 := make(map[uint32]int)
    var v19 int8 = (v1 * (v1 >> ((677) & 63)))
    v20 := GenS2{F0: v19, F1: v19}
    for fi21 := 0; fi21 < 4; fi21++ {
        v2 = (v2 >> ((4294967295) & 63))
        _ = fi21
    }
    v3 = v3
    fmt.Println(v1, v3, v2)
    for fi22 := 0; fi22 < 16; fi22++ {
        v1 -= v19
        switch "hello" {
        case "\n\t":
            var v23 string = ("" + "hello")
            _ = v23
        }
        var v24 uint32 = v2
        _ = fi22
        _ = v24
    }
    var v25 uint64 = 2
    v26 := []int{-9223372036854775808, int(v19)}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
