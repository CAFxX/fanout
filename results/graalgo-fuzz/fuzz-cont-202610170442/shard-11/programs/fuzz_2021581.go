package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 uint64) uint32 {
    fmt.Println(p0)
    v6 := []int{32768, -1, 128, -1}
    _ = p0
    _ = v6
    return 924
}

func main() {
    v1 := []int{32768, 127}
    var v2 bool = (false && (0 >= 100))
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]bool)
    var v5 uint8 = 255
    if true {
        for fi7 := 0; fi7 < 10; fi7++ {
            v5 = v5
            for fi8 := 0; fi8 < 5; fi8++ {
                v5 = v5
                var v9 int16 = 47
                v9 -= v9
                var v10 string = "a"
                _ = fi8
                _ = v10
                _ = v9
            }
            _ = fi7
        }
        if true {
            _ = len(v1)
            var v11 int16 = -32768
            _ = v11
        } else {
            if true {
                switch 2147483647 {
                case 127:
                    if true {
                        v2 = v2
                    } else {
                        var v12 int = 1
                        if false {
                            v2 = (9223372036854775807 != -52)
                        }
                        v13 := GenS1{F0: 18446744073709551615, F1: "\n\t"}
                        _ = v12
                        _ = v13
                    }
                case -9223372036854775808:
                    for fi14 := 0; fi14 < 4; fi14++ {
                        v15 := GenS2{F0: genFn0(42)}
                        switch "foo bar" {
                        case "a":
                            v16 := GenS2{F0: 42}
                            _ = v16
                        case "\n\t":
                            var v17 int = int(v5)
                            _ = v17
                        }
                        _ = fi14
                        _ = v15
                    }
                case 2147483647:
                    var v18 int8 = 42
                    _ = v18
                }
                var v19 uint64 = 1
                _ = v19
            }
            v5 -= 1
            if v2 {
                var v20 int16 = 1000
                v2 = v2
                _ = v20
            }
        }
    } else {
        if len(v1) > 0 {
            v1[0] = 453
        }
    }
    v21 := GenS1{F0: 441, F1: ("foo bar" + "a")}
    switch 32768 {
    case 42:
        var v22 int16 = 32767
        _ = v22
    case 128:
        var v23 int8 = 0
        _ = v23
    case 127:
        var v24 int64 = 758
        _ = v24
    }
    fmt.Println(v2, v2)
    v5 = v5
    switch "a" {
    case "foo bar":
        var v25 uint64 = 1
        _ = v25
    case "\n\t":
        v26 := GenS3{F0: 2}
        _ = v26
    }
    var v27 uint8 = (v5 - v5)
    v2 = (!(v2 && (!(-1 > 2147483647))))
    var v28 float64 = -2.25
    var v29 uint64 = 2
    v4[0] = v2
    var v30 int = -2147483648
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
