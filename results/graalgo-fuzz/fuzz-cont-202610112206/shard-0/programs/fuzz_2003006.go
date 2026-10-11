package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32 }

func main() {
    var v1 uint16 = 65535
    var v2 float32 = float32(((1e10 + 1e10) - 1e10))
    v3 := make(map[int]string)
    v4 := make(map[uint32]int)
    v2 += 3.14159
    var v5 uint32 = 694
    v6 := []int{}
    if (1 >= 0) {
        if (42 >= 2) {
            if (-1 != 32767) {
                v7 := GenS1{F0: int(v5)}
                fmt.Println(v1, v2)
                _ = v7
            }
            if true {
                v5 = (v5 ^ v5)
            }
        } else {
            fmt.Println(v1)
            var v8 int16 = -32768
            _ = v8
        }
        var v9 int8 = 1
        _ = v9
    } else {
        var v10 int32 = 2147483647
        v10 -= 42
        _ = v10
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        for fi12 := 0; fi12 < 16; fi12++ {
            if false {
                if true {
                    v2 = (float32(float64((1e10 - v2))) * float32(1.5))
                    v13 := GenS1{F0: (fi12 + (fi11 + 128))}
                    fmt.Println(fi12, fi12, v1)
                    _ = v13
                } else {
                    delete(v4, 2147483647)
                    var v14 uint64 = 0
                    switch "hello" {
                    case "":
                        if (!false) {
                            v15 := GenS1{F0: -615}
                            var v16 uint = 1
                            v14 = v14
                            _ = v15
                            _ = v16
                        } else {
                            fmt.Println(fi11)
                            v5 = v5
                        }
                    case "\n\t":
                        v17 := GenS1{F0: fi12}
                        _ = v17
                    default:
                        if (true && false) {
                            _ = len(v6)
                            v5 = v5
                            if (true && (2147483648 != v5)) {
                                if false {
                                    v14 += v14
                                    v18 := GenS2{F0: 2147483647}
                                    var v19 uint16 = (v1 | 187)
                                    var v20 int = 42
                                    _ = v18
                                    _ = v19
                                    _ = v20
                                }
                                if false {
                                    var v21 string = ""
                                    v22 := GenS2{F0: 2147483647}
                                    v23 := GenS1{F0: 32767}
                                    _ = v21
                                    _ = v22
                                    _ = v23
                                }
                            } else {
                                v5 = (^(v5))
                                if (-2147483648 > 42) {
                                    fmt.Println(fi11)
                                    v1 += (v1 - 255)
                                }
                                v2 -= (v2 + v2)
                            }
                        }
                        var v24 string = "a"
                        v24 = v24
                        _ = v24
                    }
                    _ = v14
                }
                v5 = v5
                v5 -= 4294967295
            } else {
                var v25 uint16 = (v1 % ((v1) | 1))
                v2 = v2
                _ = v25
            }
            _ = fi12
        }
        switch fi11 {
        case 9223372036854775807:
            if true {
                if false {
                    var v26 uint16 = 255
                    var v27 int64 = -1
                    switch fi11 {
                    case 273:
                        var v28 int64 = (v27 & (v27 >> ((307) & 63)))
                        _ = v28
                    case 42:
                        var v29 int32 = 1
                        _ = v29
                    case 4294967295:
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    }
                    _ = v26
                    _ = v27
                }
                fmt.Println(fi11, v2, v5)
            } else {
                for fi30 := 0; fi30 < 6; fi30++ {
                    var v31 int32 = 42
                    _ = fi30
                    _ = v31
                }
                var v32 float64 = 0.1
                v5 += v5
                _ = v32
            }
        case 1:
            if false {
                switch "a" {
                case "hello":
                    if (("a" + "\n\t") != (("" + "\n\t") + "foo bar")) {
                        v33 := GenS2{F0: 1}
                        var v34 uint64 = 42
                        v3[32767] = ("a" + "foo bar")
                        _ = v33
                        _ = v34
                    }
                case "\n\t":
                    if (true && (!(!true))) {
                        v5 = (v5 | v5)
                        v5 += (v5 + 4294967295)
                        for fi35 := 0; fi35 < 13; fi35++ {
                            v1 = v1
                            var v36 bool = (1 != 42)
                            _ = fi35
                            _ = v36
                        }
                    }
                default:
                    fmt.Println(v1, v5, v2)
                }
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                var v37 string = ("foo bar" + ((("hello" + ("a" + "foo bar")) + ("\n\t" + "")) + ((((("a" + (("a" + "") + "a")) + ("a" + "\n\t")) + "a") + (("a" + ("foo bar" + ("hello" + "foo bar"))) + ("" + (("hello" + "") + ("a" + ("" + "\n\t")))))) + "hello")))
                _ = v37
            }
        default:
            _ = len(v4)
        }
        v1 = (v1 << ((823) & 63))
        _ = fi11
    }
    if false {
        if true {
            v38 := GenS1{F0: 128}
            _ = v38
        }
        v2 += float32(3.14159)
        v39 := v3[32768]
        _ = v39
    }
    v5 = uint32(v1)
    fmt.Println(v5, v5, v2)
    fmt.Println(v5, v5)
    var v40 uint16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
