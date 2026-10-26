package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }

func main() {
    var v1 int8 = 1
    var v2 int32 = -1
    v3 := make(map[int]bool)
    v4 := []int{}
    v5 := []uint32{2147483647, 65535, 0, 4294967295}
    var v6 uint16 = 274
    _ = len(v3)
    v3[32767] = (!(42 != -1))
    fmt.Println(v1)
    var v7 uint16 = v6
    v8 := GenS1{F0: (-(v2)), F1: -9223372036854775808}
    v2 -= v2
    switch 127 {
    case 758:
        var v9 int64 = 0
        _ = v9
    case 32767:
        for fi10 := 0; fi10 < 2; fi10++ {
            v11 := GenS1{F0: (-(v2)), F1: 9223372036854775807}
            _ = fi10
            _ = v11
        }
    case 127:
        for fi12 := 0; fi12 < 16; fi12++ {
            var v13 bool = true
            if v13 {
                v7 = 25
                if v13 {
                    var v14 int16 = 42
                    if v13 {
                        for fi15 := 0; fi15 < 11; fi15++ {
                            var v16 int = (fi12 % ((fi12) | 1))
                            _ = fi15
                            _ = v16
                        }
                    } else {
                        v1 = (-(v1))
                        v14 = (-(v14))
                    }
                    v14 = 42
                    _ = v14
                }
                _ = len(v5)
            } else {
                if ((!(42 >= 294)) && (!v13)) {
                    var v17 uint8 = 1
                    _ = v17
                }
                v6 -= v6
            }
            fmt.Println(v6, v6)
            _ = fi12
            _ = v13
        }
    }
    v5 = append(v5, 614)
    v2 += 42
    var v18 uint64 = 1
    if false {
        v6 += 42
    } else {
        for fi19 := 0; fi19 < 10; fi19++ {
            v20 := GenS1{F0: (v2 * 1), F1: -9223372036854775808}
            _ = fi19
            _ = v20
        }
        for fi21 := 0; fi21 < 3; fi21++ {
            v2 = (-(v2))
            _ = fi21
        }
        switch (("hello" + "\n\t") + ("hello" + "foo bar")) {
        case "hello":
            for fi22 := 0; fi22 < 11; fi22++ {
                var v23 int = 42
                _ = fi22
                _ = v23
            }
        case "":
            var v24 int16 = int16(v18)
            _ = v24
        case "a":
            v18 += v18
        default:
            v25 := GenS1{F0: (v2 % ((v2) | 1)), F1: -1}
            _ = v25
        }
    }
    fmt.Println(v6, v6)
    fmt.Println(v7)
    _ = len(v4)
    var v26 int64 = -1
    v7 = (v7 / ((2) | 1))
    if (!false) {
        v1 += (v1 & v1)
    } else {
        v27 := GenS1{F0: v2, F1: v26}
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
