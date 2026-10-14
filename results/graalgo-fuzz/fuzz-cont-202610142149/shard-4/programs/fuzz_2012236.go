package main

import (
    "fmt"
)

func genFn0(p0 int8, p1 int64) uint8 {
    var v4 string = "a"
    v5 := []int{32767, -985, -1}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 42
}
func genFn1(p0 uint32, p1 uint16, p2 int) uint64 {
    var v6 uint = 102
    var v7 int = (p2 - (p2 ^ (p2 + (p2 | 42))))
    p0 += 875
    v8 := []int{(^(v7)), (v7 / (((p2 + 1)) | 1)), p2, (p2 * p2)}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 1
}

func main() {
    v1 := []int{}
    v2 := []uint32{869}
    v3 := []uint32{0, 42, 389}
    for fi9 := 0; fi9 < 16; fi9++ {
        var v10 uint8 = 204
        switch fi9 {
        case 1:
            for fi11 := 0; fi11 < 15; fi11++ {
                var v12 int8 = 42
                _ = fi11
                _ = v12
            }
        }
        _ = fi9
        _ = v10
    }
    var v13 uint32 = 2
    var v14 int16 = 1000
    for fi15 := 0; fi15 < 11; fi15++ {
        v13 -= (^(v13))
        switch ("hello" + ("\n\t" + ("a" + "foo bar"))) {
        case "hello":
            v1 = append(v1, (fi15 | fi15))
        case "":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        v13 = 2
        _ = fi15
    }
    if true {
        if true {
            v14 = v14
        } else {
            var v16 string = "\n\t"
            switch v16 {
            case "":
                fmt.Println(v13, v14, v14)
            default:
                v16 = v16
            }
            var v17 int = -9223372036854775808
            _ = v16
            _ = v17
        }
    }
    v13 += v13
    v14 = v14
    switch (("" + ("foo bar" + ("hello" + "foo bar"))) + ((("\n\t" + "hello") + "") + "hello")) {
    case "hello":
        fmt.Println(v14, v13)
    case "\n\t":
        var v18 uint8 = 128
        _ = v18
    default:
        v13 = 238
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
}
