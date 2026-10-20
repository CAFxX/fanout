package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 int; F2 string }
type GenS3 struct { F0 int8 }
func genFn0(p0 int64, p1 uint16) int8 {
    var v4 int8 = -65
    if true {
        switch int(p0) {
        case 127:
            switch 717 {
            case -2147483648:
                v5 := []int{2147483647, 4294967295, 4294967295}
                _ = v5
            default:
                var v6 float64 = (0.1 / (1.5 / 1.0))
                _ = v6
            }
        case -714:
            switch "foo bar" {
            case "foo bar":
                var v7 int = 9223372036854775807
                _ = v7
            }
        case 901:
            v8 := []int{4294967295, 32767, 2147483647}
            _ = v8
        }
        v4 += (v4 - (v4 >> ((18446744073709551615) & 63)))
        fmt.Println(p0, p1, p1)
    } else {
        if true {
            for fi9 := 0; fi9 < 18; fi9++ {
                if true {
                    var v10 uint8 = 42
                    var v11 int8 = (v4 / ((v4) | 1))
                    v11 = (v11 << ((408) & 63))
                    _ = v10
                    _ = v11
                }
                v4 = (v4 << ((756) & 63))
                if (true || true) {
                    fmt.Println(p0)
                    p1 += p1
                    var v12 uint32 = 680
                    _ = v12
                }
                _ = fi9
            }
            v13 := GenS1{F0: 0}
            _ = v13
        } else {
            var v14 int8 = (v4 + v4)
            v14 += (^(v14))
            _ = v14
        }
        v15 := GenS2{F0: (-(v4)), F1: -114, F2: "foo bar"}
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v4
    return (v4 % (((v4 * -1)) | 1))
}

func main() {
    var v1 float64 = 1e10
    var v2 float32 = (((3.14159 / float32((v1 * v1))) + (1e10 / (0.1 + 0.0))) - 1.5)
    var v3 uint = 255
    v1 = (v1 * v1)
    v2 -= float32(v1)
    switch 32767 {
    case 127:
        if false {
            v2 = v2
        }
    }
    for fi16 := 0; fi16 < 3; fi16++ {
        var v17 uint32 = 808
        for fi18 := 0; fi18 < 5; fi18++ {
            fmt.Println(v2, v3, v1)
            v19 := GenS3{F0: -94}
            v1 -= (v1 * -2.25)
            _ = fi18
            _ = v19
        }
        v20 := GenS2{F0: 1, F1: fi16, F2: ("" + "foo bar")}
        _ = fi16
        _ = v17
        _ = v20
    }
    fmt.Println(v2, v2)
    v21 := []int{-622, 127, -1, 1}
    v2 -= 1.5
    if len(v21) >= 2 {
        v21 = v21[1:len(v21)]
    }
    for fi22 := 0; fi22 < 2; fi22++ {
        if true {
            if len(v21) >= 2 {
                v21 = v21[1:len(v21)]
            }
        }
        var v23 string = "hello"
        switch fi22 {
        case 4294967295:
            for fi24 := 0; fi24 < 19; fi24++ {
                v1 -= v1
                v2 -= v2
                _ = fi24
            }
        case 230:
            v21 = append(v21, (fi22 * (-(fi22))))
        }
        _ = fi22
        _ = v23
    }
    for fi25 := 0; fi25 < 3; fi25++ {
        v1 += 1e10
        _ = fi25
    }
    if true {
        var v26 uint = v3
        _ = v26
    }
    fmt.Println(v3, v3, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
}
