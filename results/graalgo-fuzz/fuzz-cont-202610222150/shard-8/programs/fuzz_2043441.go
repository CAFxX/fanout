package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16; F1 string }
func genFn0(p0 uint) string {
    p0 = p0
    var v6 int8 = -40
    var v7 int = 0
    switch (v7 & (v7 / ((127) | 1))) {
    case -2147483648:
        if ((!true) || true) {
            v7 -= v7
            if false {
                var v8 float64 = 3.14159
                v9 := GenS1{F0: -1, F1: (p0 ^ (p0 % ((p0) | 1))), F2: 65535}
                v6 -= 21
                _ = v8
                _ = v9
            } else {
                for fi10 := 0; fi10 < 20; fi10++ {
                    v6 += -15
                    switch (fi10 * fi10) {
                    case 0:
                        p0 -= p0
                    case 32767:
                        for fi11 := 0; fi11 < 18; fi11++ {
                            var v12 uint64 = 512
                            _ = fi11
                            _ = v12
                        }
                    }
                    var v13 int16 = -32768
                    _ = fi10
                    _ = v13
                }
                v14 := GenS1{F0: 32767, F1: p0, F2: 2}
                _ = v14
            }
        }
    case 2147483647:
        v15 := GenS1{F0: -715, F1: p0, F2: 293}
        _ = v15
    case -1:
        v16 := GenS1{F0: 42, F1: (p0 | (^(p0))), F2: 2}
        _ = v16
    default:
        var v17 int32 = -2147483648
        _ = v17
    }
    v18 := GenS1{F0: -615, F1: (^(p0)), F2: 5}
    _ = p0
    _ = v18
    _ = v6
    _ = v7
    return "\n\t"
}
func genFn1(p0 int32, p1 uint16, p2 string) bool {
    switch 2147483647 {
    case 4294967295:
        p0 += p0
    case 9223372036854775807:
        var v19 int = -626
        _ = v19
    case -355:
        for fi20 := 0; fi20 < 8; fi20++ {
            if ("\n\t" <= (p2 + "a")) {
                for fi21 := 0; fi21 < 16; fi21++ {
                    v22 := []int{(fi20 + -434), fi20, (fi21 % ((-2147483648) | 1))}
                    _ = fi21
                    _ = v22
                }
                for fi23 := 0; fi23 < 12; fi23++ {
                    if true {
                        var v24 uint = 2
                        v25 := GenS2{F0: 0}
                        _ = v24
                        _ = v25
                    } else {
                        p1 += p1
                    }
                    v26 := GenS3{F0: p1, F1: p2}
                    _ = fi23
                    _ = v26
                }
                switch (fi20 >> ((791) & 63)) {
                case 4294967295:
                    var v27 float64 = float64(float32((0.0 - 0.1)))
                    _ = v27
                case 4:
                    p0 -= (p0 - p0)
                }
            }
            _ = fi20
        }
    }
    if (!(true && true)) {
        fmt.Println(p2, p1)
        fmt.Println(p2)
        v28 := GenS1{F0: 840, F1: 42, F2: p1}
        _ = v28
    }
    _ = p0
    _ = p1
    _ = p2
    return false
}

func main() {
    var v1 string = "hello"
    v2 := []string{}
    v3 := []uint32{13}
    var v4 string = (v1 + "")
    var v5 uint8 = 127
    v1 = (v4 + (((((v1 + v1) + ("hello" + (v1 + "foo bar"))) + v1) + v4) + ""))
    if ((0.0 - 1.5) != 3.14159) {
        var v29 int16 = 833
        _ = v29
    }
    if false {
        v1 = v4
        if false {
            if ((1 < v5) || (false || true)) {
                for fi30 := 0; fi30 < 1; fi30++ {
                    var v31 int64 = 9223372036854775807
                    _ = fi30
                    _ = v31
                }
            }
        }
    }
    var v32 int = 0
    var v33 uint16 = 65535
    v34 := GenS2{F0: 75}
    var v35 int = v32
    v35 = (v35 << ((uint(v5)) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
