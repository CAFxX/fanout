package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 int8 }
type GenS2 struct { F0 string; F1 uint64; F2 uint8 }
type GenS3 struct { F0 int16; F1 uint64; F2 bool }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int16; F1 int32 }
func genFn0(p0 uint64) string {
    var v5 int8 = 127
    switch ("\n\t" + "foo bar") {
    case "":
        v6 := GenS1{F0: 337, F1: 349, F2: -128}
        _ = v6
    }
    v7 := GenS2{F0: (("a" + ((("\n\t" + "") + "foo bar") + "a")) + "hello"), F1: (p0 - (p0 & p0)), F2: 128}
    fmt.Println(p0, v5)
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 uint32 = 8
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = v5
    _ = v7
    return "foo bar"
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]uint32)
    var v3 int16 = 858
    var v4 uint8 = 42
    v10 := GenS3{F0: (v3 * (v3 / ((v3) | 1))), F1: 2, F2: (!false)}
    switch ("foo bar" + (("\n\t" + "hello") + ("hello" + "\n\t"))) {
    case "hello":
        for fi11 := 0; fi11 < 5; fi11++ {
            var v12 uint8 = 226
            _ = fi11
            _ = v12
        }
    case "a":
        switch "a" {
        case "foo bar":
            fmt.Println(v4, v4)
        case "":
            if ((4294967295 == 0) && true) {
                for fi13 := 0; fi13 < 11; fi13++ {
                    for fi14 := 0; fi14 < 9; fi14++ {
                        var v15 uint8 = 103
                        _ = fi14
                        _ = v15
                    }
                    _ = fi13
                }
                for fi16 := 0; fi16 < 10; fi16++ {
                    v3 -= 32767
                    var v17 bool = false
                    _ = fi16
                    _ = v17
                }
                for fi18 := 0; fi18 < 17; fi18++ {
                    v19 := v2[(fi18 / (((^(fi18))) | 1))]
                    _ = fi18
                    _ = v19
                }
            }
        default:
            var v20 uint16 = uint16(v4)
            _ = v20
        }
    }
    v21 := []int{4294967295, -9223372036854775808, -1}
    var v22 string = "a"
    if true {
        v3 = 1
    } else {
        v23 := GenS4{F0: 42}
        for fi24 := 0; fi24 < 1; fi24++ {
            var v25 int64 = -1
            _ = fi24
            _ = v25
        }
        v4 -= v4
        _ = v23
    }
    delete(v1, 2147483647)
    var v26 uint16 = 65535
    for fi27 := 0; fi27 < 2; fi27++ {
        v28 := GenS1{F0: 616, F1: 193, F2: 1}
        _ = len(v2)
        v29 := GenS5{F0: 2147483647}
        _ = fi27
        _ = v28
        _ = v29
    }
    var v30 int = 42
    v22 = v22
    _ = len(v21)
    fmt.Println(v22, v22)
    var v31 bool = false
    v32 := v1[65535]
    delete(v1, 4294967295)
    if v31 {
        v33 := GenS6{F0: -457, F1: 1}
        var v34 int8 = int8(v26)
        _ = v33
        _ = v34
    }
    var v35 uint16 = v26
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v22
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v35
    _ = v4
}
