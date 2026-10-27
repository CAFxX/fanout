package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := make(map[uint32]bool)
    var v3 int64 = 1
    var v4 uint32 = 0
    v5 := make(map[int]string)
    v6 := make(map[string]string)
    v6["\n\t"] = ""
    switch ("foo bar" + "foo bar") {
    case "hello":
        switch 32768 {
        case -504:
            v1 -= v1
        default:
            v3 += v3
        }
    case "a":
        v7 := []int{171, 645, -252}
        _ = v7
    default:
        var v8 int = 2147483647
        _ = v8
    }
    var v9 int64 = v3
    var v10 uint16 = 26
    v9 = 637
    for fi11 := 0; fi11 < 6; fi11++ {
        fmt.Println(v1, fi11, v3)
        _ = fi11
    }
    for fi12 := 0; fi12 < 10; fi12++ {
        var v13 uint = 0
        if false {
            v3 = 1
        }
        v14 := []int{fi12, fi12}
        _ = fi12
        _ = v13
        _ = v14
    }
    v15 := GenS1{F0: 4294967295}
    v3 += 0
    switch "foo bar" {
    case "":
        var v16 int32 = 408
        _ = v16
    case "foo bar":
        for fi17 := 0; fi17 < 13; fi17++ {
            var v18 bool = (!(("" + ("\n\t" + (("a" + "") + "\n\t"))) >= ("" + "")))
            _ = fi17
            _ = v18
        }
    default:
        switch ("a" + ("a" + ((("a" + "foo bar") + "a") + "a"))) {
        case "\n\t":
            if (((127 > 1) && ((v10 + (v10 & v10)) == 2)) && true) {
                v19 := []int{9223372036854775807, 42}
                _ = v19
            }
        case "":
            for fi20 := 0; fi20 < 9; fi20++ {
                var v21 int8 = -109
                _ = fi20
                _ = v21
            }
        default:
            v22 := v5[1]
            _ = v22
        }
    }
    v23 := []int{482, -820}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
