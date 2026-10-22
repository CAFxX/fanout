package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 int8; F1 uint32; F2 string }
type GenS3 struct { F0 string; F1 uint16; F2 int }

func main() {
    v1 := make(map[string]string)
    var v2 int8 = 127
    v3 := make(map[int]int)
    v4 := []int{}
    v5 := make(map[int]uint32)
    var v6 int64 = 1
    v7 := GenS1{F0: (!(127 > 9223372036854775807)), F1: 0}
    if false {
        v6 += 0
        fmt.Println(v2, v6)
    }
    v8 := v1["hello"]
    if (((-2.25 - float32((0.0 + 1.5))) * 0.1) >= 1e10) {
        switch 127 {
        case 2147483647:
            var v9 bool = (1.5 >= 0.0)
            _ = v9
        }
    }
    if (42 >= -872) {
        if (true || true) {
            v6 += v6
        } else {
            v10 := GenS1{F0: true, F1: -9223372036854775808}
            fmt.Println(v2, v8, v8)
            _ = v10
        }
    } else {
        v2 = 100
        fmt.Println(v2, v8)
    }
    v11 := v1[v8]
    if len(v4) > 0 {
        v4[0] = 4294967295
    }
    v8 = v8
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v4 = append(v4, 2147483647)
    for fi12 := 0; fi12 < 6; fi12++ {
        if true {
            for fi13 := 0; fi13 < 10; fi13++ {
                fmt.Println(v2)
                v11 = (((v11 + v8) + "a") + (((v8 + "a") + v11) + "hello"))
                switch (fi13 * (^(fi12))) {
                case 32767:
                    v14 := v1[("hello" + "")]
                    _ = v14
                case 128:
                    if false {
                        var v15 int16 = -578
                        v15 += 972
                        v2 += v2
                        _ = v15
                    } else {
                        switch "foo bar" {
                        case "":
                            if true {
                                v16 := GenS2{F0: (v2 % (((v2 & (v2 & (-(v2))))) | 1)), F1: 4294967295, F2: v11}
                                _ = v16
                            }
                        case "a":
                            v4 = append(v4, fi13)
                        case "\n\t":
                            _ = len(v4)
                        default:
                            fmt.Println(v6)
                            var v17 string = v11
                            v18 := GenS1{F0: true, F1: fi12}
                            _ = v17
                            _ = v18
                        }
                    }
                    fmt.Println(v6, v11)
                }
                _ = fi13
            }
            var v19 uint16 = 1
            var v20 float32 = (1e10 * ((0.0 + float32(3.14159)) - (((1.5 - (float32((3.14159 / -2.25)) + 0.1)) - 3.14159) * 1.5)))
            _ = v19
            _ = v20
        } else {
            switch 868 {
            case 0:
                v21 := GenS2{F0: (^(v2)), F1: 385, F2: (("\n\t" + v11) + ("foo bar" + "\n\t"))}
                _ = v21
            }
            for fi22 := 0; fi22 < 11; fi22++ {
                if false {
                    if ((true && (-259 <= -1)) && true) {
                        if (("\n\t" + v11) != ((v11 + (v11 + v11)) + "hello")) {
                            _ = len(v5)
                            var v23 float32 = 0.1
                            switch v8 {
                            case "hello":
                                var v24 uint64 = 1
                                _ = v24
                            default:
                                v11 = v8
                            }
                            v25 := GenS2{F0: (v2 * v2), F1: 42, F2: v8}
                            switch -1 {
                            case 32767:
                                delete(v5, fi22)
                            }
                            _ = v23
                            _ = v25
                        }
                        v8 = v8
                    } else {
                        v6 -= 0
                        delete(v1, v8)
                    }
                } else {
                    var v26 int64 = (v6 | v6)
                    fmt.Println(v8)
                    _ = v26
                }
                v11 = "hello"
                v8 = ""
                _ = fi22
            }
        }
        _ = fi12
    }
    for fi27 := 0; fi27 < 5; fi27++ {
        for fi28 := 0; fi28 < 11; fi28++ {
            if (true || ((0.0 * (3.14159 / float32(0.1))) != 0.1)) {
                switch "" {
                case "":
                    _ = len(v4)
                case "foo bar":
                    v2 += (v2 >> ((4294967295) & 63))
                    fmt.Println(v11, v11, v6)
                    var v29 string = v8
                    _ = v29
                case "a":
                    var v30 int32 = -1
                    _ = v30
                default:
                    v31 := GenS3{F0: v11, F1: 255, F2: fi28}
                    _ = v31
                }
            }
            if ((!true) || (1000 == -993)) {
                v11 = (("hello" + "a") + v8)
            } else {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            _ = len(v4)
            _ = fi28
        }
        var v32 uint32 = 460
        _ = fi27
        _ = v32
    }
    for fi33 := 0; fi33 < 7; fi33++ {
        var v34 uint8 = 105
        var v35 uint16 = 180
        for fi36 := 0; fi36 < 9; fi36++ {
            v34 = v34
            v11 = v8
            var v37 int = fi33
            var v38 uint16 = 0
            _ = fi36
            _ = v37
            _ = v38
        }
        _ = fi33
        _ = v34
        _ = v35
    }
    v3[-9223372036854775808] = -633
    for fi39 := 0; fi39 < 11; fi39++ {
        switch v8 {
        case "foo bar":
            v3[(fi39 + 4294967295)] = 127
        case "\n\t":
            var v40 int16 = 0
            _ = v40
        default:
            for fi41 := 0; fi41 < 1; fi41++ {
                if (3.14159 < -2.25) {
                    fmt.Println(fi41, v11, v8)
                } else {
                    v42 := v3[(fi39 << ((65535) & 63))]
                    _ = v42
                }
                _ = fi41
            }
        }
        var v43 int = 14
        for fi44 := 0; fi44 < 19; fi44++ {
            fmt.Println(fi44, v43)
            var v45 float64 = 1.5
            _ = fi44
            _ = v45
        }
        _ = fi39
        _ = v43
    }
    v5[9223372036854775807] = 420
    if len(v4) > 0 {
        v4[0] = -777
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
