package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 int64 }
type GenS2 struct { F0 int64; F1 int64; F2 uint64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64; F1 bool }

func main() {
    var v1 float32 = 0.0
    var v2 float64 = 1e10
    var v3 bool = ((2147483647 == -493) || (42 == 0))
    v1 = v1
    if (93 == 113) {
        v4 := GenS1{F0: 65535, F1: 2147483647, F2: 42}
        _ = v4
    } else {
        for fi5 := 0; fi5 < 4; fi5++ {
            switch "foo bar" {
            case "foo bar":
                var v6 uint8 = 243
                _ = v6
            case "a":
                v7 := GenS2{F0: -9223372036854775808, F1: 93, F2: 4294967295}
                _ = v7
            case "":
                v1 -= (v1 * float32((v2 / (1e10 / v2))))
            }
            v8 := GenS2{F0: -9223372036854775808, F1: 42, F2: 18446744073709551615}
            _ = fi5
            _ = v8
        }
    }
    v1 += 3.14159
    v9 := []int{981, -9223372036854775808, -9223372036854775808}
    v2 = v2
    if ("foo bar" <= "foo bar") {
        switch "a" {
        case "\n\t":
            for fi10 := 0; fi10 < 5; fi10++ {
                _ = len(v9)
                _ = fi10
            }
        case "":
            for fi11 := 0; fi11 < 7; fi11++ {
                fmt.Println(fi11, fi11, v3)
                v2 = 0.1
                _ = fi11
            }
        }
    }
    if ((false && (false && false)) || ((!v3) || ((!(2147483648 > 65535)) && (-2.25 > float64(v1))))) {
        for fi12 := 0; fi12 < 19; fi12++ {
            fmt.Println(v1)
            var v13 int8 = 42
            v3 = (v2 != float64(-2.25))
            _ = fi12
            _ = v13
        }
    } else {
        switch "\n\t" {
        case "":
            for fi14 := 0; fi14 < 4; fi14++ {
                v3 = ((fi14 * 32767) != fi14)
                v1 = v1
                fmt.Println(fi14, fi14, v3)
                _ = fi14
            }
        }
    }
    v2 -= v2
    var v15 bool = v3
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    v3 = v15
    for fi16 := 0; fi16 < 13; fi16++ {
        v2 = (float64((0.1 + v1)) * ((1.5 * 1.5) * float64((v1 * 1.5))))
        v3 = (-2147483648 > -2147483648)
        v1 = v1
        _ = fi16
    }
    for fi17 := 0; fi17 < 18; fi17++ {
        fmt.Println(fi17, v3, v1)
        if v3 {
            switch fi17 {
            case 891:
                var v18 uint16 = 32767
                _ = v18
            default:
                v15 = (0 >= 0)
            }
            var v19 int32 = 921
            v20 := GenS3{F0: 558}
            _ = v19
            _ = v20
        } else {
            v9 = append(v9, fi17)
            switch (fi17 << ((18446744073709551615) & 63)) {
            case -1:
                var v21 int16 = -32768
                _ = v21
            case 0:
                switch "foo bar" {
                case "\n\t":
                    for fi22 := 0; fi22 < 12; fi22++ {
                        var v23 int = fi17
                        _ = fi22
                        _ = v23
                    }
                case "foo bar":
                    fmt.Println(fi17)
                case "a":
                    fmt.Println(v2)
                }
            case 1:
                v3 = v15
            }
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
        }
        _ = fi17
    }
    fmt.Println(v3)
    if (!(42 > 100)) {
        v15 = (!v15)
    }
    v15 = false
    v24 := GenS4{F0: 2, F1: (251 == 119)}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v15
    _ = v2
    _ = v24
    _ = v3
    _ = v9
}
