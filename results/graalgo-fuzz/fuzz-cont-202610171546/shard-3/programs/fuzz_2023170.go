package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16 }
func genFn0(p0 int8, p1 int64, p2 int64) string {
    var v4 int8 = (p0 ^ p0)
    for fi5 := 0; fi5 < 2; fi5++ {
        v6 := GenS1{F0: 1000}
        v7 := GenS1{F0: 594}
        _ = fi5
        _ = v6
        _ = v7
    }
    if true {
        switch (("a" + ("hello" + "")) + ("" + "hello")) {
        case "hello":
            switch "foo bar" {
            case "foo bar":
                for fi8 := 0; fi8 < 4; fi8++ {
                    var v9 uint64 = 1
                    v10 := GenS1{F0: -32768}
                    _ = fi8
                    _ = v10
                    _ = v9
                }
            case "":
                v11 := []int{609, -2147483648, 127}
                _ = v11
            default:
                fmt.Println(p2, p2, v4)
            }
        case "\n\t":
            var v12 int32 = 1
            _ = v12
        case "a":
            switch -1 {
            case -159:
                v13 := []int{315, 9223372036854775807}
                _ = v13
            case 2147483647:
                v14 := []int{-9223372036854775808, -799, -914}
                _ = v14
            case 42:
                v4 = p0
            }
        default:
            p0 = v4
        }
        p1 += (p1 % ((p2) | 1))
    } else {
        var v15 uint = 1
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return "a"
}
func genFn1(p0 uint8, p1 uint16) int8 {
    p1 = p1
    for fi16 := 0; fi16 < 20; fi16++ {
        if (true || (int64(p0) >= 9223372036854775807)) {
            p0 = p0
        }
        for fi17 := 0; fi17 < 9; fi17++ {
            var v18 uint32 = 596
            _ = fi17
            _ = v18
        }
        _ = fi16
    }
    switch ("foo bar" + "\n\t") {
    case "\n\t":
        for fi19 := 0; fi19 < 2; fi19++ {
            for fi20 := 0; fi20 < 15; fi20++ {
                var v21 uint8 = p0
                var v22 int16 = -42
                var v23 int16 = v22
                _ = fi20
                _ = v21
                _ = v22
                _ = v23
            }
            for fi24 := 0; fi24 < 12; fi24++ {
                v25 := []int{}
                v25 = append(v25, fi24)
                fmt.Println(p1)
                p0 += 128
                _ = fi24
                _ = v25
            }
            _ = fi19
        }
    case "a":
        v26 := []int{127}
        _ = v26
    case "hello":
        var v27 uint32 = 4294967295
        _ = v27
    default:
        fmt.Println(p0)
    }
    var v28 int = 42
    p0 = 128
    _ = p0
    _ = p1
    _ = v28
    return 1
}

func main() {
    var v1 uint16 = 1
    v2 := []int{9223372036854775807, -458}
    var v3 string = "hello"
    _ = len(v2)
    v3 = v3
    v1 = 332
    if len(v2) > 0 {
        v2[0] = -2147483648
    }
    v29 := GenS2{F0: -1}
    if len(v2) > 0 {
        v2[0] = 4294967295
    }
    switch v3 {
    case "hello":
        for fi30 := 0; fi30 < 19; fi30++ {
            var v31 bool = true
            v31 = v31
            v32 := GenS2{F0: 1000}
            _ = fi30
            _ = v31
            _ = v32
        }
    case "foo bar":
        v2 = append(v2, 515)
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v1, v1)
    var v33 int8 = genFn1(2, (v1 >> ((uint(v1)) & 63)))
    var v34 uint16 = 42
    v35 := GenS1{F0: -963}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v33
    _ = v34
    _ = v35
}
