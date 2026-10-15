package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string }
func genFn0(p0 int32, p1 int) uint8 {
    fmt.Println(p1, p1)
    switch (p1 + (p1 + (p1 - p1))) {
    case 9223372036854775807:
        fmt.Println(p0, p0)
    case 547:
        if true {
            p0 -= (p0 / ((p0) | 1))
        }
    case 42:
        switch p1 {
        case 1:
            var v6 uint64 = 69
            _ = v6
        default:
            v7 := []int{}
            _ = v7
        }
    }
    _ = p0
    _ = p1
    return 245
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []string{}
    v3 := []string{("a" + ("hello" + ("a" + "")))}
    var v4 string = ""
    v5 := make(map[int]int)
    v4 = ((v4 + v4) + "hello")
    var v8 int = 4294967295
    for fi9 := 0; fi9 < 11; fi9++ {
        switch v4 {
        case "a":
            if true {
                v10 := GenS1{F0: 856}
                _ = v10
            }
        case "":
            fmt.Println(fi9)
        default:
            if (-133 != 1000) {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                var v11 int64 = 9223372036854775807
                v12 := GenS2{F0: (v4 + (v4 + "a"))}
                _ = v11
                _ = v12
            } else {
                var v13 float32 = -2.25
                v14 := GenS1{F0: 822}
                for fi15 := 0; fi15 < 16; fi15++ {
                    v16 := GenS2{F0: ((v4 + (("\n\t" + (v4 + "\n\t")) + v4)) + v4)}
                    _ = fi15
                    _ = v16
                }
                _ = v13
                _ = v14
            }
        }
        _ = len(v2)
        if (!((true && ((^(v8)) != v8)) && false)) {
            v8 = 128
        }
        _ = fi9
    }
    switch "\n\t" {
    case "\n\t":
        v5[(v8 / ((v8) | 1))] = v8
    case "a":
        v4 = (v4 + v4)
    }
    _ = len(v3)
    var v17 uint64 = 155
    var v18 uint32 = 988
    for fi19 := 0; fi19 < 8; fi19++ {
        switch (v8 ^ 127) {
        case 128:
            fmt.Println(v8, v18)
        case -2147483648:
            v1[(v18 & v18)] = v4
        default:
            v8 = (-(v8))
        }
        _ = fi19
    }
    switch ((((((v4 + "a") + (v4 + v4)) + "hello") + v4) + (v4 + "hello")) + ((((((((v4 + v4) + v4) + (v4 + "")) + (v4 + "foo bar")) + v4) + v4) + "hello") + v4)) {
    case "hello":
        v17 = 4294967295
    default:
        var v20 string = ""
        _ = v20
    }
    if len(v2) > 0 {
        v2[0] = (v4 + v4)
    }
    var v21 int8 = 42
    v22 := GenS2{F0: v4}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
