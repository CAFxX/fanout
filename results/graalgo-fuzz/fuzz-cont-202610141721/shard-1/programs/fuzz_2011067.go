package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 uint16; F1 uint; F2 int32 }
type GenS3 struct { F0 uint16; F1 int16 }
type GenS4 struct { F0 bool }
func genFn0(p0 uint64, p1 int64) int16 {
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 bool = true
        _ = fi7
        _ = v8
    }
    p0 += 1
    var v9 float32 = -2.25
    if (false || true) {
        v10 := []int{32767, -620}
        _ = v10
    } else {
        p0 += (p0 >> ((4294967295) & 63))
    }
    _ = p0
    _ = p1
    _ = v9
    return 32767
}

func main() {
    v1 := []string{("" + "\n\t"), ("hello" + (("a" + ("a" + "\n\t")) + "hello")), (("a" + "\n\t") + "hello")}
    var v2 uint64 = 18446744073709551615
    v3 := make(map[uint32]uint32)
    v4 := []uint32{346}
    v5 := make(map[string]int)
    v6 := []string{("" + (("hello" + "a") + (("hello" + ("hello" + (("a" + "foo bar") + ("\n\t" + ("hello" + ("\n\t" + ("foo bar" + "a"))))))) + "foo bar"))), "a", ("" + (("a" + ("foo bar" + ((("" + "") + ("\n\t" + "hello")) + "a"))) + ("a" + "a"))), ("a" + "foo bar")}
    v11 := GenS1{F0: 4294967295, F1: 126}
    switch 127 {
    case -1:
        if (!true) {
            v12 := GenS2{F0: 32767, F1: 850, F2: 42}
            var v13 int64 = 725
            _ = v12
            _ = v13
        }
    case 522:
        if true {
            if false {
                v14 := GenS3{F0: 422, F1: -30}
                fmt.Println(v2)
                _ = v14
            } else {
                v2 -= 1
                var v15 int16 = 32767
                _ = v15
            }
        }
    default:
        fmt.Println(v2)
    }
    _ = len(v5)
    v16 := GenS4{F0: false}
    var v17 bool = (false || ((!false) || (42 == 2147483647)))
    var v18 uint32 = 2147483647
    fmt.Println(v17, v17)
    switch "hello" {
    case "":
        v18 = v18
    case "\n\t":
        fmt.Println(v17, v2)
    case "foo bar":
        var v19 int32 = -1
        _ = v19
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
