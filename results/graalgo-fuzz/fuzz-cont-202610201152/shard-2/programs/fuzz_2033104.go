package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 uint16; F1 uint64; F2 int }
type GenS3 struct { F0 int64; F1 string; F2 bool }
func genFn0(p0 bool, p1 int, p2 int32) uint64 {
    p0 = p0
    var v6 int = (p1 << ((0) & 63))
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 42
}

func main() {
    v1 := []uint32{}
    var v2 uint8 = 42
    var v3 bool = false
    v4 := make(map[int]string)
    var v5 bool = (-1 != 42)
    if v3 {
        var v7 int16 = 42
        var v8 string = "foo bar"
        if true {
            var v9 uint32 = 2
            _ = v9
        }
        _ = v7
        _ = v8
    }
    for fi10 := 0; fi10 < 7; fi10++ {
        var v11 int32 = 0
        _ = fi10
        _ = v11
    }
    var v12 int64 = -1
    var v13 int32 = int32(v12)
    switch int(v12) {
    case 1:
        var v14 uint64 = genFn0(v3, 1, (v13 << ((2) & 63)))
        _ = v14
    case 9223372036854775807:
        v5 = (!((v12 < v12) || false))
    case 4294967295:
        v1 = append(v1, 1)
    }
    v3 = false
    for fi15 := 0; fi15 < 18; fi15++ {
        for fi16 := 0; fi16 < 12; fi16++ {
            v17 := GenS1{F0: "", F1: 42}
            switch 517 {
            case 1:
                fmt.Println(fi15)
            case 0:
                v2 = v2
            default:
                if v3 {
                    if (0 > 738) {
                        if len(v1) > 0 {
                            v1[0] = 0
                        }
                        var v18 int8 = -128
                        if (!v5) {
                            if (v18 < (v18 % ((97) | 1))) {
                                var v19 int = int(v18)
                                _ = v19
                            } else {
                                v13 -= (-(v13))
                            }
                            v18 = -128
                        }
                        _ = v18
                    }
                    var v20 int8 = -128
                    _ = v20
                }
            }
            fmt.Println(fi15, v3, v12)
            _ = fi16
            _ = v17
        }
        _ = fi15
    }
    var v21 float32 = 1e10
    v22 := GenS1{F0: ((("a" + "\n\t") + "") + ((("" + "hello") + ("" + (("foo bar" + ("hello" + "a")) + "foo bar"))) + ("" + (("" + "a") + "hello")))), F1: 42}
    v23 := v4[331]
    for fi24 := 0; fi24 < 8; fi24++ {
        for fi25 := 0; fi25 < 11; fi25++ {
            v13 += v13
            v4[2147483647] = "a"
            v3 = (!false)
            _ = fi25
        }
        v26 := GenS2{F0: 433, F1: genFn0((false || (!((v12 - (v12 * -157)) != v12))), (-(fi24)), v13), F2: fi24}
        var v27 float64 = 3.14159
        _ = fi24
        _ = v26
        _ = v27
    }
    fmt.Println(v5, v3, v13)
    for fi28 := 0; fi28 < 14; fi28++ {
        v29 := GenS2{F0: 369, F1: 748, F2: 127}
        _ = fi28
        _ = v29
    }
    for fi30 := 0; fi30 < 11; fi30++ {
        if ((^(fi30)) < (fi30 ^ (-(fi30)))) {
            v13 = -50
            v5 = v5
            delete(v4, int(v12))
        } else {
            var v31 float64 = float64(float32(3.14159))
            _ = v31
        }
        _ = len(v4)
        for fi32 := 0; fi32 < 5; fi32++ {
            var v33 float32 = 0.1
            v3 = (false || true)
            if (0 <= 129) {
                switch v23 {
                case "foo bar":
                    if (float32(0.0) >= v21) {
                        v23 = (v23 + "a")
                        if true {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                    }
                    v23 = (v23 + v23)
                default:
                    v34 := GenS2{F0: 530, F1: 200, F2: (fi30 % ((fi30) | 1))}
                    _ = v34
                }
                fmt.Println(v23, v3)
                var v35 uint64 = 1
                _ = v35
            } else {
                v36 := GenS1{F0: "", F1: 65535}
                v37 := GenS1{F0: (v23 + v23), F1: 255}
                if v5 {
                    v1 = append(v1, uint32(v2))
                    v38 := GenS1{F0: ("a" + ("a" + (("foo bar" + ("a" + (((v23 + v23) + v23) + (v23 + "foo bar")))) + v23))), F1: 0}
                    v5 = false
                    _ = v38
                } else {
                    switch (v23 + v23) {
                    case "a":
                        var v39 int = fi30
                        _ = v39
                    case "hello":
                        fmt.Println(v23)
                    }
                    var v40 uint16 = 0
                    switch (fi30 ^ fi30) {
                    case 127:
                        v40 = (v40 ^ (v40 << ((0) & 63)))
                    }
                    _ = v40
                }
                _ = v36
                _ = v37
            }
            fmt.Println(v3)
            if true {
                v1 = append(v1, 42)
                v41 := GenS3{F0: v12, F1: (v23 + ("\n\t" + "a")), F2: (18446744073709551615 < uint64(v2))}
                v4[1] = (("\n\t" + v23) + v23)
                _ = v41
            } else {
                var v42 int16 = 719
                v5 = (32768 >= fi30)
                if (!true) {
                    fmt.Println(v12)
                    v2 = v2
                }
                _ = v42
            }
            _ = fi32
            _ = v33
        }
        _ = fi30
    }
    v12 += v12
    var v43 int16 = 1000
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v43
    _ = v5
}
