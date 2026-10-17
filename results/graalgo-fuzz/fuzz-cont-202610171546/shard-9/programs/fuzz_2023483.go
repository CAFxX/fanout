package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 int16 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 int32; F1 uint; F2 uint8 }
func genFn0() string {
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 int16 = -32768
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 9; fi8++ {
        var v9 uint8 = 127
        _ = fi8
        _ = v9
    }
    switch "foo bar" {
    case "\n\t":
        for fi10 := 0; fi10 < 19; fi10++ {
            var v11 float32 = ((1e10 + 3.14159) / 3.14159)
            _ = fi10
            _ = v11
        }
    case "foo bar":
        v12 := []int{9223372036854775807, 883, -9223372036854775808, 2147483647}
        _ = v12
    }
    var v13 int16 = -32768
    _ = v13
    return ("foo bar" + "\n\t")
}

func main() {
    var v1 string = "a"
    var v2 uint16 = 912
    var v3 uint32 = 4294967295
    v4 := []uint32{}
    var v5 int = -1
    var v14 string = "\n\t"
    if len(v4) > 0 {
        v4[0] = (v3 | (v3 / (((v3 / ((v3) | 1))) | 1)))
    }
    var v15 uint64 = 210
    switch (v5 & v5) {
    case -2147483648:
        for fi16 := 0; fi16 < 20; fi16++ {
            v3 -= 2
            v2 += 497
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi16
        }
    case 759:
        v17 := GenS1{F0: 619, F1: 1, F2: -1}
        _ = v17
    case 1:
        if len(v4) > 0 {
            v4[0] = (v3 & 1)
        }
    }
    for fi18 := 0; fi18 < 15; fi18++ {
        for fi19 := 0; fi19 < 13; fi19++ {
            var v20 uint = 642
            var v21 bool = false
            _ = fi19
            _ = v20
            _ = v21
        }
        if len(v4) > 0 {
            v4[0] = 778
        }
        _ = fi18
    }
    if false {
        v22 := GenS1{F0: -545, F1: 279, F2: 42}
        for fi23 := 0; fi23 < 14; fi23++ {
            switch (v5 / ((fi23) | 1)) {
            case -79:
                v24 := GenS1{F0: 0, F1: -9223372036854775808, F2: 0}
                _ = v24
            case 492:
                var v25 uint16 = v2
                _ = v25
            }
            _ = fi23
        }
        _ = v22
    } else {
        for fi26 := 0; fi26 < 12; fi26++ {
            v5 -= (^(fi26))
            var v27 float64 = (0.1 - 0.1)
            _ = fi26
            _ = v27
        }
        if true {
            v28 := GenS1{F0: 9223372036854775807, F1: -1, F2: 0}
            var v29 float64 = 1e10
            v14 = ("\n\t" + (v1 + v1))
            _ = v28
            _ = v29
        } else {
            v30 := GenS2{F0: -1, F1: 213}
            _ = v30
        }
        var v31 uint = 42
        _ = v31
    }
    var v32 int16 = 1
    if true {
        if false {
            if (false || true) {
                var v33 uint8 = 255
                var v34 int16 = 960
                var v35 bool = (true && ((v33 ^ v33) <= v33))
                _ = v33
                _ = v34
                _ = v35
            } else {
                var v36 uint = 537
                if true {
                    if false {
                        v37 := GenS2{F0: v32, F1: 128}
                        _ = v37
                    }
                }
                v38 := GenS2{F0: v32, F1: 94}
                _ = v36
                _ = v38
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v15 += (v15 & uint64(v2))
        }
    } else {
        var v39 int32 = 42
        for fi40 := 0; fi40 < 12; fi40++ {
            fmt.Println(v5, fi40, v5)
            var v41 int32 = (-(v39))
            if ("" <= (((("a" + (("a" + "\n\t") + "\n\t")) + "") + v14) + "")) {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            _ = fi40
            _ = v41
        }
        v5 -= 128
        _ = v39
    }
    v5 += (v5 - (-(v5)))
    v32 += -32768
    v1 = ""
    for fi42 := 0; fi42 < 16; fi42++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v43 int = (v5 / (((fi42 + fi42)) | 1))
        _ = fi42
        _ = v43
    }
    var v44 int8 = 62
    for fi45 := 0; fi45 < 15; fi45++ {
        var v46 string = v14
        if ((v2 & (v2 | (^(v2)))) < (v2 << ((4294967295) & 63))) {
            var v47 int64 = 68
            if true {
                v48 := GenS3{F0: 1, F1: 0, F2: 172}
                fmt.Println(v15)
                _ = len(v4)
                _ = v48
            }
            _ = v47
        } else {
            if (487 < -1) {
                switch (v5 >> ((18446744073709551615) & 63)) {
                case 127:
                    v1 = (v14 + v1)
                case 795:
                    _ = len(v4)
                }
            }
            var v49 uint = 547
            _ = v49
        }
        _ = fi45
        _ = v46
    }
    if len(v4) > 0 {
        v4[0] = (v3 & (v3 / ((v3) | 1)))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v44
    _ = v5
}
