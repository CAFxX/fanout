package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool }

func main() {
    v1 := make(map[int]uint32)
    v2 := []string{("\n\t" + "foo bar"), "\n\t", ("" + (("" + ("foo bar" + "")) + "foo bar")), ("\n\t" + ("hello" + "\n\t"))}
    var v3 int32 = -2147483648
    switch -207 {
    case 2147483647:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case 653:
        v3 = (v3 | v3)
    case 1:
        var v4 float32 = (0.1 + -2.25)
        _ = v4
    default:
        switch 599 {
        case 652:
            fmt.Println(v3)
        case 32768:
            if ("foo bar" >= "") {
                v3 += (v3 - (v3 * 1))
            }
        case 741:
            v5 := GenS1{F0: 65535}
            _ = v5
        }
    }
    v3 += v3
    v6 := GenS1{F0: 255}
    var v7 float64 = (-2.25 * 0.0)
    var v8 int64 = -817
    v9 := GenS1{F0: 302}
    fmt.Println(v7, v3, v8)
    var v10 bool = false
    for fi11 := 0; fi11 < 18; fi11++ {
        v7 = v7
        _ = fi11
    }
    v10 = (!(32767 < 42))
    var v12 int = 780
    if v10 {
        v1[int(v8)] = 536
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        switch (v12 + (v12 % ((-555) | 1))) {
        case 920:
            for fi13 := 0; fi13 < 5; fi13++ {
                v14 := GenS1{F0: 0}
                v7 -= -2.25
                _ = fi13
                _ = v14
            }
        case 507:
            for fi15 := 0; fi15 < 17; fi15++ {
                for fi16 := 0; fi16 < 17; fi16++ {
                    switch 4294967295 {
                    case -503:
                        v17 := GenS1{F0: 792}
                        _ = v17
                    }
                    _ = fi16
                }
                v10 = true
                _ = fi15
            }
        case -9223372036854775808:
            if (v10 || (419 <= 229)) {
                v7 += v7
                for fi18 := 0; fi18 < 20; fi18++ {
                    v19 := GenS1{F0: 65535}
                    if v10 {
                        v10 = false
                        fmt.Println(v8)
                        v20 := v1[(v12 | v12)]
                        _ = v20
                    } else {
                        var v21 uint32 = 410
                        fmt.Println(v12)
                        _ = v21
                    }
                    _ = fi18
                    _ = v19
                }
                var v22 uint = 723
                _ = v22
            } else {
                v7 -= 1.5
            }
        default:
            v23 := GenS2{F0: (v10 && (v12 > (v12 ^ v12)))}
            _ = v23
        }
    }
    v10 = (!false)
    v2 = append(v2, ("hello" + (("" + ("a" + "\n\t")) + "foo bar")))
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
