package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string }

func main() {
    v1 := make(map[int]string)
    var v2 uint32 = 845
    var v3 uint64 = 42
    var v4 int = -1
    switch v4 {
    case -851:
        switch "" {
        case "foo bar":
            switch (v4 << ((2) & 63)) {
            case 1:
                var v5 uint8 = 255
                _ = v5
            case -149:
                v6 := GenS1{F0: -1}
                _ = v6
            default:
                for fi7 := 0; fi7 < 18; fi7++ {
                    v2 = v2
                    var v8 uint8 = 255
                    _ = fi7
                    _ = v8
                }
            }
        }
    case 128:
        var v9 uint = 492
        _ = v9
    }
    v4 -= (v4 - v4)
    if (v4 < (^(v4))) {
        v10 := []int{v4, v4, 956}
        _ = v10
    } else {
        fmt.Println(v2, v2)
    }
    if false {
        v4 = v4
    }
    if false {
        v11 := GenS2{F0: "hello"}
        v3 = (v3 & (v3 ^ v3))
        for fi12 := 0; fi12 < 3; fi12++ {
            v13 := []int{9223372036854775807, -2147483648, fi12}
            _ = fi12
            _ = v13
        }
        _ = v11
    }
    if true {
        v3 += v3
    } else {
        if (753 > 18446744073709551615) {
            v14 := GenS3{F0: ("hello" + "foo bar")}
            _ = v14
        }
        for fi15 := 0; fi15 < 7; fi15++ {
            fmt.Println(v3)
            var v16 float64 = 3.14159
            _ = fi15
            _ = v16
        }
        v3 += 18446744073709551615
    }
    var v17 string = "foo bar"
    fmt.Println(v4)
    v3 -= (^(v3))
    var v18 uint = uint(v3)
    for fi19 := 0; fi19 < 12; fi19++ {
        fmt.Println(v18, v3)
        _ = fi19
    }
    for fi20 := 0; fi20 < 20; fi20++ {
        var v21 int = (v4 + (fi20 % ((fi20) | 1)))
        v18 -= 65535
        _ = fi20
        _ = v21
    }
    for fi22 := 0; fi22 < 1; fi22++ {
        var v23 uint32 = (v2 << ((v18) & 63))
        var v24 uint16 = 32767
        _ = fi22
        _ = v23
        _ = v24
    }
    v17 = v17
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
