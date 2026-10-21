package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int16; F1 int64 }

func main() {
    v1 := make(map[string]int)
    var v2 int8 = -116
    var v3 uint = 4294967295
    if true {
        if (!(-2.25 == (-2.25 + 0.0))) {
            for fi4 := 0; fi4 < 10; fi4++ {
                v2 -= (^(v2))
                _ = fi4
            }
        }
        var v5 uint16 = 2
        v6 := GenS1{F0: v3, F1: 2}
        _ = v5
        _ = v6
    }
    if false {
        v7 := GenS1{F0: 90, F1: 2}
        _ = v7
    } else {
        var v8 uint = 0
        v2 = v2
        if (42 <= 127) {
            var v9 uint8 = 127
            _ = v9
        } else {
            for fi10 := 0; fi10 < 14; fi10++ {
                v3 -= (v8 | v8)
                var v11 uint32 = 0
                _ = fi10
                _ = v11
            }
        }
        _ = v8
    }
    v2 = (v2 | (v2 - -121))
    v2 += v2
    var v12 uint8 = 2
    v2 += v2
    fmt.Println(v12, v2)
    var v13 int32 = -1
    if true {
        for fi14 := 0; fi14 < 16; fi14++ {
            var v15 int16 = 32767
            _ = fi14
            _ = v15
        }
    } else {
        switch -1 {
        case 32767:
            var v16 uint32 = 1
            _ = v16
        case 2147483647:
            v17 := GenS2{F0: v3}
            _ = v17
        case -505:
            for fi18 := 0; fi18 < 7; fi18++ {
                var v19 float64 = 3.14159
                var v20 string = "\n\t"
                switch "hello" {
                case "":
                    v1[v20] = fi18
                default:
                    var v21 uint16 = 171
                    _ = v21
                }
                _ = fi18
                _ = v19
                _ = v20
            }
        }
        v2 -= -128
    }
    v22 := GenS3{F0: 1000}
    if false {
        var v23 string = "foo bar"
        var v24 int32 = -540
        _ = v23
        _ = v24
    } else {
        if (2147483648 == 65535) {
            v25 := GenS3{F0: 32767}
            _ = v25
        }
        if (-90 >= -9223372036854775808) {
            var v26 uint8 = 0
            if true {
                if false {
                    _ = len(v1)
                }
                switch 9223372036854775807 {
                case -739:
                    var v27 int16 = 32767
                    _ = v27
                case -2147483648:
                    fmt.Println(v26)
                case -660:
                    var v28 int = 365
                    _ = v28
                }
                v12 = (^(v12))
            } else {
                if false {
                    v3 -= v3
                    if false {
                        if false {
                            v29 := GenS4{F0: -930, F1: -1}
                            for fi30 := 0; fi30 < 2; fi30++ {
                                v31 := []int{32767, 2147483647}
                                _ = fi30
                                _ = v31
                            }
                            var v32 int8 = 100
                            _ = v29
                            _ = v32
                        } else {
                            switch 2147483647 {
                            case 195:
                                v33 := v1["a"]
                                _ = v33
                            case -211:
                                var v34 string = ("foo bar" + ("" + "\n\t"))
                                _ = v34
                            case 2147483647:
                                for fi35 := 0; fi35 < 7; fi35++ {
                                    fmt.Println(v12)
                                    _ = fi35
                                }
                            }
                            if false {
                                fmt.Println(v2, v13)
                                v36 := GenS3{F0: 42}
                                v37 := v1[(("" + "\n\t") + "a")]
                                _ = v36
                                _ = v37
                            }
                        }
                    }
                }
                switch 4294967295 {
                case 1:
                    fmt.Println(v12)
                case 128:
                    v38 := GenS4{F0: 32767, F1: 1}
                    _ = v38
                default:
                    v39 := []int{}
                    _ = v39
                }
            }
            _ = v26
        } else {
            v2 -= (v2 - v2)
        }
    }
    var v40 int8 = (v2 + v2)
    var v41 uint64 = 18446744073709551615
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v22
    _ = v3
    _ = v40
    _ = v41
}
