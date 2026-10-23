package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool; F1 int16; F2 int32 }

func main() {
    v1 := []uint32{701, 574, 316, 804}
    var v2 uint16 = 42
    v3 := make(map[string]string)
    var v4 uint64 = 0
    var v5 int8 = 109
    if true {
        var v6 uint = 42
        v2 += uint16(v6)
        _ = v6
    }
    switch -774 {
    case 32768:
        v7 := GenS1{F0: false, F1: -2147483648, F2: 812}
        _ = v7
    case -36:
        v4 = v4
    case -397:
        if false {
            switch ("a" + "a") {
            case "\n\t":
                v8 := v3["hello"]
                _ = v8
            case "":
                _ = len(v1)
            case "a":
                for fi9 := 0; fi9 < 20; fi9++ {
                    var v10 uint = 553
                    if (true || (!(65535 >= 1))) {
                        var v11 int32 = 0
                        _ = v11
                    } else {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    switch "" {
                    case "\n\t":
                        if ("" > "a") {
                            v12 := GenS1{F0: (2 == 128), F1: 0, F2: 2147483648}
                            _ = v12
                        }
                    }
                    _ = fi9
                    _ = v10
                }
            }
        } else {
            var v13 uint8 = 2
            if len(v1) > 0 {
                v1[0] = 2147483648
            }
            _ = v13
        }
    }
    var v14 int32 = 0
    fmt.Println(v4)
    v14 -= (v14 % ((v14) | 1))
    v15 := GenS2{F0: ""}
    fmt.Println(v2)
    if false {
        _ = len(v1)
        v2 += 65535
        fmt.Println(v4, v5, v5)
    } else {
        v5 += v5
    }
    for fi16 := 0; fi16 < 10; fi16++ {
        delete(v3, "foo bar")
        v2 += (v2 | 908)
        _ = fi16
    }
    var v17 uint32 = 0
    v18 := GenS1{F0: (!false), F1: 795, F2: (v17 + 42)}
    fmt.Println(v17, v2)
    for fi19 := 0; fi19 < 13; fi19++ {
        v2 -= 0
        if true {
            var v20 string = ("a" + "a")
            var v21 uint32 = (v17 * 1)
            _ = v20
            _ = v21
        } else {
            delete(v3, ("hello" + "\n\t"))
        }
        _ = fi19
    }
    if false {
        v22 := GenS3{F0: (1 > 687), F1: 0, F2: (v14 * (v14 | int32(v5)))}
        fmt.Println(v2, v14, v5)
        _ = v22
    } else {
        v5 = v5
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
