package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 int64 }
type GenS4 struct { F0 string }
func genFn0(p0 uint64, p1 string, p2 int8) uint16 {
    var v7 int = -96
    v8 := GenS1{F0: p0}
    fmt.Println(p1, p0, p2)
    for fi9 := 0; fi9 < 18; fi9++ {
        p2 = 26
        v10 := []int{(fi9 % (((v7 + v7)) | 1)), (v7 >> ((42) & 63)), fi9}
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 6; fi11++ {
        v12 := GenS2{F0: 962}
        if true {
            v13 := GenS3{F0: -1, F1: 1}
            if false {
                for fi14 := 0; fi14 < 5; fi14++ {
                    switch "\n\t" {
                    case "a":
                        var v15 float64 = float64((3.14159 * float32(0.1)))
                        _ = v15
                    case "":
                        var v16 int = (v7 | fi11)
                        var v17 int16 = 32767
                        _ = v16
                        _ = v17
                    }
                    v7 = -398
                    var v18 uint32 = 2147483647
                    _ = fi14
                    _ = v18
                }
            } else {
                v19 := GenS1{F0: (p0 / ((p0) | 1))}
                p0 = p0
                var v20 uint64 = 1
                _ = v19
                _ = v20
            }
            _ = v13
        }
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 uint16 = 1
    v2 := make(map[uint32]bool)
    var v3 uint64 = 18446744073709551615
    var v4 int8 = 1
    var v5 int8 = v4
    v6 := make(map[int]int)
    var v21 uint32 = 0
    fmt.Println(v4, v3)
    delete(v6, 42)
    var v22 float64 = -2.25
    switch 32768 {
    case 9223372036854775807:
        var v23 string = ("a" + "foo bar")
        _ = v23
    }
    var v24 float64 = ((3.14159 + v22) + 1.5)
    for fi25 := 0; fi25 < 2; fi25++ {
        v2[v21] = (v3 < (^(v3)))
        _ = fi25
    }
    if true {
        for fi26 := 0; fi26 < 4; fi26++ {
            v1 -= (v1 + (v1 - (v1 ^ (v1 / ((v1) | 1)))))
            _ = fi26
        }
    } else {
        v3 = v3
    }
    for fi27 := 0; fi27 < 12; fi27++ {
        switch "foo bar" {
        case "a":
            var v28 uint32 = (v21 << ((65535) & 63))
            _ = v28
        case "hello":
            var v29 float32 = 1.5
            _ = v29
        case "":
            for fi30 := 0; fi30 < 13; fi30++ {
                delete(v6, int(v1))
                _ = fi30
            }
        }
        _ = fi27
    }
    v1 = (^(v1))
    var v31 uint64 = v3
    var v32 uint = 2
    v33 := GenS4{F0: ("hello" + "")}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v24
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
