package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8 }

func main() {
    var v1 int8 = 1
    var v2 int16 = 802
    v3 := make(map[uint32]uint32)
    switch -2147483648 {
    case -303:
        if (4294967295 > 0) {
            v1 = (v1 << ((4294967295) & 63))
            v1 -= 0
            if true {
                if true {
                    if (2 < 18446744073709551615) {
                        switch "hello" {
                        case "\n\t":
                            fmt.Println(v1, v1)
                        }
                        v4 := []int{9223372036854775807, 9223372036854775807, -761, -415}
                        _ = v4
                    }
                    var v5 string = (("foo bar" + "\n\t") + ((("foo bar" + "") + "") + ((("foo bar" + "hello") + "foo bar") + "\n\t")))
                    v1 = (v1 & v1)
                    _ = v5
                }
                v6 := GenS1{F0: 204}
                _ = v6
            }
        }
    }
    v2 -= (^(v2))
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 int64 = -1
        v8 += (v8 >> ((1) & 63))
        if (1 < 0) {
            var v9 float64 = 0.0
            switch "foo bar" {
            case "\n\t":
                for fi10 := 0; fi10 < 20; fi10++ {
                    switch "foo bar" {
                    case "":
                        v11 := []int{}
                        _ = v11
                    case "foo bar":
                        v12 := GenS2{F0: v1}
                        _ = v12
                    default:
                        fmt.Println(fi7)
                    }
                    _ = fi10
                }
            case "":
                v13 := []int{(fi7 * -2147483648)}
                _ = v13
            case "a":
                var v14 string = "foo bar"
                _ = v14
            default:
                var v15 int16 = 1000
                _ = v15
            }
            _ = v9
        }
        _ = fi7
        _ = v8
    }
    v2 += v2
    v2 += 0
    v1 = v1
    var v16 uint16 = 39
    fmt.Println(v1, v2)
    var v17 float32 = (0.1 / 1.0)
    v18 := []int{-973, -1}
    if ("a" != ((("a" + "\n\t") + ("foo bar" + "\n\t")) + "")) {
        v2 = 42
        for fi19 := 0; fi19 < 16; fi19++ {
            var v20 int = (fi19 ^ 4294967295)
            v16 += v16
            _ = fi19
            _ = v20
        }
        var v21 float64 = 1.5
        _ = v21
    }
    var v22 uint16 = (v16 ^ v16)
    var v23 uint8 = 127
    var v24 uint32 = 777
    var v25 int32 = int32(v24)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
