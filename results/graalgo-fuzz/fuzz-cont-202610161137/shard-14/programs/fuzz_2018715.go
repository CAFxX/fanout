package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 uint; F2 int64 }
type GenS3 struct { F0 uint8 }

func main() {
    v1 := make(map[int]bool)
    var v2 int = -94
    var v3 int16 = 1000
    fmt.Println(v3, v2)
    var v4 string = (("foo bar" + ("foo bar" + (("foo bar" + "hello") + "foo bar"))) + "")
    var v5 string = v4
    var v6 bool = (4294967295 > 2)
    switch -9223372036854775808 {
    case 4294967295:
        var v7 uint = 562
        _ = v7
    case 42:
        switch (-(v2)) {
        case 32768:
            v4 = "hello"
        case 4294967295:
            for fi8 := 0; fi8 < 15; fi8++ {
                var v9 float32 = (float32(-2.25) + -2.25)
                var v10 bool = v6
                v11 := GenS1{F0: 0}
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
        default:
            if (2 > 1) {
                if v6 {
                    fmt.Println(v5, v3, v6)
                }
                var v12 int = v2
                switch v4 {
                case "a":
                    for fi13 := 0; fi13 < 8; fi13++ {
                        switch (v4 + "") {
                        case "a":
                            if v6 {
                                switch "foo bar" {
                                case "hello":
                                    v14 := GenS1{F0: 42}
                                    _ = v14
                                default:
                                    for fi15 := 0; fi15 < 5; fi15++ {
                                        if (!false) {
                                            fmt.Println(v4, v3, v6)
                                            v1[(^(fi13))] = (502 > 1)
                                            if (4294967295 < 65535) {
                                                _ = len(v1)
                                                fmt.Println(fi15, v4, fi13)
                                                var v16 uint8 = 255
                                                _ = v16
                                            }
                                            fmt.Println(fi13)
                                            var v17 int = (^(fi13))
                                            _ = v17
                                        } else {
                                            v6 = v6
                                        }
                                        v5 = ""
                                        if true {
                                            v18 := GenS2{F0: v2, F1: 65535, F2: -464}
                                            if false {
                                                v19 := GenS1{F0: 2}
                                                var v20 int64 = 9223372036854775807
                                                var v21 uint8 = 128
                                                _ = v19
                                                _ = v20
                                                _ = v21
                                            }
                                            _ = v18
                                        } else {
                                            var v22 int32 = 0
                                            var v23 int32 = (^(v22))
                                            switch (v4 + v5) {
                                            case "a":
                                                v24 := GenS3{F0: 178}
                                                _ = v24
                                            }
                                            _ = v22
                                            _ = v23
                                        }
                                        _ = fi15
                                    }
                                }
                            }
                        default:
                            v5 = ""
                        }
                        for fi25 := 0; fi25 < 19; fi25++ {
                            if (false && (42 < 0)) {
                                v5 = "hello"
                                var v26 int64 = 618
                                _ = v26
                            } else {
                                var v27 int = 1
                                _ = v27
                            }
                            v12 += v12
                            if (v4 >= v4) {
                                v12 += 730
                                v28 := GenS2{F0: 42, F1: 2, F2: 254}
                                _ = v28
                            } else {
                                fmt.Println(v6)
                            }
                            v29 := GenS2{F0: v12, F1: 0, F2: 0}
                            v2 = (v12 | fi25)
                            fmt.Println(v4)
                            _ = fi25
                            _ = v29
                        }
                        _ = fi13
                    }
                case "hello":
                    v30 := GenS1{F0: 232}
                    _ = v30
                default:
                    fmt.Println(v6, v3)
                }
                _ = v12
            }
        }
    }
    v31 := GenS3{F0: 140}
    var v32 uint64 = 1
    v33 := GenS2{F0: (v2 ^ (v2 / ((v2) | 1))), F1: 667, F2: 0}
    var v34 int = (v2 - (v2 % (((v2 & v2)) | 1)))
    v34 = v2
    v3 = v3
    var v35 int16 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
