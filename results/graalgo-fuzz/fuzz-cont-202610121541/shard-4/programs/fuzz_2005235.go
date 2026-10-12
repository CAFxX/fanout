package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 uint64 = 4294967295
    var v2 uint8 = 2
    var v3 uint = uint(v2)
    if (v1 <= v1) {
        var v4 int32 = 51
        fmt.Println(v1, v4, v2)
        if (1 >= 1) {
            fmt.Println(v4, v1, v4)
        }
        _ = v4
    }
    v5 := GenS1{F0: 1000}
    v6 := GenS1{F0: 32767}
    v7 := GenS2{F0: 65535}
    var v8 uint64 = (^(v1))
    if (!(v8 >= (v8 >> ((v3) & 63)))) {
        v1 = (v8 | v1)
        if false {
            v3 -= 0
        }
        v9 := GenS1{F0: 1}
        _ = v9
    }
    var v10 bool = false
    var v11 bool = (!v10)
    for fi12 := 0; fi12 < 12; fi12++ {
        fmt.Println(v11, v10)
        var v13 uint16 = 307
        switch "a" {
        case "":
            v13 -= v13
        case "a":
            v14 := GenS2{F0: 65535}
            _ = v14
        default:
            v15 := []int{}
            _ = v15
        }
        _ = fi12
        _ = v13
    }
    for fi16 := 0; fi16 < 13; fi16++ {
        for fi17 := 0; fi17 < 1; fi17++ {
            if ((("" + ("foo bar" + ("" + "a"))) + "\n\t") != ((("" + "hello") + "foo bar") + (("a" + ("" + "foo bar")) + "a"))) {
                fmt.Println(v3, v11, v10)
                v3 = (v3 << ((v3) & 63))
            }
            v3 += (v3 / ((18446744073709551615) | 1))
            if (!(1 == v1)) {
                v10 = v11
            } else {
                v18 := GenS2{F0: 42}
                var v19 int16 = 1000
                v3 -= v3
                _ = v18
                _ = v19
            }
            _ = fi17
        }
        v20 := GenS1{F0: -32768}
        var v21 uint32 = 345
        _ = fi16
        _ = v20
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
