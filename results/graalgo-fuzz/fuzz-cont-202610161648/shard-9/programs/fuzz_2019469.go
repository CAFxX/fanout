package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 int32; F1 int8; F2 int64 }
func genFn0(p0 int, p1 uint64, p2 string) string {
    var v7 uint64 = (p1 % (((p1 - p1)) | 1))
    p1 += (v7 + 4294967295)
    v8 := []int{-2147483648}
    var v9 int16 = 1000
    var v10 uint = 945
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return ""
}
func genFn1(p0 bool, p1 uint16) uint32 {
    for fi11 := 0; fi11 < 9; fi11++ {
        v12 := GenS1{F0: ("a" + "foo bar"), F1: 1}
        if ((!false) && p0) {
            switch "foo bar" {
            case "":
                switch 9223372036854775807 {
                case 4294967295:
                    for fi13 := 0; fi13 < 1; fi13++ {
                        var v14 uint64 = 601
                        v15 := []int{}
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                case 0:
                    v16 := GenS1{F0: "hello", F1: 136}
                    _ = v16
                default:
                    for fi17 := 0; fi17 < 17; fi17++ {
                        p0 = p0
                        fmt.Println(fi17, fi17, p0)
                        var v18 int16 = -1
                        _ = fi17
                        _ = v18
                    }
                }
            case "a":
                v19 := GenS1{F0: "hello", F1: 18446744073709551615}
                _ = v19
            case "foo bar":
                v20 := GenS1{F0: "\n\t", F1: 0}
                _ = v20
            }
        }
        _ = fi11
        _ = v12
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    return 2147483647
}

func main() {
    var v1 uint8 = 0
    var v2 int8 = -1
    var v3 uint64 = 933
    v4 := []string{"\n\t", "", ("" + "a"), ("a" + "hello")}
    v5 := make(map[int]bool)
    var v6 uint16 = 255
    fmt.Println(v6, v3)
    if ((170 != 65535) && false) {
        var v21 bool = ("" != "\n\t")
        _ = v21
    } else {
        _ = len(v5)
    }
    if false {
        for fi22 := 0; fi22 < 11; fi22++ {
            var v23 uint32 = 2147483647
            _ = fi22
            _ = v23
        }
        for fi24 := 0; fi24 < 7; fi24++ {
            switch -9223372036854775808 {
            case -9223372036854775808:
                var v25 float32 = (1.5 - 1.5)
                _ = v25
            case 876:
                v2 -= -128
            case 32767:
                _ = len(v5)
            }
            _ = fi24
        }
    }
    v6 = 2
    fmt.Println(v1, v2, v1)
    if (!false) {
        if false {
            switch "" {
            case "a":
                fmt.Println(v2, v1)
            default:
                v3 = 18446744073709551615
            }
            for fi26 := 0; fi26 < 19; fi26++ {
                v6 = 0
                v3 -= (v3 * (v3 - (v3 * v3)))
                _ = fi26
            }
            v2 = (-(v2))
        }
        var v27 uint8 = uint8(v6)
        _ = v27
    }
    v6 = (v6 + 1)
    var v28 uint = 65535
    for fi29 := 0; fi29 < 6; fi29++ {
        var v30 uint16 = 80
        for fi31 := 0; fi31 < 9; fi31++ {
            if true {
                var v32 int16 = 32767
                delete(v5, fi29)
                var v33 int32 = -20
                _ = v32
                _ = v33
            }
            _ = fi31
        }
        _ = fi29
        _ = v30
    }
    v3 = v3
    fmt.Println(v2, v28)
    var v34 string = ("" + "foo bar")
    if (-32768 != -32768) {
        v35 := GenS1{F0: "a", F1: 255}
        if false {
            if (-538 == -32768) {
                var v36 uint64 = v3
                fmt.Println(v3, v6)
                for fi37 := 0; fi37 < 6; fi37++ {
                    v38 := GenS2{F0: 192, F1: (-(v2)), F2: 0}
                    fmt.Println(v2, v2)
                    _ = fi37
                    _ = v38
                }
                _ = v36
            }
            var v39 uint8 = v1
            switch 127 {
            case 582:
                var v40 int = 127
                _ = v40
            case 9223372036854775807:
                _ = len(v5)
            case 4294967295:
                fmt.Println(v3)
            default:
                delete(v5, 470)
            }
            _ = v39
        }
        v34 = v34
        _ = v35
    }
    v2 -= (v2 << (((v28 | 2)) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
