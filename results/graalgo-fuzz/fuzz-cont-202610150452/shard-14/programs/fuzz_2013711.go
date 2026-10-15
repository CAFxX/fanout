package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32 }
func genFn0(p0 uint8, p1 int64) int16 {
    v7 := []int{}
    v8 := GenS1{F0: "hello"}
    var v9 float32 = 1e10
    v9 = float32(1.5)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 1000
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint16 = 1
    var v3 int64 = 42
    var v4 string = (("\n\t" + "hello") + "foo bar")
    v5 := []uint32{}
    var v6 int8 = -23
    var v10 string = v4
    for fi11 := 0; fi11 < 10; fi11++ {
        v6 += int8(v3)
        var v12 uint8 = 42
        _ = fi11
        _ = v12
    }
    if (!true) {
        var v13 int8 = v6
        v14 := GenS2{F0: (!false)}
        _ = v13
        _ = v14
    } else {
        switch "hello" {
        case "\n\t":
            var v15 int8 = v6
            _ = v15
        }
    }
    v16 := GenS1{F0: v4}
    for fi17 := 0; fi17 < 16; fi17++ {
        if false {
            v18 := GenS3{F0: 42}
            _ = v18
        } else {
            v19 := GenS3{F0: 42}
            for fi20 := 0; fi20 < 11; fi20++ {
                var v21 uint64 = 0
                _ = fi20
                _ = v21
            }
            _ = v19
        }
        if ((fi17 - 295) >= fi17) {
            fmt.Println(fi17, v2)
            switch "" {
            case "\n\t":
                switch "a" {
                case "":
                    fmt.Println(v4)
                case "hello":
                    var v22 int64 = (v3 % ((v3) | 1))
                    _ = v22
                default:
                    var v23 string = "hello"
                    _ = v23
                }
            }
            v2 = 0
        } else {
            v24 := GenS3{F0: 4294967295}
            _ = v24
        }
        _ = fi17
    }
    v25 := GenS1{F0: (v4 + v10)}
    var v26 int8 = (v6 << ((0) & 63))
    fmt.Println(v6, v26)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
