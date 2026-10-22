package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 uint32, p1 int16, p2 string) int64 {
    v4 := []int{-472, 127, 32768, 540}
    switch -946 {
    case 1:
        for fi5 := 0; fi5 < 2; fi5++ {
            fmt.Println(p1, fi5, fi5)
            _ = fi5
        }
    case 2147483647:
        var v6 uint8 = 112
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 9223372036854775807
}

func main() {
    v1 := make(map[string]uint32)
    var v2 string = "foo bar"
    var v3 int = -1
    v2 = (v2 + v2)
    for fi7 := 0; fi7 < 9; fi7++ {
        switch v2 {
        case "\n\t":
            fmt.Println(v2, v3)
        case "a":
            var v8 float32 = 1.5
            _ = v8
        case "":
            if false {
                for fi9 := 0; fi9 < 9; fi9++ {
                    v3 = fi9
                    v2 = (v2 + v2)
                    v3 -= 127
                    _ = fi9
                }
                if true {
                    v3 = (v3 & v3)
                    var v10 float32 = 0.0
                    _ = v10
                }
                var v11 int32 = 161
                _ = v11
            } else {
                v2 = ""
                var v12 int16 = -417
                switch v3 {
                case -9223372036854775808:
                    v12 += v12
                default:
                    v13 := GenS1{F0: 1}
                    _ = v13
                }
                _ = v12
            }
        default:
            if true {
                for fi14 := 0; fi14 < 19; fi14++ {
                    if true {
                        v15 := []int{fi7, (v3 | v3), (fi14 | fi14), 634}
                        if true {
                            var v16 string = v2
                            var v17 uint = 42
                            _ = v16
                            _ = v17
                        } else {
                            v18 := GenS1{F0: 1}
                            fmt.Println(v3, fi14)
                            _ = v18
                        }
                        v2 = "a"
                        fmt.Println(fi14, v2)
                        fmt.Println(v2)
                        _ = v15
                    } else {
                        var v19 uint32 = 2147483648
                        v20 := []int{fi14, fi7, fi7}
                        _ = v19
                        _ = v20
                    }
                    v3 = fi7
                    v2 = v2
                    _ = fi14
                }
                var v21 uint16 = 819
                _ = v21
            } else {
                v3 += (fi7 | 0)
                v2 = ("" + "a")
            }
        }
        v22 := []int{-736, fi7, (v3 | (fi7 >> ((506) & 63))), 9223372036854775807}
        _ = fi7
        _ = v22
    }
    for fi23 := 0; fi23 < 7; fi23++ {
        var v24 uint = 4294967295
        var v25 uint32 = 2
        switch 32768 {
        case 1:
            fmt.Println(v24)
        }
        _ = fi23
        _ = v24
        _ = v25
    }
    var v26 bool = (42 == 32767)
    v27 := v1[v2]
    switch 42 {
    case 2147483647:
        v26 = (0.1 >= 0.1)
    default:
        if v26 {
            var v28 int = (v3 + (v3 & (v3 + v3)))
            fmt.Println(v28)
            _ = v28
        } else {
            _ = len(v1)
            var v29 int = 1
            _ = v29
        }
    }
    v1[("foo bar" + ("hello" + v2))] = v27
    fmt.Println(v26)
    v30 := []int{v3}
    var v31 int8 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v31
}
