package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint8; F1 int; F2 string }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int64; F1 uint16; F2 uint16 }
type GenS7 struct { F0 int8; F1 string; F2 uint8 }

func main() {
    v1 := []string{""}
    var v2 bool = (!(222 < 4294967295))
    var v3 string = "foo bar"
    var v4 float64 = ((3.14159 / (-2.25 / (1.5 / 1.0))) + 3.14159)
    var v5 bool = true
    var v6 uint = 42
    v7 := GenS1{F0: true}
    if (3.14159 == ((0.0 - (v4 - float64((((((float32(float64(-2.25)) / ((0.1 - float32(v4)) / float32(v4))) + -2.25) * 3.14159) / 1.5) + (0.1 + (1.5 / (3.14159 - ((1.5 - 0.1) / ((1e10 / float32((1.5 - v4))) - float32(v4)))))))))) + (-2.25 + 1.5))) {
        var v8 int64 = -1
        _ = v8
    } else {
        v9 := GenS1{F0: true}
        if len(v1) > 0 {
            v1[0] = v3
        }
        _ = v9
    }
    for fi10 := 0; fi10 < 6; fi10++ {
        v1 = append(v1, v3)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi10
    }
    if true {
        var v11 int8 = 7
        switch 981 {
        case 4294967295:
            switch (v3 + v3) {
            case "hello":
                _ = len(v1)
            case "":
                var v12 int = 32768
                _ = v12
            }
        default:
            v5 = v2
        }
        _ = v11
    }
    v6 += (v6 ^ (v6 * 0))
    switch 777 {
    case 108:
        v1 = append(v1, "hello")
    default:
        _ = len(v1)
    }
    if (v2 && (0.0 >= (1e10 - (0.1 + 0.0)))) {
        if (false && v5) {
            if (0.0 <= ((3.14159 + 0.0) / 1e10)) {
                switch 32767 {
                case 971:
                    for fi13 := 0; fi13 < 6; fi13++ {
                        v14 := GenS2{F0: 0, F1: 42}
                        _ = fi13
                        _ = v14
                    }
                case 32767:
                    v15 := GenS3{F0: v3}
                    _ = v15
                default:
                    fmt.Println(v4)
                }
            } else {
                var v16 uint16 = 65535
                _ = v16
            }
            switch ((v3 + (v3 + "foo bar")) + (("a" + v3) + ((v3 + ((v3 + v3) + ("a" + v3))) + (v3 + ((("hello" + ("\n\t" + v3)) + v3) + ("\n\t" + (v3 + "\n\t"))))))) {
            case "a":
                v17 := GenS4{F0: 155, F1: -43, F2: "foo bar"}
                _ = v17
            case "foo bar":
                if ((false || v5) && (uint64(v6) == 1)) {
                    if false {
                        if len(v1) > 0 {
                            v1[0] = "foo bar"
                        }
                        v3 = ""
                    } else {
                        var v18 float64 = (v4 + ((v4 - float64(float32((0.1 / (float64(0.1) / v4))))) / float64((0.1 * 1.5))))
                        _ = v18
                    }
                } else {
                    if len(v1) > 0 {
                        v1[0] = v3
                    }
                    var v19 float32 = (((0.1 + 0.0) / 0.1) / 1.0)
                    fmt.Println(v4)
                    _ = v19
                }
            case "\n\t":
                var v20 float32 = 1.5
                _ = v20
            }
            if (v2 && (!true)) {
                v1 = append(v1, "\n\t")
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        }
    } else {
        switch v3 {
        case "hello":
            v21 := GenS2{F0: 700, F1: 9223372036854775807}
            _ = v21
        case "":
            var v22 int64 = -13
            _ = v22
        case "\n\t":
            for fi23 := 0; fi23 < 20; fi23++ {
                var v24 uint = (^(v6))
                v2 = (v5 && (0 != 1000))
                _ = fi23
                _ = v24
            }
        }
        switch "" {
        case "\n\t":
            if len(v1) > 0 {
                v1[0] = v3
            }
        }
    }
    if false {
        for fi25 := 0; fi25 < 2; fi25++ {
            if (-636 == 42) {
                switch v3 {
                case "\n\t":
                    if false {
                        var v26 float32 = 1e10
                        _ = v26
                    }
                }
                if (!v5) {
                    var v27 int16 = 42
                    _ = v27
                }
                if len(v1) > 0 {
                    v1[0] = v3
                }
            } else {
                v2 = true
            }
            v4 = float64(-2.25)
            v28 := GenS5{F0: 2}
            _ = fi25
            _ = v28
        }
        v29 := GenS6{F0: -9223372036854775808, F1: 42, F2: 867}
        _ = v29
    } else {
        switch "\n\t" {
        case "hello":
            v5 = (v6 >= (v6 % (((v6 ^ (v6 % (((v6 / ((42) | 1))) | 1)))) | 1)))
        case "a":
            v4 = 0.0
        case "foo bar":
            var v30 int = 32767
            _ = v30
        }
        if v5 {
            var v31 uint16 = 56
            v31 = 65535
            _ = v31
        } else {
            v4 += v4
            if (true && v5) {
                var v32 float32 = ((float32(v4) + 0.0) + 1e10)
                v33 := GenS5{F0: 65535}
                _ = v32
                _ = v33
            }
            v34 := GenS3{F0: v3}
            _ = v34
        }
    }
    for fi35 := 0; fi35 < 10; fi35++ {
        if (v4 < v4) {
            for fi36 := 0; fi36 < 18; fi36++ {
                var v37 uint8 = 160
                v1 = append(v1, v3)
                v38 := GenS6{F0: 0, F1: 2, F2: 42}
                _ = fi36
                _ = v37
                _ = v38
            }
            v39 := GenS7{F0: 42, F1: ("hello" + v3), F2: 0}
            v3 = v3
            _ = v39
        }
        fmt.Println(v6, v3)
        _ = fi35
    }
    if true {
        fmt.Println(v5, v4)
    } else {
        switch "a" {
        case "foo bar":
            v4 -= (-2.25 / v4)
        case "":
            v40 := GenS2{F0: 44, F1: -483}
            _ = v40
        case "a":
            fmt.Println(v2)
        }
        v6 = v6
    }
    if v2 {
        v2 = false
    }
    if len(v1) > 0 {
        v1[0] = v3
    }
    if v5 {
        fmt.Println(v5, v2)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
