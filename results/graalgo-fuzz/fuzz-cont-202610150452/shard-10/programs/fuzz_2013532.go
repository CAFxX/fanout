package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 int32) int {
    fmt.Println(p0)
    p0 += (p0 - (p0 << ((524) & 63)))
    var v4 int8 = 100
    for fi5 := 0; fi5 < 1; fi5++ {
        v4 = (v4 / ((v4) | 1))
        fmt.Println(fi5, fi5)
        switch "a" {
        case "a":
            switch ("hello" + "") {
            case "a":
                v6 := GenS1{F0: 601}
                _ = v6
            case "hello":
                var v7 uint = 275
                _ = v7
            case "foo bar":
                v4 = 100
            default:
                p0 = (p0 >> ((18446744073709551615) & 63))
            }
        default:
            var v8 bool = (int64(p0) != 0)
            _ = v8
        }
        _ = fi5
    }
    _ = p0
    _ = v4
    return 9223372036854775807
}

func main() {
    var v1 int8 = 127
    var v2 float32 = (3.14159 / float32(3.14159))
    var v3 int64 = 496
    v9 := []int{0, -493, 4294967295, 315}
    for fi10 := 0; fi10 < 20; fi10++ {
        fmt.Println(v3, fi10, v3)
        _ = fi10
    }
    for fi11 := 0; fi11 < 3; fi11++ {
        v3 = (^(v3))
        v2 = v2
        _ = fi11
    }
    var v12 bool = true
    v9 = append(v9, 4294967295)
    if v12 {
        v1 = (v1 << ((1) & 63))
        switch "hello" {
        case "foo bar":
            if len(v9) > 0 {
                v9[0] = -636
            }
        case "hello":
            var v13 uint32 = 96
            _ = v13
        case "\n\t":
            v1 = (v1 * v1)
        default:
            v14 := GenS2{F0: 4294967295}
            _ = v14
        }
        v2 += v2
    }
    fmt.Println(v12)
    v2 = v2
    for fi15 := 0; fi15 < 3; fi15++ {
        var v16 string = (("hello" + ("a" + "foo bar")) + "")
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 17; fi17++ {
        fmt.Println(fi17)
        v12 = (255 != 2)
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        _ = fi17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v9
}
