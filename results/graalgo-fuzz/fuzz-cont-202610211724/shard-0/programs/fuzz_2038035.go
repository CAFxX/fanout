package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 int16; F1 int; F2 uint32 }
type GenS4 struct { F0 uint64; F1 int64; F2 uint8 }

func main() {
    var v1 string = ""
    v2 := []uint32{42, 2}
    var v3 uint16 = 202
    var v4 bool = false
    fmt.Println(v4, v4)
    v2 = append(v2, 1)
    v4 = (!false)
    if (1.5 >= 0.1) {
        if (true || v4) {
            switch "" {
            case "foo bar":
                _ = len(v2)
            case "a":
                v3 -= v3
            case "":
                if (false || true) {
                    var v5 float32 = 3.14159
                    var v6 int16 = 0
                    if len(v2) > 0 {
                        v2[0] = 483
                    }
                    _ = v5
                    _ = v6
                } else {
                    v7 := GenS1{F0: 1}
                    v1 = v1
                    if (0.1 == 1.5) {
                        v2 = append(v2, 2147483648)
                        v4 = v4
                    }
                    _ = v7
                }
            }
            var v8 float64 = -2.25
            if v4 {
                v9 := GenS1{F0: 28}
                _ = v9
            } else {
                v1 = (("hello" + "a") + "hello")
            }
            _ = v8
        } else {
            switch (("a" + "hello") + ("hello" + ("" + v1))) {
            case "a":
                _ = len(v2)
            }
            for fi10 := 0; fi10 < 16; fi10++ {
                for fi11 := 0; fi11 < 10; fi11++ {
                    v12 := GenS1{F0: -1}
                    fmt.Println(fi10)
                    _ = fi11
                    _ = v12
                }
                for fi13 := 0; fi13 < 6; fi13++ {
                    v14 := GenS1{F0: -32768}
                    if v4 {
                        if (((0.0 - ((3.14159 / -2.25) / 1.0)) + 0.0) == float64(0.0)) {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                        v15 := GenS2{F0: (v4 || v4), F1: ((0.1 * -2.25) >= (1.5 / 0.1))}
                        if v4 {
                            v16 := GenS1{F0: 293}
                            _ = v16
                        } else {
                            if v4 {
                                fmt.Println(v3, fi10, fi10)
                                if v4 {
                                    fmt.Println(fi10)
                                    if false {
                                        v17 := GenS1{F0: -1}
                                        switch v1 {
                                        case "":
                                            _ = len(v2)
                                        }
                                        if true {
                                            var v18 int64 = -1
                                            v4 = (!v4)
                                            v2 = append(v2, 2)
                                            switch (fi13 * 128) {
                                            case 2147483647:
                                                v19 := GenS1{F0: 42}
                                                _ = v19
                                            }
                                            v1 = v1
                                            _ = v18
                                        } else {
                                            switch "" {
                                            case "a":
                                                if v4 {
                                                    var v20 int64 = -1
                                                    _ = v20
                                                }
                                            }
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        }
                                        _ = v17
                                    } else {
                                        if len(v2) > 0 {
                                            v2[0] = 4294967295
                                        }
                                    }
                                } else {
                                    fmt.Println(v4, fi13)
                                    v21 := GenS3{F0: 744, F1: fi10, F2: 306}
                                    _ = v21
                                }
                            }
                            var v22 uint8 = 2
                            _ = v22
                        }
                        _ = v15
                    }
                    _ = fi13
                    _ = v14
                }
                _ = fi10
            }
            for fi23 := 0; fi23 < 6; fi23++ {
                if len(v2) > 0 {
                    v2[0] = 2147483648
                }
                v24 := GenS4{F0: 18446744073709551615, F1: 358, F2: 138}
                _ = fi23
                _ = v24
            }
        }
    } else {
        for fi25 := 0; fi25 < 7; fi25++ {
            v1 = ((v1 + (v1 + "foo bar")) + ("\n\t" + v1))
            _ = fi25
        }
    }
    if false {
        if (469 > 0) {
            v26 := GenS1{F0: 1000}
            v1 = v1
            _ = v26
        }
        v27 := GenS3{F0: 42, F1: 32767, F2: 833}
        v1 = ""
        _ = v27
    }
    var v28 int8 = 100
    var v29 uint16 = v3
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    if false {
        fmt.Println(v1, v29)
    } else {
        var v30 bool = v4
        _ = v30
    }
    v28 = v28
    var v31 int8 = v28
    v3 -= v29
    var v32 string = (v1 + ("a" + v1))
    var v33 int64 = 646
    if v4 {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = len(v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}
