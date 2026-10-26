package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint8; F1 int8 }
func genFn0(p0 int32) uint64 {
    v6 := []int{257, 32767, 9223372036854775807, 127}
    p0 = p0
    _ = p0
    _ = v6
    return 42
}

func main() {
    v1 := make(map[int]int)
    v2 := []string{"a", "a"}
    var v3 uint = 0
    v4 := make(map[uint32]string)
    v5 := []int{}
    if false {
        if false {
            var v7 float64 = 3.14159
            _ = v7
        }
        for fi8 := 0; fi8 < 18; fi8++ {
            if true {
                switch "hello" {
                case "a":
                    var v9 uint64 = 0
                    _ = v9
                case "foo bar":
                    v3 = (^(v3))
                }
            } else {
                v10 := GenS1{F0: genFn0(2147483647), F1: 44}
                for fi11 := 0; fi11 < 6; fi11++ {
                    var v12 uint32 = 2147483648
                    _ = fi11
                    _ = v12
                }
                _ = v10
            }
            if true {
                for fi13 := 0; fi13 < 11; fi13++ {
                    v14 := GenS1{F0: 18446744073709551615, F1: 0}
                    v3 = 427
                    _ = fi13
                    _ = v14
                }
            }
            if ("hello" != "") {
                var v15 int16 = 1
                v16 := GenS2{F0: (^(v3)), F1: 0}
                _ = v15
                _ = v16
            }
            _ = fi8
        }
    } else {
        fmt.Println(v3)
    }
    var v17 int16 = 1
    for fi18 := 0; fi18 < 19; fi18++ {
        if len(v5) > 0 {
            v5[0] = -2147483648
        }
        v19 := GenS3{F0: -55}
        var v20 float32 = 1.5
        _ = fi18
        _ = v19
        _ = v20
    }
    var v21 string = "\n\t"
    v21 = ("\n\t" + v21)
    v22 := GenS3{F0: 42}
    v3 = v3
    v23 := GenS3{F0: 0}
    v24 := GenS4{F0: 0}
    if len(v5) > 0 {
        v5[0] = 32768
    }
    v4[2] = v21
    switch ((v21 + v21) + ((v21 + v21) + v21)) {
    case "a":
        if true {
            v25 := GenS5{F0: 68, F1: -92}
            _ = v25
        }
    case "foo bar":
        var v26 int8 = 42
        _ = v26
    }
    if len(v5) > 0 {
        v5[0] = 1
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
