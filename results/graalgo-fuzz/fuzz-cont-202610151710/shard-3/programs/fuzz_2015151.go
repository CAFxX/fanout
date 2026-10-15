package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32; F1 uint16 }
type GenS4 struct { F0 string; F1 uint64; F2 uint }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint8; F1 int; F2 uint64 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int }
type GenS9 struct { F0 bool; F1 int16; F2 int8 }
func genFn0(p0 int64, p1 int, p2 int32) int16 {
    v5 := []int{(p1 << ((2) & 63)), (^(p1)), p1, p1}
    for fi6 := 0; fi6 < 3; fi6++ {
        if false {
            var v7 int16 = -323
            p0 -= (p0 ^ (p0 ^ (p0 - (-(p0)))))
            _ = v7
        }
        fmt.Println(fi6, fi6)
        var v8 uint = 954
        _ = fi6
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 42
}

func main() {
    var v1 int16 = 1000
    v2 := make(map[string]uint32)
    v3 := []int{42}
    var v4 int16 = 1000
    if (true || (-2.25 == (-2.25 * (1.5 + (float32((0.1 + (((-2.25 - 3.14159) / 1.5) - 0.0))) / 1.0))))) {
        for fi9 := 0; fi9 < 7; fi9++ {
            var v10 uint64 = 18446744073709551615
            _ = fi9
            _ = v10
        }
    } else {
        v4 += -349
        for fi11 := 0; fi11 < 19; fi11++ {
            switch "" {
            case "a":
                var v12 uint16 = 408
                _ = v12
            default:
                for fi13 := 0; fi13 < 1; fi13++ {
                    v14 := GenS1{F0: "a"}
                    _ = fi13
                    _ = v14
                }
            }
            _ = fi11
        }
        if false {
            for fi15 := 0; fi15 < 20; fi15++ {
                switch fi15 {
                case 42:
                    v4 = (v1 | -32768)
                case 2147483647:
                    fmt.Println(fi15)
                }
                _ = fi15
            }
            var v16 uint8 = 214
            v16 -= v16
            _ = v16
        } else {
            v4 += genFn0(-225, 4294967295, -744)
            v17 := v2["hello"]
            v2["hello"] = v17
            _ = v17
        }
    }
    switch "foo bar" {
    case "a":
        if (false && (!(-1 > -1))) {
            fmt.Println(v4, v1)
            if true {
                var v18 bool = true
                _ = v18
            } else {
                v19 := GenS2{F0: 1}
                fmt.Println(v4, v1)
                _ = v19
            }
            var v20 int32 = 0
            _ = v20
        }
    case "":
        if false {
            var v21 uint = 65535
            v22 := GenS2{F0: 65535}
            v1 += v1
            _ = v21
            _ = v22
        } else {
            if (genFn0(-9223372036854775808, 32767, -1) != v4) {
                v23 := GenS3{F0: 2147483648, F1: 1}
                fmt.Println(v4)
                v24 := GenS3{F0: 2147483648, F1: 1}
                _ = v23
                _ = v24
            }
            var v25 int16 = genFn0(-440, 32768, 0)
            _ = v25
        }
    case "foo bar":
        var v26 float32 = 0.0
        _ = v26
    default:
        v1 += v1
    }
    v4 = v4
    fmt.Println(v4, v4)
    var v27 int = 128
    if (v27 != -331) {
        if len(v3) > 0 {
            v3[0] = (v27 / ((2147483647) | 1))
        }
        for fi28 := 0; fi28 < 12; fi28++ {
            v29 := GenS1{F0: "a"}
            _ = fi28
            _ = v29
        }
    }
    var v30 uint = 2
    var v31 uint16 = 2
    fmt.Println(v1)
    if (258 >= (v31 % ((0) | 1))) {
        v1 += 42
    } else {
        if ((("" + (("" + "\n\t") + "")) == "hello") && ((0 <= v31) && ((1 <= -1) || true))) {
            var v32 int8 = 0
            switch v27 {
            case 1:
                switch (^(v27)) {
                case 59:
                    v33 := GenS4{F0: "a", F1: 18446744073709551615, F2: (^(v30))}
                    _ = v33
                default:
                    v32 -= -128
                }
            case 90:
                _ = len(v3)
            }
            _ = v32
        } else {
            for fi34 := 0; fi34 < 10; fi34++ {
                fmt.Println(v31, v30)
                v27 = (^(fi34))
                for fi35 := 0; fi35 < 15; fi35++ {
                    v36 := GenS4{F0: "a", F1: 42, F2: (^(v30))}
                    _ = fi35
                    _ = v36
                }
                _ = fi34
            }
        }
        v37 := GenS2{F0: (v31 + (v31 % ((255) | 1)))}
        v38 := GenS3{F0: 2, F1: (v31 % ((903) | 1))}
        _ = v37
        _ = v38
    }
    if false {
        v27 += -9223372036854775808
        switch int(v31) {
        case 42:
            if (1.5 <= 0.0) {
                v39 := GenS5{F0: -45}
                var v40 int32 = int32(v4)
                _ = v39
                _ = v40
            }
        default:
            v41 := GenS3{F0: 110, F1: v31}
            _ = v41
        }
        var v42 uint64 = 814
        _ = v42
    } else {
        v1 = (v1 + v1)
        if true {
            if (1 != -1) {
                v43 := GenS5{F0: 0}
                _ = v43
            }
            v44 := GenS6{F0: 2, F1: 4294967295, F2: 1}
            _ = v44
        }
    }
    v1 = genFn0(-1, (v27 - -9223372036854775808), 2147483647)
    for fi45 := 0; fi45 < 14; fi45++ {
        var v46 int16 = (v1 ^ (v4 >> ((v30) & 63)))
        for fi47 := 0; fi47 < 1; fi47++ {
            v48 := GenS7{F0: 65535}
            fmt.Println(v1, v46)
            if false {
                v49 := GenS4{F0: "a", F1: 2, F2: v30}
                v50 := GenS8{F0: fi47}
                if true {
                    switch int(v30) {
                    case -935:
                        v2[("" + "a")] = 2147483648
                    }
                    v51 := GenS9{F0: (2 < 826), F1: v1, F2: 42}
                    _ = v51
                } else {
                    fmt.Println(v27, v30)
                    switch ("\n\t" + ("a" + "hello")) {
                    case "":
                        _ = len(v2)
                    case "a":
                        if false {
                            var v52 uint16 = (v31 * (v31 / ((1) | 1)))
                            _ = v52
                        }
                    case "foo bar":
                        var v53 string = "\n\t"
                        var v54 uint64 = 2
                        v53 = ("foo bar" + "")
                        _ = v53
                        _ = v54
                    default:
                        fmt.Println(v30, v30, v4)
                    }
                    fmt.Println(fi45)
                    v30 -= (v30 | v30)
                }
                _ = v49
                _ = v50
            }
            _ = fi47
            _ = v48
        }
        var v55 uint64 = 236
        _ = fi45
        _ = v46
        _ = v55
    }
    v30 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
