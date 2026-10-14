package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 int32) int64 {
    p0 = (^(p0))
    if (127 != -36) {
        var v6 uint16 = 32767
        p0 -= 0
        _ = v6
    } else {
        p0 = (p0 % ((p0) | 1))
    }
    p0 += p0
    _ = p0
    return -549
}

func main() {
    var v1 uint32 = 42
    var v2 int16 = 0
    v3 := []string{("\n\t" + "\n\t"), "foo bar"}
    var v4 string = "\n\t"
    var v5 uint32 = v1
    var v7 uint16 = 42
    switch ((v4 + "a") + "\n\t") {
    case "hello":
        var v8 uint = 2
        _ = v8
    default:
        for fi9 := 0; fi9 < 15; fi9++ {
            v10 := GenS1{F0: 512}
            _ = fi9
            _ = v10
        }
    }
    if (25 < 0) {
        fmt.Println(v4, v2)
        switch -1 {
        case 1:
            for fi11 := 0; fi11 < 8; fi11++ {
                for fi12 := 0; fi12 < 13; fi12++ {
                    v13 := GenS1{F0: 0}
                    v1 += v5
                    _ = len(v3)
                    _ = fi12
                    _ = v13
                }
                if false {
                    var v14 float64 = -2.25
                    _ = v14
                } else {
                    v7 = 32767
                    var v15 uint = 4294967295
                    if len(v3) > 0 {
                        v3[0] = "\n\t"
                    }
                    _ = v15
                }
                switch (^(fi11)) {
                case -9223372036854775808:
                    var v16 float32 = ((1.5 / 1.0) * (1e10 - 1e10))
                    _ = v16
                default:
                    for fi17 := 0; fi17 < 6; fi17++ {
                        v2 = v2
                        _ = fi17
                    }
                }
                _ = fi11
            }
        default:
            for fi18 := 0; fi18 < 16; fi18++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi18
            }
        }
    }
    v7 = (v7 / ((v7) | 1))
    if (((!(!false)) && (false || (128 < 0))) && ((v2 & (v2 * 1000)) > v2)) {
        for fi19 := 0; fi19 < 18; fi19++ {
            for fi20 := 0; fi20 < 20; fi20++ {
                var v21 int = (-(fi20))
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                var v22 uint32 = v1
                _ = fi20
                _ = v21
                _ = v22
            }
            fmt.Println(v5, v2, v1)
            _ = fi19
        }
        var v23 int32 = 0
        _ = v23
    } else {
        v3 = append(v3, (v4 + v4))
    }
    for fi24 := 0; fi24 < 9; fi24++ {
        if false {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            var v25 int16 = 1000
            _ = v25
        }
        for fi26 := 0; fi26 < 2; fi26++ {
            var v27 int = fi26
            _ = fi26
            _ = v27
        }
        if true {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if true {
                var v28 uint32 = (^(v1))
                _ = v28
            } else {
                var v29 int16 = (v2 << ((65535) & 63))
                var v30 bool = false
                _ = v29
                _ = v30
            }
            if (1 < 0) {
                v1 += v5
            }
        } else {
            var v31 bool = (127 != -128)
            v7 += 947
            _ = v31
        }
        _ = fi24
    }
    var v32 int64 = -9223372036854775808
    var v33 int64 = (-(v32))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v7
}
