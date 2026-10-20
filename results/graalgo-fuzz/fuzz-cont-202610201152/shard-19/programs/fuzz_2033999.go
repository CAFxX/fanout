package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 uint32 = 4294967295
    var v2 uint64 = 18446744073709551615
    var v3 uint64 = (v2 ^ (v2 << ((42) & 63)))
    v4 := make(map[string]bool)
    var v5 uint = 65535
    var v6 int = 4294967295
    var v7 float64 = float64(0.0)
    var v8 string = ("foo bar" + "")
    v4["foo bar"] = (true || true)
    if (((((0.1 / -2.25) + ((float32(float64((0.0 * (((1e10 + (((float32(((float64((((0.1 - float32(((v7 / 1.5) * 1e10))) + 1e10) - float32(v7))) + 0.1) - 1.5)) + 1e10) * (1.5 * (-2.25 * 1.5))) * -2.25)) / 1.0) * float32(v7))))) + (-2.25 + 1.5)) + 3.14159)) * 0.1) > (0.1 / 1.0)) || false) {
        v9 := GenS1{F0: -1}
        v4[v8] = false
        _ = v9
    } else {
        fmt.Println(v6, v2, v2)
    }
    if ((v6 != -270) || false) {
        if false {
            v2 += (v3 & (v2 - 1))
            if false {
                v10 := v4[v8]
                var v11 float32 = 0.0
                var v12 uint16 = 255
                _ = v10
                _ = v11
                _ = v12
            } else {
                if false {
                    var v13 int32 = 0
                    if (((^(v5)) != (^(v5))) && true) {
                        v1 -= v1
                    } else {
                        if false {
                            switch ("a" + ("hello" + v8)) {
                            case "foo bar":
                                if (128 == 169) {
                                    var v14 int16 = 72
                                    if (true && (!((!true) || (((float64((float32(-2.25) * 0.1)) + v7) >= (v7 - v7)) || true)))) {
                                        var v15 uint8 = 25
                                        _ = v15
                                    }
                                    var v16 int64 = 42
                                    _ = v14
                                    _ = v16
                                } else {
                                    for fi17 := 0; fi17 < 2; fi17++ {
                                        var v18 bool = (!false)
                                        _ = fi17
                                        _ = v18
                                    }
                                    if (!((v5 % (((v5 % ((0) | 1))) | 1)) == v5)) {
                                        fmt.Println(v3)
                                    }
                                    v6 = -974
                                }
                            default:
                                v13 += v13
                            }
                        }
                    }
                    switch "foo bar" {
                    case "\n\t":
                        switch (v6 | (v6 << ((65535) & 63))) {
                        case -296:
                            v19 := []int{v6, v6}
                            _ = v19
                        default:
                            if (-1 != -9) {
                                v3 -= v3
                            } else {
                                for fi20 := 0; fi20 < 14; fi20++ {
                                    v2 = (v2 + v3)
                                    for fi21 := 0; fi21 < 19; fi21++ {
                                        switch (fi21 - v6) {
                                        case -451:
                                            var v22 string = ("hello" + (v8 + "\n\t"))
                                            _ = v22
                                        case 9223372036854775807:
                                            fmt.Println(v2, fi21, v6)
                                        }
                                        _ = fi21
                                    }
                                    _ = fi20
                                }
                                v23 := v4[(v8 + v8)]
                                var v24 int16 = -1
                                _ = v23
                                _ = v24
                            }
                        }
                    case "foo bar":
                        for fi25 := 0; fi25 < 17; fi25++ {
                            for fi26 := 0; fi26 < 6; fi26++ {
                                _ = len(v4)
                                var v27 uint32 = (v1 >> (((v5 % ((4294967295) | 1))) & 63))
                                _ = fi26
                                _ = v27
                            }
                            v5 -= 255
                            _ = fi25
                        }
                    case "a":
                        var v28 uint = v5
                        _ = v28
                    }
                    _ = v13
                }
                fmt.Println(v7, v7, v6)
            }
            fmt.Println(v5)
        }
    }
    v29 := GenS1{F0: 9223372036854775807}
    v30 := GenS1{F0: 518}
    v5 += (^(v5))
    var v31 uint64 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
