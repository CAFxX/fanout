package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0() int8 {
    v4 := GenS1{F0: 4294967295}
    v5 := GenS1{F0: 42}
    if (100 >= 100) {
        var v6 float32 = -2.25
        _ = v6
    } else {
        var v7 string = ("\n\t" + (("foo bar" + ("foo bar" + "a")) + "hello"))
        _ = v7
    }
    var v8 int32 = 247
    v8 -= 403
    _ = v4
    _ = v5
    _ = v8
    return -1
}
func genFn1(p0 int8) string {
    var v9 uint16 = 65535
    fmt.Println(p0, v9)
    _ = p0
    _ = v9
    return ("\n\t" + "a")
}

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[string]string)
    var v3 uint = 4294967295
    v10 := []int{-203, 0}
    v3 = v3
    v1 = -47
    var v11 int8 = -112
    if false {
        v11 = (v11 - genFn0())
        var v12 float32 = -2.25
        v3 = 4294967295
        _ = v12
    } else {
        v1 = 989
        switch "\n\t" {
        case "":
            var v13 int32 = (v1 % ((v1) | 1))
            _ = v13
        case "a":
            var v14 int8 = (v11 ^ v11)
            _ = v14
        default:
            fmt.Println(v1, v1, v11)
        }
        var v15 bool = ((-2.25 + 1e10) <= 1e10)
        _ = v15
    }
    var v16 int = -1
    for fi17 := 0; fi17 < 15; fi17++ {
        fmt.Println(v11, v11, fi17)
        if (42 != -924) {
            switch 50 {
            case 32768:
                var v18 int32 = -2147483648
                _ = v18
            case 42:
                if (-12 != genFn0()) {
                    _ = len(v2)
                } else {
                    v1 += (v1 & (v1 & v1))
                    v3 += (v3 ^ v3)
                }
            default:
                v1 = (v1 & -2147483648)
            }
            v19 := GenS1{F0: 2}
            _ = v19
        } else {
            fmt.Println(v16, v16, v11)
            fmt.Println(fi17, v11, fi17)
        }
        _ = fi17
    }
    if (!false) {
        fmt.Println(v11)
        v11 = (^(v11))
        var v20 uint8 = 127
        _ = v20
    } else {
        for fi21 := 0; fi21 < 11; fi21++ {
            v10 = append(v10, int(v11))
            _ = fi21
        }
        delete(v2, (("\n\t" + "") + "\n\t"))
    }
    var v22 float32 = 0.0
    var v23 uint8 = 182
    v11 = v11
    if false {
        v23 -= (^(v23))
        if true {
            var v24 int16 = 0
            _ = v24
        } else {
            fmt.Println(v16, v3)
            if false {
                delete(v2, "foo bar")
                switch -1 {
                case 509:
                    var v25 float64 = 0.0
                    _ = v25
                default:
                    if false {
                        v2["\n\t"] = "a"
                        for fi26 := 0; fi26 < 1; fi26++ {
                            v22 += v22
                            _ = fi26
                        }
                        v23 += v23
                    }
                }
                switch v16 {
                case 4294967295:
                    var v27 uint = v3
                    _ = v27
                case 127:
                    fmt.Println(v16, v1, v23)
                case 0:
                    v28 := GenS1{F0: 2}
                    _ = v28
                }
            }
            fmt.Println(v11, v22)
        }
        var v29 uint16 = 564
        _ = v29
    } else {
        v30 := GenS1{F0: 1}
        v3 -= 4294967295
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v3
}
