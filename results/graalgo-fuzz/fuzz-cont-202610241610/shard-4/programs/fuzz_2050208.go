package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]bool)
    var v3 uint16 = 0
    switch 1 {
    case 860:
        v4 := []int{}
        _ = v4
    case -104:
        switch "" {
        case "hello":
            for fi5 := 0; fi5 < 7; fi5++ {
                if false {
                    switch (^(fi5)) {
                    case 1:
                        if (true && (-1 <= 1)) {
                            for fi6 := 0; fi6 < 12; fi6++ {
                                fmt.Println(v3, v3, fi6)
                                v3 -= v3
                                if (18446744073709551615 == 2) {
                                    _ = len(v2)
                                }
                                var v7 uint8 = 255
                                _ = fi6
                                _ = v7
                            }
                            _ = len(v1)
                        }
                    case 42:
                        switch "foo bar" {
                        case "\n\t":
                            delete(v2, ("" + ("a" + "\n\t")))
                        case "a":
                            v3 = v3
                        }
                    case 855:
                        var v8 float32 = 1.5
                        _ = v8
                    default:
                        var v9 int64 = 0
                        _ = v9
                    }
                }
                v3 += 255
                for fi10 := 0; fi10 < 9; fi10++ {
                    var v11 uint32 = 0
                    switch (fi10 % ((32768) | 1)) {
                    case 609:
                        switch (fi5 >> ((18446744073709551615) & 63)) {
                        case 32768:
                            var v12 uint = 341
                            _ = v12
                        case -334:
                            var v13 float64 = (3.14159 / (3.14159 * 1.5))
                            _ = v13
                        }
                    case 2147483647:
                        v14 := GenS1{F0: (fi10 ^ fi5)}
                        _ = v14
                    case 741:
                        v15 := v2[("" + "hello")]
                        _ = v15
                    default:
                        v16 := GenS1{F0: 2147483647}
                        _ = v16
                    }
                    _ = fi10
                    _ = v11
                }
                _ = fi5
            }
        default:
            switch 127 {
            case 4294967295:
                v17 := GenS1{F0: 549}
                _ = v17
            }
        }
    case 9223372036854775807:
        v18 := GenS1{F0: 128}
        _ = v18
    }
    v3 = v3
    v3 = 42
    fmt.Println(v3)
    v3 += v3
    v19 := GenS2{F0: 666}
    if true {
        for fi20 := 0; fi20 < 15; fi20++ {
            switch "foo bar" {
            case "hello":
                v21 := GenS2{F0: 42}
                _ = v21
            default:
                if true {
                    v22 := []int{}
                    _ = v22
                }
            }
            _ = fi20
        }
        fmt.Println(v3)
    } else {
        v23 := []int{128, 127}
        for fi24 := 0; fi24 < 17; fi24++ {
            v25 := GenS2{F0: 1}
            for fi26 := 0; fi26 < 1; fi26++ {
                var v27 string = ("a" + "a")
                _ = fi26
                _ = v27
            }
            _ = fi24
            _ = v25
        }
        var v28 bool = (true || false)
        _ = v23
        _ = v28
    }
    if (1.5 == ((((3.14159 / 1e10) - -2.25) / 1.0) + (0.0 * ((3.14159 + ((((float32(1.5) / 1.5) + 0.0) * -2.25) * 1e10)) / (-2.25 + float32(1.5)))))) {
        delete(v2, "a")
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
}
