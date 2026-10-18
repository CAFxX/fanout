package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 int8 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 int32; F1 uint64 }

func main() {
    v1 := make(map[int]uint32)
    v2 := []string{""}
    var v3 int32 = -196
    var v4 int32 = v3
    v2 = append(v2, ((("\n\t" + "hello") + "") + "a"))
    v5 := GenS1{F0: -467, F1: ((((("foo bar" + "hello") + ("a" + "\n\t")) + "") + ("a" + "a")) + "foo bar"), F2: -113}
    var v6 uint8 = 255
    fmt.Println(v3, v4, v6)
    if (!true) {
        if false {
            v7 := GenS1{F0: 32768, F1: "a", F2: 0}
            var v8 uint32 = 728
            var v9 uint32 = (v8 | (v8 / (((v8 | v8)) | 1)))
            _ = v7
            _ = v8
            _ = v9
        }
        if true {
            v10 := GenS1{F0: -458, F1: "a", F2: -128}
            _ = v10
        } else {
            var v11 uint16 = 65535
            var v12 float32 = (3.14159 + 0.1)
            _ = v11
            _ = v12
        }
    }
    v13 := GenS2{F0: "", F1: 42}
    v14 := GenS1{F0: -110, F1: "a", F2: -128}
    for fi15 := 0; fi15 < 13; fi15++ {
        var v16 float32 = 3.14159
        _ = fi15
        _ = v16
    }
    switch 32767 {
    case -1:
        for fi17 := 0; fi17 < 12; fi17++ {
            v18 := GenS1{F0: 42, F1: ("hello" + ""), F2: 127}
            v4 = v4
            if len(v2) > 0 {
                v2[0] = ""
            }
            _ = fi17
            _ = v18
        }
    case -548:
        fmt.Println(v6, v4, v4)
    default:
        v3 -= -16
    }
    var v19 bool = (1 <= 1)
    var v20 int = 2147483647
    var v21 int32 = (v3 << ((101) & 63))
    v4 = (v21 - -2147483648)
    var v22 int64 = -384
    if v19 {
        v21 -= (v21 % ((v3) | 1))
        switch (v20 + (v20 - 4294967295)) {
        case -2147483648:
            switch "foo bar" {
            case "":
                v3 = (v3 * v3)
            case "a":
                v21 = -1
            case "hello":
                fmt.Println(v19)
            default:
                v23 := GenS3{F0: 0, F1: 2}
                _ = v23
            }
        case -732:
            switch ("" + "foo bar") {
            case "":
                for fi24 := 0; fi24 < 14; fi24++ {
                    for fi25 := 0; fi25 < 5; fi25++ {
                        var v26 uint16 = 255
                        _ = fi25
                        _ = v26
                    }
                    switch "a" {
                    case "":
                        var v27 uint = 898
                        _ = v27
                    case "a":
                        if v19 {
                            v20 += (v20 * (^(fi24)))
                        }
                    case "foo bar":
                        if v19 {
                            switch 9223372036854775807 {
                            case 127:
                                v28 := GenS3{F0: (v4 >> ((135) & 63)), F1: 1}
                                _ = v28
                            case 145:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            default:
                                var v29 int8 = -128
                                _ = v29
                            }
                        } else {
                            v4 += (-(v4))
                            var v30 int = v20
                            var v31 uint16 = 255
                            _ = v30
                            _ = v31
                        }
                    default:
                        fmt.Println(v6, v22)
                    }
                    _ = fi24
                }
            case "hello":
                v22 += (v22 + v22)
            case "a":
                _ = len(v2)
            }
        }
    } else {
        v32 := GenS3{F0: v4, F1: 29}
        if v19 {
            v1[int(v21)] = 45
            v33 := GenS2{F0: ((("" + "") + "foo bar") + "a"), F1: 65535}
            _ = v33
        }
        for fi34 := 0; fi34 < 16; fi34++ {
            v35 := GenS1{F0: v20, F1: "a", F2: -128}
            _ = fi34
            _ = v35
        }
        _ = v32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
