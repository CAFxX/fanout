package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64 }
func genFn0() bool {
    switch -9223372036854775808 {
    case 2147483647:
        switch (((("a" + ((("" + "hello") + "\n\t") + "\n\t")) + "hello") + "a") + "") {
        case "a":
            v5 := []int{32768, 939, 127, 42}
            _ = v5
        }
    case 0:
        var v6 string = "foo bar"
        _ = v6
    case 558:
        if true {
            var v7 int8 = 0
            v8 := []int{-9223372036854775808, 9223372036854775807, 0, 0}
            _ = v7
            _ = v8
        } else {
            var v9 uint8 = 15
            switch "hello" {
            case "a":
                fmt.Println(v9)
            case "foo bar":
                for fi10 := 0; fi10 < 19; fi10++ {
                    if (100 < -128) {
                        var v11 uint32 = 251
                        _ = v11
                    }
                    _ = fi10
                }
            }
            _ = v9
        }
    default:
        var v12 float32 = float32(0.1)
        _ = v12
    }
    var v13 uint16 = 1
    _ = v13
    return true
}

func main() {
    var v1 uint16 = 255
    var v2 int32 = -1
    var v3 int32 = (v2 & (v2 + v2))
    v4 := []int{-2147483648}
    v4 = append(v4, 4294967295)
    var v14 int = 32768
    v3 = (v3 ^ (v2 & (v3 * v2)))
    var v15 float32 = 0.0
    v16 := GenS1{F0: (((1.5 + v15) - 1e10) <= v15)}
    if true {
        var v17 uint64 = 813
        for fi18 := 0; fi18 < 3; fi18++ {
            var v19 uint16 = v1
            v20 := GenS1{F0: true}
            for fi21 := 0; fi21 < 2; fi21++ {
                fmt.Println(fi21, v15)
                v4 = append(v4, 0)
                _ = fi21
            }
            _ = fi18
            _ = v19
            _ = v20
        }
        var v22 uint16 = (v1 % ((42) | 1))
        _ = v17
        _ = v22
    }
    v3 = 1
    v23 := GenS2{F0: 0}
    v1 -= 32767
    for fi24 := 0; fi24 < 14; fi24++ {
        switch v14 {
        case -1:
            for fi25 := 0; fi25 < 5; fi25++ {
                fmt.Println(v15, v3, fi24)
                fmt.Println(v1, fi24, v2)
                switch (("a" + ("foo bar" + ("a" + ("foo bar" + "foo bar")))) + "") {
                case "foo bar":
                    var v26 int = fi25
                    _ = v26
                case "hello":
                    _ = len(v4)
                default:
                    v27 := GenS2{F0: -962}
                    _ = v27
                }
                _ = fi25
            }
        default:
            if (!(668 < 4294967295)) {
                fmt.Println(v2, v2)
            }
        }
        _ = fi24
    }
    var v28 int = v14
    if len(v4) > 0 {
        v4[0] = v28
    }
    var v29 uint32 = 92
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
