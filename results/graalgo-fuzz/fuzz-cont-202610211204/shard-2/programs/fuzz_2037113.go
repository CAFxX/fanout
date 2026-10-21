package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 uint32 }
type GenS2 struct { F0 int; F1 uint8; F2 uint64 }
func genFn0(p0 uint32) int8 {
    if (128 == -9223372036854775808) {
        p0 += p0
        switch "foo bar" {
        case "hello":
            p0 = p0
        case "\n\t":
            switch 32767 {
            case 9223372036854775807:
                fmt.Println(p0)
            case -521:
                v4 := GenS1{F0: 4294967295, F1: 1, F2: (p0 / (((^(p0))) | 1))}
                _ = v4
            case 1:
                for fi5 := 0; fi5 < 1; fi5++ {
                    var v6 uint8 = 13
                    v7 := []int{994, -504}
                    _ = fi5
                    _ = v6
                    _ = v7
                }
            default:
                fmt.Println(p0)
            }
        }
    } else {
        v8 := GenS1{F0: 0, F1: -201, F2: (p0 - 65535)}
        v9 := []int{9223372036854775807}
        switch (("" + "") + ("hello" + "")) {
        case "":
            v10 := GenS2{F0: -92, F1: 1, F2: 367}
            _ = v10
        case "\n\t":
            p0 += (p0 * (p0 >> ((42) & 63)))
        }
        _ = v8
        _ = v9
    }
    p0 = 0
    _ = p0
    return 1
}

func main() {
    v1 := []string{"a"}
    v2 := make(map[uint32]int)
    var v3 uint8 = 1
    if true {
        var v11 uint8 = v3
        v11 = (^(v11))
        if false {
            v11 -= v3
            v12 := GenS2{F0: 128, F1: v11, F2: 0}
            var v13 int64 = 815
            _ = v12
            _ = v13
        } else {
            v3 = 42
        }
        _ = v11
    } else {
        var v14 int32 = 0
        v14 = v14
        for fi15 := 0; fi15 < 11; fi15++ {
            fmt.Println(v3, v3)
            var v16 int8 = 1
            _ = fi15
            _ = v16
        }
        _ = v14
    }
    for fi17 := 0; fi17 < 11; fi17++ {
        switch ("\n\t" + "\n\t") {
        case "hello":
            switch "hello" {
            case "foo bar":
                var v18 uint8 = v3
                _ = v18
            case "a":
                fmt.Println(v3, v3)
            }
        case "a":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        fmt.Println(fi17)
        _ = fi17
    }
    fmt.Println(v3)
    v3 = (v3 << ((416) & 63))
    var v19 uint = 63
    var v20 uint32 = 4294967295
    if (1e10 > -2.25) {
        v3 = v3
    }
    v20 += (v20 + (v20 & v20))
    for fi21 := 0; fi21 < 2; fi21++ {
        if true {
            delete(v2, v20)
        }
        _ = fi21
    }
    if false {
        var v22 bool = (-2147483648 > 2147483647)
        _ = v22
    } else {
        var v23 bool = ((290 > -9223372036854775808) || false)
        var v24 uint64 = 1
        v25 := GenS2{F0: 9223372036854775807, F1: 44, F2: (v24 & (v24 & (v24 + v24)))}
        _ = v23
        _ = v24
        _ = v25
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
