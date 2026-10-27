package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint }
func genFn0(p0 int8) int32 {
    v4 := GenS1{F0: 574}
    v5 := GenS1{F0: 65535}
    fmt.Println(p0)
    var v6 float32 = float32(0.0)
    var v7 string = "\n\t"
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 int32) uint64 {
    for fi8 := 0; fi8 < 11; fi8++ {
        if true {
            switch "" {
            case "\n\t":
                if (557 == 18446744073709551615) {
                    var v9 int = 128
                    var v10 uint64 = 46
                    for fi11 := 0; fi11 < 17; fi11++ {
                        switch fi8 {
                        case 849:
                            v12 := []int{}
                            _ = v12
                        case 42:
                            v13 := []int{v9, 9223372036854775807}
                            _ = v13
                        case 104:
                            v14 := GenS1{F0: 182}
                            _ = v14
                        default:
                            var v15 int8 = 60
                            _ = v15
                        }
                        switch ("foo bar" + "a") {
                        case "hello":
                            fmt.Println(fi11, fi8, v10)
                        default:
                            if false {
                                var v16 bool = true
                                v17 := []int{fi11, fi11, -2147483648, (fi11 & fi8)}
                                var v18 int = (v9 / ((v9) | 1))
                                _ = v16
                                _ = v17
                                _ = v18
                            }
                        }
                        v9 = fi8
                        var v19 float64 = 0.1
                        if true {
                            v19 = v19
                            var v20 uint16 = 1
                            _ = v20
                        }
                        _ = fi11
                        _ = v19
                    }
                    _ = v10
                    _ = v9
                } else {
                    fmt.Println(fi8, p0)
                    fmt.Println(fi8, fi8)
                }
            case "hello":
                var v21 uint16 = 255
                _ = v21
            }
            var v22 string = ((("hello" + "foo bar") + "a") + "")
            _ = v22
        } else {
            p0 = (p0 >> ((42) & 63))
            v23 := []int{(fi8 << ((779) & 63))}
            for fi24 := 0; fi24 < 13; fi24++ {
                _ = len(v23)
                _ = fi24
            }
            _ = v23
        }
        _ = fi8
    }
    p0 -= p0
    fmt.Println(p0)
    p0 -= (p0 << ((255) & 63))
    for fi25 := 0; fi25 < 14; fi25++ {
        if false {
            var v26 uint = 18446744073709551615
            fmt.Println(p0)
            _ = v26
        } else {
            fmt.Println(p0)
        }
        p0 += (p0 << ((65535) & 63))
        var v27 int8 = 42
        _ = fi25
        _ = v27
    }
    _ = p0
    return 0
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint64 = 18446744073709551615
    var v3 uint32 = 42
    if (255 <= 1) {
        v3 += 2147483648
        var v28 int = 4294967295
        fmt.Println(v2, v2, v3)
        _ = v28
    } else {
        var v29 int8 = 0
        _ = v29
    }
    fmt.Println(v3)
    if (genFn1(931) != (v2 - (v2 * (v2 ^ (v2 - v2))))) {
        for fi30 := 0; fi30 < 18; fi30++ {
            v2 = genFn1(int32(v2))
            var v31 int64 = 1
            v31 += v31
            _ = fi30
            _ = v31
        }
    } else {
        v2 += (v2 - (v2 >> ((1) & 63)))
        fmt.Println(v2, v3)
        for fi32 := 0; fi32 < 8; fi32++ {
            if true {
                switch ((("hello" + ("" + "\n\t")) + ("foo bar" + "")) + ("" + "")) {
                case "foo bar":
                    v2 += v2
                default:
                    switch "" {
                    case "\n\t":
                        var v33 uint32 = v3
                        _ = v33
                    case "":
                        var v34 int32 = -1
                        _ = v34
                    case "hello":
                        for fi35 := 0; fi35 < 9; fi35++ {
                            switch "" {
                            case "":
                                switch "hello" {
                                case "foo bar":
                                    v3 = (^(v3))
                                }
                            default:
                                v36 := GenS2{F0: 607}
                                _ = v36
                            }
                            fmt.Println(v3)
                            _ = fi35
                        }
                    }
                }
                v37 := GenS2{F0: 42}
                _ = v37
            }
            _ = fi32
        }
    }
    for fi38 := 0; fi38 < 2; fi38++ {
        for fi39 := 0; fi39 < 19; fi39++ {
            var v40 int = -2147483648
            _ = fi39
            _ = v40
        }
        _ = fi38
    }
    fmt.Println(v2, v3)
    _ = len(v1)
    v1[(v3 * 2)] = false
    var v41 bool = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v41
}
