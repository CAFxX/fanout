package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 bool }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 int32; F1 int32 }
type GenS4 struct { F0 int8; F1 uint16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint8; F1 bool }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]string)
    var v3 int16 = -998
    v4 := GenS1{F0: 730, F1: 42, F2: true}
    for fi5 := 0; fi5 < 1; fi5++ {
        for fi6 := 0; fi6 < 9; fi6++ {
            fmt.Println(fi6)
            _ = fi6
        }
        v3 -= (v3 & (v3 >> ((255) & 63)))
        v7 := GenS1{F0: 4294967295, F1: 132, F2: true}
        _ = fi5
        _ = v7
    }
    v8 := []int{9223372036854775807, 42}
    var v9 uint64 = 4294967295
    for fi10 := 0; fi10 < 17; fi10++ {
        for fi11 := 0; fi11 < 18; fi11++ {
            switch (-(fi11)) {
            case 9223372036854775807:
                var v12 uint = 343
                _ = v12
            }
            if true {
                if (!false) {
                    v3 = v3
                    fmt.Println(fi11, fi11, fi10)
                } else {
                    v13 := GenS1{F0: 714, F1: 2, F2: false}
                    v14 := GenS1{F0: 462, F1: 128, F2: false}
                    var v15 uint32 = 42
                    if ((v3 / ((32767) | 1)) == (v3 ^ v3)) {
                        switch fi10 {
                        case 32767:
                            fmt.Println(v3)
                        case -680:
                            v16 := GenS1{F0: 0, F1: 255, F2: false}
                            _ = v16
                        case 9223372036854775807:
                            var v17 string = "hello"
                            _ = v17
                        }
                        v18 := GenS2{F0: 1, F1: 65535}
                        var v19 int64 = 9223372036854775807
                        if (-1 <= -916) {
                            var v20 int16 = 42
                            v21 := GenS3{F0: -689, F1: 8}
                            fmt.Println(fi10)
                            var v22 uint = 14
                            _ = len(v2)
                            if (v20 > (v20 / ((v3) | 1))) {
                                var v23 uint16 = 42
                                var v24 float64 = -2.25
                                _ = v23
                                _ = v24
                            }
                            fmt.Println(fi11, v9, v20)
                            _ = v20
                            _ = v21
                            _ = v22
                        }
                        _ = v18
                        _ = v19
                    }
                    v25 := GenS3{F0: -1, F1: 1}
                    _ = v13
                    _ = v14
                    _ = v15
                    _ = v25
                }
                fmt.Println(v3)
                switch fi11 {
                case -1:
                    var v26 float64 = 1e10
                    _ = v26
                case -2147483648:
                    switch 42 {
                    case -2147483648:
                        v27 := v2[0]
                        _ = v27
                    default:
                        v8 = append(v8, (fi11 + (fi10 >> ((255) & 63))))
                    }
                default:
                    v28 := GenS2{F0: 33, F1: uint16(v9)}
                    _ = v28
                }
            }
            if true {
                v9 += 18446744073709551615
                var v29 int = fi11
                v30 := GenS4{F0: 0, F1: 421}
                v31 := GenS5{F0: 984}
                if (true || false) {
                    switch "a" {
                    case "foo bar":
                        v9 = 85
                    case "":
                        var v32 int = (v29 + fi10)
                        _ = v32
                    }
                }
                _ = v29
                _ = v30
                _ = v31
            } else {
                var v33 uint8 = 127
                _ = v33
            }
            _ = fi11
        }
        v9 += (v9 << ((1) & 63))
        v3 += (v3 | v3)
        _ = fi10
    }
    v2[432] = (("foo bar" + "") + ("foo bar" + "foo bar"))
    var v34 uint32 = 4294967295
    var v35 int = -2147483648
    for fi36 := 0; fi36 < 1; fi36++ {
        v37 := GenS6{F0: 255, F1: (true || true)}
        _ = fi36
        _ = v37
    }
    v34 = 65535
    v2[(v34 >> ((400) & 63))] = (("" + "foo bar") + "\n\t")
    v34 -= 0
    v38 := v2[559]
    v35 = (v35 | -282)
    fmt.Println(v3)
    if len(v8) > 0 {
        v8[0] = v35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v38
    _ = v4
    _ = v8
    _ = v9
}
