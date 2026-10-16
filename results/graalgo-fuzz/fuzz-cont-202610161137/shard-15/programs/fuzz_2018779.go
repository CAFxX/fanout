package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 string; F1 uint8; F2 uint8 }
type GenS3 struct { F0 int64; F1 int8 }
type GenS4 struct { F0 uint16; F1 uint; F2 int }
type GenS5 struct { F0 uint64; F1 uint16; F2 int }
type GenS6 struct { F0 int8 }

func main() {
    var v1 uint8 = 42
    var v2 uint32 = 965
    var v3 float32 = (((1e10 * (1.5 * (0.1 + 1e10))) + float32((-2.25 * 1.5))) + float32(0.1))
    v4 := []uint32{140}
    v5 := make(map[uint32]uint32)
    var v6 int32 = 2147483647
    for fi7 := 0; fi7 < 6; fi7++ {
        if false {
            for fi8 := 0; fi8 < 8; fi8++ {
                if (858 >= 42) {
                    fmt.Println(fi7)
                    v9 := GenS1{F0: -290, F1: 338}
                    var v10 int32 = (-(v6))
                    _ = v10
                    _ = v9
                }
                var v11 uint8 = (^(v1))
                if len(v4) > 0 {
                    v4[0] = v2
                }
                if (!(65535 < 65535)) {
                    v12 := GenS2{F0: ("foo bar" + (("" + "") + "a")), F1: v1, F2: 128}
                    _ = v12
                } else {
                    v13 := GenS3{F0: -9223372036854775808, F1: 100}
                    v14 := GenS4{F0: 32767, F1: 0, F2: (fi7 * (fi8 >> ((65535) & 63)))}
                    _ = v13
                    _ = v14
                }
                var v15 uint8 = (v11 ^ (v11 * 128))
                _ = fi8
                _ = v11
                _ = v15
            }
            v1 += v1
            for fi16 := 0; fi16 < 3; fi16++ {
                var v17 uint16 = 2
                var v18 string = "foo bar"
                if true {
                    var v19 int16 = 1000
                    v1 = (v1 * 0)
                    v20 := GenS2{F0: v18, F1: v1, F2: v1}
                    _ = v19
                    _ = v20
                } else {
                    v21 := GenS4{F0: 0, F1: 42, F2: fi16}
                    _ = v21
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v22 bool = ((^(v2)) >= v2)
                _ = fi16
                _ = v17
                _ = v18
                _ = v22
            }
        }
        var v23 int32 = (v6 << ((471) & 63))
        _ = fi7
        _ = v23
    }
    v1 = 2
    v6 = 42
    for fi24 := 0; fi24 < 6; fi24++ {
        _ = len(v4)
        _ = fi24
    }
    for fi25 := 0; fi25 < 9; fi25++ {
        _ = len(v5)
        var v26 string = ""
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi25
        _ = v26
    }
    if (!((("a" + "a") <= "a") && (3.14159 <= (-2.25 * (-2.25 * 0.0))))) {
        v27 := GenS5{F0: 1, F1: 668, F2: 32768}
        if false {
            var v28 int32 = v6
            var v29 int64 = -9223372036854775808
            var v30 float32 = 1.5
            _ = v28
            _ = v29
            _ = v30
        } else {
            v3 = float32(float64((-2.25 / 1e10)))
            if false {
                fmt.Println(v1, v3)
            }
        }
        _ = v27
    } else {
        var v31 string = "\n\t"
        _ = v31
    }
    if false {
        var v32 uint64 = 124
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        switch "foo bar" {
        case "a":
            for fi33 := 0; fi33 < 8; fi33++ {
                var v34 int = 32767
                _ = fi33
                _ = v34
            }
        case "\n\t":
            if false {
                for fi35 := 0; fi35 < 20; fi35++ {
                    for fi36 := 0; fi36 < 17; fi36++ {
                        if false {
                            v37 := GenS6{F0: 0}
                            _ = v37
                        }
                        _ = fi36
                    }
                    _ = fi35
                }
                if true {
                    var v38 uint64 = uint64(v2)
                    v39 := v5[42]
                    _ = v38
                    _ = v39
                }
                _ = len(v5)
            }
        }
        _ = v32
    }
    var v40 bool = false
    var v41 int16 = -315
    v6 += 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
