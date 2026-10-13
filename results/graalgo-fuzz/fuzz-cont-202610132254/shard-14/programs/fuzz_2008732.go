package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16 }

func main() {
    v1 := []string{}
    var v2 string = "hello"
    var v3 int64 = 1
    var v4 string = v2
    var v5 uint8 = 1
    v6 := make(map[int]bool)
    v2 = "\n\t"
    var v7 int32 = -230
    var v8 uint16 = 65535
    var v9 uint64 = 1
    for fi10 := 0; fi10 < 13; fi10++ {
        v5 = (v5 + (v5 & 127))
        _ = fi10
    }
    var v11 float64 = 3.14159
    fmt.Println(v4)
    for fi12 := 0; fi12 < 9; fi12++ {
        v13 := GenS1{F0: "hello", F1: 863}
        var v14 uint32 = 2147483647
        v15 := GenS1{F0: v4, F1: v14}
        _ = fi12
        _ = v13
        _ = v14
        _ = v15
    }
    v3 = (v3 << ((18446744073709551615) & 63))
    var v16 string = "a"
    v16 = v2
    if false {
        var v17 bool = (true && true)
        delete(v6, -627)
        _ = v17
    }
    var v18 float32 = (-2.25 * (0.1 / float32(((float64(0.1) * v11) + float64((-2.25 + 0.0))))))
    switch ("foo bar" + ("\n\t" + "hello")) {
    case "a":
        for fi19 := 0; fi19 < 4; fi19++ {
            var v20 int16 = -1
            _ = fi19
            _ = v20
        }
    case "hello":
        for fi21 := 0; fi21 < 15; fi21++ {
            for fi22 := 0; fi22 < 7; fi22++ {
                if false {
                    var v23 bool = true
                    var v24 int64 = (-(v3))
                    var v25 uint64 = 333
                    _ = v23
                    _ = v24
                    _ = v25
                } else {
                    var v26 float64 = v11
                    fmt.Println(v16)
                    _ = v26
                }
                v18 = ((1.5 / 1e10) / (v18 * v18))
                _ = fi22
            }
            v9 = (^(v9))
            _ = fi21
        }
    case "\n\t":
        v27 := GenS1{F0: "hello", F1: 2147483648}
        _ = v27
    default:
        v4 = v2
    }
    for fi28 := 0; fi28 < 9; fi28++ {
        var v29 int = -1
        _ = fi28
        _ = v29
    }
    if false {
        for fi30 := 0; fi30 < 15; fi30++ {
            fmt.Println(v18, v9, v4)
            delete(v6, fi30)
            v31 := GenS2{F0: fi30}
            _ = fi30
            _ = v31
        }
    }
    var v32 int64 = v3
    v33 := GenS3{F0: (v8 | v8)}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
