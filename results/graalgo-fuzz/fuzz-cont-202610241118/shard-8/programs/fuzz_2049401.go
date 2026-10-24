package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 int, p1 uint64, p2 uint) int8 {
    switch (p0 + p0) {
    case 42:
        v4 := GenS1{F0: 2147483647}
        _ = v4
    case 1:
        var v5 int8 = 127
        _ = v5
    }
    v6 := []int{p0, (p0 ^ p0), (p0 + -992), 0}
    switch "a" {
    case "a":
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    fmt.Println(p2)
    var v7 int8 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return v7
}

func main() {
    var v1 uint8 = 128
    v2 := make(map[uint32]uint32)
    var v3 uint = 389
    for fi8 := 0; fi8 < 18; fi8++ {
        v1 -= v1
        v1 = (^(v1))
        _ = fi8
    }
    v1 = 0
    fmt.Println(v3, v3)
    var v9 uint = v3
    var v10 uint32 = 200
    if (("\n\t" + "hello") != ("" + ((("" + ("hello" + "a")) + "hello") + "a"))) {
        fmt.Println(v3, v1)
        for fi11 := 0; fi11 < 18; fi11++ {
            v12 := GenS1{F0: 32767}
            switch (^(fi11)) {
            case -409:
                if true {
                    fmt.Println(v3)
                }
            case -779:
                v13 := []int{}
                _ = v13
            case 884:
                switch ("\n\t" + "a") {
                case "a":
                    fmt.Println(v10, fi11)
                case "\n\t":
                    var v14 uint32 = v10
                    _ = v14
                case "foo bar":
                    fmt.Println(v10, fi11, fi11)
                default:
                    if (!false) {
                        var v15 uint = (v9 + v9)
                        v1 -= v1
                        var v16 int = 752
                        _ = v15
                        _ = v16
                    } else {
                        fmt.Println(v10)
                        v17 := GenS1{F0: (fi11 / (((fi11 + fi11)) | 1))}
                        _ = v17
                    }
                }
            }
            _ = fi11
            _ = v12
        }
        fmt.Println(v1, v9)
    } else {
        fmt.Println(v10, v9, v3)
        var v18 int = 0
        for fi19 := 0; fi19 < 4; fi19++ {
            var v20 float64 = 1.5
            _ = fi19
            _ = v20
        }
        _ = v18
    }
    for fi21 := 0; fi21 < 20; fi21++ {
        v22 := v2[v10]
        var v23 uint8 = uint8(v3)
        fmt.Println(v10)
        _ = fi21
        _ = v22
        _ = v23
    }
    v3 += (^(v3))
    v24 := GenS2{F0: 42, F1: 18446744073709551615}
    v25 := GenS3{F0: uint8(v10)}
    for fi26 := 0; fi26 < 20; fi26++ {
        v27 := []int{-693}
        _ = fi26
        _ = v27
    }
    switch ("\n\t" + "hello") {
    case "foo bar":
        fmt.Println(v1)
    default:
        switch "\n\t" {
        case "a":
            if false {
                v28 := GenS1{F0: 43}
                fmt.Println(v10, v3)
                _ = v28
            }
        }
    }
    switch "" {
    case "hello":
        for fi29 := 0; fi29 < 8; fi29++ {
            v10 = (v10 & 4294967295)
            v9 -= 18446744073709551615
            _ = fi29
        }
    }
    v30 := GenS2{F0: v1, F1: 925}
    var v31 uint32 = 2147483648
    var v32 float64 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v9
}
