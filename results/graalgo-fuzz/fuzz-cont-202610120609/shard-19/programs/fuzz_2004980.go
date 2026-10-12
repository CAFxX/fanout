package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int }

func main() {
    v1 := []int{2147483647, -2147483648, -340}
    var v2 uint64 = 4294967295
    v3 := make(map[uint32]bool)
    for fi4 := 0; fi4 < 16; fi4++ {
        var v5 uint8 = 2
        switch 128 {
        case 4294967295:
            var v6 int8 = -128
            _ = v6
        case -458:
            delete(v3, 2147483647)
        case 14:
            for fi7 := 0; fi7 < 5; fi7++ {
                v2 -= (^(v2))
                _ = fi7
            }
        default:
            if (2 <= 18446744073709551615) {
                for fi8 := 0; fi8 < 14; fi8++ {
                    var v9 string = ("" + ("foo bar" + "\n\t"))
                    _ = fi8
                    _ = v9
                }
                _ = len(v1)
            }
        }
        var v10 uint8 = (v5 % ((v5) | 1))
        _ = fi4
        _ = v10
        _ = v5
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        v2 = 433
        _ = fi11
    }
    v3[2147483647] = false
    var v12 bool = (!true)
    if v12 {
        v2 -= v2
        fmt.Println(v2)
        var v13 int32 = 2147483647
        _ = v13
    } else {
        v12 = ((-2.25 > float64(1.5)) && v12)
        fmt.Println(v2, v12)
        v12 = v12
    }
    if v12 {
        var v14 int8 = -1
        v1 = append(v1, 0)
        _ = v14
    }
    v2 -= (v2 << ((65535) & 63))
    v15 := GenS1{F0: 42}
    var v16 int16 = 1000
    v12 = v12
    var v17 int16 = (v16 & v16)
    for fi18 := 0; fi18 < 1; fi18++ {
        switch (("hello" + "\n\t") + (("\n\t" + "a") + "a")) {
        case "hello":
            for fi19 := 0; fi19 < 15; fi19++ {
                var v20 int = -2147483648
                var v21 uint64 = 42
                if false {
                    v16 = v17
                }
                _ = fi19
                _ = v20
                _ = v21
            }
        case "\n\t":
            switch "hello" {
            case "\n\t":
                v22 := GenS1{F0: 0}
                _ = v22
            case "hello":
                v16 = 0
            case "a":
                var v23 uint = 18446744073709551615
                _ = v23
            default:
                for fi24 := 0; fi24 < 11; fi24++ {
                    v25 := GenS1{F0: 127}
                    var v26 int = (fi24 & 128)
                    v27 := GenS1{F0: 31}
                    _ = fi24
                    _ = v25
                    _ = v26
                    _ = v27
                }
            }
        case "":
            var v28 int64 = 42
            _ = v28
        }
        var v29 float64 = -2.25
        switch fi18 {
        case 32768:
            if (false && v12) {
                for fi30 := 0; fi30 < 11; fi30++ {
                    v31 := GenS2{F0: fi18}
                    var v32 bool = (!false)
                    v17 += v16
                    v29 = v29
                    fmt.Println(v16, fi30)
                    _ = fi30
                    _ = v31
                    _ = v32
                }
            }
        default:
            for fi33 := 0; fi33 < 19; fi33++ {
                fmt.Println(fi18)
                var v34 uint = 18446744073709551615
                v12 = (!v12)
                _ = fi33
                _ = v34
            }
        }
        _ = fi18
        _ = v29
    }
    v12 = (0.0 != 3.14159)
    v3[2147483648] = (!(!true))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
