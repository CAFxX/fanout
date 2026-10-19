package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 int32; F1 int16; F2 uint64 }
type GenS3 struct { F0 bool; F1 int16 }
func genFn0(p0 int64, p1 uint, p2 int) bool {
    v4 := []int{(p2 % ((-778) | 1))}
    for fi5 := 0; fi5 < 7; fi5++ {
        v6 := GenS1{F0: -2147483648, F1: 1}
        _ = fi5
        _ = v6
    }
    v7 := GenS2{F0: 2147483647, F1: 42, F2: 2}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v7
    return (0 > 2147483648)
}
func genFn1() uint32 {
    var v8 uint64 = 4294967295
    if (1 <= 42) {
        fmt.Println(v8)
        var v9 float32 = (-2.25 * (3.14159 - ((-2.25 / 0.1) - -2.25)))
        _ = v9
    } else {
        v10 := []int{2147483647, -9223372036854775808, 42}
        var v11 uint64 = v8
        var v12 int16 = -907
        _ = v10
        _ = v11
        _ = v12
    }
    v13 := GenS2{F0: 0, F1: 32767, F2: v8}
    var v14 uint16 = 42
    _ = v13
    _ = v14
    _ = v8
    return 2147483647
}

func main() {
    var v1 uint = 4294967295
    var v2 float32 = (0.1 - 1e10)
    var v3 uint8 = 128
    v15 := []int{2147483647, -1, -755}
    v1 += 18446744073709551615
    for fi16 := 0; fi16 < 10; fi16++ {
        var v17 int16 = 356
        v17 += v17
        fmt.Println(fi16)
        _ = fi16
        _ = v17
    }
    if true {
        v1 = (v1 * v1)
    } else {
        for fi18 := 0; fi18 < 17; fi18++ {
            _ = len(v15)
            _ = fi18
        }
        if len(v15) > 0 {
            v15[0] = 169
        }
        switch -128 {
        case 42:
            v2 = 3.14159
        case 128:
            if len(v15) > 0 {
                v15[0] = 127
            }
        }
    }
    switch "foo bar" {
    case "foo bar":
        for fi19 := 0; fi19 < 5; fi19++ {
            v20 := GenS3{F0: false, F1: -32768}
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
            v21 := GenS3{F0: ("foo bar" <= "foo bar"), F1: 32767}
            _ = fi19
            _ = v20
            _ = v21
        }
    case "\n\t":
        if true {
            for fi22 := 0; fi22 < 15; fi22++ {
                fmt.Println(v1)
                v23 := GenS1{F0: -187, F1: 2}
                v1 = v1
                _ = fi22
                _ = v23
            }
            var v24 int = 128
            v2 = v2
            _ = v24
        }
    case "a":
        if true {
            switch "\n\t" {
            case "\n\t":
                var v25 float32 = ((v2 - float32(-2.25)) - (v2 - -2.25))
                _ = v25
            }
            if false {
                var v26 bool = true
                _ = v26
            }
        }
    }
    fmt.Println(v3, v3)
    fmt.Println(v3, v2, v1)
    var v27 uint16 = 32767
    v27 = v27
    switch "" {
    case "":
        v3 = (v3 | 46)
    case "a":
        for fi28 := 0; fi28 < 17; fi28++ {
            v27 -= (^(v27))
            _ = fi28
        }
    case "hello":
        var v29 uint32 = 2147483648
        _ = v29
    }
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v2
    _ = v27
    _ = v3
}
