package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 int }
type GenS2 struct { F0 uint32; F1 int16; F2 uint16 }
type GenS3 struct { F0 uint32; F1 int8 }

func main() {
    var v1 int16 = 32767
    var v2 uint64 = 390
    v3 := []uint32{}
    fmt.Println(v2)
    v1 += v1
    var v4 string = "hello"
    v2 += 1
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v5 bool = (42 > 238)
    var v6 uint64 = 4294967295
    if true {
        if v5 {
            v6 += (v2 / (((^(v6))) | 1))
        }
    } else {
        var v7 int8 = -128
        if (("\n\t" + v4) <= ((v4 + v4) + "foo bar")) {
            if (42 < -9223372036854775808) {
                switch v4 {
                case "a":
                    v2 = 18446744073709551615
                case "hello":
                    v6 -= v2
                }
            }
            v8 := GenS1{F0: "\n\t", F1: 1, F2: -1}
            _ = v8
        } else {
            if (790 == -2147483648) {
                v9 := GenS1{F0: (v4 + (((v4 + v4) + v4) + v4)), F1: 42, F2: 32768}
                _ = v9
            }
            if true {
                if (v4 >= (v4 + "\n\t")) {
                    fmt.Println(v1)
                    v10 := GenS2{F0: 55, F1: (v1 - v1), F2: 657}
                    _ = v10
                }
                if (!v5) {
                    _ = len(v3)
                    var v11 int16 = v1
                    for fi12 := 0; fi12 < 14; fi12++ {
                        _ = len(v3)
                        v13 := GenS1{F0: (v4 + (v4 + "")), F1: 876, F2: (fi12 << ((4294967295) & 63))}
                        _ = fi12
                        _ = v13
                    }
                    _ = v11
                } else {
                    var v14 int32 = -1
                    var v15 float32 = -2.25
                    _ = v14
                    _ = v15
                }
                if v5 {
                    var v16 int32 = 0
                    _ = v16
                } else {
                    v5 = v5
                }
            } else {
                var v17 int16 = (^(v1))
                switch (("hello" + ("\n\t" + ("hello" + (v4 + v4)))) + v4) {
                case "\n\t":
                    if len(v3) > 0 {
                        v3[0] = 916
                    }
                case "":
                    switch -2147483648 {
                    case 18:
                        if len(v3) > 0 {
                            v3[0] = 2
                        }
                    case 32767:
                        var v18 uint64 = (v2 ^ v6)
                        _ = v18
                    default:
                        v2 += (v6 & v6)
                    }
                }
                _ = v17
            }
        }
        v4 = v4
        _ = v7
    }
    v3 = append(v3, 42)
    if v5 {
        switch v4 {
        case "\n\t":
            _ = len(v3)
        case "":
            switch (v4 + v4) {
            case "\n\t":
                v5 = v5
            case "foo bar":
                v2 += (^(v2))
            default:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        case "foo bar":
            v4 = "hello"
        default:
            v1 += v1
        }
    }
    var v19 uint8 = 0
    v20 := GenS2{F0: 65535, F1: v1, F2: 32767}
    v19 = 128
    for fi21 := 0; fi21 < 4; fi21++ {
        v22 := GenS3{F0: 0, F1: 81}
        _ = fi21
        _ = v22
    }
    if (3.14159 == (0.1 * 3.14159)) {
        for fi23 := 0; fi23 < 20; fi23++ {
            for fi24 := 0; fi24 < 9; fi24++ {
                var v25 int32 = -1
                _ = fi24
                _ = v25
            }
            _ = fi23
        }
        var v26 int = 207
        _ = v26
    } else {
        switch ((v4 + "hello") + (v4 + "foo bar")) {
        case "\n\t":
            v6 = (v2 & v6)
        case "a":
            var v27 int32 = 2147483647
            _ = v27
        }
        v6 = (v2 >> ((1) & 63))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
