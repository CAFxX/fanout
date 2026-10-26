package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 uint32, p1 int8, p2 int) uint {
    v7 := []int{}
    fmt.Println(p2, p0, p2)
    var v8 int32 = 773
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    v1 := make(map[string]string)
    var v2 float32 = 1e10
    var v3 int16 = 0
    var v4 uint = 255
    v5 := []uint32{}
    var v6 uint8 = 1
    for fi9 := 0; fi9 < 17; fi9++ {
        v10 := v1["hello"]
        for fi11 := 0; fi11 < 5; fi11++ {
            var v12 uint8 = (^(v6))
            _ = fi11
            _ = v12
        }
        switch 32767 {
        case 128:
            var v13 uint = (v4 % (((^(v4))) | 1))
            _ = v13
        case 276:
            v5 = append(v5, 2147483648)
        case 4294967295:
            switch v10 {
            case "":
                if false {
                    var v14 uint8 = v6
                    delete(v1, ((v10 + ("foo bar" + v10)) + v10))
                    var v15 uint16 = uint16(v6)
                    _ = v14
                    _ = v15
                }
            case "foo bar":
                if (v3 >= (v3 << ((genFn0(264, -1, 0)) & 63))) {
                    switch -1 {
                    case 32767:
                        v16 := GenS1{F0: 2}
                        _ = v16
                    case 9223372036854775807:
                        var v17 bool = true
                        _ = v17
                    default:
                        v18 := GenS1{F0: 2147483648}
                        _ = v18
                    }
                } else {
                    fmt.Println(v10, fi9)
                    var v19 uint = v4
                    _ = v19
                }
            }
        }
        _ = fi9
        _ = v10
    }
    v20 := GenS1{F0: 733}
    v3 -= (v3 >> ((genFn0(65535, 8, -9223372036854775808)) & 63))
    v6 -= v6
    var v21 string = "hello"
    var v22 bool = true
    if ((v21 + (v21 + v21)) >= (((((v21 + v21) + (v21 + (("" + (v21 + v21)) + v21))) + "foo bar") + (v21 + v21)) + v21)) {
        if (!v22) {
            for fi23 := 0; fi23 < 4; fi23++ {
                fmt.Println(v4)
                _ = fi23
            }
            v24 := GenS1{F0: 239}
            _ = v24
        }
        if (2 == (v4 + 741)) {
            v6 = v6
            switch int(v4) {
            case 570:
                v2 += (float32((0.0 - 0.1)) + v2)
            case 2147483647:
                v4 += 1
            case 32768:
                var v25 uint8 = uint8(v4)
                _ = v25
            default:
                _ = len(v1)
            }
        }
    } else {
        for fi26 := 0; fi26 < 3; fi26++ {
            for fi27 := 0; fi27 < 14; fi27++ {
                v22 = false
                _ = fi27
            }
            fmt.Println(v3)
            _ = fi26
        }
        for fi28 := 0; fi28 < 4; fi28++ {
            if len(v5) > 0 {
                v5[0] = 2147483648
            }
            v29 := GenS1{F0: 2147483647}
            _ = fi28
            _ = v29
        }
        if false {
            v3 = v3
        } else {
            if len(v5) > 0 {
                v5[0] = 743
            }
        }
    }
    var v30 string = ("a" + "")
    if v22 {
        _ = len(v1)
        v21 = v21
        v3 = (v3 | v3)
    } else {
        v6 += 2
    }
    switch v21 {
    case "\n\t":
        fmt.Println(v21, v6)
    case "a":
        for fi31 := 0; fi31 < 16; fi31++ {
            if (42 < 0) {
                _ = len(v5)
                switch 4294967295 {
                case -1:
                    switch v30 {
                    case "\n\t":
                        var v32 uint8 = v6
                        _ = v32
                    default:
                        _ = len(v1)
                    }
                case 9223372036854775807:
                    var v33 int64 = -717
                    _ = v33
                }
            }
            v21 = v21
            v2 = -2.25
            _ = fi31
        }
    default:
        var v34 uint32 = 42
        _ = v34
    }
    v21 = v21
    var v35 float64 = 0.1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v30
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
