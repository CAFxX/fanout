package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 int32 }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 uint8; F1 int; F2 bool }
type GenS4 struct { F0 string; F1 int; F2 int32 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]int)
    v3 := []uint32{460, 42, 1, 2}
    var v4 int16 = 1
    v5 := make(map[string]uint32)
    var v6 int8 = 1
    v4 += (-(v4))
    var v7 uint = 2
    for fi8 := 0; fi8 < 20; fi8++ {
        if true {
            v6 = (v6 & v6)
            switch -2147483648 {
            case -145:
                if false {
                    v5[""] = 65535
                    v4 -= (v4 << ((v7) & 63))
                } else {
                    if len(v3) > 0 {
                        v3[0] = 65535
                    }
                    var v9 string = "hello"
                    if true {
                        v9 = ((v9 + (v9 + v9)) + v9)
                    }
                    _ = v9
                }
            default:
                var v10 int16 = 1000
                _ = v10
            }
            var v11 uint32 = 2147483648
            _ = v11
        }
        v2[994] = -221
        for fi12 := 0; fi12 < 16; fi12++ {
            _ = len(v5)
            v13 := GenS1{F0: 32767, F1: ((9223372036854775807 <= 9223372036854775807) && (true || (false || (((-2.25 * 0.0) == 1e10) || false)))), F2: -312}
            fmt.Println(v6)
            _ = fi12
            _ = v13
        }
        _ = fi8
    }
    v14 := GenS1{F0: 9223372036854775807, F1: (218 < 255), F2: -2147483648}
    var v15 uint64 = 4294967295
    var v16 bool = true
    if (0 == v4) {
        v17 := GenS1{F0: 9223372036854775807, F1: false, F2: -2147483648}
        v4 = v4
        var v18 uint8 = 42
        _ = v17
        _ = v18
    } else {
        for fi19 := 0; fi19 < 1; fi19++ {
            if ("" == "a") {
                var v20 int16 = (v4 + (v4 / ((v4) | 1)))
                _ = v20
            }
            _ = fi19
        }
        v3 = append(v3, 0)
        v15 = (v15 & 42)
    }
    v21 := GenS1{F0: 1, F1: (261 <= 351), F2: -568}
    if v16 {
        var v22 float32 = 0.1
        v6 = v6
        _ = v22
    }
    v6 = 74
    v7 = 4294967295
    v23 := GenS1{F0: 9223372036854775807, F1: ((-613 == 2147483647) || v16), F2: 42}
    switch ("a" + "a") {
    case "\n\t":
        if (v4 < 0) {
            if (4294967295 >= 32767) {
                if false {
                    fmt.Println(v6, v15, v16)
                } else {
                    switch -288 {
                    case 4294967295:
                        var v24 string = "hello"
                        _ = v24
                    default:
                        for fi25 := 0; fi25 < 9; fi25++ {
                            switch "" {
                            case "hello":
                                switch "" {
                                case "":
                                    for fi26 := 0; fi26 < 10; fi26++ {
                                        v27 := GenS2{F0: 2147483647, F1: 692}
                                        v4 += 744
                                        _ = fi26
                                        _ = v27
                                    }
                                case "hello":
                                    v6 += v6
                                case "a":
                                    for fi28 := 0; fi28 < 2; fi28++ {
                                        if len(v3) > 0 {
                                            v3[0] = 2
                                        }
                                        var v29 uint32 = 1
                                        _ = fi28
                                        _ = v29
                                    }
                                }
                            default:
                                v15 = v15
                            }
                            var v30 int = fi25
                            var v31 string = ""
                            _ = fi25
                            _ = v30
                            _ = v31
                        }
                    }
                    v3 = append(v3, 373)
                }
                var v32 float32 = 1e10
                _ = v32
            }
            v7 = 1
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            fmt.Println(v15)
        }
    case "a":
        v7 = (v7 & v7)
    default:
        if (false || (((("hello" + "\n\t") + "\n\t") >= "\n\t") || v16)) {
            switch "a" {
            case "\n\t":
                v2[207] = 0
            case "a":
                var v33 uint32 = 0
                _ = v33
            case "foo bar":
                if (!false) {
                    for fi34 := 0; fi34 < 5; fi34++ {
                        var v35 int64 = 1
                        _ = fi34
                        _ = v35
                    }
                    v36 := GenS3{F0: 1, F1: 353, F2: v16}
                    _ = v36
                }
            default:
                for fi37 := 0; fi37 < 15; fi37++ {
                    _ = len(v2)
                    if len(v3) > 0 {
                        v3[0] = 374
                    }
                    _ = fi37
                }
            }
        } else {
            fmt.Println(v7)
        }
    }
    v38 := GenS4{F0: "foo bar", F1: 2147483647, F2: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v23
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
