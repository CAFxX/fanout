package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32; F1 uint64; F2 bool }

func main() {
    v1 := make(map[string]bool)
    var v2 int8 = 84
    var v3 bool = (v2 == -65)
    switch 128 {
    case 128:
        var v4 uint = 43
        _ = v4
    case 168:
        switch "a" {
        case "\n\t":
            switch "\n\t" {
            case "hello":
                for fi5 := 0; fi5 < 5; fi5++ {
                    v6 := GenS1{F0: 255}
                    var v7 uint32 = 0
                    switch (-(fi5)) {
                    case 32768:
                        v2 = v2
                    case 128:
                        v8 := []int{9223372036854775807, (-(fi5))}
                        _ = v8
                    }
                    _ = fi5
                    _ = v6
                    _ = v7
                }
            }
        }
    case -47:
        fmt.Println(v2, v2)
    default:
        v3 = (-1 > 9223372036854775807)
    }
    if v3 {
        switch 32767 {
        case -2147483648:
            for fi9 := 0; fi9 < 16; fi9++ {
                v3 = (v3 || false)
                for fi10 := 0; fi10 < 12; fi10++ {
                    var v11 uint64 = 0
                    v12 := []int{(^(fi9))}
                    v13 := GenS1{F0: 31}
                    _ = fi10
                    _ = v11
                    _ = v12
                    _ = v13
                }
                delete(v1, ((("hello" + ("foo bar" + "a")) + "") + "hello"))
                _ = fi9
            }
        case 32768:
            var v14 int64 = -962
            _ = v14
        }
    } else {
        v2 = v2
        v3 = (v2 < (v2 + 127))
        switch -9223372036854775808 {
        case -533:
            for fi15 := 0; fi15 < 8; fi15++ {
                v3 = (!((!false) || false))
                switch "\n\t" {
                case "foo bar":
                    fmt.Println(v2)
                case "":
                    v16 := GenS1{F0: 681}
                    _ = v16
                case "hello":
                    var v17 int8 = v2
                    _ = v17
                default:
                    if ((v3 && true) || true) {
                        v18 := []int{fi15, -9223372036854775808, fi15}
                        fmt.Println(fi15, fi15)
                        _ = v18
                    } else {
                        fmt.Println(fi15, fi15)
                        switch (fi15 << ((42) & 63)) {
                        case 128:
                            var v19 bool = (v3 && false)
                            _ = v19
                        default:
                            for fi20 := 0; fi20 < 11; fi20++ {
                                v21 := GenS1{F0: 32767}
                                _ = fi20
                                _ = v21
                            }
                        }
                        v2 += (-(v2))
                    }
                }
                _ = fi15
            }
        }
    }
    v22 := []int{}
    v3 = false
    delete(v1, "hello")
    var v23 uint = 18446744073709551615
    v2 += (v2 & 127)
    var v24 int = 127
    if (((v3 || true) && (false || v3)) || v3) {
        v25 := GenS2{F0: 42, F1: 143, F2: v3}
        switch ((("hello" + "foo bar") + "") + "foo bar") {
        case "foo bar":
            fmt.Println(v24, v2, v23)
        }
        _ = v25
    } else {
        v3 = true
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
