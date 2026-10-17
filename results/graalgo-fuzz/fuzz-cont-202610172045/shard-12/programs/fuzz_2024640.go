package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 uint16; F2 int32 }
type GenS3 struct { F0 int32 }

func main() {
    v1 := []string{("a" + "a"), ("hello" + "\n\t"), "\n\t"}
    v2 := []string{(("" + ("foo bar" + "\n\t")) + "foo bar"), "foo bar", "", "hello"}
    var v3 uint64 = 193
    v4 := make(map[string]int)
    v5 := GenS1{F0: 2147483647}
    _ = len(v1)
    v3 = 4294967295
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 uint16 = 65535
        for fi8 := 0; fi8 < 11; fi8++ {
            if len(v2) > 0 {
                v2[0] = ("" + "\n\t")
            }
            _ = fi8
        }
        var v9 uint64 = 1
        _ = fi6
        _ = v7
        _ = v9
    }
    switch "a" {
    case "hello":
        if ((42 < 1) || (-2.25 < 0.0)) {
            for fi10 := 0; fi10 < 4; fi10++ {
                v11 := GenS1{F0: (fi10 << ((42) & 63))}
                var v12 uint16 = 725
                var v13 float32 = float32(0.0)
                _ = fi10
                _ = v11
                _ = v12
                _ = v13
            }
            v3 += 0
        } else {
            if len(v1) > 0 {
                v1[0] = "foo bar"
            }
        }
    default:
        fmt.Println(v3)
    }
    for fi14 := 0; fi14 < 12; fi14++ {
        if (!false) {
            v3 += v3
            var v15 int32 = -656
            _ = v15
        } else {
            v16 := GenS2{F0: 1000, F1: 934, F2: -2147483648}
            if false {
                var v17 int8 = 42
                if true {
                    fmt.Println(fi14, v3, fi14)
                    for fi18 := 0; fi18 < 6; fi18++ {
                        switch fi18 {
                        case 9223372036854775807:
                            if ("foo bar" > ("\n\t" + ("hello" + "a"))) {
                                _ = len(v4)
                                if (!false) {
                                    var v19 int16 = 1
                                    v19 = -1
                                    _ = v19
                                } else {
                                    var v20 int64 = 0
                                    var v21 string = "\n\t"
                                    _ = v20
                                    _ = v21
                                }
                                v22 := v4[""]
                                _ = v22
                            }
                        case -9223372036854775808:
                            v17 -= (-(v17))
                        default:
                            fmt.Println(fi18, fi18)
                        }
                        _ = fi18
                    }
                }
                fmt.Println(v3)
                _ = v17
            }
            _ = v16
        }
        if (900 < -698) {
            if true {
                v3 -= (v3 << ((437) & 63))
                if true {
                    var v23 uint8 = 255
                    _ = v23
                } else {
                    var v24 int16 = 1
                    _ = v24
                }
                v25 := GenS1{F0: (fi14 >> ((uint(v3)) & 63))}
                _ = v25
            }
            var v26 string = (("\n\t" + ("" + "a")) + "\n\t")
            if true {
                if true {
                    v26 = "hello"
                    if len(v1) > 0 {
                        v1[0] = v26
                    }
                    for fi27 := 0; fi27 < 18; fi27++ {
                        var v28 bool = (v26 <= (((v26 + "hello") + "hello") + ("a" + v26)))
                        _ = fi27
                        _ = v28
                    }
                } else {
                    var v29 int16 = -340
                    v29 = v29
                    var v30 bool = true
                    _ = v29
                    _ = v30
                }
                v26 = (v26 + v26)
            }
            _ = v26
        }
        _ = fi14
    }
    if (3.14159 > ((-2.25 / (-2.25 * 3.14159)) * 1e10)) {
        v3 = v3
        var v31 int16 = 777
        _ = v31
    }
    _ = len(v4)
    v32 := GenS3{F0: -1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
