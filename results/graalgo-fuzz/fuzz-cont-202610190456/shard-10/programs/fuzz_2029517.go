package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 int32; F1 uint32; F2 int64 }
type GenS3 struct { F0 uint; F1 uint }
type GenS4 struct { F0 uint; F1 int64 }

func main() {
    var v1 uint8 = 127
    v2 := []string{(("foo bar" + ("hello" + "hello")) + "\n\t")}
    v3 := []uint32{}
    v4 := make(map[int]int)
    for fi5 := 0; fi5 < 17; fi5++ {
        var v6 string = "foo bar"
        _ = fi5
        _ = v6
    }
    switch 2147483647 {
    case -1:
        v1 = v1
    }
    if true {
        var v7 int32 = 1
        _ = v7
    } else {
        v8 := GenS1{F0: 0, F1: -128}
        if false {
            switch -971 {
            case 0:
                var v9 uint16 = 265
                _ = v9
            case 4294967295:
                var v10 int16 = 32767
                _ = v10
            }
            switch "hello" {
            case "":
                if true {
                    v11 := GenS1{F0: 1, F1: -128}
                    v1 = (v1 << ((42) & 63))
                    var v12 uint8 = 1
                    _ = v11
                    _ = v12
                }
            case "a":
                if false {
                    for fi13 := 0; fi13 < 18; fi13++ {
                        v3 = append(v3, uint32(v1))
                        fmt.Println(fi13, v1)
                        _ = fi13
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                } else {
                    if (411 < 2) {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    } else {
                        var v14 string = ""
                        v15 := GenS2{F0: 222, F1: uint32(v1), F2: 42}
                        var v16 uint16 = 32767
                        _ = v14
                        _ = v15
                        _ = v16
                    }
                }
            }
            for fi17 := 0; fi17 < 18; fi17++ {
                var v18 int16 = 1
                if len(v3) > 0 {
                    v3[0] = 2147483648
                }
                if len(v2) > 0 {
                    v2[0] = "\n\t"
                }
                _ = fi17
                _ = v18
            }
        }
        if (1 >= -1) {
            v1 = v1
            v4[1] = 2147483647
            switch 4294967295 {
            case 2147483647:
                var v19 uint = 42
                _ = v19
            case 1:
                v1 = v1
            }
        }
        _ = v8
    }
    v1 -= v1
    var v20 uint8 = 2
    switch 4294967295 {
    case 128:
        if (!true) {
            switch 0 {
            case -9223372036854775808:
                v1 = v20
            case -1:
                fmt.Println(v1)
            }
        }
    case -2147483648:
        switch 762 {
        case 42:
            v21 := GenS2{F0: 1, F1: 2147483648, F2: -9223372036854775808}
            _ = v21
        }
    case 346:
        v22 := GenS1{F0: 42, F1: 100}
        _ = v22
    }
    v23 := GenS3{F0: 723, F1: 65535}
    var v24 int32 = -1
    for fi25 := 0; fi25 < 16; fi25++ {
        if true {
            var v26 bool = true
            _ = v26
        }
        v4[fi25] = fi25
        _ = fi25
    }
    switch -609 {
    case -72:
        var v27 bool = false
        _ = v27
    case -2147483648:
        v2 = append(v2, ("a" + ("\n\t" + "\n\t")))
    case 684:
        v20 += 127
    default:
        fmt.Println(v1, v20, v1)
    }
    for fi28 := 0; fi28 < 9; fi28++ {
        v24 += 0
        v29 := GenS3{F0: 18446744073709551615, F1: 844}
        _ = fi28
        _ = v29
    }
    v30 := GenS4{F0: 255, F1: -1}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v3
    _ = v30
    _ = v4
}
