package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint8; F1 uint16; F2 string }
func genFn0(p0 int64, p1 bool) bool {
    fmt.Println(p1, p1)
    p1 = p1
    v4 := []int{}
    if (!(!(42 >= 751))) {
        var v5 float64 = 0.1
        _ = v5
    } else {
        _ = len(v4)
        v6 := GenS1{F0: 0, F1: (("a" + "\n\t") + "")}
        _ = v6
    }
    p0 -= p0
    _ = p0
    _ = p1
    _ = v4
    return (false || (!(p1 || p1)))
}

func main() {
    var v1 int64 = 454
    var v2 bool = false
    v3 := make(map[uint32]uint32)
    var v7 int64 = 484
    fmt.Println(v2, v1)
    switch 128 {
    case -2147483648:
        v8 := GenS2{F0: (("\n\t" + (("" + ((((("" + ("a" + ("\n\t" + ("a" + ((("\n\t" + "a") + "a") + (((("a" + "") + "foo bar") + "foo bar") + (("foo bar" + ("\n\t" + ("foo bar" + ""))) + "hello"))))))) + ("" + "")) + "a") + "foo bar") + ("\n\t" + "a"))) + "a")) + "foo bar"), F1: 374}
        _ = v8
    case 9223372036854775807:
        var v9 uint64 = 0
        _ = v9
    default:
        if ((1e10 >= 0.0) && true) {
            fmt.Println(v1, v2, v2)
        }
    }
    var v10 uint32 = 0
    v11 := GenS3{F0: (^(v10))}
    v3[v10] = 42
    for fi12 := 0; fi12 < 3; fi12++ {
        delete(v3, (v10 & v10))
        _ = fi12
    }
    for fi13 := 0; fi13 < 14; fi13++ {
        v14 := GenS2{F0: "foo bar", F1: 4294967295}
        v15 := GenS4{F0: 128, F1: 534, F2: "a"}
        _ = fi13
        _ = v14
        _ = v15
    }
    switch "foo bar" {
    case "":
        v3[v10] = 42
    case "\n\t":
        var v16 uint32 = 809
        _ = v16
    }
    var v17 int8 = -122
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v7
}
