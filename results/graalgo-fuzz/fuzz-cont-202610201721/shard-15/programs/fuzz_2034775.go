package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int; F1 uint16; F2 int32 }
type GenS5 struct { F0 int16; F1 uint64; F2 int64 }

func main() {
    var v1 bool = false
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]string)
    var v4 float32 = -2.25
    v5 := []string{"a", "hello", ("\n\t" + "a")}
    v6 := make(map[string]int)
    if v1 {
        v1 = v1
        if (65535 >= 42) {
            var v7 uint8 = 255
            if len(v5) > 0 {
                v5[0] = "a"
            }
            _ = v7
        } else {
            switch "" {
            case "\n\t":
                if (!(v4 >= 1e10)) {
                    v8 := GenS1{F0: 42}
                    _ = v8
                }
            case "foo bar":
                v1 = v1
            case "a":
                fmt.Println(v4)
            default:
                var v9 uint64 = 42
                _ = v9
            }
            for fi10 := 0; fi10 < 13; fi10++ {
                switch (fi10 << ((0) & 63)) {
                case -301:
                    if true {
                        var v11 bool = (-1 <= 0)
                        fmt.Println(v1, fi10, v4)
                        switch (("\n\t" + ("a" + "foo bar")) + ((("" + "hello") + "") + "foo bar")) {
                        case "\n\t":
                            v11 = (!v11)
                        default:
                            for fi12 := 0; fi12 < 10; fi12++ {
                                switch (fi12 >> ((2) & 63)) {
                                case 32767:
                                    _ = len(v5)
                                default:
                                    fmt.Println(fi10, fi10, fi12)
                                    v1 = (-43 < -128)
                                    fmt.Println(fi12, v4, fi10)
                                    var v13 uint8 = 128
                                    v14 := GenS1{F0: 527}
                                    _ = v13
                                    _ = v14
                                }
                                fmt.Println(fi12, v1, v1)
                                fmt.Println(fi10, fi12)
                                _ = fi12
                            }
                        }
                        _ = v11
                    } else {
                        v5 = append(v5, "\n\t")
                        for fi15 := 0; fi15 < 17; fi15++ {
                            v16 := GenS2{F0: "hello"}
                            _ = fi15
                            _ = v16
                        }
                    }
                case 561:
                    if (65535 == 42) {
                        var v17 uint8 = 6
                        var v18 int16 = -1
                        var v19 uint = 12
                        _ = v17
                        _ = v18
                        _ = v19
                    } else {
                        var v20 uint = 463
                        _ = v20
                    }
                default:
                    if (-128 <= 126) {
                        var v21 float64 = float64((float32(0.0) + (v4 - 0.1)))
                        _ = v21
                    }
                }
                var v22 string = "a"
                _ = fi10
                _ = v22
            }
            v1 = true
        }
    }
    v23 := GenS1{F0: 42}
    for fi24 := 0; fi24 < 16; fi24++ {
        for fi25 := 0; fi25 < 8; fi25++ {
            v4 -= (3.14159 / 0.1)
            if v1 {
                v5 = append(v5, ((("a" + "") + ((("hello" + "hello") + "foo bar") + (("" + "a") + ("foo bar" + "\n\t")))) + "hello"))
                var v26 uint16 = 936
                var v27 uint16 = (v26 << ((255) & 63))
                _ = v26
                _ = v27
            }
            if (0 < 115) {
                fmt.Println(v1)
                var v28 int32 = -1
                switch "\n\t" {
                case "foo bar":
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                case "hello":
                    var v29 int8 = 4
                    _ = v29
                }
                _ = v28
            } else {
                var v30 float64 = 0.0
                v31 := GenS2{F0: ""}
                _ = v30
                _ = v31
            }
            _ = fi25
        }
        _ = fi24
    }
    var v32 string = "a"
    if true {
        fmt.Println(v32, v32, v4)
        if v1 {
            v4 = v4
        } else {
            fmt.Println(v4, v1)
            var v33 float32 = v4
            _ = v33
        }
    } else {
        v1 = (257 != 4294967295)
        v32 = v32
        v34 := GenS3{F0: 100}
        _ = v34
    }
    v35 := v6[v32]
    fmt.Println(v4, v1, v4)
    switch (v35 * 238) {
    case 1:
        if v1 {
            v4 = v4
        } else {
            if v1 {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v36 := GenS4{F0: (v35 / ((272) | 1)), F1: 255, F2: 152}
                _ = v36
            }
            v37 := v2[2147483647]
            if false {
                if (!v37) {
                    v38 := GenS2{F0: v32}
                    _ = v38
                }
                if v1 {
                    v35 = v35
                } else {
                    var v39 int16 = -1000
                    _ = v39
                }
            }
            _ = v37
        }
    case -186:
        v40 := GenS5{F0: 42, F1: 60, F2: 1}
        _ = v40
    case 9223372036854775807:
        for fi41 := 0; fi41 < 16; fi41++ {
            v3[851] = "a"
            _ = fi41
        }
    default:
        if false {
            v35 = v35
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v32
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
