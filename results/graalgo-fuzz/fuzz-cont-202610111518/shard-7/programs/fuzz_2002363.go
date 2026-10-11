package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 uint16 }
func genFn0(p0 int16, p1 int8, p2 uint16) int {
    var v6 int16 = 32767
    if false {
        for fi7 := 0; fi7 < 11; fi7++ {
            var v8 uint64 = 4294967295
            _ = fi7
            _ = v8
        }
        for fi9 := 0; fi9 < 13; fi9++ {
            v10 := []int{(fi9 << ((65535) & 63)), (fi9 & 4294967295), fi9}
            p0 = (p0 % ((v6) | 1))
            _ = fi9
            _ = v10
        }
    }
    switch 32768 {
    case -733:
        v11 := []int{480, -48}
        _ = v11
    case -2147483648:
        v12 := GenS1{F0: "", F1: -2147483648, F2: p2}
        _ = v12
    case 0:
        v13 := GenS1{F0: ("foo bar" + "\n\t"), F1: 42, F2: 65535}
        _ = v13
    }
    fmt.Println(v6, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -451
}

func main() {
    var v1 uint64 = 2
    var v2 int64 = -9223372036854775808
    v3 := []string{"\n\t", "foo bar", "foo bar", ("a" + "")}
    v4 := make(map[uint32]int)
    var v5 uint16 = 255
    switch ((("a" + "a") + "hello") + "hello") {
    case "\n\t":
        var v14 uint32 = 42
        _ = v14
    }
    switch "foo bar" {
    case "":
        for fi15 := 0; fi15 < 10; fi15++ {
            var v16 bool = (!false)
            if true {
                v2 += 42
                switch genFn0(-547, 1, (v5 - (v5 * 0))) {
                case -9223372036854775808:
                    var v17 int16 = int16(v5)
                    _ = v17
                case 127:
                    v16 = v16
                }
                _ = len(v3)
            }
            v18 := GenS1{F0: "\n\t", F1: 1, F2: (v5 & v5)}
            _ = fi15
            _ = v16
            _ = v18
        }
    case "foo bar":
        _ = len(v4)
    default:
        switch "" {
        case "hello":
            if true {
                delete(v4, 578)
            }
        case "\n\t":
            fmt.Println(v1, v5, v1)
        }
    }
    fmt.Println(v2, v1)
    var v19 int32 = -881
    if false {
        for fi20 := 0; fi20 < 12; fi20++ {
            fmt.Println(v2)
            _ = fi20
        }
        _ = len(v3)
        v21 := GenS1{F0: ("hello" + ("" + "\n\t")), F1: (v19 % (((v19 + v19)) | 1)), F2: (v5 - (v5 % ((v5) | 1)))}
        _ = v21
    } else {
        for fi22 := 0; fi22 < 18; fi22++ {
            var v23 string = "\n\t"
            _ = fi22
            _ = v23
        }
        var v24 int16 = 1
        _ = v24
    }
    v2 += 9223372036854775807
    v3 = append(v3, "")
    fmt.Println(v5, v19)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
