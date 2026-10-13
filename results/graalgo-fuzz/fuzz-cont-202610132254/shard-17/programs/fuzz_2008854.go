package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 uint32 }

func main() {
    v1 := []string{}
    var v2 string = "a"
    var v3 bool = (585 <= 4294967295)
    var v4 int = -350
    v5 := make(map[int]uint32)
    if (-1 != 897) {
        fmt.Println(v2, v4)
        if v3 {
            v2 = (v2 + ("foo bar" + "foo bar"))
            if v3 {
                v1 = append(v1, v2)
                switch 42 {
                case -9223372036854775808:
                    v6 := GenS1{F0: 32767, F1: 0}
                    _ = v6
                case -40:
                    v2 = "a"
                case -1:
                    var v7 uint64 = 414
                    _ = v7
                default:
                    if v3 {
                        v8 := GenS1{F0: v4, F1: 0}
                        _ = v8
                    }
                }
                v9 := GenS1{F0: 32768, F1: -2147483648}
                _ = v9
            }
            switch v2 {
            case "a":
                v4 -= v4
            case "hello":
                switch v4 {
                case 778:
                    switch (-(v4)) {
                    case -9223372036854775808:
                        v2 = v2
                    default:
                        v10 := GenS2{F0: 4294967295}
                        _ = v10
                    }
                case 128:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
            }
        }
        for fi11 := 0; fi11 < 13; fi11++ {
            var v12 uint64 = 0
            _ = fi11
            _ = v12
        }
    }
    var v13 uint = 65535
    _ = len(v5)
    fmt.Println(v3, v2, v4)
    var v14 uint32 = 782
    switch v4 {
    case 42:
        v15 := GenS1{F0: v4, F1: 2147483647}
        _ = v15
    case 0:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    default:
        if (!false) {
            for fi16 := 0; fi16 < 11; fi16++ {
                switch ((v2 + "") + (("hello" + (v2 + "foo bar")) + "hello")) {
                case "foo bar":
                    var v17 int = (v4 + (-(v4)))
                    _ = v17
                default:
                    v14 = v14
                }
                _ = fi16
            }
        }
    }
    var v18 float32 = -2.25
    v2 = v2
    switch v2 {
    case "hello":
        v19 := GenS2{F0: (^(v14))}
        _ = v19
    }
    v5[v4] = v14
    var v20 int16 = -1
    for fi21 := 0; fi21 < 4; fi21++ {
        var v22 uint64 = 18446744073709551615
        fmt.Println(v3)
        _ = fi21
        _ = v22
    }
    for fi23 := 0; fi23 < 17; fi23++ {
        if len(v1) > 0 {
            v1[0] = "\n\t"
        }
        v5[(v4 | (v4 + (fi23 >> (((^(v13))) & 63))))] = (v14 + v14)
        v13 += v13
        _ = fi23
    }
    if true {
        v24 := GenS2{F0: (v14 - v14)}
        _ = v24
    } else {
        switch v4 {
        case 604:
            for fi25 := 0; fi25 < 2; fi25++ {
                switch v4 {
                case 32768:
                    for fi26 := 0; fi26 < 18; fi26++ {
                        var v27 bool = false
                        v28 := GenS2{F0: (^(v14))}
                        if v3 {
                            v27 = v27
                        }
                        _ = fi26
                        _ = v27
                        _ = v28
                    }
                }
                v3 = false
                if v3 {
                    var v29 int = (v4 % (((fi25 - fi25)) | 1))
                    _ = v29
                } else {
                    switch (v4 ^ (fi25 + fi25)) {
                    case 42:
                        var v30 uint = (v13 + (v13 | v13))
                        _ = v30
                    default:
                        var v31 uint8 = 209
                        _ = v31
                    }
                    v5[fi25] = (^(v14))
                }
                _ = fi25
            }
        default:
            if (!(-128 > -1)) {
                var v32 bool = ("hello" > "")
                v33 := v5[v4]
                _ = v32
                _ = v33
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
