package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 uint16 }
type GenS4 struct { F0 bool }

func main() {
    var v1 uint8 = 42
    var v2 int16 = -243
    var v3 bool = false
    var v4 int32 = 2147483647
    if v3 {
        if v3 {
            v2 -= (v2 % (((v2 - (v2 + v2))) | 1))
            v1 = 128
        } else {
            v5 := GenS1{F0: 32767}
            _ = v5
        }
        fmt.Println(v4, v3)
        if (!(!v3)) {
            v2 = v2
        } else {
            var v6 uint = 2
            _ = v6
        }
    } else {
        v3 = v3
        fmt.Println(v2, v4)
        v1 += v1
    }
    if (1 <= 32768) {
        for fi7 := 0; fi7 < 19; fi7++ {
            v8 := GenS2{F0: (v1 % ((v1) | 1))}
            if v3 {
                for fi9 := 0; fi9 < 1; fi9++ {
                    var v10 int32 = 123
                    v11 := []int{fi7, (fi9 ^ (fi9 + -1)), fi9, 32767}
                    var v12 string = "foo bar"
                    _ = fi9
                    _ = v10
                    _ = v11
                    _ = v12
                }
                if (-9223372036854775808 > 607) {
                    if (!(true && v3)) {
                        switch int(v1) {
                        case -714:
                            v4 -= v4
                        default:
                            for fi13 := 0; fi13 < 17; fi13++ {
                                v14 := GenS1{F0: 42}
                                v1 = (^(v1))
                                var v15 float64 = -2.25
                                v1 += (v1 * v1)
                                _ = fi13
                                _ = v14
                                _ = v15
                            }
                        }
                    }
                    var v16 string = ""
                    _ = v16
                } else {
                    var v17 int8 = -128
                    v1 -= 2
                    v4 = 217
                    _ = v17
                }
            } else {
                v18 := []int{}
                _ = v18
            }
            _ = fi7
            _ = v8
        }
    }
    var v19 int = -114
    if (v19 == v19) {
        if (42 <= 473) {
            v4 += 42
        }
    } else {
        if ((v3 && (1e10 > float32(1.5))) && (v3 && ((-(v4)) != (v4 % ((0) | 1))))) {
            fmt.Println(v3)
            v2 = v2
            switch (("a" + "\n\t") + "a") {
            case "foo bar":
                v20 := GenS3{F0: -1, F1: 774}
                _ = v20
            }
        } else {
            var v21 float32 = (1e10 * (-2.25 / float32(-2.25)))
            switch ("foo bar" + ("a" + "\n\t")) {
            case "a":
                if (0.0 == ((1e10 * (float64(v21) + 3.14159)) / (0.1 + -2.25))) {
                    v22 := []int{}
                    _ = v22
                }
            case "":
                switch v19 {
                case -1:
                    for fi23 := 0; fi23 < 20; fi23++ {
                        fmt.Println(fi23, v21)
                        var v24 float32 = (1.5 + v21)
                        var v25 int32 = (v4 & 891)
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                case 9223372036854775807:
                    v26 := GenS1{F0: 65535}
                    _ = v26
                }
            }
            _ = v21
        }
        v27 := GenS4{F0: (649 <= 1)}
        fmt.Println(v1)
        _ = v27
    }
    switch v19 {
    case -1:
        var v28 uint8 = v1
        _ = v28
    }
    v3 = ((0 != 42) || (42 > 19))
    v29 := GenS3{F0: 42, F1: 2}
    for fi30 := 0; fi30 < 8; fi30++ {
        var v31 uint16 = 1
        v32 := GenS4{F0: v3}
        for fi33 := 0; fi33 < 8; fi33++ {
            if (42 <= (v31 % ((v31) | 1))) {
                if (2 <= 0) {
                    var v34 uint32 = 995
                    _ = v34
                }
                v35 := GenS2{F0: (^(v1))}
                _ = v35
            } else {
                var v36 uint = 750
                var v37 uint64 = 42
                v31 += (v31 ^ (v31 * (^(v31))))
                _ = v36
                _ = v37
            }
            var v38 string = ""
            var v39 uint = 4294967295
            v19 += (v19 | (v19 + (fi33 / ((127) | 1))))
            v40 := GenS1{F0: 42}
            _ = fi33
            _ = v38
            _ = v39
            _ = v40
        }
        _ = fi30
        _ = v31
        _ = v32
    }
    v41 := GenS2{F0: (v1 & v1)}
    switch v19 {
    case 443:
        v1 += v1
    }
    var v42 int8 = -1
    v42 += (v42 << ((42) & 63))
    switch "hello" {
    case "hello":
        fmt.Println(v42, v2)
    default:
        v43 := GenS4{F0: ((2 == 65535) && v3)}
        _ = v43
    }
    if v3 {
        if (18446744073709551615 >= 18446744073709551615) {
            v44 := GenS2{F0: 59}
            v3 = true
            if v3 {
                if ("hello" <= (("foo bar" + "foo bar") + "a")) {
                    v19 += v19
                    var v45 float64 = -2.25
                    _ = v45
                }
            } else {
                var v46 uint64 = 42
                _ = v46
            }
            _ = v44
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v19
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v41
    _ = v42
}
