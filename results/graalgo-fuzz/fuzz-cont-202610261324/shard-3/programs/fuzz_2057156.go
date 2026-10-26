package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 uint32; F2 uint64 }

func main() {
    v1 := []string{}
    v2 := []int{42, 1, -810}
    v3 := make(map[uint32]bool)
    var v4 string = "a"
    v5 := make(map[int]int)
    v6 := make(map[string]bool)
    v4 = "a"
    switch ((v4 + v4) + v4) {
    case "hello":
        v7 := GenS1{F0: -31}
        _ = v7
    }
    var v8 uint64 = 2
    if true {
        v4 = v4
    } else {
        fmt.Println(v8, v4)
        if (!(42 > -19)) {
            if false {
                var v9 uint16 = 2
                for fi10 := 0; fi10 < 1; fi10++ {
                    v9 -= (v9 / ((v9) | 1))
                    v11 := GenS1{F0: 127}
                    _ = fi10
                    _ = v11
                }
                v9 = v9
                _ = v9
            }
            if false {
                var v12 float64 = (3.14159 + 1.5)
                fmt.Println(v12, v4, v4)
                _ = v12
            }
        }
        for fi13 := 0; fi13 < 19; fi13++ {
            fmt.Println(v8, v4, fi13)
            _ = fi13
        }
    }
    if true {
        v8 -= 1
        var v14 uint64 = (v8 / (((v8 % ((v8) | 1))) | 1))
        _ = v14
    }
    var v15 int16 = 1000
    v15 += 117
    if false {
        switch v4 {
        case "":
            if false {
                fmt.Println(v8, v15, v8)
            }
        case "hello":
            for fi16 := 0; fi16 < 14; fi16++ {
                if (42 < -9223372036854775808) {
                    if len(v1) > 0 {
                        v1[0] = (v4 + ((("hello" + v4) + v4) + v4))
                    }
                    switch 99 {
                    case 128:
                        for fi17 := 0; fi17 < 3; fi17++ {
                            if false {
                                v4 = ("a" + v4)
                                v18 := GenS1{F0: 42}
                                _ = v18
                            } else {
                                var v19 int64 = 1
                                var v20 uint = 2
                                _ = v19
                                _ = v20
                            }
                            _ = fi17
                        }
                    case 433:
                        v21 := GenS1{F0: int8(v8)}
                        _ = v21
                    case 32768:
                        for fi22 := 0; fi22 < 11; fi22++ {
                            v8 = 1
                            _ = fi22
                        }
                    }
                } else {
                    if true {
                        v15 = (v15 * v15)
                    }
                    if (((v15 / (((v15 | v15)) | 1)) > v15) && (v8 == (^(v8)))) {
                        for fi23 := 0; fi23 < 2; fi23++ {
                            var v24 int64 = 9223372036854775807
                            switch 164 {
                            case 1:
                                if (!true) {
                                    _ = len(v1)
                                }
                            case -1:
                                v24 = (v24 & (v24 << ((994) & 63)))
                            }
                            _ = len(v2)
                            _ = fi23
                            _ = v24
                        }
                    } else {
                        switch (v4 + v4) {
                        case "\n\t":
                            var v25 int8 = 78
                            _ = v25
                        case "a":
                            for fi26 := 0; fi26 < 4; fi26++ {
                                v15 = v15
                                _ = fi26
                            }
                        default:
                            var v27 uint32 = 498
                            _ = v27
                        }
                        fmt.Println(fi16)
                    }
                }
                for fi28 := 0; fi28 < 11; fi28++ {
                    v4 = ((v4 + v4) + (((("hello" + v4) + v4) + v4) + "\n\t"))
                    _ = fi28
                }
                _ = fi16
            }
        case "foo bar":
            if len(v1) > 0 {
                v1[0] = v4
            }
        }
    } else {
        v8 += 1
        if ((42 >= 2147483647) || false) {
            if true {
                delete(v5, 32767)
                delete(v3, 42)
            }
        } else {
            var v29 uint8 = 0
            _ = v29
        }
        if len(v1) > 0 {
            v1[0] = (("\n\t" + ((v4 + v4) + v4)) + v4)
        }
    }
    var v30 int = 127
    v31 := GenS2{F0: (v15 >> ((736) & 63)), F1: 1, F2: 42}
    if false {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v32 int = 32768
        v32 += (v30 << ((819) & 63))
        _ = v32
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
