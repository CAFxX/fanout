package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool }
func genFn0(p0 uint64) string {
    v7 := []int{}
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 bool = ("foo bar" < "a")
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = v7
    return "\n\t"
}

func main() {
    var v1 int32 = 2147483647
    v2 := []string{"a", "a", "hello"}
    var v3 string = "foo bar"
    var v4 string = v3
    v5 := make(map[int]uint32)
    var v6 int16 = 32767
    switch (("foo bar" + "a") + "foo bar") {
    case "a":
        switch (v3 + "a") {
        case "":
            var v10 uint8 = 228
            _ = v10
        case "a":
            for fi11 := 0; fi11 < 13; fi11++ {
                for fi12 := 0; fi12 < 11; fi12++ {
                    v13 := GenS1{F0: (v1 ^ (v1 << ((255) & 63)))}
                    _ = fi12
                    _ = v13
                }
                _ = fi11
            }
        case "\n\t":
            v6 += 201
        default:
            switch "\n\t" {
            case "foo bar":
                if len(v2) > 0 {
                    v2[0] = v3
                }
            case "hello":
                var v14 float64 = (0.1 - (0.0 - 0.1))
                _ = v14
            case "\n\t":
                for fi15 := 0; fi15 < 12; fi15++ {
                    fmt.Println(v3, v6)
                    fmt.Println(fi15, fi15)
                    _ = fi15
                }
            }
        }
    case "hello":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        var v16 uint32 = 2147483648
        _ = v16
    }
    _ = len(v5)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v17 string = (v4 + v4)
    _ = len(v5)
    switch "hello" {
    case "\n\t":
        var v18 int16 = (v6 % ((v6) | 1))
        _ = v18
    default:
        fmt.Println(v1, v17)
    }
    fmt.Println(v17, v17, v3)
    switch v4 {
    case "\n\t":
        v1 -= 2147483647
    }
    v19 := GenS1{F0: v1}
    if (!true) {
        fmt.Println(v1, v3, v17)
        v4 = (("\n\t" + "\n\t") + ("a" + ("foo bar" + ("\n\t" + v3))))
    } else {
        for fi20 := 0; fi20 < 20; fi20++ {
            if (!true) {
                v3 = v4
                var v21 uint16 = 255
                v4 = "foo bar"
                _ = v21
            }
            _ = fi20
        }
        _ = len(v2)
        v3 = (v3 + v4)
    }
    fmt.Println(v1)
    v22 := GenS2{F0: ((!false) && true)}
    v23 := GenS2{F0: (!(!(!(!true))))}
    var v24 int32 = v1
    for fi25 := 0; fi25 < 17; fi25++ {
        var v26 uint8 = 0
        _ = fi25
        _ = v26
    }
    v2 = append(v2, "a")
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
