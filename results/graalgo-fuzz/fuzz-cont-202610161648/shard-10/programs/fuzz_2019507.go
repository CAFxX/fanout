package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 int16 }
func genFn0(p0 uint16, p1 uint16) int8 {
    v7 := []int{32768, -968}
    if len(v7) > 0 {
        v7[0] = 114
    }
    for fi8 := 0; fi8 < 17; fi8++ {
        var v9 uint32 = 42
        if true {
            v10 := GenS1{F0: 1, F1: 2, F2: 1}
            _ = v10
        }
        if (false || (0 < 383)) {
            for fi11 := 0; fi11 < 4; fi11++ {
                if (116 != 266) {
                    v7 = append(v7, (fi11 % (((fi8 << ((65535) & 63))) | 1)))
                }
                var v12 uint = 18446744073709551615
                fmt.Println(v9, p0, fi8)
                _ = fi11
                _ = v12
            }
        }
        _ = fi8
        _ = v9
    }
    var v13 uint32 = 1
    _ = p0
    _ = p1
    _ = v13
    _ = v7
    return 125
}

func main() {
    var v1 bool = ((float32(1e10) + ((3.14159 + 3.14159) / 3.14159)) >= 1.5)
    v2 := make(map[int]string)
    v3 := []int{}
    v4 := []uint32{}
    v5 := []string{}
    v6 := []int{1}
    switch -107 {
    case 547:
        var v14 uint = 2
        _ = v14
    case 9223372036854775807:
        for fi15 := 0; fi15 < 10; fi15++ {
            if v1 {
                var v16 string = "hello"
                var v17 bool = ((-1 != int32(fi15)) || (!(255 >= 4294967295)))
                _ = v16
                _ = v17
            }
            v18 := GenS1{F0: -1, F1: 176, F2: 42}
            v1 = ((247 != 42) || v1)
            _ = fi15
            _ = v18
        }
    case -843:
        v1 = (v1 || false)
    }
    v2[-2147483648] = ("foo bar" + "")
    v19 := v2[32768]
    v19 = ("foo bar" + v19)
    var v20 int = 32767
    for fi21 := 0; fi21 < 20; fi21++ {
        var v22 uint8 = 1
        _ = fi21
        _ = v22
    }
    v2[(v20 >> ((65535) & 63))] = v19
    var v23 bool = (!false)
    if true {
        var v24 bool = (255 == 65535)
        for fi25 := 0; fi25 < 19; fi25++ {
            if (((!v23) && (v23 && v23)) || v24) {
                v20 = fi25
                for fi26 := 0; fi26 < 12; fi26++ {
                    v27 := GenS1{F0: 1000, F1: 117, F2: 0}
                    if v1 {
                        if len(v5) > 0 {
                            v5[0] = v19
                        }
                        fmt.Println(v1, v20)
                        var v28 uint64 = 18446744073709551615
                        var v29 int8 = -45
                        _ = v28
                        _ = v29
                    } else {
                        var v30 bool = (!v23)
                        var v31 int16 = -32768
                        _ = v30
                        _ = v31
                    }
                    if false {
                        v24 = v23
                    } else {
                        var v32 bool = (65535 >= 4294967295)
                        var v33 uint64 = 2
                        _ = v32
                        _ = v33
                    }
                    _ = fi26
                    _ = v27
                }
            }
            for fi34 := 0; fi34 < 9; fi34++ {
                v2[(v20 * fi34)] = (((v19 + (v19 + v19)) + v19) + (v19 + "a"))
                switch v19 {
                case "a":
                    var v35 uint8 = 42
                    _ = v35
                default:
                    v20 = fi34
                }
                _ = fi34
            }
            _ = fi25
        }
        _ = v24
    } else {
        v20 -= v20
        for fi36 := 0; fi36 < 2; fi36++ {
            if true {
                delete(v2, -741)
            } else {
                switch (((v19 + v19) + v19) + "foo bar") {
                case "":
                    v20 = v20
                default:
                    if (42 == 423) {
                        delete(v2, (v20 % ((fi36) | 1)))
                    }
                }
                fmt.Println(v1, v1, fi36)
                var v37 uint16 = 2
                _ = v37
            }
            _ = fi36
        }
        var v38 uint64 = 592
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
