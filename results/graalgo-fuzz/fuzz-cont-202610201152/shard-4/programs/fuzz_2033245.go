package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 uint16 }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint64) uint64 {
    var v5 float32 = 0.0
    v5 += 0.1
    for fi6 := 0; fi6 < 7; fi6++ {
        if ((-806 == 42) || (993 != -124)) {
            var v7 uint = 65535
            _ = v7
        }
        _ = fi6
    }
    p0 = (p0 / ((2) | 1))
    _ = p0
    _ = v5
    return p0
}
func genFn1(p0 string) string {
    var v8 int64 = 0
    var v9 int16 = -1
    _ = p0
    _ = v8
    _ = v9
    return (p0 + p0)
}

func main() {
    v1 := make(map[int]int)
    v2 := []int{}
    v3 := []int{32767, -917}
    v4 := make(map[int]string)
    for fi10 := 0; fi10 < 6; fi10++ {
        if true {
            delete(v4, (fi10 * fi10))
        }
        _ = fi10
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        delete(v4, -515)
        var v12 uint64 = 630
        for fi13 := 0; fi13 < 5; fi13++ {
            if ((true || (false && true)) || (!false)) {
                fmt.Println(fi11, v12)
                v12 -= genFn0(v12)
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = len(v2)
                v12 -= (^(v12))
                fmt.Println(v12)
            } else {
                var v14 bool = (!("hello" >= "foo bar"))
                fmt.Println(v14, v12)
                _ = v14
            }
            v15 := GenS1{F0: 28, F1: 2, F2: 327}
            _ = fi13
            _ = v15
        }
        _ = fi11
        _ = v12
    }
    var v16 string = ("" + "foo bar")
    v17 := v1[-9223372036854775808]
    for fi18 := 0; fi18 < 17; fi18++ {
        if true {
            v17 = fi18
            v19 := GenS1{F0: 4294967295, F1: genFn0(136), F2: 65535}
            v17 = fi18
            _ = v19
        } else {
            v20 := v4[fi18]
            var v21 uint32 = 4294967295
            _ = v20
            _ = v21
        }
        var v22 int = (v17 << ((563) & 63))
        v23 := v4[v22]
        _ = fi18
        _ = v22
        _ = v23
    }
    for fi24 := 0; fi24 < 7; fi24++ {
        v25 := GenS1{F0: 465, F1: 2, F2: 65535}
        var v26 int32 = 789
        v26 -= 1
        _ = fi24
        _ = v25
        _ = v26
    }
    switch (("a" + v16) + v16) {
    case "foo bar":
        fmt.Println(v17, v16)
    case "\n\t":
        var v27 uint16 = 0
        _ = v27
    default:
        fmt.Println(v17)
    }
    v16 = v16
    if false {
        v1[(v17 % ((127) | 1))] = -909
    } else {
        v17 += (v17 / ((128) | 1))
        switch -1 {
        case 521:
            var v28 int64 = 1
            _ = v28
        case 9223372036854775807:
            v29 := GenS2{F0: 1}
            _ = v29
        default:
            var v30 int8 = 100
            _ = v30
        }
        for fi31 := 0; fi31 < 12; fi31++ {
            if (!false) {
                v2 = append(v2, fi31)
                var v32 float64 = float64(1e10)
                for fi33 := 0; fi33 < 8; fi33++ {
                    switch ("" + "hello") {
                    case "a":
                        if (true || true) {
                            var v34 int32 = 0
                            v32 = (((v32 * v32) / (1.5 * float64(-2.25))) + 1.5)
                            _ = v34
                        } else {
                            var v35 bool = false
                            _ = v35
                        }
                    case "\n\t":
                        v32 = float64(0.1)
                    default:
                        var v36 int = -1
                        _ = v36
                    }
                    var v37 string = v16
                    _ = fi33
                    _ = v37
                }
                _ = v32
            } else {
                if true {
                    if true {
                        v38 := GenS1{F0: 1, F1: 18446744073709551615, F2: 65535}
                        v39 := GenS1{F0: 65535, F1: genFn0(42), F2: 144}
                        fmt.Println(v17, v17)
                        _ = v38
                        _ = v39
                    }
                    v3 = append(v3, 417)
                    v16 = (v16 + "\n\t")
                }
            }
            for fi40 := 0; fi40 < 1; fi40++ {
                v17 = (fi31 + 4294967295)
                v41 := GenS2{F0: -605}
                _ = fi40
                _ = v41
            }
            if true {
                v17 = v17
            }
            _ = fi31
        }
    }
    v42 := GenS1{F0: 42, F1: 4294967295, F2: 255}
    if (4294967295 > 0) {
        switch ("\n\t" + ("foo bar" + (v16 + (v16 + (v16 + v16))))) {
        case "a":
            delete(v1, v17)
        }
    } else {
        v16 = "\n\t"
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v42
}
