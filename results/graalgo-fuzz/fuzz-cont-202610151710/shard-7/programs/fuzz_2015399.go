package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 uint; F1 uint32 }
type GenS4 struct { F0 int8; F1 uint }
type GenS5 struct { F0 uint32; F1 uint16; F2 int32 }
func genFn0() int32 {
    v4 := GenS1{F0: (((("\n\t" + (("\n\t" + "\n\t") + "hello")) + ("a" + "hello")) + "\n\t") + "")}
    if (42 <= 42) {
        var v5 int = 119
        _ = v5
    } else {
        var v6 uint16 = 16
        for fi7 := 0; fi7 < 14; fi7++ {
            var v8 int16 = 0
            _ = fi7
            _ = v8
        }
        switch 2147483647 {
        case 42:
            if false {
                var v9 uint32 = uint32(v6)
                v10 := []int{841, 128, -2147483648, -674}
                _ = v10
                _ = v9
            }
        default:
            v6 = v6
        }
        _ = v6
    }
    _ = v4
    return 2147483647
}

func main() {
    var v1 uint = 65535
    var v2 int32 = 2147483647
    var v3 int32 = v2
    var v11 uint64 = 18446744073709551615
    v3 = (v3 << (((v1 & (v1 * (^(v1))))) & 63))
    for fi12 := 0; fi12 < 20; fi12++ {
        v13 := GenS2{F0: (v3 & genFn0()), F1: 175}
        _ = fi12
        _ = v13
    }
    fmt.Println(v3, v11)
    if false {
        v14 := []int{int(v1), 54}
        _ = v14
    } else {
        v15 := GenS3{F0: (v1 & 1), F1: 42}
        _ = v15
    }
    v11 += (v11 & 42)
    for fi16 := 0; fi16 < 10; fi16++ {
        if (9223372036854775807 > 155) {
            switch ("" + "\n\t") {
            case "foo bar":
                v17 := GenS3{F0: (v1 + 18446744073709551615), F1: 2147483648}
                _ = v17
            case "":
                var v18 int64 = 42
                _ = v18
            }
        } else {
            v19 := []int{9223372036854775807}
            if false {
                if (!true) {
                    v11 += v11
                    fmt.Println(v3, v3)
                }
                fmt.Println(v3)
            } else {
                v20 := GenS4{F0: 127, F1: v1}
                _ = v20
            }
            if len(v19) > 0 {
                v19[0] = 963
            }
            _ = v19
        }
        v21 := GenS5{F0: 983, F1: 554, F2: (v3 ^ (^(v2)))}
        v11 = 2
        _ = fi16
        _ = v21
    }
    v11 = v11
    fmt.Println(v11, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
}
