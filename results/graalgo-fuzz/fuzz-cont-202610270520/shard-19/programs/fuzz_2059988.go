package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 uint64 = 1
    var v2 uint8 = uint8(v1)
    v3 := make(map[int]bool)
    v4 := make(map[int]uint32)
    var v5 uint32 = 2147483648
    v2 -= (v2 >> ((650) & 63))
    var v6 int64 = -9223372036854775808
    for fi7 := 0; fi7 < 12; fi7++ {
        switch (fi7 + (fi7 << ((uint(v5)) & 63))) {
        case -819:
            switch fi7 {
            case -841:
                if false {
                    var v8 bool = ((v2 + (v2 >> ((1) & 63))) == 127)
                    _ = v8
                }
            }
        case 9223372036854775807:
            if (!(!true)) {
                v9 := []int{fi7, fi7}
                switch fi7 {
                case -821:
                    _ = len(v4)
                }
                fmt.Println(fi7)
                _ = v9
            }
        case -1:
            v1 = (v1 ^ v1)
        default:
            v2 -= 1
        }
        _ = fi7
    }
    var v10 uint8 = (^(v2))
    v2 = v10
    fmt.Println(v10)
    v10 = uint8(v5)
    var v11 int32 = -1
    for fi12 := 0; fi12 < 9; fi12++ {
        v5 += v5
        switch (fi12 | fi12) {
        case 113:
            fmt.Println(v1, v1, v5)
        default:
            var v13 int = fi12
            _ = v13
        }
        fmt.Println(v1)
        _ = fi12
    }
    var v14 int8 = int8(v11)
    var v15 int16 = -1
    var v16 bool = true
    if (546 < -9223372036854775808) {
        var v17 bool = (!v16)
        fmt.Println(v5)
        var v18 int64 = 383
        _ = v17
        _ = v18
    } else {
        var v19 int = -9223372036854775808
        switch ("\n\t" + ("foo bar" + (("a" + ("hello" + (("hello" + "a") + "a"))) + ("foo bar" + "")))) {
        case "a":
            v20 := GenS1{F0: 2147483647}
            _ = v20
        case "hello":
            v21 := GenS1{F0: v5}
            _ = v21
        default:
            switch "\n\t" {
            case "\n\t":
                var v22 bool = ((v5 | v5) != v5)
                _ = v22
            case "hello":
                v1 -= v1
            }
        }
        _ = v19
    }
    var v23 uint64 = (v1 << ((1) & 63))
    var v24 uint16 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
