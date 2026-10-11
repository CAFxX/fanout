package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 int8 }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 int; F1 uint16 }

func main() {
    v1 := make(map[string]int)
    var v2 uint16 = 2
    var v3 float32 = -2.25
    var v4 uint8 = 233
    v5 := GenS1{F0: 32767, F1: 1, F2: 1}
    fmt.Println(v4)
    var v6 uint64 = 349
    if (0 > 2147483647) {
        for fi7 := 0; fi7 < 6; fi7++ {
            for fi8 := 0; fi8 < 16; fi8++ {
                var v9 int16 = 1
                _ = fi8
                _ = v9
            }
            _ = fi7
        }
        v4 += v4
        var v10 uint64 = 4294967295
        _ = v10
    } else {
        if false {
            if false {
                var v11 uint8 = 3
                v12 := v1["\n\t"]
                switch v12 {
                case 1:
                    if (false || false) {
                        var v13 float32 = (-2.25 * (v3 * v3))
                        _ = v13
                    }
                case -980:
                    v6 += (v6 - v6)
                }
                _ = v11
                _ = v12
            } else {
                if (-1 != 9223372036854775807) {
                    var v14 int16 = -32768
                    fmt.Println(v2, v4)
                    _ = v14
                } else {
                    v15 := []int{32768, -9223372036854775808, -2147483648, -9223372036854775808}
                    v16 := GenS1{F0: -32768, F1: 2147483647, F2: 100}
                    _ = v15
                    _ = v16
                }
                if true {
                    v3 = float32(float64(v3))
                    switch ("hello" + "foo bar") {
                    case "foo bar":
                        if ((v3 <= ((v3 * 0.0) * v3)) || (-309 != -987)) {
                            for fi17 := 0; fi17 < 11; fi17++ {
                                var v18 uint = 2
                                _ = fi17
                                _ = v18
                            }
                            fmt.Println(v6, v3)
                        }
                    case "a":
                        delete(v1, "foo bar")
                    }
                    v19 := GenS1{F0: 32767, F1: 0, F2: 0}
                    _ = v19
                }
                v20 := GenS2{F0: ("" + ("\n\t" + ("" + ("hello" + "")))), F1: -769}
                _ = v20
            }
        }
    }
    var v21 int32 = 42
    var v22 float32 = 1e10
    v23 := GenS3{F0: -2147483648, F1: (v2 / (((v2 - (^(v2)))) | 1))}
    var v24 uint32 = 2147483648
    fmt.Println(v2, v21, v4)
    v25 := GenS1{F0: 1000, F1: 0, F2: -79}
    fmt.Println(v4, v24)
    fmt.Println(v2, v2)
    var v26 uint64 = uint64(v24)
    for fi27 := 0; fi27 < 18; fi27++ {
        switch (("\n\t" + "") + (("" + "\n\t") + (("hello" + "foo bar") + "foo bar"))) {
        case "\n\t":
            var v28 int8 = 42
            _ = v28
        }
        var v29 int8 = 127
        _ = fi27
        _ = v29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
