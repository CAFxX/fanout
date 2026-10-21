package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16 }

func main() {
    v1 := make(map[int]int)
    var v2 int8 = -128
    v3 := make(map[uint32]string)
    var v4 bool = true
    switch 32768 {
    case 273:
        for fi5 := 0; fi5 < 6; fi5++ {
            v2 = (^(v2))
            _ = fi5
        }
    case 1:
        v6 := GenS1{F0: 32767}
        _ = v6
    }
    for fi7 := 0; fi7 < 17; fi7++ {
        if ("foo bar" >= "foo bar") {
            if v4 {
                fmt.Println(v2, v4)
            }
        }
        if (717 >= 1) {
            delete(v3, 979)
            v2 = v2
        }
        _ = fi7
    }
    var v8 int16 = 32767
    var v9 float32 = 1e10
    if (((220 > 100) || v4) && ((false || (-1 > 0)) || v4)) {
        var v10 int16 = 0
        fmt.Println(v4)
        var v11 int = 0
        _ = v10
        _ = v11
    } else {
        if (!v4) {
            v8 -= (v8 >> ((4294967295) & 63))
            v12 := []int{2147483647, 0, -285, -9223372036854775808}
            var v13 int32 = 42
            _ = v12
            _ = v13
        } else {
            v2 += (v2 << ((1) & 63))
            v8 += -1
        }
        for fi14 := 0; fi14 < 2; fi14++ {
            var v15 uint32 = 42
            _ = fi14
            _ = v15
        }
    }
    v2 = (v2 / ((-5) | 1))
    var v16 int32 = 42
    if v4 {
        if v4 {
            v8 = (v8 | (v8 >> ((255) & 63)))
            var v17 uint64 = 54
            for fi18 := 0; fi18 < 3; fi18++ {
                switch 1 {
                case 51:
                    v8 = v8
                case 32767:
                    fmt.Println(v8, v9, v16)
                case -2147483648:
                    if v4 {
                        v19 := []int{(fi18 ^ int(v16))}
                        v19 = append(v19, fi18)
                        switch "" {
                        case "a":
                            var v20 int = fi18
                            _ = v20
                        }
                        _ = v19
                    } else {
                        if (v16 <= v16) {
                            switch 2147483647 {
                            case -9223372036854775808:
                                v17 -= v17
                            case 346:
                                delete(v3, 2147483648)
                            case -1:
                                var v21 uint = 0
                                _ = v21
                            default:
                                v16 = 0
                            }
                            _ = len(v1)
                        }
                        fmt.Println(v9)
                        switch fi18 {
                        case -9223372036854775808:
                            var v22 float64 = 0.0
                            _ = v22
                        case 42:
                            v23 := GenS1{F0: 42}
                            _ = v23
                        default:
                            v2 = v2
                        }
                    }
                }
                _ = fi18
            }
            _ = v17
        } else {
            if v4 {
                fmt.Println(v2, v8)
            }
        }
        v24 := v1[32767]
        _ = v24
    }
    var v25 uint64 = 2
    for fi26 := 0; fi26 < 18; fi26++ {
        v27 := GenS1{F0: 700}
        v28 := GenS1{F0: 32767}
        _ = fi26
        _ = v27
        _ = v28
    }
    var v29 uint16 = 1
    fmt.Println(v4, v4)
    var v30 int8 = v2
    v31 := v3[323]
    var v32 uint64 = (v25 | v25)
    var v33 uint16 = (v29 - v29)
    switch v31 {
    case "\n\t":
        switch 32768 {
        case 4294967295:
            _ = len(v1)
        case 667:
            v34 := GenS1{F0: 681}
            _ = v34
        case 198:
            var v35 uint = 2
            _ = v35
        }
    case "hello":
        for fi36 := 0; fi36 < 15; fi36++ {
            switch ("\n\t" + v31) {
            case "":
                var v37 int32 = (^(v16))
                _ = v37
            default:
                v38 := GenS2{F0: 1}
                _ = v38
            }
            fmt.Println(v8, v29)
            v4 = v4
            _ = fi36
        }
    case "":
        var v39 string = v31
        _ = v39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v8
    _ = v9
}
