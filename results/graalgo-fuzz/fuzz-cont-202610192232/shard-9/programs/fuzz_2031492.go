package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 int, p1 uint64, p2 int8) uint64 {
    var v7 int8 = p2
    if false {
        v8 := []int{42, (p0 & p0), -2147483648, 358}
        fmt.Println(p2, v7, p2)
        _ = v8
    } else {
        fmt.Println(p1, p0)
        v9 := GenS1{F0: p1}
        _ = v9
    }
    for fi10 := 0; fi10 < 13; fi10++ {
        var v11 bool = (9223372036854775807 > 42)
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p1
}

func main() {
    v1 := []uint32{}
    var v2 uint8 = 42
    var v3 int8 = -1
    var v4 int16 = 42
    var v5 int32 = 0
    var v6 uint32 = 2147483647
    var v12 int8 = v3
    fmt.Println(v4)
    var v13 int64 = 1
    if false {
        v3 = -128
    } else {
        for fi14 := 0; fi14 < 6; fi14++ {
            if false {
                var v15 uint16 = 0
                _ = v15
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                switch ((("\n\t" + ("\n\t" + (("\n\t" + "hello") + ""))) + "foo bar") + ((("foo bar" + "hello") + (("\n\t" + "a") + "\n\t")) + ("hello" + "a"))) {
                case "":
                    var v16 float32 = 1e10
                    _ = v16
                }
                v6 += (v6 ^ v6)
            }
            if len(v1) > 0 {
                v1[0] = (v6 * 2147483648)
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi14
        }
    }
    switch "" {
    case "a":
        var v17 bool = false
        _ = v17
    case "foo bar":
        var v18 int32 = (v5 & (v5 >> ((708) & 63)))
        _ = v18
    }
    v1 = append(v1, (v6 ^ 2147483648))
    v1 = append(v1, (v6 % ((uint32(v2)) | 1)))
    if (128 != 80) {
        if true {
            switch 0 {
            case -9223372036854775808:
                var v19 uint16 = 255
                _ = v19
            case 0:
                v1 = append(v1, (^(v6)))
            }
            if true {
                v4 = 407
                v3 = (v3 % (((v12 << ((2) & 63))) | 1))
            } else {
                fmt.Println(v5, v6, v5)
                if true {
                    v20 := GenS1{F0: 42}
                    _ = v20
                }
                var v21 float64 = -2.25
                _ = v21
            }
            switch -428 {
            case 4294967295:
                fmt.Println(v3)
            case 1:
                v1 = append(v1, (v6 ^ 65535))
            case 32767:
                var v22 int32 = v5
                _ = v22
            default:
                var v23 uint = 18446744073709551615
                _ = v23
            }
        } else {
            for fi24 := 0; fi24 < 19; fi24++ {
                if (v13 != v13) {
                    switch fi24 {
                    case 4294967295:
                        if false {
                            var v25 int = -9223372036854775808
                            _ = v25
                        }
                    case 32768:
                        v4 = -32768
                    }
                    if len(v1) > 0 {
                        v1[0] = v6
                    }
                    var v26 float32 = float32(0.1)
                    _ = v26
                }
                fmt.Println(v5, v2)
                _ = fi24
            }
            var v27 bool = true
            for fi28 := 0; fi28 < 18; fi28++ {
                if (0.1 >= 0.0) {
                    _ = len(v1)
                    if len(v1) > 0 {
                        v1[0] = 2147483647
                    }
                    v1 = append(v1, (v6 << ((804) & 63)))
                }
                _ = fi28
            }
            _ = v27
        }
        v29 := GenS1{F0: 18446744073709551615}
        var v30 int32 = -2147483648
        _ = v29
        _ = v30
    } else {
        fmt.Println(v12, v6, v12)
        v2 = (v2 >> ((4294967295) & 63))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
