package main

import (
    "fmt"
)


func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]uint32)
    v3 := make(map[int]bool)
    v4 := make(map[uint32]int)
    var v5 uint64 = 212
    v5 += (v5 % ((v5) | 1))
    for fi6 := 0; fi6 < 3; fi6++ {
        switch ((("\n\t" + (("" + "foo bar") + "")) + "") + "foo bar") {
        case "":
            var v7 uint32 = 42
            _ = v7
        }
        _ = fi6
    }
    v8 := []int{127, 27}
    for fi9 := 0; fi9 < 1; fi9++ {
        if (((2 < 980) || (false && (!true))) && false) {
            var v10 uint = 65535
            _ = v10
        }
        switch 0 {
        case 1:
            if false {
                for fi11 := 0; fi11 < 9; fi11++ {
                    switch -9223372036854775808 {
                    case -1:
                        if (255 <= uint8(v5)) {
                            delete(v4, 718)
                            if true {
                                v5 += v5
                            } else {
                                fmt.Println(fi9, fi11)
                            }
                        }
                    }
                    _ = fi11
                }
            }
        default:
            v8 = append(v8, 32767)
        }
        _ = fi9
    }
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    var v12 string = "foo bar"
    for fi13 := 0; fi13 < 16; fi13++ {
        var v14 float32 = ((((1.5 - 0.0) * 1.5) / -2.25) + 0.0)
        _ = fi13
        _ = v14
    }
    var v15 int16 = 0
    v5 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
