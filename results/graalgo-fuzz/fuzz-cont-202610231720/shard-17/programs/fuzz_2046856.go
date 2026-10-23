package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint32 }
type GenS2 struct { F0 int64 }
func genFn0(p0 uint32, p1 bool) uint8 {
    v7 := GenS1{F0: 1, F1: 137, F2: (^(p0))}
    p0 += 4294967295
    fmt.Println(p0, p0)
    p0 = (p0 % (((p0 * (^(p0)))) | 1))
    _ = p0
    _ = p1
    _ = v7
    return 42
}

func main() {
    v1 := make(map[int]int)
    var v2 int64 = 124
    var v3 uint8 = 1
    var v4 uint32 = 65535
    var v5 float64 = 0.0
    var v6 int8 = 27
    v8 := v1[-1]
    v9 := []int{v8}
    var v10 bool = ((v2 | 830) > (v2 + v2))
    switch (v8 % ((v8) | 1)) {
    case 614:
        switch "foo bar" {
        case "a":
            if v10 {
                for fi11 := 0; fi11 < 12; fi11++ {
                    v10 = v10
                    _ = fi11
                }
                for fi12 := 0; fi12 < 4; fi12++ {
                    v13 := GenS2{F0: v2}
                    _ = fi12
                    _ = v13
                }
                if (v4 > 971) {
                    v8 = (v8 ^ v8)
                    var v14 uint64 = 42
                    _ = v14
                }
            }
        case "\n\t":
            v3 -= (v3 << ((18446744073709551615) & 63))
        case "":
            fmt.Println(v2)
        default:
            var v15 uint16 = 65535
            _ = v15
        }
    }
    if len(v9) > 0 {
        v9[0] = -374
    }
    v10 = v10
    if ((v10 && (1000 <= -32768)) && (true || (!v10))) {
        for fi16 := 0; fi16 < 1; fi16++ {
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            _ = fi16
        }
        for fi17 := 0; fi17 < 9; fi17++ {
            v18 := v1[670]
            var v19 uint8 = v3
            v4 += 474
            _ = fi17
            _ = v18
            _ = v19
        }
        if true {
            var v20 int32 = -1
            v5 += v5
            _ = v20
        } else {
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            switch -372 {
            case 341:
                for fi21 := 0; fi21 < 4; fi21++ {
                    fmt.Println(v2, v2)
                    switch (fi21 * 128) {
                    case 720:
                        v4 = (v4 ^ 710)
                    case 627:
                        var v22 bool = (243 == int(v6))
                        _ = v22
                    default:
                        v2 -= -1
                    }
                    _ = fi21
                }
            case -984:
                v9 = append(v9, (v8 % ((int(v3)) | 1)))
            case 2147483647:
                for fi23 := 0; fi23 < 13; fi23++ {
                    var v24 uint8 = 129
                    var v25 uint8 = (^(v3))
                    _ = fi23
                    _ = v24
                    _ = v25
                }
            default:
                v9 = append(v9, -570)
            }
            for fi26 := 0; fi26 < 1; fi26++ {
                switch ("hello" + "") {
                case "\n\t":
                    for fi27 := 0; fi27 < 15; fi27++ {
                        if len(v9) >= 2 {
                            v9 = v9[1:len(v9)]
                        }
                        fmt.Println(v10, v5, v2)
                        var v28 int8 = 42
                        _ = fi27
                        _ = v28
                    }
                case "foo bar":
                    v8 -= (fi26 & (fi26 / ((fi26) | 1)))
                case "a":
                    var v29 uint64 = 620
                    _ = v29
                }
                v30 := GenS1{F0: fi26, F1: 18446744073709551615, F2: (v4 % (((v4 / ((v4) | 1))) | 1))}
                var v31 float32 = 1e10
                _ = fi26
                _ = v30
                _ = v31
            }
        }
    } else {
        delete(v1, v8)
    }
    v2 = (v2 * -124)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
