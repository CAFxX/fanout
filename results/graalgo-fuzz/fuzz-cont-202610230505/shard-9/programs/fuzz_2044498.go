package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }

func main() {
    var v1 string = ("foo bar" + "\n\t")
    var v2 bool = (726 <= 963)
    var v3 int32 = 268
    v4 := make(map[string]string)
    for fi5 := 0; fi5 < 4; fi5++ {
        fmt.Println(v1, v1)
        v6 := GenS1{F0: 115, F1: 2147483647}
        var v7 uint64 = 189
        _ = fi5
        _ = v6
        _ = v7
    }
    switch ("a" + "\n\t") {
    case "hello":
        fmt.Println(v3, v2, v2)
    case "a":
        v8 := GenS1{F0: 42, F1: (v3 & v3)}
        _ = v8
    case "":
        v2 = (1 >= 2)
    }
    _ = len(v4)
    var v9 uint32 = 42
    v2 = ((1e10 <= 1e10) && false)
    switch 1 {
    case 32768:
        if v2 {
            v10 := []int{42}
            _ = v10
        } else {
            for fi11 := 0; fi11 < 3; fi11++ {
                if (!v2) {
                    v9 = (v9 >> ((599) & 63))
                    var v12 int64 = 0
                    v13 := []int{fi11, 7, (-(fi11)), (fi11 << ((2) & 63))}
                    _ = v12
                    _ = v13
                } else {
                    switch ((v1 + v1) + v1) {
                    case "\n\t":
                        fmt.Println(v1, fi11)
                    case "":
                        switch v1 {
                        case "foo bar":
                            v14 := []int{}
                            _ = v14
                        default:
                            if v2 {
                                if (false || (v2 && v2)) {
                                    v15 := GenS1{F0: 18446744073709551615, F1: v3}
                                    _ = v15
                                }
                                var v16 uint16 = 410
                                _ = v16
                            }
                        }
                    default:
                        delete(v4, v1)
                    }
                }
                _ = fi11
            }
        }
    }
    v4[(((v1 + v1) + ("" + v1)) + v1)] = v1
    if (0 >= (^(v3))) {
        var v17 int = 524
        fmt.Println(v1)
        v2 = (!(("\n\t" + "") <= v1))
        _ = v17
    } else {
        v1 = ((v1 + "") + "")
    }
    v9 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
