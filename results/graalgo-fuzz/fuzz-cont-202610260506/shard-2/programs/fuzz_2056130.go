package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]uint32)
    var v3 bool = false
    var v4 int32 = -2147483648
    var v5 int64 = -451
    var v6 uint16 = 2
    for fi7 := 0; fi7 < 5; fi7++ {
        if (true || ("a" <= "hello")) {
            var v8 int8 = 1
            switch fi7 {
            case 217:
                for fi9 := 0; fi9 < 20; fi9++ {
                    var v10 bool = v3
                    var v11 uint8 = 0
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            case -1000:
                if (("\n\t" + ("a" + "hello")) == (("a" + (("" + ("hello" + ("foo bar" + ("a" + "foo bar")))) + "foo bar")) + "foo bar")) {
                    if false {
                        switch (fi7 << ((0) & 63)) {
                        case -1:
                            v12 := GenS1{F0: (v4 / (((v4 ^ -2147483648)) | 1)), F1: 484}
                            _ = v12
                        case -984:
                            switch 2147483647 {
                            case 2147483647:
                                var v13 uint64 = 0
                                _ = v13
                            case 127:
                                fmt.Println(fi7, v3)
                            case 32768:
                                var v14 int = fi7
                                _ = v14
                            default:
                                for fi15 := 0; fi15 < 1; fi15++ {
                                    var v16 int = (fi15 % ((fi7) | 1))
                                    delete(v1, "")
                                    if true {
                                        var v17 int64 = (^(v5))
                                        v8 = (v8 / (((v8 * v8)) | 1))
                                        if (1000 >= 1) {
                                            _ = len(v1)
                                            var v18 uint = 18446744073709551615
                                            _ = v18
                                        }
                                        v3 = (true && (127 <= 127))
                                        v17 = (v5 << ((42) & 63))
                                        v3 = false
                                        v8 += v8
                                        _ = v17
                                    }
                                    _ = fi15
                                    _ = v16
                                }
                            }
                        }
                    }
                    var v19 int8 = 42
                    v4 = 0
                    _ = v19
                } else {
                    var v20 uint8 = 1
                    _ = v20
                }
            case 32768:
                var v21 int32 = int32(fi7)
                _ = v21
            default:
                v4 = (^(v4))
            }
            for fi22 := 0; fi22 < 14; fi22++ {
                if (234 >= 2) {
                    if false {
                        v23 := GenS1{F0: 1, F1: 65535}
                        _ = v23
                    }
                    v6 = v6
                    v8 -= 42
                }
                _ = fi22
            }
            _ = v8
        } else {
            v24 := v1["hello"]
            if (0.0 >= 0.0) {
                fmt.Println(v6)
                var v25 int16 = 326
                if false {
                    var v26 uint32 = 4294967295
                    v5 = int64(v4)
                    _ = v26
                }
                _ = v25
            } else {
                var v27 string = ("hello" + "foo bar")
                _ = v27
            }
            v3 = ("a" == ("hello" + "a"))
            _ = v24
        }
        _ = fi7
    }
    var v28 int = -1
    v5 += v5
    if ("" > "a") {
        var v29 string = (("hello" + "a") + "foo bar")
        fmt.Println(v29)
        var v30 uint16 = 473
        _ = v29
        _ = v30
    }
    for fi31 := 0; fi31 < 11; fi31++ {
        switch 1 {
        case 209:
            v32 := GenS1{F0: -1, F1: 255}
            _ = v32
        case 2147483647:
            v3 = (0.0 < 1e10)
        }
        _ = fi31
    }
    var v33 int8 = -128
    var v34 float64 = -2.25
    for fi35 := 0; fi35 < 8; fi35++ {
        v33 = (v33 & 0)
        if (!((0 >= 1000) && v3)) {
            switch v28 {
            case 2147483647:
                v36 := GenS1{F0: v4, F1: 0}
                _ = v36
            case 4294967295:
                fmt.Println(v6, v3)
            default:
                var v37 bool = false
                _ = v37
            }
            for fi38 := 0; fi38 < 7; fi38++ {
                v39 := GenS1{F0: v4, F1: 981}
                v34 -= v34
                _ = fi38
                _ = v39
            }
        }
        v4 = v4
        _ = fi35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
