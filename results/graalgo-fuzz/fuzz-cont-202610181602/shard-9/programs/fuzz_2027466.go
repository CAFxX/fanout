package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16; F1 uint8; F2 int64 }
func genFn0(p0 bool, p1 uint, p2 uint8) string {
    p2 = (p2 + (p2 << (((^(p1))) & 63)))
    for fi6 := 0; fi6 < 16; fi6++ {
        p1 = (p1 % (((^(p1))) | 1))
        v7 := []int{fi6, -1}
        _ = fi6
        _ = v7
    }
    var v8 bool = p0
    v9 := GenS1{F0: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return "\n\t"
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int64 = 1
    var v3 float32 = (1.5 * -2.25)
    var v4 uint16 = 32767
    var v5 float32 = v3
    if (!false) {
        if (977 != 32767) {
            v10 := GenS1{F0: v2}
            _ = v10
        } else {
            if ((((float64(v3) <= (3.14159 + 1e10)) && ((-32768 < -32768) && (true && (!true)))) && (233 <= 127)) || true) {
                for fi11 := 0; fi11 < 17; fi11++ {
                    v12 := []int{fi11}
                    fmt.Println(v2)
                    _ = fi11
                    _ = v12
                }
                v13 := []int{-102, 1, -329, 127}
                _ = v13
            } else {
                var v14 float32 = v3
                var v15 string = "a"
                switch 42 {
                case 9223372036854775807:
                    switch v15 {
                    case "foo bar":
                        var v16 int64 = (v2 - 42)
                        _ = v16
                    case "":
                        var v17 int64 = (-(v2))
                        _ = v17
                    }
                }
                _ = v14
                _ = v15
            }
            if true {
                var v18 uint8 = 221
                _ = v18
            }
        }
        v19 := []int{-9223372036854775808, 4294967295, 15}
        _ = v19
    }
    v20 := []int{74, 32768, 151}
    switch "hello" {
    case "hello":
        if true {
            var v21 uint = 0
            fmt.Println(v4)
            var v22 int = 0
            _ = v21
            _ = v22
        } else {
            for fi23 := 0; fi23 < 10; fi23++ {
                var v24 uint8 = 255
                v3 -= 3.14159
                v4 = 2
                _ = fi23
                _ = v24
            }
            if false {
                v5 += float32(float64(v5))
                v3 += v5
            } else {
                v2 -= (^(v2))
                for fi25 := 0; fi25 < 3; fi25++ {
                    for fi26 := 0; fi26 < 8; fi26++ {
                        _ = len(v1)
                        delete(v1, "foo bar")
                        _ = fi26
                    }
                    v27 := GenS1{F0: (v2 >> ((1) & 63))}
                    v5 = v5
                    _ = fi25
                    _ = v27
                }
            }
        }
    default:
        v28 := GenS1{F0: v2}
        _ = v28
    }
    var v29 uint16 = v4
    v30 := GenS2{F0: 42}
    if len(v20) >= 2 {
        v20 = v20[1:len(v20)]
    }
    v29 = (v29 ^ v4)
    fmt.Println(v29, v5)
    for fi31 := 0; fi31 < 17; fi31++ {
        v32 := GenS2{F0: 915}
        for fi33 := 0; fi33 < 1; fi33++ {
            var v34 float32 = v5
            if true {
                var v35 int = (^(fi33))
                v29 = (v4 * (^(v29)))
                if false {
                    _ = len(v20)
                }
                _ = v35
            } else {
                var v36 float64 = (0.1 - 0.0)
                v37 := GenS3{F0: v4, F1: 0, F2: 42}
                v20 = append(v20, (fi33 | (fi31 & fi33)))
                _ = v36
                _ = v37
            }
            if (("foo bar" + "a") > "hello") {
                var v38 uint = 255
                _ = v38
            } else {
                var v39 int = (fi31 | (fi31 << ((0) & 63)))
                fmt.Println(v2, v39, v3)
                switch "\n\t" {
                case "hello":
                    switch "" {
                    case "\n\t":
                        var v40 uint16 = (v4 & v29)
                        _ = v40
                    case "":
                        v3 = v34
                    case "a":
                        fmt.Println(v2, v5)
                    default:
                        v5 = (v34 / v5)
                    }
                case "\n\t":
                    if true {
                        v2 = (v2 & v2)
                        v4 = 572
                        v1["a"] = 2147483648
                    }
                default:
                    var v41 float32 = 0.0
                    _ = v41
                }
                _ = v39
            }
            _ = fi33
            _ = v34
        }
        var v42 int64 = (v2 - 9223372036854775807)
        _ = fi31
        _ = v32
        _ = v42
    }
    var v43 int = -772
    v44 := GenS2{F0: 2}
    if len(v20) > 0 {
        v20[0] = v43
    }
    v29 = (v4 << ((4294967295) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v43
    _ = v44
    _ = v5
}
