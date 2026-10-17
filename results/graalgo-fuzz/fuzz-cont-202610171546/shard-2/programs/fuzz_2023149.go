package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 int8; F1 int32; F2 bool }
type GenS3 struct { F0 int; F1 int16 }

func main() {
    var v1 int = 32768
    v2 := make(map[string]string)
    var v3 int8 = 0
    var v4 int64 = 492
    var v5 uint64 = 18446744073709551615
    v5 -= 763
    var v6 bool = true
    fmt.Println(v4, v3, v5)
    v7 := []int{-197, (v1 + v1), 4294967295}
    v8 := GenS1{F0: 1000, F1: 1}
    v6 = (1 >= 128)
    switch "" {
    case "a":
        for fi9 := 0; fi9 < 6; fi9++ {
            v10 := GenS1{F0: -775, F1: 65535}
            for fi11 := 0; fi11 < 16; fi11++ {
                var v12 uint8 = 0
                v13 := GenS1{F0: -187, F1: 255}
                _ = fi11
                _ = v12
                _ = v13
            }
            switch "a" {
            case "hello":
                if true {
                    for fi14 := 0; fi14 < 19; fi14++ {
                        var v15 int32 = int32(v1)
                        var v16 int8 = 0
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                    v17 := GenS1{F0: 153, F1: 98}
                    _ = v17
                } else {
                    v4 += (v4 - v4)
                }
            case "foo bar":
                v3 = (v3 >> ((4294967295) & 63))
            }
            _ = fi9
            _ = v10
        }
    case "foo bar":
        if (0 == -1) {
            if (9223372036854775807 == v4) {
                v18 := GenS1{F0: 42, F1: 0}
                _ = v18
            }
            v1 += v1
            for fi19 := 0; fi19 < 20; fi19++ {
                if ((v5 - v5) >= (v5 >> ((4294967295) & 63))) {
                    switch fi19 {
                    case 127:
                        fmt.Println(v3, fi19, v6)
                    default:
                        var v20 int16 = 744
                        _ = v20
                    }
                }
                _ = fi19
            }
        } else {
            v21 := GenS1{F0: -1, F1: 905}
            var v22 int32 = 0
            _ = v21
            _ = v22
        }
    case "hello":
        v23 := GenS2{F0: v3, F1: -1, F2: (0 >= uint8(v5))}
        _ = v23
    default:
        var v24 float64 = 1.5
        _ = v24
    }
    v6 = v6
    for fi25 := 0; fi25 < 8; fi25++ {
        for fi26 := 0; fi26 < 4; fi26++ {
            if v6 {
                delete(v2, "a")
                if false {
                    if v6 {
                        v3 -= v3
                    }
                    var v27 uint = 4294967295
                    _ = v27
                } else {
                    v5 -= (v5 ^ (v5 ^ (v5 * (v5 * (v5 | v5)))))
                    _ = len(v2)
                    if v6 {
                        v7 = append(v7, 42)
                        fmt.Println(v1, fi25)
                        var v28 int16 = 1000
                        _ = v28
                    } else {
                        v4 += v4
                        switch ("" + "foo bar") {
                        case "hello":
                            v5 = (^(v5))
                        default:
                            switch -463 {
                            case -2147483648:
                                switch (fi25 & 4294967295) {
                                case 0:
                                    var v29 int = v1
                                    _ = v29
                                }
                                var v30 int16 = 42
                                _ = v30
                            case 877:
                                v31 := GenS1{F0: 42, F1: 32767}
                                _ = v31
                            default:
                                if false {
                                    v32 := GenS2{F0: (v3 | (v3 & v3)), F1: 841, F2: (986 >= 18446744073709551615)}
                                    _ = v32
                                }
                            }
                        }
                        var v33 float32 = 3.14159
                        _ = v33
                    }
                    v34 := GenS3{F0: fi25, F1: 32767}
                    _ = v34
                }
            } else {
                delete(v2, "foo bar")
                v4 -= (v4 & v4)
                switch (("hello" + "hello") + (("a" + ("\n\t" + "\n\t")) + (("hello" + "\n\t") + "\n\t"))) {
                case "\n\t":
                    v6 = (((-1 == -32768) && (32767 >= 0)) || v6)
                }
            }
            var v35 int32 = -2147483648
            if (1 == 563) {
                var v36 uint16 = 65535
                v37 := GenS1{F0: 1, F1: 65535}
                _ = v36
                _ = v37
            }
            fmt.Println(v6)
            if ((3.14159 + (0.1 - 3.14159)) == ((0.0 + (float32(float64(0.0)) + 0.0)) + 0.0)) {
                var v38 bool = (1.5 != (3.14159 / 3.14159))
                fmt.Println(v3, v3)
                _ = v38
            }
            delete(v2, "")
            switch int(v4) {
            case 1:
                if (v6 && v6) {
                    if (((!v6) || v6) && v6) {
                        var v39 bool = v6
                        fmt.Println(v3)
                        fmt.Println(v3)
                        _ = v39
                    }
                }
            case 0:
                v6 = false
            case 128:
                var v40 uint64 = (v5 << ((0) & 63))
                _ = v40
            }
            var v41 float64 = -2.25
            _ = fi26
            _ = v35
            _ = v41
        }
        var v42 uint8 = 154
        v1 += (v1 >> ((42) & 63))
        _ = fi25
        _ = v42
    }
    for fi43 := 0; fi43 < 8; fi43++ {
        v4 -= -295
        v44 := GenS2{F0: 24, F1: -509, F2: v6}
        _ = fi43
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
