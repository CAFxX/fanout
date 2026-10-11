package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int; F1 int16; F2 int32 }
type GenS4 struct { F0 bool; F1 bool; F2 bool }
func genFn0() string {
    var v5 uint = 4294967295
    fmt.Println(v5)
    v5 = v5
    v5 = v5
    _ = v5
    return ("a" + ("foo bar" + ("a" + "\n\t")))
}
func genFn1(p0 uint16, p1 int8, p2 bool) uint16 {
    for fi6 := 0; fi6 < 4; fi6++ {
        fmt.Println(p0, fi6)
        _ = fi6
    }
    v7 := GenS1{F0: 543, F1: p2, F2: 65535}
    p0 += 169
    if (1 == -1) {
        fmt.Println(p0, p2, p0)
    }
    v8 := []int{-1}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (^(p0))
}

func main() {
    v1 := make(map[string]int)
    var v2 uint8 = 2
    var v3 uint = 18446744073709551615
    var v4 int16 = 32767
    for fi9 := 0; fi9 < 18; fi9++ {
        var v10 uint64 = 529
        if (!(!true)) {
            fmt.Println(v3, v3, v4)
            var v11 int32 = -1
            _ = v11
        }
        _ = fi9
        _ = v10
    }
    var v12 uint8 = (v2 / ((42) | 1))
    _ = len(v1)
    v13 := GenS1{F0: 42, F1: (1 <= 959), F2: genFn1(65535, 37, true)}
    v4 = v4
    v4 += (-(v4))
    v4 += (^(v4))
    if true {
        v14 := v1["\n\t"]
        switch ("\n\t" + "foo bar") {
        case "foo bar":
            var v15 uint8 = (^(v12))
            _ = v15
        case "":
            var v16 uint8 = (v12 << ((2) & 63))
            _ = v16
        case "hello":
            fmt.Println(v4)
        }
        _ = v14
    } else {
        v17 := GenS2{F0: 1}
        if (true || (true || true)) {
            v18 := v1["a"]
            _ = v18
        }
        if false {
            for fi19 := 0; fi19 < 5; fi19++ {
                delete(v1, "a")
                if false {
                    for fi20 := 0; fi20 < 19; fi20++ {
                        switch (fi20 ^ fi20) {
                        case 32768:
                            delete(v1, ("\n\t" + "\n\t"))
                        case 1:
                            v12 -= 255
                        default:
                            fmt.Println(fi19, v12, v12)
                        }
                        _ = fi20
                    }
                    if (("\n\t" + "\n\t") <= ("" + ("" + "a"))) {
                        var v21 int8 = -1
                        _ = v21
                    }
                    v22 := GenS3{F0: 9223372036854775807, F1: (-(v4)), F2: 0}
                    _ = v22
                }
                v23 := GenS4{F0: true, F1: false, F2: true}
                _ = fi19
                _ = v23
            }
            if (689 > 255) {
                var v24 uint = v3
                fmt.Println(v12, v4, v3)
                var v25 int = 127
                _ = v24
                _ = v25
            }
        }
        _ = v17
    }
    v2 += (v2 - (^(v2)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
