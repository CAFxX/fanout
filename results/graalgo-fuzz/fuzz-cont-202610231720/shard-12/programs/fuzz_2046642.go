package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 uint }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 bool }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 2
    v3 := make(map[uint32]int)
    v4 := make(map[string]bool)
    v5 := []string{("a" + ""), (("foo bar" + "a") + "foo bar"), "\n\t", "\n\t"}
    var v6 uint16 = 2
    v7 := GenS1{F0: (18446744073709551615 < 4294967295), F1: 2147483647, F2: 255}
    v6 = 32767
    fmt.Println(v2, v2)
    if len(v5) > 0 {
        v5[0] = "hello"
    }
    var v8 int = 2147483647
    if true {
        for fi9 := 0; fi9 < 20; fi9++ {
            switch (^(fi9)) {
            case 2147483647:
                if false {
                    var v10 int = 1
                    if ((float32(3.14159) > (0.0 * ((3.14159 * (3.14159 * (float32((float64(1.5) + 1e10)) / 1.0))) * ((1.5 / -2.25) * (((0.1 + 1.5) - (1.5 - ((1e10 * 0.0) * 1.5))) / 3.14159))))) || true) {
                        var v11 uint = 65535
                        v12 := GenS2{F0: 0, F1: -1}
                        var v13 int16 = -1
                        _ = v11
                        _ = v12
                        _ = v13
                    }
                    _ = v10
                }
            default:
                v6 = (v6 - v6)
            }
            v8 = 127
            _ = fi9
        }
        if ((v2 % ((v2) | 1)) >= 1) {
            v6 += v6
            if len(v5) > 0 {
                v5[0] = ("a" + "foo bar")
            }
            fmt.Println(v2, v8)
        }
    } else {
        var v14 uint16 = (v6 << ((0) & 63))
        _ = v14
    }
    if true {
        var v15 int32 = 1
        _ = v15
    } else {
        if true {
            if true {
                v6 = 282
            } else {
                var v16 int = (-(v8))
                _ = v16
            }
            fmt.Println(v6, v8)
        } else {
            switch -2147483648 {
            case 32767:
                var v17 int = 127
                _ = v17
            default:
                var v18 uint16 = v6
                _ = v18
            }
            for fi19 := 0; fi19 < 15; fi19++ {
                v8 -= 42
                if (447 <= 1000) {
                    if true {
                        v2 = (^(v2))
                        switch "\n\t" {
                        case "":
                            v8 = (^(v8))
                        case "foo bar":
                            var v20 float32 = 0.1
                            _ = v20
                        case "a":
                            v8 += v8
                        default:
                            v21 := GenS2{F0: -387, F1: -1}
                            _ = v21
                        }
                        var v22 bool = (true && (-9223372036854775808 < -204))
                        _ = v22
                    }
                    switch fi19 {
                    case -505:
                        for fi23 := 0; fi23 < 8; fi23++ {
                            v24 := GenS3{F0: 619}
                            _ = fi23
                            _ = v24
                        }
                    case 4294967295:
                        fmt.Println(fi19, v2)
                    }
                }
                _ = fi19
            }
        }
    }
    if true {
        v5 = append(v5, "foo bar")
        var v25 uint64 = 262
        for fi26 := 0; fi26 < 18; fi26++ {
            v27 := v4[(("" + ((("\n\t" + (("foo bar" + "\n\t") + "a")) + "") + "foo bar")) + ("hello" + "hello"))]
            _ = fi26
            _ = v27
        }
        _ = v25
    } else {
        v6 = (v6 & (^(v6)))
    }
    if true {
        v28 := GenS4{F0: v2}
        v2 = 225
        var v29 bool = false
        _ = v28
        _ = v29
    }
    v30 := v3[717]
    v5 = append(v5, ("\n\t" + ""))
    v31 := GenS5{F0: true}
    v2 -= v2
    var v32 int8 = 0
    _ = len(v5)
    if len(v5) > 0 {
        v5[0] = ("hello" + ((("" + ("foo bar" + "\n\t")) + ("" + "")) + "a"))
    }
    if (127 == v32) {
        if (!(!(true && (v6 != (v6 >> ((130) & 63)))))) {
            for fi33 := 0; fi33 < 3; fi33++ {
                v34 := GenS1{F0: (1 <= 18446744073709551615), F1: 1, F2: 812}
                _ = fi33
                _ = v34
            }
            var v35 uint64 = 945
            _ = v35
        } else {
            v36 := GenS4{F0: (v2 >> ((146) & 63))}
            for fi37 := 0; fi37 < 9; fi37++ {
                v2 -= (v2 % (((v2 << ((65535) & 63))) | 1))
                v38 := GenS2{F0: 42, F1: 42}
                v8 += 32768
                _ = fi37
                _ = v38
            }
            _ = v36
        }
        var v39 int8 = (v32 + (-(v32)))
        _ = v39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
