package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint64 }
type GenS2 struct { F0 uint64; F1 int16; F2 int }

func main() {
    var v1 float32 = -2.25
    var v2 bool = false
    var v3 uint = 214
    v4 := []int{}
    v5 := make(map[uint32]uint32)
    var v6 float32 = v1
    if (v2 || true) {
        v7 := GenS1{F0: 4294967295, F1: 4294967295, F2: 2}
        var v8 float64 = 0.0
        _ = len(v5)
        _ = v7
        _ = v8
    } else {
        v6 = 3.14159
        if true {
            var v9 float32 = 0.1
            _ = v9
        } else {
            for fi10 := 0; fi10 < 17; fi10++ {
                fmt.Println(v2, fi10)
                for fi11 := 0; fi11 < 12; fi11++ {
                    var v12 int16 = 32767
                    if len(v4) > 0 {
                        v4[0] = fi10
                    }
                    _ = fi11
                    _ = v12
                }
                if (true || v2) {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    v5[4294967295] = 2147483648
                } else {
                    var v13 uint32 = 0
                    v13 = v13
                    switch (fi10 + 4294967295) {
                    case 1:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    _ = v13
                }
                _ = fi10
            }
            v14 := GenS2{F0: 0, F1: 1, F2: -217}
            _ = v14
        }
        switch ("foo bar" + ("foo bar" + "a")) {
        case "foo bar":
            if (!true) {
                var v15 int16 = 0
                if (v1 <= ((float32(1e10) - 1.5) + 1e10)) {
                    for fi16 := 0; fi16 < 8; fi16++ {
                        for fi17 := 0; fi17 < 16; fi17++ {
                            var v18 uint8 = 126
                            _ = fi17
                            _ = v18
                        }
                        fmt.Println(v2, v2)
                        _ = fi16
                    }
                    var v19 uint32 = 2147483647
                    _ = v19
                }
                if false {
                    var v20 uint32 = 65535
                    _ = v20
                }
                _ = v15
            }
        }
    }
    v1 -= 0.1
    v3 = 985
    fmt.Println(v3, v1)
    switch 9223372036854775807 {
    case 127:
        v3 = v3
    case 2147483647:
        v6 += (1e10 - 0.1)
    case 42:
        var v21 float64 = (0.1 * 1.5)
        _ = v21
    default:
        v4 = append(v4, 127)
    }
    switch "hello" {
    case "a":
        v22 := v5[1]
        _ = v22
    default:
        for fi23 := 0; fi23 < 8; fi23++ {
            var v24 uint8 = 127
            if v2 {
                var v25 uint8 = (v24 ^ 0)
                if ("foo bar" < "\n\t") {
                    v6 += (v6 + (v6 + (v6 * v1)))
                    v5[2147483647] = 0
                    for fi26 := 0; fi26 < 15; fi26++ {
                        switch "a" {
                        case "foo bar":
                            v6 = 0.0
                        default:
                            fmt.Println(v6, v2)
                        }
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = len(v5)
                        _ = fi26
                    }
                }
                v27 := v5[0]
                _ = v25
                _ = v27
            }
            _ = fi23
            _ = v24
        }
    }
    fmt.Println(v3, v6, v1)
    v2 = (((!((float64((0.1 / 3.14159)) * (0.0 / 3.14159)) != (-2.25 - 0.0))) || (324 != 0)) || (v2 || (629 > 1)))
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
