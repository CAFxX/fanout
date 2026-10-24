package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint; F1 int8 }

func main() {
    v1 := []uint32{42, 976, 42, 2147483648}
    var v2 int = -9223372036854775808
    v3 := make(map[string]int)
    var v4 bool = false
    var v5 int8 = 42
    v6 := GenS1{F0: 42}
    v2 = (v2 & v2)
    switch v2 {
    case 1:
        v2 = v2
    case 4294967295:
        var v7 uint32 = 2
        _ = v7
    case 35:
        var v8 int32 = 0
        _ = v8
    default:
        v4 = (0 >= 153)
    }
    if v4 {
        v9 := v3[""]
        _ = v9
    }
    fmt.Println(v2)
    v4 = ((0.0 < float32(-2.25)) || (v4 && (1 == -528)))
    var v10 uint8 = 118
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v5)
    v5 = v5
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        fmt.Println(v10, fi11, v2)
        _ = fi11
    }
    switch "hello" {
    case "hello":
        v10 = v10
    case "foo bar":
        var v12 float32 = (float32((float64(3.14159) * -2.25)) + (((1.5 * float32(0.0)) - float32((0.1 - ((float64(1.5) - (1e10 * ((0.1 - -2.25) - ((1.5 * (1e10 * (((((3.14159 - 1e10) * -2.25) / (1e10 / 3.14159)) + 1e10) / 1.5))) + (3.14159 / (3.14159 * (1e10 * (0.1 / 0.1)))))))) + (1.5 / (1.5 - 3.14159)))))) + 1.5))
        _ = v12
    }
    v13 := GenS2{F0: 244, F1: 1}
    var v14 int8 = v5
    switch (v2 << ((893) & 63)) {
    case 42:
        switch ("hello" + "") {
        case "foo bar":
            switch (-(v2)) {
            case 42:
                v1 = append(v1, 65535)
            case 127:
                v15 := GenS1{F0: 360}
                _ = v15
            case 847:
                if ((((2147483647 == 2147483648) && (257 > 301)) || false) || (!(false || (1 <= 37)))) {
                    switch v2 {
                    case 32767:
                        if v4 {
                            for fi16 := 0; fi16 < 19; fi16++ {
                                v2 = v2
                                _ = fi16
                            }
                            delete(v3, "foo bar")
                            for fi17 := 0; fi17 < 5; fi17++ {
                                for fi18 := 0; fi18 < 3; fi18++ {
                                    var v19 int32 = -930
                                    _ = fi18
                                    _ = v19
                                }
                                var v20 int32 = 947
                                _ = fi17
                                _ = v20
                            }
                        }
                    case 2147483647:
                        v2 -= (v2 << ((823) & 63))
                    default:
                        if (!(!v4)) {
                            fmt.Println(v5)
                        }
                    }
                } else {
                    var v21 uint64 = 4294967295
                    _ = v21
                }
            }
        case "a":
            if (0 == -290) {
                fmt.Println(v14, v14, v10)
                v1 = append(v1, 2)
                _ = len(v1)
            } else {
                var v22 uint = 0
                fmt.Println(v2, v5)
                if (v10 > (v10 >> (((^(v22))) & 63))) {
                    for fi23 := 0; fi23 < 7; fi23++ {
                        delete(v3, "\n\t")
                        v24 := v3[""]
                        if v4 {
                            switch (((("foo bar" + ((("foo bar" + "a") + "hello") + "a")) + "hello") + "\n\t") + "foo bar") {
                            case "":
                                var v25 uint = 325
                                _ = v25
                            }
                        } else {
                            fmt.Println(v2, v24)
                        }
                        _ = fi23
                        _ = v24
                    }
                    v10 = (^(v10))
                }
                _ = v22
            }
        case "\n\t":
            var v26 uint16 = 1
            _ = v26
        }
    default:
        for fi27 := 0; fi27 < 18; fi27++ {
            var v28 float32 = -2.25
            v29 := GenS3{F0: 256}
            _ = fi27
            _ = v28
            _ = v29
        }
    }
    switch (v2 * v2) {
    case -9223372036854775808:
        switch -510 {
        case -469:
            for fi30 := 0; fi30 < 1; fi30++ {
                if v4 {
                    _ = len(v1)
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                } else {
                    var v31 int16 = -32768
                    if len(v1) > 0 {
                        v1[0] = 1
                    }
                    delete(v3, "\n\t")
                    _ = v31
                }
                v32 := GenS4{F0: 255, F1: (v14 | (v5 - v14))}
                var v33 uint32 = 2
                _ = fi30
                _ = v32
                _ = v33
            }
        }
    case 214:
        if v4 {
            switch ("foo bar" + (("" + ("a" + "foo bar")) + ((("" + "hello") + "a") + "hello"))) {
            case "foo bar":
                switch v2 {
                case 32767:
                    v34 := GenS2{F0: -1, F1: 2147483647}
                    _ = v34
                }
            }
            var v35 int16 = 578
            _ = v35
        }
    case 32767:
        v5 += v14
    default:
        var v36 int16 = -634
        _ = v36
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
