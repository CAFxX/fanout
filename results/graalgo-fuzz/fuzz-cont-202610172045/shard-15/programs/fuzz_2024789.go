package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
func genFn0(p0 int16, p1 int16) bool {
    var v6 float64 = 1e10
    fmt.Println(p0)
    v7 := []int{9223372036854775807, 4294967295}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return true
}
func genFn1(p0 int16, p1 int8, p2 uint64) int32 {
    var v8 uint8 = 2
    fmt.Println(p0, p1, p1)
    var v9 float32 = float32(0.1)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return -1
}

func main() {
    v1 := []string{}
    v2 := make(map[uint32]int)
    var v3 int8 = 0
    v4 := make(map[string]int)
    var v5 float32 = 1e10
    fmt.Println(v3, v5)
    switch 1 {
    case 9223372036854775807:
        for fi10 := 0; fi10 < 6; fi10++ {
            for fi11 := 0; fi11 < 20; fi11++ {
                v3 = 58
                var v12 bool = true
                v12 = v12
                _ = fi11
                _ = v12
            }
            var v13 int32 = -208
            _ = fi10
            _ = v13
        }
    default:
        for fi14 := 0; fi14 < 12; fi14++ {
            fmt.Println(v3)
            var v15 uint32 = 2147483647
            _ = fi14
            _ = v15
        }
    }
    var v16 int = 4294967295
    var v17 int = (v16 / ((v16) | 1))
    var v18 int64 = 221
    var v19 string = ""
    v1 = append(v1, "a")
    var v20 int64 = (v18 / ((42) | 1))
    switch v19 {
    case "\n\t":
        for fi21 := 0; fi21 < 5; fi21++ {
            fmt.Println(v3, v19)
            v20 += (v18 * v20)
            v22 := GenS1{F0: 0, F1: 109}
            _ = fi21
            _ = v22
        }
    case "":
        for fi23 := 0; fi23 < 7; fi23++ {
            v3 = v3
            switch (((v19 + "hello") + v19) + v19) {
            case "foo bar":
                fmt.Println(v3, v18)
            case "a":
                for fi24 := 0; fi24 < 19; fi24++ {
                    fmt.Println(v16, v3, v20)
                    var v25 uint32 = 0
                    if false {
                        if ((!((v18 / ((v18) | 1)) != (v20 >> ((255) & 63)))) && ("foo bar" > ("\n\t" + ""))) {
                            v18 += v20
                            var v26 string = (v19 + (v19 + "hello"))
                            v19 = "\n\t"
                            _ = v26
                        } else {
                            v25 = 2147483648
                        }
                        if len(v1) > 0 {
                            v1[0] = "a"
                        }
                        v4["hello"] = 1
                    } else {
                        v19 = (v19 + ("foo bar" + ((v19 + "\n\t") + "\n\t")))
                        var v27 string = "foo bar"
                        v5 -= 0.1
                        _ = v27
                    }
                    _ = fi24
                    _ = v25
                }
            case "\n\t":
                if (!false) {
                    v28 := GenS1{F0: 675, F1: 1}
                    _ = v28
                }
            }
            _ = fi23
        }
    }
    v18 += v18
    if (!false) {
        v3 -= 0
        v29 := v2[1]
        _ = v29
    }
    v17 -= (v17 >> ((4294967295) & 63))
    delete(v4, "\n\t")
    switch (v19 + v19) {
    case "a":
        var v30 int = v17
        _ = v30
    case "hello":
        v16 += 42
    default:
        fmt.Println(v18)
    }
    _ = len(v2)
    if len(v1) > 0 {
        v1[0] = ((v19 + "hello") + ("\n\t" + v19))
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
