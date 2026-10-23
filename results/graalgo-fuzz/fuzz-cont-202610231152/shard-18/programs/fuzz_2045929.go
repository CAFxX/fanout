package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int8 }
func genFn0(p0 uint, p1 string, p2 string) string {
    v7 := GenS1{F0: false}
    var v8 uint32 = 2147483648
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return ("\n\t" + p1)
}
func genFn1(p0 uint64) string {
    var v9 uint32 = 175
    v10 := GenS1{F0: ((p0 << ((65535) & 63)) >= (p0 | p0))}
    v9 -= v9
    _ = p0
    _ = v10
    _ = v9
    return "a"
}

func main() {
    var v1 int64 = 0
    var v2 bool = (9223372036854775807 <= v1)
    v3 := make(map[string]uint32)
    v4 := make(map[string]int)
    var v5 int16 = 1000
    var v6 int64 = v1
    v11 := GenS1{F0: v2}
    var v12 uint32 = 506
    var v13 uint8 = 255
    v6 -= v6
    for fi14 := 0; fi14 < 4; fi14++ {
        if ((float32(1.5) != float32(((3.14159 / 3.14159) / 0.1))) || true) {
            var v15 string = ("" + ("hello" + (("a" + "\n\t") + "hello")))
            fmt.Println(v6, v1)
            _ = v15
        } else {
            var v16 uint64 = 2
            for fi17 := 0; fi17 < 13; fi17++ {
                if true {
                    var v18 float32 = float32(0.1)
                    v19 := GenS1{F0: (((!false) || (("foo bar" <= "foo bar") && v2)) || (v16 > (v16 ^ (^(v16)))))}
                    var v20 float64 = 0.1
                    _ = v18
                    _ = v19
                    _ = v20
                }
                _ = fi17
            }
            for fi21 := 0; fi21 < 10; fi21++ {
                if (v2 || v2) {
                    v22 := GenS2{F0: true}
                    switch -2147483648 {
                    case -1:
                        fmt.Println(v2)
                    case 127:
                        switch "\n\t" {
                        case "\n\t":
                            var v23 int64 = v1
                            _ = v23
                        default:
                            var v24 uint8 = 1
                            _ = v24
                        }
                    default:
                        var v25 uint16 = 42
                        _ = v25
                    }
                    var v26 uint = 4294967295
                    var v27 int32 = -2147483648
                    v28 := GenS3{F0: (v13 * 189)}
                    switch (fi14 & (fi14 % ((fi21) | 1))) {
                    case 32767:
                        var v29 uint8 = 143
                        _ = v29
                    case -398:
                        switch "" {
                        case "":
                            var v30 bool = (v2 || v2)
                            _ = v30
                        case "a":
                            v31 := v4[(((("hello" + "") + "") + "\n\t") + ("foo bar" + "foo bar"))]
                            _ = v31
                        case "\n\t":
                            fmt.Println(v27)
                        }
                    }
                    _ = v22
                    _ = v26
                    _ = v27
                    _ = v28
                }
                v13 = (v13 << ((744) & 63))
                switch 1 {
                case -1:
                    if v2 {
                        var v32 int32 = int32(v12)
                        _ = v32
                    } else {
                        v6 = -1
                        fmt.Println(v1)
                    }
                case 2147483647:
                    if v2 {
                        v13 = v13
                        v6 += (^(v1))
                        v12 = (v12 - v12)
                    } else {
                        v33 := GenS1{F0: false}
                        var v34 bool = (uint64(v12) <= v16)
                        _ = v33
                        _ = v34
                    }
                case 0:
                    var v35 float64 = (3.14159 / 1.0)
                    _ = v35
                default:
                    switch "hello" {
                    case "":
                        v16 += (v16 << ((862) & 63))
                    case "a":
                        if v2 {
                            var v36 uint8 = v13
                            _ = v36
                        }
                    case "hello":
                        if v2 {
                            v5 += v5
                            v37 := []int{fi14, int(v12), (fi14 / (((fi21 & -9223372036854775808)) | 1))}
                            _ = v37
                        }
                    }
                }
                _ = fi21
            }
            _ = v16
        }
        var v38 int64 = 1
        var v39 uint32 = v12
        _ = fi14
        _ = v38
        _ = v39
    }
    _ = len(v3)
    v12 += (v12 << ((46) & 63))
    for fi40 := 0; fi40 < 16; fi40++ {
        for fi41 := 0; fi41 < 13; fi41++ {
            fmt.Println(fi41)
            _ = fi41
        }
        _ = fi40
    }
    var v42 string = "foo bar"
    for fi43 := 0; fi43 < 13; fi43++ {
        v44 := GenS3{F0: (v13 + v13)}
        for fi45 := 0; fi45 < 20; fi45++ {
            if (fi43 > (fi43 >> ((65535) & 63))) {
                var v46 float64 = (3.14159 - 0.1)
                _ = v46
            } else {
                v47 := GenS4{F0: (((("hello" + "hello") + ("\n\t" + v42)) + (v42 + ("foo bar" + "hello"))) + v42)}
                if false {
                    v13 = v13
                    var v48 uint8 = (^(v13))
                    fmt.Println(v13, fi45)
                    _ = v48
                } else {
                    v49 := GenS3{F0: v13}
                    _ = v49
                }
                fmt.Println(v42, v6, v5)
                _ = v47
            }
            v50 := GenS5{F0: 127}
            _ = fi45
            _ = v50
        }
        _ = fi43
        _ = v44
    }
    delete(v3, "\n\t")
    var v51 uint64 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v51
    _ = v6
}
