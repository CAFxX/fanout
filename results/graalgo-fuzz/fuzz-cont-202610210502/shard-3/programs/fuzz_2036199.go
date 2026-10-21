package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int32; F1 uint64 }
type GenS5 struct { F0 uint16; F1 int16; F2 uint16 }

func main() {
    v1 := make(map[uint32]int)
    var v2 int16 = 32767
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]bool)
    v5 := []string{("foo bar" + ("a" + "a")), "\n\t"}
    v6 := []string{"foo bar", ("hello" + "\n\t"), ("hello" + ("hello" + "foo bar")), "a"}
    for fi7 := 0; fi7 < 7; fi7++ {
        v2 = (v2 ^ v2)
        switch fi7 {
        case 4294967295:
            if true {
                v2 = 1
                v4[65535] = false
                fmt.Println(v2, fi7)
            }
        case 835:
            var v8 string = "hello"
            _ = v8
        case 916:
            var v9 uint32 = 2147483647
            _ = v9
        }
        fmt.Println(v2, fi7)
        _ = fi7
    }
    for fi10 := 0; fi10 < 1; fi10++ {
        _ = len(v5)
        fmt.Println(fi10)
        if false {
            if (-2147483648 <= -135) {
                fmt.Println(v2, v2)
                v2 = -32768
            }
            v2 -= 1000
        } else {
            fmt.Println(v2)
            var v11 int64 = 0
            _ = v11
        }
        _ = fi10
    }
    if len(v6) > 0 {
        v6[0] = (("\n\t" + (("hello" + (("hello" + "a") + "foo bar")) + "a")) + (("hello" + "a") + "\n\t"))
    }
    if (32767 > 255) {
        v12 := GenS1{F0: ("foo bar" + "a"), F1: 2147483648}
        v13 := GenS2{F0: 42, F1: -1}
        switch "a" {
        case "foo bar":
            if (true && (42 >= 42)) {
                v14 := GenS2{F0: 4294967295, F1: 37}
                _ = v14
            }
        case "hello":
            switch ("" + ("hello" + "")) {
            case "foo bar":
                v2 = (v2 * (v2 | (-(v2))))
            default:
                var v15 int32 = 2147483647
                _ = v15
            }
        case "a":
            for fi16 := 0; fi16 < 5; fi16++ {
                if (v2 > v2) {
                    v2 = v2
                    v17 := GenS3{F0: 4294967295}
                    if len(v6) > 0 {
                        v6[0] = ""
                    }
                    _ = v17
                } else {
                    _ = len(v6)
                }
                v18 := GenS4{F0: -113, F1: 1}
                _ = fi16
                _ = v18
            }
        default:
            var v19 float64 = 0.1
            _ = v19
        }
        _ = v12
        _ = v13
    }
    var v20 uint64 = 330
    fmt.Println(v2)
    var v21 int32 = 2147483647
    for fi22 := 0; fi22 < 17; fi22++ {
        v21 += (v21 % (((v21 | v21)) | 1))
        fmt.Println(v21)
        if len(v6) > 0 {
            v6[0] = ("" + "foo bar")
        }
        _ = fi22
    }
    fmt.Println(v20, v20)
    var v23 int32 = v21
    v24 := GenS3{F0: 713}
    for fi25 := 0; fi25 < 4; fi25++ {
        v26 := GenS5{F0: 320, F1: v2, F2: uint16(v20)}
        _ = fi25
        _ = v26
    }
    v6 = append(v6, "\n\t")
    var v27 uint = 675
    for fi28 := 0; fi28 < 16; fi28++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi28
    }
    v5 = append(v5, "\n\t")
    var v29 bool = false
    fmt.Println(v27)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v24
    _ = v27
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
