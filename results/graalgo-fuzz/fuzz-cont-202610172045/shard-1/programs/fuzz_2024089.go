package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 int32) int32 {
    if false {
        for fi7 := 0; fi7 < 10; fi7++ {
            if (((!true) && (876 > 18446744073709551615)) && (1 <= 0)) {
                v8 := []int{fi7}
                p0 = p0
                p0 -= p0
                _ = v8
            } else {
                fmt.Println(p0, fi7)
                v9 := GenS1{F0: 65535, F1: 0}
                _ = v9
            }
            p0 -= 2147483647
            _ = fi7
        }
    }
    var v10 int8 = -1
    for fi11 := 0; fi11 < 4; fi11++ {
        var v12 uint32 = 303
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = v10
    return -1
}

func main() {
    var v1 uint32 = 0
    var v2 int = 1
    var v3 float64 = 0.1
    var v4 uint16 = 716
    var v5 uint64 = 918
    v6 := []uint32{v1, uint32(v5), 939}
    _ = len(v6)
    if true {
        var v13 uint8 = 255
        _ = v13
    }
    v14 := GenS2{F0: (v5 - uint64(v4))}
    v1 = 42
    if ((-547 <= -676) && ((uint(v5) > 42) || ((^(v1)) == (^(v1))))) {
        v2 += (v2 & (v2 + 32767))
        switch (v2 % ((v2) | 1)) {
        case 0:
            fmt.Println(v1)
        case 32767:
            var v15 float64 = 1e10
            _ = v15
        default:
            var v16 uint32 = 2147483647
            _ = v16
        }
    }
    v1 -= (v1 + (v1 << ((352) & 63)))
    for fi17 := 0; fi17 < 18; fi17++ {
        var v18 uint16 = (v4 >> ((uint(v1)) & 63))
        if (!(("a" >= "hello") || true)) {
            if true {
                var v19 uint8 = 255
                if (4294967295 <= (v1 | 2)) {
                    switch "" {
                    case "":
                        switch (("foo bar" + "a") + ("hello" + "")) {
                        case "foo bar":
                            fmt.Println(v3, fi17, v18)
                        default:
                            if false {
                                v4 += (v4 * v18)
                            }
                        }
                    default:
                        if true {
                            if (202 != 1) {
                                fmt.Println(v2, fi17, v4)
                            }
                            for fi20 := 0; fi20 < 15; fi20++ {
                                if (true || false) {
                                    var v21 uint8 = (v19 + (v19 - 2))
                                    fmt.Println(fi17, v18)
                                    _ = v21
                                } else {
                                    fmt.Println(fi17)
                                }
                                var v22 uint64 = v5
                                _ = fi20
                                _ = v22
                            }
                        }
                    }
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                } else {
                    var v23 float64 = 0.0
                    v24 := GenS2{F0: v5}
                    _ = v23
                    _ = v24
                }
                _ = v19
            } else {
                for fi25 := 0; fi25 < 14; fi25++ {
                    switch v2 {
                    case -235:
                        v26 := GenS1{F0: 65535, F1: 1}
                        _ = v26
                    }
                    fmt.Println(v18, v2)
                    if (false || false) {
                        if true {
                            var v27 int16 = 42
                            _ = v27
                        } else {
                            _ = len(v6)
                            var v28 float32 = 3.14159
                            _ = v28
                        }
                        v5 = uint64(v4)
                    }
                    var v29 int32 = int32(fi25)
                    _ = fi25
                    _ = v29
                }
                fmt.Println(v1, fi17)
                for fi30 := 0; fi30 < 4; fi30++ {
                    switch fi30 {
                    case -412:
                        v6 = append(v6, v1)
                    }
                    if true {
                        v31 := GenS1{F0: 0, F1: 9223372036854775807}
                        _ = v31
                    } else {
                        if false {
                            v32 := GenS2{F0: v5}
                            _ = v32
                        }
                        v3 -= 0.0
                    }
                    v2 += v2
                    _ = fi30
                }
            }
            if (false && (!false)) {
                v33 := GenS1{F0: 950, F1: -309}
                v5 = v5
                var v34 string = "a"
                _ = v33
                _ = v34
            }
            var v35 uint32 = 4294967295
            _ = v35
        } else {
            for fi36 := 0; fi36 < 20; fi36++ {
                v3 += v3
                v5 -= 42
                var v37 int = 32768
                _ = fi36
                _ = v37
            }
        }
        _ = fi17
        _ = v18
    }
    if (0 >= 1) {
        v4 += uint16(v1)
        if true {
            _ = len(v6)
        }
    } else {
        _ = len(v6)
        v38 := GenS1{F0: 18446744073709551615, F1: -9223372036854775808}
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
