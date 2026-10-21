package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 int32 }
type GenS2 struct { F0 int; F1 uint8; F2 int64 }
type GenS3 struct { F0 int32; F1 uint8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint; F1 uint16 }
func genFn0(p0 uint8, p1 int8) uint16 {
    if (("a" <= (("hello" + "a") + "\n\t")) && (!true)) {
        fmt.Println(p1)
        var v4 uint = 65535
        _ = v4
    } else {
        for fi5 := 0; fi5 < 7; fi5++ {
            var v6 string = ("hello" + ("hello" + "\n\t"))
            var v7 string = v6
            _ = fi5
            _ = v6
            _ = v7
        }
        v8 := GenS1{F0: "a", F1: 42, F2: -1}
        v9 := GenS2{F0: 9223372036854775807, F1: p0, F2: int64(p1)}
        _ = v8
        _ = v9
    }
    if (((807 > 42) || false) || (1e10 != float32(-2.25))) {
        var v10 uint64 = 808
        switch "hello" {
        case "a":
            fmt.Println(v10, v10)
        case "\n\t":
            v11 := []int{128, 0}
            _ = v11
        default:
            p0 = 199
        }
        switch "foo bar" {
        case "a":
            v12 := GenS1{F0: "", F1: 912, F2: -1}
            _ = v12
        }
        _ = v10
    } else {
        var v13 int8 = (^(p1))
        v14 := GenS2{F0: 42, F1: 255, F2: 1}
        for fi15 := 0; fi15 < 16; fi15++ {
            p0 += 66
            switch 32767 {
            case -1:
                if true {
                    v16 := []int{}
                    fmt.Println(p0, v13)
                    _ = v16
                }
            case -3:
                v17 := GenS1{F0: ("foo bar" + ("foo bar" + (("\n\t" + "\n\t") + "hello"))), F1: 32767, F2: 177}
                _ = v17
            default:
                var v18 uint32 = 0
                _ = v18
            }
            var v19 int16 = 1000
            _ = fi15
            _ = v19
        }
        _ = v13
        _ = v14
    }
    v20 := []int{32767, 127, 4294967295}
    var v21 int64 = -525
    switch 127 {
    case -581:
        var v22 string = (("foo bar" + "a") + ("foo bar" + ""))
        _ = v22
    case 127:
        var v23 int64 = (-(v21))
        _ = v23
    case -9223372036854775808:
        fmt.Println(p0, p0)
    default:
        if (549 > 65535) {
            v24 := GenS3{F0: 0, F1: (p0 << ((810) & 63))}
            var v25 int16 = 1
            var v26 uint32 = 2
            _ = v24
            _ = v25
            _ = v26
        }
    }
    _ = p0
    _ = p1
    _ = v20
    _ = v21
    return 255
}

func main() {
    var v1 uint32 = 2147483647
    v2 := make(map[int]uint32)
    v3 := []int{629, 32768, 382}
    v1 -= v1
    v27 := GenS3{F0: 2147483647, F1: 2}
    switch 9223372036854775807 {
    case -2147483648:
        v1 += v1
    case 128:
        var v28 bool = false
        _ = v28
    case 0:
        v29 := GenS3{F0: -1, F1: 255}
        _ = v29
    }
    switch "" {
    case "hello":
        v30 := v2[42]
        _ = v30
    case "":
        v1 = (^(v1))
    case "a":
        var v31 float64 = 0.0
        _ = v31
    }
    var v32 uint = 2
    var v33 int32 = -2147483648
    for fi34 := 0; fi34 < 12; fi34++ {
        _ = len(v2)
        _ = fi34
    }
    var v35 int = 9223372036854775807
    v36 := GenS4{F0: 960}
    v33 = 970
    if (v1 > 2147483647) {
        if false {
            v32 = (^(v32))
            if len(v3) > 0 {
                v3[0] = int(v33)
            }
        }
        switch (("hello" + "a") + "\n\t") {
        case "a":
            for fi37 := 0; fi37 < 11; fi37++ {
                fmt.Println(fi37, v1, v33)
                fmt.Println(v1, v33, v35)
                if (!true) {
                    var v38 uint = (^(v32))
                    _ = v38
                }
                _ = fi37
            }
        default:
            v39 := GenS5{F0: (v32 - 2), F1: 1}
            _ = v39
        }
        var v40 uint16 = 601
        _ = v40
    }
    v33 += v33
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v32
    _ = v33
    _ = v35
    _ = v36
}
