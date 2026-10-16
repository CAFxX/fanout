package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 uint, p1 string) uint8 {
    p0 += (p0 << ((p0) & 63))
    p1 = p1
    p1 = "foo bar"
    _ = p0
    _ = p1
    return 128
}
func genFn1(p0 uint64) int {
    for fi6 := 0; fi6 < 14; fi6++ {
        var v7 int32 = 2147483647
        var v8 uint64 = 18446744073709551615
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := []int{-402}
    _ = p0
    _ = v9
    return 128
}

func main() {
    v1 := []string{("hello" + "")}
    var v2 uint16 = 32767
    v3 := make(map[string]bool)
    v4 := make(map[uint32]uint32)
    v5 := make(map[uint32]uint32)
    fmt.Println(v2)
    for fi10 := 0; fi10 < 7; fi10++ {
        fmt.Println(v2, v2)
        var v11 int16 = 32767
        _ = fi10
        _ = v11
    }
    if ("" > "\n\t") {
        var v12 int64 = -468
        v2 = 121
        _ = v12
    }
    if true {
        switch 2147483647 {
        case 32768:
            for fi13 := 0; fi13 < 6; fi13++ {
                for fi14 := 0; fi14 < 14; fi14++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    switch "hello" {
                    case "\n\t":
                        v15 := GenS1{F0: 488}
                        _ = v15
                    case "":
                        var v16 int = (fi14 * (fi14 * fi13))
                        _ = v16
                    }
                    v2 += v2
                    _ = fi14
                }
                var v17 uint8 = 128
                var v18 uint16 = (v2 * (v2 >> ((738) & 63)))
                _ = fi13
                _ = v17
                _ = v18
            }
        case 4294967295:
            v2 = 602
        }
    } else {
        v2 = (v2 + (v2 & v2))
        for fi19 := 0; fi19 < 8; fi19++ {
            switch ("hello" + ("\n\t" + "hello")) {
            case "foo bar":
                var v20 float32 = 1.5
                _ = v20
            default:
                v2 = v2
            }
            _ = fi19
        }
        fmt.Println(v2)
    }
    _ = len(v1)
    var v21 bool = true
    var v22 int = -92
    var v23 bool = true
    if v23 {
        if v21 {
            delete(v5, 2147483647)
        }
    }
    v23 = v23
    for fi24 := 0; fi24 < 12; fi24++ {
        v25 := GenS1{F0: 42}
        _ = fi24
        _ = v25
    }
    var v26 string = ""
    v27 := GenS1{F0: 0}
    var v28 int8 = -102
    v29 := v5[54]
    if v21 {
        var v30 string = v26
        var v31 int8 = (v28 & (v28 >> ((18446744073709551615) & 63)))
        _ = v30
        _ = v31
    }
    for fi32 := 0; fi32 < 2; fi32++ {
        fmt.Println(v26, fi32, v26)
        var v33 int64 = 9223372036854775807
        if (v21 && true) {
            v21 = ((v29 * (v29 & (v29 / ((0) | 1)))) == v29)
        }
        _ = fi32
        _ = v33
    }
    if false {
        for fi34 := 0; fi34 < 10; fi34++ {
            var v35 int64 = -9223372036854775808
            _ = fi34
            _ = v35
        }
    } else {
        var v36 uint = 255
        var v37 float32 = 0.0
        fmt.Println(v23)
        _ = v36
        _ = v37
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
