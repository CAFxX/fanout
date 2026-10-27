package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 int32, p1 uint16) int32 {
    v5 := GenS1{F0: (^(p1)), F1: 0}
    p1 = p1
    v6 := GenS1{F0: p1, F1: 42}
    p0 = p0
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return p0
}
func genFn1() int64 {
    for fi7 := 0; fi7 < 2; fi7++ {
        v8 := GenS1{F0: 630, F1: 42}
        var v9 int8 = 119
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 uint64 = 2
    fmt.Println(v10)
    v11 := []int{0, -2147483648, 2147483647}
    fmt.Println(v10)
    _ = v10
    _ = v11
    return -497
}

func main() {
    var v1 uint8 = 0
    var v2 int = -9223372036854775808
    v3 := make(map[string]int)
    var v4 int16 = 1000
    v2 = (v2 ^ 9223372036854775807)
    v1 = 2
    var v12 bool = false
    v13 := GenS2{F0: 649}
    v1 = 0
    v12 = v12
    var v14 uint8 = v1
    v1 += 127
    switch ((("a" + "\n\t") + "") + "") {
    case "":
        if v12 {
            var v15 string = ""
            _ = v15
        } else {
            var v16 uint64 = 188
            for fi17 := 0; fi17 < 11; fi17++ {
                switch (v2 % (((fi17 / ((776) | 1))) | 1)) {
                case 649:
                    v12 = (0.0 < 1e10)
                case 2147483647:
                    v14 = (v14 / ((v1) | 1))
                case 32767:
                    for fi18 := 0; fi18 < 11; fi18++ {
                        var v19 uint = 65535
                        _ = fi18
                        _ = v19
                    }
                default:
                    v2 = v2
                }
                for fi20 := 0; fi20 < 3; fi20++ {
                    v4 = v4
                    _ = len(v3)
                    switch "foo bar" {
                    case "a":
                        var v21 string = "foo bar"
                        _ = v21
                    }
                    v12 = (598 > 42)
                    v4 += v4
                    fmt.Println(v14)
                    var v22 uint = 1
                    _ = fi20
                    _ = v22
                }
                if true {
                    var v23 int16 = 567
                    _ = v23
                }
                _ = fi17
            }
            _ = v16
        }
    case "a":
        v24 := []int{(v2 | v2), v2}
        _ = v24
    case "\n\t":
        v3[("a" + "a")] = v2
    default:
        if (0 == -273) {
            v4 = 1
            v4 = (^(v4))
        } else {
            var v25 uint32 = 42
            fmt.Println(v4)
            _ = v25
        }
    }
    fmt.Println(v1, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
