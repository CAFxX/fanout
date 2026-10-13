package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 int) int8 {
    fmt.Println(p0)
    var v6 float32 = (0.0 + 3.14159)
    v6 = ((-2.25 - v6) * 0.0)
    for fi7 := 0; fi7 < 2; fi7++ {
        for fi8 := 0; fi8 < 15; fi8++ {
            v9 := GenS1{F0: 18}
            _ = fi8
            _ = v9
        }
        var v10 uint8 = 0
        _ = fi7
        _ = v10
    }
    _ = p0
    _ = v6
    return -1
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 uint8 = 128
    var v3 int32 = -1
    var v4 string = (("hello" + ("a" + "foo bar")) + "hello")
    var v5 string = ((v4 + "") + ((v4 + v4) + v4))
    if (42 > genFn0(-173)) {
        var v11 int16 = 0
        v3 = -2147483648
        v4 = (("foo bar" + v4) + v4)
        _ = v11
    }
    var v12 int32 = 767
    v5 = ""
    v13 := GenS1{F0: 2}
    switch v5 {
    case "foo bar":
        v2 -= (v2 * v2)
    case "hello":
        v14 := GenS1{F0: -33}
        _ = v14
    default:
        v15 := GenS2{F0: 2}
        _ = v15
    }
    var v16 string = v4
    for fi17 := 0; fi17 < 15; fi17++ {
        var v18 uint64 = 4294967295
        var v19 uint32 = 2
        _ = fi17
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 9; fi20++ {
        v21 := GenS1{F0: -128}
        switch (fi20 | (fi20 - fi20)) {
        case 258:
            v22 := GenS1{F0: 0}
            _ = v22
        default:
            for fi23 := 0; fi23 < 20; fi23++ {
                var v24 int32 = -282
                v25 := []int{(fi20 >> ((2) & 63)), fi23}
                _ = fi23
                _ = v24
                _ = v25
            }
        }
        var v26 uint16 = 967
        _ = fi20
        _ = v21
        _ = v26
    }
    var v27 int16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
