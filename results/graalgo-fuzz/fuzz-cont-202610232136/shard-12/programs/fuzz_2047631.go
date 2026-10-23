package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
func genFn0(p0 uint16, p1 uint16, p2 uint64) uint16 {
    switch 434 {
    case 4294967295:
        v6 := []int{128, 4294967295, 1, 0}
        _ = v6
    }
    v7 := []int{490, -969, 4294967295}
    switch 42 {
    case 128:
        if len(v7) > 0 {
            v7[0] = int(p2)
        }
    case 32768:
        fmt.Println(p2)
    case 215:
        fmt.Println(p2, p1, p1)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p0
}

func main() {
    var v1 int = 127
    var v2 string = ("a" + "\n\t")
    var v3 uint16 = 0
    v4 := []uint32{}
    v5 := []uint32{2147483647, 0, 410}
    v2 = ""
    switch -67 {
    case 741:
        v1 -= (v1 + (v1 - 872))
    case 42:
        v3 = v3
    }
    v2 = v2
    if true {
        var v8 int = v1
        var v9 float32 = 1e10
        _ = v8
        _ = v9
    }
    switch v2 {
    case "a":
        for fi10 := 0; fi10 < 4; fi10++ {
            for fi11 := 0; fi11 < 3; fi11++ {
                var v12 uint32 = 889
                switch fi11 {
                case 760:
                    v2 = v2
                }
                var v13 int = (fi10 & (fi11 >> ((42) & 63)))
                _ = fi11
                _ = v12
                _ = v13
            }
            fmt.Println(fi10, fi10)
            var v14 string = ((((v2 + (v2 + "foo bar")) + "\n\t") + (v2 + v2)) + v2)
            _ = fi10
            _ = v14
        }
    case "foo bar":
        if true {
            switch "hello" {
            case "\n\t":
                var v15 int32 = 0
                _ = v15
            case "foo bar":
                v16 := GenS1{F0: ((735 != 42) && (!false)), F1: 1}
                _ = v16
            }
            v3 = v3
        } else {
            var v17 int16 = 1
            if false {
                v17 = (-(v17))
            } else {
                var v18 uint16 = (^(v3))
                var v19 uint = 18446744073709551615
                _ = v18
                _ = v19
            }
            _ = v17
        }
    case "hello":
        if (776 < -1) {
            v3 -= v3
            v2 = (v2 + v2)
        }
    default:
        if (750 != 1000) {
            switch "\n\t" {
            case "hello":
                v20 := GenS1{F0: false, F1: 9223372036854775807}
                _ = v20
            }
        } else {
            if (128 > 0) {
                v3 += (^(v3))
                var v21 int16 = -1
                _ = v21
            }
        }
    }
    v22 := GenS1{F0: false, F1: -705}
    fmt.Println(v1, v3)
    var v23 int64 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
