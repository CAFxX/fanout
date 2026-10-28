package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 string }
func genFn0(p0 int16) string {
    if false {
        fmt.Println(p0)
    } else {
        v6 := GenS1{F0: (43 >= 42)}
        _ = v6
    }
    v7 := GenS1{F0: (4294967295 >= 0)}
    for fi8 := 0; fi8 < 6; fi8++ {
        p0 = p0
        _ = fi8
    }
    v9 := GenS1{F0: (255 != 42)}
    _ = p0
    _ = v7
    _ = v9
    return "hello"
}

func main() {
    var v1 int = 321
    var v2 int32 = -2147483648
    var v3 uint16 = 2
    v4 := make(map[uint32]string)
    var v5 int32 = (v2 << ((1) & 63))
    v10 := GenS2{F0: 2, F1: 6}
    var v11 bool = ((-2.25 < 1.5) && ("\n\t" == "hello"))
    for fi12 := 0; fi12 < 3; fi12++ {
        v13 := []int{-84, (fi12 >> ((42) & 63)), fi12}
        for fi14 := 0; fi14 < 8; fi14++ {
            var v15 uint8 = 15
            _ = fi14
            _ = v15
        }
        _ = fi12
        _ = v13
    }
    v1 += (v1 + 1)
    if (638 != 2147483647) {
        v16 := GenS3{F0: (("a" + "") + "\n\t")}
        var v17 int32 = -400
        v1 = (v1 << ((18446744073709551615) & 63))
        _ = v16
        _ = v17
    }
    var v18 string = ""
    for fi19 := 0; fi19 < 15; fi19++ {
        v1 = 1
        _ = fi19
    }
    if v11 {
        fmt.Println(v18, v18)
        for fi20 := 0; fi20 < 19; fi20++ {
            switch (v1 + fi20) {
            case -736:
                var v21 int32 = v2
                _ = v21
            default:
                fmt.Println(v3, v3, v1)
            }
            var v22 uint32 = 2147483647
            _ = fi20
            _ = v22
        }
        var v23 float64 = (((3.14159 / -2.25) + 0.1) / 0.1)
        _ = v23
    }
    for fi24 := 0; fi24 < 18; fi24++ {
        v25 := v4[403]
        var v26 float64 = 0.1
        var v27 uint32 = 0
        _ = fi24
        _ = v25
        _ = v26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
