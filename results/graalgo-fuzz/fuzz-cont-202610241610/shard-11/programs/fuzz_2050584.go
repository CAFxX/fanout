package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8; F1 int32; F2 uint64 }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 bool; F1 bool; F2 string }
type GenS7 struct { F0 uint16 }
func genFn0(p0 int8, p1 uint64) bool {
    for fi7 := 0; fi7 < 10; fi7++ {
        p1 += p1
        _ = fi7
    }
    var v8 int64 = -67
    var v9 int16 = 0
    v10 := GenS1{F0: (!((p1 ^ p1) < p1))}
    _ = p0
    _ = p1
    _ = v10
    _ = v8
    _ = v9
    return false
}

func main() {
    v1 := make(map[string]bool)
    var v2 int64 = 1
    var v3 uint = 120
    v4 := make(map[int]string)
    v5 := make(map[string]bool)
    var v6 uint32 = 2147483648
    switch "foo bar" {
    case "hello":
        var v11 uint8 = 2
        _ = v11
    default:
        v12 := GenS2{F0: 255, F1: 2147483647, F2: 936}
        _ = v12
    }
    v13 := GenS1{F0: (-2.25 > ((-2.25 / float32(-2.25)) / 1e10))}
    var v14 int = 337
    if (((("\n\t" <= "foo bar") && ((!false) || (0.0 < 0.1))) || (!(1 == -555))) || false) {
        v15 := GenS3{F0: 0, F1: 4294967295}
        v16 := GenS4{F0: 18446744073709551615}
        _ = len(v4)
        _ = v15
        _ = v16
    }
    v6 -= 710
    v17 := []int{}
    v14 += -452
    var v18 uint16 = 2
    delete(v5, "foo bar")
    for fi19 := 0; fi19 < 2; fi19++ {
        var v20 int16 = 1
        var v21 int8 = 0
        for fi22 := 0; fi22 < 8; fi22++ {
            if len(v17) >= 2 {
                v17 = v17[1:len(v17)]
            }
            v23 := GenS5{F0: 214}
            _ = fi22
            _ = v23
        }
        _ = fi19
        _ = v20
        _ = v21
    }
    if true {
        switch "a" {
        case "a":
            var v24 int16 = -32768
            _ = v24
        case "":
            v25 := GenS6{F0: true, F1: false, F2: ""}
            _ = v25
        case "hello":
            fmt.Println(v6)
        }
        var v26 string = "a"
        _ = v26
    } else {
        v27 := GenS6{F0: ((-700 > (^(v14))) && (!(-2147483648 < 1))), F1: true, F2: ""}
        v28 := GenS4{F0: 0}
        v29 := GenS7{F0: v18}
        _ = v27
        _ = v28
        _ = v29
    }
    for fi30 := 0; fi30 < 16; fi30++ {
        v3 = 18446744073709551615
        var v31 uint = (v3 % ((v3) | 1))
        v6 = 2147483648
        _ = fi30
        _ = v31
    }
    if len(v17) >= 2 {
        v17 = v17[1:len(v17)]
    }
    for fi32 := 0; fi32 < 4; fi32++ {
        v14 -= (v14 % (((-(fi32))) | 1))
        _ = fi32
    }
    v6 = 436
    v6 = uint32(v18)
    var v33 uint8 = 42
    var v34 uint32 = (v6 ^ v6)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
