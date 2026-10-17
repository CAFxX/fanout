package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 int16; F2 bool }
type GenS3 struct { F0 int64 }
func genFn0(p0 int) uint8 {
    if false {
        var v5 uint32 = 2
        if false {
            switch 9223372036854775807 {
            case 151:
                var v6 uint64 = 2
                _ = v6
            default:
                v5 += 65535
            }
            v5 -= (v5 ^ 4294967295)
            var v7 float32 = 3.14159
            _ = v7
        } else {
            for fi8 := 0; fi8 < 14; fi8++ {
                if true {
                    var v9 uint32 = 1
                    _ = v9
                }
                var v10 bool = (-1 >= -1)
                _ = fi8
                _ = v10
            }
            v11 := []int{1}
            _ = v11
        }
        _ = v5
    }
    v12 := []int{}
    _ = p0
    _ = v12
    return 154
}

func main() {
    v1 := make(map[int]string)
    v2 := []int{-2147483648, 334, -9223372036854775808}
    v3 := make(map[uint32]string)
    v4 := make(map[int]bool)
    switch "" {
    case "foo bar":
        var v13 float32 = 0.0
        _ = v13
    case "":
        var v14 string = "foo bar"
        _ = v14
    case "\n\t":
        v1[32768] = ""
    default:
        v15 := GenS1{F0: -239}
        _ = v15
    }
    for fi16 := 0; fi16 < 20; fi16++ {
        var v17 float32 = (-2.25 / 0.1)
        for fi18 := 0; fi18 < 2; fi18++ {
            v19 := GenS1{F0: -162}
            if false {
                v17 += (v17 * (((1.5 / 3.14159) + v17) / (v17 * (v17 * (v17 * v17)))))
            } else {
                v20 := GenS2{F0: "hello", F1: 1000, F2: (!false)}
                if true {
                    fmt.Println(v17)
                }
                if (false || (0 == -9223372036854775808)) {
                    var v21 string = ""
                    _ = v21
                }
                _ = v20
            }
            v22 := GenS2{F0: "\n\t", F1: -439, F2: true}
            _ = fi18
            _ = v19
            _ = v22
        }
        _ = fi16
        _ = v17
    }
    v23 := GenS3{F0: 42}
    for fi24 := 0; fi24 < 4; fi24++ {
        v1[32767] = ("foo bar" + "\n\t")
        _ = fi24
    }
    var v25 bool = false
    v25 = false
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch 971 {
    case 989:
        v26 := GenS3{F0: -9223372036854775808}
        _ = v26
    }
    v27 := GenS1{F0: 9223372036854775807}
    for fi28 := 0; fi28 < 16; fi28++ {
        fmt.Println(fi28, v25)
        _ = fi28
    }
    for fi29 := 0; fi29 < 1; fi29++ {
        var v30 uint16 = 215
        v31 := GenS1{F0: 9223372036854775807}
        var v32 uint = 65535
        _ = fi29
        _ = v30
        _ = v31
        _ = v32
    }
    var v33 uint16 = 920
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v27
    _ = v3
    _ = v33
    _ = v4
}
