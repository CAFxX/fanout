package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 int32 }
func genFn0(p0 uint32, p1 uint16, p2 int8) uint {
    for fi4 := 0; fi4 < 20; fi4++ {
        if false {
            v5 := GenS1{F0: (p0 >> ((934) & 63)), F1: 1}
            var v6 uint16 = 839
            for fi7 := 0; fi7 < 20; fi7++ {
                fmt.Println(p1)
                fmt.Println(p0, p1, fi4)
                var v8 uint16 = v6
                _ = fi7
                _ = v8
            }
            _ = v5
            _ = v6
        } else {
            var v9 uint8 = 42
            for fi10 := 0; fi10 < 6; fi10++ {
                if (!((p2 * p2) == p2)) {
                    p2 = int8(p0)
                    if (42 == 930) {
                        var v11 uint32 = 2147483648
                        _ = v11
                    } else {
                        var v12 uint32 = uint32(v9)
                        _ = v12
                    }
                    var v13 uint = 255
                    _ = v13
                } else {
                    v14 := []int{fi10, fi4, fi4, (fi10 << ((18446744073709551615) & 63))}
                    var v15 uint8 = (^(v9))
                    switch "hello" {
                    case "a":
                        var v16 int16 = 42
                        _ = v16
                    default:
                        var v17 int64 = 321
                        _ = v17
                    }
                    var v18 int16 = -32768
                    _ = v14
                    _ = v15
                    _ = v18
                }
                v9 += (v9 & (v9 ^ 202))
                if false {
                    switch (fi4 << ((42) & 63)) {
                    case 128:
                        p0 = p0
                    default:
                        if false {
                            p2 = p2
                            v19 := GenS2{F0: 0}
                            _ = v19
                        }
                    }
                    v20 := GenS2{F0: -1}
                    _ = v20
                }
                _ = fi10
            }
            var v21 uint = 234
            _ = v21
            _ = v9
        }
        v22 := []int{fi4, -270}
        p0 += 42
        _ = fi4
        _ = v22
    }
    v23 := GenS1{F0: 2147483647, F1: 127}
    p1 = 0
    p2 += (p2 & p2)
    if true {
        var v24 float32 = 0.1
        v25 := []int{2147483647, 4294967295, 1}
        _ = v24
        _ = v25
    } else {
        p0 = (^(p0))
        var v26 int16 = -537
        _ = v26
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    return 1
}

func main() {
    v1 := []uint32{42}
    v2 := make(map[int]int)
    var v3 string = ("" + "a")
    v3 = (v3 + v3)
    _ = len(v2)
    if true {
        switch v3 {
        case "a":
            delete(v2, -375)
        case "foo bar":
            var v27 uint8 = 42
            _ = v27
        case "hello":
            switch -1 {
            case 2147483647:
                if ((!(!(!(false && ((88 == 128) || ((((!false) || (-138 == 42)) && false) || false)))))) || false) {
                    switch v3 {
                    case "hello":
                        v3 = v3
                    case "a":
                        fmt.Println(v3)
                    }
                } else {
                    var v28 float32 = 3.14159
                    for fi29 := 0; fi29 < 14; fi29++ {
                        for fi30 := 0; fi30 < 5; fi30++ {
                            v3 = (v3 + "a")
                            v31 := v2[fi29]
                            fmt.Println(v3, v31, v28)
                            v3 = ("a" + "\n\t")
                            delete(v2, (fi30 ^ v31))
                            _ = fi30
                            _ = v31
                        }
                        _ = len(v1)
                        _ = fi29
                    }
                    _ = v28
                }
            default:
                var v32 int = 834
                _ = v32
            }
        default:
            v1 = append(v1, 2147483647)
        }
        v3 = v3
        if false {
            switch 4294967295 {
            case 32767:
                v3 = v3
            case 0:
                v3 = v3
            case 4294967295:
                if false {
                    v33 := GenS1{F0: 2147483647, F1: 131}
                    _ = v33
                }
            default:
                switch 42 {
                case -933:
                    for fi34 := 0; fi34 < 9; fi34++ {
                        v3 = (v3 + v3)
                        _ = fi34
                    }
                default:
                    var v35 int16 = 1
                    _ = v35
                }
            }
            v36 := GenS2{F0: 2147483647}
            _ = v36
        } else {
            switch (v3 + ("" + v3)) {
            case "\n\t":
                fmt.Println(v3)
            default:
                v3 = (v3 + "hello")
            }
        }
    } else {
        if (311 < 65535) {
            var v37 string = (v3 + "")
            v37 = ""
            var v38 bool = false
            _ = v37
            _ = v38
        }
        var v39 int = 42
        fmt.Println(v39, v3)
        _ = v39
    }
    delete(v2, 2147483647)
    v3 = v3
    v3 = v3
    for fi40 := 0; fi40 < 6; fi40++ {
        switch ("hello" + "hello") {
        case "\n\t":
            if len(v1) > 0 {
                v1[0] = 65535
            }
        }
        _ = fi40
    }
    v3 = "hello"
    var v41 uint32 = 454
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v41
}
