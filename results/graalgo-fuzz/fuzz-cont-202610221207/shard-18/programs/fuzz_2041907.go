package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16 }
func genFn0() int16 {
    v6 := []int{-1}
    var v7 uint32 = 2147483647
    switch -2147483648 {
    case 128:
        v7 += v7
    case -162:
        v8 := GenS1{F0: 857, F1: 42, F2: (("\n\t" + "hello") + ("hello" + "foo bar"))}
        _ = v8
    case 1:
        var v9 uint = 195
        _ = v9
    default:
        if (!true) {
            v7 -= v7
            v10 := GenS2{F0: -759}
            _ = v10
        }
    }
    v11 := GenS2{F0: 0}
    _ = v11
    _ = v6
    _ = v7
    return 1000
}

func main() {
    var v1 uint32 = 417
    v2 := make(map[int]int)
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]uint32)
    var v5 uint64 = 18446744073709551615
    v12 := GenS2{F0: genFn0()}
    var v13 int32 = 1
    fmt.Println(v5, v13)
    switch "\n\t" {
    case "":
        v5 = 0
    case "\n\t":
        if (1 != 42) {
            v13 = v13
        } else {
            v1 = 0
        }
    case "hello":
        v13 -= v13
    }
    for fi14 := 0; fi14 < 10; fi14++ {
        fmt.Println(v13, v1, v1)
        switch ("a" + "\n\t") {
        case "hello":
            var v15 int32 = (v13 | v13)
            _ = v15
        }
        _ = fi14
    }
    if (100 != -1) {
        switch ("hello" + "hello") {
        case "foo bar":
            if (((128 > 15) || ((v5 ^ v5) > v5)) || false) {
                var v16 int8 = 73
                _ = v16
            } else {
                v5 = 169
                v2[1] = 1
                fmt.Println(v1, v5)
            }
        }
    }
    for fi17 := 0; fi17 < 18; fi17++ {
        v18 := GenS3{F0: 726}
        var v19 uint32 = (v1 * (v1 << ((0) & 63)))
        _ = fi17
        _ = v18
        _ = v19
    }
    v20 := GenS3{F0: 618}
    v13 -= v13
    v5 = (v5 ^ (v5 - v5))
    v13 = v13
    for fi21 := 0; fi21 < 14; fi21++ {
        switch ("\n\t" + (("foo bar" + ("a" + "a")) + "foo bar")) {
        case "foo bar":
            v22 := []int{(fi21 % ((0) | 1))}
            _ = v22
        case "a":
            var v23 uint64 = v5
            _ = v23
        }
        v24 := []int{fi21, 967, 1}
        if len(v24) >= 2 {
            v24 = v24[1:len(v24)]
        }
        _ = fi21
        _ = v24
    }
    delete(v3, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
