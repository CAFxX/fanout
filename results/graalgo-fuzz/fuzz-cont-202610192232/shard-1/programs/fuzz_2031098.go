package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 uint16 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]string)
    var v3 string = (("foo bar" + ("" + "hello")) + "")
    v4 := make(map[int]string)
    var v5 bool = (true && false)
    var v6 uint16 = 47
    for fi7 := 0; fi7 < 10; fi7++ {
        for fi8 := 0; fi8 < 1; fi8++ {
            v6 += 2
            var v9 int16 = -32768
            _ = fi8
            _ = v9
        }
        switch (fi7 % (((fi7 * fi7)) | 1)) {
        case -487:
            var v10 int8 = 15
            _ = v10
        default:
            fmt.Println(fi7, v3)
        }
        for fi11 := 0; fi11 < 8; fi11++ {
            fmt.Println(fi11, v6, fi7)
            v5 = (1 <= 171)
            _ = len(v2)
            if (42 == 18446744073709551615) {
                if ((!(("a" + (("a" + v3) + ("foo bar" + (v3 + v3)))) == ((((v3 + ("\n\t" + "hello")) + v3) + (v3 + (v3 + ((v3 + v3) + v3)))) + v3))) || false) {
                    var v12 float32 = 0.1
                    v5 = false
                    if true {
                        v2[823] = v3
                    }
                    _ = v12
                }
            }
            switch v3 {
            case "":
                fmt.Println(v3, v6)
            case "foo bar":
                v13 := []int{(fi11 << ((62) & 63)), (fi7 % ((fi7) | 1)), 32767}
                _ = v13
            case "\n\t":
                v14 := []int{-9223372036854775808}
                _ = v14
            }
            _ = fi11
        }
        _ = fi7
    }
    v15 := GenS1{F0: (-1 > 0), F1: 155}
    fmt.Println(v5, v6)
    if true {
        var v16 uint = 6
        _ = v16
    } else {
        fmt.Println(v6, v6, v6)
        v17 := GenS1{F0: v5, F1: 127}
        _ = v17
    }
    var v18 int16 = 1000
    fmt.Println(v5)
    fmt.Println(v5, v18, v18)
    var v19 int = 1
    delete(v4, v19)
    var v20 float64 = float64((1e10 / 1.0))
    v19 = 212
    v2[0] = ""
    delete(v4, (^(v19)))
    if false {
        v18 -= (v18 * (v18 % ((v18) | 1)))
    }
    var v21 uint8 = 2
    for fi22 := 0; fi22 < 15; fi22++ {
        v23 := GenS2{F0: 2}
        v20 = v20
        _ = fi22
        _ = v23
    }
    if v5 {
        delete(v2, 691)
        v21 = (^(v21))
    } else {
        for fi24 := 0; fi24 < 1; fi24++ {
            fmt.Println(v18)
            _ = fi24
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
