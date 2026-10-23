package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
func genFn0() int32 {
    var v5 int16 = 0
    switch 127 {
    case 42:
        if false {
            v5 = -32768
        } else {
            for fi6 := 0; fi6 < 20; fi6++ {
                if false {
                    var v7 bool = false
                    for fi8 := 0; fi8 < 9; fi8++ {
                        v5 = (v5 ^ (v5 * v5))
                        v7 = ((127 <= (fi6 << ((1) & 63))) || true)
                        var v9 uint32 = 4294967295
                        v7 = (v9 == 984)
                        if true {
                            fmt.Println(v9, v7)
                            v7 = true
                            v10 := []int{fi8}
                            var v11 uint16 = 255
                            v5 = (v5 >> ((18446744073709551615) & 63))
                            _ = v10
                            _ = v11
                        } else {
                            v9 = v9
                            v9 -= v9
                            v12 := GenS1{F0: -1, F1: 128}
                            _ = v12
                        }
                        _ = fi8
                        _ = v9
                    }
                    var v13 bool = true
                    _ = v13
                    _ = v7
                }
                v5 = 1000
                _ = fi6
            }
            var v14 uint32 = 2147483647
            v5 -= 0
            _ = v14
        }
    case 32767:
        switch 42 {
        case 32768:
            v15 := GenS1{F0: 42, F1: 128}
            _ = v15
        default:
            switch ("a" + (("a" + ("foo bar" + "")) + "foo bar")) {
            case "hello":
                var v16 uint = 894
                _ = v16
            case "\n\t":
                v5 -= v5
            default:
                v5 -= (^(v5))
            }
        }
    default:
        v5 += (v5 >> ((42) & 63))
    }
    if false {
        v17 := []int{0, -818}
        var v18 int8 = -1
        _ = len(v17)
        _ = v17
        _ = v18
    } else {
        fmt.Println(v5)
    }
    v19 := []int{}
    _ = v19
    _ = v5
    return 1
}

func main() {
    v1 := make(map[string]string)
    v2 := []int{32768, 32767, 607, 9223372036854775807}
    var v3 uint8 = 127
    var v4 uint16 = 361
    v3 -= v3
    var v20 bool = (false || (-334 == -610))
    v4 += uint16(v3)
    if false {
        for fi21 := 0; fi21 < 2; fi21++ {
            for fi22 := 0; fi22 < 2; fi22++ {
                delete(v1, "foo bar")
                _ = fi22
            }
            _ = fi21
        }
        v4 += v4
    }
    var v23 uint16 = 32767
    for fi24 := 0; fi24 < 20; fi24++ {
        var v25 int32 = -444
        var v26 string = "\n\t"
        _ = fi24
        _ = v25
        _ = v26
    }
    if (v3 < 2) {
        var v27 uint = 211
        _ = v27
    }
    delete(v1, "a")
    switch 290 {
    case 961:
        switch -2147483648 {
        case 864:
            var v28 uint64 = 4294967295
            _ = v28
        case -886:
            if false {
                fmt.Println(v4, v3)
            }
        case -9223372036854775808:
            v29 := GenS1{F0: 1, F1: v3}
            _ = v29
        default:
            if false {
                var v30 uint8 = v3
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v31 int = 260
                _ = v30
                _ = v31
            }
        }
    default:
        var v32 float32 = 0.1
        _ = v32
    }
    if v20 {
        fmt.Println(v3)
        fmt.Println(v23, v3, v3)
    }
    fmt.Println(v23, v4)
    var v33 float32 = 1.5
    v1["foo bar"] = ("a" + "foo bar")
    for fi34 := 0; fi34 < 16; fi34++ {
        for fi35 := 0; fi35 < 18; fi35++ {
            v4 = 978
            _ = fi35
        }
        delete(v1, "foo bar")
        for fi36 := 0; fi36 < 3; fi36++ {
            var v37 float32 = v33
            var v38 bool = true
            _ = fi36
            _ = v37
            _ = v38
        }
        _ = fi34
    }
    v39 := v1["a"]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v33
    _ = v39
    _ = v4
}
