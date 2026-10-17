package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint16 }
type GenS2 struct { F0 uint64; F1 uint; F2 int8 }
type GenS3 struct { F0 int32; F1 int8 }
func genFn0(p0 bool) uint64 {
    if p0 {
        var v7 int32 = -2147483648
        for fi8 := 0; fi8 < 19; fi8++ {
            var v9 float32 = -2.25
            _ = fi8
            _ = v9
        }
        _ = v7
    } else {
        p0 = (0 >= 1)
        p0 = (p0 || ((!((p0 && p0) || p0)) || (!(float32(3.14159) < 0.1))))
        v10 := GenS1{F0: 85, F1: 2, F2: 0}
        _ = v10
    }
    for fi11 := 0; fi11 < 1; fi11++ {
        var v12 uint16 = 255
        var v13 int = fi11
        _ = fi11
        _ = v12
        _ = v13
    }
    _ = p0
    return 4294967295
}

func main() {
    var v1 int8 = -1
    v2 := make(map[int]uint32)
    var v3 uint8 = 1
    v4 := []uint32{65535, 2147483647}
    var v5 int16 = 42
    v6 := []string{"foo bar", ("" + ""), (("a" + ("\n\t" + "foo bar")) + "a"), (((("\n\t" + "a") + "hello") + ("\n\t" + "a")) + "hello")}
    for fi14 := 0; fi14 < 1; fi14++ {
        switch ("" + "a") {
        case "a":
            v1 = (v1 + (v1 - (v1 << ((65535) & 63))))
        case "hello":
            var v15 int8 = (v1 % ((v1) | 1))
            _ = v15
        case "\n\t":
            fmt.Println(v1)
        }
        v5 += v5
        var v16 float32 = 1.5
        _ = fi14
        _ = v16
    }
    if false {
        v4 = append(v4, 2147483648)
        var v17 int8 = (-(v1))
        if true {
            for fi18 := 0; fi18 < 1; fi18++ {
                if false {
                    v3 = 2
                }
                if (65535 < 42) {
                    v19 := GenS2{F0: 646, F1: 18446744073709551615, F2: (v1 + (^(v17)))}
                    _ = v19
                }
                _ = fi18
            }
            if (-1 >= -828) {
                var v20 float64 = 3.14159
                if true {
                    v1 = 97
                }
                _ = v20
            } else {
                var v21 uint16 = 2
                v22 := GenS1{F0: 2, F1: 0, F2: (v21 ^ (v21 & 751))}
                v23 := v2[338]
                _ = v21
                _ = v22
                _ = v23
            }
            v5 -= (-(v5))
        } else {
            for fi24 := 0; fi24 < 3; fi24++ {
                fmt.Println(fi24)
                if ("a" >= "hello") {
                    fmt.Println(v1)
                    var v25 int = fi24
                    if (true || (-1 >= 1)) {
                        for fi26 := 0; fi26 < 3; fi26++ {
                            fmt.Println(v1)
                            switch (fi24 + (fi24 & fi24)) {
                            case 567:
                                var v27 bool = (true && (true && (true || true)))
                                _ = v27
                            }
                            var v28 string = (("" + "foo bar") + "\n\t")
                            v5 = v5
                            v17 = (v17 / ((42) | 1))
                            _ = fi26
                            _ = v28
                        }
                    }
                    _ = v25
                }
                _ = fi24
            }
        }
        _ = v17
    }
    v1 = v1
    for fi29 := 0; fi29 < 20; fi29++ {
        if ("a" < ("a" + "a")) {
            v1 += -44
            v2[fi29] = 2147483647
            v1 -= (v1 % ((v1) | 1))
        } else {
            v30 := GenS3{F0: int32(fi29), F1: 100}
            _ = v30
        }
        v1 -= int8(v3)
        _ = fi29
    }
    var v31 uint32 = 42
    fmt.Println(v1, v5)
    v31 -= 0
    var v32 uint8 = v3
    v5 += (-(v5))
    for fi33 := 0; fi33 < 14; fi33++ {
        if (!true) {
            var v34 uint = 60
            _ = v34
        } else {
            if true {
                fmt.Println(v32)
                var v35 float32 = float32((3.14159 + float64(((0.0 - 3.14159) + float32(-2.25)))))
                var v36 float32 = v35
                _ = v35
                _ = v36
            } else {
                v4 = append(v4, (^(v31)))
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v37 uint32 = v31
            _ = v37
        }
        if len(v4) > 0 {
            v4[0] = v31
        }
        v31 += (^(v31))
        _ = fi33
    }
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
