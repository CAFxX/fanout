package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 uint64 }
type GenS2 struct { F0 uint8; F1 int8 }
func genFn0(p0 uint) int64 {
    if false {
        v7 := GenS1{F0: 0, F1: 2147483647, F2: 1}
        _ = v7
    } else {
        switch ("\n\t" + "hello") {
        case "a":
            switch ("hello" + "") {
            case "a":
                switch (("" + "") + "hello") {
                case "foo bar":
                    if (!(!true)) {
                        fmt.Println(p0)
                        v8 := []int{-770, 128, 127, -9223372036854775808}
                        p0 = 2
                        _ = v8
                    } else {
                        switch "" {
                        case "hello":
                            fmt.Println(p0)
                        }
                        var v9 int64 = 42
                        _ = v9
                    }
                case "\n\t":
                    var v10 float64 = 0.0
                    _ = v10
                case "":
                    p0 += 2
                default:
                    p0 = (p0 % (((p0 & (p0 | (p0 & p0)))) | 1))
                }
            }
        default:
            if false {
                var v11 int8 = -91
                _ = v11
            } else {
                for fi12 := 0; fi12 < 18; fi12++ {
                    p0 = (p0 >> ((p0) & 63))
                    var v13 int32 = -1
                    _ = fi12
                    _ = v13
                }
                if true {
                    p0 -= p0
                    p0 = 4294967295
                    v14 := []int{57}
                    _ = v14
                }
                p0 = (p0 | p0)
            }
        }
        fmt.Println(p0)
    }
    var v15 uint = p0
    fmt.Println(p0, p0)
    if (316 <= 1) {
        p0 -= v15
    }
    v15 += (v15 ^ (v15 - v15))
    _ = p0
    _ = v15
    return -9223372036854775808
}

func main() {
    v1 := []int{-2147483648}
    v2 := make(map[uint32]string)
    var v3 uint16 = 0
    v4 := make(map[int]int)
    var v5 uint = 4294967295
    var v6 int = 32767
    v3 = v3
    var v16 int32 = 874
    v6 = (v6 / ((42) | 1))
    if true {
        v17 := GenS2{F0: 169, F1: 67}
        delete(v4, (^(v6)))
        _ = v17
    }
    v5 += 255
    v3 = v3
    _ = len(v2)
    fmt.Println(v5)
    v5 -= v5
    v18 := GenS1{F0: 1, F1: 240, F2: 0}
    var v19 string = "a"
    var v20 int = (v6 - v6)
    for fi21 := 0; fi21 < 2; fi21++ {
        v3 = (^(v3))
        _ = fi21
    }
    for fi22 := 0; fi22 < 11; fi22++ {
        var v23 int16 = -1
        _ = fi22
        _ = v23
    }
    if (v3 == 255) {
        v6 -= v6
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
