package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 int64 }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 string = (("" + "hello") + "foo bar")
    var v2 uint8 = 128
    var v3 bool = (!((!true) && true))
    var v4 uint64 = 18446744073709551615
    var v5 int8 = -1
    v3 = (!v3)
    v2 = v2
    var v6 uint32 = 1
    if false {
        var v7 uint32 = v6
        _ = v7
    } else {
        for fi8 := 0; fi8 < 17; fi8++ {
            v9 := []int{fi8}
            v3 = (false || ((fi8 % ((fi8) | 1)) >= (fi8 | fi8)))
            v3 = v3
            _ = fi8
            _ = v9
        }
        switch v1 {
        case "foo bar":
            var v10 int8 = (-(v5))
            _ = v10
        case "hello":
            switch ("\n\t" + v1) {
            case "foo bar":
                v6 -= (v6 / (((v6 >> ((18446744073709551615) & 63))) | 1))
            case "\n\t":
                v1 = ((v1 + (("hello" + v1) + (v1 + v1))) + v1)
            case "a":
                var v11 float64 = float64(3.14159)
                _ = v11
            default:
                if ((670 > 255) || ((!(!v3)) && true)) {
                    v6 -= (v6 ^ 564)
                    var v12 float64 = (((1e10 * 0.1) - 3.14159) * (1e10 + float64(0.1)))
                    _ = v12
                }
            }
        default:
            v6 += 834
        }
    }
    for fi13 := 0; fi13 < 16; fi13++ {
        var v14 int32 = -389
        _ = fi13
        _ = v14
    }
    v15 := []int{-9223372036854775808, 127, 27}
    switch (v1 + (v1 + v1)) {
    case "\n\t":
        var v16 string = v1
        _ = v16
    case "":
        v3 = (!false)
    default:
        for fi17 := 0; fi17 < 5; fi17++ {
            v18 := GenS1{F0: (v5 + (v5 | 100)), F1: v6, F2: 517}
            if ("foo bar" <= v1) {
                v15 = append(v15, fi17)
                switch (fi17 * 2147483647) {
                case -420:
                    fmt.Println(v1, v2, fi17)
                case 905:
                    for fi19 := 0; fi19 < 14; fi19++ {
                        if len(v15) > 0 {
                            v15[0] = -2147483648
                        }
                        v20 := GenS2{F0: 2147483647}
                        v1 = v1
                        _ = fi19
                        _ = v20
                    }
                case -638:
                    v4 -= 2
                default:
                    var v21 string = "hello"
                    _ = v21
                }
                v22 := GenS1{F0: v5, F1: 65535, F2: -206}
                _ = v22
            } else {
                fmt.Println(v1)
                v2 += (v2 << ((1) & 63))
            }
            v23 := GenS1{F0: v5, F1: (v6 << ((4294967295) & 63)), F2: 1}
            _ = fi17
            _ = v18
            _ = v23
        }
    }
    for fi24 := 0; fi24 < 5; fi24++ {
        switch 127 {
        case 1:
            v25 := GenS2{F0: v6}
            _ = v25
        case -445:
            v26 := GenS1{F0: v5, F1: (v6 + (v6 ^ 4294967295)), F2: -9223372036854775808}
            _ = v26
        default:
            if false {
                v27 := GenS2{F0: (v6 / ((4294967295) | 1))}
                var v28 int8 = (v5 % ((v5) | 1))
                _ = v27
                _ = v28
            } else {
                var v29 float32 = 0.0
                var v30 uint64 = v4
                var v31 int8 = v5
                _ = v29
                _ = v30
                _ = v31
            }
        }
        _ = fi24
    }
    var v32 uint = 2
    var v33 uint16 = uint16(v6)
    switch "\n\t" {
    case "":
        var v34 bool = v3
        _ = v34
    case "a":
        for fi35 := 0; fi35 < 9; fi35++ {
            if true {
                for fi36 := 0; fi36 < 9; fi36++ {
                    v32 = v32
                    _ = fi36
                }
            }
            v15 = append(v15, fi35)
            fmt.Println(v3, v2)
            _ = fi35
        }
    }
    for fi37 := 0; fi37 < 1; fi37++ {
        var v38 uint16 = (v33 - (v33 >> ((255) & 63)))
        fmt.Println(v5)
        if true {
            v1 = ("a" + "foo bar")
            for fi39 := 0; fi39 < 10; fi39++ {
                v40 := GenS2{F0: (v6 + (v6 + 4294967295))}
                if len(v15) > 0 {
                    v15[0] = (fi37 % ((fi39) | 1))
                }
                var v41 bool = (0 < int32(v33))
                _ = fi39
                _ = v40
                _ = v41
            }
            v38 -= v38
        }
        _ = fi37
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
