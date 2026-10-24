package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 int8; F1 bool }
func genFn0(p0 uint8, p1 uint64, p2 int8) int32 {
    if (true && ("" > "")) {
        v4 := []int{32767, -291}
        _ = v4
    }
    var v5 bool = (false && (1 < 2147483647))
    fmt.Println(p0, p0)
    v6 := []int{32767}
    _ = len(v6)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 0
}
func genFn1(p0 bool, p1 int16) bool {
    p0 = (9223372036854775807 >= -9223372036854775808)
    switch ("" + "") {
    case "hello":
        var v7 uint64 = 713
        _ = v7
    case "\n\t":
        fmt.Println(p1, p0)
    }
    var v8 uint = 2
    _ = p0
    _ = p1
    _ = v8
    return p0
}

func main() {
    var v1 uint8 = 255
    v2 := make(map[string]string)
    v3 := []string{"a", ("hello" + "hello")}
    v3 = append(v3, "")
    var v9 uint = 847
    var v10 uint = 1
    switch "foo bar" {
    case "a":
        if true {
            v10 -= (v10 >> ((415) & 63))
            if false {
                v11 := GenS1{F0: 2, F1: -1}
                _ = v11
            }
            v9 += (v10 % ((uint(v1)) | 1))
        }
    }
    var v12 int = -847
    switch "" {
    case "foo bar":
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
    case "":
        v10 = v10
    }
    v13 := GenS1{F0: 1, F1: -1}
    if false {
        if true {
            for fi14 := 0; fi14 < 11; fi14++ {
                var v15 uint8 = v1
                switch ("hello" + ("\n\t" + "hello")) {
                case "\n\t":
                    fmt.Println(v9)
                }
                _ = fi14
                _ = v15
            }
            delete(v2, (("a" + "") + "foo bar"))
        } else {
            fmt.Println(v10)
        }
        fmt.Println(v10, v10)
    } else {
        for fi16 := 0; fi16 < 10; fi16++ {
            for fi17 := 0; fi17 < 8; fi17++ {
                v1 = (v1 | 0)
                _ = fi17
            }
            _ = fi16
        }
    }
    v9 = (v10 | 149)
    v2["\n\t"] = "hello"
    var v18 uint16 = 255
    for fi19 := 0; fi19 < 5; fi19++ {
        switch fi19 {
        case -676:
            switch (v12 * v12) {
            case 0:
                v9 += 157
            case 808:
                if false {
                    v1 = (v1 ^ v1)
                    var v20 float64 = (0.1 - (1.5 / 0.1))
                    v18 = (v18 % (((v18 * v18)) | 1))
                    _ = v20
                }
            default:
                switch "" {
                case "\n\t":
                    switch "hello" {
                    case "":
                        v10 = 4294967295
                    default:
                        v1 = (v1 % (((v1 & v1)) | 1))
                    }
                case "a":
                    v21 := GenS1{F0: 595, F1: 1}
                    _ = v21
                }
            }
        case 98:
            v22 := GenS2{F0: 100, F1: (127 > 100)}
            _ = v22
        case 0:
            v18 = (v18 - v18)
        default:
            v23 := GenS2{F0: 1, F1: true}
            _ = v23
        }
        _ = fi19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v9
}
