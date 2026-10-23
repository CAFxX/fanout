package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int }

func main() {
    var v1 uint8 = 42
    var v2 int16 = 1000
    v3 := make(map[string]bool)
    v4 := make(map[string]int)
    var v5 uint = 18446744073709551615
    v6 := make(map[uint32]string)
    v7 := []int{4294967295, 42, 0}
    var v8 int16 = (v2 & (-(v2)))
    v5 += (v5 / (((v5 / ((v5) | 1))) | 1))
    var v9 int8 = -1
    var v10 int64 = -678
    var v11 string = (((((("foo bar" + "") + (("\n\t" + "foo bar") + ("hello" + "\n\t"))) + (((("" + "") + "") + "hello") + "\n\t")) + "") + "\n\t") + "\n\t")
    for fi12 := 0; fi12 < 16; fi12++ {
        if true {
            var v13 uint64 = 4294967295
            fmt.Println(v1)
            _ = v13
        }
        _ = fi12
    }
    var v14 uint32 = 0
    var v15 bool = (true && (280 > 296))
    if (!true) {
        var v16 int8 = (-(v9))
        v17 := GenS1{F0: uint16(v1)}
        _ = v16
        _ = v17
    } else {
        var v18 uint32 = (v14 >> ((v5) & 63))
        var v19 int = 0
        _ = len(v7)
        _ = v18
        _ = v19
    }
    if v15 {
        v14 += (v14 - 2147483648)
    }
    if ((v9 & v9) > 42) {
        var v20 string = "hello"
        if v15 {
            if (2 != 32767) {
                v9 += v9
            } else {
                var v21 uint16 = 2
                if false {
                    v9 -= 0
                } else {
                    for fi22 := 0; fi22 < 8; fi22++ {
                        fmt.Println(v1, v11)
                        _ = fi22
                    }
                }
                var v23 int16 = v8
                _ = v21
                _ = v23
            }
            var v24 float64 = float64(0.0)
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            _ = v24
        }
        _ = v20
    }
    v9 = (v9 << (((v5 % (((^(v5))) | 1))) & 63))
    var v25 int64 = (v10 | (v10 ^ 9223372036854775807))
    if ((-2.25 / 1.5) == 0.1) {
        v15 = v15
    }
    switch 0 {
    case 32768:
        v26 := GenS1{F0: 42}
        _ = v26
    case 9223372036854775807:
        fmt.Println(v10, v15, v5)
    case -486:
        for fi27 := 0; fi27 < 11; fi27++ {
            switch (("\n\t" + v11) + "hello") {
            case "hello":
                if (!(!(1.5 <= 1e10))) {
                    if v15 {
                        if (2 > 4294967295) {
                            if ((!(true || v15)) && (!true)) {
                                v11 = "hello"
                            } else {
                                var v28 uint16 = 0
                                _ = v28
                            }
                            v29 := GenS2{F0: (fi27 & (fi27 & fi27))}
                            switch (fi27 % ((fi27) | 1)) {
                            case 2147483647:
                                if v15 {
                                    var v30 float64 = -2.25
                                    _ = v30
                                }
                            }
                            _ = v29
                        } else {
                            switch 32767 {
                            case -9223372036854775808:
                                var v31 int8 = 0
                                _ = v31
                            case 4294967295:
                                var v32 int16 = v2
                                _ = v32
                            }
                            var v33 uint8 = (^(v1))
                            if (1e10 <= ((-2.25 - 1.5) + ((float32(0.1) * 3.14159) + 3.14159))) {
                                var v34 string = (v11 + v11)
                                v1 = (v33 * 127)
                                _ = v34
                            }
                            _ = v33
                        }
                    } else {
                        var v35 int16 = (v2 & 42)
                        v36 := GenS2{F0: 0}
                        _ = v35
                        _ = v36
                    }
                    v11 = "a"
                    fmt.Println(v10)
                } else {
                    switch 9223372036854775807 {
                    case 4294967295:
                        for fi37 := 0; fi37 < 20; fi37++ {
                            v4[((("" + v11) + "") + "")] = (fi27 << (((v5 >> ((v5) & 63))) & 63))
                            v7 = append(v7, fi37)
                            switch v11 {
                            case "":
                                var v38 int = 1
                                _ = v38
                            case "hello":
                                if v15 {
                                    if len(v7) > 0 {
                                        v7[0] = (fi27 - (fi27 / (((fi27 | 32767)) | 1)))
                                    }
                                } else {
                                    v10 += (^(v10))
                                }
                            case "a":
                                fmt.Println(v5)
                            default:
                                if len(v7) > 0 {
                                    v7[0] = (fi37 >> ((754) & 63))
                                }
                            }
                            _ = fi37
                        }
                    case 128:
                        fmt.Println(fi27, v15)
                    }
                    if (255 <= 32767) {
                        if (false || (false || v15)) {
                            fmt.Println(v5)
                            var v39 uint = (v5 + 14)
                            v7 = append(v7, -313)
                            _ = v39
                        }
                        v10 = 42
                    }
                    if v15 {
                        for fi40 := 0; fi40 < 1; fi40++ {
                            var v41 bool = v15
                            _ = fi40
                            _ = v41
                        }
                    }
                }
            }
            var v42 float64 = (float64(3.14159) - 0.0)
            _ = fi27
            _ = v42
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
