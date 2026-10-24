package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int }
type GenS5 struct { F0 string; F1 int32 }

func main() {
    v1 := []uint32{357, 42}
    v2 := make(map[string]uint32)
    v3 := make(map[string]string)
    var v4 uint8 = 127
    v4 = v4
    v4 = v4
    v4 = (v4 | (v4 & v4))
    for fi5 := 0; fi5 < 17; fi5++ {
        v6 := GenS1{F0: 65535, F1: 1}
        fmt.Println(v4, v4)
        _ = fi5
        _ = v6
    }
    var v7 int32 = 2147483647
    var v8 uint = 42
    if (!(!true)) {
        v4 = (^(v4))
        if true {
            if (("\n\t" + "\n\t") < "\n\t") {
                for fi9 := 0; fi9 < 17; fi9++ {
                    for fi10 := 0; fi10 < 18; fi10++ {
                        switch "a" {
                        case "":
                            var v11 int16 = 42
                            _ = v11
                        }
                        v1 = append(v1, 0)
                        delete(v3, ("foo bar" + "\n\t"))
                        _ = fi10
                    }
                    var v12 float32 = 1e10
                    fmt.Println(v8)
                    _ = fi9
                    _ = v12
                }
                for fi13 := 0; fi13 < 15; fi13++ {
                    for fi14 := 0; fi14 < 12; fi14++ {
                        v7 -= (v7 ^ v7)
                        _ = fi14
                    }
                    v15 := GenS1{F0: 1, F1: 9223372036854775807}
                    fmt.Println(v8)
                    _ = fi13
                    _ = v15
                }
                for fi16 := 0; fi16 < 11; fi16++ {
                    for fi17 := 0; fi17 < 8; fi17++ {
                        var v18 int32 = (v7 ^ v7)
                        v19 := GenS2{F0: 1, F1: 1}
                        v7 -= (v7 & (v7 | v7))
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                    _ = fi16
                }
            }
            var v20 string = ("hello" + ((("hello" + "\n\t") + ("foo bar" + "foo bar")) + "foo bar"))
            _ = v20
        } else {
            var v21 string = "\n\t"
            var v22 int8 = -115
            if (32767 == 1) {
                for fi23 := 0; fi23 < 9; fi23++ {
                    v4 = v4
                    _ = fi23
                }
            } else {
                v24 := GenS3{F0: true}
                for fi25 := 0; fi25 < 7; fi25++ {
                    _ = len(v1)
                    v7 += v7
                    _ = fi25
                }
                for fi26 := 0; fi26 < 18; fi26++ {
                    v21 = v21
                    _ = fi26
                }
                _ = v24
            }
            _ = v21
            _ = v22
        }
    }
    if false {
        switch "a" {
        case "":
            var v27 int16 = 172
            _ = v27
        case "foo bar":
            v8 = 2
        case "\n\t":
            var v28 int8 = 42
            _ = v28
        }
        v8 += v8
        for fi29 := 0; fi29 < 4; fi29++ {
            for fi30 := 0; fi30 < 4; fi30++ {
                var v31 int = (fi30 * (fi30 - (^(fi30))))
                var v32 int = fi29
                v32 = v31
                _ = fi30
                _ = v31
                _ = v32
            }
            _ = fi29
        }
    } else {
        for fi33 := 0; fi33 < 1; fi33++ {
            for fi34 := 0; fi34 < 8; fi34++ {
                var v35 int16 = -32768
                _ = fi34
                _ = v35
            }
            for fi36 := 0; fi36 < 18; fi36++ {
                switch "" {
                case "\n\t":
                    v7 = v7
                case "foo bar":
                    v37 := GenS4{F0: fi36}
                    _ = v37
                }
                if false {
                    v1 = append(v1, 4294967295)
                    v38 := GenS5{F0: "\n\t", F1: v7}
                    v7 = v7
                    _ = v38
                } else {
                    v1 = append(v1, 1)
                }
                v8 = v8
                _ = fi36
            }
            var v39 int8 = 42
            _ = fi33
            _ = v39
        }
        var v40 uint8 = (v4 | 1)
        _ = v40
    }
    v41 := GenS3{F0: (405 <= 42)}
    var v42 string = "hello"
    var v43 int64 = -1
    fmt.Println(v4, v42)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v7
    _ = v8
}
