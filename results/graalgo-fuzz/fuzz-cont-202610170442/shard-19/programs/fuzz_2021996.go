package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 uint8 }
func genFn0() uint64 {
    var v7 float32 = 3.14159
    fmt.Println(v7)
    _ = v7
    return 57
}

func main() {
    v1 := make(map[string]bool)
    v2 := []int{1}
    var v3 uint16 = 1
    v4 := make(map[uint32]string)
    var v5 int8 = 100
    var v6 int32 = 0
    fmt.Println(v5, v6)
    for fi8 := 0; fi8 < 20; fi8++ {
        v5 += 1
        for fi9 := 0; fi9 < 14; fi9++ {
            v4[2] = ("\n\t" + "\n\t")
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi9
        }
        _ = fi8
    }
    v3 = (v3 << ((2) & 63))
    v3 = (v3 & v3)
    if (((1.5 - float32(-2.25)) - float32(3.14159)) > 1e10) {
        switch 32767 {
        case 4294967295:
            for fi10 := 0; fi10 < 9; fi10++ {
                var v11 float64 = 0.0
                var v12 uint64 = 0
                _ = fi10
                _ = v11
                _ = v12
            }
        }
        fmt.Println(v5)
        if (false || true) {
            var v13 uint8 = 0
            v14 := GenS1{F0: ("a" + "a"), F1: -594, F2: (v13 * (v13 << ((0) & 63)))}
            _ = v13
            _ = v14
        }
    }
    v5 = (v5 << ((0) & 63))
    if (42 >= 2) {
        for fi15 := 0; fi15 < 9; fi15++ {
            v3 += v3
            v16 := GenS1{F0: "hello", F1: (v6 - (v6 * (^(v6)))), F2: 127}
            _ = fi15
            _ = v16
        }
        for fi17 := 0; fi17 < 7; fi17++ {
            v18 := GenS1{F0: "a", F1: 2147483647, F2: 128}
            _ = fi17
            _ = v18
        }
        if false {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            fmt.Println(v3, v6)
        }
    }
    v19 := v4[2147483648]
    fmt.Println(v3, v19)
    v20 := GenS1{F0: ("hello" + ((v19 + v19) + ((((v19 + v19) + "hello") + ("a" + ("\n\t" + ("hello" + ((v19 + v19) + (v19 + ((("" + v19) + v19) + "\n\t"))))))) + "foo bar"))), F1: (^(v6)), F2: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
