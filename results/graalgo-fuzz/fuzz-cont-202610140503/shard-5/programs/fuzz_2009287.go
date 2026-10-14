package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 uint64; F1 uint; F2 int32 }
func genFn0(p0 bool, p1 int32) bool {
    var v5 uint = 65535
    p0 = (false && p0)
    var v6 uint = v5
    for fi7 := 0; fi7 < 15; fi7++ {
        v8 := GenS1{F0: p0, F1: 42}
        _ = fi7
        _ = v8
    }
    var v9 float64 = float64(3.14159)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v9
    return p0
}

func main() {
    v1 := make(map[string]bool)
    var v2 int64 = 1
    var v3 int = 448
    var v4 string = ("\n\t" + "")
    var v10 int8 = -128
    switch v4 {
    case "foo bar":
        for fi11 := 0; fi11 < 4; fi11++ {
            v12 := GenS2{F0: 42, F1: 255, F2: -700}
            var v13 float32 = (1.5 / float32(1e10))
            switch v3 {
            case 127:
                var v14 uint16 = 1
                _ = v14
            case 42:
                v3 -= 9223372036854775807
            case 14:
                v4 = v4
            }
            _ = fi11
            _ = v12
            _ = v13
        }
    case "hello":
        fmt.Println(v3, v2, v2)
    }
    for fi15 := 0; fi15 < 9; fi15++ {
        if true {
            v16 := GenS2{F0: 18446744073709551615, F1: 234, F2: 2147483647}
            var v17 uint32 = 65535
            _ = v16
            _ = v17
        }
        switch ("a" + v4) {
        case "hello":
            v18 := v1[v4]
            _ = v18
        case "foo bar":
            fmt.Println(v3, fi15, v4)
        }
        _ = fi15
    }
    for fi19 := 0; fi19 < 5; fi19++ {
        v4 = "\n\t"
        var v20 float64 = 0.1
        _ = fi19
        _ = v20
    }
    v10 += (v10 - (v10 >> ((575) & 63)))
    v21 := []int{}
    if len(v21) >= 2 {
        v21 = v21[1:len(v21)]
    }
    if len(v21) > 0 {
        v21[0] = v3
    }
    var v22 uint8 = 181
    switch (v4 + (v4 + v4)) {
    case "a":
        var v23 uint = 521
        _ = v23
    case "":
        fmt.Println(v4)
    case "\n\t":
        for fi24 := 0; fi24 < 1; fi24++ {
            switch (fi24 + v3) {
            case 2147483647:
                if len(v21) > 0 {
                    v21[0] = fi24
                }
            case -1:
                for fi25 := 0; fi25 < 5; fi25++ {
                    if false {
                        _ = len(v21)
                        v26 := GenS2{F0: 1, F1: 280, F2: 2147483647}
                        _ = v26
                    }
                    fmt.Println(v3)
                    _ = fi25
                }
            case 9223372036854775807:
                var v27 bool = (165 > 42)
                _ = v27
            }
            _ = fi24
        }
    default:
        v3 += (v3 * v3)
    }
    if false {
        v28 := GenS1{F0: true, F1: 0}
        for fi29 := 0; fi29 < 11; fi29++ {
            fmt.Println(v2, v3)
            _ = fi29
        }
        switch v3 {
        case -2147483648:
            var v30 int32 = -708
            _ = v30
        }
        _ = v28
    }
    for fi31 := 0; fi31 < 12; fi31++ {
        v2 = v2
        _ = fi31
    }
    for fi32 := 0; fi32 < 19; fi32++ {
        v10 -= 72
        _ = fi32
    }
    if ("" == "\n\t") {
        var v33 uint = 18446744073709551615
        switch (v3 >> (((v33 % ((v33) | 1))) & 63)) {
        case 128:
            var v34 uint64 = 42
            _ = v34
        case 394:
            switch (v4 + v4) {
            case "\n\t":
                var v35 uint32 = 42
                _ = v35
            case "hello":
                fmt.Println(v10, v2, v2)
            }
        default:
            v3 += 1
        }
        _ = v33
    } else {
        var v36 uint8 = 42
        _ = v36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
