package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
func genFn0(p0 uint, p1 int, p2 uint32) int {
    p0 -= p0
    v5 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 623
}
func genFn1(p0 uint8, p1 int16) int32 {
    v6 := GenS1{F0: 0, F1: -944}
    var v7 int64 = -1
    switch ("a" + ("a" + "hello")) {
    case "foo bar":
        var v8 int16 = p1
        _ = v8
    case "":
        var v9 uint16 = 32767
        _ = v9
    case "\n\t":
        for fi10 := 0; fi10 < 19; fi10++ {
            var v11 uint = 65535
            switch "foo bar" {
            case "foo bar":
                fmt.Println(p1, p1, p1)
            }
            _ = fi10
            _ = v11
        }
    }
    p0 = p0
    switch ("a" + "foo bar") {
    case "hello":
        if (true || (65535 >= 1)) {
            for fi12 := 0; fi12 < 8; fi12++ {
                if true {
                    if (true || true) {
                        for fi13 := 0; fi13 < 19; fi13++ {
                            v7 -= v7
                            fmt.Println(fi13, fi12, fi12)
                            v14 := []int{-2147483648, fi13, fi13}
                            _ = fi13
                            _ = v14
                        }
                        var v15 string = "hello"
                        var v16 uint = 255
                        _ = v15
                        _ = v16
                    }
                    var v17 string = ""
                    v17 = v17
                    _ = v17
                }
                _ = fi12
            }
        }
    case "\n\t":
        v18 := []int{127}
        _ = v18
    case "foo bar":
        v19 := []int{2147483647}
        _ = v19
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 2147483647
}

func main() {
    var v1 uint32 = 1
    var v2 float32 = 3.14159
    var v3 uint64 = 4294967295
    v4 := make(map[uint32]string)
    var v20 uint8 = 127
    v3 += v3
    v21 := GenS1{F0: 65535, F1: -32768}
    v22 := []int{}
    for fi23 := 0; fi23 < 10; fi23++ {
        switch (fi23 | fi23) {
        case -413:
            var v24 uint8 = (v20 << ((uint(v1)) & 63))
            _ = v24
        case 42:
            _ = len(v22)
        case 9223372036854775807:
            v3 += v3
        }
        switch fi23 {
        case 32767:
            fmt.Println(v2, v2)
        }
        _ = fi23
    }
    if len(v22) > 0 {
        v22[0] = 701
    }
    v20 = 2
    for fi25 := 0; fi25 < 13; fi25++ {
        v22 = append(v22, int(v3))
        v26 := GenS1{F0: uint16(v3), F1: 0}
        if true {
            v27 := v4[v1]
            if len(v22) > 0 {
                v22[0] = (fi25 << ((18446744073709551615) & 63))
            }
            _ = len(v22)
            _ = v27
        }
        _ = fi25
        _ = v26
    }
    v20 -= (v20 - (^(v20)))
    if len(v22) > 0 {
        v22[0] = genFn0(42, -9223372036854775808, 884)
    }
    _ = len(v22)
    fmt.Println(v2, v20, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
