package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 bool; F2 string }
func genFn0(p0 uint, p1 int8) int32 {
    v6 := GenS1{F0: 0}
    var v7 int64 = 0
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 42
}

func main() {
    var v1 uint = 316
    var v2 uint32 = 65535
    var v3 uint64 = 702
    var v4 string = ("a" + "foo bar")
    var v5 bool = false
    v5 = false
    switch "" {
    case "":
        switch "hello" {
        case "foo bar":
            switch 4294967295 {
            case 183:
                for fi8 := 0; fi8 < 5; fi8++ {
                    var v9 int16 = -32768
                    _ = fi8
                    _ = v9
                }
            case 764:
                var v10 uint16 = 646
                _ = v10
            case 32768:
                if v5 {
                    v2 -= (v2 | 2147483647)
                    var v11 uint32 = v2
                    _ = v11
                }
            }
        case "\n\t":
            var v12 uint16 = 1
            _ = v12
        default:
            for fi13 := 0; fi13 < 20; fi13++ {
                var v14 uint = 0
                if v5 {
                    v15 := []int{int(v3), (fi13 + -9223372036854775808), (fi13 ^ 1), fi13}
                    _ = v15
                } else {
                    v16 := []int{fi13, 0, fi13, 128}
                    v14 = (v14 + (^(v14)))
                    _ = v16
                }
                if true {
                    var v17 string = ""
                    var v18 uint16 = 65535
                    fmt.Println(v17)
                    _ = v17
                    _ = v18
                }
                _ = fi13
                _ = v14
            }
        }
    default:
        v19 := []int{32767, 127, 1}
        _ = v19
    }
    var v20 uint16 = 2
    v21 := GenS1{F0: (^(v2))}
    fmt.Println(v3, v20, v2)
    v22 := GenS2{F0: 0, F1: (((((v4 + "foo bar") + (v4 + "\n\t")) <= v4) && v5) || (1000 == 1000)), F2: "hello"}
    var v23 uint64 = (v3 - (v3 - 817))
    switch 2147483647 {
    case 4294967295:
        v24 := []int{int(v3), 127, -2147483648}
        _ = v24
    case 2147483647:
        v1 += v1
    default:
        fmt.Println(v20, v23)
    }
    v25 := []int{32768}
    fmt.Println(v23, v5)
    if (!v5) {
        if (v5 || v5) {
            if ((^(v20)) == uint16(v1)) {
                fmt.Println(v2, v23)
                switch v4 {
                case "a":
                    for fi26 := 0; fi26 < 12; fi26++ {
                        v1 = v1
                        switch (fi26 >> (((v1 | (v1 & (v1 ^ v1)))) & 63)) {
                        case 0:
                            var v27 uint32 = 99
                            _ = v27
                        case 127:
                            var v28 uint64 = v23
                            _ = v28
                        case -9223372036854775808:
                            v20 += v20
                        }
                        _ = fi26
                    }
                case "hello":
                    v29 := GenS2{F0: int(v20), F1: v5, F2: "a"}
                    _ = v29
                case "foo bar":
                    v2 -= (v2 | 2)
                default:
                    v4 = (v4 + "\n\t")
                }
            } else {
                v2 = 42
                v30 := GenS2{F0: 32768, F1: (!(v1 > uint(v2))), F2: "a"}
                if (!v5) {
                    for fi31 := 0; fi31 < 20; fi31++ {
                        v25 = append(v25, (fi31 | (fi31 | (fi31 & (fi31 >> ((v1) & 63))))))
                        _ = fi31
                    }
                }
                _ = v30
            }
        } else {
            v32 := GenS2{F0: -1, F1: v5, F2: ((("" + v4) + ((v4 + ((("foo bar" + "hello") + (v4 + v4)) + v4)) + "a")) + (v4 + ""))}
            if len(v25) > 0 {
                v25[0] = 42
            }
            _ = v32
        }
        fmt.Println(v20)
        var v33 float64 = float64(float32((((1.5 / 1.0) + (1.5 / 0.1)) - -2.25)))
        _ = v33
    } else {
        var v34 string = "foo bar"
        for fi35 := 0; fi35 < 3; fi35++ {
            if len(v25) >= 2 {
                v25 = v25[1:len(v25)]
            }
            _ = fi35
        }
        _ = len(v25)
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
