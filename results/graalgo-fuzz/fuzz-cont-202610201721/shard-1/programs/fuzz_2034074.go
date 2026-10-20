package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 int16) int16 {
    fmt.Println(p0)
    p0 -= p0
    p0 -= 42
    fmt.Println(p0)
    _ = p0
    return (^(p0))
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int8 = -1
    var v3 uint64 = 18446744073709551615
    v4 := make(map[string]string)
    _ = len(v4)
    if true {
        v5 := GenS1{F0: 4294967295}
        v6 := []int{128, 32768}
        if true {
            fmt.Println(v2)
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        }
        _ = v5
        _ = v6
    } else {
        fmt.Println(v2, v2)
        v7 := []int{-9223372036854775808}
        for fi8 := 0; fi8 < 8; fi8++ {
            fmt.Println(fi8, v2)
            var v9 uint16 = 963
            _ = fi8
            _ = v9
        }
        _ = v7
    }
    v3 = (v3 * v3)
    v10 := GenS1{F0: uint(v3)}
    fmt.Println(v2, v3)
    if ((v3 ^ 4294967295) < 0) {
        var v11 uint8 = 125
        _ = v11
    } else {
        if true {
            for fi12 := 0; fi12 < 18; fi12++ {
                v3 = v3
                switch fi12 {
                case -573:
                    switch "a" {
                    case "":
                        var v13 bool = false
                        _ = v13
                    case "\n\t":
                        if ((!(("\n\t" + "") > "foo bar")) || true) {
                            v14 := GenS1{F0: 1}
                            fmt.Println(v3)
                            var v15 int8 = -117
                            _ = v14
                            _ = v15
                        } else {
                            switch "\n\t" {
                            case "":
                                for fi16 := 0; fi16 < 5; fi16++ {
                                    v2 = v2
                                    fmt.Println(v3, fi12)
                                    var v17 uint32 = 168
                                    _ = fi16
                                    _ = v17
                                }
                            }
                            v3 = v3
                        }
                    case "foo bar":
                        fmt.Println(fi12, fi12)
                    }
                }
                _ = fi12
            }
        } else {
            v3 -= (v3 & 18446744073709551615)
            v18 := []int{-9223372036854775808, 2147483647, 127, 2147483647}
            _ = v18
        }
    }
    var v19 uint8 = 255
    fmt.Println(v3, v19)
    switch -2147483648 {
    case 974:
        v20 := []int{904, 9223372036854775807, -163}
        _ = v20
    case 158:
        v21 := GenS1{F0: 42}
        _ = v21
    case 127:
        var v22 float32 = float32(0.1)
        _ = v22
    }
    v19 += 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
