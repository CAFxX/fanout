package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 uint8; F2 uint16 }
type GenS3 struct { F0 int16 }
func genFn0(p0 uint32, p1 int8) uint64 {
    var v7 int = 0
    var v8 uint64 = uint64(p0)
    fmt.Println(v7, v8)
    var v9 int64 = 9223372036854775807
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return (v8 >> ((65535) & 63))
}
func genFn1(p0 int32, p1 int32) uint16 {
    v10 := []int{}
    if true {
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        p1 += p1
    } else {
        p0 += 1
        p0 -= p1
        if (!(1 <= 9223372036854775807)) {
            var v11 uint8 = 42
            _ = v11
        } else {
            v10 = append(v10, 42)
            p1 -= (-(p1))
            var v12 int32 = p1
            _ = v12
        }
    }
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    if false {
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        if (2147483648 < 711) {
            var v13 uint32 = 65535
            p1 -= p0
            p0 = p1
            _ = v13
        }
    }
    _ = p0
    _ = p1
    _ = v10
    return 0
}

func main() {
    var v1 bool = (18446744073709551615 != 4294967295)
    v2 := []int{}
    v3 := make(map[string]uint32)
    v4 := []string{}
    var v5 uint = 83
    var v6 uint32 = 4294967295
    var v14 float64 = 1e10
    for fi15 := 0; fi15 < 4; fi15++ {
        v16 := GenS1{F0: (!v1)}
        _ = fi15
        _ = v16
    }
    var v17 uint32 = v6
    if ("\n\t" >= "") {
        v18 := GenS1{F0: true}
        for fi19 := 0; fi19 < 3; fi19++ {
            for fi20 := 0; fi20 < 8; fi20++ {
                var v21 string = ("a" + "\n\t")
                var v22 int = (fi20 << ((v5) & 63))
                _ = fi20
                _ = v21
                _ = v22
            }
            for fi23 := 0; fi23 < 16; fi23++ {
                if (true && (v14 > v14)) {
                    switch (-(fi23)) {
                    case 32768:
                        v24 := GenS1{F0: (fi19 < -1)}
                        _ = v24
                    case 1:
                        switch 920 {
                        case 95:
                            if len(v2) > 0 {
                                v2[0] = fi19
                            }
                        default:
                            v14 = 3.14159
                        }
                    case 128:
                        if v1 {
                            var v25 int16 = 32767
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            _ = v25
                        }
                    default:
                        v26 := GenS2{F0: 4294967295, F1: 127, F2: 255}
                        _ = v26
                    }
                } else {
                    fmt.Println(v17, v17, v1)
                    v1 = true
                }
                if v1 {
                    v27 := GenS2{F0: uint64(v5), F1: 129, F2: 286}
                    _ = v27
                }
                _ = fi23
            }
            _ = fi19
        }
        _ = v18
    }
    for fi28 := 0; fi28 < 2; fi28++ {
        v14 -= 0.0
        v29 := GenS3{F0: -960}
        v17 -= v17
        _ = fi28
        _ = v29
    }
    delete(v3, "")
    for fi30 := 0; fi30 < 13; fi30++ {
        switch "foo bar" {
        case "":
            v6 = (v17 / ((44) | 1))
        case "a":
            v3[""] = 2147483648
        case "\n\t":
            v14 -= 0.1
        }
        _ = fi30
    }
    v31 := v3[(("\n\t" + "\n\t") + ("" + ("a" + "hello")))]
    if len(v2) > 0 {
        v2[0] = 32767
    }
    var v32 float64 = v14
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
