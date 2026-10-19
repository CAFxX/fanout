package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }

func main() {
    var v1 uint8 = 127
    var v2 int16 = 0
    v3 := make(map[uint32]string)
    v4 := v3[345]
    for fi5 := 0; fi5 < 2; fi5++ {
        fmt.Println(v2, v2, v4)
        if false {
            var v6 uint = 0
            v7 := []int{fi5, fi5, (fi5 ^ fi5)}
            _ = v6
            _ = v7
        } else {
            v4 = ("hello" + ((v4 + v4) + "hello"))
            v4 = (("hello" + "foo bar") + "a")
        }
        fmt.Println(fi5, v2)
        _ = fi5
    }
    v8 := []int{32767, 42, -736, -46}
    for fi9 := 0; fi9 < 12; fi9++ {
        v4 = (("a" + ("hello" + (v4 + ("a" + "")))) + (v4 + ""))
        if true {
            if (255 < 2) {
                for fi10 := 0; fi10 < 8; fi10++ {
                    v11 := GenS1{F0: 42, F1: -9223372036854775808}
                    _ = fi10
                    _ = v11
                }
                switch ((((v4 + ((v4 + (v4 + ((v4 + v4) + (("" + v4) + ((v4 + (v4 + v4)) + ("\n\t" + "\n\t")))))) + v4)) + "") + ("hello" + v4)) + "hello") {
                case "hello":
                    v12 := GenS1{F0: 184, F1: -1}
                    _ = v12
                case "foo bar":
                    switch v4 {
                    case "\n\t":
                        if true {
                            switch fi9 {
                            case -2147483648:
                                v13 := GenS1{F0: uint64(v1), F1: -9223372036854775808}
                                _ = v13
                            }
                            if (39 >= v1) {
                                var v14 float32 = (3.14159 - 3.14159)
                                _ = v14
                            }
                        } else {
                            for fi15 := 0; fi15 < 3; fi15++ {
                                var v16 int = (-(fi9))
                                _ = fi15
                                _ = v16
                            }
                        }
                    case "foo bar":
                        v17 := GenS1{F0: 42, F1: -234}
                        _ = v17
                    default:
                        if (1 == v1) {
                            for fi18 := 0; fi18 < 16; fi18++ {
                                var v19 int8 = 1
                                v20 := GenS1{F0: 1, F1: 794}
                                _ = fi18
                                _ = v19
                                _ = v20
                            }
                        } else {
                            var v21 uint = 42
                            v4 = ("a" + v4)
                            _ = len(v8)
                            _ = v21
                        }
                    }
                default:
                    var v22 uint = 0
                    _ = v22
                }
                v23 := GenS1{F0: 4294967295, F1: -9223372036854775808}
                _ = v23
            } else {
                v2 -= -32768
            }
            for fi24 := 0; fi24 < 12; fi24++ {
                switch v4 {
                case "a":
                    var v25 uint32 = 2
                    _ = v25
                }
                v2 = (v2 >> ((4294967295) & 63))
                _ = fi24
            }
        } else {
            v8 = append(v8, 912)
            v26 := GenS1{F0: 4294967295, F1: 88}
            _ = v26
        }
        var v27 int16 = (v2 * -32768)
        _ = fi9
        _ = v27
    }
    fmt.Println(v4, v4)
    if (-1 > 1) {
        v3[999] = ""
    }
    switch ("foo bar" + v4) {
    case "foo bar":
        for fi28 := 0; fi28 < 6; fi28++ {
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            _ = fi28
        }
    default:
        v29 := GenS1{F0: 910, F1: -527}
        _ = v29
    }
    var v30 int = 42
    if len(v8) > 0 {
        v8[0] = (v30 >> ((18446744073709551615) & 63))
    }
    var v31 int8 = 100
    v8 = append(v8, int(v2))
    var v32 uint8 = v1
    for fi33 := 0; fi33 < 2; fi33++ {
        for fi34 := 0; fi34 < 6; fi34++ {
            fmt.Println(v2, v1)
            _ = fi34
        }
        if (271 >= 2147483648) {
            var v35 uint16 = 0
            _ = v35
        }
        delete(v3, 2147483647)
        _ = fi33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v8
}
